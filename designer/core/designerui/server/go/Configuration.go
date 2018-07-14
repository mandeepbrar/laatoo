package main

import (
  
  "laatoo/sdk/server/components/data"
)

type Configuration_Ref struct {
  Id    string
  Name string
}

type Configuration struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Values	[]KeyValue `json:"Values" bson:"Values" datastore: "Values"`
}

func (ent *Configuration) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Configuration",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "Configuration",
		Cacheable:       false,
	}
}
