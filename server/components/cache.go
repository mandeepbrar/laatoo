package components

import (
	"time"

	"laatoo.io/sdk/server/core"
)

type CacheComponent interface {
	PutTempObject(ctx core.RequestContext, bucket string, key string, item interface{}, ttl time.Duration) error
	PutObject(ctx core.RequestContext, bucket string, key string, item interface{}) error
	PutObjects(ctx core.RequestContext, bucket string, vals core.StringMap) error
	GetObject(ctx core.RequestContext, bucket string, key string, objectType string) (interface{}, bool)
	Get(ctx core.RequestContext, bucket string, key string) (interface{}, bool)
	GetObjects(ctx core.RequestContext, bucket string, keys []string, objectType string) core.StringMap
	GetMulti(ctx core.RequestContext, bucket string, keys []string) core.StringMap
	Delete(ctx core.RequestContext, bucket string, key string) error
	Increment(ctx core.RequestContext, bucket string, key string) error
	Decrement(ctx core.RequestContext, bucket string, key string) error
}
