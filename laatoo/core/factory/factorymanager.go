package factory

import (
	"laatoo/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
)

const (
	CONF_SERVICEFACTORYGROUPS = "factorygroups"
	CONF_SERVICEFACTORIES     = "factories"
	CONF_SERVICEFACTORY       = "factory"
	CONF_SERVICEFACTORYCONFIG = "config"
)

type factoryManager struct {
	parent core.ServerElement
	proxy  server.FactoryManager
	//store for service factory in an application
	serviceFactoryStore map[string]*serviceFactory
}

func (facMgr *factoryManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	facmgrInitializeCtx := facMgr.createContext(ctx, "Initialize factory manager")
	err := facMgr.createServiceFactories(facmgrInitializeCtx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	err = facMgr.createServiceFactory(facmgrInitializeCtx, CONF_DEFAULTFACTORY_NAME, &config.GenericConfig{CONF_SERVICEFACTORY: CONF_DEFAULTFACTORY_NAME})
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	err = facMgr.createServiceFactory(facmgrInitializeCtx, CONF_DEFAULTMETHODFACTORY_NAME, &config.GenericConfig{CONF_SERVICEFACTORY: CONF_DEFAULTMETHODFACTORY_NAME})
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return facMgr.initializeFactories(facmgrInitializeCtx)
}

func (facMgr *factoryManager) Start(ctx core.ServerContext) error {
	facmgrStartCtx := facMgr.createContext(ctx, "Start factory manager")
	for facname, facStruct := range facMgr.serviceFactoryStore {
		if facStruct.owner == facMgr {
			log.Logger.Debug(ctx, "Starting factory", "factory name", facname)
			facStartCtx := facmgrStartCtx.NewContextWithElements("Start"+facname, core.ContextMap{core.ServerElementServiceFactory: facStruct}, core.ServerElementServiceFactory)
			err := facStruct.factory.Start(facStartCtx)
			if err != nil {
				return errors.WrapError(facStartCtx, err)
			}
		}
	}
	return nil
}

func (facMgr *factoryManager) createServiceFactories(ctx core.ServerContext, conf config.Config) error {
	//get a map of all the services
	allgroups, ok := conf.GetSubConfig(CONF_SERVICEFACTORYGROUPS)
	if ok {
		groups := allgroups.AllConfigurations()
		for _, groupname := range groups {
			log.Logger.Trace(ctx, "Process Service Factory group", "groupname", groupname)
			facgrpConfig, err, _ := config.ConfigFileAdapter(allgroups, groupname)
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
		log.Logger.Trace(ctx, "Process Factory ", "Factory name", factoryName)
		factoryConfig, err, _ := config.ConfigFileAdapter(factoriesConfig, factoryName)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		facCtx := ctx.SubContext("Create Factory:" + factoryName)
		err = facMgr.createServiceFactory(facCtx, factoryName, factoryConfig)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}

func (facMgr *factoryManager) createServiceFactory(ctx core.ServerContext, factoryAlias string, factoryConfig config.Config) error {
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

	factoryInt, err := ctx.CreateObject(factoryName, nil)

	if err != nil {
		return errors.WrapError(ctx, err)
	}
	factory, ok := factoryInt.(core.ServiceFactory)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Object is not factory for Factory Name", factoryName)
	}

	if factory != nil {
		//creates a factory context from the parent
		facElem := facMgr.parent.NewCtx(factoryAlias).(*common.Context)
		fac := &serviceFactory{Context: facElem, name: factoryAlias, factory: factory, owner: facMgr, conf: factoryConfig}

		middleware, ok := factoryConfig.GetStringArray(config.CONF_MIDDLEWARE)
		if ok {
			parentMw, ok := facMgr.parent.GetStringArray(config.CONF_MIDDLEWARE)
			if ok {
				middleware = append(parentMw, middleware...)
			}
			facElem.Set(config.CONF_MIDDLEWARE, middleware)
		}

		cacheToUse, ok := factoryConfig.GetString(config.CONF_CACHE_NAME)
		if ok {
			fac.Set("__cache", cacheToUse)
		}

		//add the service to the application
		facMgr.serviceFactoryStore[factoryAlias] = fac
	}
	return nil
}

//initialize services within an application
func (facMgr *factoryManager) initializeFactories(ctx core.ServerContext) error {
	for facname, facStruct := range facMgr.serviceFactoryStore {
		if facStruct.owner == facMgr {
			log.Logger.Debug(ctx, "Initializing factory", "factory name", facname)
			facInitializeCtx := ctx.NewContextWithElements("Initialize "+facname, core.ContextMap{core.ServerElementServiceFactory: facStruct}, core.ServerElementServiceFactory)
			err := facStruct.factory.Initialize(facInitializeCtx, facStruct.conf)
			if err != nil {
				return errors.WrapError(facInitializeCtx, err)
			}
		}
	}
	return nil
}

//creates a context specific to factory manager
func (facMgr *factoryManager) createContext(ctx core.ServerContext, name string) core.ServerContext {
	return ctx.NewContextWithElements(name,
		core.ContextMap{core.ServerElementFactoryManager: facMgr.proxy}, core.ServerElementFactoryManager)
}
