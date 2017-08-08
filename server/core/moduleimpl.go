package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

func newModuleImpl() *moduleImpl {
	return &moduleImpl{state: Created, configurableObject: newConfigurableObject()}
}

type moduleImpl struct {
	*configurableObject
	state State
}

func (impl *moduleImpl) Initialize(ctx core.ServerContext) error {
	return nil
}
func (impl *moduleImpl) Start(ctx core.ServerContext) error {
	return nil
}

func (impl *moduleImpl) Factories(ctx core.ServerContext) map[string]config.Config {
	return nil
}

func (impl *moduleImpl) Services(ctx core.ServerContext) map[string]config.Config {
	return nil
}

func (impl *moduleImpl) Rules(ctx core.ServerContext) map[string]config.Config {
	return nil
}

func (impl *moduleImpl) Channels(ctx core.ServerContext) map[string]config.Config {
	return nil
}

func (impl *moduleImpl) Tasks(ctx core.ServerContext) map[string]config.Config {
	return nil
}
