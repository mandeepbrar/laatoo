package main

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/utils"
	"laatoo/services/cache/common"
)

const (
	CONF_MEMORYCACHE_SVC = "memory_cache"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_MEMORYCACHE_SVC, Object: MemoryCacheService{}}}
}

type MemoryCacheService struct {
	core.Service
	memoryStorer *utils.MemoryStorer
	name         string
	cacheEncoder *common.CacheEncoder
}

func (svc *MemoryCacheService) Delete(ctx core.RequestContext, bucket string, key string) error {
	return svc.memoryStorer.DeleteObject(common.GetCacheKey(bucket, key))
}

func (svc *MemoryCacheService) PutObject(ctx core.RequestContext, bucket string, key string, val interface{}) error {
	if svc.cacheEncoder == nil {
		svc.memoryStorer.PutObject(common.GetCacheKey(bucket, key), val)
		return nil
	}
	b, err := svc.cacheEncoder.Encode(val)
	if err != nil {
		return err
	}
	svc.memoryStorer.PutObject(common.GetCacheKey(bucket, key), b)
	return nil
}

func (svc *MemoryCacheService) PutObjects(ctx core.RequestContext, bucket string, vals map[string]interface{}) error {
	if svc.cacheEncoder != nil {
		for k, v := range vals {
			b, err := svc.cacheEncoder.Encode(v)
			if err != nil {
				return err
			}
			svc.memoryStorer.PutObject(common.GetCacheKey(bucket, k), b)
		}
	} else {
		for k, v := range vals {
			svc.memoryStorer.PutObject(common.GetCacheKey(bucket, k), v)
		}
	}
	return nil
}

func (svc *MemoryCacheService) Get(ctx core.RequestContext, bucket string, key string) (interface{}, bool) {
	obj, err := svc.memoryStorer.GetObject(common.GetCacheKey(bucket, key))
	if err != nil || obj == nil {
		return nil, false
	}
	return obj, true
}

func (svc *MemoryCacheService) GetObject(ctx core.RequestContext, bucket string, key string, objectType string) (interface{}, bool) {
	if svc.cacheEncoder == nil {
		return svc.Get(ctx, bucket, key)
	}
	obj, err := svc.memoryStorer.GetObject(common.GetCacheKey(bucket, key))
	if err != nil || obj == nil {
		return nil, false
	}
	svrctx := ctx.ServerContext()
	val, err := svrctx.CreateObject(objectType)
	if err != nil {
		return nil, false
	}
	err = svc.cacheEncoder.Decode(obj.([]byte), val)
	if err != nil {
		return nil, false
	}
	return val, true
}

func (svc *MemoryCacheService) GetMulti(ctx core.RequestContext, bucket string, keys []string) map[string]interface{} {
	val := make(map[string]interface{})
	for _, key := range keys {
		cval, err := svc.memoryStorer.GetObject(common.GetCacheKey(bucket, key))
		if err != nil || cval == nil {
			val[key] = nil
		} else {
			val[key] = cval
		}
	}
	return val
}

func (svc *MemoryCacheService) GetObjects(ctx core.RequestContext, bucket string, keys []string, objectType string) map[string]interface{} {
	if svc.cacheEncoder == nil {
		return svc.GetMulti(ctx, bucket, keys)
	}
	val := make(map[string]interface{})
	svrctx := ctx.ServerContext()
	objectcreator, err := svrctx.GetObjectCreator(objectType)
	if err != nil {
		return val
	}
	for _, key := range keys {
		cval, err := svc.memoryStorer.GetObject(common.GetCacheKey(bucket, key))
		if err != nil || cval == nil {
			val[key] = nil
		} else {
			obj := objectcreator()
			err = svc.cacheEncoder.Decode(cval.([]byte), obj)
			if err != nil {
				val[key] = nil
				continue
			}
			val[key] = obj
		}
	}
	return val
}

func (svc *MemoryCacheService) Increment(ctx core.RequestContext, bucket string, key string) error {
	return svc.memoryStorer.Increment(common.GetCacheKey(bucket, key), 1)
}
func (svc *MemoryCacheService) Decrement(ctx core.RequestContext, bucket string, key string) error {
	return svc.memoryStorer.Decrement(common.GetCacheKey(bucket, key), 1)
}

/*
func (ms *MemoryCacheService) Describe(ctx core.ServerContext) {
	ms.SetComponent(ctx, true)
	ms.SetDescription(ctx, "Memory cache component service")
	ms.AddOptionalConfigurations(ctx, map[string]string{config.ENCODING: config.OBJECTTYPE_STRING}, nil)
}*/

func (ms *MemoryCacheService) Start(ctx core.ServerContext) error {
	ms.memoryStorer = utils.NewMemoryStorer()
	encoding, ok := ms.GetStringConfiguration(ctx, config.ENCODING)
	if ok {
		ms.cacheEncoder = common.NewCacheEncoder(ctx, encoding)
	} else {
		ms.cacheEncoder = nil
	}

	return nil
}
