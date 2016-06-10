package middleware

import (
	"laatoo/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
)

const (
	SVC_CHECKPERMISSION  = "CheckPermission"
	SVC_PERMISSION_PARAM = "permission"
)

func init() {
	objects.RegisterObject(SVC_CHECKPERMISSION, createCheckPermissionsService, nil)
}

func createCheckPermissionsService(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	return &checkPermissionService{}, nil
}

type checkPermissionService struct {
	perm    string
	svcperm bool
}

//The services start serving when this method is called
func (svc *checkPermissionService) Initialize(ctx core.ServerContext, conf config.Config) error {
	svc.perm, svc.svcperm = conf.GetString(SVC_PERMISSION_PARAM)
	return nil
}

func (svc *checkPermissionService) Start(ctx core.ServerContext) error {
	return nil
}

//The services start serving when this method is called
func (svc *checkPermissionService) Invoke(ctx core.RequestContext) error {
	var perm string
	var ok bool
	if svc.svcperm {
		perm = svc.perm
	} else {
		perm, ok = ctx.GetString(SVC_PERMISSION_PARAM)
		if !ok {
			ctx.SetResponse(core.StatusUnauthorizedResponse)
			return nil
		}
	}
	hasperm := ctx.HasPermission(perm)
	log.Logger.Trace(ctx, "Checked permission", "perm", perm, "result", hasperm)
	if !hasperm {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
	}
	return nil
}