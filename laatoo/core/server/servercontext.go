package server

import (
	"laatoo/core/common"
	"laatoo/core/services"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
)

type serverContext struct {
	*common.Context
	conf        config.Config
	application *Application
}

func NewServerContext(name string, conf config.Config, app *Application) *serverContext {
	return &serverContext{Context: common.NewContext(name), conf: conf, application: app}
}

func (ctx *serverContext) EngineContext() core.EngineServerContext {
	return ctx.application.appEngine.GetContext()
}

func (ctx *serverContext) GetServerName() string {
	return ctx.application.Name
}
func (ctx *serverContext) GetServerType() string {
	return ctx.application.ServerType
}

func (ctx *serverContext) SubContext(name string, conf config.Config) core.ServerContext {
	return ctx.subCtx(name, conf, ctx.application)
}

func (ctx *serverContext) subCtx(name string, conf config.Config, app *Application) *serverContext {
	if app == nil {
		app = ctx.application
	}
	if conf == nil {
		conf = ctx.conf
	}
	duplicateContext := &serverContext{Context: ctx.DupCtx(name), conf: conf, application: app}
	return duplicateContext
}

func (ctx *serverContext) GetServerVariable(variable core.ServerVariable) interface{} {
	return ctx.application.GetVariable(variable)
}

func (ctx *serverContext) GetService(alias string) (core.Service, error) {
	return ctx.application.GetService(ctx, alias)
}

func (ctx *serverContext) SubscribeTopic(topic string, handler core.TopicListener) error {
	return ctx.application.SubscribeTopic(ctx, topic, handler)
}

func (ctx *serverContext) PublishMessage(topic string, message interface{}) error {
	return ctx.application.PublishMessage(ctx, topic, message)
}

func (ctx *serverContext) GetConf() config.Config {
	return ctx.conf
}

func (ctx *serverContext) PutInCache(key string, item interface{}) error {
	if ctx.application.Cache != nil {
		return ctx.application.Cache.PutObject(ctx, key, item)
	}
	return errors.ThrowError(ctx, CORE_ERROR_NO_CACHE_SVC)
}

func (ctx *serverContext) GetFromCache(key string, val interface{}) bool {
	if ctx.application.Cache != nil {
		return ctx.application.Cache.GetObject(ctx, key, val)
	}
	return false
}

func (ctx *serverContext) GetMultiFromCache(keys []string, val map[string]interface{}) bool {
	if ctx.application.Cache != nil {
		return ctx.application.Cache.GetMulti(ctx, keys, val)
	}
	return false
}

func (ctx *serverContext) DeleteFromCache(key string) error {
	if ctx.application.Cache != nil {
		return ctx.application.Cache.Delete(ctx, key)
	}
	return errors.ThrowError(ctx, CORE_ERROR_NO_CACHE_SVC)
}
func (ctx *serverContext) HasPermission(req core.RequestContext, perm string) bool {
	return ctx.application.Security.HasPermission(req, perm)
}
func (ctx *serverContext) GetRolePermissions(role []string) ([]string, bool) {
	return ctx.application.Security.GetRolePermissions(role)
}
func (ctx *serverContext) CreateNewRequest(name string) core.RequestContext {
	reqCtx := services.NewRequestContext(name, ctx.conf, ctx, nil)
	reqCtx.SetUser(ctx.application.serverUser)
	reqCtx.SetAdmin(true)
	return reqCtx
}

func (ctx *serverContext) ApplicationContext() core.ApplicationContext {
	return ctx.application.appContext
}
