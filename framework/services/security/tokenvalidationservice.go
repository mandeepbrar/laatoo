package security

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
)

type TokenValidationService struct {
	sechandler server.SecurityHandler
}

func (ls *TokenValidationService) Initialize(ctx core.ServerContext, conf config.Config) error {
	sechandler := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sechandler == nil {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ls.sechandler = sechandler.(server.SecurityHandler)
	log.Logger.Debug(ctx, "Token validation service initialized")
	return nil
}

//Expects Local user to be provided inside the request
func (ls *TokenValidationService) Invoke(ctx core.RequestContext) error {
	err := ls.sechandler.AuthenticateRequest(ctx)
	usr := ctx.GetUser()
	log.Logger.Error(ctx, "checked token", "err", err)
	if err == nil && usr != nil {
		log.Logger.Error(ctx, "checked token", "usr", usr.GetId())
		if usr.GetId() != "Anonymous" {
			ctx.SetResponse(core.StatusSuccessResponse)
			return nil
		}
	}
	log.Logger.Error(ctx, "checked token - sending unauthorized response")
	ctx.SetResponse(core.StatusUnauthorizedResponse)
	return nil
}

func (ls *TokenValidationService) Start(ctx core.ServerContext) error {
	return nil
}
