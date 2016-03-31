package security

/*
import (
	"fmt"
	"laatoosdk/auth"
	"laatoosdk/core"
	"laatoosdk/log"
	"reflect"
)

//register the roles and permissions
func (env *Environment) RegisterRoles(ctx *serverContext, rolesInt interface{}) {
	if rolesInt != nil {
		arr := reflect.ValueOf(rolesInt).Elem()
		length := arr.Len()
		for i := 0; i < length; i++ {
			role := arr.Index(i).Addr().Interface().(auth.Role)
			env.RegisterRolePermissions(ctx, role)
		}
	}
}

func (env *Environment) RegisterRolePermissions(ctx *serverContext, role auth.Role) {
	permissions := role.GetPermissions()
	for _, perm := range permissions {
		key := fmt.Sprintf("%s#%s", role.GetId(), perm)
		env.RolePermissions[key] = true
	}
	log.Logger.Trace(ctx, "Registered Role permissions", "Role Permissions", env.RolePermissions)
}

func (env *Environment) HasPermission(ctx core.RequestContext, perm string) bool {
	if perm == "" {
		return true
	}
	bypass, ok := ctx.Get(CONF_SERVICE_AUTHBYPASS)
	if ok && bypass.(bool) {
		log.Logger.Trace(ctx, "Bypassed permission", "perm", perm, "bypass", bypass)
		return true
	}
	rolesInt, ok := ctx.Get("Roles")
	if !ok {
		return false
	}
	roles := rolesInt.([]string)
	log.Logger.Trace(ctx, "Checking roles for permission", "perm", perm, "bypass", bypass, "roles", roles)
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
*/
