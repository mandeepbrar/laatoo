package cache

import (
	"laatoo/framework/core/common"
	"laatoo/sdk/components"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
)

type cacheManager struct {
	registeredCacheNames map[string]string
	registeredCaches     map[string]components.CacheComponent
	proxy                *cacheManagerProxy
}

func (cm *cacheManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	cacheNames := conf.AllConfigurations()
	for _, cacheName := range cacheNames {
		cacheConf, err, _ := common.ConfigFileAdapter(ctx, conf, cacheName)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		cacheSvcName, ok := cacheConf.GetString(config.CONF_CACHE_SVC)
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", config.CONF_CACHE_SVC)
		}
		cm.registeredCacheNames[cacheName] = cacheSvcName
	}
	return nil
}
func (cm *cacheManager) Start(ctx core.ServerContext) error {
	for cacheName, svcName := range cm.registeredCacheNames {
		svc, err := ctx.GetService(svcName)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		cacheSvc, ok := svc.(components.CacheComponent)
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Conf", config.CONF_CACHE_SVC, "Cache Name", cacheName)
		}
		cm.registeredCaches[cacheName] = cacheSvc
	}
	return nil
}
