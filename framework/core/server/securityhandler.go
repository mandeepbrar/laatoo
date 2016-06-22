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
	roleObject    string
	authHeader    string
	publicKey     *rsa.PublicKey
}

func newSecurityHandler(ctx core.ServerContext, name string, parent core.ServerElement) (server.ServerElementHandle, core.ServerElement) {
	shCtx := parent.NewCtx(name)
	sh := &securityHandler{Context: shCtx.(*common.Context)}
	return sh, sh
}

func (sh *securityHandler) Initialize(ctx core.ServerContext, conf config.Config) error {
	initCtx := sh.createContext(ctx, "Initialize Security Handler")
	adminRole, ok := conf.GetString(config.ADMINROLE)
	if !ok {
		adminRole = config.DEFAULT_ADMIN
	}
	sh.Set(config.ADMINROLE, adminRole)

	roleobject, ok := conf.GetString(config.ROLE)
	if !ok {
		roleobject = config.DEFAULT_ROLE
	}
	sh.Set(config.ROLE, roleobject)

	roleCreator, err := initCtx.GetObjectCreator(roleobject)
	if err != nil {
		return errors.WrapError(initCtx, err)
	}
	sh.roleCreator = roleCreator

	userObject, ok := conf.GetString(config.USER)
	if !ok {
		userObject = config.DEFAULT_USER
	}
	sh.Set(config.USER, userObject)

	userCreator, err := initCtx.GetObjectCreator(userObject)
	if err != nil {
		return errors.WrapError(initCtx, err)
	}
	sh.userCreator = userCreator

	publicKeyPath, ok := conf.GetString(config.CONF_PUBLICKEYPATH)
	if !ok {
		return errors.ThrowError(initCtx, errors.CORE_ERROR_MISSING_CONF, "conf", config.CONF_PUBLICKEYPATH)
	}
	publicKey, err := utils.LoadPublicKey(publicKeyPath)
	if err != nil {
		return errors.RethrowError(initCtx, errors.CORE_ERROR_BAD_CONF, err, "conf", config.CONF_PUBLICKEYPATH)
	}
	sh.publicKey = publicKey

	usr, err := userCreator(initCtx, nil)
	if err != nil {
		return err
	}
	anonymousUser, ok := usr.(auth.RbacUser)
	if !ok {
		return errors.ThrowError(initCtx, errors.CORE_ERROR_TYPE_MISMATCH)
	}
	anonymousUser.SetId("Anonymous")
	anonymousUser.SetRoles([]string{"Anonymous"})
	sh.anonymousUser = anonymousUser

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

	switch mode {
	case config.CONF_SECURITY_LOCAL:
		plugin, err := security.NewLocalSecurityHandler(initCtx, conf, adminRole, sh.roleCreator)
		if err != nil {
			return err
		}
		sh.handler = plugin
		return nil
	case config.CONF_SECURITY_REMOTE:
		plugin, err := security.NewRemoteSecurityHandler(initCtx, conf, adminRole, authToken, roleobject)
		if err != nil {
			return err
		}
		sh.handler = plugin
		return nil
	default:
		return errors.ThrowError(initCtx, errors.CORE_ERROR_BAD_CONF, "conf", config.CONF_SECURITY_MODE)
	}
}

func (sh *securityHandler) Start(ctx core.ServerContext) error {
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
