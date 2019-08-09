package main

import (
  
  "laatoo/sdk/server/components/data"
)

type Security_Ref struct {
  Id    string
  Name string
}

type Security struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Mode	string `json:"Mode" bson:"Mode" datastore:"Mode"`
	RoleSvc	string `json:"RoleSvc" bson:"RoleSvc" datastore:"RoleSvc"`
	Publickey	string `json:"Publickey" bson:"Publickey" datastore:"Publickey"`
	Privatekey	string `json:"Privatekey" bson:"Privatekey" datastore:"Privatekey"`
	Realm	string `json:"Realm" bson:"Realm" datastore:"Realm"`
	Supportedrealms	[]string `json:"Supportedrealms" bson:"Supportedrealms" datastore:"Supportedrealms"`
	Authservices	[]string `json:"Authservices" bson:"Authservices" datastore:"Authservices"`
	Permissions	[]string `json:"Permissions" bson:"Permissions" datastore:"Permissions"`
}

func (ent *Security) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Security",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "Security",
		Cacheable:       false,
	}
}
