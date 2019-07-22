package main

import (
	"laatoo/sdk/server/core"
)


func ObjectsManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "Solution", Object: Solution{}, Metadata: core.NewInfo("","Solution", map[string]interface{}{"descriptor":"{\"name\":\"Solution\",\"inherits\":\"\",\"imports\":[],\"multitenant\":true,\"form\":{\"overlay\":true,\"layout\":[\"Name\",\"Description\",\"Modules\"]},\"titleField\":\"Name\",\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"User\":{\"type\":\"string\"},\"Modules\":{\"type\":\"storableref\",\"entity\":\"ModuleDefinition\",\"label\":\"Solution Modules\",\"addwidget\":\"ModuleSelect\",\"addwidgetmodule\":\"modulesrepository\",\"list\":true,\"widget\":\"SubEntity\",\"module\":\"subentity\"}}}"})},
  }
}
