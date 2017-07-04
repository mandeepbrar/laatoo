package security

import (
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

func GetAllPermissions(ctx core.RequestContext) error {
	elem := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if elem != nil {
		sh, ok := elem.(server.SecurityHandler)
		if ok {
			perms := sh.AllPermissions(ctx)
			ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, perms, nil))
			return nil
		}
	}
	ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, []string{}, nil))
	return nil
}
