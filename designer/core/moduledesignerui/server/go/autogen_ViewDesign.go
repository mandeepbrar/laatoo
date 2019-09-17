package main


import (
  
  "laatoo/sdk/server/components/data"
)

type ViewDesign_Ref struct {
  Id    string
}

type ViewDesign struct {
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Service	string `json:"Service" bson:"Service" datastore:"Service"`
	PostArgs	map[string]interface{} `json:"PostArgs" bson:"PostArgs" datastore:"PostArgs"`
	URLParams	map[string]interface{} `json:"URLParams" bson:"URLParams" datastore:"URLParams"`
	ItemType	ViewItemDesign `json:"ItemType" bson:"ItemType" datastore:"ItemType"`
}

func (ent *ViewDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ViewDesign",
	}
}

