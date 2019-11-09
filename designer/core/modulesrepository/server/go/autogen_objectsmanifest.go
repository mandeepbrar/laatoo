package main

import (
	"laatoo/sdk/server/core"
  "laatoo/sdk/modules/modulesrepository"
)


func ObjectsManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "Entitlement", Object: modulesrepository.Entitlement{}, Metadata: core.NewInfo("","Entitlement", map[string]interface{}{"descriptor":"{\"name\":\"Entitlement\",\"sdkinclude\":true,\"inherits\":\"\",\"imports\":[\"laatoo/sdk/modules/modulesbase\",\"laatoo/sdk/modules/laatoositeui\"],\"multitenant\":true,\"presave\":true,\"form\":{\"layout\":[\"Name\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Solution\":{\"type\":\"storableref\",\"entity\":\"laatoositeui.Solution\"},\"Local\":{\"type\":\"bool\"},\"Module\":{\"type\":\"storableref\",\"entity\":\"modulesbase.ModuleDefinition\"}}}"})},
  }
}
