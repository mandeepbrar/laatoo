package main

import (
  
  "laatoo/sdk/components/data"
)

type ConfigurationDefinition_Ref struct {
  Id    string
  Name string
}

type ConfigurationDefinition struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Type	string `json:"Type" bson:"Type" datastore:"Type"`
	Default	string `json:"Default" bson:"Default" datastore:"Default"`
	Required	bool `json:"Required" bson:"Required" datastore:"Required"`
}

func (ent *ConfigurationDefinition) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ConfigurationDefinition",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "ConfigurationDefinition",
		Cacheable:       false,
	}
}
