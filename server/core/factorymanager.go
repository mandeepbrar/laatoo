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
	CONF_SERVICEFACTORYGROUPS = "factorygroups"

	CONF_SERVICEFACTORYCONFIG = "config"
)

type factoryManager struct {
	name   string
	svrref *abstractserver
	parent core.ServerElement
	proxy  elements.FactoryManager
	//store for service factory in an application
	serviceFactoryStore map[string]*serviceFactoryProxy
}

func (facMgr *factoryManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	err := facMgr.createServiceFactories(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	c := ctx.CreateConfig()
	c.SetVals(ctx, map[string]interface{}{constants.CONF_SERVICEFACTORY: common.CONF_DEFAULTFACTORY_NAME})
	err = facMgr.createServiceFactory(ctx, c, common.CONF_DEFAULTFACTORY_NAME)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	baseDir, _ := ctx.GetString(config.BASEDIR)

	modManager := ctx.GetServerElement(core.ServerElementModuleManager).(*moduleManagerProxy).modMgr

	if err = modManager.loadFactories(ctx, facMgr.createServiceFactory); err != nil {
		return err
	}

	if err := facMgr.processFactoriesFromFolder(ctx, baseDir); err != nil {
		return errors.WrapError(ctx, err)
	}

	return facMgr.initializeFactories(ctx)
}

func (facMgr *factoryManager) Start(ctx core.ServerContext) error {
	facmgrStartCtx := ctx.SubContext("Start factory manager")
	for facname, facProxy := range facMgr.serviceFactoryStore {
		log.Debug(facmgrStartCtx, "Starting factory", "factory name", facname)
		facStruct := facProxy.fac
		if facStruct.owner == facMgr {
			err := facMgr.startFactory(facmgrStartCtx, facProxy)
			if err != nil {
				return errors.WrapError(facmgrStartCtx, err)
			}
		}
	}
	return nil
}

func (facMgr *factoryManager) startModuleInstanceFactories(ctx core.ServerContext, mod *serverModule) error {
	for facname, _ := range mod.factories {
		fac, _ := facMgr.serviceFactoryStore[facname]
		if err := facMgr.startFactory(ctx, fac); err != nil {
			return err
		}
	}
	return nil
}

func (facMgr *factoryManager) startFactory(ctx core.ServerContext, facProxy *serviceFactoryProxy) error {
	facStartCtx := facProxy.fac.svrContext.subContext("Start " + facProxy.fac.name)
	log.Debug(facStartCtx, "Starting factory ", "factory name", facProxy.fac.name)
	err := facProxy.fac.start(facStartCtx)
	if err != nil {
		return errors.WrapError(facStartCtx, err)
	}
	log.Info(facStartCtx, "Started factory ", "name", facProxy.fac.name)

	//, core.ContextMap{core.ServerElementService: svcProxy, core.ServerElementServiceFactory: svcProxy.svc.factory}, core.ServerElementService
	return nil
}

func (facMgr *factoryManager) processFactoriesFromFolder(ctx core.ServerContext, folder string) error {
	objs, err := facMgr.loadFactoriesFromFolder(ctx, folder)
	if err != nil {
		return err
	}

	if err = common.ProcessObjects(ctx, objs, facMgr.createServiceFactory); err != nil {
		return err
	}
	return nil
}

func (facMgr *factoryManager) loadFactoriesFromFolder(ctx core.ServerContext, folder string) (map[string]config.Config, error) {
	return common.ProcessDirectoryFiles(ctx, folder, constants.CONF_SERVICEFACTORIES, true)
}

func (facMgr *factoryManager) createServiceFactories(ctx core.ServerContext, conf config.Config) error {
	//get a map of all the services
	allgroups, ok := conf.GetSubConfig(ctx, CONF_SERVICEFACTORYGROUPS)
	if ok {
		groups := allgroups.AllConfigurations(ctx)
		for _, groupname := range groups {
			log.Trace(ctx, "Process Service Factory group", "groupname", groupname)
			facgrpConfig, err, _ := common.ConfigFileAdapter(ctx, allgroups, groupname)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			grpCtx := ctx.SubContext("Group:" + groupname)
			err = facMgr.createServiceFactories(grpCtx, facgrpConfig)
			if err != nil {
				return err
			}
		}
	}
	//get a map of all the services
	factoriesConfig, ok := conf.GetSubConfig(ctx, constants.CONF_SERVICEFACTORIES)
	if !ok {
		return nil
	}
	factories := factoriesConfig.AllConfigurations(ctx)
	for _, factoryName := range factories {
		log.Trace(ctx, "Process Factory ", "Factory name", factoryName)
		factoryConfig, err, _ := common.ConfigFileAdapter(ctx, factoriesConfig, factoryName)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		facCtx := ctx.SubContext("Create Factory:" + factoryName)
		err = facMgr.createServiceFactory(facCtx, factoryConfig, factoryName)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}

func (facMgr *factoryManager) createServiceFactory(ctx core.ServerContext, factoryConfig config.Config, factoryAlias string) error {
	ctx = ctx.SubContext("Create Service Factory")
	factoryName, ok := factoryConfig.GetString(ctx, constants.CONF_SERVICEFACTORY)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Wrong config for Factory Name", factoryAlias, "Missing Config", constants.CONF_SERVICEFACTORY)
	}

	_, ok = facMgr.serviceFactoryStore[factoryAlias]
	if ok {
		return nil
	}

	log.Trace(ctx, "Creating factory object", "Name", factoryName)
	factoryInt, err := ctx.CreateObject(factoryName)

	if err != nil {
		return errors.WrapError(ctx, err)
	}
	factory, ok := factoryInt.(core.ServiceFactory)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Object is not factory for Factory Name", factoryName)
	}

	if factory != nil {

		var facCtx *serverContext

		mod := ctx.GetServerElement(core.ServerElementModule)

		if mod != nil {
			log.Trace(ctx, "Creating factory object for module", "facmgr", facMgr)
			facCtx = mod.(*moduleProxy).mod.svrContext.newContext("Factory: " + factoryAlias)
		} else {
			//derivce new context from abstract server context
			facCtx = facMgr.svrref.svrContext.newContext("Factory: " + factoryAlias)
		}

		if err := processLogging(facCtx, factoryConfig, factoryAlias); err != nil {
			return errors.WrapError(facCtx, err)
		}

		fac := &serviceFactory{name: factoryAlias, objectName: factoryName, factory: factory, owner: facMgr, conf: factoryConfig, svrContext: facCtx}
		facProxy := &serviceFactoryProxy{fac: fac}
		facCtx.setElements(core.ContextMap{core.ServerElementServiceFactory: facProxy})

		common.SetupMiddleware(facCtx, factoryConfig)

		cacheToUse, ok := factoryConfig.GetString(facCtx, constants.CONF_CACHE_NAME)
		if ok {
			facCtx.Set("__cache", cacheToUse)
		}

		if err := fac.loadMetaData(facCtx); err != nil {
			return errors.WrapError(facCtx, err)
		}

		//add the service to the application
		facMgr.serviceFactoryStore[factoryAlias] = facProxy
		log.Trace(ctx, "factory store in manager", "serviceFactoryStore", facMgr.serviceFactoryStore)

	}
	return nil
}

