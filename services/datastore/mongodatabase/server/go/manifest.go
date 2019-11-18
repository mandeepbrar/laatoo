package main

import (
	"laatoo/sdk/server/core"
	"mongodatabase/mongodatabase"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: mongodatabase.MongoDataServicesFactory{}}}
}
