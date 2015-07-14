package ginauth_rbac

import (
	"github.com/storageutils"
)

type Role struct {
	Role        string                 `json:"Role" form:"Role" bson:"Role"`
	Permissions storageutils.StringSet `json:"Permissions" bson:"Permissions"`
}

func (r *Role) GetId() string {
	return r.Role
}
func (r *Role) SetId(id string) {
	r.Role = id
}
func (r *Role) GetIdField() string {
	return "Role"
}
func CreateRole() interface{} {
	return &Role{}
}
