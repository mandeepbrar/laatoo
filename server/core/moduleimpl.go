package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
)

type moduleInfo struct {
	*configurableObject
}

func newModuleInfo(description string, configurations []core.Configuration) *moduleInfo {
	f := &moduleInfo{newConfigurableObject(description, "Module")}
	f.setConfigurations(configurations)
	return f
}

func buildModuleInfo(ctx core.ServerContext, conf config.Config) *moduleInfo {
	return &moduleInfo{buildConfigurableObject(ctx, conf)}
}

func (modInfo *moduleInfo) clone() *moduleInfo {
	return &moduleInfo{modInfo.configurableObject.clone()}
}

type moduleImpl struct {
	*moduleInfo
	state State
}

func newModuleImpl() *moduleImpl {
	return &moduleImpl{state: Created, moduleInfo: newModuleInfo("", nil)}
}

func (impl *moduleImpl) setModuleInfo(inf *moduleInfo) {
	impl.moduleInfo = inf
}

func (impl *moduleImpl) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

func (impl *moduleImpl) Start(ctx core.ServerContext) error {
	return nil
}

func (impl *moduleImpl) Describe(ctx core.ServerContext) error {
	return nil
}

func (impl *moduleImpl) MetaInfo(ctx core.ServerContext) map[string]interface{} {
	return map[string]interface{}{}
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
