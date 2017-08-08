package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

func newFactoryImpl() *factoryImpl {
	return &factoryImpl{state: Created, configurableObject: newConfigurableObject()}
}

type factoryImpl struct {
	*configurableObject
	state State
}

func (impl *factoryImpl) Initialize(ctx core.ServerContext) error {
	return nil
}
func (impl *factoryImpl) Start(ctx core.ServerContext) error {
	return nil
}

//Create the services configured for factory.
func (impl *factoryImpl) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	return nil, nil
}
