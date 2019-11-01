package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

type VerifyEmailService struct {
	core.Service
	verifier           *EmailVerifier
	VerifyWithWorkflow bool
}

func (svc *VerifyEmailService) Initialize(ctx core.ServerContext, conf config.Config) error {
	key, _ := conf.GetString(ctx, "Key")
	svc.verifier = &EmailVerifier{key: key}
	return nil
}

func (svc *VerifyEmailService) Invoke(ctx core.RequestContext) error {
	token, ok := ctx.GetStringParam("token")
	if !ok {
		return errors.MissingArg(ctx, "token")
	}
	email, err := svc.verifier.verifyToken(ctx, token)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Debug(ctx, "email token verified", "email", email)
	return nil
}
