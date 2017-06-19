package main

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/services/security/common"
)

const (
	CONF_SECURITYSERVICE_TOKENVALIDATION = "TOKEN_VALIDATE"
)

func Manifest() []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_SECURITYSERVICE_TOKENVALIDATION, Object: TokenValidationService{}}}
}

type TokenValidationService struct {
	sechandler server.SecurityHandler
	authHeader string
}

func (ls *TokenValidationService) Initialize(ctx core.ServerContext, conf config.Config) error {
	sechandler := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sechandler == nil {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ls.sechandler = sechandler.(server.SecurityHandler)
	authHeader, ok := sechandler.GetString(config.AUTHHEADER)
	if !ok {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	if !ok {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ls.authHeader = authHeader

	log.Logger.Debug(ctx, "Token validation service initialized")
	return nil
}

//Expects Local user to be provided inside the request
func (ls *TokenValidationService) Invoke(ctx core.RequestContext) error {
	token, err := ls.sechandler.AuthenticateRequest(ctx, true)
	usr := ctx.GetUser()
	log.Logger.Error(ctx, "checked token", "err", err, "usr", usr)
	if err == nil && usr != nil {
		log.Logger.Error(ctx, "checked token", "usr", usr.GetId())
		if usr.GetId() != "Anonymous" {
			resp := core.NewServiceResponse(core.StatusSuccess, usr, map[string]interface{}{ls.authHeader: token})
			ctx.SetResponse(resp)
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
