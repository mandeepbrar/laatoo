// +build !appengine

package laatooauthentication

import (
	glctx "golang.org/x/net/context"
	"golang.org/x/oauth2"
	"laatoosdk/core"
)

func GetOAuthContext(ctx core.Context) glctx.Context {
	return oauth2.NoContext
}
