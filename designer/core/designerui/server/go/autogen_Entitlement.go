package main

import (
  
  "laatoo/sdk/server/components/data"
)

type Entitlement_Ref struct {
  Id    string
  Title string
}

type Entitlement struct {
	data.SoftDeleteAuditableMT `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Solution	data.StorableRef `json:"Solution" bson:"Solution"  entity:"Solution" datastore: "Solution"`
	Module	[]Module `json:"Module" bson:"Module" datastore: "Module"`
}

func (ent *Entitlement) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Title",
		Type:            "Entitlement",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     true,
		Collection:      "Entitlement",
		Cacheable:       false,
	}
}
