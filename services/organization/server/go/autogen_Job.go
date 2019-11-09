package main

import (
  
  "laatoo/sdk/server/components/data"
)

type Job_Ref struct {
  Id    string
  Account string
}

type Job struct {
	data.Storable `laatoo:"auditable, softdelete"`
  
	JobID	string `json:"JobID" bson:"JobID" datastore:"JobID"`
	Title	string `json:"Title" bson:"Title" datastore:"Title"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	OrgUnit	OrgNode_Ref `json:"OrgUnit" bson:"OrgUnit" datastore:"OrgUnit"`
}

func (ent *Job) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Account",
		Type:            "Job",
		PreSave:         true,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "Job",
		Cacheable:       false,
	}
}
