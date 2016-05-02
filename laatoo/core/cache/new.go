package cache

import (
	"laatoo/core/common"
	"laatoo/sdk/core"
	"laatoo/sdk/server"
	"laatoo/sdk/services"
)

func NewCacheManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (*cacheManager, *cacheManagerProxy) {
	cacheMgr := &cacheManager{registeredCacheNames: make(map[string]string, 10), registeredCaches: make(map[string]services.Cache, 10)}
	cacheElemCtx := parentElem.NewCtx("Cache Manager:" + name)
	cacheElem := &cacheManagerProxy{Context: cacheElemCtx.(*common.Context), manager: cacheMgr}
	cacheMgr.proxy = cacheElem
	return cacheMgr, cacheElem
}

func ChildCacheManager(ctx core.ServerContext, name string, parentCacheManager core.ServerElement, parentElem core.ServerElement, filters ...server.Filter) (*cacheManager, *cacheManagerProxy) {
	cacheMgrProxy := parentCacheManager.(*cacheManagerProxy)
	cacheMgr := cacheMgrProxy.manager
	registeredCaches := make(map[string]services.Cache, len(cacheMgr.registeredCaches))
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
	childcacheMgr := &cacheManager{registeredCaches: registeredCaches, registeredCacheNames: registeredCacheNames}
	childcacheElemCtx := parentElem.NewCtx("Cache Manager:" + name)
	childcacheMgrElem := &cacheManagerProxy{Context: childcacheElemCtx.(*common.Context), manager: childcacheMgr}
	childcacheMgr.proxy = childcacheMgrElem
	return childcacheMgr, childcacheMgrElem
}
