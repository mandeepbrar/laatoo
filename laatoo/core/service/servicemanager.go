package service

import (
	"laatoo/core/common"
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
	err := svcMgr.createServices(svcmgrInitializeCtx, conf)
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
			svc := svcStruct.service
			err := svc.Start(svcStartCtx)
			if err != nil {
				return errors.WrapError(svcStartCtx, err)
			}
			log.Logger.Info(svcmgrStartCtx, "Started service ", "name", svcname)
			svcChannels, ok := svcStruct.conf.GetSubConfig(config.CONF_ENGINE_CHANNELS)
			if ok {
				channelnames := svcChannels.AllConfigurations()
				for _, channelName := range channelnames {
					svcChannelConfig, err := config.ConfigFileAdapter(svcChannels, channelName)
					svcServeCtx := ctx.NewContextWithElements("Serve: "+svcStruct.name, core.ContextMap{core.ServerElementService: svcStruct, core.ServerElementServiceFactory: svcStruct.factory}, core.ServerElementService)
					err = chanMgr.Serve(svcServeCtx, channelName, svcStruct, svcChannelConfig)
					if err != nil {
						return errors.WrapError(svcServeCtx, err)
					}
					log.Logger.Info(svcmgrStartCtx, "Serving service ", "name", svcname, "channel", channelName)
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
			log.Logger.Debug(ctx, "Process Service group", "groupname", groupname)
			svcgrpConfig, err := config.ConfigFileAdapter(allgroups, groupname)
			if err != nil {
				return errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_CONF, err, "Wrong config for service group", groupname)
			}
			grpCtx := ctx.SubContext("ServiceGroup:" + groupname)
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
			serviceConfig, err := config.ConfigFileAdapter(svcsConf, svcAlias)
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

			svc, err := svcMgr.createService(svcCtx, svcAlias, svcFactory, svcMethod, serviceConfig)
			if err != nil {
				return errors.WrapError(svcCtx, err)
			}
			/*svcMw := createMW(serviceConfig, app.ServiceFactoryMiddleware[factoryAlias])

			if len(*svcMw) > 0 {
				serviceMiddleware[svcAlias] = svcMw
			} else {
			services[svcAlias] = svc
			}*/
			svcMgr.servicesStore[svcAlias] = svc
			log.Logger.Debug(ctx, "Registered service", "service name", svcAlias)
		}
	}
	return nil
}

//create service
func (svcMgr *serviceManager) createService(ctx core.ServerContext, serviceAlias string, factoryname string, serviceMethod string, conf config.Config) (*service, error) {
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
	//pass a server context to service with element set to service
	svcCreationCtx := ctx.NewContextWithElements("Create"+serviceAlias, core.ContextMap{core.ServerElementService: svcStruct, core.ServerElementServiceFactory: facElem}, core.ServerElementService)
	svc, err := factory.CreateService(svcCreationCtx, serviceAlias, serviceMethod)
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