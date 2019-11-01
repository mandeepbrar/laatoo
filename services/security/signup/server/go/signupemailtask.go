package main

import (
	"encoding/json"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

type SignupEmailTask struct {
	core.Service
	SiteName string
	SiteLink string
	MailBody string
	verifier *EmailVerifier
}

func (signup *SignupEmailTask) Initialize(ctx core.ServerContext, conf config.Config) error {
	key, _ := conf.GetString(ctx, "Key")
	signup.verifier = &EmailVerifier{SiteName: signup.SiteName, SiteLink: signup.SiteLink, MailBody: signup.MailBody, key: key}
	return nil
}

func (signup *SignupEmailTask) Invoke(ctx core.RequestContext) error {
	tsk, ok := ctx.GetParamValue("Task")
	if !ok {
		return errors.MissingArg(ctx, "Task")
	}
	task, ok := tsk.([]byte)
	if !ok {
		return errors.BadArg(ctx, "task", "Err", "Argument not bytes")
	}
	mailArgs := make(map[string]string)
	err := json.Unmarshal(task, &mailArgs)
	if err != nil {
		return errors.BadArg(ctx, "task", "Err", "Could not unmarshal task data")
	}
	log.Error(ctx, "Mail to be sent", "args", mailArgs)
	mailId, ok := mailArgs["email"]
	if !ok {
		return errors.BadArg(ctx, "task", "Err", "missing email in task")
	}
	name, ok := mailArgs["name"]

	if mailId != "" {
		err := signup.verifier.sendCommunication(ctx, name, mailId)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	} else {
		return errors.BadRequest(ctx, "Missing email in request map", "email")
	}

	return nil
}
