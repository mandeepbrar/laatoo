package http

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
)

type httpChannelProxy struct {
	channel *httpChannel
}

func (channel *httpChannelProxy) Serve(ctx core.ServerContext) error {
	return channel.channel.serve(ctx)
}

func (channel *httpChannelProxy) GetServiceName() string {
	return channel.channel.svcName
}

func (channel *httpChannelProxy) Child(ctx core.ServerContext, name string, channelConfig config.Config) (server.Channel, error) {
	log.Trace(ctx, "Creating child channel ", "Parent", channel.channel.name, "Name", name)
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
