// +build appengine

package server

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	glctx "golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/cloud"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"net/http"
	"sync"
)

//Create a new server
func NewServer(configName string) (*Server, error) {
	ctx := &Context{Context: echo.NewContext(nil, nil, router)}
	server := &Server{ServerType: CONF_SERVERTYPE_GOOGLEAPP}
	err := server.InitServer(ctx, configName, router)
	if err != nil {
		return nil, errors.WrapError(err)
	}
	http.Handle("/", router)
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
	return server, nil
}
