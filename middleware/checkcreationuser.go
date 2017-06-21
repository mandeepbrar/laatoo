package main

import (
	"laatoo/sdk/core"
	//"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	CHECK_CREATION_USER = "CheckCreationUser"
)

func CheckCreationUser(ctx core.RequestContext) error {
	id, ok := ctx.GetString("id")
	if !ok {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	usr := ctx.GetUser()
	if usr == nil {
		log.Trace(ctx, "Entity not accessible by anonymous user", "user", id)
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	if id != usr.GetId() {
		log.Trace(ctx, "Entity accessible only by owner", "entity", id, "user", usr.GetId())
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	log.Trace(ctx, "User allowed", "entity", id, "user", usr.GetId())
	return nil
}
