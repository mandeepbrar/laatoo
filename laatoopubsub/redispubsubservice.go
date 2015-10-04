package laatoopubsub

import (
	"github.com/garyburd/redigo/redis"
	"laatoocore"
	//	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"time"
)

type RedisPubSubService struct {
	name             string
	context          service.ServiceContext
	connectionstring string
	database         string
	conn             redis.Conn
	pool             *redis.Pool
}

const (
	LOGGING_CONTEXT             = "redis"
	CONF_REDISPUBSUB_NAME       = "redis_pubsub"
	CONF_REDIS_CONNECTIONSTRING = "server"
	CONF_REDIS_DATABASE         = "db"
)

func init() {
	laatoocore.RegisterObjectProvider(CONF_REDISPUBSUB_NAME, RedisPubSubServiceFactory)
}

func RedisPubSubServiceFactory(ctx interface{}, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Creating redis pubsub service ")
	redisSvc := &RedisPubSubService{name: CONF_REDISPUBSUB_NAME}

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
				log.Logger.Error(ctx, LOGGING_CONTEXT, "TestOnBorrow", "Error", err)
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

func (svc *RedisPubSubService) GetServiceType() string {
	return service.SERVICE_TYPE_DATA
}

//name of the service
func (svc *RedisPubSubService) GetName() string {
	return svc.name
}

//Initialize the service. Consumer of a service passes the data
func (svc *RedisPubSubService) Initialize(ctx service.ServiceContext) error {
	svc.context = ctx
	return nil
}

//The service starts serving when this method is called
func (svc *RedisPubSubService) Serve(ctx interface{}) error {
	return nil
}

func (svc *RedisPubSubService) Publish(ctx interface{}, topic string, message interface{}) error {
	conn := svc.pool.Get()
	defer conn.Close()
	_, err := conn.Do("PUBLISH", topic, message)
	if err != nil {
		return err
	}
	conn.Flush()
	return nil
}

func (svc *RedisPubSubService) Subscribe(ctx interface{}, topics []string, lstnr service.TopicListener) error {
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
				lstnr(ctx, v.Channel, v.Data)
			case redis.Subscription:
			case error:
				log.Logger.Info(ctx, LOGGING_CONTEXT, "Pubsub error ", "Error", v)
			}
		}
	}()
	return nil
}

//Execute method
func (svc *RedisPubSubService) Execute(ctx interface{}, name string, params map[string]interface{}) (interface{}, error) {
	return nil, nil
}
