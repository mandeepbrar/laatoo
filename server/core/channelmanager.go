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

	//secondPass map[string]config.Config

	//store for service factory in an application
	channelStore   map[string]elements.Channel
	channelConfs   map[string]config.Config
	parentChannels map[string]string
	serviceManager elements.ServiceManager
	svrContext     core.ServerContext
}

func (chanMgr *channelManager) Initialize(ctx core.ServerContext, conf config.Config) error {

	log.Trace(ctx, "Channel manager initialize")

	err := chanMgr.processChannelsFromServerConf(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	baseDir, _ := ctx.GetString(config.BASEDIR)

	modManager := ctx.GetServerElement(core.ServerElementModuleManager).(*moduleManagerProxy).modMgr
	svcMgr := ctx.GetServerElement(core.ServerElementServiceManager).(elements.ServiceManager)
	chanMgr.serviceManager = svcMgr

	if err = modManager.loadChannels(ctx, chanMgr.storeConf); err != nil {
		return err
	}

	if err := chanMgr.processChannelsFromFolder(ctx, baseDir); err != nil {
		return errors.WrapError(ctx, err)
	}

	if _, err := chanMgr.createChannels(ctx, chanMgr.channelConfs); err != nil {
		return errors.WrapError(ctx, err)
	}

	return nil
}

func (chanMgr *channelManager) processChannelsFromFolder(ctx core.ServerContext, folder string) error {
	objs, err := common.ProcessDirectoryFiles(ctx, folder, constants.CONF_CHANNELS, true)
	if err != nil {
		return err
	}

	if err = common.ProcessObjects(ctx, objs, chanMgr.storeConf); err != nil {
		return err
	}
	return nil
}

func (chanMgr *channelManager) processChannelsFromServerConf(ctx core.ServerContext, conf config.Config) error {
	channelNames := conf.AllConfigurations(ctx)
	for _, channelName := range channelNames {
		channelConf, err, _ := common.ConfigFileAdapter(ctx, conf, channelName)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		name, ok := conf.GetString(ctx, constants.CONF_OBJECT_NAME)
		if !ok {
			name = channelName
		}
		chanMgr.storeConf(ctx, channelConf, name)
	}
	return nil
}

func (chanMgr *channelManager) storeConf(ctx core.ServerContext, channelConf config.Config, channelName string) error {
	arr, ok := channelConf.GetConfigArray(ctx, constants.CONF_CHANNELS)
	if ok {
		for _, conf := range arr {
			name, ok := conf.GetString(ctx, constants.CONF_OBJECT_NAME)
			if !ok {
				return errors.MissingConf(ctx, constants.CONF_OBJECT_NAME, "Channels", channelName)
			}
			//			chanMgr.channelConfs[name] = channelConf
			/*
				/*_, childChannels := channelConf.Get(createCtx, constants.CONF_ENGINE_CHANNELS)
				if childChannels {
					err := chanMgr.createChannels(createCtx, channelConf)
					if err != nil {
						return errors.WrapError(createCtx, err)
					}
				}*/

			/*			channelConf, err, _ := common.ConfigFileAdapter(ctx, channelsConf, channelName)
						if err != nil {
							return errors.WrapError(ctx, err)
						}*/
			chanMgr.channelConfs[name] = conf
		}
	} else {
		chanMgr.channelConfs[channelName] = channelConf
	}
	return nil
}

func (chanMgr *channelManager) createChannels(ctx core.ServerContext, channelConfs map[string]config.Config) (map[string]elements.Channel, error) {
	pendingConfs := make(map[string]config.Config)
	channelsReturned := make(map[string]elements.Channel)
	for channelName, channelConf := range channelConfs {
		channel, err := chanMgr.createChannel(ctx, channelConf, channelName)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		if channel == nil {
			pendingConfs[channelName] = channelConf
		}
		channelsReturned[channelName] = channel
	}
	if len(pendingConfs) != 0 {
		if len(pendingConfs) == len(channelConfs) {
			return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Missing Parents for Channels", pendingConfs)
		}
		channelsCreated, err := chanMgr.createChannels(ctx, pendingConfs)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		for k, v := range channelsCreated {
			channelsReturned[k] = v
		}
	}
	return channelsReturned, nil
}

func (chanMgr *channelManager) Start(ctx core.ServerContext) error {
	/*
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
		}*/

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
	log.Info(chanCtx, "Starting channels ", "chanName", chanName)
	svcName := channel.GetServiceName()
	if svcName != "" {
		svcProxy, err := chanMgr.serviceManager.GetService(ctx, svcName)
		if err != nil {
			return err
		}
		proxy := svcProxy.(*serviceProxy)
		svcServeCtx := proxy.svc.svrContext.newContext("Serve: " + proxy.svc.name)

		respHandlerConf, rhpresent := channel.GetContext().Get(constants.CONF_SERVICE_RH)
		if !rhpresent {
			respHandlerConf, rhpresent = svcServeCtx.Get(constants.CONF_SERVICE_RH)
		}

		if rhpresent {
			rhConf, ok := respHandlerConf.(config.Config)
			if !ok {
				return errors.BadConf(ctx, constants.CONF_SERVICE_RH, "Info", "Config object required")
			}
			respHandlerName, ok := rhConf.GetString(ctx, constants.CONF_OBJECT_NAME)
			if !ok {
				return errors.MissingConf(ctx, constants.CONF_OBJECT_NAME, "Config", constants.CONF_SERVICE_RH)
			}
			respHandler, err := svcServeCtx.CreateObject(respHandlerName)
			if err != nil {
				return errors.BadConf(ctx, "Response handler object not found", constants.CONF_SERVICE_RH, respHandlerName)
			}
			rh, ok := respHandler.(elements.ServiceResponseHandler)
			if !ok {
				return errors.BadConf(ctx, "Provided objects is not response handler ", constants.CONF_SERVICE_RH, respHandlerName)
			}
			err = rh.Initialize(svcServeCtx, rhConf)
			if err != nil {
				return errors.WrapError(svcServeCtx, err)
			}
			log.Trace(ctx, "Setting response handler ", "respHandlerName", respHandlerName, "svc", svcName)
			svcServeCtx.setElements(core.ContextMap{core.ServerElementServiceResponseHandler: rh})
		}

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
/*
func (chanMgr *channelManager) reviewMissingChannels(ctx core.ServerContext, chansToReview map[string]config.Config) error {
	for channelName, channelConf := range chansToReview {
		if err := chanMgr.createChannel(ctx, channelConf, channelName); err != nil {
			return errors.WrapError(ctx, err)
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
}*/

func (chanMgr *channelManager) createChannel(ctx core.ServerContext, channelConf config.Config, channelName string) (elements.Channel, error) {
	ctx = ctx.SubContext("Create channel " + channelName)
	chann, ok := chanMgr.channelStore[channelName]
	if ok {
		return chann, nil
	}

	createCtx := ctx.(*serverContext).newContext("Channel: " + channelName)
	log.Info(createCtx, "Creating channel with conf ", "channelName", channelName, "conf", channelConf)

	respHandlerConf, ok := channelConf.GetSubConfig(ctx, constants.CONF_SERVICE_RH)
	if ok {
		log.Trace(ctx, "Found response handler for channel", "name", channelName, "resphandler", respHandlerConf)
		createCtx.Set(constants.CONF_SERVICE_RH, respHandlerConf)
	}

	parentChannelName, ok := channelConf.GetString(ctx, constants.CONF_ENGINE_PARENTCHANNEL)
	if !ok {
		return nil, errors.ThrowError(createCtx, errors.CORE_ERROR_MISSING_CONF, "channelname", channelName, "channelConf", channelConf, "conf", constants.CONF_ENGINE_PARENTCHANNEL)
	}
	parentChannel, ok := chanMgr.channelStore[parentChannelName]
	if !ok {
		//chanMgr.secondPass[channelName] = channelConf
		//"Could not find parent channel"
		return nil, nil
		//return errors.ThrowError(createCtx, errors.CORE_ERROR_BAD_CONF, "conf", constants.CONF_ENGINE_PARENTCHANNEL)
	}
	log.Trace(createCtx, "Found parent channel ", "Parent Channel Name", parentChannelName, "Found", ok)
	channel, err := parentChannel.Child(createCtx, channelName, channelConf)
	if err != nil {
		return nil, errors.WrapError(createCtx, err)
	}
	chanMgr.channelStore[channelName] = channel
	chanMgr.parentChannels[channelName] = parentChannelName
	/*_, childChannels := channelConf.Get(createCtx, constants.CONF_ENGINE_CHANNELS)
	if childChannels {
		err := chanMgr.createChannels(createCtx, channelConf)
		if err != nil {
			return errors.WrapError(createCtx, err)
		}
	}*/
	//log.Trace(ctx, "Creating channel", "Name:", channelName)
	log.Info(createCtx, "Created channel", "Name:", channelName)
	return channel, nil
}

func (chanMgr *channelManager) unloadModuleChannels(ctx core.ServerContext, mod *serverModule) error {
	ctx = ctx.SubContext("unload channels")
	if err := common.ProcessObjects(ctx, mod.channels, chanMgr.unloadChannel); err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

/*
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
*/
func (chanMgr *channelManager) unloadChannel(ctx core.ServerContext, channelConf config.Config, channelName string) error {
	unloadCtx := ctx.SubContext("Unload Channel: " + channelName)

	channel, ok := chanMgr.channelStore[channelName]
	if !ok {
		return nil
	}

	for childChann, parentChann := range chanMgr.parentChannels {
		if parentChann == channelName {
			childConf, ok := chanMgr.channelConfs[childChann]
			if !ok {
				return errors.NotImplemented(ctx, "conf not found of child for unload", "channel name", channelName, "child", childChann)
			}
			err := chanMgr.unloadChannel(unloadCtx, childConf, childChann)
			if err != nil {
				return errors.WrapError(unloadCtx, err)
			}
		}
	}

	parentChannelName, ok := chanMgr.parentChannels[channelName]
	if !ok {
		//this should not happen
		return errors.NotImplemented(ctx, "parent channel not found for unload", "channel name", channelName)
	}
	parentChannel, ok := chanMgr.channelStore[parentChannelName]
	if ok {
		err := channel.Destruct(unloadCtx, parentChannel)
		if err != nil {
			return errors.WrapError(unloadCtx, err)
		}
	} else {
		//this should not happen
		return errors.NotImplemented(ctx, "parent channel not found for unload", "channel name", channelName, "parentChannelName", parentChannelName)
	}
	delete(chanMgr.channelStore, channelName)
	delete(chanMgr.channelConfs, channelName)
	delete(chanMgr.parentChannels, channelName)
	//log.Trace(ctx, "Creating channel", "Name:", channelName)
	log.Info(unloadCtx, "Unloaded channel", "Name:", channelName)
	return nil
}

func (chanMgr *channelManager) startModuleInstanceChannels(ctx core.ServerContext, mod *serverModule) error {
	ctx = ctx.SubContext("Start module channels " + mod.name)
	log.Info(ctx, "Starting channels for module instance ", "module name", mod.name)
	for chanName, _ := range mod.channels {
		chn, _ := chanMgr.channelStore[chanName]
		log.Info(ctx, "Starting channels for module instance ", "chanName", chanName)
		if err := chanMgr.startChannel(ctx, chanName, chn); err != nil {
			return err
		}
	}
	return nil
}

func (chanMgr *channelManager) createModuleChannels(ctx core.ServerContext, mod *serverModule) error {
	//channels := make(map[string]elements.Channel)
	if mod.channels != nil {
		_, err := chanMgr.createChannels(ctx, mod.channels)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		/*for chanName, chanConf := range mod.channels {
		chnCtx := ctx.SubContext("Create channel:" + chanName)
		log.Info(chnCtx, "Creating channel ", "mod name", mod.name, "channel", chanName)
		channel, err := chanMgr.createChannel(chnCtx, chanConf, chanName)
		if err != nil {
			return errors.WrapError(chnCtx, err)
		}
		channels[chanName] = channel
		*/
		for k, v := range mod.channels {
			chanMgr.channelConfs[k] = v
		}
	}
	//channels do not require initialization... so leaving the stored channels here
	return nil
}
