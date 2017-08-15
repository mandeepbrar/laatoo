package core

import (
	"laatoo/sdk/components"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/server/common"
	"laatoo/server/constants"
)

type cacheManager struct {
	name                 string
	registeredCacheNames map[string]string
	registeredCaches     map[string]components.CacheComponent
	proxy                *cacheManagerProxy
}

func (cm *cacheManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	log.Trace(ctx, "Process Caches")

	cacheManagerConf, err, ok := common.ConfigFileAdapter(ctx, conf, constants.CONF_CACHES)
	if err != nil {
		return err
	}

	if ok {
		cacheNames := cacheManagerConf.AllConfigurations()
		for _, cacheName := range cacheNames {
			cacheConf, err, _ := common.ConfigFileAdapter(ctx, cacheManagerConf, cacheName)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			err = cm.processCache(ctx, cacheConf, cacheName)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	}

	log.Trace(ctx, "Process Caches directory")
	baseDir, _ := ctx.GetString(config.BASEDIR)

	return cm.processCachesFromFolder(ctx, baseDir)
}

func (cm *cacheManager) processCachesFromFolder(ctx core.ServerContext, folder string) error {
	objs, err := cm.loadCachesFromDirectory(ctx, folder)
	if err != nil {
		return err
	}

	if err = common.ProcessObjects(ctx, objs, cm.processCache); err != nil {
		return err
	}
	return nil
}

func (cm *cacheManager) loadCachesFromDirectory(ctx core.ServerContext, folder string) (map[string]config.Config, error) {
	return common.ProcessDirectoryFiles(ctx, folder, constants.CONF_CACHES, true)
}

func (cm *cacheManager) processCache(ctx core.ServerContext, cacheConf config.Config, cacheName string) error {
	cacheSvcName, ok := cacheConf.GetString(constants.CONF_CACHE_SVC)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", constants.CONF_CACHE_SVC)
	}
	cm.registeredCacheNames[cacheName] = cacheSvcName
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
			return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Conf", constants.CONF_CACHE_SVC, "Cache Name", cacheName)
		}
		cm.registeredCaches[cacheName] = cacheSvc
	}
	return nil
}
