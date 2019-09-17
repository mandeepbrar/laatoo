package main

import (
	"laatoo/sdk/server/core"
  "laatoo/sdk/modules/moduledesignerbase"
)


func ObjectsManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "ParamDesign", Object: moduledesignerbase.ParamDesign{}, Metadata: core.NewInfo("","ParamDesign", map[string]interface{}{"descriptor":"{\"name\":\"ParamDesign\",\"sdkinclude\":true,\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"collection\":\"<nocollection>\",\"form\":{\"layout\":[\"Name\",\"Type\",\"Description\",\"Required\",\"Default\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"string\",\"value\":\"string\"},{\"text\":\"int\",\"value\":\"int\"},{\"text\":\"bool\",\"value\":\"bool\"},{\"text\":\"stringmap\",\"value\":\"stringmap\"},{\"text\":\"stringsmap\",\"value\":\"stringsmap\"},{\"text\":\"stringarray\",\"value\":\"stringarray\"},{\"text\":\"config\",\"value\":\"config\"}]}}},\"Description\":{\"type\":\"string\"},\"Required\":{\"type\":\"bool\"},\"Default\":{\"type\":\"string\"}}}"})},
  }
}
