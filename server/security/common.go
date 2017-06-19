package security

import (
	"fmt"
	"laatoo/sdk/auth"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
)

func hasPermission(ctx core.RequestContext, perm string, rolePermissions map[string]bool) bool {
	if perm == "" {
		return true
	}
	/*bypass, ok := ctx.Get(CONF_SERVICE_AUTHBYPASS)
	if ok && bypass.(bool) {
		log.Logger.Trace(ctx, "Bypassed permission", "perm", perm, "bypass", bypass)
		return true
	}*/
	if ctx.IsAdmin() {
		return true
	}
	usr := ctx.GetUser().(auth.RbacUser)
	roles, _ := usr.GetRoles()
	log.Logger.Trace(ctx, "Checking roles for permission", "perm", perm, "roles", roles)
	for _, role := range roles {
		key := fmt.Sprintf("%s#%s", role, perm)
		val, ok := rolePermissions[key]
		if ok {
			return val
		}
	}
	return false
}
