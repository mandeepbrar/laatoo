package main

import (
	"laatoo/sdk/core"
)


func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "Application", Object: Application{}, Metadata: core.NewInfo("","Application", map[string]interface{}{"descriptor":"{\"name\":\"Application\",\"inherits\":\"\",\"imports\":[],\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"}}}"})},
  }
}
