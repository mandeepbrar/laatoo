package websocket

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/constants"
)

type wsChannel struct {
	name         string
	disabled     bool
	config       config.Config
	svcName      string
	svc          elements.Service
	engine       *wsEngine
	staticValues map[string]interface{}
	respHandler  elements.ServiceResponseHandler
}

func (channel *wsChannel) configure(ctx core.ServerContext) error {
	channel.disabled, _ = channel.config.GetBool(ctx, constants.CONF_WSENGINE_DISABLEROUTE)
	if channel.disabled {
		return nil
	}

	////build value parameters
	staticValues := make(map[string]interface{})
	staticValuesConf, ok := channel.config.GetSubConfig(ctx, constants.CONF_WSENGINE_STATICVALUES)
	if ok {
		values := staticValuesConf.AllConfigurations(ctx)
		for _, paramname := range values {
			staticValues[paramname], _ = staticValuesConf.Get(ctx, paramname)
		}
	}
	channel.staticValues = staticValues
	return nil
}

func (channel *wsChannel) child(ctx core.ServerContext, name string, channelConfig config.Config) (*wsChannel, error) {
	ctx = ctx.SubContext("Channel " + name)
	svc, found := channelConfig.GetString(ctx, constants.CONF_CHANNEL_SERVICE)
	if !found {
		return nil, errors.BadConf(ctx, constants.CONF_CHANNEL_SERVICE)
	}
	log.Trace(ctx, "Creating child channel ", "Parent", channel.name, "Name", name, "Service", svc)
	childChannel := &wsChannel{name: name, config: channelConfig, engine: channel.engine, svcName: svc}
	err := childChannel.configure(ctx)
	if err != nil {
		return nil, err
	}
	return childChannel, nil
}

func (channel *wsChannel) serve(ctx core.ServerContext) error {
	ctx = ctx.SubContext("Serve")
	if !channel.disabled {
		log.Trace(ctx, "Channel config", "name", channel.name, "config", channel.config)

		svcManager := ctx.GetServerElement(core.ServerElementServiceManager).(elements.ServiceManager)
		svc, err := svcManager.GetService(ctx, channel.svcName)
		if err != nil {
			return err
		}
		channel.svc = svc

		handler := ctx.GetServerElement(core.ServerElementServiceResponseHandler)
		if handler != nil {
			channel.respHandler = handler.(elements.ServiceResponseHandler)
		} else {
			channel.respHandler = DefaultResponseHandler(ctx, channel.engine.codec)
		}
	}
	return nil
}

func (channel *wsChannel) removeChild(ctx core.ServerContext, name string, channelConfig config.Config) error {
	channel.disabled = true
	return nil
}
