package main

import (
	"laatoo/sdk/server/core"
	"tunnymemtasks/tunnymemtasks"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: tunnymemtasks.MemTaskProcessor{}}}
}
