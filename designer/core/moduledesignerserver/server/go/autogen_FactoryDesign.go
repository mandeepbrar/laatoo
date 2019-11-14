package main


import (
  
  "laatoo/sdk/server/components/data"
)

/*
type FactoryDesign_Ref struct {
  Id    string
}*/

type FactoryDesign struct {
	data.Storable 
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Object	string `json:"Object" bson:"Object" datastore:"Object"`
	Settings	map[string]interface{} `json:"Settings" bson:"Settings" datastore:"Settings"`
}

func (ent *FactoryDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "FactoryDesign",
	}
}

