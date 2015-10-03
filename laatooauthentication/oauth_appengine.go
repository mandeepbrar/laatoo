// +build appengine

package laatooauthentication

import (
	"github.com/labstack/echo"
	"golang.org/x/net/context"
	appctx "laatoosdk/context"
)

func GetOAuthContext(ctx *echo.Context) context.Context {
	return appctx.GetAppengineContext(ctx)
}
