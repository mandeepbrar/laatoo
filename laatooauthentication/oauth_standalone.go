// +build !appengine

package laatooauthentication

import (
	"github.com/labstack/echo"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func GetOAuthContext(ctx *echo.Context) context.Context {
	return oauth2.NoContext
}
