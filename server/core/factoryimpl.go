package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

type factoryInfo struct {
	*configurableObject
}

func newFactoryInfo(description string, configurations []core.Configuration) *factoryInfo {
	f := &factoryInfo{newConfigurableObject(description, "Factory")}
	f.setConfigurations(configurations)
	return f
}

func buildFactoryInfo(ctx core.ServerContext, conf config.Config) *factoryInfo {
	return &factoryInfo{buildConfigurableObject(ctx, conf)}
}
func (facInfo *factoryInfo) clone() *factoryInfo {
	return &factoryInfo{facInfo.configurableObject.clone()}
}

type factoryImpl struct {
	*factoryInfo
	state State
}

func newFactoryImpl() *factoryImpl {
	return &factoryImpl{state: Created, factoryInfo: newFactoryInfo("", nil)}
}

func (impl *factoryImpl) setFactoryInfo(fi *factoryInfo) {
	impl.factoryInfo = fi
}

func (impl *factoryImpl) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}
func (impl *factoryImpl) Start(ctx core.ServerContext) error {
	return nil
}
func (impl *factoryImpl) Describe(ctx core.ServerContext) {
}

//Create the services configured for factory.
func (impl *factoryImpl) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	return nil, nil
}
