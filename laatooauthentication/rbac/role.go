package ginauth_rbac

import (
	"laatoo/commonobjects"
)

type Role struct {
	Role        string                  `json:"Role" form:"Role" bson:"Role"`
	Permissions commonobjects.StringSet `json:"Permissions" bson:"Permissions"`
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
