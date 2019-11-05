package http

import (
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"
	"laatoo/server/engine/http/net"
	"strings"
)

type defaultResponseHandler struct {
	svrContext core.ServerContext
}

func DefaultResponseHandler(ctx core.ServerContext) *defaultResponseHandler {
	return &defaultResponseHandler{}
}
func (rh *defaultResponseHandler) Initialize(ctx core.ServerContext, conf config.Config) error {
	rh.svrContext = ctx
	return nil
}

func (rh *defaultResponseHandler) HandleResponse(ctx core.RequestContext, resp *core.Response, handlingError error) error {
	log.Trace(ctx, "Returning request with default response handler", "resp", resp)
	return handleResponse(ctx, resp, rh.handleMetaInfo)
}

func (rh *defaultResponseHandler) Reference() core.ServerElement {
	anotherref := rh
	return anotherref
}
func (rh *defaultResponseHandler) GetContext() core.ServerContext {
	return rh.svrContext
}
func (rh *defaultResponseHandler) GetProperty(name string) interface{} {
	return nil
}

func (rh *defaultResponseHandler) GetName() string {
	return "DefaultResponseHandler"
}
func (rh *defaultResponseHandler) GetType() core.ServerElementType {
	return core.ServerElementServiceResponseHandler
}

func (rh *defaultResponseHandler) handleMetaInfo(ctx core.RequestContext, webctx net.WebContext, info map[string]interface{}) error {
	if info != nil {
		keyNames := make([]string, len(info))
		i := 0
		for key, val := range info {
			webctx.SetHeader(key, fmt.Sprint(val))
			keyNames[i] = key
			i++
		}
		webctx.SetHeader("Access-Control-Expose-Headers", strings.Join(keyNames, ","))
	}
	return nil
}
