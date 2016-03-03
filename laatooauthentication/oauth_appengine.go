// +build appengine

package laatooauthentication

import (
	glctx "golang.org/x/net/context"
	"laatoosdk/core"
)

func GetOAuthContext(ctx core.Context) glctx.Context {
	return ctx.GetAppengineContext()
}
