package services

import (
	"fmt"
	"laatoo/sdk/core"
)

type Cache interface {
	core.Service
	PutObject(ctx core.Context, key string, item interface{}) error
	GetObject(ctx core.Context, key string, val interface{}) bool
	GetMulti(ctx core.Context, keys []string, val map[string]interface{}) bool
	Delete(ctx core.Context, key string) error
}

func GetCacheKey(objectType string, variants ...interface{}) string {
	return fmt.Sprintf("%s_%#v", objectType, variants)
}
