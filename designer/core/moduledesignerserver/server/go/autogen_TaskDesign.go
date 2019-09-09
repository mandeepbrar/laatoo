package main


import (
  
  "laatoo/sdk/server/components/data"
)

type TaskDesign_Ref struct {
  Id    string
}

type TaskDesign struct {
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Receiver	string `json:"Receiver" bson:"Receiver" datastore:"Receiver"`
	Processor	string `json:"Processor" bson:"Processor" datastore:"Processor"`
}

func (ent *TaskDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "TaskDesign",
	}
}

