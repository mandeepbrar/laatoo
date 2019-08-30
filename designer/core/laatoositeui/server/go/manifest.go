package main

import (
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	objects := ObjectsManifest(provider)
	services := ServicesManifest(provider)
	objs := append([]core.PluginComponent{}, objects...)
	objs = append(objs, services...)
	return objs
}
