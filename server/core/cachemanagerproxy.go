package core

import (
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
)

type cacheManagerProxy struct {
	manager *cacheManager
}

func (cm *cacheManagerProxy) GetCache(ctx core.ServerContext, name string) components.CacheComponent {
	cacheObj, ok := cm.manager.registeredCaches[name]
	if !ok {
		return nil
	}
	return cacheObj
}

func (proxy *cacheManagerProxy) Reference() core.ServerElement {
	return &cacheManagerProxy{manager: proxy.manager}
}
func (proxy *cacheManagerProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *cacheManagerProxy) GetName() string {
	return proxy.manager.name
}
func (proxy *cacheManagerProxy) GetType() core.ServerElementType {
	return core.ServerElementCacheManager
}
func (proxy *cacheManagerProxy) GetContext() core.ServerContext {
	return proxy.manager.svrContext
}
