package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/sdk/utils"
	"laatoo/server/cache"
	"laatoo/server/common"
	"laatoo/server/constants"
	"laatoo/server/engine/http"
	slog "laatoo/server/log"
	"laatoo/server/objects"
	"laatoo/server/rules"
	"laatoo/server/tasks"
	"path"
)

func (as *abstractserver) createNonConfComponents(svrCtx *serverContext, name string, parent *abstractserver, proxy core.ServerElement) {
	if parent == nil {
		loggerCreateCtx := svrCtx.SubContext("Create Logger")
		loggerHandle, logger := slog.NewLogger(loggerCreateCtx, name)
		as.logger = logger
		as.loggerHandle = loggerHandle
		svrCtx.setElements(core.ContextMap{core.ServerElementLogger: logger})

		objCreateCtx := svrCtx.SubContext("Create Object Loader")
		objectLoaderHandle, objectLoader := objects.NewObjectLoader(objCreateCtx, name, proxy)
		as.objectLoaderHandle = objectLoaderHandle
		as.objectLoader = objectLoader.(server.ObjectLoader)
		svrCtx.setElements(core.ContextMap{core.ServerElementLoader: objectLoader})

		fmCreateCtx := svrCtx.SubContext("Create Factory Manager")
		factoryManagerHandle, factoryManager := as.newFactoryManager(fmCreateCtx, name, proxy)
		as.factoryManagerHandle = factoryManagerHandle
		as.factoryManager = factoryManager.(server.FactoryManager)
		svrCtx.setElements(core.ContextMap{core.ServerElementFactoryManager: factoryManager})

		smCreateCtx := svrCtx.SubContext("Create Service Manager")
		serviceManagerHandle, serviceManager := as.newServiceManager(smCreateCtx, name, proxy)
		as.serviceManagerHandle = serviceManagerHandle
		as.serviceManager = serviceManager.(server.ServiceManager)
		svrCtx.setElements(core.ContextMap{core.ServerElementServiceManager: serviceManager})

		cmCreateCtx := svrCtx.SubContext("Create Channel Manager")
		channelMgrHandle, channelMgr := newChannelManager(cmCreateCtx, name, proxy)
		as.channelManager = channelMgr
		as.channelManagerHandle = channelMgrHandle
		svrCtx.setElements(core.ContextMap{core.ServerElementChannelManager: channelMgr})

	} else {

		logger := parent.logger
		loader := parent.objectLoader
		factoryManager := parent.factoryManager
		serviceManager := parent.serviceManager
		channelMgr := parent.channelManager

		loggerHandle, logger := slog.ChildLogger(svrCtx, name, logger)
		as.logger = logger
		as.loggerHandle = loggerHandle
		svrCtx.setElements(core.ContextMap{core.ServerElementLogger: logger})

		objCreateCtx := svrCtx.SubContext("Create Object Loader")
		childLoaderHandle, childLoader := objects.ChildLoader(objCreateCtx, name, loader, proxy)
		as.objectLoaderHandle = childLoaderHandle
		as.objectLoader = childLoader.(server.ObjectLoader)
		svrCtx.setElements(core.ContextMap{core.ServerElementLoader: childLoader})

		fmCreateCtx := svrCtx.SubContext("Create Factory Manager")
		childFactoryManagerHandle, childFactoryManager := as.childFactoryManager(fmCreateCtx, name, factoryManager, proxy)
		as.factoryManagerHandle = childFactoryManagerHandle
		as.factoryManager = childFactoryManager.(server.FactoryManager)
		svrCtx.setElements(core.ContextMap{core.ServerElementFactoryManager: childFactoryManager})

		smCreateCtx := svrCtx.SubContext("Create Service Manager")
		childServiceManagerHandle, childServiceManager := as.childServiceManager(smCreateCtx, name, serviceManager, proxy)
		as.serviceManagerHandle = childServiceManagerHandle
		as.serviceManager = childServiceManager.(server.ServiceManager)
		svrCtx.setElements(core.ContextMap{core.ServerElementServiceManager: childServiceManager})

		cmCreateCtx := svrCtx.SubContext("Create Channel Manager")
		childChanMgrHandle, childChannelMgr := childChannelManager(cmCreateCtx, name, channelMgr, proxy)
		as.channelManagerHandle = childChanMgrHandle
		as.channelManager = childChannelMgr.(server.ChannelManager)
		svrCtx.setElements(core.ContextMap{core.ServerElementChannelManager: childChannelMgr})
	}

	taskCreateCtx := svrCtx.SubContext("Create Task Manager")
	taskMgrHandle, taskMgr := tasks.NewTaskManager(taskCreateCtx, name)
	if taskMgr != nil {
		as.taskManager = taskMgr
		as.taskManagerHandle = taskMgrHandle
	}
	svrCtx.setElements(core.ContextMap{core.ServerElementTaskManager: taskMgr})

	rulesCreateCtx := svrCtx.SubContext("Create Rules Manager")
	rulesMgrHandle, rulesMgr := rules.NewRulesManager(rulesCreateCtx, name)
	if rulesMgr != nil {
		as.rulesManager = rulesMgr
		as.rulesManagerHandle = rulesMgrHandle
	}
	svrCtx.setElements(core.ContextMap{core.ServerElementRulesManager: rulesMgr})

}

