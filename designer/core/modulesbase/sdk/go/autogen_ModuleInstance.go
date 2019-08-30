package modulesbase

import (
  
  "laatoo/sdk/server/components/data"
)

type ModuleInstance_Ref struct {
  Id    string
  Name string
}

type ModuleInstance struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Module	data.StorableRef `json:"Module" bson:"Module" datastore:"Module"`
	LoggingLevel	string `json:"LoggingLevel" bson:"LoggingLevel" datastore:"LoggingLevel"`
	LoggingFormat	string `json:"LoggingFormat" bson:"LoggingFormat" datastore:"LoggingFormat"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Settings	map[string]interface{} `json:"Settings" bson:"Settings" datastore:"Settings"`
}

func (ent *ModuleInstance) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ModuleInstance",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "ModuleInstance",
		Cacheable:       false,
	}
}