package main

import (
	"laatoo/sdk/server/core"
)


func ObjectsManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "ModuleDesignGeneral", Object: ModuleDesignGeneral{}, Metadata: core.NewInfo("","ModuleDesignGeneral", map[string]interface{}{"descriptor":"{\"name\":\"ModuleDesignGeneral\",\"imports\":[\"laatoo/sdk/modules/modulesbase\"],\"inherits\":\"\",\"multitenant\":true,\"form\":{\"tabs\":{\"Info\":[\"Name\",\"Description\"],\"Params\":[\"Params\"],\"UIDependencies\":[\"UIDependencies\"],\"Dependencies\":[\"Dependencies\"]},\"verticaltabs\":true,\"overlay\":true,\"submit\":\"Save\",\"layout\":[\"Info\",\"Params\",\"Dependencies\",\"UIDependencies\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Params\":{\"type\":\"subentity\",\"entity\":\"modulesbase.Param\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"}},\"UIDependencies\":{\"type\":\"subentity\",\"entity\":\"modulesbase.Dependency\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"}},\"Dependencies\":{\"type\":\"subentity\",\"entity\":\"modulesbase.Dependency\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"}}}}"})},
  }
}
