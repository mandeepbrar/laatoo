package elements

import "laatoo.io/sdk/server/core"

type Module interface {
	core.ServerElement
	GetObject() core.Module
	GetModuleProperties() map[string]interface{}
}
