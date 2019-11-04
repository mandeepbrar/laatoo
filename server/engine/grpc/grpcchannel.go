package grpc

import (
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/constants"

	"google.golang.org/grpc"
)

type grpcChannel struct {
	name          string
	disabled      bool
	config        config.Config
	svcName       string
	path          string
	svc           elements.Service
	engine        *grpcEngine
	handleRequest func(core.ServerContext, grpc.ServerStream) error
	staticValues  map[string]interface{}
	respHandler   elements.ServiceResponseHandler
	svrContext    core.ServerContext
}

func (channel *grpcChannel) configure(ctx core.ServerContext) error {
	channel.disabled, _ = channel.config.GetBool(ctx, constants.CONF_GRPCENGINE_DISABLEROUTE)
	if channel.disabled {
		return nil
	}

	////build value parameters
	staticValues := make(map[string]interface{})
	staticValuesConf, ok := channel.config.GetSubConfig(ctx, constants.CONF_GRPCENGINE_STATICVALUES)
	if ok {
		values := staticValuesConf.AllConfigurations(ctx)
		for _, paramname := range values {
			staticValues[paramname], _ = staticValuesConf.Get(ctx, paramname)
		}
	}
	channel.staticValues = staticValues

	return nil
}

func (channel *grpcChannel) child(ctx core.ServerContext, name string, channelConfig config.Config) (*grpcChannel, error) {
	ctx = ctx.SubContext("Channel " + name)

	svc, found := channelConfig.GetString(ctx, constants.CONF_CHANNEL_SERVICE)
	if !found {
		return nil, errors.BadConf(ctx, constants.CONF_CHANNEL_SERVICE)
	}
	log.Trace(ctx, "Creating child channel ", "Parent", channel.name, "Name", name, "Service", svc)

	path, ok := channelConfig.GetString(ctx, constants.CONF_GRPCENGINE_PATH)
	if !ok {
		errors.BadConf(ctx, constants.CONF_GRPCENGINE_PATH)
	}

	route := fmt.Sprintf("%s/%s", channel.path, path)

	log.Error(ctx, "Adding grpc channel", "route", route)

	childChannel := &grpcChannel{name: name, config: channelConfig, engine: channel.engine, svcName: svc, path: path, svrContext: ctx}
	err := childChannel.configure(ctx)
	if err != nil {
		return nil, err
	}

	channel.engine.addRoute(ctx, route, childChannel)

	return childChannel, nil
}

func (channel *grpcChannel) destruct(ctx core.ServerContext, parentChannel *grpcChannel) error {
	return nil
}
