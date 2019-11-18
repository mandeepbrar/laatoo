package main

import (
	"dataadapter/dataadapter"
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: dataadapter.DataAdapterFactory{}},
		core.PluginComponent{Object: dataadapter.DataAdapterModule{}}}
}
