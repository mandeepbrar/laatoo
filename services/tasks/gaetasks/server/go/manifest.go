package main

import (
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: GaeProducer{}},
		core.PluginComponent{Object: GaeConsumer{}}}
}
