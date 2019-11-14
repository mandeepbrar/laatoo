package main

import (
  
  "laatoo/sdk/server/components/data"
)

/*type ModuleDesignEntities_Ref struct {
  Id    string
  Title string
}*/

type ModuleDesignEntities struct {
	data.Storable `laatoo:"auditable, softdelete, multitenant"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Entities	[]EntityDesign `json:"Entities" bson:"Entities" datastore:"Entities"`
}

func (ent *ModuleDesignEntities) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Title",
		Type:            "ModuleDesignEntities",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     true,
		Collection:      "ModuleDesignEntities",
		Cacheable:       false,
	}
}
