package main

import (
	"laatoo/sdk/core"
)


func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "Environment", Object: Environment{}, Metadata: core.NewInfo("","Environment", map[string]interface{}{"descriptor":"{\"name\":\"Environment\",\"inherits\":\"\",\"imports\":[],\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"}}}"})},
  }
}
