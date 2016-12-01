package echo

import (
	"laatoo/framework/core/engine/http/net"
	"laatoo/sdk/log"
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
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
	router.Use(mw.LoggerWithConfig(loggerconfig))
	router.Use(mw.Recover())
	router.Use(mw.Gzip())
	wf.rootRouter = router
	return nil
}

func (wf *EchoWebFramework) GetParentRouter(path string) net.Router {
	return &EchoRouter{routerGrp: wf.rootRouter.Group(path)}
}

func (wf *EchoWebFramework) GetRootHandler() http.Handler {
	return wf.rootRouter
}

func (wf *EchoWebFramework) StartServer(address string) error {
	wf.rootRouter.Start(address)
	return nil
}

func (wf *EchoWebFramework) StartSSLServer(address string, certpath string, keypath string) error {
	wf.rootRouter.StartTLS(address, certpath, keypath)
	return nil
}
