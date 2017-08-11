package core

import (
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
)

//start application with object loader, factory manager, service manager
func (as *abstractserver) start(ctx *serverContext) error {
	if err := as.loggerHandle.Start(ctx); err != nil {
		return errors.WrapError(ctx, err)
	}

	err := as.startSecurityHandler(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	objldrCtx := ctx.SubContext("Start Object Loader")
	err = as.objectLoaderHandle.Start(objldrCtx)
	if err != nil {
		return errors.WrapError(objldrCtx, err)
	}
	log.Trace(objldrCtx, "Started Object Loader")

	fmCtx := ctx.SubContext("Start Factory Manager")
	err = as.factoryManagerHandle.Start(fmCtx)
	if err != nil {
		return errors.WrapError(fmCtx, err)
	}
	log.Trace(fmCtx, "Started factory manager")

	smCtx := ctx.SubContext("Start Service Manager")
	err = as.serviceManagerHandle.Start(smCtx)
	if err != nil {
		return errors.WrapError(smCtx, err)
	}
	log.Trace(smCtx, "Started service manager")

	if (as.cacheManagerHandle != nil) && ((as.parent == nil) || (as.cacheManager != as.parent.cacheManager)) {
		cmCtx := ctx.SubContext("Start Cache Manager")
		err = as.cacheManagerHandle.Start(cmCtx)
		if err != nil {
			return errors.WrapError(smCtx, err)
		}
	}

	engstart := ctx.SubContext("Start Engines")
	err = as.startEngines(engstart)
	if err != nil {
		return errors.WrapError(engstart, err)
	}
	log.Trace(engstart, "Started Engines")

	chanstart := ctx.SubContext("Start Channel manager")
	log.Trace(chanstart, "Starting channel managers")
	err = as.channelManagerHandle.Start(chanstart)
	if err != nil {
		return errors.WrapError(chanstart, err)
	}
	log.Trace(chanstart, "Started channel managers")

	if as.messagingManagerHandle != nil {
		msgstart := ctx.SubContext("Start messaging manager")
		err := as.messagingManagerHandle.Start(msgstart)
		if err != nil {
			return errors.WrapError(msgstart, err)
		}
	}

	if as.rulesManagerHandle != nil {
		rulesHCtx := ctx.SubContext("Start Rules Manager")
		log.Trace(rulesHCtx, "Starting Rules Manager")
		err := as.rulesManagerHandle.Start(rulesHCtx)
		if err != nil {
			return errors.WrapError(rulesHCtx, err)
		}
	}

	if as.taskManagerHandle != nil {
		taskHCtx := ctx.SubContext("Start Task Manager")
		log.Trace(taskHCtx, "Starting Task Manager")
		err := as.taskManagerHandle.Start(taskHCtx)
		if err != nil {
			return errors.WrapError(taskHCtx, err)
		}
	}
	return nil
}

func (as *abstractserver) startSecurityHandler(ctx *serverContext) error {
	if (as.securityHandlerHandle != nil) && ((as.parent == nil) || (as.securityHandler != as.parent.securityHandler)) {
		secCtx := ctx.SubContext("Start Security Handler")
		log.Trace(secCtx, "Starting Security Handler")
		return as.securityHandlerHandle.Start(secCtx)
	}
	return nil
}

func (as *abstractserver) startEngines(ctx core.ServerContext) error {
	engStartCtx := ctx.SubContext("Start Engines")
	for engName, engineHandle := range as.engineHandles {
		go func(ctx core.ServerContext, engHandle server.ServerElementHandle, name string) {
			log.Info(ctx, "Starting engine*****", "name", name)
			err := engHandle.Start(ctx)
			if err != nil {
				panic(err.Error())
			}
		}(engStartCtx, engineHandle, engName)
	}
	log.Trace(engStartCtx, "Started engines")
	return nil
}
