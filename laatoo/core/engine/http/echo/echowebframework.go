package echo

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	mw "github.com/labstack/echo/middleware"
	"laatoo/core/engine/http/net"
	"laatoo/sdk/log"
	"net/http"
)

type EchoWebFramework struct {
	rootRouter *echo.Echo
}

func (wf *EchoWebFramework) Initialize() error {
	//create all service factories in the application
	//initialize router
	router := echo.New()
	// Middleware
	loggerconfig := mw.LoggerConfig{Output: log.Logger, Format: "{\"time\":\"${time_rfc3339}\", \"remote_ip\":\"${remote_ip}\", \"method\":\"${method}\", \"uri\":\"${uri}\", \"status\":\"${status}\", \"took\":\"${response_time}\", \"sent\":\"${response_size} bytes\"}\n"}
	router.Use(mw.LoggerFromConfig(loggerconfig))
	router.Use(mw.Recover())
	router.Use(mw.Gzip())
	wf.rootRouter = router
	return nil
}

func (wf *EchoWebFramework) GetParentRouter() net.Router {
	return &EchoRouter{routerGrp: wf.rootRouter.Group("")}
}

func (wf *EchoWebFramework) GetRootHandler() http.Handler {
	s := standard.New("")
	s.SetHandler(wf.rootRouter)
	return s
}

func (wf *EchoWebFramework) StartServer(address string) error {
	wf.rootRouter.Run(standard.New(address))
	return nil
}

func (wf *EchoWebFramework) StartSSLServer(address string, certpath string, keypath string) error {
	wf.rootRouter.Run(standard.NewFromTLS(address, certpath, keypath))
	return nil
}
