package websocket

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	//"laatoo/sdk/log"
	"laatoo/sdk/server"
)

func NewEngine(ctx core.ServerContext, name string, conf config.Config) (server.ServerElementHandle, server.Engine) {
	eng := &wsEngine{name: name, conf: conf}
	proxy := &wsEngineProxy{engine: eng}
	eng.proxy = proxy
	return eng, proxy
}
