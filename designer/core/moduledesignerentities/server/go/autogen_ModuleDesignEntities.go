package main

import (
  
  "laatoo/sdk/server/components/data"
)

type ModuleDesignEntities_Ref struct {
  Id    string
  Title string
}

type ModuleDesignEntities struct {
	data.SoftDeleteAuditableMT `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Entities	[]EntityDesign `json:"Entities" bson:"Entities" datastore:"Entities"`
}

func (ent *ModuleDesignEntities) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Title",
		Type:            "ModuleDesignEntities",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     true,
		Collection:      "ModuleDesignEntities",
		Cacheable:       false,
	}
}
