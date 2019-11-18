package main

import (
	"laatoo/sdk/server/core"
)

const (
	CONF_SECURITYSERVICE_TOKENVALIDATION = "TOKEN_VALIDATE"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: TokenValidationService{}}}
}
