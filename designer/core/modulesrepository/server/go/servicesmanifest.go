package main

import (
	"laatoo/sdk/server/core"
)

func ServicesManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
		core.PluginComponent{Name: "repositoryimport", Object: RepositoryImport{}},
		core.PluginComponent{Name: "repositoryupdate", Object: RepositoryUpdate{}},
		core.PluginComponent{Name: "entitlementcreator", Object: EntitlementCreationService{}}}
}
