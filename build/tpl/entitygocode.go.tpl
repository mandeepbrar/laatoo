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
  {{#titleField titleField}}{{/titleField}} string
}

type {{#type name}}{{/type}} struct {
  {{#if multitenant}}
	data.SoftDeleteAuditableMT `bson:",inline"`
	{{else}}
	data.SoftDeleteAuditable `bson:",inline"`
  {{/if}}
  {{#fields fields}}{{/fields}}
}

func (ent *{{#type name}}{{/type}}) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "{{#titleField titleField}}{{/titleField}}",
		Type:            "{{#type name}}{{/type}}",
		SoftDeleteField: "Deleted",
		PreSave:         {{#presave presave}}{{/presave}},
		PostSave:        {{#postsave postsave}}{{/postsave}},
		PostLoad:        {{#postload postload}}{{/postload}},
		Auditable:       true,
		Multitenant:     {{#multitenant multitenant}}{{/multitenant}},
		Collection:      "{{#collection collection name}}{{/collection}}",
		Cacheable:       {{#cacheable cacheable}}{{/cacheable}},
	}
}
