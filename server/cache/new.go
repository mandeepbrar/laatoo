package cache

import (
	"laatoo/sdk/components"
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

func NewCacheManager(ctx core.ServerContext, name string) (*cacheManager, *cacheManagerProxy) {
	cacheMgr := &cacheManager{name: name, registeredCacheNames: make(map[string]string, 10), registeredCaches: make(map[string]components.CacheComponent, 10)}
	cacheElem := &cacheManagerProxy{manager: cacheMgr}
	cacheMgr.proxy = cacheElem
	return cacheMgr, cacheElem
}

func ChildCacheManager(ctx core.ServerContext, name string, parentCacheManager core.ServerElement, filters ...server.Filter) (*cacheManager, *cacheManagerProxy) {
	cacheMgrProxy := parentCacheManager.(*cacheManagerProxy)
	cacheMgr := cacheMgrProxy.manager
	registeredCaches := make(map[string]components.CacheComponent, len(cacheMgr.registeredCaches))
	registeredCacheNames := make(map[string]string, len(cacheMgr.registeredCacheNames))
	for k, v := range cacheMgr.registeredCaches {
		allowed := true
		for _, filter := range filters {
			if !filter.Allowed(ctx, k) {
				allowed = false
				break
			}
		}
		if allowed {
			registeredCaches[k] = v
		}
	}
	for k, v := range cacheMgr.registeredCacheNames {
		allowed := true
		for _, filter := range filters {
			if !filter.Allowed(ctx, k) {
				allowed = false
				break
			}
		}
		if allowed {
			registeredCacheNames[k] = v
		}
	}
	childcacheMgr := &cacheManager{name: name, registeredCaches: registeredCaches, registeredCacheNames: registeredCacheNames}
	childcacheMgrElem := &cacheManagerProxy{manager: childcacheMgr}
	childcacheMgr.proxy = childcacheMgrElem
	return childcacheMgr, childcacheMgrElem
}
