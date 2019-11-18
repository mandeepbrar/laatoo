package main

import (
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: BeanstalkProducer{}},
		core.PluginComponent{Object: BeanstalkConsumer{}}}
}
