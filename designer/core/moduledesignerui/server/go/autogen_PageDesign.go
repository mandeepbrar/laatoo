package main


import (
  
	"laatoo/sdk/modules/moduledesignerbase" 
  "laatoo/sdk/server/components/data"
)

type PageDesign_Ref struct {
  Id    string
}

type PageDesign struct {
	data.Storable 
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Route	string `json:"Route" bson:"Route" datastore:"Route"`
	Block	string `json:"Block" bson:"Block" datastore:"Block"`
	Params	[]moduledesignerbase.ParamDesign `json:"Params" bson:"Params" datastore:"Params"`
}

func (ent *PageDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "PageDesign",
	}
}

