package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/sdk/utils"
	"laatoo/server/common"
	"laatoo/server/engine/http"
	"laatoo/server/factory"
	"laatoo/server/objects"
	"laatoo/server/service"
	"path"
)

type abstractserver struct {
	name string

	objectLoader       server.ObjectLoader
	objectLoaderHandle server.ServerElementHandle

	channelMgr       server.ChannelManager
	channelMgrHandle server.ServerElementHandle

	factoryManager       server.FactoryManager
	factoryManagerHandle server.ServerElementHandle

	serviceManager       server.ServiceManager
	serviceManagerHandle server.ServerElementHandle

	securityHandler       server.SecurityHandler
	securityHandlerHandle server.ServerElementHandle

	messagingManager       server.MessagingManager
	messagingManagerHandle server.ServerElementHandle

	taskManager       server.TaskManager
	taskManagerHandle server.ServerElementHandle

	rulesManager       server.RulesManager
	rulesManagerHandle server.ServerElementHandle

	cacheManager       server.CacheManager
	cacheManagerHandle server.ServerElementHandle

	//engines configured on the abstract server
	engineHandles map[string]server.ServerElementHandle
	engines       map[string]server.Engine

	parent *abstractserver

	proxy core.ServerElement
}

func newAbstractServer(svrCtx *serverContext, name string, parent *abstractserver, proxy core.ServerElement, filterConf config.Config) *abstractserver {
	as := &abstractserver{name: name, parent: parent, proxy: proxy}
	if parent == nil {

		objCreateCtx := svrCtx.SubContext("Create Object Loader")
		objectLoaderHandle, objectLoader := objects.NewObjectLoader(objCreateCtx, name, proxy)
		as.objectLoaderHandle = objectLoaderHandle
		as.objectLoader = objectLoader.(server.ObjectLoader)
		svrCtx.objectLoader = as.objectLoader

		fmCreateCtx := svrCtx.SubContext("Create Factory Manager")
		factoryManagerHandle, factoryManager := factory.NewFactoryManager(fmCreateCtx, name, proxy)
		as.factoryManagerHandle = factoryManagerHandle
		as.factoryManager = factoryManager.(server.FactoryManager)
		svrCtx.factoryManager = as.factoryManager

		smCreateCtx := svrCtx.SubContext("Create Service Manager")
		serviceManagerHandle, serviceManager := service.NewServiceManager(smCreateCtx, name, proxy)
		as.serviceManagerHandle = serviceManagerHandle
		as.serviceManager = serviceManager.(server.ServiceManager)
		svrCtx.serviceManager = as.serviceManager

		cmCreateCtx := svrCtx.SubContext("Create Channel Manager")
		channelMgrHandle, channelMgr := newChannelManager(cmCreateCtx, name, proxy)
		as.channelMgr = channelMgr
		as.channelMgrHandle = channelMgrHandle
		svrCtx.channelManager = as.channelMgr

	} else {

		loader := parent.objectLoader
		factoryManager := parent.factoryManager
		serviceManager := parent.serviceManager
		channelMgr := parent.channelMgr

		objCreateCtx := svrCtx.SubContext("Create Object Loader")
		childLoaderHandle, childLoader := objects.ChildLoader(objCreateCtx, name, loader, proxy)
		as.objectLoaderHandle = childLoaderHandle
		as.objectLoader = childLoader.(server.ObjectLoader)
		svrCtx.objectLoader = as.objectLoader

		fmCreateCtx := svrCtx.SubContext("Create Factory Manager")
		childFactoryManagerHandle, childFactoryManager := factory.ChildFactoryManager(fmCreateCtx, name, factoryManager, proxy)
		as.factoryManagerHandle = childFactoryManagerHandle
		as.factoryManager = childFactoryManager.(server.FactoryManager)
		svrCtx.factoryManager = as.factoryManager

		smCreateCtx := svrCtx.SubContext("Create Service Manager")
		childServiceManagerHandle, childServiceManager := service.ChildServiceManager(smCreateCtx, name, serviceManager, proxy)
		as.serviceManagerHandle = childServiceManagerHandle
		as.serviceManager = childServiceManager.(server.ServiceManager)
		svrCtx.serviceManager = as.serviceManager

		cmCreateCtx := svrCtx.SubContext("Create Channel Manager")
		childChanMgrHandle, childChannelMgr := childChannelManager(cmCreateCtx, name, channelMgr, proxy)
		as.channelMgrHandle = childChanMgrHandle
		as.channelMgr = childChannelMgr.(server.ChannelManager)
		svrCtx.channelManager = as.channelMgr
	}
	return as
}

