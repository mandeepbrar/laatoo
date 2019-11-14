package main


import (
  
	"laatoo/sdk/modules/laatoositeui" 
  "laatoo/sdk/server/components/data"
)

/*
type Deployment_Ref struct {
  Id    string
}*/

type Deployment struct {
	data.Storable 
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Solution	data.StorableRef `json:"Solution" bson:"Solution" datastore:"Solution"`
	Environment	*Environment `json:"Environment" bson:"Environment" datastore:"Environment"`
	Server	*Server `json:"Server" bson:"Server" datastore:"Server"`
	Application	*Application `json:"Application" bson:"Application" datastore:"Application"`
}

func (ent *Deployment) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Deployment",
	}
}

