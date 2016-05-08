// +build appengine

package common

import (
	glctx "golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/cloud"
	"net/http"
)

func GetAppengineContext(ctx *Context) glctx.Context {
	if ctx == nil || ctx.GaeReq == nil {
		return nil
	}
	return appengine.NewContext(ctx.GaeReq)
}

func GetCloudContext(ctx *Context, scope string) glctx.Context {
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

func HttpClient(ctx *Context) *http.Client {
	appenginectx := GetAppengineContext(ctx)
	return &http.Client{
		Transport: &urlfetch.Transport{
			Context: appenginectx,
			AllowInvalidServerCertificate: true,
		},
	}
}

func GetOAuthContext(ctx *Context) glctx.Context {
	return GetAppengineContext(ctx)
}
