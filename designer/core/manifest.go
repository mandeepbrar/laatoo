package main

import (
	"laatoo/designer/core/entities"
	"laatoo/sdk/core"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: entities.ENTITY_APPL_NAME, Object: entities.Application{}},
		core.PluginComponent{Name: entities.ENTITY_ENV_NAME, Object: entities.Environment{}},
		core.PluginComponent{Name: entities.ENTITY_SERVER_NAME, Object: entities.Server{}},
		core.PluginComponent{Name: entities.ENTITY_SOLUTION_NAME, Object: entities.Solution{}},
	}
}
