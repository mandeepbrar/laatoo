package core

import (
	"laatoo/sdk/config"
)

const (
	CONF_SERVERTYPE_STANDALONE = "STANDALONE"
	CONF_SERVERTYPE_GOOGLEAPP  = "GOOGLE_APP"
)

type ServerContext interface {
	Context
	SubContext(name string, conf config.Config) ServerContext
	GetServerType() string
	GetConf() config.Config
	GetService(alias string) (Service, error)
	SubscribeTopic(topic string, handler TopicListener) error
	PublishMessage(topic string, message interface{}) error
	PutInCache(key string, item interface{}) error
	GetFromCache(key string, val interface{}) bool
	GetMultiFromCache(keys []string, val map[string]interface{}) bool
	DeleteFromCache(key string) error
	GetServerVariable(variable int) interface{}
	GetServerName() string
	EngineContext() EngineContext
}