//initialize services within an application
func (facMgr *factoryManager) initializeFactories(ctx core.ServerContext) error {
	for facname, facProxy := range facMgr.serviceFactoryStore {
		facStruct := facProxy.fac
		if facStruct.owner == facMgr {
			facInitializeCtx := facStruct.svrContext.SubContext("Initialize: " + facname)
			log.Debug(facInitializeCtx, "Initializing factory", "factory name", facname)
			err := facStruct.initialize(facInitializeCtx, facStruct.conf)
			if err != nil {
				return errors.WrapError(facInitializeCtx, err)
			}
			log.Trace(facInitializeCtx, "Initialized factory", "factory name", facname)
		}
	}
	return nil
}

func (facMgr *factoryManager) unloadModuleFactories(ctx core.ServerContext, mod *serverModule) error {
	ctx = ctx.SubContext("unload factories")
	if err := common.ProcessObjects(ctx, mod.factories, facMgr.unloadFactory); err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (facMgr *factoryManager) unloadFactory(ctx core.ServerContext, conf config.Config, fac string) error {
	unloadfac := ctx.SubContext("Unload factory")
	facprxy, ok := facMgr.serviceFactoryStore[fac]
	if ok {
		err := facprxy.fac.stop(unloadfac)
		if err != nil {
			log.Error(unloadfac, "Error while stopping factory", "err", err)
		}
		err = facprxy.fac.unload(unloadfac)
		if err != nil {
			log.Error(unloadfac, "Error while stopping factory", "err", err)
		}
		delete(facMgr.serviceFactoryStore, fac)
	}
	return nil
}
