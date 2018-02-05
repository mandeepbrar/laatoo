package main

import (
  
  "laatoo/sdk/components/data"
)

type Param_Ref struct {
  Id    string
  Name string
}

type Param struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Value	string `json:"Value" bson:"Value" datastore:"Value"`
}

func (ent *Param) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Param",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "Param",
		Cacheable:       false,
	}
}
