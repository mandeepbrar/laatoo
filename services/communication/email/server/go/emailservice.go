package main

import (
	"encoding/json"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
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
	sender     string
	queueComm  bool
	emailqueue string
}

func (svc *EmailService) Initialize(ctx core.ServerContext, conf config.Config) error {
	var ok bool
	svc.mailServer, _ = svc.GetStringConfiguration(ctx, "server")
	svc.mailUser, _ = svc.GetStringConfiguration(ctx, "user")
	svc.mailPass, _ = svc.GetStringConfiguration(ctx, "password")
	svc.sender, _ = svc.GetStringConfiguration(ctx, "mailsender")
	svc.queueComm, _ = svc.GetBoolConfiguration(ctx, "queuetasks")
	svc.emailqueue, ok = svc.GetStringConfiguration(ctx, "emailqueue")
	if svc.queueComm && !ok {
		return errors.MissingConf(ctx, "emailqueue")
	}
	return nil
}

func (svc *EmailService) Invoke(ctx core.RequestContext) error {

	taskI, _ := ctx.GetParamValue("Task")
	task, ok := taskI.(*components.Task)
	if ok {
		communication := make(map[interface{}]interface{})
		err := json.Unmarshal(task.Data, &communication)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		return svc.sendCommunication(ctx, communication)
	} else {
		return errors.BadRequest(ctx, "Error", "Task not present")
	}
	return nil
}

func (svc *EmailService) SendCommunication(ctx core.RequestContext, communication map[interface{}]interface{}) (err error) {
	if svc.queueComm {
		err = ctx.PushTask(svc.emailqueue, communication)
	} else {
		err = svc.sendCommunication(ctx, communication)
	}
	if err != nil {
		err = errors.WrapError(ctx, err)
	}
	return err
}

func (svc *EmailService) sendCommunication(ctx ctx.Context, communication map[interface{}]interface{}) error {
	for recipients, msg := range communication {
		err := svc.sendMessage(ctx, recipients.([]string), svc.sender, msg)
		if err != nil {
			log.Error(ctx, "Could not send email to recepients", "recepients", recipients)
		}
	}
	return nil
}

func (svc *EmailService) sendMessage(ctx ctx.Context, recipients []string, sender string, msg interface{}) error {
	auth := smtp.PlainAuth("", svc.mailUser, svc.mailPass, svc.mailServer)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(svc.mailServer+":25", auth, sender, recipients, msg.([]byte))
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}
