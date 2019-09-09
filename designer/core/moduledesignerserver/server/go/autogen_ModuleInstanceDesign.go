package main


import (
  
  "laatoo/sdk/server/components/data"
)

type ModuleInstanceDesign_Ref struct {
  Id    string
}

type ModuleInstanceDesign struct {
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Module	data.StorableRef `json:"Module" bson:"Module" datastore:"Module"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Settings	map[string]interface{} `json:"Settings" bson:"Settings" datastore:"Settings"`
}

func (ent *ModuleInstanceDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ModuleInstanceDesign",
	}
}

