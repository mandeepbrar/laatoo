package service

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/common"
	"laatoo/server/constants"
	slog "laatoo/server/log"
)

const (
	CONF_SERVICEGROUPS = "servicegroups"
	CONF_SERVICES      = "services"
	CONF_FACTORY       = "factory"
	CONF_SERVICEMETHOD = "servicemethod"
)

type serviceManager struct {
	name   string
	parent core.ServerElement
	proxy  server.ServiceManager
	//store for service factory in an application
	servicesStore  map[string]*serviceProxy
	factoryManager server.FactoryManager
}

func (svcMgr *serviceManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	elem := ctx.GetServerElement(core.ServerElementFactoryManager)
	svcMgr.factoryManager = elem.(server.FactoryManager)
	svcmgrInitializeCtx := svcMgr.createContext(ctx, "Initialize service manager ")
	err := svcMgr.createServices(svcmgrInitializeCtx, conf)
	if err != nil {
		return errors.WrapError(svcmgrInitializeCtx, err)
	}
	basedir, _ := ctx.GetString(constants.CONF_BASE_DIR)
	log.Trace(ctx, "*************** Processing service manager", " base directory", basedir)
	if err := common.ProcessDirectoryFiles(svcmgrInitializeCtx, svcMgr.parent, constants.CONF_SERVICES, svcMgr.createService, true); err != nil {
		return errors.WrapError(svcmgrInitializeCtx, err)
	}

	err = svcMgr.initializeServices(svcmgrInitializeCtx)
	if err != nil {
		return errors.WrapError(svcmgrInitializeCtx, err)
	}
	return nil
}

func (svcMgr *serviceManager) Start(ctx core.ServerContext) error {
	svcmgrStartCtx := svcMgr.createContext(ctx, "Start service manager")
	chanMgr := ctx.GetServerElement(core.ServerElementChannelManager).(server.ChannelManager)
	for svcname, svcProxy := range svcMgr.servicesStore {
		if svcProxy.svc.owner == svcMgr {
			log.Debug(svcmgrStartCtx, "Starting service ", "service name", svcname)
			svcStartCtx := svcmgrStartCtx.NewContextWithElements("Start "+svcname, core.ContextMap{core.ServerElementService: svcProxy, core.ServerElementServiceFactory: svcProxy.svc.factory}, core.ServerElementService)
			err := svcProxy.svc.start(svcStartCtx)
			if err != nil {
				return errors.WrapError(svcStartCtx, err)
			}
			log.Info(svcmgrStartCtx, "Started service ", "name", svcname)
		}
	}

	for svcname, svcProxy := range svcMgr.servicesStore {
		if svcProxy.svc.owner == svcMgr {
			svcChannels, ok := svcProxy.svc.conf.GetSubConfig(constants.CONF_ENGINE_CHANNELS)
			if ok {
				channelnames := svcChannels.AllConfigurations()
				for _, channelName := range channelnames {
					svcChannelConfigs, ok := svcChannels.GetConfigArray(channelName)
					if !ok {
						channelConfig, _ := svcChannels.GetSubConfig(channelName)
						svcChannelConfigs = []config.Config{channelConfig}
					}
					svcServeCtx := ctx.NewContextWithElements("Serve: "+svcProxy.svc.name, core.ContextMap{core.ServerElementService: svcProxy, core.ServerElementServiceFactory: svcProxy.svc.factory}, core.ServerElementService)
					for _, conf := range svcChannelConfigs {
						err := chanMgr.Serve(svcServeCtx, channelName, svcProxy, conf)
						if err != nil {
							return errors.WrapError(svcServeCtx, err)
						}
					}
					log.Info(svcmgrStartCtx, "Serving service ", "name", svcname, "channel", channelName)
				}
			}
		}
	}
	return nil
}

//create services within an application
func (svcMgr *serviceManager) createServices(ctx core.ServerContext, conf config.Config) error {
	//get a map of all the services
	allgroups, ok := conf.GetSubConfig(CONF_SERVICEGROUPS)
	if ok {
		groups := allgroups.AllConfigurations()
		for _, groupname := range groups {
			log.Debug(ctx, "Process Service group", "groupname", groupname)
			svcgrpConfig, err, _ := common.ConfigFileAdapter(ctx, allgroups, groupname)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			grpCtx := ctx.SubContext("ServiceGroup:" + groupname)
			//middleware, _ := conf.GetStringArray(config.CONF_MIDDLEWARE)
			err = svcMgr.createServices(grpCtx, svcgrpConfig)
			if err != nil {
				return err
			}
		}
	}

	svcsConf, ok := conf.GetSubConfig(CONF_SERVICES)
	if ok {
		svcAliases := svcsConf.AllConfigurations()
		for _, svcAlias := range svcAliases {
			_, ok := svcMgr.servicesStore[svcAlias]
			if ok {
				continue
			}
			serviceConfig, err, _ := common.ConfigFileAdapter(ctx, svcsConf, svcAlias)
			if err != nil {
				return errors.WrapError(ctx, err)
			}

			svcCtx := ctx.SubContext("Create Service:" + svcAlias)

			err = svcMgr.createService(svcCtx, serviceConfig, svcAlias)
			if err != nil {
				return errors.WrapError(svcCtx, err)
			}
			/*svcMw := createMW(serviceConfig, app.ServiceFactoryMiddleware[factoryAlias])

			if len(*svcMw) > 0 {
				serviceMiddleware[svcAlias] = svcMw
			} else {
			services[svcAlias] = svc
			}*/
			log.Debug(ctx, "Registered service", "service name", svcAlias)
		}
	}
	return nil
}

