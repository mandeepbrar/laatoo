package main

import (
	"laatoo/sdk/server/core"
)


func ObjectsManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "ModuleDesignGeneral", Object: ModuleDesignGeneral{}, Metadata: core.NewInfo("","ModuleDesignGeneral", map[string]interface{}{"descriptor":"{\"name\":\"ModuleDesignGeneral\",\"imports\":[\"laatoo/sdk/modules/modulesbase\"],\"inherits\":\"\",\"multitenant\":true,\"form\":{\"tabs\":{\"Info\":[\"Name\",\"Description\"],\"Params\":[\"Params\"],\"Dependencies\":[\"Dependencies\"]},\"verticaltabs\":true,\"overlay\":true,\"submit\":\"Save\",\"layout\":[\"Info\",\"Params\",\"Dependencies\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Params\":{\"type\":\"subentity\",\"entity\":\"modulesbase.Param\",\"widget\":{\"props\":{\"label\":\"Parameters\",\"itemLabel\":\"Parameter\"}},\"list\":true},\"Dependencies\":{\"type\":\"subentity\",\"entity\":\"modulesbase.Dependency\",\"widget\":{\"props\":{\"label\":\"Dependencies\",\"itemLabel\":\"Dependency\"}},\"list\":true}}}"})},
  }
}
