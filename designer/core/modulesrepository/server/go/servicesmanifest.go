package main

import (
	"laatoo/sdk/server/core"
)

func ServicesManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
		core.PluginComponent{Name: "repositoryimport", Object: RepositoryImport{}},
		core.PluginComponent{Name: "entitlementcreator", Object: EntitlementCreationService{}},
		core.PluginComponent{Name: "configformsbuilder", Object: ConfigFormsBuilder{}},
		core.PluginComponent{Name: "repositoryupdate", Object: RepositoryUpdate{}}}
}
