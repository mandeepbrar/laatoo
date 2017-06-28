package service

import (
	"laatoo/sdk/core"
	"laatoo/sdk/server"
	//	"laatoo/sdk/errors"
	//	"laatoo/sdk/log"
)

func NewServiceManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (server.ServerElementHandle, core.ServerElement) {
	sm := &serviceManager{name: name, parent: parentElem, servicesStore: make(map[string]*serviceProxy, 100)}
	smElem := &serviceManagerProxy{manager: sm}
	sm.proxy = smElem
	return sm, smElem
}

func ChildServiceManager(ctx core.ServerContext, name string, parentSvcMgr core.ServerElement, parent core.ServerElement, filters ...server.Filter) (server.ServerElementHandle, core.ServerElement) {
	svcMgrProxy := parentSvcMgr.(*serviceManagerProxy)
	svcMgr := svcMgrProxy.manager
	store := make(map[string]*serviceProxy, len(svcMgr.servicesStore))
	for k, v := range svcMgr.servicesStore {
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
	sm := &serviceManager{name: name, parent: parent, servicesStore: store}
	smElem := &serviceManagerProxy{manager: sm}
	sm.proxy = smElem
	return sm, smElem
}
