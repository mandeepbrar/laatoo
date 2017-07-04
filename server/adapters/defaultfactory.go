package adapters

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	//	"laatoo/sdk/log"
	"laatoo/sdk/errors"
)

type DefaultFactory struct {
}

//Create the services configured for factory.
func (mi *DefaultFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	obj, err := ctx.CreateObject(method)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	svc, ok := obj.(core.Service)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Conf", method)
	}
	return svc, nil
}

func (ds *DefaultFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//The services start serving when this method is called
func (ds *DefaultFactory) Start(ctx core.ServerContext) error {
	return nil
}
