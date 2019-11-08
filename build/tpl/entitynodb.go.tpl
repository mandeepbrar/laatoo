{{#if sdkinclude}}
package {{modulename}}
{{else}}
package main
{{/if}}


import (
  {{#imports imports}}{{/imports}}
  "laatoo/sdk/server/components/data"
)

type {{#type name}}{{/type}}_Ref struct {
  Id    string
}

type {{#type name}}{{/type}} struct {
	data.Storable 
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  {{#fields fields}}{{/fields}}
}

func (ent *{{#type name}}{{/type}}) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "{{#titleField titleField}}{{/titleField}}",
		Type:            "{{#type name}}{{/type}}",
	}
}

{{#if sdkinclude}}
{{#fieldFuncs fields name}}{{/fieldFuncs}}
{{/if}}
