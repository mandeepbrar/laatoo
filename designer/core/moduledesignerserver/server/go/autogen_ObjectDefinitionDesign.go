package main


import (
  
	"laatoo/sdk/modules/moduledesignerbase" 
  "laatoo/sdk/server/components/data"
)

type ObjectDefinitionDesign_Ref struct {
  Id    string
}

type ObjectDefinitionDesign struct {
	*data.SerializableBase `initialize:"SerializableBase"`
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Type	string `json:"Type" bson:"Type" datastore:"Type"`
	RequestType	string `json:"RequestType" bson:"RequestType" datastore:"RequestType"`
	RequestParams	[]moduledesignerbase.ParamDesign `json:"RequestParams" bson:"RequestParams" datastore:"RequestParams"`
	Configurations	[]moduledesignerbase.ParamDesign `json:"Configurations" bson:"Configurations" datastore:"Configurations"`
}

func (ent *ObjectDefinitionDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ObjectDefinitionDesign",
	}
}

