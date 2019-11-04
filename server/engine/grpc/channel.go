package grpc

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

type grpcChannelProxy struct {
	channel *grpcChannel
}

func (channel *grpcChannelProxy) Serve(ctx core.ServerContext) error {
	return channel.channel.serve(ctx)
}

func (channel *grpcChannelProxy) GetServiceName() string {
	return channel.channel.svcName
}

func (channel *grpcChannelProxy) Child(ctx core.ServerContext, name string, channelConfig config.Config) (elements.Channel, error) {
	log.Trace(ctx, "Creating child channel ", "Parent", channel.channel.name, "Name", name)
	childChannel, err := channel.channel.child(ctx, name, channelConfig)
	if err != nil {
		return nil, err
	}
	proxy := &grpcChannelProxy{channel: childChannel}
	return proxy, nil

	return nil, errors.NotImplemented(ctx, "Child")
}

func (proxy *grpcChannelProxy) RemoveChild(ctx core.ServerContext, name string, channelConfig config.Config) error {
	return nil
}

func (proxy *grpcChannelProxy) Reference() core.ServerElement {
	return &grpcChannelProxy{channel: proxy.channel}
}
func (proxy *grpcChannelProxy) GetContext() core.ServerContext {
	return proxy.channel.svrContext
}
func (proxy *grpcChannelProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *grpcChannelProxy) GetName() string {
	return proxy.channel.name
}
func (proxy *grpcChannelProxy) GetType() core.ServerElementType {
	return core.ServerElementChannel
}

func (proxy *grpcChannelProxy) Destruct(ctx core.ServerContext, parentChannel elements.Channel) error {
	return proxy.channel.destruct(ctx, parentChannel.(*grpcChannelProxy).channel)
}
