package main

import (
	"laatoo/sdk/server/core"
{{#if hasSDK}}
  "laatoo/sdk/modules/{{name}}"
{{/if}}  
)


func ObjectsManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    {{#plugins entities}}{{/plugins}}
  }
}
