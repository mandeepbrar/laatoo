package main

import (
	"filesystemstorage/filesystemstorage"
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: filesystemstorage.FileSystemSvc{}}}
}
