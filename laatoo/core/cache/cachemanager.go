package cache

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/services"
)

type cacheManager struct {
	registeredCacheNames map[string]string
	registeredCaches     map[string]services.Cache
	proxy                *cacheManagerProxy
}

func (cm *cacheManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	cacheNames := conf.AllConfigurations()
	for _, cacheName := range cacheNames {
		cacheConf, err, _ := config.ConfigFileAdapter(conf, cacheName)
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
		cacheSvc, ok := svc.(services.Cache)
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Conf", config.CONF_CACHE_SVC, "Cache Name", cacheName)
		}
		cm.registeredCaches[cacheName] = cacheSvc
	}
	return nil
}
