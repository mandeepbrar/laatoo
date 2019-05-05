package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/common"
	"laatoo/server/constants"
)

type channelManager struct {
	name string

	proxy elements.ChannelManager

	parent core.ServerElement

	secondPass map[string]config.Config

	//store for service factory in an application
	channelStore   map[string]elements.Channel
	channelConfs   map[string]config.Config
	serviceManager elements.ServiceManager
}

func (chanMgr *channelManager) Initialize(ctx core.ServerContext, conf config.Config) error {

	log.Trace(ctx, "Create Channels")
	err := chanMgr.createChannels(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	baseDir, _ := ctx.GetString(config.BASEDIR)

	modManager := ctx.GetServerElement(core.ServerElementModuleManager).(*moduleManagerProxy).modMgr
	svcMgr := ctx.GetServerElement(core.ServerElementServiceManager).(elements.ServiceManager)
	chanMgr.serviceManager = svcMgr

	if err = modManager.loadChannels(ctx, chanMgr.processChannel); err != nil {
		return err
	}

	if err := chanMgr.processChannelsFromFolder(ctx, baseDir); err != nil {
		return errors.WrapError(ctx, err)
	}

	return nil
}

func (chanMgr *channelManager) Start(ctx core.ServerContext) error {

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

	for chanName, channel := range chanMgr.channelStore {
		err := chanMgr.startChannel(ctx, chanName, channel)
		if err != nil {
			return errors.WrapError(ctx, err)
		}

	}
	return nil
}

func (chanMgr *channelManager) startChannel(ctx core.ServerContext, chanName string, channel elements.Channel) error {
	chanCtx := ctx.SubContext(chanName)
	svcName := channel.GetServiceName()
	if svcName != "" {
		svcProxy, err := chanMgr.serviceManager.GetService(ctx, svcName)
		if err != nil {
			return err
		}
		proxy := svcProxy.(*serviceProxy)
		svcServeCtx := proxy.svc.svrContext.newContext("Serve: " + proxy.svc.name)

		chanConf := chanMgr.channelConfs[chanName]

		if err := processLogging(svcServeCtx, chanConf, chanName); err != nil {
			return errors.WrapError(svcServeCtx, err)
		}

		err = channel.Serve(svcServeCtx)
		if err != nil {
			return err
		}
		log.Info(chanCtx, "Serving channel ", "channel", chanName)
	} else {
		log.Info(chanCtx, "No service configured channel ", "channel", chanName)
	}
	return nil
}

/*
func (cm *channelManagerProxy) Serve(ctx core.ServerContext, channelName string, svc elements.Service, channelConfig config.Config) error {
	channel, ok := cm.manager.channelStore[channelName]
	if ok {
		return channel.Serve(ctx, svc, channelConfig)
	} else {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "No such channel", channelName)
	}
}*/

func (chanMgr *channelManager) processChannelsFromFolder(ctx core.ServerContext, folder string) error {
	objs, err := chanMgr.loadChannelsFromFolder(ctx, folder)
	if err != nil {
		return err
	}

	if err = common.ProcessObjects(ctx, objs, chanMgr.processChannel); err != nil {
		return err
	}
	return nil
}

func (chanMgr *channelManager) loadChannelsFromFolder(ctx core.ServerContext, folder string) (map[string]config.Config, error) {
	return common.ProcessDirectoryFiles(ctx, folder, constants.CONF_CHANNELS, true)
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
	channelsConf, ok := conf.GetSubConfig(ctx, constants.CONF_ENGINE_CHANNELS)
	if ok {
		channelNames := channelsConf.AllConfigurations(ctx)
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

func (chanMgr *channelManager) processChannel(ctx core.ServerContext, channelConf config.Config, channelName string) error {
	arr, ok := channelConf.GetConfigArray(ctx, constants.CONF_CHANNELS)
	if ok {
		for _, conf := range arr {
			name, ok := conf.GetString(ctx, constants.CONF_OBJECT_NAME)
			if !ok {
				return errors.MissingConf(ctx, constants.CONF_OBJECT_NAME, "Channels", channelName)
			}
			if err := chanMgr.createChannel(ctx, conf, name); err != nil {
				return err
			}
		}
	} else {
		return chanMgr.createChannel(ctx, channelConf, channelName)
	}
	return nil
}

func (chanMgr *channelManager) createChannel(ctx core.ServerContext, channelConf config.Config, channelName string) error {

	_, ok := chanMgr.channelConfs[channelName]
	if ok {
		return nil
	}

	createCtx := ctx.SubContext("Create Channel: " + channelName)
	log.Trace(createCtx, "Creating channel with conf ", "conf", channelConf)
	parentChannelName, ok := channelConf.GetString(ctx, constants.CONF_ENGINE_PARENTCHANNEL)
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
	_, childChannels := channelConf.Get(createCtx, constants.CONF_ENGINE_CHANNELS)
	if childChannels {
		err := chanMgr.createChannels(createCtx, channelConf)
		if err != nil {
			return errors.WrapError(createCtx, err)
		}
	}
	chanMgr.channelConfs[channelName] = channelConf
	//log.Trace(ctx, "Creating channel", "Name:", channelName)
	log.Info(createCtx, "Created channel", "Name:", channelName)
	return nil
}

func (chanMgr *channelManager) unloadModuleChannels(ctx core.ServerContext, mod *serverModule) error {
	ctx = ctx.SubContext("unload channels")
	if err := common.ProcessObjects(ctx, mod.channels, chanMgr.unloadChannel); err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (chanMgr *channelManager) unloadChildChannels(ctx core.ServerContext, conf config.Config) error {
	channelsConf, ok := conf.GetSubConfig(ctx, constants.CONF_ENGINE_CHANNELS)
	if ok {
		channelNames := channelsConf.AllConfigurations(ctx)
		for _, channelName := range channelNames {
			channelConf, err, _ := common.ConfigFileAdapter(ctx, channelsConf, channelName)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			if err := chanMgr.unloadChannel(ctx, channelConf, channelName); err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	}
	return nil
}

func (chanMgr *channelManager) unloadChannel(ctx core.ServerContext, channelConf config.Config, channelName string) error {
	unloadCtx := ctx.SubContext("Unload Channel: " + channelName)

	_, childChannels := channelConf.Get(unloadCtx, constants.CONF_ENGINE_CHANNELS)
	if childChannels {
		err := chanMgr.unloadChildChannels(unloadCtx, channelConf)
		if err != nil {
			return errors.WrapError(unloadCtx, err)
		}
	}

	parentChannelName, ok := channelConf.GetString(unloadCtx, constants.CONF_ENGINE_PARENTCHANNEL)
	if !ok {
		return errors.ThrowError(unloadCtx, errors.CORE_ERROR_MISSING_CONF, "conf", constants.CONF_ENGINE_PARENTCHANNEL)
	}
	parentChannel, ok := chanMgr.channelStore[parentChannelName]
	if ok {
		err := parentChannel.RemoveChild(unloadCtx, channelName, channelConf)
		if err != nil {
			return errors.WrapError(unloadCtx, err)
		}
	}
	delete(chanMgr.channelStore, channelName)
	//log.Trace(ctx, "Creating channel", "Name:", channelName)
	log.Info(unloadCtx, "Unloaded channel", "Name:", channelName)
	return nil
}

func (chanMgr *channelManager) startModuleInstanceChannels(ctx core.ServerContext, mod *serverModule) error {
	for chanName, _ := range mod.channels {
		chn, _ := chanMgr.channelStore[chanName]
		if err := chanMgr.startChannel(ctx, chanName, chn); err != nil {
			return err
		}
	}
	return nil
}
