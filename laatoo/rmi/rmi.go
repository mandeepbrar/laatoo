package rmi

import (
	"laatoo/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	//	"laatoo/sdk/log"
	"laatoo/core/common"
	"laatoo/sdk/errors"
)

type MethodInvokerFactory struct {
}

const (
	CONF_METHOD_INVOKER_NAME = "method_invoker"
)

func init() {
	objects.RegisterObject(CONF_METHOD_INVOKER_NAME, createMethodInvokerFactory, nil)
}

func createMethodInvokerFactory(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	return &MethodInvokerFactory{}, nil
}

//Create the services configured for factory.
func (mi *MethodInvokerFactory) CreateService(ctx core.ServerContext, name string, method string) (core.Service, error) {
	svcFunc, err := ctx.GetMethod(method)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	return common.NewService(ctx, name, svcFunc), nil
}

func (ds *MethodInvokerFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//The services start serving when this method is called
func (ds *MethodInvokerFactory) Start(ctx core.ServerContext) error {
	return nil
}
