// +build appengine

package laatoocore

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	glctx "golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/cloud"
	"laatoosdk/log"
	"net/http"
	"sync"
)

//Create a new server
func NewServer(configName string, serverType string) (*Server, error) {
	//initialize router
	router := echo.New()
	// Middleware
	router.Use(mw.Logger())
	router.Use(mw.Recover())
	router.Use(mw.Gzip())
	ctx := &Context{Context: echo.NewContext(nil, nil, router)}
	server := &Server{ServerType: serverType}
	server.InitServer(ctx, configName, router)
	http.Handle("/", router)
	if server.ServerType == CONF_SERVERTYPE_GOOGLEAPP {
		log.Logger.Error(ctx, "core.appengine.warmup", "setting up router for warmup")
		var req *Context
		var once sync.Once
		warmupFunc := func() {
			log.Logger.Error(req, "core.appengine.warmup", "starting server")
			server.Start(req)
		}
		router.Use(func(ctx *echo.Context) error {
			req = &Context{Context: ctx}
			once.Do(warmupFunc)
			return nil
		})
		router.Get("/_ah/warmup", func(ctx *echo.Context) error {
			req = &Context{Context: ctx}
			once.Do(warmupFunc)
			return nil
		})
	}
	return server, nil
}

func GetAppengineContext(ctx *Context) glctx.Context {
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
