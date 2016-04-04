package services

import (
	"laatoo/core/common"
	"laatoo/sdk/auth"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"strconv"
	"time"
)

type requestContext struct {
	*common.Context
	parentContext interface{}
	engineContext core.EngineRequestContext
	ParamsStore   map[string]interface{}
	User          auth.User
	Admin         bool
	responseData  *core.ServiceResponse
	requestBody   interface{}
	conf          config.Config
	serverContext core.ServerContext
	appContext    core.ApplicationContext
	createTime    time.Time
	subRequest    bool
}

func NewRequestContext(name string, conf config.Config, server core.ServerContext, engineCtx interface{}) *requestContext {
	return &requestContext{Context: common.NewContext(name), conf: conf, serverContext: server, ParamsStore: make(map[string]interface{}),
		engineContext: engineCtx, appContext: server.ApplicationContext(), createTime: time.Now(), subRequest: false}
}

func (ctx *requestContext) ParentContext() interface{} {
	return ctx.parentContext
}

func (ctx *requestContext) EngineContext() core.EngineRequestContext {
	return ctx.engineContext
}

func (ctx *requestContext) SubRequest(name string, conf config.Config) core.RequestContext {
	return ctx.subReq(name, conf, ctx.serverContext)
}
func (ctx *requestContext) subReq(name string, conf config.Config, serverContext core.ServerContext) *requestContext {
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
		parentContext: ctx, engineContext: ctx.engineContext, appContext: ctx.appContext, subRequest: true}
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

func (ctx *requestContext) SubscribeTopic(topic string, handler core.TopicListener) error {
	return ctx.serverContext.SubscribeTopic(topic, handler)
}

func (ctx *requestContext) PublishMessage(topic string, message interface{}) error {
	return ctx.serverContext.PublishMessage(topic, message)
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
func (ctx *requestContext) SetAdmin(val bool) {
	ctx.Admin = val
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
func (ctx *requestContext) GetServerVariable(variable core.ServerVariable) interface{} {
	return ctx.serverContext.GetServerVariable(variable)
}
func (ctx *requestContext) HasPermission(perm string) bool {
	return ctx.serverContext.HasPermission(ctx, perm)
}
func (ctx *requestContext) GetRolePermissions(role []string) ([]string, bool) {
	return ctx.serverContext.GetRolePermissions(role)
}
func (ctx *requestContext) ApplicationContext() core.ApplicationContext {
	return ctx.appContext
}
func (ctx *requestContext) CompleteRequest() {
	if !ctx.subRequest {
		completionTime := time.Now()
		elapsedTime := completionTime.Sub(ctx.createTime)
		log.Logger.Info(ctx, "Request Complete", "Time taken", elapsedTime)
	}
	ctx.parentContext = nil
	ctx.engineContext = nil
	ctx.ParamsStore = nil
	ctx.User = nil
	ctx.Admin = false
	ctx.responseData = nil
	ctx.requestBody = nil
	ctx.conf = nil
	ctx.serverContext = nil
	ctx.appContext = nil
}
