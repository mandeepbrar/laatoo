package cache

import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	//	"laatoo/sdk/log"
	"laatoo/sdk/utils"
	"reflect"
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
func (mf *MemoryCacheFactory) CreateService(ctx core.ServerContext, name string, method string) (core.Service, error) {
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

func (svc *MemoryCacheService) Delete(ctx core.RequestContext, key string) error {
	return svc.memoryStorer.DeleteObject(key)
}

func (svc *MemoryCacheService) PutObject(ctx core.RequestContext, key string, val interface{}) error {
	svc.memoryStorer.PutObject(key, val)
	return nil
}

func (svc *MemoryCacheService) PutDerivedObject(ctx core.RequestContext, key string, val interface{}) error {
	svc.memoryStorer.PutObject(key, val)
	return nil
}

func (svc *MemoryCacheService) GetObject(ctx core.RequestContext, key string, val interface{}) bool {
	obj, err := svc.memoryStorer.GetObject(key)
	if err != nil {
		return false
	}
	if obj == nil || reflect.ValueOf(obj).IsNil() {
		return false
	}
	reflect.ValueOf(val).Elem().Set(reflect.ValueOf(obj).Elem())
	return true
}

func (svc *MemoryCacheService) GetMulti(ctx core.RequestContext, keys []string, val map[string]interface{}) bool {
	/*_, err := memcache.GetMulti(ctx.GetAppengineContext(), keys)
	if err != nil {
		return err
	} else {
		/*arr := reflect.ValueOf(val).Elem()
		length := arr.Len()
		for i := 0; i < length; i++ {
			itemVal, ok := items[keys[i]]
			if ok {
				dec := gob.NewDecoder(bytes.NewReader(itemVal.Value))
				stor := arr.Index(i).Addr().Interface()
				err := dec.Decode(stor)
				if err != nil {
					return err
				}
			}
		}
	}	*/
	return false
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
