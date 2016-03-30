package services

import (
	"laatoo/core/common"
	"laatoo/sdk/auth"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"strconv"
)

type requestContext struct {
	*common.Context
	parentContext interface{}
	engineContext interface{}
	ParamsStore   map[string]interface{}
	User          auth.User
	Admin         bool
	responseData  *core.ServiceResponse
	requestBody   interface{}
	conf          config.Config
	serverContext core.ServerContext
}

func NewRequestContext(name string, conf config.Config, server core.ServerContext, engineCtx interface{}) *requestContext {
	return &requestContext{Context: common.NewContext(name), conf: conf, serverContext: server, ParamsStore: make(map[string]interface{}), engineContext: engineCtx}
}

func (ctx *requestContext) ParentContext() interface{} {
	return ctx.parentContext
}

func (ctx *requestContext) EngineContext() core.EngineContext {
	return ctx.engineContext
}

func (ctx *requestContext) SubContext(name string, conf config.Config) core.RequestContext {
	return ctx.subCtx(name, conf, ctx.serverContext)
}
func (ctx *requestContext) subCtx(name string, conf config.Config, serverContext core.ServerContext) *requestContext {
	duplicateMap := make(map[string]interface{}, len(ctx.ParamsStore))
	for k, v := range ctx.ParamsStore {
		duplicateMap[k] = v
	}
	if serverContext == nil {
		serverContext = ctx.serverContext
	}
	if conf == nil {
		conf = ctx.conf
	}
	duplicateContext := &requestContext{Context: ctx.DupCtx(name), conf: conf, serverContext: serverContext, ParamsStore: duplicateMap,
		parentContext: ctx, engineContext: ctx.engineContext}
	return duplicateContext
}

func (ctx *requestContext) Get(key string) (interface{}, bool) {
	val, ok := ctx.ParamsStore[key]
	return val, ok
}

func (ctx *requestContext) GetString(key string) (string, bool) {
	valInt, ok := ctx.Get(key)
	if ok {
		val, ok := valInt.(string)
		if ok {
			return val, true
		}
	}
	return "", false
}
func (ctx *requestContext) GetInt(key string) (int, bool) {
	valInt, ok := ctx.Get(key)
	if ok {
		val, ok := valInt.(string)
		if ok {
			intVal, err := strconv.Atoi(val)
			if err == nil {
				return intVal, true
			}
		}
	}
	return -1, false
}
func (ctx *requestContext) GetStringArray(key string) ([]string, bool) {
	valInt, ok := ctx.Get(key)
	if ok {
		val, ok := valInt.([]string)
		if ok {
			return val, true
		}
	}
	return nil, false
}

func (ctx *requestContext) GetConf() config.Config {
	return ctx.conf
}

func (ctx *requestContext) Set(key string, val interface{}) {
	ctx.ParamsStore[key] = val
}

func (ctx *requestContext) GetService(alias string) (core.Service, error) {
	return ctx.serverContext.GetService(alias)
}

func (ctx *requestContext) HasPermission(perm string) bool {
	//	return ctx.environment.HasPermission(ctx, perm)
	return false
}

func (ctx *requestContext) SubscribeTopic(topic string, handler core.TopicListener) error {
	//return ctx.environment.SubscribeTopic(ctx, topic, handler)
	return nil
}

func (ctx *requestContext) PublishMessage(topic string, message interface{}) error {
	//return ctx.environment.PublishMessage(ctx, topic, message)
	return nil
}

func (ctx *requestContext) PutInCache(key string, item interface{}) error {
	return ctx.serverContext.PutInCache(key, item)
}

func (ctx *requestContext) GetFromCache(key string, val interface{}) bool {
	return ctx.serverContext.GetFromCache(key, val)
}

func (ctx *requestContext) GetMultiFromCache(keys []string, val map[string]interface{}) bool {
	return ctx.serverContext.GetMultiFromCache(keys, val)
}

func (ctx *requestContext) DeleteFromCache(key string) error {
	return ctx.serverContext.DeleteFromCache(key)
}

func (ctx *requestContext) GetUser() auth.User {
	return ctx.User
}

func (ctx *requestContext) SetUser(usr auth.User) {
	ctx.User = usr
}

func (ctx *requestContext) IsAdmin() bool {
	return ctx.Admin
}

func (ctx *requestContext) SetResponse(responseData *core.ServiceResponse) {
	ctx.responseData = responseData
}
func (ctx *requestContext) GetResponse() *core.ServiceResponse {
	return ctx.responseData
}
func (ctx *requestContext) GetRequestBody() interface{} {
	return ctx.requestBody
}
func (ctx *requestContext) SetRequestBody(requestBody interface{}) {
	ctx.requestBody = requestBody
}
