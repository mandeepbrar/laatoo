package laatooauthentication

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"laatoocore"
	"laatoosdk/auth"
	"laatoosdk/errors"
	"laatoosdk/log"
)

const (
	//login path to be used for local and oauth authentication
	CONF_AUTHSERVICE_KEYPATH    = "/keyauth"
	CONF_AUTHSERVICE_PVTKEYPATH = "pvtkey"
	CONF_AUTHSERVICE_DOMAINS    = "domains"
)

type keyAuthType struct {
	//reference to the main auth service
	securityService *SecurityService
	//method called in case of callback
	authCallback echo.HandlerFunc
	pvtKeyPath   string
	domains      map[string]interface{}
}

//method called for creating new auth type
func NewKeyAuth(ctx interface{}, conf map[string]interface{}, svc *SecurityService) (*keyAuthType, error) {
	//create the new auth type
	keyauth := &keyAuthType{}

	pvtKeyPath, ok := conf[CONF_AUTHSERVICE_PVTKEYPATH]
	if ok {
		keyauth.pvtKeyPath = pvtKeyPath.(string)
	}
	domains, ok := conf[CONF_AUTHSERVICE_DOMAINS]
	if ok {
		keyauth.domains = domains.(map[string]interface{})
	}
	//store the reference to the parent
	keyauth.securityService = svc
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "keyAuthProvider: Initializing")

	return keyauth, nil
}

//initialize auth type called by base auth for initializing
func (keyauth *keyAuthType) InitializeType(ctx interface{}, authStart echo.HandlerFunc, authCallback echo.HandlerFunc) error {
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "Settingup Api Auth")
	//setup path for listening to login post request
	keyauth.securityService.Router.Post(CONF_AUTHSERVICE_KEYPATH, authStart)
	keyauth.authCallback = authCallback
	return nil
}

//validate the local user
//derive the data from context object
func (keyauth *keyAuthType) ValidateUser(ctx *echo.Context) error {
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "keyauth: Validating Credentials")

	if keyauth.domains == nil {
		return errors.ThrowError(ctx, AUTH_ERROR_DOMAIN_NOT_ALLOWED)
	}
	//create the user
	usrInt, err := keyauth.securityService.CreateUser(ctx)
	if err != nil {
		return errors.RethrowError(ctx, laatoocore.AUTH_ERROR_USEROBJECT_NOT_CREATED, err)
	}

	authform := &laatoocore.KeyAuth{}

	err = ctx.Bind(authform)
	if err != nil {
		return errors.RethrowError(ctx, AUTH_ERROR_INCORRECT_REQ_FORMAT, err)
	}

	pvtKey, err := loadPrivateKey(keyauth.pvtKeyPath)
	if err != nil {
		return errors.RethrowError(ctx, AUTH_ERROR_KEYAUTH_MISSING_PVTKEY, err)
	}

	out, err := rsa.DecryptOAEP(md5.New(), rand.Reader, pvtKey, authform.Key, []byte(""))
	if err != nil {
		return err
	}

	domain := string(out)
	role, ok := keyauth.domains[domain]
	if ok {
		usr := usrInt.(auth.RbacUser)
		usr.SetId("system")
		usr.SetRoles([]string{role.(string)})
		ctx.Set("User", usr)
	} else {
		return errors.ThrowError(ctx, AUTH_ERROR_DOMAIN_NOT_ALLOWED, "Domain", domain)
	}
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "Auth Key Validated", "Domain", domain, " Role assigned", role.(string))

	return keyauth.authCallback(ctx)
}

func (keyauth *keyAuthType) GetName() string {
	return "keyauth"
}

//complete authentication
func (keyauth *keyAuthType) CompleteAuthentication(ctx *echo.Context) error {
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "keyAuthProvider: Authentication Successful")
	return nil
}

// loadPrivateKey loads an parses a PEM encoded private key file.
func loadPrivateKey(path string) (*rsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("ssh: no key found")
	}

	switch block.Type {
	case "RSA PRIVATE KEY":
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}
}
