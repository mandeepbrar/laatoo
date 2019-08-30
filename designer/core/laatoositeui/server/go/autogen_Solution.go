package main

import (
  
  "laatoo/sdk/server/components/data"
)

type Solution_Ref struct {
  Id    string
  Name string
}

type Solution struct {
	data.SoftDeleteAuditableMT `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	User	string `json:"User" bson:"User" datastore:"User"`
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
		Multitenant:     true,
		Collection:      "Solution",
		Cacheable:       false,
	}
}


