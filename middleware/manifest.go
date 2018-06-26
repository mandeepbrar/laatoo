package main

import (
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CHECK_CREATION_USER, Object: CheckCreationUser{}},
		core.PluginComponent{Name: CHECK_USER_OWN_ACCOUNT, Object: OwnUserAccountEnforce{}},
		core.PluginComponent{Name: SVC_CHECKPERMISSION, Object: checkPermissionService{}}}
}
