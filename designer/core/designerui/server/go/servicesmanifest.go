package main

import (
	"laatoo/sdk/server/core"
)

func ServicesManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
		core.PluginComponent{Name: "objectsresolver", Object: ObjectResolver{}}}
}
