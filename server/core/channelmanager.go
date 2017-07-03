package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/common"
	"laatoo/server/constants"
)

type channelManager struct {
	name string

	proxy server.ChannelManager

	parent core.ServerElement
	//store for service factory in an application
	channelStore map[string]server.Channel
}

func (chanMgr *channelManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	chanmgrInitializeCtx := ctx.SubContext("Initialize channel manager")
	log.Trace(chanmgrInitializeCtx, "Create Channels")
	err := chanMgr.createChannels(chanmgrInitializeCtx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	if err := common.ProcessDirectoryFiles(chanmgrInitializeCtx, constants.CONF_CHANNELS, chanMgr.createChannel, true); err != nil {
		return errors.WrapError(chanmgrInitializeCtx, err)
	}

	return nil
}

func (chanMgr *channelManager) Start(ctx core.ServerContext) error {
	return nil
}

func (chanMgr *channelManager) createChannels(ctx core.ServerContext, conf config.Config) error {
	channelsConf, ok := conf.GetSubConfig(constants.CONF_ENGINE_CHANNELS)
	if ok {
		channelNames := channelsConf.AllConfigurations()
		for _, channelName := range channelNames {
			channelConf, err, _ := common.ConfigFileAdapter(ctx, channelsConf, channelName)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			if err := chanMgr.createChannel(ctx, channelConf, channelName); err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	}
	return nil
}

func (chanMgr *channelManager) createChannel(ctx core.ServerContext, channelConf config.Config, channelName string) error {
	createCtx := ctx.SubContext("Create Channel" + channelName)
	parentChannelName, ok := channelConf.GetString(constants.CONF_ENGINE_PARENTCHANNEL)
	if !ok {
		return errors.ThrowError(createCtx, errors.CORE_ERROR_MISSING_CONF, "conf", constants.CONF_ENGINE_PARENTCHANNEL)
	}
	parentChannel, ok := chanMgr.channelStore[parentChannelName]
	if !ok {
		return errors.ThrowError(createCtx, errors.CORE_ERROR_BAD_CONF, "conf", constants.CONF_ENGINE_PARENTCHANNEL)
	}
	channel, err := parentChannel.Child(createCtx, channelName, channelConf)
	if err != nil {
		return errors.WrapError(createCtx, err)
	}
	chanMgr.channelStore[channelName] = channel
	_, childChannels := channelConf.Get(constants.CONF_ENGINE_CHANNELS)
	if childChannels {
		err := chanMgr.createChannels(createCtx, channelConf)
		if err != nil {
			return errors.WrapError(createCtx, err)
		}
	}
	//log.Trace(ctx, "Creating channel", "Name:", channelName)
	log.Info(createCtx, "Created channel", "Name:", channelName)
	return nil
}
