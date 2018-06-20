package main

import (
  
  "laatoo/sdk/components/data"
)

type Solution_Ref struct {
  Id    string
  Name string
}

type Solution struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Modules	[]Module `json:"Modules" bson:"Modules" datastore: "Modules"`
}

func (ent *Solution) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Solution",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "Solution",
		Cacheable:       false,
	}
}
