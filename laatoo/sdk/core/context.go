package core

type Context interface {
	GetId() string
	GetName() string
	GetParent() Context
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
	GetString(key string) (string, bool)
	GetInt(key string) (int, bool)
	GetStringArray(key string) ([]string, bool)
	SubCtx(name string) Context
	NewCtx(flow string) Context
}
