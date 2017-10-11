package main

import (
  
  "laatoo/sdk/components/data"
)

var (
	entityConf = &data.StorableConfig{
		IdField:         "Id",
		Type:            "Comment",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "Comment",
		Cacheable:       false,
	}
)

type Comment struct {
	data.SoftDeleteAuditable `bson:",inline"`
  
	Post	string `json:"Post" bson:"Post"  `
	PostTitle	string `json:"PostTitle" bson:"PostTitle"  `
	Blocked	bool `json:"Blocked" bson:"Blocked"  `
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

func (ent *Comment) Config() *data.StorableConfig {
	return entityConf
}
