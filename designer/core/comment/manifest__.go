package main

import (
	"laatoo/sdk/core"
)


func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "Comment", Object: Comment{}},
  }
}
