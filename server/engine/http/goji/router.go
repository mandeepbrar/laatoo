package goji

import (
	"context"
	"fmt"
	"laatoo/engine/http/net"
	"net/http"
	"strings"

	"goji.io"
	"goji.io/pat"
)

type GojiRouter struct {
	routerGrp *goji.Mux
	rootMux   *goji.Mux
	pattern   string
}

//Get a sub router
func (router *GojiRouter) Group(path string) net.Router {
	submux := goji.SubMux()
	fmt.Println("********Registering for group path", path)
	newpath := router.pattern + path
	newpath = strings.Replace(newpath, "//", "/", -1)
	fmt.Println("********Goji mux path", (newpath + "/*"))
	router.routerGrp.HandleC(pat.New(path+"/*"), submux)

	fmt.Println("********Goji mux pathregistering", (newpath + "/me"))
	mehandler := func(pathCtx context.Context, res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		fmt.Fprint(res, "hello")
	}
	submux.HandleFuncC(pat.Get("/me"), mehandler)

	return &GojiRouter{routerGrp: submux, rootMux: router.rootMux, pattern: newpath}
}

func (router *GojiRouter) Get(path string, handler net.HandlerFunc) {
	router.routerGrp.HandleFuncC(pat.Get(path), router.httpAdapater(handler))
}

func (router *GojiRouter) Options(path string, handler net.HandlerFunc) {
	router.routerGrp.HandleFuncC(pat.Options(path), router.httpAdapater(handler))
}

func (router *GojiRouter) Put(path string, handler net.HandlerFunc) {
	router.routerGrp.HandleFuncC(pat.Put(path), router.httpAdapater(handler))
}

func (router *GojiRouter) Post(path string, handler net.HandlerFunc) {
	router.routerGrp.HandleFuncC(pat.Post(path), router.httpAdapater(handler))
}

func (router *GojiRouter) Delete(path string, handler net.HandlerFunc) {
	router.routerGrp.HandleFuncC(pat.Delete(path), router.httpAdapater(handler))
}

func (router *GojiRouter) httpAdapater(handler net.HandlerFunc) goji.HandlerFunc {
	return func(pathCtx context.Context, res http.ResponseWriter, req *http.Request) {
		corectx := &GojiContext{baseCtx: pathCtx, req: req, res: res}
		handler(corectx)
	}
}

func (router *GojiRouter) Use(handler net.HandlerFunc) {
	router.routerGrp.UseC(func(handler goji.Handler) goji.Handler {
		return nil /*func(echoCtx echo.Context) error {
			corectx := &EchoContext{baseCtx: echoCtx}
			//defer corectx.CompleteRequest()
			return handler(corectx)
		}*/
	})
}

func (router *GojiRouter) UseMW(handler func(http.Handler) http.Handler) {
	router.routerGrp.Use(handler)
}
func (router *GojiRouter) UseMiddleware(handler http.HandlerFunc) {
	panic("not implemented")
}
