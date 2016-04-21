package server

import (
	"laatoo/core/common"
	"laatoo/core/security"
	"laatoo/sdk/auth"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"
)

type securityHandler struct {
	*common.Context
	handler       security.SecurityPlugin
	anonymousUser auth.User
	adminRole     string
	roleObject    string
	userObject    string
	jwtSecret     string
	authHeader    string
}

func newSecurityHandler(ctx core.ServerContext, name string, parent core.ServerElement) (server.ServerElementHandle, core.ServerElement) {
	shCtx := parent.NewCtx(name)
	sh := &securityHandler{Context: shCtx.(*common.Context)}
	return sh, sh
}

func (sh *securityHandler) Initialize(ctx core.ServerContext, conf config.Config) error {
	mode, ok := conf.GetString(config.CONF_SECURITY_MODE)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", config.CONF_SECURITY_MODE)
	}
	adminRole, ok := ctx.GetString(config.ADMINROLE)
	if ok {
		sh.adminRole = adminRole
	} else {
		sh.adminRole = config.DEFAULT_ADMIN
	}
	sh.Set(config.ADMINROLE, sh.adminRole)

	roleobject, ok := ctx.GetString(config.ROLE)
	if !ok {
		sh.roleObject = config.DEFAULT_ROLE
	} else {
		sh.roleObject = roleobject
	}
	sh.Set(config.ROLE, sh.roleObject)

	userObject, ok := conf.GetString(config.USER)
	if !ok {
		userObject = config.DEFAULT_USER
	}
	sh.userObject = userObject
	sh.Set(config.USER, sh.userObject)
	usr, err := ctx.CreateObject(userObject, nil)
	if err != nil {
		return err
	}
	anonymousUser, ok := usr.(auth.RbacUser)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_TYPE_MISMATCH)
	}
	anonymousUser.SetId("Anonymous")
	anonymousUser.SetRoles([]string{"Anonymous"})
	sh.anonymousUser = anonymousUser

	jwtSecret, ok := conf.GetString(config.JWTSECRET)
	if !ok {
		jwtSecret = config.DEFAULT_JWTSECRET
	}
	sh.jwtSecret = jwtSecret
	sh.Set(config.JWTSECRET, jwtSecret)

	authToken, ok := conf.GetString(config.AUTHHEADER)
	if !ok {
		authToken = config.DEFAULT_AUTHHEADER
	}
	sh.authHeader = authToken
	sh.Set(config.AUTHHEADER, authToken)

	switch mode {
	case config.CONF_SECURITY_LOCAL:
		plugin, err := security.NewLocalSecurityHandler(ctx, sh, conf, sh.roleObject, sh.adminRole, sh.userObject, sh.jwtSecret, sh.authHeader)
		if err != nil {
			return err
		}
		sh.handler = plugin
		return nil
	case config.CONF_SECURITY_REMOTE:
		return nil
	default:
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", config.CONF_SECURITY_MODE)
	}
}

func (sh *securityHandler) Start(ctx core.ServerContext) error {
	return sh.handler.Start(ctx)
}

func (sh *securityHandler) CreateSystemRequest(ctx core.ServerContext, name string) core.RequestContext {
	reqCtx := ctx.CreateNewRequest(name, nil).(*requestContext)
	reqCtx.user = nil
	reqCtx.admin = true
	return reqCtx
}

func (sh *securityHandler) AuthenticateRequest(ctx core.RequestContext) error {
	usr, isadmin, err := sh.handler.GetUser(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if usr == nil {
		usr = sh.anonymousUser
		isadmin = false
	}
	reqCtx := ctx.(*requestContext)
	reqCtx.user = usr
	reqCtx.admin = isadmin
	return nil
}

func (sh *securityHandler) HasPermission(ctx core.RequestContext, perm string) bool {
	return sh.handler.HasPermission(ctx, perm)
}
func (sh *securityHandler) GetRolePermissions(ctx core.RequestContext, roles []string) ([]string, bool) {
	return sh.handler.GetRolePermissions(ctx, roles)
}

//creates a context specific to environment
func (sh *securityHandler) createContext(ctx core.ServerContext, name string) core.ServerContext {
	return ctx.NewContextWithElements(name,
		core.ContextMap{core.ServerElementSecurityHandler: sh}, core.ServerElementSecurityHandler)
}
