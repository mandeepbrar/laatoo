package service

import (
	"laatoo/framework/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
)

const (
	CONF_SERVICEGROUPS = "servicegroups"
	CONF_SERVICES      = "services"
	CONF_FACTORY       = "factory"
	CONF_SERVICEMETHOD = "servicemethod"
)

type serviceManager struct {
	parent core.ServerElement
	proxy  server.ServiceManager
	//store for service factory in an application
	servicesStore  map[string]*service
	factoryManager server.FactoryManager
}

func (svcMgr *serviceManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	elem := ctx.GetServerElement(core.ServerElementFactoryManager)
	svcMgr.factoryManager = elem.(server.FactoryManager)
	svcmgrInitializeCtx := svcMgr.createContext(ctx, "Initialize service manager")
	err := svcMgr.createServices(svcmgrInitializeCtx, conf, nil)
	if err != nil {
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
	for svcname, svcStruct := range svcMgr.servicesStore {
		if svcStruct.owner == svcMgr {
			log.Logger.Debug(svcmgrStartCtx, "Starting service", "service name", svcname)
			svcStartCtx := svcmgrStartCtx.NewContextWithElements("Start "+svcname, core.ContextMap{core.ServerElementService: svcStruct, core.ServerElementServiceFactory: svcStruct.factory}, core.ServerElementService)
			err := svcStruct.start(svcStartCtx)
			if err != nil {
				return errors.WrapError(svcStartCtx, err)
			}
			log.Logger.Info(svcmgrStartCtx, "Started service ", "name", svcname)
			svcChannels, ok := svcStruct.conf.GetSubConfig(config.CONF_ENGINE_CHANNELS)
			if ok {
				channelnames := svcChannels.AllConfigurations()
				for _, channelName := range channelnames {
					svcChannelConfigs, ok := svcChannels.GetConfigArray(channelName)
					if !ok {
						channelConfig, _ := svcChannels.GetSubConfig(channelName)
						svcChannelConfigs = []config.Config{channelConfig}
					}
					svcServeCtx := ctx.NewContextWithElements("Serve: "+svcStruct.name, core.ContextMap{core.ServerElementService: svcStruct, core.ServerElementServiceFactory: svcStruct.factory}, core.ServerElementService)
					for _, conf := range svcChannelConfigs {
						err = chanMgr.Serve(svcServeCtx, channelName, svcStruct, conf)
						if err != nil {
							return errors.WrapError(svcServeCtx, err)
						}
					}
					log.Logger.Info(svcmgrStartCtx, "Serving service ", "name", svcname, "channel", channelName)
				}
			}
		}
	}
	return nil
}

//create services within an application
func (svcMgr *serviceManager) createServices(ctx core.ServerContext, conf config.Config, grpMw []string) error {
	//get a map of all the services
	allgroups, ok := conf.GetSubConfig(CONF_SERVICEGROUPS)
	if ok {
		groups := allgroups.AllConfigurations()
		for _, groupname := range groups {
			log.Logger.Debug(ctx, "Process Service group", "groupname", groupname)
			svcgrpConfig, err, _ := common.ConfigFileAdapter(ctx, allgroups, groupname)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			grpCtx := ctx.SubContext("ServiceGroup:" + groupname)
			middleware, _ := conf.GetStringArray(config.CONF_MIDDLEWARE)
			err = svcMgr.createServices(grpCtx, svcgrpConfig, middleware)
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

			svcFactory, ok := serviceConfig.GetString(CONF_FACTORY)
			if !ok {
				return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Service", svcAlias, "Conf", CONF_FACTORY)
			}

			svcMethod, ok := serviceConfig.GetString(CONF_SERVICEMETHOD)
			if !ok {
				return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Service", svcAlias, "Conf", CONF_SERVICEMETHOD)
			}

			svc, err := svcMgr.createService(svcCtx, svcAlias, svcFactory, svcMethod, serviceConfig, grpMw)
			if err != nil {
				return errors.WrapError(svcCtx, err)
			}
			/*svcMw := createMW(serviceConfig, app.ServiceFactoryMiddleware[factoryAlias])

			if len(*svcMw) > 0 {
				serviceMiddleware[svcAlias] = svcMw
			} else {
			services[svcAlias] = svc
			}*/
			_, ok = svcMgr.servicesStore[svcAlias]
			if ok {
				return errors.ThrowError(svcCtx, errors.CORE_ERROR_BAD_CONF, "Error", "Service with this alias already exists")
			}
			svcMgr.servicesStore[svcAlias] = svc
			log.Logger.Debug(ctx, "Registered service", "service name", svcAlias)
		}
	}
	return nil
}

//create service
func (svcMgr *serviceManager) createService(ctx core.ServerContext, serviceAlias string, factoryname string, serviceMethod string, conf config.Config, grpMw []string) (*service, error) {
	//get the factory from factory manager
	facElem, err := svcMgr.factoryManager.GetFactory(ctx, factoryname)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	//get the factory from proxy
	factory := facElem.Factory()

	//create a subcontext from proxy
	//service subcontext will be a child of its factory
	svcElemCtx := facElem.NewCtx(serviceAlias)
	//proxy for the service
	svcStruct := &service{Context: svcElemCtx.(*common.Context), name: serviceAlias, conf: conf, owner: svcMgr, factory: facElem}

	parentMw, ok := facElem.GetStringArray(config.CONF_MIDDLEWARE)
	if ok {
		if grpMw != nil {
			parentMw = append(parentMw, grpMw...)
		}
	} else {
		parentMw = grpMw
	}
	middleware, ok := conf.GetStringArray(config.CONF_MIDDLEWARE)
	if ok {
		if parentMw != nil {
			middleware = append(parentMw, middleware...)
		}
	}
	if middleware != nil {
		svcElemCtx.Set(config.CONF_MIDDLEWARE, middleware)
	}

	cacheToUse, ok := conf.GetString(config.CONF_CACHE_NAME)
	if ok {
		svcStruct.Set("__cache", cacheToUse)
		log.Logger.Error(ctx, "Setting cache for service ", "cacheToUse", cacheToUse)
	}

	//pass a server context to service with element set to service
	svcCreationCtx := ctx.NewContextWithElements("Create"+serviceAlias, core.ContextMap{core.ServerElementService: svcStruct, core.ServerElementServiceFactory: facElem}, core.ServerElementService)
	log.Logger.Trace(ctx, "Creating service", "service name", serviceAlias, "method", serviceMethod, "factory", factoryname)
	svc, err := factory.CreateService(svcCreationCtx, serviceAlias, serviceMethod, conf)
	if err != nil {
		return nil, errors.WrapError(svcCreationCtx, err)
	}
	if svc == nil {
		return nil, errors.ThrowError(svcCreationCtx, errors.CORE_ERROR_MISSING_SERVICE, "Name", serviceAlias)
	}
	svcStruct.service = svc
	return svcStruct, nil
}

//initialize services within an application
func (svcMgr *serviceManager) initializeServices(ctx core.ServerContext) error {
	for svcname, svcStruct := range svcMgr.servicesStore {
		if svcStruct.owner == svcMgr {
			log.Logger.Debug(ctx, "Initializing service", "service name", svcname)
			svcInitializeCtx := ctx.NewContextWithElements("Initialize"+svcname, core.ContextMap{core.ServerElementService: svcStruct, core.ServerElementServiceFactory: svcStruct.factory}, core.ServerElementService)
			svc := svcStruct.service
			log.Logger.Trace(ctx, "Initializing service", "conf", svcStruct.conf)
			err := svc.Initialize(svcInitializeCtx, svcStruct.conf)
			if err != nil {
				return errors.WrapError(svcInitializeCtx, err)
			}
		}
	}
	return nil
}

//creates a context specific to service manager
func (svcMgr *serviceManager) createContext(ctx core.ServerContext, name string) core.ServerContext {
	return ctx.NewContextWithElements(name,
		core.ContextMap{core.ServerElementServiceManager: svcMgr.proxy,
			core.ServerElementFactoryManager: svcMgr.factoryManager}, core.ServerElementServiceManager)
}
