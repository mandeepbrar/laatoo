package laatooauthentication

import (
	"laatoocore"
)

type Role struct {
	Role        string   `json:"Role" form:"Role" bson:"Role"`
	Permissions []string `json:"Permissions" bson:"Permissions"`
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
func (r *Role) GetPermissions() []string {
	return r.Permissions
}

func (r *Role) SetPermissions(permissions []string) {
	r.Permissions = permissions
}
func (ent *Role) PreSave() error {
	return nil
}
func (ent *Role) PostLoad() error {
	return nil
}

func init() {
	laatoocore.RegisterObjectProvider(laatoocore.DEFAULT_ROLE, CreateRole)
}

func CreateRole(ctx interface{}, conf map[string]interface{}) (interface{}, error) {
	return &Role{}, nil
}
