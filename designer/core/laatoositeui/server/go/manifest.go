package main

import (
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	objects := ObjectsManifest(provider)
	objs := append([]core.PluginComponent{}, objects...)
	return objs
}
