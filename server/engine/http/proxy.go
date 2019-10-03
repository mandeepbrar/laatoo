package http

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/server/engine/http/net"
)

type httpEngineProxy struct {
	engine *httpEngine
}

func (eng *httpEngineProxy) GetRootChannel(ctx core.ServerContext) elements.Channel {
	return &httpChannelProxy{channel: eng.engine.rootChannel}
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

func (proxy *httpEngineProxy) GetContext() core.ServerContext {
	return proxy.engine.svrContext
}

func (proxy *httpEngineProxy) GetRequestParams(ctx core.RequestContext) map[string]interface{} {
	params := make(map[string]interface{})
	webCtx := ctx.EngineRequestContext().(net.WebContext)
	paramNames := webCtx.GetRouteParamNames()
	for _, pname := range paramNames {
		pval := webCtx.GetRouteParam(pname)
		params[pname] = pval
	}
	return params
}
