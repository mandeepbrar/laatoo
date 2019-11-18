package main

import (
	"laatoo/sdk/server/core"
	"tokenvalidation/tokenvalidation"
)

const (
	CONF_SECURITYSERVICE_TOKENVALIDATION = "TOKEN_VALIDATE"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: tokenvalidation.TokenValidationService{}}}
}
