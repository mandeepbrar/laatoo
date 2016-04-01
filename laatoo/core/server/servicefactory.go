package server

import (
	"fmt"
	"laatoo/core/registry"
	"laatoo/sdk/config"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	CONF_ENV_SERVICEFACTORIES       = "servicefactories"
	CONF_ENV_SERVICEFACTORYGROUPS   = "servicefactorygroups"
	CONF_ENV_SERVICEFACTORYPROVIDER = "provider"
	CONF_ENV_SERVICEFACTORYCONFIG   = "config"
)

//creates a new service factory
func (env *Environment) createServiceFactories(ctx *serverContext) error {
	return env.processServiceFactoryGrp(ctx, env.Config)
}
func (env *Environment) processServiceFactoryGrp(ctx *serverContext, conf config.Config) error {
	//get a map of all the services
	allgroups, ok := conf.GetSubConfig(CONF_ENV_SERVICEFACTORYGROUPS)
	if ok {
		groups := allgroups.AllConfigurations()
		for _, groupname := range groups {
			groupconfigname, _ := allgroups.GetString(groupname)
			grpFileName := fmt.Sprintf("%s/%s", env.Name, groupconfigname)
			log.Logger.Info(ctx, "Process Service Factory group", "groupname", groupname, "filename", grpFileName)
			facgrpConfig, err := config.NewConfigFromFile(grpFileName)
			if err != nil {
				return errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_CONF, err, "Wrong config for Factory group", groupname, "Missing Config", groupconfigname, "allgroups", allgroups, "groups", groups)
			}
			err = env.processServiceFactoryGrp(ctx, facgrpConfig)
			if err != nil {
				return err
			}
		}
	}
	//get a map of all the services
	svcs, ok := conf.GetSubConfig(CONF_ENV_SERVICEFACTORIES)
	if !ok {
		return nil
	}
	factories := svcs.AllConfigurations()
	for _, factoryName := range factories {
		factoryConfig, ok := svcs.GetSubConfig(factoryName)
		if !ok {
			return errors.ThrowError(ctx, CORE_FACTORY_NOT_CREATED, "Wrong config for Factory Name", factoryName)
		}
		env.processServiceFactoryConfig(ctx, factoryName, factoryConfig)
	}
	return nil
}

func (env *Environment) processServiceFactoryConfig(ctx *serverContext, factoryName string, factoryConfig config.Config) error {
	providerName, ok := factoryConfig.GetString(CONF_ENV_SERVICEFACTORYPROVIDER)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Wrong config for Factory Name", factoryName, "Missing Config", CONF_ENV_SERVICEFACTORYPROVIDER)
	}
	configuration, ok := factoryConfig.Get(CONF_ENV_SERVICEFACTORYCONFIG)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Wrong config for Factory Name", factoryName, "Missing Config", CONF_ENV_SERVICEFACTORYCONFIG)
	}
	//supports both file names as well as subconfig
	factoryConfigFile, ok := configuration.(string)
	var svcfacConfig config.Config
	var err error
	if ok {
		svcFileName := fmt.Sprintf("%s/%s", env.Name, factoryConfigFile)
		log.Logger.Info(ctx, "Creating Service Factory", "factoryName", factoryName, "filename", svcFileName)
		svcfacConfig, err = config.NewConfigFromFile(svcFileName)
		if err != nil {
			return errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_CONF, err, "Could not read from file to create factory", factoryName)
		}
	} else {
		svcfacConfig, ok = config.Cast(configuration)
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Wrong config for Factory Name", factoryName)
		}
	}

	//facmw := createMW(svcfacConfig, env.middleware)

	//env.ServiceFactoryMiddleware[factoryName] = facmw

	//create the service factory
	err = env.createServiceFactory(ctx, factoryName, providerName, svcfacConfig)
	if err != nil {
		return errors.RethrowError(ctx, CORE_FACTORY_NOT_CREATED, err, "Could not create factory", factoryName)
	}
	return nil
}

//create service factory
func (env *Environment) createServiceFactory(ctx *serverContext, alias string, factoryProviderName string, conf config.Config) error {
	_, ok := env.ServiceFactoryStore[alias]
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
		//add the service to the environment
		env.ServiceFactoryStore[alias] = factory
	}
	env.ServiceFactoryConfig[alias] = conf
	return nil
}
