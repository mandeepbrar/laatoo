package factory

import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	//	"laatoo/sdk/log"
	"laatoo/sdk/errors"
)

type defaultMethodFactory struct {
}

const (
	CONF_DEFAULTMETHODFACTORY_NAME = "__defaultmethodfactory__"
)

func init() {
	objects.RegisterObject(CONF_DEFAULTMETHODFACTORY_NAME, createDefaultMethodFactory, nil)
}

func createDefaultMethodFactory() interface{} {
	return &defaultMethodFactory{}
}

//Create the services configured for factory.
func (mi *defaultMethodFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	svcFunc, err := ctx.GetMethod(method)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	return core.NewService(ctx, name, svcFunc), nil
}

func (ds *defaultMethodFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//The services start serving when this method is called
func (ds *defaultMethodFactory) Start(ctx core.ServerContext) error {
	return nil
}
