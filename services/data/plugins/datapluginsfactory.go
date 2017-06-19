package main

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	//"laatoosdk/log"
)

const (
	DATAPLUGINS_FACTORY    = "dataplugins"
	DATASVC_CACHE_PLUGIN   = "cache"
	DATASVC_JOIN_PLUGIN    = "join"
	DATASVC_CHECKOWNER     = "checkowner"
	DATASVC_PLUGINS_HOOKER = "hook"
)

type dataPluginsFactory struct {
}

func Manifest() []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: DATAPLUGINS_FACTORY, Object: dataPluginsFactory{}}}
}

func (df *dataPluginsFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//Create the services configured for factory.
func (df *dataPluginsFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	switch method {
	case DATASVC_PLUGINS_HOOKER:
		{
			return NewPluginHookService(ctx, name, conf)
		}
	case DATASVC_CACHE_PLUGIN:
		{
			return NewDataCacheService(ctx), nil
		}
	case DATASVC_JOIN_PLUGIN:
		{
			return NewJoinService(ctx), nil
		}
	case DATASVC_CHECKOWNER:
		{
			return NewCheckOwnerService(ctx), nil
		}
	}
	return nil, nil
}

//The services start serving when this method is called
func (df *dataPluginsFactory) Start(ctx core.ServerContext) error {
	return nil
}
