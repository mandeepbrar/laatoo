package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/components/rules"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"laatoo/server/common"
	"laatoo/server/constants"
	"laatoo/server/engine/http"
	"laatoo/server/engine/websocket"
	slog "laatoo/server/log"
	"path"

	"github.com/blang/semver"
)

func (as *abstractserver) createNonConfComponents(svrCtx *serverContext, name string, parent *abstractserver, proxy core.ServerElement) {
	if parent == nil {
		loggerCreateCtx := svrCtx.SubContext("Create Logger")
		loggerHandle, logger := slog.NewLogger(loggerCreateCtx, name)
		as.logger = logger
		as.loggerHandle = loggerHandle
		svrCtx.setElements(core.ContextMap{core.ServerElementLogger: logger})

		objCreateCtx := svrCtx.SubContext("Create Object Loader")
		objectLoaderHandle, objectLoader := newObjectLoader(objCreateCtx, name, proxy)
		as.objectLoaderHandle = objectLoaderHandle
		as.objectLoader = objectLoader.(elements.ObjectLoader)
		svrCtx.setElements(core.ContextMap{core.ServerElementLoader: objectLoader})

		fmCreateCtx := svrCtx.SubContext("Create Factory Manager")
		factoryManagerHandle, factoryManager := as.newFactoryManager(fmCreateCtx, name, proxy)
		as.factoryManagerHandle = factoryManagerHandle
		as.factoryManager = factoryManager.(elements.FactoryManager)
		svrCtx.setElements(core.ContextMap{core.ServerElementFactoryManager: factoryManager})

		smCreateCtx := svrCtx.SubContext("Create Service Manager")
		serviceManagerHandle, serviceManager := as.newServiceManager(smCreateCtx, name, proxy)
		as.serviceManagerHandle = serviceManagerHandle
		as.serviceManager = serviceManager.(elements.ServiceManager)
		svrCtx.setElements(core.ContextMap{core.ServerElementServiceManager: serviceManager})

		cmCreateCtx := svrCtx.SubContext("Create Channel Manager")
		channelMgrHandle, channelMgr := newChannelManager(cmCreateCtx, name, proxy)
		as.channelManager = channelMgr
		as.channelManagerHandle = channelMgrHandle
		svrCtx.setElements(core.ContextMap{core.ServerElementChannelManager: channelMgr})

		modCreateCtx := svrCtx.SubContext("Create Module Manager")
		modMgrHandle, modMgr := as.newModuleManager(modCreateCtx, name, proxy)
		if modMgr != nil {
			as.moduleManager = modMgr
			as.moduleManagerHandle = modMgrHandle
		}
		svrCtx.setElements(core.ContextMap{core.ServerElementModuleManager: modMgr})

	} else {

		logger := parent.logger
		loader := parent.objectLoader
		factoryManager := parent.factoryManager
		serviceManager := parent.serviceManager
		channelMgr := parent.channelManager
		moduleMgr := parent.moduleManager

		loggerHandle, logger := slog.ChildLogger(svrCtx, name, logger)
		as.logger = logger
		as.loggerHandle = loggerHandle
		svrCtx.setElements(core.ContextMap{core.ServerElementLogger: logger})

		objCreateCtx := svrCtx.SubContext("Create Object Loader")
		childLoaderHandle, childLoader := childLoader(objCreateCtx, name, loader, proxy)
		as.objectLoaderHandle = childLoaderHandle
		as.objectLoader = childLoader.(elements.ObjectLoader)
		svrCtx.setElements(core.ContextMap{core.ServerElementLoader: childLoader})

		fmCreateCtx := svrCtx.SubContext("Create Factory Manager")
		childFactoryManagerHandle, childFactoryManager := as.childFactoryManager(fmCreateCtx, name, factoryManager, proxy)
		as.factoryManagerHandle = childFactoryManagerHandle
		as.factoryManager = childFactoryManager.(elements.FactoryManager)
		svrCtx.setElements(core.ContextMap{core.ServerElementFactoryManager: childFactoryManager})

		smCreateCtx := svrCtx.SubContext("Create Service Manager")
		childServiceManagerHandle, childServiceManager := as.childServiceManager(smCreateCtx, name, serviceManager, proxy)
		as.serviceManagerHandle = childServiceManagerHandle
		as.serviceManager = childServiceManager.(elements.ServiceManager)
		svrCtx.setElements(core.ContextMap{core.ServerElementServiceManager: childServiceManager})

		cmCreateCtx := svrCtx.SubContext("Create Channel Manager")
		childChanMgrHandle, childChannelMgr := childChannelManager(cmCreateCtx, name, channelMgr, proxy)
		as.channelManagerHandle = childChanMgrHandle
		as.channelManager = childChannelMgr.(elements.ChannelManager)
		svrCtx.setElements(core.ContextMap{core.ServerElementChannelManager: childChannelMgr})

		modCreateCtx := svrCtx.SubContext("Create Module Manager")
		modMgrHandle, modMgr := as.childModuleManager(modCreateCtx, name, moduleMgr, proxy)
		if modMgr != nil {
			as.moduleManager = modMgr
			as.moduleManagerHandle = modMgrHandle
		}
		svrCtx.setElements(core.ContextMap{core.ServerElementModuleManager: modMgr})
	}

	taskCreateCtx := svrCtx.SubContext("Create Task Manager")
	taskMgrHandle, taskMgr := newTaskManager(taskCreateCtx, name)
	if taskMgr != nil {
		as.taskManager = taskMgr
		as.taskManagerHandle = taskMgrHandle
	}
	svrCtx.setElements(core.ContextMap{core.ServerElementTaskManager: taskMgr})

	rulesCreateCtx := svrCtx.SubContext("Create Rules Manager")
	rulesMgrHandle, rulesMgr := newRulesManager(rulesCreateCtx, name)
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

	createsessionctx := ctx.subContext("Create Session Manager " + as.name)
	if err := as.createSessionManager(createsessionctx, conf); err != nil {
		return errors.WrapError(createsessionctx, err)
	}

	createmsgctx := ctx.subContext("Create Messaging Manager: " + as.name)
	if err := as.createMessagingManager(createmsgctx, conf, as.parent); err != nil {
		return errors.WrapError(createmsgctx, err)
	}
	ctx.setElements(core.ContextMap{core.ServerElementMessagingManager: as.messagingManager})

	return nil
}

