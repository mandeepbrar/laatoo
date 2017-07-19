package core

import "laatoo/sdk/core"

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
