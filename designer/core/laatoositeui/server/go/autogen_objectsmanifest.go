package main

import (
	"laatoo/sdk/server/core"
)


func ObjectsManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "Solution", Object: Solution{}, Metadata: core.NewInfo("","Solution", map[string]interface{}{"descriptor":"{\"name\":\"Solution\",\"inherits\":\"\",\"imports\":[],\"multitenant\":true,\"form\":{\"overlay\":true,\"layout\":[\"Name\",\"Description\"]},\"titleField\":\"Name\",\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"User\":{\"type\":\"string\"}}}"})},
  }
}
