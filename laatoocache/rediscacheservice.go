package laatoopubsub

import (
	"github.com/garyburd/redigo/redis"
	"laatoocore"
	//	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"time"
)

type RedisCacheService struct {
	name             string
	context          service.ServiceContext
	connectionstring string
	database         string
	conn             redis.Conn
	pool             *redis.Pool
}

const (
	LOGGING_CONTEXT             = "rediscache"
	CONF_REDISCACHE_NAME        = "redis_cache"
	CONF_REDIS_CONNECTIONSTRING = "server"
	CONF_REDIS_DATABASE         = "db"
)

func init() {
	laatoocore.RegisterObjectProvider(CONF_REDISCACHE_NAME, RedisCacheServiceFactory)
}

func RedisCacheServiceFactory(conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(LOGGING_CONTEXT, "Creating redis cache service ")
	redisSvc := &RedisCacheService{name: CONF_REDISCACHE_NAME}

	connectionStringInt, ok := conf[CONF_REDIS_CONNECTIONSTRING]
	if ok {
		redisSvc.connectionstring = connectionStringInt.(string)
	} else {
		redisSvc.connectionstring = ":6379"
	}

	connectiondbInt, ok := conf[CONF_REDIS_DATABASE]
	if ok {
		redisSvc.database = connectiondbInt.(string)
	} else {
		redisSvc.database = "0"
	}

	redisSvc.pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				log.Logger.Error(LOGGING_CONTEXT, "TestOnBorrow", "Error", err)
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

func (svc *RedisCacheService) GetServiceType() string {
	return service.SERVICE_TYPE_DATA
}

//name of the service
func (svc *RedisCacheService) GetName() string {
	return svc.name
}

//Initialize the service. Consumer of a service passes the data
func (svc *RedisCacheService) Initialize(ctx service.ServiceContext) error {
	svc.context = ctx
	return nil
}

//The service starts serving when this method is called
func (svc *RedisCacheService) Serve() error {
	return nil
}

func (svc *RedisCacheService) Delete(key string) error {
	return nil
}

func (svc *RedisCacheService) PutObject(key string, val interface{}) error {
	conn := svc.pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, val)
	if err != nil {
		return err
	}
	conn.Flush()
	return nil
}

func (svc *RedisCacheService) GetObject(key string) (interface{}, error) {
	conn := svc.pool.Get()
	defer conn.Close()
	return conn.Do("GET", key)
}

//Execute method
func (svc *RedisCacheService) Execute(name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}
