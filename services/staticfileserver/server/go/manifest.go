package main

import (
	"laatoo/sdk/server/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: StaticFiles{}},
		core.PluginComponent{Object: NamedFileService{}},
		core.PluginComponent{Object: TemplateFileService{}},
		core.PluginComponent{Object: BundledFileService{}}}
}
