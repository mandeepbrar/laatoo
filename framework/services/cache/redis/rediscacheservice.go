package redis

import (
	"laatoo/framework/core/objects"
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
	objects.RegisterObject(CONF_REDISCACHE_NAME, createRedisCacheServiceFactory, nil)
}

func createRedisCacheServiceFactory(ctx core.Context, args core.MethodArgs, conf config.Config) (interface{}, error) {
	return &RedisCacheFactory{}, nil
}

//Create the services configured for factory.
func (mf *RedisCacheFactory) CreateService(ctx core.ServerContext, name string, method string) (core.Service, error) {
	if name == CONF_REDISCACHE_SVC {
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
	conn := svc.pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, val)
	if err != nil {
		return err
	}
	conn.Flush()
	return nil
}

func (svc *RedisCacheService) GetObject(ctx core.RequestContext, bucket string, key string, val interface{}) bool {
	conn := svc.pool.Get()
	defer conn.Close()
	k, err := conn.Do("GET", key)
	if err != nil {
		return false
	}
	val = &k
	return true
}

func (svc *RedisCacheService) GetMulti(ctx core.RequestContext, bucket string, keys []string, val map[string]interface{}) {
	return false
}

func (rs *RedisCacheService) Invoke(ctx core.RequestContext) error {
	return nil
}

func (rs *RedisCacheService) Start(ctx core.ServerContext) error {
	return nil
}
