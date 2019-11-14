package main


import (
  
  "laatoo/sdk/server/components/data"
)

/*
type FieldWidgetDesign_Ref struct {
  Id    string
}*/

type FieldWidgetDesign struct {
	data.Storable 
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Module	string `json:"Module" bson:"Module" datastore:"Module"`
	Props	map[string]interface{} `json:"Props" bson:"Props" datastore:"Props"`
}

func (ent *FieldWidgetDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "FieldWidgetDesign",
	}
}

