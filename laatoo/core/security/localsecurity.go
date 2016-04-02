package security

import (
	"fmt"
	"laatoo/core/registry"
	"laatoo/sdk/auth"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/data"
	"laatoo/sdk/errors"
	"laatoo/sdk/utils"
)

const (
	CONF_SECURIY_ROLEDATASERVICE = "role_data_svc"
	CONF_SECURIY_PERMISSIONS     = "permissions"
)

type LocalSecurityHandler struct {
	conf config.Config
	//userCreator core.ObjectCreator
	roleCreator core.ObjectCreator
	adminRole   string
	//data service to use for users
	//UserDataService    data.DataService
	roleDataService data.DataService
	rolesMap        map[string]auth.Role
	allPermissions  []string
	//permissions assigned to a role
	rolePermissions map[string]bool
}

func NewLocalSecurityHandler(ctx core.ServerContext, conf config.Config) core.SecurityHandler {
	lsh := &LocalSecurityHandler{conf: conf}
	lsh.rolesMap = make(map[string]auth.Role, 10)
	//map containing roles and permissions
	lsh.rolePermissions = make(map[string]bool, 50)

	lsh.adminRole = ctx.GetServerVariable(core.ADMINROLE).(string)
	return lsh
}

func (lsh *LocalSecurityHandler) Initialize(ctx core.ServerContext, conf config.Config) error {
	if conf == nil {
		conf = lsh.conf
	}
	permissions, ok := conf.GetStringArray(CONF_SECURIY_PERMISSIONS)
	if ok {
		lsh.allPermissions = permissions
	} else {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_SECURIY_PERMISSIONS)
	}
	roleobject := ctx.GetServerVariable(core.ROLE)
	roleCreator, err := registry.GetObjectCreator(ctx, roleobject.(string))
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	lsh.roleCreator = roleCreator

	roleDataSvcName, ok := conf.GetString(CONF_SECURIY_ROLEDATASERVICE)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", CONF_SECURIY_ROLEDATASERVICE)
	}
	roleService, err := ctx.GetService(roleDataSvcName)
	if err != nil {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", CONF_SECURIY_ROLEDATASERVICE)
	}
	roleDataService, ok := roleService.(data.DataService)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", CONF_SECURIY_ROLEDATASERVICE)
	}
	lsh.roleDataService = roleDataService
	return lsh.loadRoles(ctx)
}

func (lsh *LocalSecurityHandler) GetRolePermissions(roles []string) ([]string, bool) {
	permissions := utils.NewStringSet([]string{})
	for _, rolename := range roles {
		role, ok := lsh.rolesMap[rolename]
		if ok {
			permissions.Append(role.GetPermissions())
		}
	}
	return permissions.Values(), false
}

func (lsh *LocalSecurityHandler) HasPermission(ctx core.RequestContext, perm string) bool {
	return hasPermission(ctx, perm, lsh.rolePermissions)
}

func (lsh *LocalSecurityHandler) loadRoles(ctx core.ServerContext) error {
	roles, _, _, err := lsh.roleDataService.GetList(ctx.CreateNewRequest("LoadRoles"), -1, -1, "", "")
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

func (lsh *LocalSecurityHandler) createInitRoles(ctx core.ServerContext, anonExists bool, adminExists bool) error {

	if !anonExists {
		ent, _ := lsh.roleCreator(ctx, nil)
		anonymousRole := ent.(auth.Role)
		anonymousRole.SetId("Anonymous")
		storable, ok := ent.(data.Storable)
		if !ok {
			errors.ThrowError(ctx, errors.CORE_ERROR_TYPE_MISMATCH)
		}
		err := lsh.roleDataService.Save(ctx.CreateNewRequest("Save Anonymous Role"), storable)
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
		err := lsh.roleDataService.Save(ctx.CreateNewRequest("Create Admin Role"), storable)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}
