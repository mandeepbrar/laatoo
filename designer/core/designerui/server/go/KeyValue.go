package main

import (
  
  "laatoo/sdk/server/components/data"
)

type KeyValue_Ref struct {
  Id    string
  Name string
}

type KeyValue struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Value	string `json:"Value" bson:"Value" datastore:"Value"`
}

func (ent *KeyValue) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "KeyValue",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "KeyValue",
		Cacheable:       false,
	}
}
