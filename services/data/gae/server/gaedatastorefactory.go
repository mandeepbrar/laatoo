package main

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	//"laatoo/sdk/errors"
	//"laatoosdk/log"
)

type gaeDataServicesFactory struct {
	core.ServiceFactory
}

const (
	GAEDATA_SERVICES = "gaedatastore_services"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: GAEDATA_SERVICES, Object: gaeDataServicesFactory{}}}
}

//Create the services configured for factory.
func (gf *gaeDataServicesFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	return newGaeDataService(ctx, name)
}
