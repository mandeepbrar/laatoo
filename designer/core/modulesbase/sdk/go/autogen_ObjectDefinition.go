package modulesbase

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
	RequestType	string `json:"RequestType" bson:"RequestType" datastore:"RequestType"`
	RequestParams	[]Param `json:"RequestParams" bson:"RequestParams" datastore:"RequestParams"`
	Configurations	[]Param `json:"Configurations" bson:"Configurations" datastore:"Configurations"`
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
