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
	/***try ***/
	svc.AddParamWithType(ctx, "object", svc.objtype)
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
	obj, _ := ctx.GetParamValue("object")
	svr := obj.(*Server)
	log.Error(ctx, "Received server", "svr", svr)
	return nil
}

func (svc *ObjectResolver) resolveEnvironment(ctx core.RequestContext) error {
	ctx = ctx.SubContext("Resolve environment")
	obj, _ := ctx.GetParamValue("object")
	env := obj.(*Environment)
	log.Error(ctx, "Received environment", "env", env)
	return nil
}

func (svc *ObjectResolver) resolveApplication(ctx core.RequestContext) error {
	ctx = ctx.SubContext("Resolve application")
	obj, _ := ctx.GetParamValue("object")
	app := obj.(*Application)
	log.Error(ctx, "Received application", "app", app)
	return nil
}
