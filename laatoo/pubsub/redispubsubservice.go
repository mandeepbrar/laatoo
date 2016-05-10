package pubsub

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	//"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"time"
)

const (
	CONF_REDIS_CONNECTIONSTRING = "server"
	CONF_REDIS_DATABASE         = "db"
)

type RedisPubSubService struct {
	connectionstring string
	name             string
	database         string
	conn             redis.Conn
	pool             *redis.Pool
}

func (svc *RedisPubSubService) Publish(ctx core.RequestContext, topic string, message interface{}) error {
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

func (svc *RedisPubSubService) Subscribe(ctx core.ServerContext, topics []string, lstnr core.ServiceFunc) error {
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
				req := ctx.CreateSystemRequest("Message Received")
				req.Set("messagetype", v.Channel)
				req.SetRequest(v.Data)
				lstnr(req)
			case redis.Subscription:
			case error:
				log.Logger.Info(ctx, "Pubsub error ", "Error", v)
			}
		}
	}()
	return nil
}

func (redisSvc *RedisPubSubService) Initialize(ctx core.ServerContext, conf config.Config) error {
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

func (rs *RedisPubSubService) Invoke(ctx core.RequestContext) error {
	return nil
}
func (rs *RedisPubSubService) Start(ctx core.ServerContext) error {
	return nil
}
