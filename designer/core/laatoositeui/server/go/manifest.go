package main

import (
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	objects := ObjectsManifest(provider)
	objs := []core.PluginComponent{
		core.PluginComponent{Name: "accountuser_createrule", Object: AccountUserCreateRule{}},
	}
	objs = append(objs, objects...)
	return objs
}
