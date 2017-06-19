package service

import (
	"laatoo/server/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
)

type service struct {
	*common.Context
	name       string
	service    core.Service
	factory    server.Factory
	conf       config.Config
	owner      *serviceManager
	funcs      []core.ServiceFunc
	paramsConf config.Config
}

func (svc *service) start(ctx core.ServerContext) error {
	var mergedConf config.Config
	svcParamsConf, ok := svc.conf.GetSubConfig(config.CONF_SVCPARAMS)
	if ok {
		mergedConf = config.Merge(mergedConf, svcParamsConf)
	}
	funcs := make([]core.ServiceFunc, 0)
	middlewareNames, ok := svc.GetStringArray(config.CONF_MIDDLEWARE)
	if ok {
		for _, mwName := range middlewareNames {
			//			mwSvc, err := ctx.GetService(mwName)
			mwSvcStruct, err := svc.owner.proxy.GetService(ctx, mwName)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			mwSvc := mwSvcStruct.(*service)
			funcs = append(funcs, mwSvc.Invoke)
			mwsvcParamsConf, ok := mwSvc.conf.GetSubConfig(config.CONF_SVCPARAMS)
			if ok {
				mergedConf = config.Merge(mergedConf, mwsvcParamsConf)
			}
		}
	}
	funcs = append(funcs, svc.service.Invoke)
	svc.funcs = funcs
	svc.paramsConf = mergedConf
	log.Logger.Trace(ctx, "params config ", "service name", svc.GetName(), "params conf", svc.paramsConf)
	return svc.service.Start(ctx)
}

func (svc *service) Service() core.Service {
	return svc.service
}

func (svc *service) ParamsConfig() config.Config {
	return svc.paramsConf
}

func (svc *service) Invoke(ctx core.RequestContext) error {
	for _, function := range svc.funcs {
		err := function(ctx)
		if err != nil {
			log.Logger.Trace(ctx, "got error ", "service name", svc.GetName())
			return errors.WrapError(ctx, err)
		}
		if ctx.GetResponse() != nil {
			log.Logger.Trace(ctx, "Got response ", "service name", svc.GetName())
			break
		}
	}
	return nil
}
