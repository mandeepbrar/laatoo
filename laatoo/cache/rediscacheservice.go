package cache

import (
	"github.com/garyburd/redigo/redis"
	"laatoo/core/registry"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/data"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"time"
)

type RedisCacheFactory struct {
	Conf config.Config
}

const (
	CONF_REDISCACHE_NAME        = "redis_cache"
	CONF_REDISCACHE_SVC         = "cache"
	CONF_REDIS_CONNECTIONSTRING = "server"
	CONF_REDIS_DATABASE         = "db"
)

func init() {
	registry.RegisterServiceFactoryProvider(CONF_REDISCACHE_NAME, RedisCacheServiceFactory)
}

func RedisCacheServiceFactory(ctx core.ServerContext, conf config.Config) (core.ServiceFactory, error) {
	log.Logger.Trace(ctx, "Creating redis cache service ")
	redisFac := &RedisCacheFactory{conf}
	return redisFac, nil
}

//Create the services configured for factory.
func (mf *RedisCacheFactory) CreateService(ctx core.ServerContext, name string, conf config.Config) (core.Service, error) {
	if name == CONF_REDISCACHE_SVC {
		svc, err := NewRedisCacheService(ctx, conf)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		return svc, nil
	}
	return nil, nil
}

//The services start serving when this method is called
func (ds *RedisCacheFactory) StartServices(ctx core.ServerContext) error {
	return nil
}

type RedisCacheService struct {
	connectionstring string
	database         string
	conn             redis.Conn
	pool             *redis.Pool
	conf             config.Config
}

func NewRedisCacheService(ctx core.Context, conf config.Config) (data.Cache, error) {
	log.Logger.Info(ctx, "Creating redis cache service ")
	redisSvc := &RedisCacheService{}

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

	return redisSvc, nil
}

func (svc *RedisCacheService) Delete(ctx core.Context, key string) error {
	return nil
}

func (svc *RedisCacheService) PutObject(ctx core.Context, key string, val interface{}) error {
	conn := svc.pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, val)
	if err != nil {
		return err
	}
	conn.Flush()
	return nil
}

func (svc *RedisCacheService) GetObject(ctx core.Context, key string, val interface{}) bool {
	conn := svc.pool.Get()
	defer conn.Close()
	k, err := conn.Do("GET", key)
	if err != nil {
		return false
	}
	val = &k
	return true
}

func (svc *RedisCacheService) GetMulti(ctx core.Context, keys []string, val map[string]interface{}) bool {
	return false
}

func (rs *RedisCacheService) Initialize(ctx core.ServerContext) error {
	return nil
}

func (rs *RedisCacheService) Invoke(ctx core.RequestContext) error {
	return nil
}

func (rs *RedisCacheService) GetConf() config.Config {
	return rs.conf
}
func (rs *RedisCacheService) GetResponseHandler() core.ServiceResponseHandler {
	return nil
}
