package data

import (
	"laatoo/framework/core/objects"
	"laatoo/framework/services/data/common"
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
	objects.Register(CONF_GAEDATA_SERVICES, gaeDataServicesFactory{})
}

func (gf *gaeDataServicesFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//Create the services configured for factory.
func (gf *gaeDataServicesFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	switch method {
	case common.CONF_DATA_SVCS:
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
