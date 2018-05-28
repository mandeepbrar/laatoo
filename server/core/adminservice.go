package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
)

type adminService struct {
	core.Service
}

func (svc *adminService) Describe(ctx core.ServerContext) {
	svc.SetDescription(ctx, "Server admin service")
	svc.AddStringParam(ctx, "request")
	svc.AddParam(ctx, "requestparam", config.OBJECTTYPE_STRING, false, false)
	svc.SetRequestType(ctx, config.OBJECTTYPE_BYTES, false, false)
}

func (svc *adminService) Invoke(ctx core.RequestContext) error {
	req, _ := ctx.GetStringParam("request")
	switch req {
	case "shutdown":
		log.Error(ctx, "Shutting down server")
	case "list":
		reqparam, _ := ctx.GetStringParam("requestparam")
		log.Error(ctx, "List modules request", "param", reqparam)
	}
	return nil
}
