package main

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/services/cache/common"

	"laatoo/sdk/log"
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisCacheFactory struct {
}

const (
	CONF_REDISCACHE_NAME        = "redis_cache"
	CONF_REDIS_CONNECTIONSTRING = "server"
	CONF_REDIS_DATABASE         = "db"
)

func Manifest() []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_REDISCACHE_NAME, Object: RedisCacheFactory{}}}
}

//Create the services configured for factory.
func (mf *RedisCacheFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	return &RedisCacheService{name: name}, nil
}

func (ds *RedisCacheFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//The services start serving when this method is called
func (ds *RedisCacheFactory) Start(ctx core.ServerContext) error {
	return nil
}

type RedisCacheService struct {
	core.Service
	connectionstring string
	database         string
	conn             redis.Conn
	pool             *redis.Pool
	name             string
	cacheEncoder     *common.CacheEncoder
}

func (svc *RedisCacheService) Delete(ctx core.RequestContext, bucket string, key string) error {
	conn := svc.pool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", common.GetCacheKey(bucket, key))
	if err != nil {
		return err
	}
	return nil
}

func (svc *RedisCacheService) PutObject(ctx core.RequestContext, bucket string, key string, val interface{}) error {
	b, err := svc.cacheEncoder.Encode(val)
	if err != nil {
		return err
	}
	conn := svc.pool.Get()
	defer conn.Close()
	_, err = conn.Do("SET", common.GetCacheKey(bucket, key), b)
	if err != nil {
		return err
	}
	conn.Flush()
	return nil
}

func (svc *RedisCacheService) PutObjects(ctx core.RequestContext, bucket string, vals map[string]interface{}) error {
	var args []interface{}
	for k, v := range vals {
		b, err := svc.cacheEncoder.Encode(v)
		if err != nil {
			return err
		}
		args = append(args, common.GetCacheKey(bucket, k), b)
	}
	conn := svc.pool.Get()
	defer conn.Close()
	_, err := conn.Do("MSET", args...)
	if err != nil {
		return err
	}
	conn.Flush()
	return nil
}

func (svc *RedisCacheService) Get(ctx core.RequestContext, bucket string, key string) (interface{}, bool) {
	conn := svc.pool.Get()
	defer conn.Close()
	cval, err := conn.Do("GET", common.GetCacheKey(bucket, key))
	if err != nil {
		return nil, false
	}
	return cval, true
}

func (svc *RedisCacheService) GetObject(ctx core.RequestContext, bucket string, key string, objectType string) (interface{}, bool) {
	conn := svc.pool.Get()
	defer conn.Close()
	k, err := conn.Do("GET", common.GetCacheKey(bucket, key))
	if err != nil {
		return nil, false
	}
	obj, err := ctx.ServerContext().CreateObject(objectType)
	if err != nil {
		return nil, false
	}
	err = svc.cacheEncoder.Decode(k.([]byte), obj)
	if err != nil {
		return nil, false
	}
	return obj, true
}

func (svc *RedisCacheService) GetMulti(ctx core.RequestContext, bucket string, keys []string) map[string]interface{} {
	var args []interface{}
	for _, k := range keys {
		args = append(args, common.GetCacheKey(bucket, k))
	}
	retval := make(map[string]interface{})
	conn := svc.pool.Get()
	defer conn.Close()
	cvals, err := redis.Values(conn.Do("MGET", args...))
	if err != nil {
		return retval
	}
	for index, val := range cvals {
		key := keys[index]
		if val == nil {
			retval[key] = nil
		} else {
			if err != nil {
				retval[key] = nil
			} else {
				retval[key] = val
			}
		}
	}
	return retval
}

func (svc *RedisCacheService) GetObjects(ctx core.RequestContext, bucket string, keys []string, objectType string) map[string]interface{} {
	var args []interface{}
	for _, k := range keys {
		args = append(args, common.GetCacheKey(bucket, k))
	}
	retval := make(map[string]interface{})
	svrctx := ctx.ServerContext()
	objectcreator, err := svrctx.GetObjectCreator(objectType)
	if err != nil {
		return retval
	}
	conn := svc.pool.Get()
	defer conn.Close()
	k, err := redis.Values(conn.Do("MGET", args...))
	if err != nil {
		return retval
	}
	for index, val := range k {
		key := keys[index]
		if val == nil {
			retval[key] = nil
		} else {
			obj := objectcreator()
			err := svc.cacheEncoder.Decode(val.([]byte), obj)
			if err != nil {
				retval[key] = nil
			} else {
				retval[key] = obj
			}
		}
	}
	return retval
}

func (svc *RedisCacheService) Increment(ctx core.RequestContext, bucket string, key string) error {
	conn := svc.pool.Get()
	defer conn.Close()
	_, err := conn.Do("INCR", common.GetCacheKey(bucket, key))
	if err != nil {
		return err
	}
	conn.Flush()
	return nil
}
func (svc *RedisCacheService) Decrement(ctx core.RequestContext, bucket string, key string) error {
	conn := svc.pool.Get()
	defer conn.Close()
	_, err := conn.Do("DECR", common.GetCacheKey(bucket, key))
	if err != nil {
		return err
	}
	conn.Flush()
	return nil
}

func (redisSvc *RedisCacheService) Initialize(ctx core.ServerContext) error {
	redisSvc.SetDescription("Redis cache component service")
	redisSvc.SetComponent(true)
	redisSvc.AddStringConfigurations([]string{CONF_REDIS_CONNECTIONSTRING, CONF_REDIS_DATABASE, config.ENCODING}, []string{":6379", "0", "binary"})

	return nil
}

func (redisSvc *RedisCacheService) Start(ctx core.ServerContext) error {
	connectionString, _ := redisSvc.GetStringConfiguration(CONF_REDIS_CONNECTIONSTRING)
	redisSvc.connectionstring = connectionString

	connectiondb, _ := redisSvc.GetStringConfiguration(CONF_REDIS_DATABASE)
	redisSvc.database = connectiondb

	encoding, _ := redisSvc.GetStringConfiguration(config.ENCODING)
	redisSvc.cacheEncoder = common.NewCacheEncoder(ctx, encoding)

	redisSvc.pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				log.Error(ctx, "TestOnBorrow", "Error", err)
			}
			return err
		},
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", redisSvc.connectionstring)
			if err != nil {
				return nil, err
			}
			_, err = conn.Do("SELECT", redisSvc.database)
			if err != nil {
				conn.Close()
				return nil, err
			}
			_, err = conn.Do("FLUSHDB")
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}

	return nil
}
