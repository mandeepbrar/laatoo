package gin

import (
	"laatoo/sdk/log"
	"laatoo/server/engine/http/net"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinWebFramework struct {
	rootRouter *gin.Engine
	Name       string
}

func (wf *GinWebFramework) Initialize() error {
	//create all service factories in the application
	//initialize router
	router := gin.New()
	// Middleware
	//loggerconfig := mw.LoggerConfig{Output: log.Logger, Format: "{\"time\":\"${time_rfc3339}\", \"remote_ip\":\"${remote_ip}\", \"method\":\"${method}\", \"uri\":\"${uri}\", \"status\":\"${status}\", \"took\":\"${response_time}\", \"sent\":\"${response_size} bytes\"}\n"}
	router.Use(gin.LoggerWithWriter(log.Logger))
	router.Use(gin.Recovery())
	wf.rootRouter = router
	return nil
}

func (wf *GinWebFramework) GetParentRouter(path string) net.Router {
	return &GinRouter{routerGrp: wf.rootRouter.Group(path)}
}

func (wf *GinWebFramework) GetRootHandler() http.Handler {
	return wf.rootRouter
}

func (wf *GinWebFramework) StartServer(address string) error {
	return wf.rootRouter.Run(address)
}

func (wf *GinWebFramework) StartSSLServer(address string, certpath string, keypath string) error {
	return wf.rootRouter.RunTLS(address, certpath, keypath)
}
