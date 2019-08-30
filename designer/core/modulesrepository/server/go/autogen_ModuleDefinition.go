package main

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
	Params	[]ModuleParam `json:"Params" bson:"Params" datastore:"Params"`
	Dependencies	map[string]string `json:"Dependencies" bson:"Dependencies" datastore:"Dependencies"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	UIDependencies	map[string]string `json:"UIDependencies" bson:"UIDependencies" datastore:"UIDependencies"`
	Objects	[]ObjectDefinition `json:"Objects" bson:"Objects" datastore:"Objects"`
	Services	[]string `json:"Services" bson:"Services" datastore:"Services"`
	Factories	[]string `json:"Factories" bson:"Factories" datastore:"Factories"`
	Channels	[]string `json:"Channels" bson:"Channels" datastore:"Channels"`
	Engines	[]string `json:"Engines" bson:"Engines" datastore:"Engines"`
	Rules	[]string `json:"Rules" bson:"Rules" datastore:"Rules"`
	Tasks	[]string `json:"Tasks" bson:"Tasks" datastore:"Tasks"`
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


