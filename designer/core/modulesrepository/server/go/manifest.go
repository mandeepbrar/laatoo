package main

import (
	"laatoo/sdk/server/core"
	"modulesrepository/autogen"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	objects := autogen.ObjectsManifest(provider)
	services := ServicesManifest(provider)
	objs := []core.PluginComponent{}
	objs = append(objs, objects...)
	objs = append(objs, services...)
	return objs
}
