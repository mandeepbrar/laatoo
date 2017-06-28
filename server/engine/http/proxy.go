package http

import (
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

type httpEngineProxy struct {
	engine *httpEngine
}

func (eng *httpEngineProxy) GetRootChannel(ctx core.ServerContext) server.Channel {
	return &httpChannelProxy{eng.engine.rootChannel}
}

func (proxy *httpEngineProxy) Reference() core.ServerElement {
	return &httpEngineProxy{engine: proxy.engine}
}

func (proxy *httpEngineProxy) GetProperty(name string) interface{} {
	return nil
}

func (proxy *httpEngineProxy) GetName() string {
	return proxy.engine.name
}
func (proxy *httpEngineProxy) GetType() core.ServerElementType {
	return core.ServerElementEngine
}
