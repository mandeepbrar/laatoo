package cache

import (
	"laatoo/framework/core/common"
	"laatoo/sdk/components"
	"laatoo/sdk/core"
)

type cacheManagerProxy struct {
	*common.Context
	manager *cacheManager
}

func (cm *cacheManagerProxy) GetCache(ctx core.ServerContext, name string) components.CacheComponent {
	cacheObj, ok := cm.manager.registeredCaches[name]
	if !ok {
		return nil
	}
	return cacheObj
}
