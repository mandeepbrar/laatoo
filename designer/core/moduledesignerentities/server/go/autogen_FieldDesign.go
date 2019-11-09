package main


import (
  
  "laatoo/sdk/server/components/data"
)

type FieldDesign_Ref struct {
  Id    string
}

type FieldDesign struct {
	data.Storable 
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Type	string `json:"Type" bson:"Type" datastore:"Type"`
	List	bool `json:"List" bson:"List" datastore:"List"`
	Reference	bool `json:"Reference" bson:"Reference" datastore:"Reference"`
	Imports	[]string `json:"Imports" bson:"Imports" datastore:"Imports"`
	Titlefield	string `json:"Titlefield" bson:"Titlefield" datastore:"Titlefield"`
	Collection	string `json:"Collection" bson:"Collection" datastore:"Collection"`
	Widget	FieldWidgetDesign `json:"Widget" bson:"Widget" datastore:"Widget"`
}

func (ent *FieldDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "FieldDesign",
	}
}

