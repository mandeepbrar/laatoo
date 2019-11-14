package main

import "laatoo/sdk/server/core"

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: "restimporter", Object: restImporter{}},
		core.PluginComponent{Name: "restexporter", Object: restExporter{}}}
}
