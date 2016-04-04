package server

import (
	"laatoo/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	CONF_APP_SERVICEGROUPS = "servicegroups"
	CONF_APP_SERVICES      = "services"
	CONF_APP_SERVICENAME   = "servicename"
)

//create services within an application
func (app *Application) createServices(ctx *serverContext) error {
	//	serviceMiddleware := make(map[string]*Middleware, 100)
	for factoryAlias, factoryConf := range app.ServiceFactoryConfig {
		grpCtx := ctx.subCtx("Factory:"+factoryAlias, factoryConf, app)
		err := app.processServiceGrp(grpCtx, factoryAlias, factoryAlias, factoryConf)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil // app.processMiddleware(ctx, services, serviceMiddleware)
}

func (app *Application) processServiceGrp(ctx *serverContext, svcgroup string, factoryAlias string, conf config.Config) error {
	//get a map of all the services
	allgroups, ok := conf.GetSubConfig(CONF_APP_SERVICEGROUPS)
	if ok {
		groups := allgroups.AllConfigurations()
		for _, groupname := range groups {
			log.Logger.Debug(ctx, "Process Service group", "groupname", groupname)
			svcgrpConfig, err := common.ConfigFileAdapter(allgroups, groupname)
			if err != nil {
				return errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_CONF, err, "Wrong config for service group", groupname)
			}
			grpCtx := ctx.subCtx("ServiceGroup:"+groupname, svcgrpConfig, app)
			err = app.processServiceGrp(grpCtx, groupname, factoryAlias, svcgrpConfig)
			if err != nil {
				return err
			}
		}
	}

	svcsConf, ok := conf.GetSubConfig(CONF_APP_SERVICES)
	if ok {
		svcAliases := svcsConf.AllConfigurations()
		for _, svcAlias := range svcAliases {
			_, ok := app.ServicesStore[svcAlias]
			if ok {
				continue
			}
			serviceConfig, err := common.ConfigFileAdapter(svcsConf, svcAlias)
			if err != nil {
				return errors.RethrowError(ctx, CORE_ERROR_SERVICE_CREATION, err, "Wrong Config for Service Name", svcAlias)
			}
			svcCtx := ctx.subCtx("Service:"+svcAlias, serviceConfig, app)

			svcName, ok := serviceConfig.GetString(CONF_APP_SERVICENAME)
			if !ok {
				return errors.ThrowError(ctx, CORE_ERROR_SERVICE_CREATION, "Service Name not provided for service alias", svcAlias, "Factory", factoryAlias)
			}
			svc, err := app.createService(svcCtx, svcAlias, svcName, factoryAlias, serviceConfig)
			if err != nil {
				return errors.WrapError(svcCtx, err)
			}
			/*svcMw := createMW(serviceConfig, app.ServiceFactoryMiddleware[factoryAlias])

			if len(*svcMw) > 0 {
				serviceMiddleware[svcAlias] = svcMw
			} else {
			services[svcAlias] = svc
			}*/
			app.ServicesStore[svcAlias] = svc
			log.Logger.Info(ctx, "Registered service", "service name", svcAlias)
		}
	}
	return nil
}

//create service
func (app *Application) createService(ctx *serverContext, serviceAlias string, serviceName string, factoryAlias string, conf config.Config) (core.Service, error) {
	//get the factory from the register
	factory, ok := app.ServiceFactoryStore[factoryAlias]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Factory Alias", factoryAlias)
	}
	svc, err := factory.CreateService(ctx, serviceName, conf)
	if err != nil || svc == nil {
		return nil, errors.RethrowError(ctx, CORE_ERROR_SERVICE_CREATION, err, "Service alias", serviceAlias, "Factory", factoryAlias)
	}
	return svc, nil
}

//initialize services within an application
func (app *Application) initializeServices(ctx *serverContext) error {
	for svcname, service := range app.ServicesStore {
		log.Logger.Debug(ctx, "Initializing service", "service name", svcname)
		err := service.Initialize(ctx)
		if err != nil {
			errors.WrapError(ctx, err)
		}
	}
	return nil
}
