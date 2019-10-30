package main

import (
	"encoding/json"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"

	mailsender "github.com/go-mail/mail"
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
	MailPort   string
	mailPass   string
	sender     string
	queueComm  bool
	emailqueue string
}

func (svc *EmailService) Initialize(ctx core.ServerContext, conf config.Config) error {
	var ok bool
	svc.mailServer, _ = svc.GetStringConfiguration(ctx, "mailserver")
	svc.mailPass, _ = svc.GetStringConfiguration(ctx, "mailpass")
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
		communication := &components.Communication{}
		err := json.Unmarshal(task.Data, communication)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		return svc.sendCommunication(ctx, communication)
	} else {
		return errors.BadRequest(ctx, "Error", "Task not present")
	}
	return nil
}

func (svc *EmailService) SendCommunication(ctx core.RequestContext, communication *components.Communication) (err error) {
	if svc.queueComm {
		err = ctx.PushTask(svc.emailqueue, communication)
	} else {
		err = svc.sendCommunication(ctx, communication)
	}
	if err != nil {
		err = errors.WrapError(ctx, err)
	}
	log.Debug(ctx, "Sent email to recipients", "recipients", communication.Recipients)

	return err
}

func (svc *EmailService) sendCommunication(ctx ctx.Context, communication *components.Communication) error {
	msg := string(communication.Message)
	for mail, name := range communication.Recipients {
		err := svc.sendMessage(ctx, mail, name, svc.sender, communication.Subject, communication.Mime, msg, communication.Attachments)
		if err != nil {
			log.Error(ctx, "Could not send email to recepients", "recipients", communication.Recipients)
		}
	}
	return nil
}

func (svc *EmailService) sendMessage(ctx ctx.Context, mail, name, sender, subject, mime, msg string, attachments []string) error {

	m := mailsender.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", mail)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", string(msg))
	//	m.Attach("/home/Alex/lolcat.jpg")

	d := mailsender.NewDialer(svc.mailServer, 587, sender, svc.mailPass)
	d.StartTLSPolicy = mailsender.MandatoryStartTLS

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

/*

func (svc *EmailService) sendMessage(ctx ctx.Context, recipients []string, sender string, msg interface{}) error {
	auth := smtp.PlainAuth("", sender, svc.mailPass, svc.mailServer)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         svc.mailServer,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	c, err := smtp.Dial(svc.mailServer + svc.MailPort)
	if nil != err {
		return errors.WrapError(ctx, err)
	}
	defer c.Close()
	if err = c.StartTLS(tlsconfig); err != nil {
		log.Error(ctx, "tls error "+err.Error())
		return errors.WrapError(ctx, err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		return errors.WrapError(ctx, err)
	}

	for _, recipient := range recipients {

		// To && From
		if err = c.Mail(sender); err != nil {
			return errors.WrapError(ctx, err)
		}

		if err = c.Rcpt(recipient); err != nil {
			return errors.WrapError(ctx, err)
		}

		// Data
		w, err := c.Data()
		if err != nil {
			return errors.WrapError(ctx, err)
		}

		_, err = w.Write(msg.([]byte))
		if err != nil {
			return errors.WrapError(ctx, err)
		}

		err = w.Close()
		if err != nil {
			return errors.WrapError(ctx, err)
		}

		c.Quit()
	}
	return nil
}
*/
