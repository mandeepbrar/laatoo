package main

import "laatoo/sdk/core"

const (
	NAMED_FILES_SVC   = "named_files_service"
	STATIC_FILES_SVC  = "static_files_service"
	FILEBUNDLE_SVC    = "filebundle_service"
	TEMPLATE_FILE_SVC = "template_files_service"
	CONF_STATIC_CACHE = "cache"
	CONF_STATIC_DIR   = "directory"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: STATIC_FILES_SVC, Object: StaticFiles{}},
		core.PluginComponent{Name: NAMED_FILES_SVC, Object: FileService{}},
		core.PluginComponent{Name: TEMPLATE_FILE_SVC, Object: TemplateFileService{}},
		core.PluginComponent{Name: FILEBUNDLE_SVC, Object: BundledFileService{}}}
}
