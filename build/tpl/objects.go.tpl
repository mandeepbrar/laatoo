package main

import (
	"laatoo/sdk/core"
)


func ObjectsManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    {{#plugins entities}}{{/plugins}}
  }
}
