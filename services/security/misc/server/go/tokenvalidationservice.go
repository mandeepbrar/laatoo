package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"securitycommon"
)

const (
	CONF_SECURITYSERVICE_TOKENVALIDATION = "TOKEN_VALIDATE"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_SECURITYSERVICE_TOKENVALIDATION, Object: TokenValidationService{}}}
}

type TokenValidationService struct {
	core.Service
	sechandler elements.SecurityHandler
	authHeader string
}

func (ls *TokenValidationService) Initialize(ctx core.ServerContext, conf config.Config) error {
	sechandler := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sechandler == nil {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ls.sechandler = sechandler.(elements.SecurityHandler)
	authHeader := sechandler.GetProperty(config.AUTHHEADER)
	if authHeader == nil {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ls.authHeader = authHeader.(string)

	log.Debug(ctx, "Token validation service initialized")
	return nil
}

//Expects Local user to be provided inside the request
func (ls *TokenValidationService) Invoke(ctx core.RequestContext) error {
	token, err := ls.sechandler.AuthenticateRequest(ctx, true)
	usr := ctx.GetUser()
	log.Error(ctx, "checked token", "err", err, "usr", usr)
	if err == nil && usr != nil {
		log.Error(ctx, "checked token", "usr", usr.GetId())
		if usr.GetId() != "Anonymous" {
			ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, usr, map[string]interface{}{ls.authHeader: token}))
			return nil
		}
	}
	log.Error(ctx, "checked token - sending unauthorized response")
	ctx.SetResponse(core.StatusUnauthorizedResponse)
	return nil
}