//create service
func (svcMgr *serviceManager) createService(ctx core.ServerContext, conf config.Config, serviceAlias string) error {
	if !common.CheckContextCondition(ctx, conf) {
		return nil
	}

	factoryname, ok := conf.GetString(CONF_FACTORY)
	if !ok {
		factoryname = common.CONF_DEFAULTFACTORY_NAME
	}

	serviceMethod, ok := conf.GetString(CONF_SERVICEMETHOD)
	if !ok {
		log.Debug(ctx, "No service method provided for service", "Service", serviceAlias)
	}

	//get the factory from factory manager
	facElem, err := svcMgr.factoryManager.GetFactory(ctx, factoryname)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	//get the factory from proxy
	factory := facElem.Factory()

	//proxy for the service
	svcStruct := &service{name: serviceAlias, conf: conf, owner: svcMgr, factory: facElem}

	svcProxy := &serviceProxy{svc: svcStruct}

	svcCtx := svcMgr.getServiceContext(ctx)

	parentMw, ok := svcCtx.GetStringArray(constants.CONF_MIDDLEWARE)
	/*if ok {
		if grpMw != nil {
			parentMw = append(parentMw, grpMw...)
		}
	} else {
		parentMw = grpMw
	}*/
	middleware, ok := conf.GetStringArray(constants.CONF_MIDDLEWARE)
	if ok {
		if parentMw != nil {
			middleware = append(parentMw, middleware...)
		}
	}
	if middleware != nil {
		svcCtx.Set(constants.CONF_MIDDLEWARE, middleware)
	}

	cacheToUse, ok := conf.GetString(constants.CONF_CACHE_NAME)
	if ok {
		svcCtx.Set("__cache", cacheToUse)
		log.Error(ctx, "Setting cache for service ", "cacheToUse", cacheToUse)
	}

	elem := ctx.GetServerElement(core.ServerElementLogger)
	_, logger := slog.ChildLoggerWithConf(ctx, serviceAlias, elem.(server.Logger), svcProxy, conf)
	svcCtx.Set("__logger", logger)

	//pass a server context to service with element set to service
	svcCreationCtx := ctx.NewContextWithElements("Create: "+serviceAlias, core.ContextMap{core.ServerElementService: svcProxy, core.ServerElementServiceFactory: facElem}, core.ServerElementService)
	log.Trace(svcCreationCtx, "Creating service", "service name", serviceAlias, "method", serviceMethod, "factory", factoryname)
	svc, err := factory.CreateService(svcCreationCtx, serviceAlias, serviceMethod, conf)
	if err != nil {
		return errors.WrapError(svcCreationCtx, err)
	}
	if svc == nil {
		return errors.ThrowError(svcCreationCtx, errors.CORE_ERROR_MISSING_SERVICE, "Name", serviceAlias)
	}
	svcStruct.service = svc

	_, ok = svcMgr.servicesStore[serviceAlias]
	if ok {
		return errors.ThrowError(svcCreationCtx, errors.CORE_ERROR_BAD_CONF, "Error", "Service with this alias already exists")
	}
	svcMgr.servicesStore[serviceAlias] = svcProxy

	log.Trace(svcCreationCtx, "Created service", "service name", serviceAlias)

	return nil
}

//initialize services within an application
func (svcMgr *serviceManager) initializeServices(ctx core.ServerContext) error {
	for svcname, svcProxy := range svcMgr.servicesStore {
		if svcProxy.svc.owner == svcMgr {
			svcInitializeCtx := ctx.NewContextWithElements("Initialize: "+svcname, core.ContextMap{core.ServerElementService: svcProxy, core.ServerElementServiceFactory: svcProxy.svc.factory}, core.ServerElementService)
			log.Debug(svcInitializeCtx, "Initializing service", "service name", svcname)
			svc := svcProxy.svc.service
			err := svc.Initialize(svcInitializeCtx, svcProxy.svc.conf)
			if err != nil {
				return errors.WrapError(svcInitializeCtx, err)
			}
			log.Trace(svcInitializeCtx, "Initialized service")
		}
	}
	return nil
}

func (svcMgr *serviceManager) getServiceContext(ctx core.ServerContext) core.ServerContext {
	return ctx
}

//creates a context specific to service manager
func (svcMgr *serviceManager) createContext(ctx core.ServerContext, name string) core.ServerContext {
	return ctx.NewContextWithElements(name,
		core.ContextMap{core.ServerElementServiceManager: svcMgr.proxy,
			core.ServerElementFactoryManager: svcMgr.factoryManager}, core.ServerElementServiceManager)
}
