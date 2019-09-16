package main

import (
	"laatoo/sdk/server/core"
)


func ObjectsManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "ModuleDesignUI", Object: ModuleDesignUI{}, Metadata: core.NewInfo("","ModuleDesignUI", map[string]interface{}{"descriptor":"{\"name\":\"ModuleDesignUI\",\"imports\":[\"laatoo/sdk/modules/modulesbase\"],\"inherits\":\"\",\"multitenant\":true,\"form\":{\"tabs\":{\"UIDependencies\":[\"UIDependencies\"],\"Externals\":[\"Externals\"]},\"verticaltabs\":true,\"overlay\":true,\"submit\":\"Save\",\"layout\":[\"UIDependencies\",\"Externals\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Externals\":{\"type\":\"stringmap\"},\"UIDependencies\":{\"type\":\"subentity\",\"entity\":\"modulesbase.Dependency\",\"list\":true}}}"})},
  }
}
