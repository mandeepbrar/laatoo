package main

import (
  {{#imports imports}}{{/imports}}
  "laatoo/sdk/components/data"
)

type {{#type name}}{{/type}}_Ref struct {
  Id    string
  {{#titleField titleField}}{{/titleField}} string
}

type {{#type name}}{{/type}} struct {
	data.SoftDeleteAuditable `bson:",inline"`
  {{#fields fields}}{{/fields}}
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

func (ent *{{#type name}}{{/type}}) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "{{#titleField titleField}}{{/titleField}}",
		Type:            "{{#type name}}{{/type}}",
		SoftDeleteField: "Deleted",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Collection:      "{{#collection collection name}}{{/collection}}",
		Cacheable:       {{#cacheable cacheable}}{{/cacheable}},
	}
}
