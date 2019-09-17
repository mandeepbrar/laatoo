package main


import (
  
  "laatoo/sdk/server/components/data"
)

type ActionDesign_Ref struct {
  Id    string
}

type ActionDesign struct {
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Type	string `json:"Type" bson:"Type" datastore:"Type"`
	BlockID	string `json:"BlockID" bson:"BlockID" datastore:"BlockID"`
	URL	string `json:"URL" bson:"URL" datastore:"URL"`
}

func (ent *ActionDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ActionDesign",
	}
}

