package common

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"laatoo/server/constants"
	"regexp"
)

func CheckContextCondition(ctx core.ServerContext, conf config.Config) bool {
	cond, ok := conf.GetSubConfig(constants.CONF_CONDITION)
	if ok {
		keys := cond.AllConfigurations()
		for _, key := range keys {
			str, ok := cond.GetString(key)
			if ok {
				contextVal, cok := ctx.GetVariable(key)
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
	middleware, ok := conf.GetStringArray(constants.CONF_MIDDLEWARE)
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

var varReplacer = regexp.MustCompile(`\[(.*?)\]`)

func FillVariables(ctx core.ServerContext, name string) string {
	return varReplacer.ReplaceAllStringFunc(name, func(variableName string) string {
		val, _ := ctx.GetVariable(variableName)
		return val
	})
}
