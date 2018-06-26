package main

import (
  
  "laatoo/sdk/server/components/data"
)

type Factory_Ref struct {
  Id    string
  Name string
}

type Factory struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Object	string `json:"Object" bson:"Object" datastore:"Object"`
}

func (ent *Factory) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Factory",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "Factory",
		Cacheable:       false,
	}
}
