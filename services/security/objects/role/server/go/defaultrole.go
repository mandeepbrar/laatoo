package main

import (
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
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

func (ent *Role) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	var err error

	if _, err = rdr.ReadString(c, cdc, "Role", &ent.Role); err != nil {
		return err
	}

	if _, err = rdr.ReadString(c, cdc, "Realm", &ent.Realm); err != nil {
		return err
	}

	if _, err = rdr.ReadArray(c, cdc, "Permissions", &ent.Permissions); err != nil {
		return err
	}

	return ent.Storable.ReadAll(c, cdc, rdr)
}

func (ent *Role) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	var err error

	if err = wtr.WriteString(c, cdc, "Role", &ent.Role); err != nil {
		return err
	}

	if err = wtr.WriteArray(c, cdc, "Permissions", &ent.Permissions); err != nil {
		return err
	}

	if err = wtr.WriteString(c, cdc, "Realm", &ent.Realm); err != nil {
		return err
	}

	return ent.Storable.WriteAll(c, cdc, wtr)
}
