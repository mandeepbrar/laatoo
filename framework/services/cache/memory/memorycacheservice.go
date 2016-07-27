package memory

import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	//	"laatoo/sdk/log"
	"laatoo/sdk/utils"
)

type MemoryCacheFactory struct {
}

const (
	CONF_MEMORYCACHE_NAME = "memory_cache"
	CONF_MEMORYCACHE_SVC  = "cache"
)

func init() {
	objects.RegisterObject(CONF_MEMORYCACHE_NAME, createMemoryCacheFactory, nil)
}

func createMemoryCacheFactory(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	return &MemoryCacheFactory{}, nil
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
	svc.memoryStorer.PutObject(key, val)
	return nil
}

func (svc *MemoryCacheService) PutDerivedObject(ctx core.RequestContext, bucket string, key string, val interface{}) error {
	svc.memoryStorer.PutObject(key, val)
	return nil
}

func (svc *MemoryCacheService) GetObject(ctx core.RequestContext, bucket string, key string, val interface{}) bool {
	obj, err := svc.memoryStorer.GetObject(key)
	if err != nil || obj == nil {
		return false
	}
	val = &obj
	//	reflect.ValueOf(val).Elem().Set(reflect.ValueOf(obj).Elem())
	return true
}

func (svc *MemoryCacheService) GetMulti(ctx core.RequestContext, bucket string, keys []string, val map[string]interface{}) {
	for _, key := range keys {
		obj, err := svc.memoryStorer.GetObject(key)
		if err != nil || obj == nil {
			val[key] = nil
		}
		val[key] = &obj
	}
	return
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
