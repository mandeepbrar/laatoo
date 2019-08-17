package main

import (
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	objects := ObjectsManifest(provider)
	services := ServicesManifest(provider)
	objs := []core.PluginComponent{
		core.PluginComponent{Name: "newentitlementrule", Object: NewEntitlementRule{}},
	}
	objs = append(objs, objects...)
	objs = append(objs, services...)
	return objs
}
