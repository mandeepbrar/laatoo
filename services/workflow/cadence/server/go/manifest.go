package main

import (
	"cadence/cadence"
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: cadence.CadenceWorkerService{}},
		core.PluginComponent{Object: cadence.SimpleWorkflow{}},
		core.PluginComponent{Object: cadence.WorkflowInitiator{}},
		core.PluginComponent{Object: cadence.DSLWorkflow{}},
	}
}
