package main

import (
  
  "laatoo/sdk/server/components/data"
)

type Task_Ref struct {
  Id    string
  Name string
}

type Task struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Receiver	string `json:"Receiver" bson:"Receiver" datastore:"Receiver"`
	Processor	string `json:"Processor" bson:"Processor" datastore:"Processor"`
}

func (ent *Task) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Task",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "Task",
		Cacheable:       false,
	}
}
