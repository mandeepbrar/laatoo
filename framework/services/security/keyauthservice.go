package security

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"laatoo/sdk/auth"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	//	"laatoo/sdk/log"
	"laatoo/sdk/utils"
)

const (
	//login path to be used for local and oauth authentication
	CONF_KEYAUTH_CLIENTS           = "clients"
	CONF_KEYAUTH_CLIENT_ROLE       = "role"
	CONF_KEYAUTH_CLIENT_IDENTIFIER = "id"
)

type client struct {
	role       string
	key        *rsa.PublicKey
	identifier string
}

type KeyAuthService struct {
	clients        map[string]*client
	userCreator    core.ObjectCreator
	tokenGenerator func(auth.User, string) (string, auth.User, error)
	adminRole      string
	jwtSecret      string
	authHeader     string
	userObject     string
	name           string
	localRealm     string
}

func (ks *KeyAuthService) Initialize(ctx core.ServerContext, conf config.Config) error {
	sechandler := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sechandler == nil {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	authHeader, ok := sechandler.GetString(config.AUTHHEADER)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ks.authHeader = authHeader
	localRealm, ok := sechandler.GetString(config.REALM)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ks.localRealm = localRealm
	userObject, ok := sechandler.GetString(config.USER)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ks.userObject = userObject
	adminRole, ok := sechandler.GetString(config.ADMINROLE)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ks.adminRole = adminRole

	userCreator, err := ctx.GetObjectCreator(userObject)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	ks.userCreator = userCreator
	clientsConf, ok := conf.GetSubConfig(CONF_KEYAUTH_CLIENTS)
	if ok {
		allclients := clientsConf.AllConfigurations()
		ks.clients = make(map[string]*client, len(allclients))
		for _, clientName := range allclients {
			clientConf, _ := clientsConf.GetSubConfig(clientName)
			role, ok := clientConf.GetString(CONF_KEYAUTH_CLIENT_ROLE)
			if ok {
				pubKeyPath, _ := clientConf.GetString(config.CONF_PUBLICKEYPATH)
				pubKey, err := utils.LoadPublicKey(pubKeyPath)
				if err != nil {
					return errors.RethrowError(ctx, errors.CORE_ERROR_BAD_CONF, err, "path", pubKeyPath)
				}
				ks.clients[clientName] = &client{role: role, key: pubKey, identifier: clientName}
			}
		}
	} else {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_KEYAUTH_CLIENTS)
	}
	return nil
}

func (ks *KeyAuthService) Invoke(ctx core.RequestContext) error {
	id, ok := ctx.GetString(CONF_KEYAUTH_CLIENT_IDENTIFIER)
	if !ok {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	client, ok := ks.clients[id]
	if !ok {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	req := ctx.GetRequest()
	data, ok := req.([]byte)
	if !ok {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}

	// compute the sha1
	h := sha256.New()
	h.Write([]byte(id))

	err := rsa.VerifyPKCS1v15(client.key, crypto.SHA256, h.Sum(nil), data)
	if err != nil {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}

	//create the user
	usrInt := ks.userCreator()
	init := usrInt.(core.Initializable)
	init.Init(ctx, core.MethodArgs{"Id": "system", "Roles": []string{client.role}, "Realm": ks.localRealm})

	usr := usrInt.(auth.RbacUser)
	/*usr.SetId("system")
	usr.SetRoles([]string{client.role})*/
	token, user, err := ks.tokenGenerator(usr, ks.localRealm)
	if err != nil {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	resp := core.NewServiceResponse(core.StatusSuccess, user, map[string]interface{}{ks.authHeader: token})
	ctx.SetResponse(resp)
	return nil
}
func (ks *KeyAuthService) SetTokenGenerator(ctx core.ServerContext, gen func(auth.User, string) (string, auth.User, error)) {
	ks.tokenGenerator = gen
}
func (ks *KeyAuthService) Start(ctx core.ServerContext) error {
	return nil
}
