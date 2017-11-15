package main

import (
	"laatoo/sdk/core"
)


func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "Comment", Object: Comment{}, Metadata: core.NewInfo("","Comment", map[string]interface{}{"descriptor":"{\"name\":\"Comment\",\"inherits\":\"\",\"imports\":[],\"fields\":{\"Post\":{\"type\":\"string\"},\"PostTitle\":{\"type\":\"string\"},\"Blocked\":{\"type\":\"bool\"}}}"})},
  }
}
