package core

import (
	"laatoo/sdk/auth"
	"laatoo/sdk/config"
)

type RequestContext interface {
	Context
	ParentContext() interface{}
	EngineContext() EngineRequestContext
	SubContext(name string, conf config.Config) RequestContext
	Get(key string) (interface{}, bool)
	Set(key string, val interface{})
	GetString(key string) (string, bool)
	GetInt(key string) (int, bool)
	GetStringArray(key string) ([]string, bool)
	GetUser() auth.User
	SetUser(usr auth.User)
	GetConf() config.Config
	GetService(alias string) (Service, error)
	HasPermission(perm string) bool
	SubscribeTopic(topic string, handler TopicListener) error
	PublishMessage(topic string, message interface{}) error
	PutInCache(key string, item interface{}) error
	GetFromCache(key string, val interface{}) bool
	GetMultiFromCache(keys []string, val map[string]interface{}) bool
	DeleteFromCache(key string) error
	SetAdmin(val bool)
	IsAdmin() bool
	SetRequestBody(interface{})
	GetRequestBody() interface{}
	SetResponse(*ServiceResponse)
	GetResponse() *ServiceResponse
	GetServerVariable(variable ServerVariable) interface{}
	GetRolePermissions(role []string) ([]string, bool)
}
