package security

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"laatoo/core/registry"
	"laatoo/sdk/auth"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	//login path to be used for local and oauth authentication
	CONF_KEYAUTH_PVTKEYPATH = "pvtkey"
	CONF_KEYAUTH_DOMAINS    = "domains"
)

//service method for doing various tasks
func NewKeyAuthService(ctx core.ServerContext, conf config.Config) (core.Service, error) {
	return &KeyAuthService{conf: conf}, nil
}

type KeyAuthService struct {
	conf        config.Config
	pvtKey      *rsa.PrivateKey
	domains     map[string]string
	userCreator core.ObjectCreator
	adminRole   string
}

func (ks *KeyAuthService) Initialize(ctx core.ServerContext) error {
	pvtKeyPath, ok := ks.conf.GetString(CONF_KEYAUTH_PVTKEYPATH)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_KEYAUTH_PVTKEYPATH)
	}
	pvtKey, err := loadPrivateKey(ctx, pvtKeyPath)
	if err != nil {
		return errors.RethrowError(ctx, errors.CORE_ERROR_BAD_CONF, err, "conf", CONF_KEYAUTH_PVTKEYPATH)
	}
	ks.pvtKey = pvtKey
	userobject := ctx.GetServerVariable(core.USER)
	userCreator, err := registry.GetObjectCreator(ctx, userobject.(string))
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	ks.userCreator = userCreator
	domainsConf, ok := ks.conf.GetSubConfig(CONF_KEYAUTH_DOMAINS)
	if ok {
		alldomains := domainsConf.AllConfigurations()
		ks.domains = make(map[string]string, len(alldomains))
		for _, domain := range alldomains {
			role, _ := domainsConf.GetString(domain)
			ks.domains[domain] = role
		}
	} else {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_KEYAUTH_DOMAINS)
	}
	return nil
}

//Expects domain data encrypted with public key to be provided inside the request
func (ks *KeyAuthService) Invoke(ctx core.RequestContext) error {
	//create the user
	usrInt, _ := ks.userCreator(ctx, nil)

	publicKey := ctx.GetRequestBody().([]byte)

	out, err := rsa.DecryptOAEP(md5.New(), rand.Reader, ks.pvtKey, publicKey, []byte(""))
	if err != nil {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}

	domain := string(out)
	role, ok := ks.domains[domain]
	if ok {
		usr := usrInt.(auth.RbacUser)
		usr.SetId("system")
		usr.SetRoles([]string{role})
		resp, err := completeAuthentication(ctx, usr)
		if err != nil {
			ctx.SetResponse(core.StatusUnauthorizedResponse)
			return nil
		}
		ctx.SetResponse(resp)
	} else {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	log.Logger.Debug(ctx, "Auth Key Validated", "Domain", domain, " Role assigned", role)
	return nil
}

func (ks *KeyAuthService) GetConf() config.Config {
	return ks.conf
}
func (ks *KeyAuthService) GetResponseHandler() core.ServiceResponseHandler {
	return nil
}

// loadPrivateKey loads an parses a PEM encoded private key file.
func loadPrivateKey(ctx core.ServerContext, path string) (*rsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.ThrowError(ctx, "ssh: no key found")
	}

	switch block.Type {
	case "RSA PRIVATE KEY":
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	default:
		return nil, errors.ThrowError(ctx, fmt.Sprintf("ssh: unsupported key type %q", block.Type))
	}
}
