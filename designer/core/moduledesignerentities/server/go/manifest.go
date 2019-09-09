package main

import (
	"laatoo/sdk/server/core"
)


func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	objects := ObjectsManifest(provider)
	objs := []core.PluginComponent{
		core.PluginComponent{Name: "moduledesign_entityrule", Object: ModuleDesignEntitiesRule{}},
	}
	objs = append(objs, objects...)
	return objs
}
