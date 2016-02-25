// +build appengine

package context

import (
	"github.com/labstack/echo"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/cloud"
	"net/http"
)

func GetAppengineContext(ctx *echo.Context) context.Context {
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

func GetCloudContext(ctx *echo.Context, scope string) context.Context {
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

func HttpClient(ctx *echo.Context) *http.Client {
	appenginectx := GetAppengineContext(ctx)
	return &http.Client{
		Transport: &urlfetch.Transport{
			Context: appenginectx,
			AllowInvalidServerCertificate: true,
		},
	}
}
