package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/sdk/utils"
	"laatoo/server/common"
	"laatoo/server/constants"
	"path"
)

//initialize application with object loader, factory manager, service manager
func (as *abstractserver) initialize(ctx *serverContext, conf config.Config) error {
	ctx = ctx.subContext("initialize components: " + as.name)

	if err := as.loggerHandle.Initialize(ctx, conf); err != nil {
		return errors.WrapError(ctx, err)
	}

	contextFile := path.Join(as.baseDir, constants.CONF_CONTEXT, constants.CONF_CONFIG_FILE)
	ok, _, _ := utils.FileExists(contextFile)
	if ok {
		contextVars, err := common.NewConfigFromFile(contextFile)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		keys := contextVars.AllConfigurations()
		for _, key := range keys {
			val, _ := contextVars.Get(key)
			ctx.SetVariable(key, val.(string))
		}
	}

	createctx := ctx.subContext("Create Conf components")
	if err := as.createConfBasedComponents(createctx, conf); err != nil {
		return err
	}

	secinit := ctx.subContext("Initialize security handleer")
	err := as.initializeSecurityHandler(secinit, conf)
	if err != nil {
		return errors.WrapError(secinit, err)
	}
	log.Trace(secinit, "Initialized security handler")

	common.SetupMiddleware(ctx, conf)

	modsctx := ctx.SubContext("Modules Manager: " + as.name)
	err = as.moduleManagerHandle.Initialize(modsctx, conf)
	if err != nil {
		return errors.WrapError(modsctx, err)
	}
	log.Debug(ctx, "Initialized modules manager")

	if err = as.initializeServicesCore(ctx, conf); err != nil {
		return errors.WrapError(ctx, err)
	}

	if err := as.initializeCacheManager(ctx, conf); err != nil {
		return err
	}

	enginit := ctx.SubContext("Initialize engines")
	err = as.initializeEngines(enginit, conf)
	if err != nil {
		return errors.WrapError(enginit, err)
	}
	log.Trace(enginit, "Initialized engines")

	chaninit := ctx.SubContext("Initialize channel manager")
	err = initializeChannelManager(chaninit, conf, as.channelManagerHandle)
	if err != nil {
		return errors.WrapError(chaninit, err)
	}
	log.Debug(chaninit, "Initialized channel manager")

	if as.messagingManagerHandle != nil {
		msginit := ctx.SubContext("Initialize messaging manager")
		err := as.initializeMessagingManager(msginit, as.name, conf, as.parent)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		log.Debug(msginit, "Initialized messaging manager")
	}

	taskctx := ctx.SubContext("Task Manager: " + as.name)
	err = as.taskManagerHandle.Initialize(taskctx, conf)
	if err != nil {
		return errors.WrapError(taskctx, err)
	}
	log.Debug(ctx, "Initialized task manager")

	rulesctx := ctx.SubContext("Rules Manager: " + as.name)
	err = as.rulesManagerHandle.Initialize(rulesctx, conf)
	if err != nil {
		return errors.WrapError(rulesctx, err)
	}
	log.Debug(ctx, "Initialized rules manager")

	return nil
}

func (as *abstractserver) initializeServicesCore(ctx *serverContext, conf config.Config) error {

	objinit := ctx.subContext("Initialize object loader")
	err := as.objectLoaderHandle.Initialize(objinit, conf)
	if err != nil {
		return errors.WrapError(objinit, err)
	}
	log.Trace(objinit, "Initialized object loader")

	facinit := ctx.SubContext("Initialize factory manager")
	err = initializeFactoryManager(facinit, conf, as.factoryManagerHandle)
	if err != nil {
		return err
	}
	log.Trace(facinit, "Initialized factory manager")

	svcinit := ctx.SubContext("Initialize service manager")
	err = initializeServiceManager(svcinit, conf, as.serviceManagerHandle)
	if err != nil {
		return errors.WrapError(svcinit, err)
	}
	log.Trace(svcinit, "Initialized service manager")

	return nil
}

