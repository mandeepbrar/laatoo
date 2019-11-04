package core

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
)

type channelManagerProxy struct {
	manager *channelManager
}

func (proxy *channelManagerProxy) Reference() core.ServerElement {
	return &channelManagerProxy{manager: proxy.manager}
}
func (proxy *channelManagerProxy) GetChannel(ctx core.ServerContext, name string) (elements.Channel, bool) {
	if proxy.manager != nil && proxy.manager.channelStore != nil {
		channel, ok := proxy.manager.channelStore[name]
		return channel, ok
	}
	return nil, false
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
func (proxy *channelManagerProxy) GetContext() core.ServerContext {
	return proxy.manager.svrContext
}
