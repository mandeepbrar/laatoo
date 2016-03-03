package data

import (
	"laatoosdk/core"
)

type Cache interface {
	PutObject(ctx core.Context, key string, item interface{}) error
	GetObject(ctx core.Context, key string, val interface{}) error
	GetMulti(ctx core.Context, keys []string, val map[string]interface{}) error
	Delete(ctx core.Context, key string) error
}
