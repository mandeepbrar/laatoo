package laatoopubsub

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"laatoo/core/registry"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"time"
)

type RedisPubSubFactory struct {
	Conf config.Config
}

const (
	CONF_REDISPUBSUB_NAME       = "redis_pubsub"
	CONF_REDISPUBSUB_SVC        = "pubsub"
	CONF_REDIS_CONNECTIONSTRING = "server"
	CONF_REDIS_DATABASE         = "db"
)

func init() {
	registry.RegisterServiceFactoryProvider(CONF_REDISPUBSUB_NAME, NewRedisPubSubFactory)
}

func NewRedisPubSubFactory(ctx core.ServerContext, conf config.Config) (core.ServiceFactory, error) {
	log.Logger.Info(ctx, "Creating redis pubsub factory ")
	redisFac := &RedisPubSubFactory{conf}
	return redisFac, nil
}

//Create the services configured for factory.
func (mf *RedisPubSubFactory) CreateService(ctx core.ServerContext, name string, conf config.Config) (core.Service, error) {
	if name == CONF_REDISPUBSUB_SVC {
		svc, err := NewRedisPubSubService(ctx, conf)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		return svc, nil
	}
	return nil, nil
}

//The services start serving when this method is called
func (ds *RedisPubSubFactory) StartServices(ctx core.ServerContext) error {
	return nil
}

type RedisPubSubService struct {
	connectionstring string
	database         string
	conn             redis.Conn
	pool             *redis.Pool
	conf             config.Config
}

func NewRedisPubSubService(ctx core.Context, conf config.Config) (core.PubSub, error) {
	log.Logger.Info(ctx, "Creating redis pubsub service ")
	redisSvc := &RedisPubSubService{conf: conf}

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

func (svc *RedisPubSubService) Publish(ctx core.Context, topic string, message interface{}) error {
	conn := svc.pool.Get()
	defer conn.Close()
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	_, err = conn.Do("PUBLISH", topic, bytes)
	log.Logger.Trace(ctx, "Published message on topic", "topic", topic)
	if err != nil {
		return err
	}
	conn.Flush()
	return nil
}

func (svc *RedisPubSubService) Subscribe(ctx core.Context, topics []string, lstnr core.TopicListener) error {
	conn := svc.pool.Get()
	psc := redis.PubSubConn{Conn: conn}
	for _, topic := range topics {
		err := psc.Subscribe(topic)
		if err != nil {
			return err
		}
	}
	go func() {
		for {
			switch v := psc.Receive().(type) {
			case redis.Message:
				log.Logger.Trace(ctx, "Message received on Queue")
				lstnr(ctx, v.Channel, v.Data)
			case redis.Subscription:
			case error:
				log.Logger.Info(ctx, "Pubsub error ", "Error", v)
			}
		}
	}()
	return nil
}

func (rs *RedisPubSubService) Initialize(ctx core.ServerContext) error {
	return nil
}

func (rs *RedisPubSubService) Invoke(ctx core.RequestContext) error {
	return nil
}

func (rs *RedisPubSubService) GetConf() config.Config {
	return rs.conf
}

func (rs *RedisPubSubService) GetResponseHandler() core.ServiceResponseHandler {
	return nil
}
