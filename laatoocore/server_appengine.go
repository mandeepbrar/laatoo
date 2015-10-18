// +build appengine

package laatoocore

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
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
	ctx := echo.NewContext(nil, nil, router)
	server := &Server{ServerType: serverType}
	server.InitServer(ctx, configName, router)
	http.Handle("/", router)
	if server.ServerType == CONF_SERVERTYPE_GOOGLEAPP {
		log.Logger.Error(nil, "core.appengine.warmup", "setting up router for warmup")
		var req *echo.Context
		var once sync.Once
		warmupFunc := func() {
			log.Logger.Error(req, "core.appengine.warmup", "starting server")
			server.Start(req)
		}
		router.Use(func(ctx *echo.Context) error {
			req = ctx
			once.Do(warmupFunc)
			return nil
		})
		router.Get("/_ah/warmup", func(ctx *echo.Context) error {
			req = ctx
			once.Do(warmupFunc)
			return nil
		})
	}
	return server, nil
}
