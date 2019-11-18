package main

import (
	"gaetasks/gaetasks"
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: gaetasks.GaeProducer{}},
		core.PluginComponent{Object: gaetasks.GaeConsumer{}}}
}
