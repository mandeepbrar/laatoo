package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/common"
	"laatoo/server/constants"
)

const (
	CONF_SERVICEGROUPS = "servicegroups"
	CONF_FACTORY       = "factory"
	CONF_SERVICEMETHOD = "servicemethod"
)

type serviceManager struct {
	name   string
	parent core.ServerElement
	proxy  elements.ServiceManager
	//store for service factory in an application
	servicesStore  map[string]*serviceProxy
	factoryManager elements.FactoryManager
}

func (svcMgr *serviceManager) Initialize(ctx core.ServerContext, conf config.Config) error {

	modManager := ctx.GetServerElement(core.ServerElementModuleManager).(*moduleManagerProxy).modMgr

	if err := modManager.loadServices(ctx, svcMgr.createService); err != nil {
		return err
	}

	err := svcMgr.createServices(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	basedir, _ := ctx.GetString(config.BASEDIR)
	log.Trace(ctx, "*************** Processing service manager", " base directory", basedir)
	err = svcMgr.processServicesFromFolder(ctx, basedir)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	err = svcMgr.initializeServices(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (svcMgr *serviceManager) Start(ctx core.ServerContext) error {
	//	svcmgrStartCtx := ctx.(*serverContext)
	for _, svcProxy := range svcMgr.servicesStore {
		if svcProxy.svc.owner == svcMgr {
			err := svcMgr.startService(ctx, svcProxy)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (svcMgr *serviceManager) startModuleInstanceServices(ctx core.ServerContext, mod *serverModule) error {
	for svcname, _ := range mod.services {
		svc, _ := svcMgr.servicesStore[svcname]
		if err := svcMgr.startService(ctx, svc); err != nil {
			return err
		}
	}
	return nil
}

func (svcMgr *serviceManager) startService(ctx core.ServerContext, svcProxy *serviceProxy) error {
	svcStartCtx := svcProxy.svc.svrContext.subContext("Start " + svcProxy.svc.name)
	log.Debug(svcStartCtx, "Starting service ", "service name", svcProxy.svc.name)
	//, core.ContextMap{core.ServerElementService: svcProxy, core.ServerElementServiceFactory: svcProxy.svc.factory}, core.ServerElementService
	err := svcProxy.svc.start(svcStartCtx)
	if err != nil {
		return errors.WrapError(svcStartCtx, err)
	}
	log.Info(svcStartCtx, "Started service ", "name", svcProxy.svc.name)
	return nil
}

func (svcMgr *serviceManager) processServicesFromFolder(ctx core.ServerContext, folder string) error {
	objs, err := svcMgr.loadServicesFromFolder(ctx, folder)
	if err != nil {
		return err
	}

	if err = common.ProcessObjects(ctx, objs, svcMgr.createService); err != nil {
		return err
	}
	return nil
}

func (svcMgr *serviceManager) loadServicesFromFolder(ctx core.ServerContext, folder string) (map[string]config.Config, error) {
	return common.ProcessDirectoryFiles(ctx, folder, constants.CONF_SERVICES, true)
}

func (svcMgr *serviceManager) getService(ctx core.ServerContext, serviceName string) (elements.Service, error) {
	elem, ok := svcMgr.servicesStore[serviceName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, "Service Alias", serviceName)
	}
	return elem, nil
}

//create services within an application
func (svcMgr *serviceManager) createServices(ctx core.ServerContext, conf config.Config) error {
	//get a map of all the services
	allgroups, ok := conf.GetSubConfig(ctx, CONF_SERVICEGROUPS)
	if ok {
		groups := allgroups.AllConfigurations(ctx)
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

	svcsConf, ok := conf.GetSubConfig(ctx, constants.CONF_SERVICES)
	if ok {
		svcAliases := svcsConf.AllConfigurations(ctx)
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
	svcCreateCtx := ctx.(*serverContext)

	_, ok := svcMgr.servicesStore[serviceAlias]
	if ok {
		//return errors.ThrowError(svcCtx, errors.CORE_ERROR_BAD_CONF, "Error", "Service with this alias already exists")
		return nil
	}

	factoryname, factoryok := conf.GetString(ctx, CONF_FACTORY)
	if !factoryok {
		factoryname = common.CONF_DEFAULTFACTORY_NAME
	}

	serviceMethod, ok := conf.GetString(ctx, CONF_SERVICEMETHOD)
	if !ok {
		log.Debug(ctx, "No service method provided for service", "Service", serviceAlias)
	}

	//get the factory from factory manager
	facElem, err := svcMgr.factoryManager.GetFactory(ctx, factoryname)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	svcfactoryProxy := facElem.(*serviceFactoryProxy)
	facCtx := svcfactoryProxy.fac.svrContext

	var svcCtx *serverContext

	mod := svcCreateCtx.GetServerElement(core.ServerElementModule)

	if mod != nil {
		svcCtx = mod.(*moduleProxy).mod.svrContext.newContext("Service: " + serviceAlias)
	} else {
		//get the factory from proxy
		svcCtx = facCtx.newContext("Service: " + serviceAlias)
	}

	//log.Trace(ctx, "levels", "factory", facCtx.level, "server", svcCreateCtx.level)
	//use the latest context... i.e. server.. environment or application....
	//if factory is from earlier level then override elements with latest context
	if facCtx.level <= svcCreateCtx.level {
		//	log.Trace(ctx, "factory from a lower level than context")
		cmap := svcCreateCtx.getElementsContextMap()
		svcCtx.setElementReferences(cmap, true)
	}

	factory := svcfactoryProxy.Factory()
	//proxy for the service
	svcStruct := &serverService{name: serviceAlias, conf: conf, owner: svcMgr, factory: facElem, svrContext: svcCtx}
	if !factoryok {
		svcStruct.objectName = serviceMethod
	}

	svcProxy := &serviceProxy{svc: svcStruct}
	svcCtx.setElements(core.ContextMap{core.ServerElementService: svcProxy})

	/*if ok {
		if grpMw != nil {
			parentMw = append(parentMw, grpMw...)
		}
	} else {
		parentMw = grpMw
	}*/
	common.SetupMiddleware(svcCtx, conf)

	cacheToUse, ok := conf.GetString(ctx, constants.CONF_CACHE_NAME)
	if ok {
		svcCtx.Set("__cache", cacheToUse)
		log.Info(svcCtx, "Setting cache for service ", "cacheToUse", cacheToUse)
	}

	if err := processLogging(svcCtx, conf, serviceAlias); err != nil {
		return errors.WrapError(svcCtx, err)
	}

	//pass a server context to service with element set to service
	//, core.ContextMap{core.ServerElementService: svcProxy, core.ServerElementServiceFactory: facElem}, core.ServerElementService
	log.Trace(svcCtx, "Creating service", "service name", serviceAlias, "method", serviceMethod, "factory", factoryname)
	svc, err := factory.CreateService(svcCtx, serviceAlias, serviceMethod, conf)
	if err != nil {
		return errors.WrapError(svcCtx, err)
	}
	if svc == nil {
		return errors.ThrowError(svcCtx, errors.CORE_ERROR_MISSING_SERVICE, "Name", serviceAlias)
	}
	svcStruct.service = svc

	if err := svcStruct.loadMetaData(svcCtx); err != nil {
		return errors.WrapError(svcCtx, err)
	}

	svcMgr.servicesStore[serviceAlias] = svcProxy

	log.Trace(svcCtx, "Created service", "service name", serviceAlias)

	return nil
}

//initialize services within an application
func (svcMgr *serviceManager) initializeServices(ctx core.ServerContext) error {
	for svcname, svcProxy := range svcMgr.servicesStore {
		if svcProxy.svc.owner == svcMgr {
			//, core.ContextMap{core.ServerElementService: svcProxy, core.ServerElementServiceFactory: svcProxy.svc.factory}, core.ServerElementService
			svcInitializeCtx := svcProxy.svc.svrContext.SubContext("Initialize: " + svcname)
			log.Debug(svcInitializeCtx, "Initializing service", "service name", svcname)
			//svc := svcProxy.svc.service
			err := svcProxy.svc.initialize(svcInitializeCtx, svcProxy.svc.conf)
			if err != nil {
				return errors.WrapError(svcInitializeCtx, err)
			}
			log.Trace(svcInitializeCtx, "Initialized service")
		}
	}
	return nil
}

func (svcMgr *serviceManager) unloadModuleServices(ctx core.ServerContext, mod *serverModule) error {
	ctx = ctx.SubContext("unload services")
	if err := common.ProcessObjects(ctx, mod.services, svcMgr.unloadService); err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (svcMgr *serviceManager) unloadService(ctx core.ServerContext, conf config.Config, svcName string) error {
	unloadSvc := ctx.SubContext("Unload service")
	svcprxy, ok := svcMgr.servicesStore[svcName]
	if ok {
		err := svcprxy.svc.stop(unloadSvc)
		if err != nil {
			log.Error(unloadSvc, "Error while stopping service", "err", err)
		}
		err = svcprxy.svc.unload(unloadSvc)
		if err != nil {
			log.Error(unloadSvc, "Error while stopping service", "err", err)
		}
		delete(svcMgr.servicesStore, svcName)
	}
	return nil
}
