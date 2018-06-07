package websocket

import (
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

type wsEngineProxy struct {
	engine *wsEngine
}

func (eng *wsEngineProxy) GetRootChannel(ctx core.ServerContext) server.Channel {
	return &wsChannelProxy{channel: eng.engine.rootChannel}
}

func (proxy *wsEngineProxy) Reference() core.ServerElement {
	return &wsEngineProxy{engine: proxy.engine}
}

func (proxy *wsEngineProxy) GetProperty(name string) interface{} {
	return nil
}

func (proxy *wsEngineProxy) GetName() string {
	return proxy.engine.name
}
func (proxy *wsEngineProxy) GetType() core.ServerElementType {
	return core.ServerElementEngine
}
