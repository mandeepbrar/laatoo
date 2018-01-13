package main

import (
	"laatoo/sdk/core"
	"laatoo/sdk/log"
)

type ObjectResolver struct {
	core.Service
	objtype string
}

func (svc *ObjectResolver) Start(ctx core.ServerContext) error {
	svc.objtype, _ = svc.GetStringConfiguration(ctx, "objecttype")
	svc.SetRequestType(ctx, svc.objtype, false, false)
	return nil
}

func (svc *ObjectResolver) Invoke(ctx core.RequestContext) error {
	switch svc.objtype {
	case "Server":
		return svc.resolveServer(ctx)
	case "Environment":
		return svc.resolveEnvironment(ctx)
	case "Application":
		return svc.resolveApplication(ctx)
	}
	return nil
}

func (svc *ObjectResolver) resolveServer(ctx core.RequestContext) error {
	ctx = ctx.SubContext("Resolve server")
	svr := ctx.GetBody().(*Server)
	log.Error(ctx, "Received server", "svr", svr)
	return nil
}

func (svc *ObjectResolver) resolveEnvironment(ctx core.RequestContext) error {
	ctx = ctx.SubContext("Resolve environment")
	env := ctx.GetBody().(*Environment)
	log.Error(ctx, "Received environment", "env", env)
	return nil
}

func (svc *ObjectResolver) resolveApplication(ctx core.RequestContext) error {
	ctx = ctx.SubContext("Resolve application")
	app := ctx.GetBody().(*Application)
	log.Error(ctx, "Received application", "app", app)
	return nil
}
