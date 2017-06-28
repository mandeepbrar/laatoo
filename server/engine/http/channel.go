package http

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

type httpChannelProxy struct {
	channel *httpChannel
}

func (channel *httpChannelProxy) Serve(ctx core.ServerContext, svc server.Service, channelConfig config.Config) error {
	return channel.channel.serve(ctx, svc, channelConfig)
}

func (channel *httpChannelProxy) Child(ctx core.ServerContext, name string, channelConfig config.Config) (server.Channel, error) {
	childChannel := channel.channel.group(ctx, name, channelConfig)
	proxy := &httpChannelProxy{channel: childChannel}
	return proxy, nil
}

func (proxy *httpChannelProxy) Reference() core.ServerElement {
	return &httpChannelProxy{channel: proxy.channel}
}

func (proxy *httpChannelProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *httpChannelProxy) GetName() string {
	return proxy.channel.name
}
func (proxy *httpChannelProxy) GetType() core.ServerElementType {
	return core.ServerElementChannel
}
