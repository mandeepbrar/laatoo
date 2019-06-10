package main

import (
  
  "laatoo/sdk/server/components/data"
)

type Field_Ref struct {
  Id    string
}

type Field struct {
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Type	string `json:"Type" bson:"Type" datastore:"Type"`
}

func (ent *Field) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Field",
	}
}
