package main

import (
  
	"laatoo/sdk/modules/modulesbase" 
	"laatoo/sdk/modules/laatoositeui" 
  "laatoo/sdk/server/components/data"
)

type Environment_Ref struct {
  Id    string
  Title string
}

type Environment struct {
	data.Storable `laatoo:"auditable, softdelete, multitenant"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Solution	laatoositeui.Solution_Ref `json:"Solution" bson:"Solution" datastore:"Solution"`
	ServerTemp	Server `json:"ServerTemp" bson:"ServerTemp" datastore:"ServerTemp"`
	LoggingLevel	string `json:"LoggingLevel" bson:"LoggingLevel" datastore:"LoggingLevel"`
	LoggingFormat	string `json:"LoggingFormat" bson:"LoggingFormat" datastore:"LoggingFormat"`
	Objects	[]string `json:"Objects" bson:"Objects" datastore:"Objects"`
	Instances	[]modulesbase.ModuleInstance `json:"Instances" bson:"Instances" datastore:"Instances"`
	Services	[]modulesbase.Service `json:"Services" bson:"Services" datastore:"Services"`
	Entities	[]modulesbase.Entity `json:"Entities" bson:"Entities" datastore:"Entities"`
	Factories	[]modulesbase.Factory `json:"Factories" bson:"Factories" datastore:"Factories"`
	Channels	[]modulesbase.Channel `json:"Channels" bson:"Channels" datastore:"Channels"`
	Engines	[]modulesbase.Engine `json:"Engines" bson:"Engines" datastore:"Engines"`
	Rules	[]modulesbase.Rule `json:"Rules" bson:"Rules" datastore:"Rules"`
	Tasks	[]modulesbase.Task `json:"Tasks" bson:"Tasks" datastore:"Tasks"`
	Security	Security `json:"Security" bson:"Security" datastore:"Security"`
}

func (ent *Environment) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Title",
		Type:            "Environment",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     true,
		Collection:      "Environment",
		Cacheable:       false,
	}
}
