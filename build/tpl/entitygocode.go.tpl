{{#if sdkinclude}}
package {{modulename}}
{{else}}
package main
{{/if}}

import (
  {{#imports imports}}{{/imports}}
  "laatoo/sdk/server/components/data"
  "laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

/*type {{#type name}}{{/type}}_Ref struct {
  Id    string
  {{#titleField titleField}}{{/titleField}} string
}*/

type {{#type name}}{{/type}} struct {
  {{#if multitenant}}
	data.Storable `json:",inline" bson:",inline" laatoo:"auditable, softdelete, multitenant"`
	{{else}}
	data.Storable `json:",inline" bson:",inline" laatoo:"auditable, softdelete"`
  {{/if}}
  {{#fields fields}}{{/fields}}
}

func (ent *{{#type name}}{{/type}}) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "{{#titleField titleField}}{{/titleField}}",
		Type:            "{{#type name}}{{/type}}",
		PreSave:         {{#presave presave}}{{/presave}},
		PostSave:        {{#postsave postsave}}{{/postsave}},
		PostLoad:        {{#postload postload}}{{/postload}},
		Auditable:       true,
		Multitenant:     {{#multitenant multitenant}}{{/multitenant}},
		Collection:      "{{#collection collection name}}{{/collection}}",
		Cacheable:       {{#cacheable cacheable}}{{/cacheable}},
	}
}



func (ent *{{#type name}}{{/type}}) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	var err error

  {{#fieldReadAlls fields}}{{/fieldReadAlls}}

	return ent.Storable.ReadAll(c, cdc, rdr)
}

func (ent *{{#type name}}{{/type}}) ReadProps(c ctx.Context, cdc core.Codec, rdr core.SerializableReader, props map[string]interface{}) error {
  var ok bool
	var err error

  {{#fieldReadProps fields}}{{/fieldReadProps}}
  
	return ent.Storable.ReadProps(c, cdc, rdr, props)
}

func (ent *{{#type name}}{{/type}}) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	var err error

  {{#fieldWriteAlls fields}}{{/fieldWriteAlls}}

	return ent.Storable.WriteAll(c, cdc, wtr)
}

func (ent *{{#type name}}{{/type}}) WriteProps(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter, props map[string]interface{}) error {
  var ok bool
	var err error

  {{#fieldWriteProps fields}}{{/fieldWriteProps}}

	return ent.Storable.WriteProps(c, cdc, wtr, props)
}