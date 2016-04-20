package server

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
)

type channelManager struct {
	parent core.ServerElement
	proxy  server.ChannelManager
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
			channelConf, err := config.ConfigFileAdapter(channelsConf, channelName)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			parentChannelName, ok := channelConf.GetString(config.CONF_ENGINE_PARENTCHANNEL)
			if !ok {
				return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", config.CONF_ENGINE_PARENTCHANNEL)
			}
			parentChannel, ok := chanMgr.channelStore[parentChannelName]
			if !ok {
				return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", config.CONF_ENGINE_PARENTCHANNEL)
			}
			channel, err := parentChannel.Child(ctx, channelName, channelConf)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			chanMgr.channelStore[channelName] = channel
			_, childChannels := channelConf.Get(config.CONF_ENGINE_CHANNELS)
			if childChannels {
				err := chanMgr.createChannels(ctx, channelConf)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
			}
			//log.Logger.Trace(ctx, "Creating channel", "Name:", channelName)
			log.Logger.Info(ctx, "Created channel", "Name:", channelName)
		}
	}
	return nil
}

//creates a context specific to factory manager
func (chanMgr *channelManager) createContext(ctx core.ServerContext, name string) core.ServerContext {
	return ctx.NewContextWithElements(name,
		core.ContextMap{core.ServerElementChannelManager: chanMgr.proxy}, core.ServerElementChannelManager)
}
