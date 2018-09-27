package main

import (
	"laatoo/sdk/server/core"
)

func (svc *UI) Invoke(ctx core.RequestContext) error {
	mod, ok := ctx.GetStringParam("module")
	if ok {
		cont, ok := svc.uiFiles[mod]
		if ok {
			ctx.SetResponse(core.NewServiceResponseWithInfo(core.StatusServeBytes, &cont, nil))
		} else {
			ctx.SetResponse(core.StatusNotFoundResponse)
		}
	}
	return nil
}