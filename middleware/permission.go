package main

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"
)

const (
	SVC_CHECKPERMISSION  = "CheckPermission"
	SVC_PERMISSION_PARAM = "permission"
)

type checkPermissionService struct {
	core.Service
	perm string
}

//The services start serving when this method is called
func (svc *checkPermissionService) Initialize(ctx core.ServerContext) error {
	svc.SetDescription("Check permission middleware service. Checks if the permission required by a service has been assigned to a user")
	svc.AddStringConfigurations([]string{SVC_PERMISSION_PARAM}, nil)
	return nil
}

func (svc *checkPermissionService) Start(ctx core.ServerContext) error {
	permission, _ := svc.GetStringConfiguration(SVC_PERMISSION_PARAM)
	svc.perm = permission
	return nil
}

//The services start serving when this method is called
func (svc *checkPermissionService) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	//	log.Trace(ctx, "Checking permissions")
	/*var perm string
	var ok bool
	if svc.svcperm {
		perm = svc.perm
	} else {
		perm, ok = ctx.GetString(SVC_PERMISSION_PARAM)
		if !ok {
			log.Trace(ctx, "Unauthorized response for ", "perm", perm)
			return core.StatusUnauthorizedResponse, nil
		}
	}*/
	hasperm := ctx.HasPermission(svc.perm)
	log.Trace(ctx, "Checked permission", "perm", svc.perm, "result", hasperm)
	if !hasperm {
		return core.StatusUnauthorizedResponse, nil
	}
	return nil, nil
}
