package core

import (
	"crypto/rsa"
	"fmt"
	"laatoo/sdk/auth"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/sdk/utils"
	"laatoo/server/common"
	"laatoo/server/security"

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
	anonRole      string
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

	realm, _ := conf.GetString(config.REALM)
	sh.realm = realm
	sh.Set(config.REALM, realm)

	adminRole, ok := conf.GetString(config.ADMINROLE)
	if !ok {
		adminRole = config.DEFAULT_ADMIN
	}
	if realm != "" {
		adminRole = fmt.Sprintf("%s_%s", adminRole, realm)
	}
	sh.adminRole = adminRole
	sh.Set(config.ADMINROLE, adminRole)

	anonRole := "Anonymous"
	if realm != "" {
		anonRole = fmt.Sprintf("%s_%s", anonRole, realm)
	}
	sh.anonRole = anonRole

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

	anonymousUser := userCreator()
	init := anonymousUser.(core.Initializable)
	err = init.Init(startCtx, core.MethodArgs{"Id": "Anonymous", "Roles": []string{sh.anonRole}})
	if err != nil {
		return err
	}
	sh.anonymousUser = anonymousUser.(auth.User)

	switch sh.securityMode {
	case config.CONF_SECURITY_LOCAL:
		plugin, err := security.NewLocalSecurityHandler(startCtx, sh.securityConf, sh.adminRole, sh.anonRole, sh.roleCreator, sh.realm)
		if err != nil {
			return err
		}
		sh.handler = plugin
	case config.CONF_SECURITY_REMOTE:
		plugin, err := security.NewRemoteSecurityHandler(startCtx, sh.securityConf, sh.adminRole, sh.anonRole, sh.authHeader, sh.roleObject, sh.realm)
		if err != nil {
			return err
		}
		sh.handler = plugin
	default:
		return errors.ThrowError(startCtx, errors.CORE_ERROR_BAD_CONF, "conf", config.CONF_SECURITY_MODE)
	}

	return sh.handler.Start(ctx)
}

func (sh *securityHandler) AuthenticateRequest(ctx core.RequestContext, loadFresh bool) (string, error) {
	usr, isadmin, token, err := sh.getUserFromToken(ctx, loadFresh)
	if err != nil {
		return token, errors.WrapError(ctx, err)
	}
	if usr == nil {
		usr = sh.anonymousUser
		isadmin = false
	}
	log.Logger.Info(ctx, "Authenticated request", "User", usr.GetId())
	reqCtx := ctx.(*requestContext)
	reqCtx.user = usr
	reqCtx.admin = isadmin
	return token, nil
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

func (sh *securityHandler) getUserFromToken(ctx core.RequestContext, loadFresh bool) (auth.User, bool, string, error) {
	tokenVal, ok := ctx.GetString(sh.authHeader)
	log.Logger.Trace(ctx, "Token received", "token", tokenVal)
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
			log.Logger.Trace(ctx, "Token validated")
			userInt := sh.userCreator()
			user, ok := userInt.(auth.RbacUser)
			if !ok {
				return nil, false, tokenVal, errors.ThrowError(ctx, errors.CORE_ERROR_TYPE_MISMATCH)
			}
			claims := token.Claims.(jwt.MapClaims)
			realm := claims[config.REALM]
			log.Logger.Info(ctx, "token realms", "realm", realm, "expected", sh.realm)
			if realm != sh.realm {
				return nil, false, tokenVal, nil
			}

			user.LoadClaims(claims)
			admin := false
			if loadFresh {
				roles, err := user.GetRoles()
				if err != nil {
					return nil, false, tokenVal, err
				}
				var perms []string
				perms, admin = sh.handler.GetRolePermissions(roles, sh.realm)
				user.SetPermissions(perms)
			} else {
				adminClaim := claims["Admin"]
				if adminClaim != nil {
					adminVal, ok := adminClaim.(bool)
					if ok {
						admin = adminVal
					}
				}
			}
			return user, admin, tokenVal, nil
		}
	}
	log.Logger.Trace(ctx, "Token invalid")
	return nil, false, tokenVal, nil
}