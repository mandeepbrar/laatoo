package main

import (
  
  "laatoo/sdk/server/components/data"
)

type ModuleParam_Ref struct {
  Id    string
  Name string
}

type ModuleParam struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Type	string `json:"Type" bson:"Type" datastore:"Type"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Required	bool `json:"Required" bson:"Required" datastore:"Required"`
	Default	string `json:"Default" bson:"Default" datastore:"Default"`
}

func (ent *ModuleParam) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ModuleParam",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "ModuleParam",
		Cacheable:       false,
	}
}