//initialize application with object loader, factory manager, service manager
func (as *abstractserver) initialize(ctx *serverContext, conf config.Config) error {

	secinit := ctx.subContext("Initialize security handleer")
	err := as.initializeSecurityHandler(secinit, conf)
	if err != nil {
		return errors.WrapError(secinit, err)
	}
	log.Logger.Trace(secinit, "Initialized security handler")
	ctx.securityHandler = as.securityHandler

	objinit := ctx.subContext("Initialize object loader")
	err = as.objectLoaderHandle.Initialize(objinit, conf)
	if err != nil {
		return errors.WrapError(objinit, err)
	}
	log.Logger.Trace(objinit, "Initialized object loader")

	enginit := ctx.SubContext("Initialize engines")
	err = as.initializeEngines(ctx, conf)
	if err != nil {
		return errors.WrapError(objinit, err)
	}
	log.Logger.Trace(objinit, "Initialized engines")

	chaninit := ctx.SubContextWithElement("Initialize channel manager", core.ServerElementChannelManager)
	err = initializeChannelManager(chaninit, conf, as.channelMgrHandle)
	if err != nil {
		return errors.WrapError(chaninit, err)
	}
	log.Logger.Debug(chaninit, "Initialized channel manager")

	middleware, ok := conf.GetStringArray(config.CONF_MIDDLEWARE)
	if ok {
		parentMw, ok := as.proxy.GetStringArray(config.CONF_MIDDLEWARE)
		if ok {
			middleware = append(parentMw, middleware...)
		}
		as.proxy.Set(config.CONF_MIDDLEWARE, middleware)
	}

	if (as.parent == nil) || (as.parent.messagingManager == nil) {
		msgSvcName, ok := conf.GetString(config.CONF_MESSAGING_SVC)
		if ok {
			msgCtx := ctx.SubContext("Create Messaging Manager")
			msgHandle, msgElem := newMessagingManager(msgCtx, "Root", as.proxy, msgSvcName)
			as.messagingManager = msgElem
			as.messagingManagerHandle = msgHandle
			log.Logger.Trace(msgCtx, "Created messaging manager")
		}
	} else {
		childMsgMgrHandle, childMsgMgr := childMessagingManager(ctx, as.name, as.parent.messagingManager, as.proxy)
		as.messagingManagerHandle = childMsgMgrHandle
		as.messagingManager = childMsgMgr.(server.MessagingManager)
	}
	ctx.msgManager = as.messagingManager
	if as.messagingManagerHandle != nil {
		msginit := ctx.SubContextWithElement("Initialize messaging manager", core.ServerElementMessagingManager)
		err := initializeMessagingManager(msginit, conf, as.messagingManagerHandle)
		if err != nil {
			return err
		}
		log.Logger.Debug(msginit, "Initialized messaging manager")
	}

	taskMgrHandle, taskMgr, err := createTaskManager(ctx, as.name, conf, as.proxy)
	if err != nil {
		return err
	} else {
		if taskMgr != nil {
			as.taskManager = taskMgr
			as.taskManagerHandle = taskMgrHandle
			log.Logger.Debug(ctx, "Created task manager")
		}
	}
	ctx.taskManager = as.taskManager

	var parentCacheMgr server.CacheManager
	if as.parent != nil {
		parentCacheMgr = as.parent.cacheManager
	}
	cacheMgrHandle, cacheMgr, err := createCacheManager(ctx, as.name, conf, parentCacheMgr, as.proxy)
	if err != nil {
		return err
	}
	if cacheMgr == nil {
		as.cacheManager = parentCacheMgr
	} else {
		as.cacheManager = cacheMgr
		as.cacheManagerHandle = cacheMgrHandle
	}
	ctx.cacheManager = as.cacheManager
	cacheToUse, ok := conf.GetString(config.CONF_CACHE_NAME)
	if ok {
		as.proxy.Set("__cache", cacheToUse)
	}

	facinit := ctx.SubContext("Initialize factory manager")
	err = initializeFactoryManager(facinit, conf, as.factoryManagerHandle)
	if err != nil {
		return err
	}
	log.Logger.Trace(facinit, "Initialized factory manager")

	svcinit := ctx.SubContext("Initialize service manager")
	err = initializeServiceManager(svcinit, conf, as.serviceManagerHandle)
	if err != nil {
		return err
	}
	log.Logger.Trace(svcinit, "Initialized service manager")

	rulesMgrHandle, rulesMgr, err := createRulesManager(ctx, as.name, conf, as.proxy)
	if err != nil {
		return err
	} else {
		if rulesMgr != nil {
			as.rulesManager = rulesMgr
			as.rulesManagerHandle = rulesMgrHandle
			log.Logger.Debug(ctx, "Created rules manager")
		}
	}
	ctx.rulesManager = as.rulesManager

	return nil
}

