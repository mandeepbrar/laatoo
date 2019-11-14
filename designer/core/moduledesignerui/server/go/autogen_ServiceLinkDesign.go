package main


import (
  
  "laatoo/sdk/server/components/data"
)

/*
type ServiceLinkDesign_Ref struct {
  Id    string
}*/

type ServiceLinkDesign struct {
	data.Storable 
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	URL	string `json:"URL" bson:"URL" datastore:"URL"`
	Method	string `json:"Method" bson:"Method" datastore:"Method"`
}

func (ent *ServiceLinkDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ServiceLinkDesign",
	}
}

