package main

import (
	"laatoo/sdk/server/core"
)


func ObjectsManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "AccountUser", Object: AccountUser{}, Metadata: core.NewInfo("","AccountUser", map[string]interface{}{"descriptor":"{\"name\":\"AccountUser\",\"inherits\":\"\",\"imports\":[],\"multitenant\":true,\"postsave\":true,\"form\":{\"overlay\":true,\"layout\":[\"FName\",\"LName\",\"User\",\"Email\",\"Picture\"]},\"titleField\":\"Name\",\"fields\":{\"FName\":{\"type\":\"string\"},\"LName\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Picture\":{\"type\":\"string\"},\"User\":{\"type\":\"string\"},\"Email\":{\"type\":\"string\"},\"Disabled\":{\"type\":\"bool\"},\"Account\":{\"type\":\"string\"}}}"})},core.PluginComponent{Name: "Solution", Object: Solution{}, Metadata: core.NewInfo("","Solution", map[string]interface{}{"descriptor":"{\"name\":\"Solution\",\"inherits\":\"\",\"imports\":[],\"multitenant\":true,\"form\":{\"overlay\":true,\"layout\":[\"Name\",\"Description\"]},\"titleField\":\"Name\",\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"User\":{\"type\":\"string\"}}}"})},
  }
}
