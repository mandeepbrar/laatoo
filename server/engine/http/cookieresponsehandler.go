package http

import (
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"
	"laatoo/server/engine/http/net"
	"net/http"
)

type cookiesResponseHandler struct {
	svrContext core.ServerContext
	jsonCodec  core.Codec
}

func (rh *cookiesResponseHandler) Initialize(ctx core.ServerContext, conf config.Config) error {
	rh.svrContext = ctx
	rh.jsonCodec, _ = ctx.GetCodec("json")
	return nil
}

func (rh *cookiesResponseHandler) HandleResponse(ctx core.RequestContext, resp *core.Response, handlererr error) error {
	log.Trace(ctx, "Returning request with cookies response handler", "resp", resp)
	return handleResponse(ctx, resp, rh.jsonCodec, rh.handleMetaInfo, handlererr)
}

func (rh *cookiesResponseHandler) Reference() core.ServerElement {
	anotherref := rh
	return anotherref
}

func (rh *cookiesResponseHandler) GetProperty(name string) interface{} {
	return nil
}

func (rh *cookiesResponseHandler) GetName() string {
	return "CookiesResponseHandler"
}
func (rh *cookiesResponseHandler) GetType() core.ServerElementType {
	return core.ServerElementServiceResponseHandler
}
func (rh *cookiesResponseHandler) GetContext() core.ServerContext {
	return rh.svrContext
}

func (rh *cookiesResponseHandler) handleMetaInfo(ctx core.RequestContext, webctx net.WebContext, info map[string]interface{}) error {
	log.Error(ctx, "cookies to be sent", "info", info)
	if info != nil {
		for key, val := range info {
			switch key {
			case core.ContentType:
				webctx.SetHeader(core.ContentType, fmt.Sprint(val))
			case core.ContentEncoding:
				webctx.SetHeader(core.ContentEncoding, fmt.Sprint(val))
			case core.LastModified:
				webctx.SetHeader(core.LastModified, fmt.Sprint(val))
			default:
				{
					cookie := new(http.Cookie)
					cookie.Name = key
					if val == "<delete>" {
						cookie.MaxAge = 0
					} else {
						cookie.Value = fmt.Sprint(val)
					}
					cookie.HttpOnly = true
					//cookie.Expires = time.Now().Add(24 * time.Hour)
					webctx.SetCookie(cookie)
				}
			}
		}
	}
	return nil
}
