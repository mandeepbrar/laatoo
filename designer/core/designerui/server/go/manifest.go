package main

import (
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	objects := ObjectsManifest(provider)
	services := ServicesManifest(provider)
	objs := []core.PluginComponent{}
	objs = append(objs, objects...)
	objs = append(objs, services...)
	return objs
}
