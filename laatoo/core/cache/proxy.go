package cache

import (
	"laatoo/core/common"
	"laatoo/sdk/core"
	"laatoo/sdk/services"
)

type cacheManagerProxy struct {
	*common.Context
	manager *cacheManager
}

func (cm *cacheManagerProxy) GetCache(ctx core.ServerContext, name string) services.Cache {
	cacheObj, ok := cm.manager.registeredCaches[name]
	if !ok {
		return nil
	}
	return cacheObj
}
