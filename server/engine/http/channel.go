package http

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/log"
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

func (channel *httpChannelProxy) Child(ctx core.ServerContext, name string, channelConfig config.Config) (elements.Channel, error) {
	log.Trace(ctx, "Creating child channel ", "Parent", channel.channel.name, "Name", name)
	childChannel, err := channel.channel.child(ctx, name, channelConfig)
	if err != nil {
		return nil, err
	}
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
