package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/common"
	"laatoo/server/constants"
)

const (
	CONF_SERVICEFACTORYGROUPS = "factorygroups"
	CONF_SERVICEFACTORIES     = "factories"
	CONF_SERVICEFACTORY       = "factory"
	CONF_SERVICEFACTORYCONFIG = "config"
)

type factoryManager struct {
	name   string
	svrref *abstractserver
	parent core.ServerElement
	proxy  server.FactoryManager
	//store for service factory in an application
	serviceFactoryStore map[string]*serviceFactoryProxy
}

func (facMgr *factoryManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	err := facMgr.createServiceFactories(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	err = facMgr.createServiceFactory(ctx, &common.GenericConfig{CONF_SERVICEFACTORY: common.CONF_DEFAULTFACTORY_NAME}, common.CONF_DEFAULTFACTORY_NAME)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	baseDir, _ := ctx.GetString(constants.CONF_BASE_DIR)

	if err := facMgr.loadFactoriesFromFolder(ctx, baseDir); err != nil {
		return errors.WrapError(ctx, err)
	}

	return facMgr.initializeFactories(ctx)
}

func (facMgr *factoryManager) Start(ctx core.ServerContext) error {
	facmgrStartCtx := ctx.SubContext("Start factory manager")
	for facname, facProxy := range facMgr.serviceFactoryStore {
		facStruct := facProxy.fac
		if facStruct.owner == facMgr {
			log.Debug(ctx, "Starting factory", "factory name", facname)
			facStartCtx := facmgrStartCtx.SubContext("Start" + facname).(*serverContext)
			err := facStruct.factory.Start(facStartCtx)
			if err != nil {
				return errors.WrapError(facStartCtx, err)
			}
		}
	}
	return nil
}

func (facMgr *factoryManager) loadFactoriesFromFolder(ctx core.ServerContext, baseDir string) error {
	return common.ProcessDirectoryFiles(ctx, baseDir, constants.CONF_FACTORIES, facMgr.createServiceFactory, true)
}

func (facMgr *factoryManager) createServiceFactories(ctx core.ServerContext, conf config.Config) error {
	//get a map of all the services
	allgroups, ok := conf.GetSubConfig(CONF_SERVICEFACTORYGROUPS)
	if ok {
		groups := allgroups.AllConfigurations()
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
	factoriesConfig, ok := conf.GetSubConfig(CONF_SERVICEFACTORIES)
	if !ok {
		return nil
	}
	factories := factoriesConfig.AllConfigurations()
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
	factoryAlias = common.FillVariables(ctx, factoryAlias)
	factoryName, ok := factoryConfig.GetString(CONF_SERVICEFACTORY)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Wrong config for Factory Name", factoryAlias, "Missing Config", CONF_SERVICEFACTORY)
	}

	/*svcfacConfig, err, ok := config.ConfigFileAdapter(factoryConfig, CONF_SERVICEFACTORYCONFIG)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if !ok {
		svcfacConfig = make(config.GenericConfig, 0)
	}*/

	//facmw := createMW(svcfacConfig, app.middleware)

	//app.ServiceFactoryMiddleware[factoryName] = facmw

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
			facCtx = mod.(*module).svrContext.newContext("Factory: " + factoryAlias)
		} else {
			//derivce new context from abstract server context
			facCtx = facMgr.svrref.svrContext.newContext("Factory: " + factoryAlias)
		}

		if err := processLogging(facCtx, factoryConfig, factoryAlias); err != nil {
			return errors.WrapError(facCtx, err)
		}

		fac := &serviceFactory{name: factoryAlias, factory: factory, owner: facMgr, conf: factoryConfig, svrContext: facCtx}
		facProxy := &serviceFactoryProxy{fac: fac}
		facCtx.setElements(core.ContextMap{core.ServerElementServiceFactory: facProxy})

		common.SetupMiddleware(facCtx, factoryConfig)

		cacheToUse, ok := factoryConfig.GetString(constants.CONF_CACHE_NAME)
		if ok {
			facCtx.Set("__cache", cacheToUse)
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
			facInitializeCtx := ctx.SubContext("Initialize " + facname)
			log.Debug(facInitializeCtx, "Initializing factory", "factory name", facname)
			err := facStruct.factory.Initialize(facInitializeCtx, facStruct.conf)
			if err != nil {
				return errors.WrapError(facInitializeCtx, err)
			}
			log.Trace(facInitializeCtx, "Initialized factory", "factory name", facname)
		}
	}
	return nil
}
