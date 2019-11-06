package main

import (
  
  "laatoo/sdk/server/components/data"
)

type AccountUser_Ref struct {
  Id    string
  Name string
}

type AccountUser struct {
	*data.SoftDeleteAuditableMT `initialize:"SoftDeleteAuditableMT"`
  
	FName	string `json:"FName" bson:"FName" datastore:"FName"`
	LName	string `json:"LName" bson:"LName" datastore:"LName"`
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Picture	string `json:"Picture" bson:"Picture" datastore:"Picture"`
	User	string `json:"User" bson:"User" datastore:"User"`
	Email	string `json:"Email" bson:"Email" datastore:"Email"`
	Disabled	bool `json:"Disabled" bson:"Disabled" datastore:"Disabled"`
	Account	string `json:"Account" bson:"Account" datastore:"Account"`
}

func (ent *AccountUser) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "AccountUser",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        true,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     true,
		Collection:      "AccountUser",
		Cacheable:       false,
	}
}
