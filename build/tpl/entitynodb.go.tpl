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

/*
type {{#type name}}{{/type}}_Ref struct {
  Id    string
}*/

type {{#type name}}{{/type}} struct {
	data.Storable ``
  {{#fields fields}}{{/fields}}
}

func (ent *{{#type name}}{{/type}}) Config() *data.StorableConfig {
	return &data.StorableConfig{
    LabelField:      "{{#titleField titleField}}{{/titleField}}",
	}
}

{{#if sdkinclude}}
{{#fieldFuncs fields name}}{{/fieldFuncs}}
{{/if}}



func (ent *{{#type name}}{{/type}}) ReadAll(c ctx.Context, cdc core.Codec, rdr core.SerializableReader) error {
	var err error

  {{#fieldReadAlls fields}}{{/fieldReadAlls}}

	return ent.Storable.ReadAll(c, cdc, rdr)
}

func (ent *{{#type name}}{{/type}}) WriteAll(c ctx.Context, cdc core.Codec, wtr core.SerializableWriter) error {
	var err error

  {{#fieldWriteAlls fields}}{{/fieldWriteAlls}}

	return ent.Storable.WriteAll(c, cdc, wtr)
}
