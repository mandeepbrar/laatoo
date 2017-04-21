package http

import (
	"laatoo/framework/core/common"
	//	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

type httpEngineProxy struct {
	*common.Context
	engine *httpEngine
}

func (eng *httpEngineProxy) GetRootChannel(ctx core.ServerContext) server.Channel {
	engCtx := ctx.GetServerElement(core.ServerElementServer).NewCtx("Channel:" + eng.engine.name)
	return &httpChannelProxy{engCtx.(*common.Context), eng.engine.rootChannel}
}
