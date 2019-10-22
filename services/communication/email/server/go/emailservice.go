package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/errors"
	"net/smtp"
)

const (
	CONF_EMAIL_SVC = "email_service"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_EMAIL_SVC, Object: EmailService{}}}
}

type EmailService struct {
	core.Service
	mailServer string
	mailUser   string
	mailPass   string
}

func (svc *EmailService) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

func (svc *EmailService) SendMessage(ctx ctx.Context, recipients []string, sender string, msg interface{}) error {
	auth := smtp.PlainAuth("", svc.mailUser, svc.mailPass, svc.mailServer)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(svc.mailServer+":25", auth, sender, recipients, msg.([]byte))
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (svc *EmailService) Start(ctx core.ServerContext) error {
	svc.mailServer, _ = svc.GetStringConfiguration(ctx, "server")
	svc.mailUser, _ = svc.GetStringConfiguration(ctx, "user")
	svc.mailPass, _ = svc.GetStringConfiguration(ctx, "password")
	return nil
}
