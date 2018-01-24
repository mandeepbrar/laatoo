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
