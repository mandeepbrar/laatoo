package main

import (
	"laatoo/sdk/server/core"
	"vaultsecrets/vaultsecrets"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: vaultsecrets.VaultSecretsSvc{}}}
}
