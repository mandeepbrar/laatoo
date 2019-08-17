package core

import (
	"crypto/rsa"
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/auth"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"laatoo/server/constants"
	"laatoo/server/security"

	jwt "github.com/dgrijalva/jwt-go"
)

type securityHandler struct {
	name          string
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
	skipSecurity  bool
}

func newSecurityHandler(ctx core.ServerContext, name string, parent core.ServerElement) (elements.ServerElementHandle, core.ServerElement) {
	sh := &securityHandler{name: name}
	proxy := &securityHandlerProxy{secHandler: sh}
	return sh, proxy
}

func (sh *securityHandler) Initialize(ctx core.ServerContext, conf config.Config) error {
	if conf == nil {
		sh.skipSecurity = true
		log.Trace(ctx, "Skipping security handler start")
		return nil
	}

	initCtx := ctx.SubContext("Initialize Security Handler")
	sh.securityConf = conf

	realm, _ := conf.GetString(ctx, config.REALM)
	sh.realm = realm

	adminRole, ok := conf.GetString(ctx, config.ADMINROLE)
	if !ok {
		adminRole = config.DEFAULT_ADMIN
	}
	if realm != "" {
		adminRole = fmt.Sprintf("%s_%s", adminRole, realm)
	}
	sh.adminRole = adminRole

	anonRole := "Anonymous"
	if realm != "" {
		anonRole = fmt.Sprintf("%s_%s", anonRole, realm)
	}
	sh.anonRole = anonRole

	roleobject, ok := conf.GetString(ctx, config.ROLE)
	if !ok {
		roleobject = config.DEFAULT_ROLE
	}
	sh.roleObject = roleobject

	userObject, ok := conf.GetString(ctx, config.USER)
	if !ok {
		userObject = config.DEFAULT_USER
	}
	sh.userObject = userObject

	if !sh.skipSecurity {
		publicKeyPath, ok := conf.GetString(ctx, config.CONF_PUBLICKEYPATH)
		if !ok {
			return errors.ThrowError(initCtx, errors.CORE_ERROR_MISSING_CONF, "conf", config.CONF_PUBLICKEYPATH)
		}
		publicKey, err := utils.LoadPublicKey(publicKeyPath)
		if err != nil {
			return errors.RethrowError(initCtx, errors.CORE_ERROR_BAD_CONF, err, "conf", config.CONF_PUBLICKEYPATH)
		}
		sh.publicKey = publicKey
	}

	authToken, ok := conf.GetString(ctx, config.AUTHHEADER)
	if !ok {
		authToken = config.DEFAULT_AUTHHEADER
	}
	sh.authHeader = authToken
	mode, ok := conf.GetString(ctx, constants.CONF_SECURITY_MODE)
	if !ok {
		mode = constants.CONF_SECURITY_LOCAL
	}
	sh.securityMode = mode

	log.Trace(initCtx, "Security handler initialized")

	return nil
}

func (sh *securityHandler) Start(ctx core.ServerContext) error {
	startCtx := ctx.SubContext("Starting Security Handler")
	if sh.skipSecurity {
		log.Trace(startCtx, "Skipping security handler start")
		return nil
	}
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
	userConf := ctx.CreateConfig()
	userConf.Set(ctx, "Id", "Anonymous")
	userConf.Set(ctx, "Roles", []string{sh.anonRole})
	err = init.Initialize(startCtx, userConf)
	if err != nil {
		return err
	}
	sh.anonymousUser = anonymousUser.(auth.User)

	if sh.skipSecurity {
		return nil
	}

	switch sh.securityMode {
	case constants.CONF_SECURITY_LOCAL:
		plugin, err := security.NewLocalSecurityHandler(startCtx, sh.securityConf, sh.adminRole, sh.anonRole, sh.roleCreator, sh.realm)
		if err != nil {
			return err
		}
		sh.handler = plugin
	case constants.CONF_SECURITY_REMOTE:
		plugin, err := security.NewRemoteSecurityHandler(startCtx, sh.securityConf, sh.adminRole, sh.anonRole, sh.authHeader, sh.roleObject, sh.realm)
		if err != nil {
			return err
		}
		sh.handler = plugin
	default:
		return errors.ThrowError(startCtx, errors.CORE_ERROR_BAD_CONF, "conf", constants.CONF_SECURITY_MODE)
	}
	return nil
}

func (sh *securityHandler) onServerReady(ctx core.ServerContext) error {
	if sh.skipSecurity {
		return nil
	}
	return sh.handler.Start(ctx)
}

func (sh *securityHandler) authenticateRequest(ctx core.RequestContext, loadFresh bool) (string, error) {
	if sh.skipSecurity {
		return "", nil
	}
	usr, isadmin, token, err := sh.getUserFromToken(ctx, loadFresh)
	if err != nil {
		return token, errors.WrapError(ctx, err)
	}
	if usr == nil {
		usr = sh.anonymousUser
		isadmin = false
	}
	log.Info(ctx, "Authenticated request", "User", usr.GetId())
	reqCtx := ctx.(*requestContext)
	reqCtx.user = usr
	reqCtx.admin = isadmin
	return token, nil
}

func (sh *securityHandler) hasPermission(ctx core.RequestContext, perm string) bool {
	if sh.skipSecurity {
		return true
	}
	return sh.handler.HasPermission(ctx, perm)
}

func (sh *securityHandler) allPermissions(ctx core.RequestContext) []string {
	if sh.skipSecurity {
		return []string{}
	}
	return sh.handler.AllPermissions(ctx)
}

/*//creates a context specific to environment
func (sh *securityHandler) createContext(ctx core.ServerContext, name string) core.ServerContext {
	return ctx.NewContextWithElements(name,
		core.ContextMap{core.ServerElementSecurityHandler: sh}, core.ServerElementSecurityHandler)
}*/

func (sh *securityHandler) getUserFromToken(ctx core.RequestContext, loadFresh bool) (auth.User, bool, string, error) {
	tokenVal, ok := ctx.GetString(sh.authHeader)
	if ok {
		token, err := jwt.Parse(tokenVal, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if method, ok := token.Method.(*jwt.SigningMethodRSA); !ok || method != jwt.SigningMethodRS512 {
				log.Trace(ctx, "Invalid Token", "method", method)
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_REQUEST)
			}
			return sh.publicKey, nil
		})
		if err == nil && token.Valid {
			log.Trace(ctx, "Token validated")
			userInt := sh.userCreator()
			user, ok := userInt.(auth.RbacUser)
			if !ok {
				return nil, false, tokenVal, errors.ThrowError(ctx, errors.CORE_ERROR_TYPE_MISMATCH)
			}
			claims := token.Claims.(jwt.MapClaims)
			realm := claims[config.REALM]
			log.Info(ctx, "token realms", "realm", realm, "expected", sh.realm)
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
	log.Trace(ctx, "Token invalid")
	return nil, false, tokenVal, nil
}