func (as *abstractserver) newServiceManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (elements.ServerElementHandle, core.ServerElement) {
	sm := &serviceManager{name: name, parent: parentElem, servicesStore: make(map[string]*serviceProxy, 100), factoryManager: as.factoryManager}
	smElem := &serviceManagerProxy{manager: sm}
	sm.proxy = smElem
	return sm, smElem
}

func (as *abstractserver) childServiceManager(ctx core.ServerContext, name string, parentSvcMgr core.ServerElement, parent core.ServerElement, filters ...elements.Filter) (elements.ServerElementHandle, core.ServerElement) {
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
	sm := &serviceManager{name: name, parent: parent, servicesStore: store, factoryManager: as.factoryManager}
	smElem := &serviceManagerProxy{manager: sm}
	sm.proxy = smElem
	return sm, smElem
}

func (as *abstractserver) newFactoryManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (elements.ServerElementHandle, core.ServerElement) {
	fm := &factoryManager{name: name, parent: parentElem, serviceFactoryStore: make(map[string]*serviceFactoryProxy, 30), svrref: as}
	fmElem := &factoryManagerProxy{manager: fm}
	fm.proxy = fmElem
	return fm, fmElem
}

func (as *abstractserver) childFactoryManager(ctx core.ServerContext, name string, parentFacMgr core.ServerElement, parent core.ServerElement, filters ...elements.Filter) (elements.ServerElementHandle, core.ServerElement) {
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
	cm := &channelManager{name: name, channelStore: make(map[string]elements.Channel, 10), parent: parentElem, channelConfs: make(map[string]config.Config), parentChannels: make(map[string]string)}
	cmElem := &channelManagerProxy{manager: cm}
	cm.proxy = cmElem
	return cm, cmElem
}

