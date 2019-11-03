package core

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
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

	if (as.secretsManagerHandle != nil) && ((as.parent == nil) || (as.secretsManagerHandle != as.parent.secretsManagerHandle)) {
		smCtx := ctx.SubContext("Start Secrets Manager")
		err = as.secretsManagerHandle.Start(smCtx)
		if err != nil {
			return errors.WrapError(smCtx, err)
		}
	}

	//module manager must be started before service manager and factory manager initializations
	//module manager is used during init of other managers
	modstart := ctx.SubContext("Start module manager")
	err = as.moduleManagerHandle.Start(modstart)
	if err != nil {
		return errors.WrapError(modstart, err)
	}
	log.Trace(modstart, "Started module managers")

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

	if as.communicatorHandle != nil {
		commstart := ctx.SubContext("Start communicator")
		err := as.communicatorHandle.Start(commstart)
		if err != nil {
			return errors.WrapError(commstart, err)
		}
	}

	if as.messagingManagerHandle != nil {
		msgstart := ctx.SubContext("Start messaging manager")
		err := as.messagingManagerHandle.Start(msgstart)
		if err != nil {
			return errors.WrapError(msgstart, err)
		}
	}

	if as.sessionManagerHandle != nil {
		sessCtx := ctx.SubContext("Start Session Manager")
		err := as.sessionManagerHandle.Start(sessCtx)
		if err != nil {
			return errors.WrapError(sessCtx, err)
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

	return as.onReady(ctx)
}

func (as *abstractserver) startSecurityHandler(ctx *serverContext) error {
	if (as.securityHandlerHandle != nil) && ((as.parent == nil) || (as.securityHandler != as.parent.securityHandler)) {
		secCtx := ctx.SubContext("Start Security Handler")
		log.Trace(secCtx, "Starting Security Handler")
		return as.securityHandlerHandle.Start(secCtx)
	}
	return nil
}

func (as *abstractserver) startSessionManager(ctx *serverContext) error {
	if (as.sessionManagerHandle != nil) && ((as.parent == nil) || (as.sessionManagerHandle != as.parent.sessionManagerHandle)) {
		sesCtx := ctx.SubContext("Start Session Manager")
		log.Trace(sesCtx, "Starting Session Manager")
		return as.sessionManagerHandle.Start(sesCtx)
	}
	return nil
}

func (as *abstractserver) startEngines(ctx core.ServerContext) error {
	engStartCtx := ctx.SubContext("Start Engines")
	for engName, engineHandle := range as.engineHandles {
		log.Error(ctx, "Starting engine*****", "name", engName)
		go func(ctx core.ServerContext, engHandle elements.ServerElementHandle, name string) {
			log.Info(ctx, "Starting engine*****", "name", name)
			err := engHandle.Start(ctx)
			if err != nil {
				panic(err.Error())
			}
		}(engStartCtx, engineHandle, engName)
	}
	log.Error(engStartCtx, "Started engines")
	return nil
}
