package main


import (
  
  "laatoo/sdk/server/components/data"
)

/*
type RuleDesign_Ref struct {
  Id    string
}*/

type RuleDesign struct {
	data.Storable 
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Trigger	string `json:"Trigger" bson:"Trigger" datastore:"Trigger"`
	MessageType	string `json:"MessageType" bson:"MessageType" datastore:"MessageType"`
	Rule	string `json:"Rule" bson:"Rule" datastore:"Rule"`
	Settings	map[string]interface{} `json:"Settings" bson:"Settings" datastore:"Settings"`
}

func (ent *RuleDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "RuleDesign",
	}
}

