package http

import (
	"laatoo/framework/core/common"
	"laatoo/sdk/core"
	//"laatoo/sdk/log"
	"laatoo/sdk/server"
)

func NewEngine(ctx core.ServerContext, name string) (server.ServerElementHandle, server.Engine) {
	eng := &httpEngine{ssl: false, name: name}
	engCtx := ctx.GetServerElement(core.ServerElementServer).NewCtx("Engine" + name)
	proxy := &httpEngineProxy{Context: engCtx.(*common.Context), engine: eng}
	eng.proxy = proxy
	return eng, proxy
}
