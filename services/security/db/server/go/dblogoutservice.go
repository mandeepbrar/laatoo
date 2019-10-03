package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	common "securitycommon"
)

type LogoutService struct {
	core.Service
	authHeader string
}

func (ls *LogoutService) Initialize(ctx core.ServerContext, conf config.Config) error {
	sechandler := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sechandler == nil {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	authHeader := sechandler.GetProperty(config.AUTHHEADER)
	if authHeader == nil {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ls.authHeader = authHeader.(string)

	return nil
}

func (ls *LogoutService) Invoke(ctx core.RequestContext) error {
	info := map[string]interface{}{ls.authHeader: "<delete>"}
	ctx.SetResponse(core.SuccessResponseWithInfo(nil, info))
	return nil
}
