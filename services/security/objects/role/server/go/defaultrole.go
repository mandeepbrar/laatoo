package main

import (
	"laatoo/sdk/server/components/data"
)

var (
	rc = &data.StorableConfig{
		PreSave:    false,
		PostSave:   false,
		PostLoad:   false,
		Auditable:  true,
		Collection: "Role",
		Cacheable:  false,
	}
)

type Role struct {
	data.Storable `bson:"-" json:",inline" laatoo:"softdelete, auditable"`
	Role          string   `json:"Role" form:"Role" bson:"Role"`
	Permissions   []string `json:"Permissions" bson:"Permissions"`
	Realm         string   `json:"Realm" bson:"Realm"`
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
