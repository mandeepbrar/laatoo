package server

import (
	"laatoo/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
)

type serverContext struct {
	*common.Context
	conf        config.Config
	environment *Environment
}

func NewServerContext(name string, conf config.Config, env *Environment) *serverContext {
	return &serverContext{Context: common.NewContext(name), conf: conf, environment: env}
}

func (ctx *serverContext) EngineContext() core.EngineContext {
	return ctx.environment.envEngine.GetContext()
}

func (ctx *serverContext) GetServerName() string {
	return ctx.environment.Name
}
func (ctx *serverContext) GetServerType() string {
	return ctx.environment.ServerType
}

func (ctx *serverContext) SubContext(name string, conf config.Config) core.ServerContext {
	return ctx.subCtx(name, conf, ctx.environment)
}

func (ctx *serverContext) subCtx(name string, conf config.Config, env *Environment) *serverContext {
	if env == nil {
		env = ctx.environment
	}
	if conf == nil {
		conf = ctx.conf
	}
	duplicateContext := &serverContext{Context: ctx.DupCtx(name), conf: conf, environment: env}
	return duplicateContext
}

func (ctx *serverContext) GetServerVariable(variable int) interface{} {
	return ctx.environment.GetVariable(variable)
}

func (ctx *serverContext) GetService(alias string) (core.Service, error) {
	return ctx.environment.GetService(ctx, alias)
}

func (ctx *serverContext) SubscribeTopic(topic string, handler core.TopicListener) error {
	//return ctx.environment.SubscribeTopic(ctx, topic, handler)
	return nil
}

func (ctx *serverContext) PublishMessage(topic string, message interface{}) error {
	//return ctx.environment.PublishMessage(ctx, topic, message)
	return nil
}

func (ctx *serverContext) GetConf() config.Config {
	return ctx.conf
}

func (ctx *serverContext) PutInCache(key string, item interface{}) error {
	if ctx.environment.Cache != nil {
		return ctx.environment.Cache.PutObject(ctx, key, item)
	}
	return errors.ThrowError(ctx, CORE_ERROR_NO_CACHE_SVC)
}

func (ctx *serverContext) GetFromCache(key string, val interface{}) bool {
	if ctx.environment.Cache != nil {
		return ctx.environment.Cache.GetObject(ctx, key, val)
	}
	return false
}

func (ctx *serverContext) GetMultiFromCache(keys []string, val map[string]interface{}) bool {
	if ctx.environment.Cache != nil {
		return ctx.environment.Cache.GetMulti(ctx, keys, val)
	}
	return false
}

func (ctx *serverContext) DeleteFromCache(key string) error {
	if ctx.environment.Cache != nil {
		return ctx.environment.Cache.Delete(ctx, key)
	}
	return errors.ThrowError(ctx, CORE_ERROR_NO_CACHE_SVC)
}
