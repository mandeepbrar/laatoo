package main

import (
	"laatoo/sdk/core"
)


func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	objects := ObjectsManifest(provider)
	return append([]core.PluginComponent{}, objects...)
}
