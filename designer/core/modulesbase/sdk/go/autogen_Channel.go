package modulesbase

import (
  
  "laatoo/sdk/server/components/data"
)

type Channel_Ref struct {
  Id    string
  Name string
}

type Channel struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Factory	string `json:"Factory" bson:"Factory" datastore:"Factory"`
	ServiceMethod	string `json:"ServiceMethod" bson:"ServiceMethod" datastore:"ServiceMethod"`
	Settings	map[string]interface{} `json:"Settings" bson:"Settings" datastore:"Settings"`
}

func (ent *Channel) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Channel",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "Channel",
		Cacheable:       false,
	}
}
