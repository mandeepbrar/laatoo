package main

import (
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

func Manifest() []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_GAEDATA_SERVICES, Object: gaeDataServicesFactory{}}}
}

func (gf *gaeDataServicesFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//Create the services configured for factory.
func (gf *gaeDataServicesFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	return newGaeDataService(ctx, name)
}

//The services start serving when this method is called
func (gf *gaeDataServicesFactory) Start(ctx core.ServerContext) error {
	return nil
}
