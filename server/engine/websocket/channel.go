package websocket

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"
)

type wsChannelProxy struct {
	channel *wsChannel
}

func (channel *wsChannelProxy) Serve(ctx core.ServerContext) error {
	return channel.channel.serve(ctx)
}

func (channel *wsChannelProxy) GetServiceName() string {
	return channel.channel.svcName
}

func (channel *wsChannelProxy) Child(ctx core.ServerContext, name string, channelConfig config.Config) (server.Channel, error) {
	/*log.Trace(ctx, "Creating child channel ", "Parent", channel.channel.name, "Name", name)
	childChannel, err := channel.channel.child(ctx, name, channelConfig)
	if err != nil {
		return nil, err
	}
	proxy := &wsChannelProxy{channel: childChannel}*/
	return nil, errors.NotImplemented(ctx, "Child")
}

func (proxy *wsChannelProxy) Reference() core.ServerElement {
	return &wsChannelProxy{channel: proxy.channel}
}

func (proxy *wsChannelProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *wsChannelProxy) GetName() string {
	return proxy.channel.name
}
func (proxy *wsChannelProxy) GetType() core.ServerElementType {
	return core.ServerElementChannel
}
