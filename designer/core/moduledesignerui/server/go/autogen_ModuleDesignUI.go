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
	*data.SoftDeleteAuditableMT `initialize:"SoftDeleteAuditableMT"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Externals	map[string]interface{} `json:"Externals" bson:"Externals" datastore:"Externals"`
	UIDependencies	[]modulesbase.Dependency `json:"UIDependencies" bson:"UIDependencies" datastore:"UIDependencies"`
	ServiceLinks	[]ServiceLinkDesign `json:"ServiceLinks" bson:"ServiceLinks" datastore:"ServiceLinks"`
	Actions	[]ActionDesign `json:"Actions" bson:"Actions" datastore:"Actions"`
	Views	[]ViewDesign `json:"Views" bson:"Views" datastore:"Views"`
	Pages	[]PageDesign `json:"Pages" bson:"Pages" datastore:"Pages"`
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
