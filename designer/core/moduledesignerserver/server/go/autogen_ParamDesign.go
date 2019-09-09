package main


import (
  
  "laatoo/sdk/server/components/data"
)

type ParamDesign_Ref struct {
  Id    string
}

type ParamDesign struct {
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Type	string `json:"Type" bson:"Type" datastore:"Type"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Required	bool `json:"Required" bson:"Required" datastore:"Required"`
	Default	string `json:"Default" bson:"Default" datastore:"Default"`
}

func (ent *ParamDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ParamDesign",
	}
}

