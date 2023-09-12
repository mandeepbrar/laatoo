package rbac

import (
	"laatoo/sdk/server/auth"
	"laatoo/sdk/server/components/data"
)

type RbacUser interface {
	auth.User
	GetRoles() ([]data.StorableRef, error)
	SetRoles([]data.StorableRef) error
}

/*GetId() string
SetId(string)
GetUsernameField() string
GetUserName() string
LoadClaims(map[string]interface{})
PopulateClaims(map[string]interface{})
GetRealm() string
GetTenant() string
*/
