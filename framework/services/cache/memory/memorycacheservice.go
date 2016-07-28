package memory

import (
	"laatoo/framework/core/objects"
	"laatoo/framework/services/cache/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/utils"
)

type MemoryCacheFactory struct {
}

const (
	CONF_MEMORYCACHE_NAME = "memory_cache"
	CONF_MEMORYCACHE_SVC  = "cache"
)

func init() {
	objects.Register(CONF_MEMORYCACHE_NAME, MemoryCacheFactory{})
}

//Create the services configured for factory.
func (mf *MemoryCacheFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	if method == CONF_MEMORYCACHE_SVC {
		return &MemoryCacheService{memoryStorer: utils.NewMemoryStorer(), name: name}, nil
	}
	return nil, nil
}

func (ds *MemoryCacheFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//The services start serving when this method is called
func (ds *MemoryCacheFactory) Start(ctx core.ServerContext) error {
	return nil
}

type MemoryCacheService struct {
	memoryStorer *utils.MemoryStorer
	name         string
}

func (svc *MemoryCacheService) Delete(ctx core.RequestContext, bucket string, key string) error {
	return svc.memoryStorer.DeleteObject(key)
}

func (svc *MemoryCacheService) PutObject(ctx core.RequestContext, bucket string, key string, val interface{}) error {
	b, err := common.Encode(val)
	if err != nil {
		return err
	}
	svc.memoryStorer.PutObject(key, b)
	return nil
}

func (svc *MemoryCacheService) GetObject(ctx core.RequestContext, bucket string, key string, objectType string) (interface{}, bool) {
	obj, err := svc.memoryStorer.GetObject(key)
	if err != nil || obj == nil {
		return nil, false
	}
	svrctx := ctx.ServerContext()
	val, err := svrctx.CreateObject(objectType)
	if err != nil {
		return nil, false
	}
	err = common.Decode(obj.([]byte), val)
	if err != nil {
		return nil, false
	}
	//	reflect.ValueOf(val).Elem().Set(reflect.ValueOf(obj).Elem())
	return val, true
}

func (svc *MemoryCacheService) GetMulti(ctx core.RequestContext, bucket string, keys []string, objectType string) map[string]interface{} {
	val := make(map[string]interface{})
	svrctx := ctx.ServerContext()
	objectcreator, err := svrctx.GetObjectCreator(objectType)
	if err != nil {
		return val
	}
	for _, key := range keys {
		b, err := svc.memoryStorer.GetObject(key)
		if err != nil || b == nil {
			val[key] = nil
		} else {
			obj := objectcreator()
			err = common.Decode(b.([]byte), obj)
			if err != nil {
				val[key] = nil
				continue
			}
			val[key] = obj
		}
	}
	return val
}

func (ms *MemoryCacheService) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

func (ms *MemoryCacheService) Invoke(ctx core.RequestContext) error {
	return nil
}

func (ms *MemoryCacheService) Start(ctx core.ServerContext) error {
	return nil
}
