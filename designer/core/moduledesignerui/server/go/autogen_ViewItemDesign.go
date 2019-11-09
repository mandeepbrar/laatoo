package main


import (
  
  "laatoo/sdk/server/components/data"
)

type ViewItemDesign_Ref struct {
  Id    string
}

type ViewItemDesign struct {
	data.Storable 
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Type	string `json:"Type" bson:"Type" datastore:"Type"`
	EntityName	string `json:"EntityName" bson:"EntityName" datastore:"EntityName"`
	EntityDisplay	string `json:"EntityDisplay" bson:"EntityDisplay" datastore:"EntityDisplay"`
}

func (ent *ViewItemDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ViewItemDesign",
	}
}

