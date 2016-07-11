package server

import (
	"crypto/rsa"
	"laatoo/framework/core/common"
	"laatoo/framework/core/security"
	"laatoo/sdk/auth"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/sdk/utils"

	jwt "github.com/dgrijalva/jwt-go"
)

type securityHandler struct {
	*common.Context
	handler       security.SecurityPlugin
	userCreator   core.ObjectCreator
	roleCreator   core.ObjectCreator
	anonymousUser auth.User
	securityMode  string
	userObject    string
	roleObject    string
	authHeader    string
	adminRole     string
	securityConf  config.Config
	publicKey     *rsa.PublicKey
	realm         string
}

func newSecurityHandler(ctx core.ServerContext, name string, parent core.ServerElement) (server.ServerElementHandle, core.ServerElement) {
	shCtx := parent.NewCtx(name)
	sh := &securityHandler{Context: shCtx.(*common.Context)}
	return sh, sh
}

func (sh *securityHandler) Initialize(ctx core.ServerContext, conf config.Config) error {
	initCtx := sh.createContext(ctx, "Initialize Security Handler")
	sh.securityConf = conf
	adminRole, ok := conf.GetString(config.ADMINROLE)
	if !ok {
		adminRole = config.DEFAULT_ADMIN
	}
	sh.adminRole = adminRole
	sh.Set(config.ADMINROLE, adminRole)

	roleobject, ok := conf.GetString(config.ROLE)
	if !ok {
		roleobject = config.DEFAULT_ROLE
	}
	sh.Set(config.ROLE, roleobject)
	sh.roleObject = roleobject

	userObject, ok := conf.GetString(config.USER)
	if !ok {
		userObject = config.DEFAULT_USER
	}
	sh.Set(config.USER, userObject)
	sh.userObject = userObject

	realm, _ := conf.GetString(config.REALM)
	sh.realm = realm
	sh.Set(config.REALM, realm)

	publicKeyPath, ok := conf.GetString(config.CONF_PUBLICKEYPATH)
	if !ok {
		return errors.ThrowError(initCtx, errors.CORE_ERROR_MISSING_CONF, "conf", config.CONF_PUBLICKEYPATH)
	}
	publicKey, err := utils.LoadPublicKey(publicKeyPath)
	if err != nil {
		return errors.RethrowError(initCtx, errors.CORE_ERROR_BAD_CONF, err, "conf", config.CONF_PUBLICKEYPATH)
	}
	sh.publicKey = publicKey

	authToken, ok := conf.GetString(config.AUTHHEADER)
	if !ok {
		authToken = config.DEFAULT_AUTHHEADER
	}
	sh.authHeader = authToken
	sh.Set(config.AUTHHEADER, authToken)

	mode, ok := conf.GetString(config.CONF_SECURITY_MODE)
	if !ok {
		return errors.ThrowError(initCtx, errors.CORE_ERROR_MISSING_CONF, "conf", config.CONF_SECURITY_MODE)
	}
	sh.securityMode = mode
	return nil
}

func (sh *securityHandler) Start(ctx core.ServerContext) error {
	startCtx := sh.createContext(ctx, "Starting Security Handler")
	userCreator, err := startCtx.GetObjectCreator(sh.userObject)
	if err != nil {
		return errors.WrapError(startCtx, err)
	}
	sh.userCreator = userCreator

	roleCreator, err := startCtx.GetObjectCreator(sh.roleObject)
	if err != nil {
		return errors.WrapError(startCtx, err)
	}
	sh.roleCreator = roleCreator

	anonymousUser, err := userCreator(startCtx, core.MethodArgs{"Id": "Anonymous", "Roles": []string{"Anonymous"}})
	if err != nil {
		return err
	}
	sh.anonymousUser = anonymousUser.(auth.User)

	switch sh.securityMode {
	case config.CONF_SECURITY_LOCAL:
		plugin, err := security.NewLocalSecurityHandler(startCtx, sh.securityConf, sh.adminRole, sh.roleCreator, sh.realm)
		if err != nil {
			return err
		}
		sh.handler = plugin
	case config.CONF_SECURITY_REMOTE:
		plugin, err := security.NewRemoteSecurityHandler(startCtx, sh.securityConf, sh.adminRole, sh.authHeader, sh.roleObject, sh.realm)
		if err != nil {
			return err
		}
		sh.handler = plugin
	default:
		return errors.ThrowError(startCtx, errors.CORE_ERROR_BAD_CONF, "conf", config.CONF_SECURITY_MODE)
	}

	return sh.handler.Start(ctx)
}

func (sh *securityHandler) AuthenticateRequest(ctx core.RequestContext) error {
	usr, isadmin, err := sh.getUserFromToken(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if usr == nil {
		usr = sh.anonymousUser
		isadmin = false
	}
	log.Logger.Info(ctx, "Authenticated request", "User", usr.GetId())
	reqCtx := ctx.(*requestContext)
	reqCtx.user = usr
	reqCtx.admin = isadmin
	return nil
}

func (sh *securityHandler) HasPermission(ctx core.RequestContext, perm string) bool {
	return sh.handler.HasPermission(ctx, perm)
}

func (sh *securityHandler) AllPermissions(ctx core.RequestContext) []string {
	return sh.handler.AllPermissions(ctx)
}

//creates a context specific to environment
func (sh *securityHandler) createContext(ctx core.ServerContext, name string) core.ServerContext {
	return ctx.NewContextWithElements(name,
		core.ContextMap{core.ServerElementSecurityHandler: sh}, core.ServerElementSecurityHandler)
}

func (sh *securityHandler) getUserFromToken(ctx core.RequestContext) (auth.User, bool, error) {
	tokenVal, ok := ctx.GetString(sh.authHeader)
	if ok {
		token, err := jwt.Parse(tokenVal, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if method, ok := token.Method.(*jwt.SigningMethodRSA); !ok || method != jwt.SigningMethodRS512 {
				log.Logger.Trace(ctx, "Invalid Token", "method", method)
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_REQUEST)
			}
			return sh.publicKey, nil
		})
		if err == nil && token.Valid {
			userInt, err := sh.userCreator(ctx, nil)
			if err != nil {
				return nil, false, errors.WrapError(ctx, err)
			}
			user, ok := userInt.(auth.RbacUser)
			if !ok {
				return nil, false, errors.ThrowError(ctx, errors.CORE_ERROR_TYPE_MISMATCH)
			}
			realm := token.Claims[config.REALM]
			log.Logger.Info(ctx, "token realms", "realm", realm, "expected", sh.realm)
			if realm != sh.realm {
				return nil, false, nil
			}

			user.LoadJWTClaims(token)
			admin := false
			adminClaim := token.Claims["Admin"]
			if adminClaim != nil {
				adminVal, ok := adminClaim.(bool)
				if ok {
					admin = adminVal
				}
			}
			return user, admin, nil
		}
	}
	return nil, false, nil
}
