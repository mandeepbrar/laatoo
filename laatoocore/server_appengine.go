// +build appengine

package laatoocore

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"laatoosdk/log"
	"net/http"
	"sync"
)

var (
	APPENGINE_CONTEXT context.Context
)

//Create a new server
func NewServer(configName string, serverType string) (*Server, error) {
	//initialize router
	router := echo.New()
	// Middleware
	router.Use(mw.Logger())
	router.Use(mw.Recover())
	server := &Server{ServerType: serverType}
	server.InitServer(configName, router)
	http.Handle("/", router)
	if server.ServerType == CONF_SERVERTYPE_GOOGLEAPP {
		log.Logger.Error("setting up router for warmup")
		var req *echo.Context
		var once sync.Once
		warmupFunc := func() {
			APPENGINE_CONTEXT = appengine.NewContext(req.Request())
			log.Logger.Error("context", APPENGINE_CONTEXT)
			server.Start()
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
