package core

import "laatoo/sdk/core"

type moduleManagerProxy struct {
	modMgr *moduleManager
}

func (proxy *moduleManagerProxy) Reference() core.ServerElement {
	return &moduleManagerProxy{modMgr: proxy.modMgr}
}

func (proxy *moduleManagerProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *moduleManagerProxy) GetName() string {
	return proxy.modMgr.name
}
func (proxy *moduleManagerProxy) GetType() core.ServerElementType {
	return core.ServerElementModuleManager
}