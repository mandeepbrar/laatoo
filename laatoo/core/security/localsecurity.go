package security

import (
	"fmt"
	"laatoo/sdk/auth"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/data"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"
	"laatoo/sdk/utils"
)

const (
	CONF_SECURIY_ROLEDATASERVICE = "role_data_svc"
	CONF_SECURIY_PERMISSIONS     = "permissions"
)

type localSecurityHandler struct {
	userCreator core.ObjectCreator
	roleCreator core.ObjectCreator
	adminRole   string
	jwtsecret   string
	authheader  string
	//data service to use for users
	//UserDataService    data.DataService
	roleDataService data.DataService
	parent          server.SecurityHandler
	rolesMap        map[string]auth.Role
	allPermissions  []string
	//permissions assigned to a role
	rolePermissions map[string]bool
}

func NewLocalSecurityHandler(ctx core.ServerContext, parent server.SecurityHandler, conf config.Config, role string, adminrole string, user string, jwtsecret string, authheader string) (SecurityPlugin, error) {
	lsh := &localSecurityHandler{adminRole: adminrole, parent: parent, jwtsecret: jwtsecret, authheader: authheader}
	lsh.rolesMap = make(map[string]auth.Role, 10)
	//map containing roles and permissions
	lsh.rolePermissions = make(map[string]bool, 50)

	permissions, ok := conf.GetStringArray(CONF_SECURIY_PERMISSIONS)
	if ok {
		lsh.allPermissions = permissions
	} else {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_SECURIY_PERMISSIONS)
	}

	roleCreator, err := ctx.GetObjectCreator(role)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	lsh.roleCreator = roleCreator

	userCreator, err := ctx.GetObjectCreator(user)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	lsh.userCreator = userCreator

	roleDataSvcName, ok := conf.GetString(CONF_SECURIY_ROLEDATASERVICE)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", CONF_SECURIY_ROLEDATASERVICE)
	}
	roleService, err := ctx.GetService(roleDataSvcName)
	if err != nil {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", CONF_SECURIY_ROLEDATASERVICE)
	}
	roleDataService, ok := roleService.(data.DataService)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", CONF_SECURIY_ROLEDATASERVICE)
	}
	lsh.roleDataService = roleDataService
	err = lsh.loadRoles(ctx)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	return lsh, nil
}

func (lsh *localSecurityHandler) GetRolePermissions(ctx core.RequestContext, roles []string) ([]string, bool) {
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
func (lsh *localSecurityHandler) GetUser(ctx core.RequestContext) (auth.User, bool, error) {
	return getUserFromToken(ctx, lsh.userCreator, lsh.authheader, lsh.jwtsecret)
}
func (lsh *localSecurityHandler) loadRoles(ctx core.ServerContext) error {
	loadRolesReq := lsh.parent.CreateSystemRequest(ctx, "LoadRoles")
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
		if id == "Anonymous" {
			anonExists = true
		}

		if id == lsh.adminRole {
			adminExists = true
		}
		lsh.rolesMap[id] = role
		permissions := role.GetPermissions()
		for _, perm := range permissions {
			key := fmt.Sprintf("%s#%s", id, perm)
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
		storable, ok := ent.(data.Storable)
		if !ok {
			errors.ThrowError(ctx, errors.CORE_ERROR_TYPE_MISMATCH)
		}
		anonRolesReq := lsh.parent.CreateSystemRequest(ctx, "Save Anonymous Role")
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
		storable, ok := ent.(data.Storable)
		if !ok {
			errors.ThrowError(ctx, errors.CORE_ERROR_TYPE_MISMATCH)
		}
		adminRolesReq := lsh.parent.CreateSystemRequest(ctx, "Create Admin Role")
		defer adminRolesReq.CompleteRequest()
		err := lsh.roleDataService.Save(adminRolesReq, storable)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}
