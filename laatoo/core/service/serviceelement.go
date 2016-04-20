package service

import (
	"laatoo/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/server"
	"laatoo/sdk/services"
)

type service struct {
	*common.Context
	name    string
	service core.Service
	factory server.Factory
	conf    config.Config
	owner   *serviceManager
}

func (svc *service) Service() core.Service {
	return svc.service
}
func (svc *service) PutInCache(key string, item interface{}) error {
	return nil
}
func (svc *service) GetFromCache(key string, val interface{}) bool {
	return false
}
func (svc *service) GetMultiFromCache(keys []string, val map[string]interface{}) bool {
	return false
}
func (svc *service) InvalidateCache(key string) error {
	return nil
}
func (svc *service) SubscribeTopic(topic string, handler services.TopicListener) error {
	return nil
}
func (svc *service) PublishMessage(topic string, message interface{}) error {
	return nil
}
func (svc *service) CreateNewRequest(name string) core.RequestContext {
	return nil
}