func (as *abstractserver) createConfBasedComponents(ctx *serverContext, conf config.Config) error {
	createsecctx := ctx.subContext("Create Security Manager: " + as.name)
	if err := as.createSecurityHandler(createsecctx, conf); err != nil {
		return errors.WrapError(createsecctx, err)
	}
	ctx.setElements(core.ContextMap{core.ServerElementSecurityHandler: as.securityHandler})

	createcachectx := ctx.subContext("Create Cache manager: " + as.name)
	if err := as.createCacheManager(createcachectx, conf); err != nil {
		return errors.WrapError(createcachectx, err)
	}
	ctx.setElements(core.ContextMap{core.ServerElementCacheManager: as.cacheManager})

	createenginectx := ctx.SubContext("Create Engines: " + as.name)
	if err := as.createEngines(createenginectx, conf); err != nil {
		return errors.WrapError(createenginectx, err)
	}

	createmsgctx := ctx.subContext("Create Messaging Manager: " + as.name)
	if err := as.createMessagingManager(createmsgctx, conf, as.parent); err != nil {
		return errors.WrapError(createmsgctx, err)
	}
	ctx.setElements(core.ContextMap{core.ServerElementMessagingManager: as.messagingManager})

	return nil
}

func (as *abstractserver) newServiceManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (server.ServerElementHandle, core.ServerElement) {
	sm := &serviceManager{name: name, parent: parentElem, servicesStore: make(map[string]*serviceProxy, 100)}
	smElem := &serviceManagerProxy{manager: sm}
	sm.proxy = smElem
	return sm, smElem
}

func (as *abstractserver) childServiceManager(ctx core.ServerContext, name string, parentSvcMgr core.ServerElement, parent core.ServerElement, filters ...server.Filter) (server.ServerElementHandle, core.ServerElement) {
	svcMgrProxy := parentSvcMgr.(*serviceManagerProxy)
	svcMgr := svcMgrProxy.manager
	store := make(map[string]*serviceProxy, len(svcMgr.servicesStore))
	for k, v := range svcMgr.servicesStore {
		allowed := true
		for _, filter := range filters {
			if !filter.Allowed(ctx, k) {
				allowed = false
				break
			}
		}
		if allowed {
			store[k] = v
		}
	}
	sm := &serviceManager{name: name, parent: parent, servicesStore: store}
	smElem := &serviceManagerProxy{manager: sm}
	sm.proxy = smElem
	return sm, smElem
}

func (as *abstractserver) newFactoryManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (server.ServerElementHandle, core.ServerElement) {
	fm := &factoryManager{name: name, parent: parentElem, serviceFactoryStore: make(map[string]*serviceFactoryProxy, 30), svrref: as}
	fmElem := &factoryManagerProxy{manager: fm}
	fm.proxy = fmElem
	return fm, fmElem
}

func (as *abstractserver) childFactoryManager(ctx core.ServerContext, name string, parentFacMgr core.ServerElement, parent core.ServerElement, filters ...server.Filter) (server.ServerElementHandle, core.ServerElement) {
	facMgrProxy := parentFacMgr.(*factoryManagerProxy)
	facMgr := facMgrProxy.manager
	store := make(map[string]*serviceFactoryProxy, len(facMgr.serviceFactoryStore))
	for k, v := range facMgr.serviceFactoryStore {
		allowed := true
		for _, filter := range filters {
			if !filter.Allowed(ctx, k) {
				allowed = false
				break
			}
		}
		if allowed {
			store[k] = v
		}
	}
	fm := &factoryManager{parent: parent, serviceFactoryStore: store, svrref: as}
	fmElem := &factoryManagerProxy{manager: fm}
	fm.proxy = fmElem
	return fm, fmElem
}

func newChannelManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (*channelManager, *channelManagerProxy) {
	cm := &channelManager{name: name, channelStore: make(map[string]server.Channel, 10), parent: parentElem}
	cmElem := &channelManagerProxy{manager: cm}
	cm.proxy = cmElem
	return cm, cmElem
}

func childChannelManager(ctx core.ServerContext, name string, parentChannelMgr core.ServerElement, parent core.ServerElement, filters ...server.Filter) (server.ServerElementHandle, core.ServerElement) {
	chanMgrProxy := parentChannelMgr.(*channelManagerProxy)
	chanMgr := chanMgrProxy.manager
	store := make(map[string]server.Channel, len(chanMgr.channelStore))
	for k, v := range chanMgr.channelStore {
		allowed := true
		for _, filter := range filters {
			if !filter.Allowed(ctx, k) {
				allowed = false
				break
			}
		}
		if allowed {
			store[k] = v
		}
	}
	cm := &channelManager{name: name, channelStore: store, parent: parent}
	cmElem := &channelManagerProxy{manager: cm}
	cm.proxy = cmElem
	return cm, cmElem
}

