package main

import (
  
  "laatoo/sdk/components/data"
)

type Channel_Ref struct {
  Id    string
  Name string
}

type Channel struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Service	string `json:"Service" bson:"Service" datastore:"Service"`
}

func (ent *Channel) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Channel",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "Channel",
		Cacheable:       false,
	}
}
