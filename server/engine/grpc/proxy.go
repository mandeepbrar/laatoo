package grpc

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
)

type grpcEngineProxy struct {
	engine *grpcEngine
}

func (eng *grpcEngineProxy) GetRootChannel(ctx core.ServerContext) elements.Channel {
	return &grpcChannelProxy{channel: eng.engine.rootChannel}
}

func (proxy *grpcEngineProxy) Reference() core.ServerElement {
	return &grpcEngineProxy{engine: proxy.engine}
}

func (proxy *grpcEngineProxy) GetProperty(name string) interface{} {
	return nil
}

func (proxy *grpcEngineProxy) GetName() string {
	return proxy.engine.name
}
func (proxy *grpcEngineProxy) GetType() core.ServerElementType {
	return core.ServerElementEngine
}
func (proxy *grpcEngineProxy) GetRequestParams(ctx core.RequestContext) map[string]interface{} {
	return nil
}
func (proxy *grpcEngineProxy) GetContext() core.ServerContext {
	return proxy.engine.svrContext
}