/*
Messaging manager needs communication service
*/
func (as *abstractserver) createMessagingManager(ctx *serverContext, conf config.Config, parent *abstractserver) error {
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
			msgConf = make(common.GenericConfig, 0)
		}
	}
	if found {
		msgSvcName, ok := msgConf.GetString(constants.CONF_MESSAGING_SVC)
		if ok {
			msgCtx := ctx.SubContext("Create Messaging Manager")
			msgHandle, msgElem := newMessagingManager(msgCtx, as.name, msgSvcName)
			as.messagingManager = msgElem
			as.messagingManagerHandle = msgHandle
			log.Trace(msgCtx, "Created messaging manager")
		}
	}

	if (as.messagingManager == nil) && (parent != nil) && (parent.messagingManager != nil) {
		childMsgMgrHandle, childMsgMgr := childMessagingManager(ctx, as.name, parent.messagingManager, nil)
		as.messagingManagerHandle = childMsgMgrHandle
		as.messagingManager = childMsgMgr.(server.MessagingManager)
	}

	return nil
}

/*
override security handler depending upon the conf existence
*/
func (as *abstractserver) createSecurityHandler(ctx *serverContext, conf config.Config) error {
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
	createSecHandler := func(ctx core.ServerContext) {
		log.Trace(ctx, "Creating security handler")
		shElem, sh := newSecurityHandler(ctx, "Security Handler:"+as.name, as.proxy)
		as.securityHandlerHandle = shElem
		as.securityHandler = sh.(server.SecurityHandler)
	}

	if secConf != nil {
		createSecHandler(ctx)
	} else {
		if as.parent != nil {
			as.securityHandler = as.parent.securityHandler
			as.securityHandlerHandle = as.parent.securityHandlerHandle
		} else {
			createSecHandler(ctx)
		}
	}
	return nil
}

func (as *abstractserver) createEngines(ctx core.ServerContext, conf config.Config) error {
	ctx = ctx.SubContext("Create Engines")
	engines, ok := conf.GetSubConfig(constants.CONF_ENGINES)
	if ok {
		engineNames := engines.AllConfigurations()
		for _, engName := range engineNames {
			engConf, err, ok := common.ConfigFileAdapter(ctx, engines, engName)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			if ok {
				if err := as.processEngineConf(ctx, engConf, engName); err != nil {
					return err
				}
			}
			log.Trace(ctx, "Registered engine", "Name", engName)
		}
	}

	if err := common.ProcessDirectoryFiles(ctx, constants.CONF_ENGINES, as.processEngineConf, true); err != nil {
		return err
	}
	return nil
}

func (as *abstractserver) processEngineConf(ctx core.ServerContext, conf config.Config, name string) error {
	_, found := as.engines[name]
	if !found {
		engCreateCtx := ctx.SubContext("Create Engine: " + name)
		log.Trace(engCreateCtx, "Creating Engine", "Engine", name)
		engHandle, eng, err := as.createEngine(engCreateCtx, conf, name)
		if err != nil {
			return errors.WrapError(engCreateCtx, err)
		}

		as.engines[name] = eng
		as.engineConf[name] = conf
		as.engineHandles[name] = engHandle
	} else {
		log.Info(ctx, "Engine already exists", "Name", name)
	}
	return nil
}

func (as *abstractserver) createCacheManager(ctx *serverContext, conf config.Config) error {
	cacheMgrCreateCtx := ctx.SubContext("Create Cache Manager")
	var cacheMgrHandle server.ServerElementHandle
	var cacheMgr server.CacheManager
	if as.parent == nil || as.parent.cacheManager == nil {
		cacheMgrHandle, cacheMgr = cache.NewCacheManager(cacheMgrCreateCtx, "Root")
	} else {
		cacheMgrHandle, cacheMgr = cache.ChildCacheManager(cacheMgrCreateCtx, as.name, as.parent.cacheManager)
	}

	if cacheMgr == nil && as.parent != nil {
		as.cacheManager = as.parent.cacheManager
	} else {
		as.cacheManager = cacheMgr
		as.cacheManagerHandle = cacheMgrHandle
	}
	log.Trace(cacheMgrCreateCtx, "Cache Manager Created")
	return nil
}

func (as *abstractserver) createEngine(ctx core.ServerContext, engConf config.Config, engName string) (server.ServerElementHandle, server.Engine, error) {
	enginetype, ok := engConf.GetString(constants.CONF_ENGINE_TYPE)
	if !ok {
		return nil, nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Config Name", constants.CONF_ENGINE_TYPE)
	}
	var engineHandle server.ServerElementHandle
	var engine server.Engine
	switch enginetype {
	case constants.CONF_ENGINETYPE_HTTP:
		engineHandle, engine = http.NewEngine(ctx, engName, engConf)
	case core.CONF_ENGINE_TCP:
	default:
		return nil, nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Config Name", constants.CONF_ENGINE_TYPE)
	}
	return engineHandle, engine, nil
}
