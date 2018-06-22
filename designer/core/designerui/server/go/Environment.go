package main

import (
  
  "laatoo/sdk/components/data"
)

type Environment_Ref struct {
  Id    string
  Title string
}

type Environment struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	SolutionRef	*Solution `json:"SolutionRef" bson:"SolutionRef" datastore:"SolutionRef"`
	Solution	string `json:"Solution" bson:"Solution" datastore: "Solution"`
	ServerTempRef	*Server `json:"ServerTempRef" bson:"ServerTempRef" datastore:"ServerTempRef"`
	ServerTemp	string `json:"ServerTemp" bson:"ServerTemp" datastore: "ServerTemp"`
	LoggingLevel	string `json:"LoggingLevel" bson:"LoggingLevel" datastore:"LoggingLevel"`
	LoggingFormat	string `json:"LoggingFormat" bson:"LoggingFormat" datastore:"LoggingFormat"`
	Objects	[]string `json:"Objects" bson:"Objects" datastore:"Objects"`
	Modules	[]Module `json:"Modules" bson:"Modules" datastore: "Modules"`
	Instances	[]ModuleInstance `json:"Instances" bson:"Instances" datastore: "Instances"`
	Services	[]Service `json:"Services" bson:"Services" datastore: "Services"`
	Entities	[]Entity `json:"Entities" bson:"Entities" datastore: "Entities"`
	Factories	[]Factory `json:"Factories" bson:"Factories" datastore: "Factories"`
	Channels	[]Channel `json:"Channels" bson:"Channels" datastore: "Channels"`
	Engines	[]Engine `json:"Engines" bson:"Engines" datastore: "Engines"`
	Rules	[]Rule `json:"Rules" bson:"Rules" datastore: "Rules"`
	Tasks	[]Task `json:"Tasks" bson:"Tasks" datastore: "Tasks"`
	Security	Security `json:"Security" bson:"Security" datastore: "Security"`
}

func (ent *Environment) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Title",
		Type:            "Environment",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "Environment",
		Cacheable:       false,
	}
}
