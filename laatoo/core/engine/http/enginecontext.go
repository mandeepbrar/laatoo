package http

type HttpEngineContext struct {
	eng *HttpEngine
}

func (engctx *HttpEngineContext) GetName() string {
	return CONF_ENGINE_NAME
}
