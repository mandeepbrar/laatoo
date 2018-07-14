package http

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	//"laatoo/sdk/server/log"
	"laatoo/sdk/server/elements"
)

func NewEngine(ctx core.ServerContext, name string, conf config.Config) (elements.ServerElementHandle, elements.Engine) {
	eng := &httpEngine{ssl: false, name: name, conf: conf}
	proxy := &httpEngineProxy{engine: eng}
	eng.proxy = proxy
	return eng, proxy
}
