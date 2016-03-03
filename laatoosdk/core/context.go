package core

import (
	glctx "golang.org/x/net/context"
	"net/http"
)

type Context interface {
	NoContent(errorcode int) error
	Request() *http.Request
	ResponseWriter() http.ResponseWriter
	JSON(code int, val interface{}) error
	HTML(code int, format string, val ...interface{}) error
	Bind(i interface{}) error
	SetHeader(key string, val string)
	Redirect(code int, url string) error
	Get(key string) interface{}
	Set(key string, val interface{})
	Param(key string) string
	ParamByIndex(index int) string
	Query(key string) string
	GetConf() map[string]interface{}
	GetVariable(variable string) interface{}
	GetService(alias string) (Service, error)
	IsAllowed(perm string) bool
	SubscribeTopic(topic string, handler TopicListener) error
	PublishMessage(topic string, message interface{}) error
	PutInCache(key string, item interface{}) error
	GetFromCache(key string, val interface{}) error
	GetMultiFromCache(keys []string, val map[string]interface{}) error
	DeleteFromCache(key string) error
	HttpClient() *http.Client
	GetAppengineContext() glctx.Context
	GetCloudContext(scope string) glctx.Context
}
