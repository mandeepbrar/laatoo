package goji

import (
	"goji.io"
	"laatoo/framework/core/engine/http/net"
	//	"laatoo/sdk/log"
	"net/http"
)

type GojiWebFramework struct {
	rootRouter *goji.Mux
}

func (wf *GojiWebFramework) Initialize() error {
	//create all service factories in the application
	//initialize router
	mux := goji.NewMux()
	wf.rootRouter = mux
	return nil
}

func (wf *GojiWebFramework) GetParentRouter(path string) net.Router {
	return &GojiRouter{routerGrp: wf.rootRouter, rootMux: wf.rootRouter, pattern: path}
}

func (wf *GojiWebFramework) GetRootHandler() http.Handler {
	return wf.rootRouter
}

func (wf *GojiWebFramework) StartServer(address string) error {
	return http.ListenAndServe(address, wf.rootRouter)
}

func (wf *GojiWebFramework) StartSSLServer(address string, certpath string, keypath string) error {
	return http.ListenAndServeTLS(address, certpath, keypath, wf.rootRouter)
}
