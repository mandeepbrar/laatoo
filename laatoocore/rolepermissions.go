package laatoocore

import (
	"fmt"
	"laatoosdk/auth"
	"laatoosdk/core"
	"laatoosdk/log"
	"reflect"
)

//register the roles and permissions
func (env *Environment) RegisterRoles(ctx *Context, rolesInt interface{}) {
	if rolesInt != nil {
		arr := reflect.ValueOf(rolesInt).Elem()
		length := arr.Len()
		for i := 0; i < length; i++ {
			role := arr.Index(i).Addr().Interface().(auth.Role)
			env.RegisterRolePermissions(ctx, role)
		}
	}
}

func (env *Environment) RegisterRolePermissions(ctx *Context, role auth.Role) {
	permissions := role.GetPermissions()
	for _, perm := range permissions {
		key := fmt.Sprintf("%s#%s", role.GetId(), perm)
		env.RolePermissions[key] = true
	}
	log.Logger.Trace(ctx, "core.permissions", "Registered Role permissions", "Role Permissions", env.RolePermissions)
}

func (env *Environment) HasPermission(ctx core.Context, perm string) bool {
	if perm == "" {
		return true
	}
	bypass := ctx.Get(CONF_SERVICE_AUTHBYPASS)
	if bypass != nil && bypass.(bool) {
		log.Logger.Trace(ctx, "core.permissions", "Bypassed permission", "perm", perm, "bypass", bypass)
		return true
	}
	rolesInt := ctx.Get("Roles")
	if rolesInt == nil {
		return false
	}
	roles := rolesInt.([]string)
	log.Logger.Trace(ctx, "core.permissions", "Checking roles for permission", "perm", perm, "bypass", bypass, "roles", roles)
	for _, role := range roles {
		if role == env.AdminRole {
			return true
		}
		key := fmt.Sprintf("%s#%s", role, perm)
		val, ok := env.RolePermissions[key]
		if ok {
			return val
		}
	}
	return false
}
