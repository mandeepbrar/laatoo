package common

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"
	"laatoo/server/constants"
)

func CheckContextCondition(ctx core.ServerContext, conf config.Config) bool {
	cond, ok := conf.GetSubConfig(ctx, constants.CONF_CONDITION)
	if ok {
		keys := cond.AllConfigurations(ctx)
		for _, key := range keys {
			str, ok := cond.GetString(ctx, key)
			if ok {
				contextVal, cok := ctx.GetString(key)
				if !cok || contextVal != str {
					return false
				}
			}
		}
	}
	return true
}

func SetupMiddleware(ctx core.ServerContext, conf config.Config) {
	parentMw, pok := ctx.GetStringArray(constants.CONF_MIDDLEWARE)
	middleware, ok := conf.GetStringArray(ctx, constants.CONF_MIDDLEWARE)
	if pok {
		if !ok {
			middleware = parentMw
		} else {
			middleware = append(parentMw, middleware...)
		}
	}
	if middleware != nil {
		ctx.Set(constants.CONF_MIDDLEWARE, middleware)
	}
	log.Trace(ctx, "Middleware setup", "middleware", middleware)
}
