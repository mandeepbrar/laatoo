package http

import (
	"laatoo/sdk/core"
	//"laatoo/sdk/log"
	"laatoo/sdk/server"
)

func NewEngine(ctx core.ServerContext, name string) (server.ServerElementHandle, server.Engine) {
	eng := &httpEngine{ssl: false, name: name}
	proxy := &httpEngineProxy{engine: eng}
	eng.proxy = proxy
	return eng, proxy
}
