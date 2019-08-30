package modulesbase

import (
  
  "laatoo/sdk/server/components/data"
)

type ModuleDefinition_Ref struct {
  Id    string
  Name string
}

type ModuleDefinition struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Version	string `json:"Version" bson:"Version" datastore:"Version"`
	Params	[]Param `json:"Params" bson:"Params" datastore:"Params"`
	Dependencies	map[string]string `json:"Dependencies" bson:"Dependencies" datastore:"Dependencies"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	UIDependencies	map[string]string `json:"UIDependencies" bson:"UIDependencies" datastore:"UIDependencies"`
	Objects	[]ObjectDefinition `json:"Objects" bson:"Objects" datastore:"Objects"`
	Modules	[]ModuleInstance `json:"Modules" bson:"Modules" datastore:"Modules"`
	Services	[]Service `json:"Services" bson:"Services" datastore:"Services"`
	Entities	[]Entity `json:"Entities" bson:"Entities" datastore:"Entities"`
	Factories	[]Factory `json:"Factories" bson:"Factories" datastore:"Factories"`
	Channels	[]Channel `json:"Channels" bson:"Channels" datastore:"Channels"`
	Engines	[]Engine `json:"Engines" bson:"Engines" datastore:"Engines"`
	Rules	[]Rule `json:"Rules" bson:"Rules" datastore:"Rules"`
	Tasks	[]Task `json:"Tasks" bson:"Tasks" datastore:"Tasks"`
}

func (ent *ModuleDefinition) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ModuleDefinition",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "ModuleDefinition",
		Cacheable:       false,
	}
}
