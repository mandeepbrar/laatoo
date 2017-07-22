package main

import (
	"encoding/json"
	"laatoo/sdk/config"
	"laatoo/sdk/core"

	"github.com/garyburd/redigo/redis"
	//"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"time"
)

const (
	CONF_REDISPUBSUB_FACTORY    = "redispubsubfactory"
	CONF_REDISPUBSUB_SVC        = "redispubsub"
	CONF_REDIS_CONNECTIONSTRING = "server"
	CONF_REDIS_DATABASE         = "db"
)

func Manifest() []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_REDISPUBSUB_SVC, Object: RedisPubSubService{}},
		core.PluginComponent{Name: CONF_REDISPUBSUB_FACTORY, ObjectCreator: core.NewFactory(func() interface{} { return &RedisPubSubService{} })}}
}

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
	log.Trace(ctx, "Published message on topic", "topic", topic)
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
				log.Trace(ctx, "Message received on Queue")
				reqctx := ctx.CreateSystemRequest("Message Received")
				req := reqctx.CreateRequest()
				req.AddParam("messagetype", v.Channel, "", false)
				req.SetBody(v.Data)
				lstnr(reqctx, req)
			case redis.Subscription:
			case error:
				log.Info(ctx, "Pubsub error ", "Error", v)
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
		}}
	return nil
}

func (ds *RedisPubSubService) Info() *core.ServiceInfo {
	return &core.ServiceInfo{Description: "Redis pubsub component service"}
}

func (ms *RedisPubSubService) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	return nil, nil
}
func (rs *RedisPubSubService) Start(ctx core.ServerContext) error {
	return nil
}
