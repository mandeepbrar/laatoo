package modulesbase

import (
  
  "laatoo/sdk/server/components/data"
)

type Entity_Ref struct {
  Id    string
  Name string
}

type Entity struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Fields	[]Field `json:"Fields" bson:"Fields" datastore:"Fields"`
}

func (ent *Entity) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Entity",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "Entity",
		Cacheable:       false,
	}
}
