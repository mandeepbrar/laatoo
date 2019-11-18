package main

import (
	"laatoo/sdk/modules/user"
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: user.DefaultUser{}}}
}
