package main

import (
	"laatoo/sdk/core"
)


func ObjectsManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "ConfigurationDefinition", Object: ConfigurationDefinition{}, Metadata: core.NewInfo("","ConfigurationDefinition", map[string]interface{}{"descriptor":"{\"name\":\"ConfigurationDefinition\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"form\":{\"layout\":[\"Name\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Default\":{\"type\":\"string\"},\"Required\":{\"type\":\"bool\"}}}"})},core.PluginComponent{Name: "ModuleDefinition", Object: ModuleDefinition{}, Metadata: core.NewInfo("","ModuleDefinition", map[string]interface{}{"descriptor":"{\"name\":\"ModuleDefinition\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"form\":{\"layout\":[\"Name\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Version\":{\"type\":\"string\"},\"Params\":{\"type\":\"stringmap\"},\"Dependencies\":{\"type\":\"stringmap\"},\"Description\":{\"type\":\"string\"},\"UIDependencies\":{\"type\":\"stringmap\"},\"Objects\":{\"type\":\"subentity\",\"entity\":\"ObjectDefinition\",\"list\":true},\"Services\":{\"type\":\"string\",\"list\":true},\"Factories\":{\"type\":\"string\",\"list\":true},\"Channels\":{\"type\":\"string\",\"list\":true},\"Engines\":{\"type\":\"string\",\"list\":true},\"Rules\":{\"type\":\"string\",\"list\":true},\"Tasks\":{\"type\":\"string\",\"list\":true}}}"})},core.PluginComponent{Name: "ObjectDefinition", Object: ObjectDefinition{}, Metadata: core.NewInfo("","ObjectDefinition", map[string]interface{}{"descriptor":"{\"name\":\"ObjectDefinition\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"form\":{\"layout\":[\"Name\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Configurations\":{\"type\":\"subentity\",\"entity\":\"ConfigurationDefinition\",\"list\":true}}}"})},
  }
}
