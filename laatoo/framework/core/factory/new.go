package factory

import (
	"laatoo/framework/core/common"
	"laatoo/sdk/core"
	"laatoo/sdk/server"
	//	"laatoo/sdk/errors"
	//	"laatoo/sdk/log"
)

func NewFactoryManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (server.ServerElementHandle, core.ServerElement) {
	fm := &factoryManager{parent: parentElem, serviceFactoryStore: make(map[string]*serviceFactory, 30)}
	fmElemCtx := parentElem.NewCtx("Factory Manager:" + name)
	fmElem := &factoryManagerProxy{Context: fmElemCtx.(*common.Context), manager: fm}
	fm.proxy = fmElem
	return fm, fmElem
}

func ChildFactoryManager(ctx core.ServerContext, name string, parentFacMgr core.ServerElement, parent core.ServerElement, filters ...server.Filter) (server.ServerElementHandle, core.ServerElement) {
	facMgrProxy := parentFacMgr.(*factoryManagerProxy)
	facMgr := facMgrProxy.manager
	store := make(map[string]*serviceFactory, len(facMgr.serviceFactoryStore))
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
	fmElemCtx := parentFacMgr.NewCtx("Factory Manager:" + name)
	fmElem := &factoryManagerProxy{Context: fmElemCtx.(*common.Context), manager: fm}
	fm.proxy = fmElem
	return fm, fmElem
}
