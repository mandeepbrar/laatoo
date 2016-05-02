package service

import (
	"laatoo/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"
)

type service struct {
	*common.Context
	name    string
	service core.Service
	factory server.Factory
	conf    config.Config
	owner   *serviceManager
	funcs   []core.ServiceFunc
}

func (svc *service) start(ctx core.ServerContext) error {
	funcs := make([]core.ServiceFunc, 0)
	middlewareNames, ok := svc.GetStringArray(config.CONF_MIDDLEWARE)
	if ok {
		for _, mwName := range middlewareNames {
			mwSvc, err := ctx.GetService(mwName)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			funcs = append(funcs, mwSvc.Invoke)
		}
	}
	funcs = append(funcs, svc.service.Invoke)
	svc.funcs = funcs
	return svc.service.Start(ctx)
}

func (svc *service) Service() core.Service {
	return svc.service
}

func (svc *service) Invoke(ctx core.RequestContext) error {
	for _, function := range svc.funcs {
		err := function(ctx)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		if ctx.GetResponse() != nil {
			break
		}
	}
	return nil
}
