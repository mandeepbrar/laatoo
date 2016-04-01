package security

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/utils"
)

type LocalSecurityHandler struct {
	allPermissions     []string
	rolePermissionsMap map[string][]string
	//permissions set for the environment
	Permissions utils.StringSet
	//permissions assigned to a role
	RolePermissions map[string]bool
}

func NewLocalSecurityHandler(ctx core.ServerContext) core.SecurityHandler {
	lsh := &LocalSecurityHandler{}
	lsh.allPermissions = make([]string, 0, 10)
	lsh.rolePermissionsMap = make(map[string][]string, 10)
	//construct permissions set
	lsh.Permissions = utils.NewStringSet([]string{})
	//map containing roles and permissions
	lsh.RolePermissions = make(map[string]bool)
	return lsh
}

func (lsh *LocalSecurityHandler) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}
func (lsh *LocalSecurityHandler) HasPermission(ctx core.RequestContext, permission string) bool {
	return false
}
func (lsh *LocalSecurityHandler) GetRolePermissions(role []string) ([]string, bool) {
	perms := make([]string, 0, 10)
	/*permissions := utils.NewStringSet([]string{})
	adminRole := ctx.GetServerVariable(core.ADMINROLE)
	for _, rolename := range roles {
		if rolename == adminRole {
			usr.SetPermissions(*allpermissions)
			return true
		}
		rolepermissions, ok := rolesMap[rolename]
		if ok {
			permissions.Append(rolepermissions)
		}
	}*/

	return perms, false
}