func childChannelManager(ctx core.ServerContext, name string, parentChannelMgr core.ServerElement, parent core.ServerElement, filters ...elements.Filter) (elements.ServerElementHandle, core.ServerElement) {
	chanMgrProxy := parentChannelMgr.(*channelManagerProxy)
	chanMgr := chanMgrProxy.manager
	store := make(map[string]elements.Channel, len(chanMgr.channelStore))
	confs := make(map[string]config.Config)
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
			confs[k] = chanMgr.channelConfs[k]
		}
	}
	cm := &channelManager{name: name, channelStore: store, parent: parent, channelConfs: confs, parentChannels: make(map[string]string)}
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
		basedir, _ := ctx.GetString(config.BASEDIR)
		confFile := path.Join(basedir, constants.CONF_MESSAGING, constants.CONF_CONFIG_FILE)
		found, _, _ = utils.FileExists(confFile)
		if found {
			var err error
			if msgConf, err = common.NewConfigFromFile(ctx, confFile, nil); err != nil {
				return errors.WrapError(ctx, err)
			}
		} else {
			msgConf = ctx.CreateConfig()
		}
	}
	if found {
		msgSvcName, ok := msgConf.GetString(ctx, constants.CONF_MESSAGING_SVC)
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
		as.messagingManager = childMsgMgr.(elements.MessagingManager)
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
			if secConf, err = common.NewConfigFromFile(ctx, confFile, nil); err != nil {
				return err
			}
		}
	}
	createSecHandler := func(ctx core.ServerContext) {
		log.Trace(ctx, "Creating security handler")
		shElem, sh := newSecurityHandler(ctx, "Security Handler:"+as.name, as.proxy)
		as.securityHandlerHandle = shElem
		as.securityHandler = sh.(elements.SecurityHandler)
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

func (as *abstractserver) createSessionManager(ctx *serverContext, conf config.Config) error {
	sesConf, err, ok := common.ConfigFileAdapter(ctx, conf, constants.CONF_SESSION)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	if ok {
		smElem, sm := newSessionManager(ctx, "Session Manager:"+as.name, sesConf)
		as.sessionManagerHandle = smElem
		as.sessionManager = sm
	} else {
		if as.parent != nil {
			as.sessionManager = as.parent.sessionManager
			as.sessionManagerHandle = as.parent.sessionManagerHandle
		}
	}
	return nil
}

func (as *abstractserver) createEngines(ctx core.ServerContext, conf config.Config) error {
	ctx = ctx.SubContext("Create Engines")
	engines, ok := conf.GetSubConfig(ctx, constants.CONF_ENGINES)
	if ok {
		engineNames := engines.AllConfigurations(ctx)
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

	engs, err := common.ProcessDirectoryFiles(ctx, as.baseDir, constants.CONF_ENGINES, true)
	if err != nil {
		return err
	}

	if err = common.ProcessObjects(ctx, engs, as.processEngineConf); err != nil {
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
	var cacheMgrHandle elements.ServerElementHandle
	var cacheMgr elements.CacheManager
	if as.parent == nil || as.parent.cacheManager == nil {
		cacheMgrHandle, cacheMgr = newCacheManager(cacheMgrCreateCtx, "Root")
	} else {
		cacheMgrHandle, cacheMgr = childCacheManager(cacheMgrCreateCtx, as.name, as.parent.cacheManager)
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

func (as *abstractserver) createEngine(ctx core.ServerContext, engConf config.Config, engName string) (elements.ServerElementHandle, elements.Engine, error) {
	enginetype, ok := engConf.GetString(ctx, constants.CONF_ENGINE_TYPE)
	if !ok {
		return nil, nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Config Name", constants.CONF_ENGINE_TYPE)
	}
	var engineHandle elements.ServerElementHandle
	var engine elements.Engine
	switch enginetype {
	case constants.CONF_ENGINETYPE_HTTP:
		engineHandle, engine = http.NewEngine(ctx, engName, engConf)
	case constants.CONF_ENGINETYPE_WS:
		engineHandle, engine = websocket.NewEngine(ctx, engName, engConf)
	case core.CONF_ENGINE_TCP:
	default:
		return nil, nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Config Name", constants.CONF_ENGINE_TYPE)
	}
	return engineHandle, engine, nil
}

func (as *abstractserver) newModuleManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (elements.ServerElementHandle, core.ServerElement) {
	mm := &moduleManager{name: name, parent: parentElem, moduleInstances: make(map[string]*serverModule), installedModules: make(map[string]*semver.Version), availableModules: make(map[string]string),
		loadedModules: make(map[string]*semver.Version), parentModules: make(map[string]*serverModule), modulePlugins: make(map[string]components.ModuleManagerPlugin), moduleConf: make(map[string]config.Config),
		hotModules: make(map[string]string), moduleInstallationConf: make(map[string]config.Config), svrref: as}
	mmElem := &moduleManagerProxy{modMgr: mm}
	mm.proxy = mmElem
	return mm, mmElem
}

func (as *abstractserver) childModuleManager(ctx core.ServerContext, name string, parentModMgr core.ServerElement, parent core.ServerElement, filters ...elements.Filter) (elements.ServerElementHandle, core.ServerElement) {
	modMgrProxy := parentModMgr.(*moduleManagerProxy)
	modMgr := modMgrProxy.modMgr
	moduleInstances := make(map[string]*serverModule, len(modMgr.moduleInstances))
	for k, v := range modMgr.moduleInstances {
		moduleInstances[k] = v
	}
	installedModules := make(map[string]*semver.Version, len(modMgr.installedModules))
	for k, v := range modMgr.installedModules {
		installedModules[k] = v
	}
	loadedModules := make(map[string]*semver.Version, len(modMgr.loadedModules))
	for k, v := range modMgr.installedModules {
		loadedModules[k] = v
	}
	availableModules := make(map[string]string, len(modMgr.availableModules))
	for k, v := range modMgr.availableModules {
		availableModules[k] = v
	}
	moduleConf := make(map[string]config.Config, len(modMgr.moduleConf))
	for k, v := range modMgr.moduleConf {
		moduleConf[k] = v
	}
	modulePlugins := make(map[string]components.ModuleManagerPlugin)
	for k, v := range modMgr.modulePlugins {
		modulePlugins[k] = v
	}
	hotModules := make(map[string]string)
	for k, v := range modMgr.hotModules {
		hotModules[k] = v
	}

	childModMgr := &moduleManager{name: name, parent: parent, moduleInstances: moduleInstances, installedModules: installedModules, parentModules: make(map[string]*serverModule),
		loadedModules: loadedModules, moduleConf: moduleConf, svrref: as, availableModules: availableModules, modulePlugins: modulePlugins,
		moduleInstallationConf: make(map[string]config.Config), hotModules: hotModules}
	childModMgrProxy := &moduleManagerProxy{modMgr: childModMgr}
	return childModMgr, childModMgrProxy
}

func newObjectLoader(ctx core.ServerContext, name string, parentElem core.ServerElement) (elements.ServerElementHandle, core.ServerElement) {
	ldr := &objectLoader{objectsFactoryRegister: make(map[string]core.ObjectFactory, 30), objModMap: make(utils.StringMap), name: name, parentElem: parentElem, provider: &metadataProvider{}}
	ldrElem := &objectLoaderProxy{loader: ldr}
	return ldr, ldrElem
}

func childLoader(ctx core.ServerContext, name string, parentLdr core.ServerElement, parent core.ServerElement, filters ...elements.Filter) (elements.ServerElementHandle, core.ServerElement) {
	objLdrProxy := parentLdr.(*objectLoaderProxy)
	objLoader := objLdrProxy.loader
	registry := make(map[string]core.ObjectFactory, len(objLoader.objectsFactoryRegister))
	for k, v := range objLoader.objectsFactoryRegister {
		allowed := true
		for _, filter := range filters {
			if !filter.Allowed(ctx, k) {
				allowed = false
				break
			}
		}
		if allowed {
			registry[k] = v
		}
	}

	childModMap := objLoader.objModMap.Clone()

	log.Trace(ctx, "carrying over the following objects to the child", "objects", registry)
	ldr := &objectLoader{objectsFactoryRegister: registry, objModMap: childModMap, name: name, parentElem: parent, provider: objLoader.provider}
	ldrElem := &objectLoaderProxy{loader: ldr}
	return ldr, ldrElem
}

func newRulesManager(ctx core.ServerContext, name string) (*rulesManager, *rulesManagerProxy) {
	rulesMgr := &rulesManager{registeredRules: make(map[string][]rules.Rule, 10), rulesStore: make(map[string]rules.Rule, 10), name: name}
	rulesElem := &rulesManagerProxy{manager: rulesMgr}
	rulesMgr.proxy = rulesElem
	return rulesMgr, rulesElem
}

func newTaskManager(ctx core.ServerContext, name string) (*taskManager, *taskManagerProxy) {
	tskMgr := &taskManager{name: name, taskPublisherSvcs: make(map[string]components.TaskQueue, 10), taskProcessors: make(map[string]elements.Service, 10),
		taskPublishers: make(map[string]string, 10), taskConsumerNames: make(map[string]string, 10), taskProcessorNames: make(map[string]string, 10)}
	tskElem := &taskManagerProxy{manager: tskMgr}
	tskMgr.proxy = tskElem
	return tskMgr, tskElem
}

func newCacheManager(ctx core.ServerContext, name string) (*cacheManager, *cacheManagerProxy) {
	cacheMgr := &cacheManager{name: name, registeredCacheNames: make(map[string]string, 10), registeredCaches: make(map[string]components.CacheComponent, 10)}
	cacheElem := &cacheManagerProxy{manager: cacheMgr}
	cacheMgr.proxy = cacheElem
	return cacheMgr, cacheElem
}

func newSessionManager(ctx core.ServerContext, name string, conf config.Config) (*sessionManager, *sessionManagerProxy) {
	sessionMgr := &sessionManager{name: name}
	sessElem := &sessionManagerProxy{manager: sessionMgr}
	sessionMgr.proxy = sessElem
	return sessionMgr, sessElem
}

func childCacheManager(ctx core.ServerContext, name string, parentCacheManager core.ServerElement, filters ...elements.Filter) (*cacheManager, *cacheManagerProxy) {
	cacheMgrProxy := parentCacheManager.(*cacheManagerProxy)
	cacheMgr := cacheMgrProxy.manager
	registeredCaches := make(map[string]components.CacheComponent, len(cacheMgr.registeredCaches))
	registeredCacheNames := make(map[string]string, len(cacheMgr.registeredCacheNames))
	for k, v := range cacheMgr.registeredCaches {
		allowed := true
		for _, filter := range filters {
			if !filter.Allowed(ctx, k) {
				allowed = false
				break
			}
		}
		if allowed {
			registeredCaches[k] = v
		}
	}
	for k, v := range cacheMgr.registeredCacheNames {
		allowed := true
		for _, filter := range filters {
			if !filter.Allowed(ctx, k) {
				allowed = false
				break
			}
		}
		if allowed {
			registeredCacheNames[k] = v
		}
	}
	childcacheMgr := &cacheManager{name: name, registeredCaches: registeredCaches, registeredCacheNames: registeredCacheNames}
	childcacheMgrElem := &cacheManagerProxy{manager: childcacheMgr}
	childcacheMgr.proxy = childcacheMgrElem
	return childcacheMgr, childcacheMgrElem
}
