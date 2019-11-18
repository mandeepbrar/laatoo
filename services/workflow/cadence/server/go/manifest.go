package main

import (
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: CadenceWorkerService{}},
		core.PluginComponent{Object: SimpleWorkflow{}},
		core.PluginComponent{Object: WorkflowInitiator{}},
		core.PluginComponent{Object: DSLWorkflow{}},
	}
}
