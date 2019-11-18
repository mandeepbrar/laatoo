package main

import (
	"googlestorage/googlestorage"
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: googlestorage.GoogleStorageSvc{}}}
}