//start application with object loader, factory manager, service manager
func (as *abstractserver) start(ctx *serverContext) error {
	err := as.startSecurityHandler(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	objldrCtx := ctx.SubContextWithElement("Start ObjectLoader", core.ServerElementLoader)
	err = as.objectLoaderHandle.Start(objldrCtx)
	if err != nil {
		return errors.WrapError(objldrCtx, err)
	}
	log.Logger.Trace(objldrCtx, "Started Object Loader")

	engstart := ctx.SubContext("Start Engines")
	err = as.startEngines(engstart)
	if err != nil {
		return errors.WrapError(engstart, err)
	}
	log.Logger.Trace(engstart, "Started Engines")

	chanstart := ctx.SubContextWithElement("Start Channel manager", core.ServerElementChannelManager)
	err = as.channelMgrHandle.Start(chanstart)
	if err != nil {
		return errors.WrapError(chanstart, err)
	}

	fmCtx := ctx.SubContextWithElement("Start Factory Manager", core.ServerElementFactoryManager)
	err = as.factoryManagerHandle.Start(fmCtx)
	if err != nil {
		return errors.WrapError(fmCtx, err)
	}
	log.Logger.Trace(fmCtx, "Started factory manager")

	smCtx := ctx.SubContextWithElement("Start Service Manager", core.ServerElementServiceManager)
	err = as.serviceManagerHandle.Start(smCtx)
	if err != nil {
		return errors.WrapError(smCtx, err)
	}
	log.Logger.Trace(smCtx, "Started service manager")

	if (as.cacheManagerHandle != nil) && ((as.parent == nil) || (as.cacheManager != as.parent.cacheManager)) {
		cmCtx := ctx.SubContextWithElement("Start Cache Manager", core.ServerElementCacheManager)
		err = as.cacheManagerHandle.Start(cmCtx)
		if err != nil {
			return errors.WrapError(smCtx, err)
		}
	}

	if as.messagingManagerHandle != nil {
		msgstart := ctx.SubContextWithElement("Start messaging manager", core.ServerElementMessagingManager)
		err := as.messagingManagerHandle.Start(msgstart)
		if err != nil {
			return errors.WrapError(msgstart, err)
		}
	}

	if as.rulesManagerHandle != nil {
		rulesHCtx := ctx.SubContextWithElement("Start Rules Manager", core.ServerElementRulesManager)
		log.Logger.Trace(rulesHCtx, "Starting Rules Manager")
		err := as.rulesManagerHandle.Start(rulesHCtx)
		if err != nil {
			return errors.WrapError(rulesHCtx, err)
		}
	}

	if as.taskManagerHandle != nil {
		taskHCtx := ctx.SubContextWithElement("Start Task Manager", core.ServerElementTaskManager)
		log.Logger.Trace(taskHCtx, "Starting Task Manager")
		err := as.taskManagerHandle.Start(taskHCtx)
		if err != nil {
			return errors.WrapError(taskHCtx, err)
		}
	}

	return nil
}

func (as *abstractserver) initializeEngines(ctx *serverContext, conf config.Config) error {
	log.Logger.Trace(ctx, "Initializing engines")
	engines, ok := conf.GetSubConfig(config.CONF_ENGINES)
	if ok {
		engineNames := engines.AllConfigurations()
		for _, engName := range engineNames {
			engConf, err, _ := common.ConfigFileAdapter(ctx, engines, engName)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			engCreateCtx := initctx.SubContext("Create Engine: " + engName)
			log.Logger.Trace(engCreateCtx, "Creating")
			engHandle, eng, err := svr.createEngine(engCreateCtx, engName, engConf)
			if err != nil {
				return errors.WrapError(engCreateCtx, err)
			}

			engInitCtx := initctx.SubContext("Initialize Engine: " + engName)
			log.Logger.Trace(engInitCtx, "Initializing engine")
			err = engHandle.Initialize(engInitCtx, engConf)
			if err != nil {
				return errors.WrapError(engInitCtx, err)
			}

			//get a root channel and assign it to server channel manager
			svr.channelMgr.(*channelManagerProxy).manager.channelStore[engName] = eng.GetRootChannel(engInitCtx)
			log.Logger.Info(engInitCtx, "Registered root channel", "Name", engName)

			svr.engines[engName] = eng
			svr.engineHandles[engName] = engHandle
			log.Logger.Trace(initctx, "Registered engine", "Name", engName)
		}
	}
	log.Logger.Debug(initctx, "Initialized Engines")
}

func (as *abstractserver) initializeSecurityHandler(ctx *serverContext, conf config.Config) error {
	secConf, err, ok := common.ConfigFileAdapter(ctx, conf, config.CONF_SECURITY)
	if !ok {
		baseDir, _ := ctx.GetString(config.CONF_BASE_DIR)
		confFile := path.Join(baseDir, config.CONF_SECURITY, config.CONF_CONFIG_FILE)
		ok, _, _ := utils.FileExists(confFile)
		if ok {
			if secConf, err := common.NewConfigFromFile(confFile); err != nil {
				return err
			}
		}
	}
	if ok {
		secCtx := ctx.SubContext("Create Security Handler")
		shElem, sh := newSecurityHandler(secCtx, "Security Handler:"+as.name, as.proxy)
		secInitCtx := ctx.NewContextWithElements("Initialize Security Handler", core.ElementMap{core.ServerElementSecurityHandler: sh}, core.ServerElementSecurityHandler)
		err := shElem.Initialize(secInitCtx, secConf)
		if err != nil {
			return errors.WrapError(secInitCtx, err)
		}
		as.securityHandlerHandle = shElem
		as.securityHandler = sh.(server.SecurityHandler)
	} else {
		if as.parent != nil {
			as.securityHandler = as.parent.securityHandler
		}
	}
	return nil
}

func (as *abstractserver) startSecurityHandler(ctx *serverContext) error {
	if (as.securityHandlerHandle != nil) && ((as.parent == nil) || (as.securityHandler != as.parent.securityHandler)) {
		secCtx := ctx.SubContextWithElement("Start Security Handler", core.ServerElementSecurityHandler)
		log.Logger.Trace(secCtx, "Starting Security Handler")
		return as.securityHandlerHandle.Start(secCtx)
	}
	return nil
}

func (as *abstractserver) startEngines(ctx *serverContext) error {
	engStartCtx := startCtx.SubContext("Start Engines")
	for engName, engineHandle := range svr.engineHandles {
		errorsChan := make(chan error)
		go func(ctx core.ServerContext, engHandle server.ServerElementHandle, name string) {
			log.Logger.Info(ctx, "Starting engine*****", "name", name)
			errorsChan <- engHandle.Start(ctx)
			err := <-errorsChan
			if err != nil {
				panic(err.Error())
			}
		}(engStartCtx, engineHandle, engName)
	}
	log.Logger.Trace(startCtx, "Started engines")

}

func (as *abstractserver) createEngine(ctx core.ServerContext, name string, engConf config.Config) (server.ServerElementHandle, server.Engine, error) {
	enginetype, ok := engConf.GetString(config.CONF_ENGINE_TYPE)
	if !ok {
		return nil, nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Config Name", config.CONF_ENGINE_TYPE)
	}
	var engineHandle server.ServerElementHandle
	var engine server.Engine
	switch enginetype {
	case config.CONF_ENGINETYPE_HTTP:
		engineHandle, engine = http.NewEngine(ctx, name)
	case core.CONF_ENGINE_TCP:
	default:
		return nil, nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Config Name", config.CONF_ENGINE_TYPE)
	}
	return engineHandle, engine, nil
}

//creates a context specific to environment
func (as *abstractserver) contextMap(ctx core.ServerContext) core.ContextMap {
	return core.ContextMap{
		core.ServerElementLoader:           as.objectLoader,
		core.ServerElementSecurityHandler:  as.securityHandler,
		core.ServerElementMessagingManager: as.messagingManager,
		core.ServerElementChannelManager:   as.channelMgr,
		core.ServerElementCacheManager:     as.cacheManager,
		core.ServerElementTaskManager:      as.taskManager,
		core.ServerElementRulesManager:     as.rulesManager,
		core.ServerElementFactoryManager:   as.factoryManager,
		core.ServerElementServiceManager:   as.serviceManager}
}
