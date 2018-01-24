package main

import (
  
  "laatoo/sdk/components/data"
)

type Rule_Ref struct {
  Id    string
  Name string
}

type Rule struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Trigger	string `json:"Trigger" bson:"Trigger" datastore:"Trigger"`
	MessageType	string `json:"MessageType" bson:"MessageType" datastore:"MessageType"`
	Rule	string `json:"Rule" bson:"Rule" datastore:"Rule"`
	LoggingLevel	string `json:"LoggingLevel" bson:"LoggingLevel" datastore:"LoggingLevel"`
	LoggingFormat	string `json:"LoggingFormat" bson:"LoggingFormat" datastore:"LoggingFormat"`
	Params	[]Param `json:"Params" bson:"Params" datastore: "Params"`
}

func (ent *Rule) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Rule",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "Rule",
		Cacheable:       false,
	}
}
