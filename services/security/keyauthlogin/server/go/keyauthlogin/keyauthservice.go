package keyauthlogin

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/auth"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	common "securitycommon"
)

const (
	//login path to be used for local and oauth authentication
	CONF_KEYAUTH_CLIENTS           = "clients"
	CONF_KEYAUTH_CLIENT_ROLE       = "role"
	CONF_KEYAUTH_CLIENT_IDENTIFIER = "id"
	CONF_SECURITYSERVICE_KEYAUTH   = "KEYAUTH"
)

type client struct {
	role       string
	key        *rsa.PublicKey
	identifier string
}

type KeyAuthService struct {
	core.Service
	clients        map[string]*client
	tokenGenerator func(auth.User, string) (string, auth.User, error)
	adminRole      string
	jwtSecret      string
	authHeader     string
	userObject     string
	name           string
	localRealm     string
}

/*
func (ks *KeyAuthService) Describe(ctx core.ServerContext) {
	ks.SetDescription("Keyauth service")
	ks.AddParam(CONF_KEYAUTH_CLIENT_IDENTIFIER, config.OBJECTTYPE_STRING, false)
	ks.SetRequestType(config.OBJECTTYPE_BYTES, false, false)
	ks.AddConfigurations(map[string]string{CONF_KEYAUTH_CLIENTS: config.OBJECTTYPE_CONFIG})
}*/

func (ks *KeyAuthService) Initialize(ctx core.ServerContext, conf config.Config) error {

	sechandler := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sechandler == nil {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	authHeader := sechandler.GetProperty(config.AUTHHEADER)
	if authHeader == nil {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ks.authHeader = authHeader.(string)
	localRealm := sechandler.GetProperty(config.REALM)
	if localRealm == nil {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ks.localRealm = localRealm.(string)
	userObject := sechandler.GetProperty(config.USER)
	if userObject == nil {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ks.userObject = userObject.(string)
	adminRole := sechandler.GetProperty(config.ADMINROLE)
	if adminRole == nil {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ks.adminRole = adminRole.(string)

	return nil
}

func (ks *KeyAuthService) Invoke(ctx core.RequestContext, req core.Request) error {
	idparam, ok := req.GetParam(ctx, CONF_KEYAUTH_CLIENT_IDENTIFIER)
	id := idparam.GetValue().(string)
	client, ok := ks.clients[id]

	if !ok {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	bodyParam := "Data"
	reqbytes, ok := ctx.GetParamValue(bodyParam)
	if !ok {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	//:= req.GetBody()
	data, ok := reqbytes.([]byte)
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
	usrInt, _ := ctx.CreateObject(ks.userObject)

	init := usrInt.(core.Initializable)

	userConf := ctx.ServerContext().CreateConfig()
	userConf.Set(ctx, "Id", fmt.Sprint("system_", id))
	userConf.Set(ctx, "Roles", []string{client.role})
	userConf.Set(ctx, "Realm", ks.localRealm)

	init.Initialize(ctx, userConf)

	usr := usrInt.(auth.RbacUser)
	/*usr.SetId("system")
	usr.SetRoles([]string{client.role})*/
	token, user, err := ks.tokenGenerator(usr, ks.localRealm)
	if err != nil {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}

	info := map[string]interface{}{ks.authHeader: token}

	err = ctx.SendSynchronousMessage(common.EVT_LOGIN_SUCCESS, map[string]interface{}{"Data": user, "info": info})
	if err != nil {
		log.Error(ctx, "Encountered Error in sending event", "error", err)
	}
	ctx.SetResponse(core.SuccessResponseWithInfo(user, info))
	return nil
}

func (ks *KeyAuthService) SetTokenGenerator(ctx core.ServerContext, gen func(auth.User, string) (string, auth.User, error)) {
	ks.tokenGenerator = gen
}

func (ks *KeyAuthService) Start(ctx core.ServerContext) error {

	c, ok := ks.GetConfiguration(ctx, CONF_KEYAUTH_CLIENTS)
	if ok {
		clientsConf := c.(config.Config)
		allclients := clientsConf.AllConfigurations(ctx)
		ks.clients = make(map[string]*client, len(allclients))
		for _, clientName := range allclients {
			clientConf, _ := clientsConf.GetSubConfig(ctx, clientName)
			role, ok := clientConf.GetString(ctx, CONF_KEYAUTH_CLIENT_ROLE)
			if ok {
				pubKeyPath, _ := clientConf.GetString(ctx, config.CONF_PUBLICKEYPATH)
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
