// +build appengine

package context

import (
	"github.com/labstack/echo"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func GetAppengineContext(ctx interface{}) context.Context {
	if ctx == nil {
		return nil
	}
	echoContext, ok := ctx.(*echo.Context)
	if !ok {
		appengineCtx, ok := ctx.(context.Context)
		if ok {
			return appengineCtx
		}
		return nil
	}
	return appengine.NewContext(echoContext.Request())

}
