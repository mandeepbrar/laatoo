package main

import (
	"laatoo/sdk/server/core"
	"oauthlogin/oauthlogin"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: oauthlogin.OAuthLoginService{}}}
}
