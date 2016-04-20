// +build appengine

package server

import (
	glctx "golang.org/x/net/context"
	"net/http"
)

const (
	SERVER_TYPE = "CONF_SERVERTYPE_GOOGLEAPP"
)

func GetAppengineContext(ctx RequestContext) glctx.Context {
	if ctx == nil || ctx.Request() == nil {
		return nil
	}
	/*	echoContext, ok := ctx.(*echo.Context)
		if !ok {
			appengineCtx, ok := ctx.(context.Context)
			if ok {
				return appengineCtx
			}
			return nil
		}*/
	return appengine.NewContext(ctx.Request())

}

func GetCloudContext(ctx RequestContext, scope string) glctx.Context {
	appenginectx := GetAppengineContext(ctx)
	hc := &http.Client{
		Transport: &oauth2.Transport{
			Source: google.AppEngineTokenSource(appenginectx, scope),
			Base: &urlfetch.Transport{
				Context: appenginectx,
			},
		},
	}
	return cloud.NewContext(appengine.AppID(appenginectx), hc)
}

func HttpClient(ctx RequestContext) *http.Client {
	appenginectx := GetAppengineContext(ctx)
	return &http.Client{
		Transport: &urlfetch.Transport{
			Context: appenginectx,
			AllowInvalidServerCertificate: true,
		},
	}
}

func GetOAuthContext(ctx Context) glctx.Context {
	return GetAppengineContext()
}
