package security

import (
	"crypto/rsa"
	"fmt"
	"laatoo/sdk/auth"
	"laatoo/sdk/components"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/utils"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	CONF_SECURITY_ROLEDATASERVICE = "role_data_svc"
	CONF_SECURITY_PERMISSIONS     = "permissions"
	CONF_SECURITY_AUTHSERVICES    = "authservices"
)

type localSecurityHandler struct {
	roleCreator     core.ObjectCreator
	adminRole       string
	pvtKey          *rsa.PrivateKey
	roleDataService data.DataComponent
	roleDataSvcName string
	rolesMap        map[string]auth.Role
	authServices    []string
	allPermissions  []string
	//permissions assigned to a role
	rolePermissions map[string]bool
}

func NewLocalSecurityHandler(ctx core.ServerContext, conf config.Config, adminrole string, roleCreator core.ObjectCreator) (SecurityPlugin, error) {
	lsh := &localSecurityHandler{adminRole: adminrole, roleCreator: roleCreator}
	lsh.rolesMap = make(map[string]auth.Role, 10)
	//map containing roles and permissions
	lsh.rolePermissions = make(map[string]bool, 50)

	permissions, ok := conf.GetStringArray(CONF_SECURITY_PERMISSIONS)
	if ok {
		lsh.allPermissions = permissions
	} else {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_SECURITY_PERMISSIONS)
	}

	pvtKeyPath, ok := conf.GetString(config.CONF_PVTKEYPATH)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", config.CONF_PVTKEYPATH)
	}
	pvtKey, err := utils.LoadPrivateKey(pvtKeyPath)
	if err != nil {
		return nil, errors.RethrowError(ctx, errors.CORE_ERROR_BAD_CONF, err, "conf", config.CONF_PVTKEYPATH)
	}
	lsh.pvtKey = pvtKey

	roleDataSvcName, ok := conf.GetString(CONF_SECURITY_ROLEDATASERVICE)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", CONF_SECURITY_ROLEDATASERVICE)
	}
	lsh.roleDataSvcName = roleDataSvcName
	authServices, ok := conf.GetStringArray(CONF_SECURITY_AUTHSERVICES)
	if ok {
		lsh.authServices = authServices
	}
	return lsh, nil
}

func (lsh *localSecurityHandler) Start(ctx core.ServerContext) error {
	roleService, err := ctx.GetService(lsh.roleDataSvcName)
	if err != nil {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", CONF_SECURITY_ROLEDATASERVICE)
	}
	roleDataService, ok := roleService.(data.DataComponent)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", CONF_SECURITY_ROLEDATASERVICE)
	}
	lsh.roleDataService = roleDataService

	err = lsh.loadRoles(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if lsh.authServices != nil {
		for _, svcName := range lsh.authServices {
			svc, err := ctx.GetService(svcName)
			if err != nil {
				return err
			}
			authComp, ok := svc.(components.AuthenticationComponent)
			if !ok {
				return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", CONF_SECURITY_AUTHSERVICES, "svcname", svcName)
			}
			authComp.SetTokenGenerator(ctx, lsh.tokenGenerator(ctx))
		}
	}
	return nil
}

func (lsh *localSecurityHandler) tokenGenerator(ctx core.ServerContext) func(auth.User) (string, auth.User, error) {
	return func(user auth.User) (string, auth.User, error) {
		token := jwt.New(jwt.SigningMethodRS512)
		rbac, ok := user.(auth.RbacUser)
		if ok {
			roles, _ := rbac.GetRoles()
			permissions, admin := lsh.getRolePermissions(roles)
			rbac.SetPermissions(permissions)
			token.Claims["Admin"] = admin
		}
		user.SetJWTClaims(token)
		token.Claims["UserId"] = user.GetId()
		//token.Claims["IP"] = ctx.ClientIP()
		token.Claims["exp"] = time.Now().Add(time.Hour * 4).Unix()
		tokenString, err := token.SignedString(lsh.pvtKey)
		if err != nil {
			return "", nil, errors.WrapError(ctx, err)
		}
		return tokenString, user, nil
	}
}

func (lsh *localSecurityHandler) getRolePermissions(roles []string) ([]string, bool) {
	permissions := utils.NewStringSet([]string{})
	for _, rolename := range roles {
		role, ok := lsh.rolesMap[rolename]
		if ok {
			permissions.Append(role.GetPermissions())
		}
	}
	return permissions.Values(), false
}

func (lsh *localSecurityHandler) HasPermission(ctx core.RequestContext, perm string) bool {
	return hasPermission(ctx, perm, lsh.rolePermissions)
}

func (lsh *localSecurityHandler) AllPermissions(core.RequestContext) []string {
	return lsh.allPermissions
}

func (lsh *localSecurityHandler) loadRoles(ctx core.ServerContext) error {
	loadRolesReq := ctx.CreateSystemRequest("LoadRoles")
	defer loadRolesReq.CompleteRequest()
	roles, _, _, err := lsh.roleDataService.GetList(loadRolesReq, -1, -1, "", "")
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	adminExists := false
	anonExists := false
	for _, val := range roles {
		role, ok := val.(auth.Role)
		if !ok {
			errors.ThrowError(ctx, errors.CORE_ERROR_TYPE_MISMATCH)
		}
		id := role.GetId()
		roleName := role.GetName()
		if id == "Anonymous" {
			anonExists = true
		}

		if id == lsh.adminRole {
			adminExists = true
		}
		lsh.rolesMap[roleName] = role
		permissions := role.GetPermissions()
		for _, perm := range permissions {
			key := fmt.Sprintf("%s#%s", roleName, perm)
			lsh.rolePermissions[key] = true
		}
	}
	return lsh.createInitRoles(ctx, anonExists, adminExists)
}

func (lsh *localSecurityHandler) createInitRoles(ctx core.ServerContext, anonExists bool, adminExists bool) error {

	if !anonExists {
		ent, _ := lsh.roleCreator(ctx, nil)
		anonymousRole := ent.(auth.Role)
		anonymousRole.SetId("Anonymous")
		anonymousRole.SetName("Anonymous")
		storable, ok := ent.(data.Storable)
		if !ok {
			errors.ThrowError(ctx, errors.CORE_ERROR_TYPE_MISMATCH)
		}
		anonRolesReq := ctx.CreateSystemRequest("Save Anonymous Role")
		defer anonRolesReq.CompleteRequest()
		err := lsh.roleDataService.Save(anonRolesReq, storable)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	if !adminExists {
		ent, _ := lsh.roleCreator(ctx, nil)
		adminRole := ent.(auth.Role)
		adminRole.SetPermissions(lsh.allPermissions)
		adminRole.SetId(lsh.adminRole)
		adminRole.SetName(lsh.adminRole)
		storable, ok := ent.(data.Storable)
		if !ok {
			errors.ThrowError(ctx, errors.CORE_ERROR_TYPE_MISMATCH)
		}
		adminRolesReq := ctx.CreateSystemRequest("Create Admin Role")
		defer adminRolesReq.CompleteRequest()
		err := lsh.roleDataService.Save(adminRolesReq, storable)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}
