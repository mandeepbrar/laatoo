package http

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/server"
	"laatoo/server/common"
)

type httpChannelProxy struct {
	*common.Context
	channel *httpChannel
}

func (channel *httpChannelProxy) Serve(ctx core.ServerContext, svc server.Service, channelConfig config.Config) error {
	return channel.channel.serve(ctx, svc, channelConfig)
}

func (channel *httpChannelProxy) Child(ctx core.ServerContext, name string, channelConfig config.Config) (server.Channel, error) {
	childChannel := channel.channel.group(ctx, name, channelConfig)
	childCtx := channel.NewCtx("Channel:" + name)
	proxy := &httpChannelProxy{Context: childCtx.(*common.Context), channel: childChannel}
	return proxy, nil
}
