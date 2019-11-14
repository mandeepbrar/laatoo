package main

import (
  
  "laatoo/sdk/server/components/data"
)

/*type ModuleDesignServer_Ref struct {
  Id    string
  Title string
}*/

type ModuleDesignServer struct {
	data.Storable `laatoo:"auditable, softdelete, multitenant"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Objects	[]ObjectDefinitionDesign `json:"Objects" bson:"Objects" datastore:"Objects"`
	Modules	[]ModuleInstanceDesign `json:"Modules" bson:"Modules" datastore:"Modules"`
	Services	[]ServiceDesign `json:"Services" bson:"Services" datastore:"Services"`
	Factories	[]FactoryDesign `json:"Factories" bson:"Factories" datastore:"Factories"`
	Channels	[]ChannelDesign `json:"Channels" bson:"Channels" datastore:"Channels"`
	Engines	[]EngineDesign `json:"Engines" bson:"Engines" datastore:"Engines"`
	Rules	[]RuleDesign `json:"Rules" bson:"Rules" datastore:"Rules"`
	Tasks	[]TaskDesign `json:"Tasks" bson:"Tasks" datastore:"Tasks"`
}

func (ent *ModuleDesignServer) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Title",
		Type:            "ModuleDesignServer",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     true,
		Collection:      "ModuleDesignServer",
		Cacheable:       false,
	}
}
