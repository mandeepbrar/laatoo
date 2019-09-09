package main


import (
  
  "laatoo/sdk/server/components/data"
)

type EntityDesign_Ref struct {
  Id    string
}

type EntityDesign struct {
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Fields	[]FieldDesign `json:"Fields" bson:"Fields" datastore:"Fields"`
}

func (ent *EntityDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "EntityDesign",
	}
}

