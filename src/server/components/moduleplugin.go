package components

import (
	"laatoo/sdk/config"
	"laatoo/sdk/server/core"
)

type ModInfo struct {
	ModName            string
	PluginName         string
	PluginDir          string
	ParentModName      string
	Mod                core.Module
	UserObj            core.Module
	PluginConf         config.Config
	ModConf            config.Config
	ModSettings        config.Config
	Configurations     map[string]core.Configuration
	ModProps           map[string]interface{}
	IsExtended         bool
	ExtendedPluginName string
	ExtendedPluginConf config.Config
	ExtendedPluginDir  string
	Hot                bool
}

func (info *ModInfo) GetContext(ctx core.ServerContext, variable string) (interface{}, bool) {
	return info.Mod.GetContext(ctx, variable)
}

type ModuleManagerPlugin interface {
	GetName() string
	Load(ctx core.ServerContext, modInfo *ModInfo) error
	Loaded(ctx core.ServerContext) error
	Unloaded(ctx core.ServerContext, insName, modName string) error
	Unloading(ctx core.ServerContext, insName, modName string) error
}
