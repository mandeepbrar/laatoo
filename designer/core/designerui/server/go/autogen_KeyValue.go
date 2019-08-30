package main

import (
  
  "laatoo/sdk/server/components/data"
)

type KeyValue_Ref struct {
  Id    string
}

type KeyValue struct {
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Value	string `json:"Value" bson:"Value" datastore:"Value"`
}

func (ent *KeyValue) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "KeyValue",
	}
}

