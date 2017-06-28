package factory

import (
	"laatoo/sdk/core"
	"laatoo/sdk/server"
	//	"laatoo/sdk/errors"
	//	"laatoo/sdk/log"
)

func NewFactoryManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (server.ServerElementHandle, core.ServerElement) {
	fm := &factoryManager{name: name, parent: parentElem, serviceFactoryStore: make(map[string]*serviceFactoryProxy, 30)}
	fmElem := &factoryManagerProxy{manager: fm}
	fm.proxy = fmElem
	return fm, fmElem
}

func ChildFactoryManager(ctx core.ServerContext, name string, parentFacMgr core.ServerElement, parent core.ServerElement, filters ...server.Filter) (server.ServerElementHandle, core.ServerElement) {
	facMgrProxy := parentFacMgr.(*factoryManagerProxy)
	facMgr := facMgrProxy.manager
	store := make(map[string]*serviceFactoryProxy, len(facMgr.serviceFactoryStore))
	for k, v := range facMgr.serviceFactoryStore {
		allowed := true
		for _, filter := range filters {
			if !filter.Allowed(ctx, k) {
				allowed = false
				break
			}
		}
		if allowed {
			store[k] = v
		}
	}
	fm := &factoryManager{parent: parent, serviceFactoryStore: store}
	fmElem := &factoryManagerProxy{manager: fm}
	fm.proxy = fmElem
	return fm, fmElem
}
