package main

import "laatoo/sdk/server/core"

const (
	DATAPIPELINE_FACTORY = "datapipelinefactory"
	DATAPIPELINE_SERVICE = "datapipelineservice"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: DATAPIPELINE_FACTORY, Object: dataPipelineFactory{}},
		core.PluginComponent{Name: DATAPIPELINE_SERVICE, Object: pipelineService{}}}
}
