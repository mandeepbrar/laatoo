package plugins

import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	//"laatoosdk/log"
)

const (
	DATAPLUGINS_FACTORY  = "dataplugins"
	DATASVC_CACHE_PLUGIN = "cache"
	DATASVC_JOIN_PLUGIN  = "join"
)

type dataPluginsFactory struct {
}

func init() {
	objects.Register(DATAPLUGINS_FACTORY, dataPluginsFactory{})
}

func (df *dataPluginsFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//Create the services configured for factory.
func (df *dataPluginsFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	switch method {
	case DATASVC_CACHE_PLUGIN:
		{
			return NewDataCacheService(ctx), nil
		}
	case DATASVC_JOIN_PLUGIN:
		{
			return NewJoinCacheService(ctx), nil
		}
	}
	return nil, nil
}

//The services start serving when this method is called
func (df *dataPluginsFactory) Start(ctx core.ServerContext) error {
	return nil
}
