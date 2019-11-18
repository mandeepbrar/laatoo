package main

import (
	"laatoo/sdk/server/core"
	"staticfileserver/staticfileserver"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: staticfileserver.StaticFiles{}},
		core.PluginComponent{Object: staticfileserver.NamedFileService{}},
		core.PluginComponent{Object: staticfileserver.TemplateFileService{}},
		core.PluginComponent{Object: staticfileserver.BundledFileService{}}}
}
