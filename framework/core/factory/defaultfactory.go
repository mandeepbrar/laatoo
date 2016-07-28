package factory

import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	//	"laatoo/sdk/log"
	//	"laatoo/core/common"
	"laatoo/sdk/errors"
)

type defaultFactory struct {
}

const (
	CONF_DEFAULTFACTORY_NAME = "__defaultfactory__"
)

func init() {
	objects.RegisterObject(CONF_DEFAULTFACTORY_NAME, createDefaultFactory, nil)
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
