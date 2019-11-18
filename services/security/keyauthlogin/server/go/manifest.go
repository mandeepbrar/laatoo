package main

import (
	"keyauthlogin/keyauthlogin"
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: keyauthlogin.KeyAuthService{}}}
}
