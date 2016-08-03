package redis

import (
	"laatoo/framework/core/objects"
	"laatoo/framework/services/cache/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"

	"github.com/garyburd/redigo/redis"
	//	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"time"
)

type RedisCacheFactory struct {
}

const (
	CONF_REDISCACHE_NAME        = "redis_cache"
	CONF_REDISCACHE_SVC         = "cache"
	CONF_REDIS_CONNECTIONSTRING = "server"
	CONF_REDIS_DATABASE         = "db"
)

func init() {
	objects.Register(CONF_REDISCACHE_NAME, RedisCacheFactory{})
}

//Create the services configured for factory.
func (mf *RedisCacheFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	if method == CONF_REDISCACHE_SVC {
		return &RedisCacheService{name: name}, nil
	}
	return nil, nil
}

func (ds *RedisCacheFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//The services start serving when this method is called
func (ds *RedisCacheFactory) Start(ctx core.ServerContext) error {
	return nil
}

type RedisCacheService struct {
	connectionstring string
	database         string
	conn             redis.Conn
	pool             *redis.Pool
	name             string
	cacheEncoder     *common.CacheEncoder
}

func (redisSvc *RedisCacheService) Initialize(ctx core.ServerContext, conf config.Config) error {

	connectionString, ok := conf.GetString(CONF_REDIS_CONNECTIONSTRING)
	if ok {
		redisSvc.connectionstring = connectionString
	} else {
		redisSvc.connectionstring = ":6379"
	}

	connectiondb, ok := conf.GetString(CONF_REDIS_DATABASE)
	if ok {
		redisSvc.database = connectiondb
	} else {
		redisSvc.database = "0"
	}

	encoding, ok := conf.GetString(config.CONF_CACHE_ENC)
	if ok {
		redisSvc.cacheEncoder = common.NewCacheEncoder(ctx, encoding)
	} else {
		redisSvc.cacheEncoder = common.NewCacheEncoder(ctx, "binary")
	}

	redisSvc.pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				log.Logger.Error(ctx, "TestOnBorrow", "Error", err)
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

func (rs *RedisCacheService) Invoke(ctx core.RequestContext) error {
	return nil
}

func (rs *RedisCacheService) Start(ctx core.ServerContext) error {
	return nil
}
