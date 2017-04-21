package middleware

import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/core"
	//"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	CHECK_USER_OWN_ACCOUNT = "OwnUserAccountEnforce"
)

func init() {
	objects.RegisterInvokableMethod(CHECK_USER_OWN_ACCOUNT, OwnUserAccountEnforce)
}

func OwnUserAccountEnforce(ctx core.RequestContext) error {
	id, ok := ctx.GetString("id")
	if !ok {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	usr := ctx.GetUser()
	if usr == nil {
		log.Logger.Trace(ctx, "Entity not accessible by anonymous user", "user", id)
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	if id != usr.GetId() {
		log.Logger.Trace(ctx, "Entity accessible only by owner", "entity", id, "user", usr.GetId())
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	log.Logger.Trace(ctx, "User allowed", "entity", id, "user", usr.GetId())
	return nil
}
