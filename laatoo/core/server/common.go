package server

import (
	"laatoo/core/cache"
	"laatoo/core/rules"
	"laatoo/core/tasks"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"
)

func initializeObjectLoader(ctx core.ServerContext, conf config.Config, objectLoaderHandle server.ServerElementHandle) error {
	ldrconf, err, ok := config.ConfigFileAdapter(conf, config.CONF_OBJECTLDR)
	if err != nil {
		return err
	}
	if !ok {
		ldrconf = make(config.GenericConfig, 0)
	}
	err = objectLoaderHandle.Initialize(ctx, ldrconf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func initializeChannelManager(ctx core.ServerContext, conf config.Config, channelManagerHandle server.ServerElementHandle) error {
	chmgrconf, err, ok := config.ConfigFileAdapter(conf, config.CONF_CHANNEL_MGR)
	if err != nil {
		return err
	}
	if !ok {
		chmgrconf = make(config.GenericConfig, 0)
	}
	err = channelManagerHandle.Initialize(ctx, chmgrconf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func initializeFactoryManager(ctx core.ServerContext, conf config.Config, factoryManagerHandle server.ServerElementHandle) error {
	facConf, err, ok := config.ConfigFileAdapter(conf, config.CONF_FACMGR)
	if err != nil {
		return err
	}
	if !ok {
		facConf = make(config.GenericConfig, 0)
	}
	err = factoryManagerHandle.Initialize(ctx, facConf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func initializeServiceManager(ctx core.ServerContext, conf config.Config, serviceManagerHandle server.ServerElementHandle) error {
	svcConf, err, ok := config.ConfigFileAdapter(conf, config.CONF_SVCMGR)
	if err != nil {
		return err
	}
	if !ok {
		svcConf = make(config.GenericConfig, 0)
	}
	err = serviceManagerHandle.Initialize(ctx, svcConf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func initializeMessagingManager(ctx core.ServerContext, conf config.Config, messagingManagerHandle server.ServerElementHandle) error {
	msgConf, err, ok := config.ConfigFileAdapter(conf, config.CONF_MSGMGR)
	if err != nil {
		return err
	}
	if !ok {
		msgConf = make(config.GenericConfig, 0)
	}
	err = messagingManagerHandle.Initialize(ctx, msgConf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func createRulesManager(ctx core.ServerContext, name string, conf config.Config, parent core.ServerElement) (server.ServerElementHandle, server.RulesManager, error) {
	rulesConf, err, ok := config.ConfigFileAdapter(conf, config.CONF_RULESMGR)
	if err != nil {
		return nil, nil, err
	}
	if !ok {
		return nil, nil, nil
	}
	rulesCreateCtx := ctx.SubContext("Create Rules Manager")
	rulesMgrHandle, rulesMgr := rules.NewRulesManager(rulesCreateCtx, name, parent)
	err = rulesMgrHandle.Initialize(ctx, rulesConf)
	if err != nil {
		return nil, nil, errors.WrapError(ctx, err)
	}
	return rulesMgrHandle, rulesMgr, nil
}

func createTaskManager(ctx core.ServerContext, name string, conf config.Config, parent core.ServerElement) (server.ServerElementHandle, server.TaskManager, error) {
	taskMgrConf, err, ok := config.ConfigFileAdapter(conf, config.CONF_TASKMGR)
	if err != nil {
		return nil, nil, err
	}
	if !ok {
		return nil, nil, nil
	}
	taskCreateCtx := ctx.SubContext("Create Task Manager")
	taskMgrHandle, taskMgr := tasks.NewTaskManager(taskCreateCtx, name, parent)
	err = taskMgrHandle.Initialize(ctx, taskMgrConf)
	if err != nil {
		return nil, nil, errors.WrapError(ctx, err)
	}
	return taskMgrHandle, taskMgr, nil
}

func createCacheManager(ctx core.ServerContext, name string, conf config.Config, parentCacheMgr core.ServerElement, parent core.ServerElement) (server.ServerElementHandle, server.CacheManager, error) {
	cacheManagerConf, err, ok := config.ConfigFileAdapter(conf, config.CONF_CACHE_MGR)
	if err != nil {
		return nil, nil, err
	}
	if !ok {
		return nil, nil, nil
	}
	cacheMgrCreateCtx := ctx.SubContext("Create Cache Manager")
	var cacheMgrHandle server.ServerElementHandle
	var cacheMgr server.CacheManager
	if parentCacheMgr == nil {
		cacheMgrHandle, cacheMgr = cache.NewCacheManager(cacheMgrCreateCtx, "Root", parent)
	} else {
		cacheMgrHandle, cacheMgr = cache.ChildCacheManager(cacheMgrCreateCtx, name, parent, parentCacheMgr)
	}
	err = cacheMgrHandle.Initialize(ctx, cacheManagerConf)
	if err != nil {
		return nil, nil, errors.WrapError(ctx, err)
	}
	return cacheMgrHandle, cacheMgr, nil
}
