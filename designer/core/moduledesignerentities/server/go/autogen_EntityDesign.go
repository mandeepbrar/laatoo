package main


import (
  
  "laatoo/sdk/server/components/data"
)

type EntityDesign_Ref struct {
  Id    string
}

type EntityDesign struct {
	*data.SerializableBase `initialize:"SerializableBase"`
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Imports	[]string `json:"Imports" bson:"Imports" datastore:"Imports"`
	TitleField	string `json:"TitleField" bson:"TitleField" datastore:"TitleField"`
	Collection	string `json:"Collection" bson:"Collection" datastore:"Collection"`
	SkipStorage	bool `json:"SkipStorage" bson:"SkipStorage" datastore:"SkipStorage"`
	Fields	[]FieldDesign `json:"Fields" bson:"Fields" datastore:"Fields"`
}

func (ent *EntityDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "EntityDesign",
	}
}

