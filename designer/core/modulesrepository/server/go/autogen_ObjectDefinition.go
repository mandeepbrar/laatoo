package main

import (
  
  "laatoo/sdk/server/components/data"
)

type ObjectDefinition_Ref struct {
  Id    string
  Name string
}

type ObjectDefinition struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Type	string `json:"Type" bson:"Type" datastore:"Type"`
	Request	[]RequestDefinition `json:"Request" bson:"Request" datastore:"Request"`
	Configurations	[]ConfigurationDefinition `json:"Configurations" bson:"Configurations" datastore:"Configurations"`
}

func (ent *ObjectDefinition) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ObjectDefinition",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "ObjectDefinition",
		Cacheable:       false,
	}
}


