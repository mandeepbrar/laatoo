package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
)

type factoryInfo struct {
	*configurableObject
}

func newFactoryInfo(name, description string, configurations []core.Configuration) *factoryInfo {
	f := &factoryInfo{newConfigurableObject(name, description, "Factory")}
	f.setConfigurations(configurations)
	return f
}

func buildFactoryInfo(ctx core.ServerContext, name string, conf config.Config) *factoryInfo {
	return &factoryInfo{buildConfigurableObject(ctx, name, conf)}
}
func (facInfo *factoryInfo) clone() *factoryInfo {
	return &factoryInfo{facInfo.configurableObject.clone()}
}

type factoryImpl struct {
	*factoryInfo
	state State
}

func newFactoryImpl(name string) *factoryImpl {
	return &factoryImpl{state: Created, factoryInfo: newFactoryInfo(name, "", nil)}
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
func (impl *factoryImpl) Describe(ctx core.ServerContext) error {
	return nil
}

func (impl *factoryImpl) Stop(ctx core.ServerContext) error {
	return nil
}
func (impl *factoryImpl) Unload(ctx core.ServerContext) error {
	return nil
}

//Create the services configured for factory.
func (impl *factoryImpl) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	return nil, nil
}
