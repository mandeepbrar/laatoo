package laatoocore

import (
	"fmt"
	"github.com/labstack/echo"
	"laatoosdk/auth"
	"laatoosdk/utils"
)

var (
	SystemUser      = ""
	SystemRole      = ""
	Permissions     = utils.NewStringSet([]string{})
	RolePermissions = make(map[string]bool)
)

//register the object factory in the global register
func RegisterPermissions(perm []string) {
	Permissions.Append(perm)
}

func ListAllPermissions() []string {
	return Permissions.Values()
}

func RegisterRolePermissions(role auth.Role) {
	permissions := role.GetPermissions()
	for _, perm := range permissions {
		key := fmt.Sprintf("%s#%s", role.GetId(), perm)
		RolePermissions[key] = true
	}
}

func IsAllowed(ctx *echo.Context, perm string) bool {
	if perm == "" {
		return true
	}
	rolesInt := ctx.Get("Roles")
	if rolesInt == nil {
		return false
	}
	roles := rolesInt.([]string)
	for _, role := range roles {
		key := fmt.Sprintf("%s#%s", role, perm)
		val, ok := RolePermissions[key]
		if ok {
			return val
		}
	}
	return false
}
