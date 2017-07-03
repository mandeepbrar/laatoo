package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"
)

type channelManagerProxy struct {
	manager *channelManager
}

func (proxy *channelManagerProxy) Reference() core.ServerElement {
	return &channelManagerProxy{manager: proxy.manager}
}
func (proxy *channelManagerProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *channelManagerProxy) GetName() string {
	return proxy.manager.name
}
func (proxy *channelManagerProxy) GetType() core.ServerElementType {
	return core.ServerElementChannelManager
}

func (cm *channelManagerProxy) Serve(ctx core.ServerContext, channelName string, svc server.Service, channelConfig config.Config) error {
	channel, ok := cm.manager.channelStore[channelName]
	if ok {
		return channel.Serve(ctx, svc, channelConfig)
	} else {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "No such channel", channelName)
	}
}
