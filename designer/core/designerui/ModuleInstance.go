package main

import (
  
  "laatoo/sdk/components/data"
)

type ModuleInstance_Ref struct {
  Id    string
  Name string
}

type ModuleInstance struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	ModuleRef	*Module `json:"ModuleRef" bson:"ModuleRef" datastore:"ModuleRef"`
	Module	string `json:"Module" bson:"Module" datastore: "Module"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Settings	[]Param `json:"Settings" bson:"Settings" datastore: "Settings"`
/*	Post                     string `json:"Post" bson:"Post"`
	PostTitle                string
	PostTitleEng             string
	BodyGur                  string `json:"BodyGur" bson:"BodyGur" datastore:",noindex"`
	BodyEng                  string `json:"BodyEng" bson:"BodyEng" datastore:",noindex"`
	UserName                 string `json:"UserName" bson:"UserName" datastore:",noindex"`
	UserId                   string `json:"UserId" bson:"UserId"`
	UserPic                  string `json:"UserPic" bson:"UserPic" datastore:",noindex"`
	Status                   string `json:"Status" bson:"Status"`
	Blocked                  bool   `json:"Blocked" bson:"Blocked"`*/
}

func (ent *ModuleInstance) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "ModuleInstance",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "ModuleInstance",
		Cacheable:       false,
	}
}
