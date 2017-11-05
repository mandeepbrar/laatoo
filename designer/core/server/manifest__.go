package main

import (
	"laatoo/sdk/core"
)


func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "Server", Object: Server{}, Metadata: core.NewInfo("","Server", map[string]interface{}{"descriptor":"{\"name\":\"Server\",\"inherits\":\"\",\"imports\":[],\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"}}}"})},
  }
}
