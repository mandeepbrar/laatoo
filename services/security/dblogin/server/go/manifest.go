package main

import (
	"laatoo/sdk/server/core"
)

const (
	//login path to be used for local and oauth authentication
	CONF_SECURITYSERVICE_LOGOUT = "DB_LOGOUT"
	CONF_SECURITYSERVICE_DB     = "DB_LOGIN"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: LogoutService{}},
		core.PluginComponent{Object: LoginService{}}}
}
