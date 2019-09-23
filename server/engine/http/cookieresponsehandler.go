package http

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/log"
	"laatoo/server/engine/http/net"
	"net/http"
	"fmt"
)

type cookiesResponseHandler struct {
}

func (rh *cookiesResponseHandler) Initialize(conf config.Config) error {
	return nil
}

func (rh *cookiesResponseHandler) HandleResponse(ctx core.RequestContext, resp *core.Response) error {
	log.Trace(ctx, "Returning request with default response handler", "resp", resp)
	return handleResponse(ctx, resp, rh.handleMetaInfo)
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


func(rh *cookiesResponseHandler) handleMetaInfo(ctx net.WebContext, info map[string]interface{}) error {
	if(info != nil) {
		for key, val := range info {
			cookie := new(http.Cookie)
			cookie.Name = key
			cookie.Value = fmt.Sprint(val)
			cookie.HttpOnly = true
			//cookie.Expires = time.Now().Add(24 * time.Hour)		
			ctx.SetCookie(cookie)
		}
	}
	return nil
}

