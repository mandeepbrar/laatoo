package laatooauthentication

import (
	"laatoocore"
	"laatoosdk/core"
	"laatoosdk/errors"
	"laatoosdk/log"
)

const (
	CHECK_USER_OWN_ACCOUNT = "OwnAccountEnforce"
	USER_AUTH_CONTEXT      = "laatooauth.userauthorization"
)

func init() {
	laatoocore.RegisterInvokableMethod(CHECK_USER_OWN_ACCOUNT, OwnAccountEnforce)
}

func OwnAccountEnforce(ctx core.Context, conf map[string]interface{}) error {
	id := ctx.ParamByIndex(0)
	if id == "" {
		return errors.ThrowError(ctx, AUTH_ERROR_NOT_ALLOWED)
	}
	usr := ctx.GetUser()
	if usr == nil {
		log.Logger.Trace(ctx, USER_AUTH_CONTEXT, "Entity not accessible by anonymous user", "user", id)
		return errors.ThrowError(ctx, AUTH_ERROR_NOT_ALLOWED)
	}
	if id != usr.GetId() {
		log.Logger.Trace(ctx, USER_AUTH_CONTEXT, "Entity accessible only by owner", "entity", id, "user", usr.GetId())
		return errors.ThrowError(ctx, AUTH_ERROR_NOT_ALLOWED)
	}
	log.Logger.Trace(ctx, USER_AUTH_CONTEXT, "User allowed", "entity", id, "user", usr.GetId())
	return nil
}
