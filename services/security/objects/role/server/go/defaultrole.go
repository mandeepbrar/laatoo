package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
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

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: config.DEFAULT_ROLE, Object: Role{}}}
}

type Role struct {
	*data.SoftDeleteAuditable `json:",inline" initialize:"SoftDeleteAuditable"`
	Role                      string   `json:"Role" form:"Role" bson:"Role"`
	Permissions               []string `json:"Permissions" bson:"Permissions"`
	Realm                     string   `json:"Realm" bson:"Realm"`
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
