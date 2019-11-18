package main

import (
	"beanstalktasks/beanstalktasks"
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: beanstalktasks.BeanstalkProducer{}},
		core.PluginComponent{Object: beanstalktasks.BeanstalkConsumer{}}}
}
