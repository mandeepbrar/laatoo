package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/sdk/utils"
	"laatoo/server/cache"
	"laatoo/server/common"
	"laatoo/server/constants"
	"laatoo/server/rules"
	"laatoo/server/tasks"
	"path"
)

func initializeChannelManager(ctx core.ServerContext, conf config.Config, channelManagerHandle server.ServerElementHandle) error {
	chmgrconf, err, ok := common.ConfigFileAdapter(ctx, conf, constants.CONF_CHANNELS)
	if err != nil {
		return err
	}
	if !ok {
		chmgrconf = make(common.GenericConfig, 0)
	}
	err = channelManagerHandle.Initialize(ctx, chmgrconf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func initializeFactoryManager(ctx core.ServerContext, conf config.Config, factoryManagerHandle server.ServerElementHandle) error {
	facConf, err, ok := common.ConfigFileAdapter(ctx, conf, constants.CONF_FACTORIES)
	if err != nil {
		return err
	}
	if !ok {
		facConf = make(common.GenericConfig, 0)
	}
	err = factoryManagerHandle.Initialize(ctx, facConf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func initializeServiceManager(ctx core.ServerContext, conf config.Config, serviceManagerHandle server.ServerElementHandle) error {
	svcConf, err, ok := common.ConfigFileAdapter(ctx, conf, constants.CONF_SERVICES)
	if err != nil {
		return err
	}
	if !ok {
		svcConf = make(common.GenericConfig, 0)
	}
	err = serviceManagerHandle.Initialize(ctx, svcConf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func createMessagingManager(ctx core.ServerContext, name string, conf config.Config, proxy core.ServerElement, parent *abstractserver) (server.ServerElementHandle, server.MessagingManager, error) {
	msgConf, err, found := common.ConfigFileAdapter(ctx, conf, constants.CONF_MESSAGING)
	if err != nil {
		return nil, nil, errors.WrapError(ctx, err)
	}
	var messagingManager server.MessagingManager
	var messagingManagerHandle server.ServerElementHandle
	if !found {
		baseDir, _ := ctx.GetString(constants.CONF_BASE_DIR)
		confFile := path.Join(baseDir, constants.CONF_MESSAGING, constants.CONF_CONFIG_FILE)
		found, _, _ = utils.FileExists(confFile)
		if found {
			var err error
			if msgConf, err = common.NewConfigFromFile(confFile); err != nil {
				return nil, nil, errors.WrapError(ctx, err)
			}
		} else {
			msgConf = make(common.GenericConfig, 0)
		}
	}
	if found {
		msgSvcName, ok := msgConf.GetString(constants.CONF_MESSAGING_SVC)
		if ok {
			msgCtx := ctx.SubContext("Create Messaging Manager")
			msgHandle, msgElem := newMessagingManager(msgCtx, name, proxy, msgSvcName)
			messagingManager = msgElem
			messagingManagerHandle = msgHandle
			log.Trace(msgCtx, "Created messaging manager")
		}
	}

	if (messagingManager == nil) && (parent != nil) && (parent.messagingManager != nil) {
		childMsgMgrHandle, childMsgMgr := childMessagingManager(ctx, name, parent.messagingManager, proxy)
		messagingManagerHandle = childMsgMgrHandle
		messagingManager = childMsgMgr.(server.MessagingManager)
	}

	if messagingManagerHandle != nil {
		msginit := ctx.SubContextWithElement("Initialize messaging manager", core.ServerElementMessagingManager)
		err := messagingManagerHandle.Initialize(msginit, msgConf)
		if err != nil {
			return nil, nil, errors.WrapError(ctx, err)
		}
		log.Debug(msginit, "Initialized messaging manager")
	}
	return messagingManagerHandle, messagingManager, nil
}

func createRulesManager(ctx core.ServerContext, name string, conf config.Config, parent core.ServerElement) (server.ServerElementHandle, server.RulesManager, error) {
	rulesCreateCtx := ctx.SubContext("Create Rules Manager")
	rulesMgrHandle, rulesMgr := rules.NewRulesManager(rulesCreateCtx, name, parent)
	err := rulesMgrHandle.Initialize(ctx, conf)
	if err != nil {
		return nil, nil, errors.WrapError(ctx, err)
	}
	return rulesMgrHandle, rulesMgr, nil
}

func createTaskManager(ctx core.ServerContext, name string, conf config.Config, parent core.ServerElement) (server.ServerElementHandle, server.TaskManager, error) {
	taskCreateCtx := ctx.SubContext("Create Task Manager")
	taskMgrHandle, taskMgr := tasks.NewTaskManager(taskCreateCtx, name, parent)
	err := taskMgrHandle.Initialize(ctx, conf)
	if err != nil {
		return nil, nil, errors.WrapError(ctx, err)
	}
	return taskMgrHandle, taskMgr, nil
}

func createCacheManager(ctx core.ServerContext, name string, conf config.Config, parentCacheMgr core.ServerElement, parent core.ServerElement) (server.ServerElementHandle, server.CacheManager, error) {
	cacheManagerConf, err, ok := common.ConfigFileAdapter(ctx, conf, constants.CONF_CACHE)
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
