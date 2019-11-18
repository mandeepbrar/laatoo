package signup

import (
	"bytes"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"net/url"
	"text/template"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type EmailVerifier struct {
	SiteName string
	SiteLink string
	MailBody string
	key      []byte
}

func (verifier *EmailVerifier) sendCommunication(ctx core.RequestContext, name, mailId string) error {

	token, err := verifier.createToken(ctx, mailId)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	mail, err := verifier.createEmail(ctx, name, mailId, token)
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

func (verifier *EmailVerifier) createEmail(ctx core.RequestContext, name, mailId, token string) (string, error) {
	email, err := template.New("Mail").Delims("<<", ">>").Parse(verifier.MailBody)
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
		name, mailId, verifier.SiteLink, url.PathEscape(token), verifier.SiteName,
	}
	if err := email.Execute(&tpl, data); err != nil {
		return "", err
	}
	return tpl.String(), nil
}

func (verifier *EmailVerifier) createToken(ctx core.RequestContext, mail string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["email"] = mail
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte(verifier.key))
	if err != nil {
		return "", errors.WrapError(ctx, err)
	}
	return tokenString, nil
}

func (verifier *EmailVerifier) verifyToken(ctx core.RequestContext, tokenVal string) (string, error) {

	token, err := jwt.Parse(tokenVal, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || method != jwt.SigningMethodHS512 {
			log.Error(ctx, "Invalid Token", "method", method)
			return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_REQUEST)
		}
		return verifier.key, nil
	})
	if err == nil && token.Valid {
		log.Error(ctx, "Token validated")
		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"]
		return email.(string), nil
	}
	return "", err
}
