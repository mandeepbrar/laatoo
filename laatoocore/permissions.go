package laatoocore

import (
	"fmt"
	"laatoosdk/auth"
	"laatoosdk/core"
	"laatoosdk/log"
	"reflect"
)

/*type PermissionsExchange struct {
	Permissions []string
}

//register the object factory in the global register
func (env *Environment) RegisterPermissions(ctx interface{}, perm []string) {
	env.Permissions.Append(perm)
}

func (env *Environment) ListAllPermissions() []string {
	return env.Permissions.Values()
}
*/
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

func (env *Environment) IsAllowed(ctx core.Context, perm string) bool {
	if perm == "" {
		return true
	}
	bypass := ctx.Get(CONF_SERVICE_AUTHBYPASS)
	if bypass != nil && bypass.(bool) {
		log.Logger.Trace(ctx, "core.permissions", "Registered Role permissions", "perm", perm, "bypass", bypass)
		return true
	}
	log.Logger.Trace(ctx, "core.permissions", "Registered Role permissions", "perm", perm, "bypass", bypass)
	rolesInt := ctx.Get("Roles")
	if rolesInt == nil {
		return false
	}
	roles := rolesInt.([]string)
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
