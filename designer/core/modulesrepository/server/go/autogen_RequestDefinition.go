package main

import (
  
  "laatoo/sdk/server/components/data"
)

type RequestDefinition_Ref struct {
  Id    string
  Title string
}

type RequestDefinition struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Type	string `json:"Type" bson:"Type" datastore:"Type"`
}

func (ent *RequestDefinition) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Title",
		Type:            "RequestDefinition",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "RequestDefinition",
		Cacheable:       false,
	}
}


