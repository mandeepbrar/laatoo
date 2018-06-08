package security

import (
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

type AllPermissions struct {
	core.Service
}

func (ap *AllPermissions) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	elem := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if elem != nil {
		sh, ok := elem.(server.SecurityHandler)
		if ok {
			perms := sh.AllPermissions(ctx)
			return core.SuccessResponse(perms), nil
		}
	}
	return core.SuccessResponse([]string{}), nil
}
