package echo

import (
	"laatoo/server/engine/http/net"
	"net/http"

	"github.com/labstack/echo"
)

type EchoRouter struct {
	routerGrp *echo.Group
}

//Get a sub router
func (router *EchoRouter) Group(path string) net.Router {
	return &EchoRouter{routerGrp: router.routerGrp.Group(path)}
}

func (router *EchoRouter) Get(path string, handler net.HandlerFunc) {
	router.routerGrp.GET(path, router.httpAdapater(handler))
}
func (router *EchoRouter) Options(path string, handler net.HandlerFunc) {
	router.routerGrp.OPTIONS(path, router.httpAdapater(handler))
}
func (router *EchoRouter) Put(path string, handler net.HandlerFunc) {
	router.routerGrp.PUT(path, router.httpAdapater(handler))
}

func (router *EchoRouter) Post(path string, handler net.HandlerFunc) {
	router.routerGrp.POST(path, router.httpAdapater(handler))
}

func (router *EchoRouter) Delete(path string, handler net.HandlerFunc) {
	router.routerGrp.DELETE(path, router.httpAdapater(handler))
}

func (router *EchoRouter) httpAdapater(handler net.HandlerFunc) echo.HandlerFunc {
	return func(pathCtx echo.Context) error {
		corectx := &EchoContext{baseCtx: pathCtx}
		return handler(corectx)
	}
}

func (router *EchoRouter) Use(handler net.HandlerFunc) {
	router.routerGrp.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCtx echo.Context) error {
			corectx := &EchoContext{baseCtx: echoCtx}
			//defer corectx.CompleteRequest()
			err := handler(corectx)
			if err != nil {
				return err
			}
			return next(echoCtx)
		}
	})
}

func (router *EchoRouter) UseMW(handler func(http.Handler) http.Handler) {
	router.routerGrp.Use(echo.WrapMiddleware(handler))
}
func (router *EchoRouter) UseMiddleware(handler http.HandlerFunc) {
	panic("not implemented")
}
