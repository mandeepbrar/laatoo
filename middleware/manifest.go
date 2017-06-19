package main

import (
	"laatoo/sdk/core"
)

func Manifest() []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CHECK_CREATION_USER, ServiceFunc: CheckCreationUser},
		core.PluginComponent{Name: CHECK_USER_OWN_ACCOUNT, ServiceFunc: OwnUserAccountEnforce},
		core.PluginComponent{Name: SVC_CHECKPERMISSION, Object: checkPermissionService{}}}
}
