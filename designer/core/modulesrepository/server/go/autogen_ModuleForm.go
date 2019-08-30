package main

import (
  
  "laatoo/sdk/server/components/data"
)

type ModuleForm_Ref struct {
  Id    string
  Title string
}

type ModuleForm struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Form	string `json:"Form" bson:"Form" datastore:"Form"`
}

func (ent *ModuleForm) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Title",
		Type:            "ModuleForm",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "ModuleForm",
		Cacheable:       false,
	}
}


