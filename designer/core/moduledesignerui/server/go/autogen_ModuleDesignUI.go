package main

import (
  
	"laatoo/sdk/modules/modulesbase" 
  "laatoo/sdk/server/components/data"
)

type ModuleDesignUI_Ref struct {
  Id    string
  Title string
}

type ModuleDesignUI struct {
	data.SoftDeleteAuditableMT `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Externals	map[string]interface{} `json:"Externals" bson:"Externals" datastore:"Externals"`
	UIDependencies	[]modulesbase.Dependency `json:"UIDependencies" bson:"UIDependencies" datastore:"UIDependencies"`
}

func (ent *ModuleDesignUI) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Title",
		Type:            "ModuleDesignUI",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     true,
		Collection:      "ModuleDesignUI",
		Cacheable:       false,
	}
}
