package security

import (
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

func GetAllPermissions(ctx core.RequestContext, req core.ServiceRequest) (*core.ServiceResponse, error) {
	elem := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if elem != nil {
		sh, ok := elem.(server.SecurityHandler)
		if ok {
			perms := sh.AllPermissions(ctx)
			return core.NewServiceResponse(core.StatusSuccess, perms, nil), nil
		}
	}
	return core.NewServiceResponse(core.StatusSuccess, []string{}, nil), nil
}
