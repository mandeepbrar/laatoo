package main

import (
	"laatoo/sdk/server/core"
	"role/role"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: role.Role{}}}
}
