package gin

import (
	"github.com/gin-gonic/gin"
	"laatoo/core/engine/http/net"
	"net/http"
)

type GinRouter struct {
	routerGrp *gin.RouterGroup
}

//Get a sub router
func (router *GinRouter) Group(path string) net.Router {
	return &GinRouter{routerGrp: router.routerGrp.Group(path)}
}

func (router *GinRouter) Get(path string, handler net.HandlerFunc) {
	router.routerGrp.GET(path, router.httpAdapater(handler))
}

func (router *GinRouter) Options(path string, handler net.HandlerFunc) {
	router.routerGrp.OPTIONS(path, router.httpAdapater(handler))
}

func (router *GinRouter) Put(path string, handler net.HandlerFunc) {
	router.routerGrp.PUT(path, router.httpAdapater(handler))
}

func (router *GinRouter) Post(path string, handler net.HandlerFunc) {
	router.routerGrp.POST(path, router.httpAdapater(handler))
}

func (router *GinRouter) Delete(path string, handler net.HandlerFunc) {
	router.routerGrp.DELETE(path, router.httpAdapater(handler))
}

func (router *GinRouter) httpAdapater(handler net.HandlerFunc) gin.HandlerFunc {
	return func(pathCtx *gin.Context) {
		corectx := &GinContext{baseCtx: pathCtx}
		handler(corectx)
	}
}

func (router *GinRouter) Use(handler net.HandlerFunc) {
	router.routerGrp.Use(func(ginctx *gin.Context) {
		corectx := &GinContext{baseCtx: ginctx}
		//defer corectx.CompleteRequest()
		handler(corectx)
		ginctx.Next()
	})
}

func (router *GinRouter) UseMW(handler func(http.Handler) http.Handler) {
	panic("not implemented")
}
func (router *GinRouter) UseMiddleware(handler http.HandlerFunc) {
	router.routerGrp.Use(gin.WrapF(handler))
}
