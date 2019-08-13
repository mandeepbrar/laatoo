package main

import (
  
  "laatoo/sdk/server/components/data"
)

type Module_Ref struct {
  Id    string
  Name string
}

type Module struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Instance	bool `json:"Instance" bson:"Instance" datastore:"Instance"`
	Objects	[]string `json:"Objects" bson:"Objects" datastore:"Objects"`
	Params	map[string]ModuleParam `json:"Params" bson:"Params" datastore:"Params"`
	ParamsForm	map[string]interface{} `json:"ParamsForm" bson:"ParamsForm" datastore:"ParamsForm"`
	Dependencies	map[string]interface{} `json:"Dependencies" bson:"Dependencies" datastore:"Dependencies"`
	Services	[]string `json:"Services" bson:"Services" datastore:"Services"`
	Factories	[]string `json:"Factories" bson:"Factories" datastore:"Factories"`
	Channels	[]string `json:"Channels" bson:"Channels" datastore:"Channels"`
	Engines	[]string `json:"Engines" bson:"Engines" datastore:"Engines"`
	Rules	[]string `json:"Rules" bson:"Rules" datastore:"Rules"`
	Tasks	[]string `json:"Tasks" bson:"Tasks" datastore:"Tasks"`
}

func (ent *Module) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Module",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "Module",
		Cacheable:       false,
	}
}
