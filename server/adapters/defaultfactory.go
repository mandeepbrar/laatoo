package adapters

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/server/objects"
	//	"laatoo/sdk/log"
	"laatoo/sdk/errors"
	"laatoo/server/common"
)

type defaultFactory struct {
}

func init() {
	objects.RegisterObject(common.CONF_DEFAULTFACTORY_NAME, createDefaultFactory, nil)
}

func createDefaultFactory() interface{} {
	return &defaultFactory{}
}

//Create the services configured for factory.
func (mi *defaultFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
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

func (ds *defaultFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//The services start serving when this method is called
func (ds *defaultFactory) Start(ctx core.ServerContext) error {
	return nil
}
