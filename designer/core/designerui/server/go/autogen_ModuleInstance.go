package main

import (
  
  "laatoo/sdk/server/components/data"
)

type ModuleInstance_Ref struct {
  Id    string
  Name string
}

type ModuleInstance struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Module	Module `json:"Module" bson:"Module" datastore: "Module"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Settings	Configuration `json:"Settings" bson:"Settings" datastore: "Settings"`
}

func (ent *ModuleInstance) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ModuleInstance",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "ModuleInstance",
		Cacheable:       false,
	}
}
