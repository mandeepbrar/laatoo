package main

import "laatoo/sdk/server/core"

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
		core.PluginComponent{Name: "memoryerrorprocessor", Object: memoryErrorsProcessor{}}}
}