func initializeChannelManager(ctx core.ServerContext, conf config.Config, channelManagerHandle server.ServerElementHandle) error {
	chmgrconf, err, ok := common.ConfigFileAdapter(ctx, conf, constants.CONF_CHANNELS)
	if err != nil {
		return err
	}
	if !ok {
		chmgrconf = make(config.GenericConfig, 0)
	}
	err = channelManagerHandle.Initialize(ctx, chmgrconf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (as *abstractserver) initializeMessagingManager(ctx core.ServerContext, name string, conf config.Config, parent *abstractserver) error {
	msgConf, err, found := common.ConfigFileAdapter(ctx, conf, constants.CONF_MESSAGING)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if !found {
		basedir, _ := ctx.GetString(constants.CONF_BASE_DIR)
		confFile := path.Join(basedir, constants.CONF_MESSAGING, constants.CONF_CONFIG_FILE)
		found, _, _ = utils.FileExists(confFile)
		if found {
			var err error
			if msgConf, err = common.NewConfigFromFile(confFile); err != nil {
				return errors.WrapError(ctx, err)
			}
		} else {
			msgConf = make(config.GenericConfig, 0)
		}
	}
	if as.messagingManagerHandle != nil {
		msginit := ctx.SubContext("Initialize messaging manager")
		err := as.messagingManagerHandle.Initialize(msginit, msgConf)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		log.Debug(msginit, "Initialized messaging manager")
	}
	return nil
}

func initializeFactoryManager(ctx core.ServerContext, conf config.Config, factoryManagerHandle server.ServerElementHandle) error {
	facConf, err, ok := common.ConfigFileAdapter(ctx, conf, constants.CONF_SERVICEFACTORIES)
	if err != nil {
		return err
	}
	if !ok {
		facConf = make(config.GenericConfig, 0)
	}
	err = factoryManagerHandle.Initialize(ctx, facConf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func initializeServiceManager(ctx core.ServerContext, conf config.Config, serviceManagerHandle server.ServerElementHandle) error {
	svcConf, err, ok := common.ConfigFileAdapter(ctx, conf, constants.CONF_SERVICES)
	if err != nil {
		return err
	}
	if !ok {
		svcConf = make(config.GenericConfig, 0)
	}
	err = serviceManagerHandle.Initialize(ctx, svcConf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (as *abstractserver) initializeEngines(ctx core.ServerContext, conf config.Config) error {
	for name, engHandle := range as.engineHandles {
		engInitCtx := ctx.SubContext("Initialize Engine: " + name)
		engConf := as.engineConf[name]
		err := engHandle.Initialize(engInitCtx, engConf)
		if err != nil {
			return errors.WrapError(engInitCtx, err)
		}

		//register engines root channel
		eng := as.engines[name]
		rootChannel := eng.GetRootChannel(engInitCtx)
		//get a root channel and assign it to server channel manager
		as.channelManagerHandle.(*channelManager).channelStore[name] = rootChannel

		log.Trace(engInitCtx, "Initialized engine", "Engine", name)
	}
	log.Debug(ctx, "Initialized Engines")
	return nil
}

func (as *abstractserver) initializeSecurityHandler(ctx *serverContext, conf config.Config) error {
	secConf, err, ok := common.ConfigFileAdapter(ctx, conf, constants.CONF_SECURITY)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	if !ok {
		confFile := path.Join(as.baseDir, constants.CONF_SECURITY, constants.CONF_CONFIG_FILE)
		ok, _, _ = utils.FileExists(confFile)
		if ok {
			var err error
			if secConf, err = common.NewConfigFromFile(confFile); err != nil {
				return err
			}
		}
	}
	initializeHandler := as.securityHandler != nil
	if as.parent != nil && as.securityHandler == as.parent.securityHandler {
		initializeHandler = false
	}
	if initializeHandler {
		secInitCtx := ctx.newContext("Initialize Security Handler")
		secInitCtx.setElements(core.ContextMap{core.ServerElementSecurityHandler: as.securityHandler})
		err := as.securityHandlerHandle.Initialize(secInitCtx, secConf)
		if err != nil {
			return errors.WrapError(secInitCtx, err)
		}
	}
	return nil
}

func (as *abstractserver) initializeCacheManager(ctx core.ServerContext, conf config.Config) error {
	cacheMgrInitCtx := ctx.SubContext("Initialize Cache Manager: " + as.name)
	err := as.cacheManagerHandle.Initialize(cacheMgrInitCtx, conf)
	if err != nil {
		return errors.WrapError(cacheMgrInitCtx, err)
	}

	cacheToUse, ok := conf.GetString(constants.CONF_CACHE_NAME)
	if ok {
		ctx.Set("__cache", cacheToUse)
		log.Debug(ctx, "Cache Set ", "Cache name", cacheToUse)
	}

	log.Trace(cacheMgrInitCtx, "Cache Manager Initialized")
	return nil
}
