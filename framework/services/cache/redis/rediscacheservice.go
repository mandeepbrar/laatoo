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
		}}

	return nil
}

func (svc *RedisCacheService) Delete(ctx core.RequestContext, bucket string, key string) error {
	conn := svc.pool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	if err != nil {
		return err
	}
	return nil
}

func (svc *RedisCacheService) PutObject(ctx core.RequestContext, bucket string, key string, val interface{}) error {
	b, err := common.Encode(val)
	if err != nil {
		return err
	}
	conn := svc.pool.Get()
	defer conn.Close()
	_, err = conn.Do("SET", key, b)
	if err != nil {
		return err
	}
	conn.Flush()
	log.Logger.Error(ctx, "redis putting", "keys", key, "val", val)
	return nil
}

func (svc *RedisCacheService) GetObject(ctx core.RequestContext, bucket string, key string, objectType string) (interface{}, bool) {
	conn := svc.pool.Get()
	defer conn.Close()
	k, err := conn.Do("GET", key)
	if err != nil {
		return nil, false
	}
	return k, true
}

func (svc *RedisCacheService) GetMulti(ctx core.RequestContext, bucket string, keys []string, objectType string) map[string]interface{} {
	tim1 := time.Now()
	var args []interface{}
	for _, k := range keys {
		args = append(args, k)
	}
	retval := make(map[string]interface{})
	svrctx := ctx.ServerContext()
	objectcreator, err := svrctx.GetObjectCreator(objectType)
	if err != nil {
		return retval
	}
	tim11 := time.Now()
	conn := svc.pool.Get()
	defer conn.Close()
	k, err := redis.Values(conn.Do("MGET", args...))
	if err != nil {
		return retval
	}
	tim2 := time.Now()
	for index, val := range k {
		key := keys[index]
		if val == nil {
			retval[key] = nil
		} else {
			obj := objectcreator()
			err := common.Decode(val.([]byte), obj)
			if err != nil {
				retval[key] = nil
			} else {
				retval[key] = obj
			}
		}
	}
	tim3 := time.Now()
	log.Logger.Error(ctx, "time inside getmulti", "housekeeping", tim11.Sub(tim1), "redis", tim2.Sub(tim11), " decoding", tim3.Sub(tim2))
	//	val = &k
	return retval
}

func (rs *RedisCacheService) Invoke(ctx core.RequestContext) error {
	return nil
}

func (rs *RedisCacheService) Start(ctx core.ServerContext) error {
	return nil
}
