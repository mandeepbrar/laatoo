package main

import (
  
  "laatoo/sdk/server/components/data"
)

type Configuration_Ref struct {
  Id    string
}

type Configuration struct {
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Values	[]KeyValue `json:"Values" bson:"Values" datastore: "Values"`
}

func (ent *Configuration) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Configuration",
	}
}
