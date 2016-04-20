package server

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"
)

func initializeObjectLoader(ctx core.ServerContext, conf config.Config, objectLoaderHandle server.ServerElementHandle) error {
	ldrconf, ok := conf.GetSubConfig(config.CONF_OBJECTLDR)
	if !ok {
		ldrconf = make(config.GenericConfig, 0)
	}
	err := objectLoaderHandle.Initialize(ctx, ldrconf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func initializeChannelManager(ctx core.ServerContext, conf config.Config, channelManagerHandle server.ServerElementHandle) error {
	chmgrconf, ok := conf.GetSubConfig(config.CONF_CHANNEL_MGR)
	if !ok {
		chmgrconf = make(config.GenericConfig, 0)
	}
	err := channelManagerHandle.Initialize(ctx, chmgrconf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func initializeFactoryManager(ctx core.ServerContext, conf config.Config, factoryManagerHandle server.ServerElementHandle) error {
	facConf, ok := conf.GetSubConfig(config.CONF_FACMGR)
	if !ok {
		facConf = make(config.GenericConfig, 0)
	}
	err := factoryManagerHandle.Initialize(ctx, facConf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func initializeServiceManager(ctx core.ServerContext, conf config.Config, serviceManagerHandle server.ServerElementHandle) error {
	svcConf, ok := conf.GetSubConfig(config.CONF_SVCMGR)
	if !ok {
		svcConf = make(config.GenericConfig, 0)
	}
	err := serviceManagerHandle.Initialize(ctx, svcConf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}
