package main

import (
  
  "laatoo/sdk/components/data"
)

type ModuleInstance_Ref struct {
  Id    string
  Name string
}

type ModuleInstance struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	ModuleRef	*Module `json:"ModuleRef" bson:"ModuleRef" datastore:"ModuleRef"`
	Module	string `json:"Module" bson:"Module" datastore: "Module"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Settings	[]Param `json:"Settings" bson:"Settings" datastore: "Settings"`
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
		Collection:      "ModuleInstance",
		Cacheable:       false,
	}
}
