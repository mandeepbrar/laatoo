package security

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/server/objects"
)

var (
	rc = &data.StorableConfig{
		IdField:         "Id",
		Type:            config.DEFAULT_ROLE,
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "Role",
		Cacheable:       false,
	}
)

func init() {
	objects.Register(config.DEFAULT_ROLE, Role{})
}

type Role struct {
	data.SoftDeleteAuditable `bson:",inline"`
	Role                     string   `json:"Role" form:"Role" bson:"Role"`
	Permissions              []string `json:"Permissions" bson:"Permissions"`
	Realm                    string   `json:"Realm" bson:"Realm"`
}

func (r *Role) Config() *data.StorableConfig {
	return rc
}

func (r *Role) GetPermissions() []string {
	return r.Permissions
}

func (r *Role) SetPermissions(permissions []string) {
	r.Permissions = permissions
}

func (ent *Role) GetName() string {
	return ent.Role
}
func (ent *Role) SetName(val string) {
	ent.Role = val
}
func (ent *Role) GetRealm() string {
	return ent.Realm
}
func (ent *Role) SetRealm(val string) {
	ent.Realm = val
}