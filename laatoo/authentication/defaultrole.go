package laatooauthentication

import (
	"laatoocore"
	"laatoosdk/core"
)

type Role struct {
	Role        string   `json:"Role" form:"Role" bson:"Role"`
	Permissions []string `json:"Permissions" bson:"Permissions"`
	CreatedBy   string   `json:"CreatedBy" bson:"CreatedBy"`
	UpdatedBy   string   `json:"UpdatedBy" bson:"UpdatedBy"`
	UpdatedOn   string   `json:"UpdatedOn" bson:"UpdatedOn"`
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
func (ent *Role) PreSave(ctx core.Context) error {
	return nil
}
func (ent *Role) PostSave(ctx core.Context) error {
	return nil
}
func (ent *Role) PostLoad(ctx core.Context) error {
	return nil
}

func init() {
	laatoocore.RegisterObjectProvider(laatoocore.DEFAULT_ROLE, CreateRole)
}

func CreateRole(ctx core.Context, conf map[string]interface{}) (interface{}, error) {
	return &Role{}, nil
}

func (ent *Role) IsNew() bool {
	return ent.CreatedBy == ""
}
func (ent *Role) SetUpdatedOn(val string) {
	ent.UpdatedOn = val
}
func (ent *Role) SetUpdatedBy(val string) {
	ent.UpdatedBy = val
}
func (ent *Role) SetCreatedBy(val string) {
	ent.CreatedBy = val
}
func (ent *Role) GetCreatedBy() string {
	return ent.CreatedBy
}
