package main

import (
  
  "laatoo/sdk/components/data"
)

type Service_Ref struct {
  Id    string
  Name string
}

type Service struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Factory	string `json:"Factory" bson:"Factory" datastore:"Factory"`
	ServiceMethod	string `json:"ServiceMethod" bson:"ServiceMethod" datastore:"ServiceMethod"`
	LoggingLevel	string `json:"LoggingLevel" bson:"LoggingLevel" datastore:"LoggingLevel"`
	LoggingFormat	string `json:"LoggingFormat" bson:"LoggingFormat" datastore:"LoggingFormat"`
	Params	[]Param `json:"Params" bson:"Params" datastore: "Params"`
}

func (ent *Service) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Service",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "Service",
		Cacheable:       false,
	}
}
