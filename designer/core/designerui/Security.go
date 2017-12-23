package main

import (
  
  "laatoo/sdk/components/data"
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
	supportedrealms	[]string `json:"supportedrealms" bson:"supportedrealms" datastore:"supportedrealms"`
	authservices	[]string `json:"authservices" bson:"authservices" datastore:"authservices"`
	permissions	[]string `json:"permissions" bson:"permissions" datastore:"permissions"`
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
		Collection:      "Security",
		Cacheable:       false,
	}
}
