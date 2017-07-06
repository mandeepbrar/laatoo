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

	secondPass map[string]config.Config

	//store for service factory in an application
	channelStore map[string]server.Channel
}

func (chanMgr *channelManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	chanMgr.secondPass = make(map[string]config.Config)

	log.Trace(ctx, "Create Channels")
	err := chanMgr.createChannels(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	if err := common.ProcessDirectoryFiles(ctx, constants.CONF_CHANNELS, chanMgr.createChannel, true); err != nil {
		return errors.WrapError(ctx, err)
	}

	for {
		channelsToCreate := len(chanMgr.secondPass)
		if channelsToCreate == 0 {
			break
		}
		chansToCreate := chanMgr.secondPass
		chanMgr.secondPass = make(map[string]config.Config)
		if err := chanMgr.reviewMissingChannels(ctx, chansToCreate); err != nil {
			return errors.WrapError(ctx, err)
		}
		moreChannelsToCreate := len(chanMgr.secondPass)
		if moreChannelsToCreate == 0 {
			break
		}
		if moreChannelsToCreate == channelsToCreate {
			return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Reason", "parents are missing for channels", "channels", chanMgr.secondPass)
		}
	}

	return nil
}

func (chanMgr *channelManager) Start(ctx core.ServerContext) error {
	return nil
}

func (chanMgr *channelManager) reviewMissingChannels(ctx core.ServerContext, chansToReview map[string]config.Config) error {
	for channelName, channelConf := range chansToReview {
		if err := chanMgr.createChannel(ctx, channelConf, channelName); err != nil {
			return errors.WrapError(ctx, err)
		}
	}
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
	createCtx := ctx.SubContext("Create Channel: " + channelName)
	parentChannelName, ok := channelConf.GetString(constants.CONF_ENGINE_PARENTCHANNEL)
	if !ok {
		return errors.ThrowError(createCtx, errors.CORE_ERROR_MISSING_CONF, "conf", constants.CONF_ENGINE_PARENTCHANNEL)
	}
	parentChannel, ok := chanMgr.channelStore[parentChannelName]
	if !ok {
		chanMgr.secondPass[channelName] = channelConf
		log.Trace(createCtx, "Could not find parent channel", "Channel Name", channelName)
		return nil
		//return errors.ThrowError(createCtx, errors.CORE_ERROR_BAD_CONF, "conf", constants.CONF_ENGINE_PARENTCHANNEL)
	}
	log.Trace(createCtx, "Found parent channel ", "Parent Channel Name", parentChannelName, "Found", ok)
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
