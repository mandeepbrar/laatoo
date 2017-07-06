package common

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/server/constants"
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
