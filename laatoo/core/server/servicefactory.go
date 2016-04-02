package server

import (
	"laatoo/core/common"
	"laatoo/core/registry"
	"laatoo/sdk/config"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	CONF_APP_SERVICEFACTORIES       = "servicefactories"
	CONF_APP_SERVICEFACTORYGROUPS   = "servicefactorygroups"
	CONF_APP_SERVICEFACTORYPROVIDER = "provider"
	CONF_APP_SERVICEFACTORYCONFIG   = "config"
)

//creates a new service factory
func (app *Application) createServiceFactories(ctx *serverContext) error {
	return app.processServiceFactoryGrp(ctx, app.Config)
}
func (app *Application) processServiceFactoryGrp(ctx *serverContext, conf config.Config) error {
	//get a map of all the services
	allgroups, ok := conf.GetSubConfig(CONF_APP_SERVICEFACTORYGROUPS)
	if ok {
		groups := allgroups.AllConfigurations()
		for _, groupname := range groups {
			log.Logger.Trace(ctx, "Process Service Factory group", "groupname", groupname)
			facgrpConfig, err := common.ConfigFileAdapter(allgroups, groupname)
			if err != nil {
				return errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_CONF, err, "Wrong config for Factory group", groupname)
			}
			grpCtx := ctx.subCtx("Group:"+groupname, facgrpConfig, app)
			err = app.processServiceFactoryGrp(grpCtx, facgrpConfig)
			if err != nil {
				return err
			}
		}
	}
	//get a map of all the services
	svcs, ok := conf.GetSubConfig(CONF_APP_SERVICEFACTORIES)
	if !ok {
		return nil
	}
	factories := svcs.AllConfigurations()
	for _, factoryName := range factories {
		log.Logger.Trace(ctx, "Process Factory ", "Factory name", factoryName)
		factoryConfig, err := common.ConfigFileAdapter(svcs, factoryName)
		if err != nil {
			return errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_CONF, err, "Wrong config for factory", factoryName)
		}
		facCtx := ctx.subCtx("Factory:"+factoryName, factoryConfig, app)
		app.processServiceFactoryConfig(facCtx, factoryName, factoryConfig)
	}
	return nil
}

func (app *Application) processServiceFactoryConfig(ctx *serverContext, factoryName string, factoryConfig config.Config) error {
	providerName, ok := factoryConfig.GetString(CONF_APP_SERVICEFACTORYPROVIDER)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Wrong config for Factory Name", factoryName, "Missing Config", CONF_APP_SERVICEFACTORYPROVIDER)
	}

	svcfacConfig, err := common.ConfigFileAdapter(factoryConfig, CONF_APP_SERVICEFACTORYCONFIG)
	if err != nil {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Wrong config for Factory Name", factoryName)
	}

	//facmw := createMW(svcfacConfig, app.middleware)

	//app.ServiceFactoryMiddleware[factoryName] = facmw

	//create the service factory
	err = app.createServiceFactory(ctx, factoryName, providerName, svcfacConfig)
	if err != nil {
		return errors.RethrowError(ctx, CORE_FACTORY_NOT_CREATED, err, "Could not create factory", factoryName)
	}
	return nil
}

//create service factory
func (app *Application) createServiceFactory(ctx *serverContext, alias string, factoryProviderName string, conf config.Config) error {
	_, ok := app.ServiceFactoryStore[alias]
	if ok {
		return nil
	}
	//get the factory provider from the register
	factoryProvider, ok := registry.ServiceFactoryProviders[factoryProviderName]
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Factory Name", factoryProviderName)
	}
	//create factory with given config from factory provider
	factory, err := factoryProvider(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if factory != nil {
		//add the service to the application
		app.ServiceFactoryStore[alias] = factory
	}
	app.ServiceFactoryConfig[alias] = conf
	return nil
}
