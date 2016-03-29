package core

type Engine interface {
	InitializeEngine(ctx ServerContext) error
	StartEngine(ctx ServerContext) error
	GetContext() EngineContext
}
