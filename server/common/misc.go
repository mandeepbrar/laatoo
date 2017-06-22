package common

import "laatoo/sdk/components"

type RequestContextParams struct {
	EngineContext interface{}
	Cache         components.CacheComponent
	Logger        components.Logger
}
