package data

import (
	"laatoo/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	//"laatoo/sdk/errors"
	//"laatoosdk/log"
)

type gaeDataServicesFactory struct {
}

const (
	CONF_GAEDATA_SERVICES = "gaedatastore"
)

func init() {
	objects.RegisterObject(CONF_GAEDATA_SERVICES, createGAEDataServicesFactory, nil)
}

func createGAEDataServicesFactory(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	return &gaeDataServicesFactory{}, nil
}

func (gf *gaeDataServicesFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//Create the services configured for factory.
func (gf *gaeDataServicesFactory) CreateService(ctx core.ServerContext, name string, method string) (core.Service, error) {
	switch method {
	case CONF_DATA_SVCS:
		{

			return newGaeDataService(ctx, name)

		}
	}
	return nil, nil
}

//The services start serving when this method is called
func (gf *gaeDataServicesFactory) Start(ctx core.ServerContext) error {
	return nil
}