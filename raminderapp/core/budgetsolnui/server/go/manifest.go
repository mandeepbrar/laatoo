package main

import (
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	objects := ObjectsManifest(provider)
	objs := []core.PluginComponent{}
	svcs := []core.PluginComponent{
		core.PluginComponent{Object: BudgetCreationService{}},
		core.PluginComponent{Object: SubAccountCreationService{}},

		//core.PluginComponent{Name: "importglservice", Object: importGLService{}},
	}
	objs = append(objs, objects...)
	objs = append(objs, svcs...)
	return objs
}
