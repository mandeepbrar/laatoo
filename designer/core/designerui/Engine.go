package main

import (
  
  "laatoo/sdk/components/data"
)

type Engine_Ref struct {
  Id    string
  Name string
}

type Engine struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	EngineType	string `json:"EngineType" bson:"EngineType" datastore:"EngineType"`
	Address	string `json:"Address" bson:"Address" datastore:"Address"`
	Framework	string `json:"Framework" bson:"Framework" datastore:"Framework"`
	SSL	bool `json:"SSL" bson:"SSL" datastore:"SSL"`
	CORS	bool `json:"CORS" bson:"CORS" datastore:"CORS"`
	Path	string `json:"Path" bson:"Path" datastore:"Path"`
	CORSHosts	string `json:"CORSHosts" bson:"CORSHosts" datastore:"CORSHosts"`
	QueryParams	string `json:"QueryParams" bson:"QueryParams" datastore:"QueryParams"`
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

func (ent *Engine) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Engine",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "Engine",
		Cacheable:       false,
	}
}
