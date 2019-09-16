package main

import (
  
	"laatoo/sdk/modules/modulesbase" 
  "laatoo/sdk/server/components/data"
)

type ModuleDesignGeneral_Ref struct {
  Id    string
  Title string
}

type ModuleDesignGeneral struct {
	data.SoftDeleteAuditableMT `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Params	[]modulesbase.Param `json:"Params" bson:"Params" datastore:"Params"`
	Dependencies	[]modulesbase.Dependency `json:"Dependencies" bson:"Dependencies" datastore:"Dependencies"`
}

func (ent *ModuleDesignGeneral) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Title",
		Type:            "ModuleDesignGeneral",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     true,
		Collection:      "ModuleDesignGeneral",
		Cacheable:       false,
	}
}
