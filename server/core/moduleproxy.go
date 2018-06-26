package core

import "laatoo/sdk/server/core"

type moduleProxy struct {
	mod *serverModule
}

func (mod *moduleProxy) Reference() core.ServerElement {
	return &moduleProxy{mod: mod.mod}
}

func (mod *moduleProxy) GetProperty(string) interface{} {
	return nil
}
func (mod *moduleProxy) GetName() string {
	return mod.mod.name
}
func (mod *moduleProxy) GetType() core.ServerElementType {
	return core.ServerElementModule
}

func (mod *moduleProxy) GetObject() core.Module {
	return mod.mod.userModule
}
func (mod *moduleProxy) GetModuleProperties() map[string]interface{} {
	return mod.mod.properties
}
