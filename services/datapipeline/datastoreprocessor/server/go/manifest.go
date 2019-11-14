package main

import "laatoo/sdk/server/core"

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: "datastoreexporter", Object: dataExporter{}},
		core.PluginComponent{Name: "datastoreimporter", Object: dataImporter{}}}
}
