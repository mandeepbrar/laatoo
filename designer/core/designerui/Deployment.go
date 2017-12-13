package main

import (
  
  "laatoo/sdk/components/data"
)

type Deployment_Ref struct {
  Id    string
  Name string
}

type Deployment struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	SolutionRef	*Solution `json:"SolutionRef" bson:"SolutionRef" datastore:"SolutionRef"`
	Solution	string `json:"Solution" bson:"Solution" datastore: "Solution"`
	EnvironmentRef	*Environment `json:"EnvironmentRef" bson:"EnvironmentRef" datastore:"EnvironmentRef"`
	Environment	string `json:"Environment" bson:"Environment" datastore: "Environment"`
	ServerRef	*Server `json:"ServerRef" bson:"ServerRef" datastore:"ServerRef"`
	Server	string `json:"Server" bson:"Server" datastore: "Server"`
	ApplicationRef	*Application `json:"ApplicationRef" bson:"ApplicationRef" datastore:"ApplicationRef"`
	Application	string `json:"Application" bson:"Application" datastore: "Application"`
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

func (ent *Deployment) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Deployment",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "Deployment",
		Cacheable:       false,
	}
}
