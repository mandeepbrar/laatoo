package server

import (
	"laatoo/framework/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
)

type channelManager struct {
	proxy server.ChannelManager
	//store for service factory in an application
	channelStore map[string]server.Channel
}

func (chanMgr *channelManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	chanmgrInitializeCtx := chanMgr.createContext(ctx, "Initialize channel manager")
	log.Logger.Trace(chanmgrInitializeCtx, "Create Channels")
	err := chanMgr.createChannels(chanmgrInitializeCtx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (chanMgr *channelManager) Start(ctx core.ServerContext) error {
	return nil
}

func (chanMgr *channelManager) createChannels(ctx core.ServerContext, conf config.Config) error {
	channelsConf, ok := conf.GetSubConfig(config.CONF_ENGINE_CHANNELS)
	if ok {
		channelNames := channelsConf.AllConfigurations()
		for _, channelName := range channelNames {
			createCtx := chanMgr.createContext(ctx, "Create Channel"+channelName)
			channelConf, err, _ := common.ConfigFileAdapter(ctx, channelsConf, channelName)
			if err != nil {
				return errors.WrapError(createCtx, err)
			}
			parentChannelName, ok := channelConf.GetString(config.CONF_ENGINE_PARENTCHANNEL)
			if !ok {
				return errors.ThrowError(createCtx, errors.CORE_ERROR_MISSING_CONF, "conf", config.CONF_ENGINE_PARENTCHANNEL)
			}
			parentChannel, ok := chanMgr.channelStore[parentChannelName]
			if !ok {
				return errors.ThrowError(createCtx, errors.CORE_ERROR_BAD_CONF, "conf", config.CONF_ENGINE_PARENTCHANNEL)
			}
			channel, err := parentChannel.Child(createCtx, channelName, channelConf)
			if err != nil {
				return errors.WrapError(createCtx, err)
			}
			chanMgr.channelStore[channelName] = channel
			_, childChannels := channelConf.Get(config.CONF_ENGINE_CHANNELS)
			if childChannels {
				err := chanMgr.createChannels(createCtx, channelConf)
				if err != nil {
					return errors.WrapError(createCtx, err)
				}
			}
			//log.Logger.Trace(ctx, "Creating channel", "Name:", channelName)
			log.Logger.Info(createCtx, "Created channel", "Name:", channelName)
		}
	}
	return nil
}

//creates a context specific to factory manager
func (chanMgr *channelManager) createContext(ctx core.ServerContext, name string) core.ServerContext {
	return ctx.NewContextWithElements(name,
		core.ContextMap{core.ServerElementChannelManager: chanMgr.proxy}, core.ServerElementChannelManager)
}
