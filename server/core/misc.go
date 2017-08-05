package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/server/constants"
	"laatoo/server/log"
)

func processLogging(ctx *serverContext, conf config.Config, name string) error {
	_, ok := conf.GetSubConfig(constants.CONF_LOGGING)
	if ok {
		elem := ctx.GetServerElement(core.ServerElementLogger)
		childLoggerHandle, childLogger := log.ChildLogger(ctx, name, elem)
		if err := childLoggerHandle.Initialize(ctx, conf); err != nil {
			return err
		}
		if err := childLoggerHandle.Start(ctx); err != nil {
			return err
		}
		ctx.setElements(core.ContextMap{core.ServerElementLogger: childLogger})
	}
	return nil
}
