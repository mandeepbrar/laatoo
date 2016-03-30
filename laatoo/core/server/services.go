package server

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	CONF_ENV_SERVICES    = "services"
	CONF_ENV_SERVICENAME = "servicename"
)

//create services within an environment
func (env *Environment) createServices(ctx *serverContext) error {
	//	serviceMiddleware := make(map[string]*Middleware, 100)
	services := make(map[string]core.Service, 100)
	for factoryAlias, factoryConf := range env.ServiceFactoryConfig {
		svcsConf, ok := factoryConf.GetSubConfig(CONF_ENV_SERVICES)
		if ok {
			svcAliases := svcsConf.AllConfigurations()
			for _, svcAlias := range svcAliases {
				_, ok := services[svcAlias]
				if ok {
					continue
				}
				serviceConfig, ok := svcsConf.GetSubConfig(svcAlias)
				if !ok {
					return errors.ThrowError(ctx, CORE_ERROR_SERVICE_CREATION, "Wrong Config for Service Name", svcAlias, "Factory", factoryAlias)
				}
				svcName, ok := serviceConfig.GetString(CONF_ENV_SERVICENAME)
				if !ok {
					return errors.ThrowError(ctx, CORE_ERROR_SERVICE_CREATION, "Service Name not provided for service alias", svcAlias, "Factory", factoryAlias)
				}
				svc, err := env.createService(ctx, svcAlias, svcName, factoryAlias, serviceConfig)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				/*svcMw := createMW(serviceConfig, env.ServiceFactoryMiddleware[factoryAlias])

				if len(*svcMw) > 0 {
					serviceMiddleware[svcAlias] = svcMw
				} else {
				services[svcAlias] = svc
				}*/
				env.ServicesStore[svcAlias] = svc
				log.Logger.Debug(ctx, "Registered service", "service name", svcAlias)
			}
		}
	}

	return nil // env.processMiddleware(ctx, services, serviceMiddleware)
}

//create service
func (env *Environment) createService(ctx *serverContext, serviceAlias string, serviceName string, factoryAlias string, conf config.Config) (core.Service, error) {
	//get the factory from the register
	factory, ok := env.ServiceFactoryStore[factoryAlias]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Factory Alias", factoryAlias)
	}
	svc, err := factory.CreateService(ctx, serviceName, conf)
	if err != nil || svc == nil {
		return nil, errors.RethrowError(ctx, CORE_ERROR_SERVICE_CREATION, err, "Service alias", serviceAlias, "Factory", factoryAlias)
	}
	return svc, nil
}
