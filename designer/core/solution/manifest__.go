package main

import (
	"laatoo/sdk/core"
)


func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "Solution", Object: Solution{}, Metadata: core.NewInfo("","Solution", map[string]interface{}{"descriptor":"{\"name\":\"Solution\",\"inherits\":\"\",\"imports\":[],\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"}}}"})},
  }
}
