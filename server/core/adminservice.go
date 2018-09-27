package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"
)

type adminService struct {
	core.Service
}

func (svc *adminService) Describe(ctx core.ServerContext) error {
	svc.SetDescription(ctx, "Server admin service")
	svc.AddStringParam(ctx, "request")
	return svc.AddParam(ctx, "requestparam", config.OBJECTTYPE_STRING, false, false, false)
	//svc.SetRequestType(ctx, config.OBJECTTYPE_BYTES, false, false)
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