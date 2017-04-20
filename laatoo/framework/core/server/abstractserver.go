package server

import (
	"laatoo/framework/core/common"
	"laatoo/framework/core/factory"
	"laatoo/framework/core/objects"
	"laatoo/framework/core/service"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
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

	parent *abstractserver

	proxy core.ServerElement
}

func newAbstractServer(svrCtx *serverContext, name string, parent *abstractserver, proxy core.ServerElement, filterConf config.Config) *abstractserver {
	as := &abstractserver{name: name, parent: parent, proxy: proxy}
	if parent == nil {
		objCreateCtx := svrCtx.SubContext("Create Object Loader")
		objectLoaderHandle, objectLoader := objects.NewObjectLoader(objCreateCtx, "Root", proxy)
		as.objectLoaderHandle = objectLoaderHandle
		as.objectLoader = objectLoader.(server.ObjectLoader)

		fmCreateCtx := svrCtx.SubContext("Create Factory Manager")
		factoryManagerHandle, factoryManager := factory.NewFactoryManager(fmCreateCtx, "Root", proxy)
		as.factoryManagerHandle = factoryManagerHandle
		as.factoryManager = factoryManager.(server.FactoryManager)

		smCreateCtx := svrCtx.SubContext("Create Service Manager")
		serviceManagerHandle, serviceManager := service.NewServiceManager(smCreateCtx, "Root", proxy)
		as.serviceManagerHandle = serviceManagerHandle
		as.serviceManager = serviceManager.(server.ServiceManager)

		cmCreateCtx := svrCtx.SubContext("Create Channel Manager")
		channelMgrHandle, channelMgr := newChannelManager(cmCreateCtx, "Root", proxy)
		as.channelMgr = channelMgr
		as.channelMgrHandle = channelMgrHandle
	} else {
		loader := parent.objectLoader
		factoryManager := parent.factoryManager
		serviceManager := parent.serviceManager
		channelMgr := parent.channelMgr
		childLoaderHandle, childLoader := objects.ChildLoader(svrCtx, name, loader, proxy)
		as.objectLoaderHandle = childLoaderHandle
		as.objectLoader = childLoader.(server.ObjectLoader)
		childFactoryManagerHandle, childFactoryManager := factory.ChildFactoryManager(svrCtx, name, factoryManager, proxy)
		as.factoryManagerHandle = childFactoryManagerHandle
		as.factoryManager = childFactoryManager.(server.FactoryManager)
		childServiceManagerHandle, childServiceManager := service.ChildServiceManager(svrCtx, name, serviceManager, proxy)
		as.serviceManagerHandle = childServiceManagerHandle
		as.serviceManager = childServiceManager.(server.ServiceManager)
		childChanMgrHandle, childChannelMgr := childChannelManager(svrCtx, name, channelMgr, proxy)
		as.channelMgrHandle = childChanMgrHandle
		as.channelMgr = childChannelMgr.(server.ChannelManager)

	}
	return as
}

//initialize application with object loader, factory manager, service manager
func (as *abstractserver) initialize(ctx *serverContext, conf config.Config) error {

	err := as.initializeSecurityHandler(ctx, conf)
	if err != nil {
		return err
	}
	log.Logger.Trace(ctx, "Initialized security handler")
	ctx.securityHandler = as.securityHandler

	objinit := ctx.SubContext("Initialize object loader")
	err = initializeObjectLoader(objinit, conf, as.objectLoaderHandle)
	if err != nil {
		return err
	}
	log.Logger.Trace(objinit, "Initialized object loader")

	chaninit := ctx.SubContextWithElement("Initialize channel manager", core.ServerElementChannelManager)
	err = initializeChannelManager(chaninit, conf, as.channelMgrHandle)
	if err != nil {
		return err
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

func (as *abstractserver) initializeSecurityHandler(ctx *serverContext, conf config.Config) error {
	secConf, err, ok := common.ConfigFileAdapter(ctx, conf, config.CONF_SECURITY)
	if err != nil {
		return err
	}
	if ok {
		secCtx := ctx.SubContext("Initialize Security Handler")
		shElem, sh := newSecurityHandler(secCtx, "Security Handler:"+as.name, as.proxy)
		err := shElem.Initialize(secCtx, secConf)
		if err != nil {
			return errors.WrapError(secCtx, err)
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

//creates a context specific to environment
func (as *abstractserver) contextMap(ctx core.ServerContext, name string) core.ContextMap {
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
