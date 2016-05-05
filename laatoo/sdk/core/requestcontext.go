package core

import (
	"laatoo/sdk/auth"
)

type RequestContext interface {
	Context
	ServerContext() ServerContext
	EngineRequestContext() interface{}
	SubContext(name string) RequestContext
	GetServerElement(elemType ServerElementType) ServerElement
	NewContext(name string) RequestContext
	GetUser() auth.User
	HasPermission(perm string) bool
	GetRolePermissions(role []string) ([]string, bool)
	PublishMessage(topic string, message interface{})
	FireEvent(eventType string, eventObject string, data map[string]interface{})
	PutInCache(key string, item interface{}) error
	GetFromCache(key string, val interface{}) bool
	GetMultiFromCache(keys []string, val map[string]interface{}) bool
	InvalidateCache(key string) error
	IsAdmin() bool
	SetRequest(interface{})
	GetRequest() interface{}
	SetResponse(*ServiceResponse)
	GetResponse() *ServiceResponse
	PrintElapsedTime()
	CompleteRequest()
}
