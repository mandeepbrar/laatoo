package main


import (
  
  "laatoo/sdk/server/components/data"
)

type ServiceDesign_Ref struct {
  Id    string
}

type ServiceDesign struct {
	*data.SerializableBase `initialize:"SerializableBase"`
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Factory	string `json:"Factory" bson:"Factory" datastore:"Factory"`
	ServiceMethod	string `json:"ServiceMethod" bson:"ServiceMethod" datastore:"ServiceMethod"`
	LoggingLevel	string `json:"LoggingLevel" bson:"LoggingLevel" datastore:"LoggingLevel"`
	LoggingFormat	string `json:"LoggingFormat" bson:"LoggingFormat" datastore:"LoggingFormat"`
	Settings	map[string]interface{} `json:"Settings" bson:"Settings" datastore:"Settings"`
}

func (ent *ServiceDesign) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ServiceDesign",
	}
}

