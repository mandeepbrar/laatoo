package main

import (
	"laatoo/sdk/core"
)


func ObjectsManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "ModuleDefinition", Object: ModuleDefinition{}, Metadata: core.NewInfo("","ModuleDefinition", map[string]interface{}{"descriptor":"{\"name\":\"ModuleDefinition\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"form\":{\"layout\":[\"Name\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Objects\":{\"type\":\"string\",\"list\":true},\"Services\":{\"type\":\"string\",\"list\":true},\"Factories\":{\"type\":\"string\",\"list\":true},\"Channels\":{\"type\":\"string\",\"list\":true},\"Engines\":{\"type\":\"string\",\"list\":true},\"Rules\":{\"type\":\"string\",\"list\":true},\"Tasks\":{\"type\":\"string\",\"list\":true}}}"})},
  }
}
