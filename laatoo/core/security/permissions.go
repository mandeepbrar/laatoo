package security

import (
	"laatoo/core/objects"
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

const (
	CONST_ALL_PERMISSIONS = "AllPermissions"
)

func init() {
	objects.RegisterInvokableMethod(CONST_ALL_PERMISSIONS, GetAllPermissions)
}

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
