package main

import (
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: PipelineService{}},
		core.PluginComponent{Object: datapipeline.PipelineRecord{}},
	}
}
