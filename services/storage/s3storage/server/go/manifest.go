package main

import (
	"laatoo/sdk/server/core"
	"s3storage/s3storage"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: s3storage.S3StorageSvc{}}}
}
