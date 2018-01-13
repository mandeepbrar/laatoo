package main

import (
  
  "laatoo/sdk/components/data"
)

type Module_Ref struct {
  Id    string
  Name string
}

type Module struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Instance	bool `json:"Instance" bson:"Instance" datastore:"Instance"`
	Objects	[]string `json:"Objects" bson:"Objects" datastore:"Objects"`
	Services	[]string `json:"Services" bson:"Services" datastore:"Services"`
	Factories	[]string `json:"Factories" bson:"Factories" datastore:"Factories"`
	Channels	[]string `json:"Channels" bson:"Channels" datastore:"Channels"`
	Engines	[]string `json:"Engines" bson:"Engines" datastore:"Engines"`
	Rules	[]string `json:"Rules" bson:"Rules" datastore:"Rules"`
	Tasks	[]string `json:"Tasks" bson:"Tasks" datastore:"Tasks"`
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

func (ent *Module) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Module",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "Module",
		Cacheable:       false,
	}
}
