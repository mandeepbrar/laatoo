package main


import (
  
  "laatoo/sdk/server/components/data"
)

type FieldDesign_Ref struct {
  Id    string
}

type FieldDesign struct {
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Type	string `json:"Type" bson:"Type" datastore:"Type"`
}

func (ent *FieldDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "FieldDesign",
	}
}

