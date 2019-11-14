package main

import (
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
)

const (
	DATAPIPELINE_RECORD  = "PipelineRecord"
	DATAPIPELINE_SERVICE = "datapipelineservice"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: DATAPIPELINE_SERVICE, Object: pipelineService{}},
		core.PluginComponent{Name: DATAPIPELINE_RECORD, Object: datapipeline.PipelineRecord{}},
	}
}
