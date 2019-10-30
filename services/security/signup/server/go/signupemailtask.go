package main

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"text/template"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type SignupEmailTask struct {
	core.Service
	SiteName string
	SiteLink string
	MailBody string
	KeyPath  string
	pvtKey   *rsa.PrivateKey
}

func (signup *SignupEmailTask) Initialize(ctx core.ServerContext, conf config.Config) error {
	log.Error(ctx, "loading key from path", "path", signup, "conf", conf)
	ctx.Dump()
	pvtKey, err := utils.LoadPrivateKey(signup.KeyPath)
	if err != nil {
		return errors.BadConf(ctx, "KeyPath", "err", err)
	}
	signup.pvtKey = pvtKey
	return nil
}
func (signup *SignupEmailTask) Start(ctx core.ServerContext) error {
	log.Error(ctx, "loading key from path", "path", signup)
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
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	tokenString, err := token.SignedString(signup.pvtKey)
	if err != nil {
		return "", errors.WrapError(ctx, err)
	}
	return tokenString, nil
}
