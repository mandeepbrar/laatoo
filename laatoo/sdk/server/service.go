package server

import (
	"laatoo/sdk/core"
	"laatoo/sdk/services"
)

type Service interface {
	core.ServerElement
	Service() core.Service
	PutInCache(key string, item interface{}) error
	GetFromCache(key string, val interface{}) bool
	GetMultiFromCache(keys []string, val map[string]interface{}) bool
	InvalidateCache(key string) error
	SubscribeTopic(topic string, handler services.TopicListener) error
	PublishMessage(topic string, message interface{}) error
	CreateNewRequest(name string) core.RequestContext
}
