package main

import (
  
  "laatoo/sdk/server/components/data"
)

type OrgNode_Ref struct {
  Id    string
  Title string
}

type OrgNode struct {
	data.Storable `laatoo:"auditable, softdelete"`
  
	Title	string `json:"Title" bson:"Title" datastore:"Title"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Parent	OrgNode_Ref `json:"Parent" bson:"Parent" datastore:"Parent"`
	Data1	data.StorableRef `json:"Data1" bson:"Data1" datastore:"Data1"`
	Data2	data.StorableRef `json:"Data2" bson:"Data2" datastore:"Data2"`
	Data3	data.StorableRef `json:"Data3" bson:"Data3" datastore:"Data3"`
}

func (ent *OrgNode) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Title",
		Type:            "OrgNode",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "OrgNode",
		Cacheable:       false,
	}
}
