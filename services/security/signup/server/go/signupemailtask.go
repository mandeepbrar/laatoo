package main

import (
	"bytes"
	"encoding/json"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"text/template"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

/*

  SignupEmailTask:
    type: service
    description: Task processor for signup email tasks.
    configurations:
      SiteName:
        type: string
        description: Site name to use in email body
        variable: SiteName
        required: true
      SiteLink:
        type: string
        description: Link to use for verifying email
        required: true
        variable: SiteLink
      MailBody:
        type: string
        description: Mail body to use for emails
        required: true
        variable: MailBody
      Key:
        type: string
        required: true
        description: Key path to use for signing token
        variable: KeyPath
    request:
      params:
        Task:
          type: bytes
          description: Bytes containing string map for signup task

*/

type SignupEmailTask struct {
	core.Service
	SiteName string
	SiteLink string
	MailBody string
	key      string
}

func (signup *SignupEmailTask) Initialize(ctx core.ServerContext, conf config.Config) error {
	key, _ := conf.GetString(ctx, "Key")
	signup.key = key
	return nil
}
func (signup *SignupEmailTask) Invoke(ctx core.RequestContext) error {
	ctx.Dump()
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

	token, err := signup.createToken(ctx, mailId)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	mail, err := signup.createEmail(ctx, name, mailId, token)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	recipients := map[string]string{mailId: name}
	log.Error(ctx, "Communication created", "mail", mail, "recipients", recipients)
	mailPack := &components.Communication{Recipients: recipients, Message: []byte(mail), Subject: "Verify your Laatoo account"}

	log.Error(ctx, "Communication created", "mailPack", mailPack)

	err = ctx.SendCommunication(mailPack)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

/*\

 */
func (signup *SignupEmailTask) createEmail(ctx core.RequestContext, name, mailId, token string) (string, error) {
	email, err := template.New("Mail").Delims("<<", ">>").Parse(signup.MailBody)
	if err != nil {
		return "", errors.WrapError(ctx, err)
	}
	var tpl bytes.Buffer

	data := struct {
		Name     string
		Mail     string
		SiteLink string
		Token    string
		SiteName string
	}{
		name, mailId, signup.SiteLink, token, signup.SiteName,
	}

	if err := email.Execute(&tpl, data); err != nil {
		return "", err
	}
	return tpl.String(), nil
}

func (signup *SignupEmailTask) createToken(ctx core.RequestContext, mail string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["email"] = mail
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(signup.key)
	if err != nil {
		return "", errors.WrapError(ctx, err)
	}
	return tokenString, nil
}
