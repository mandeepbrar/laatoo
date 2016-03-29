package cache

import (
	"laatoo/core/registry"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"laatoo/sdk/utils"
	"reflect"
)

type MemoryCacheFactory struct {
	Conf config.Config
}

const (
	CONF_MEMORYCACHE_NAME = "memory_cache"
	CONF_MEMORYCACHE_SVC  = "cache"
)

func init() {
	registry.RegisterServiceFactoryProvider(CONF_MEMORYCACHE_NAME, MemoryCacheServiceFactory)
}

func MemoryCacheServiceFactory(ctx core.ServerContext, conf config.Config) (core.ServiceFactory, error) {
	log.Logger.Info(ctx, "Creating memory cache service ")
	memoryFac := &MemoryCacheFactory{conf}
	return memoryFac, nil
}

//Create the services configured for factory.
func (mf *MemoryCacheFactory) CreateService(ctx core.ServerContext, name string, conf config.Config) (core.Service, error) {
	if name == CONF_MEMORYCACHE_SVC {
		return &MemoryCacheService{memoryStorer: utils.NewMemoryStorer(), conf: conf}, nil
	}
	return nil, nil
}

//The services start serving when this method is called
func (ds *MemoryCacheFactory) StartServices(ctx core.ServerContext) error {
	return nil
}

type MemoryCacheService struct {
	memoryStorer *utils.MemoryStorer
	conf         config.Config
}

func (svc *MemoryCacheService) Delete(ctx core.Context, key string) error {
	return svc.memoryStorer.DeleteObject(key)
}

func (svc *MemoryCacheService) PutObject(ctx core.Context, key string, val interface{}) error {
	svc.memoryStorer.PutObject(key, val)
	return nil
}

func (svc *MemoryCacheService) GetObject(ctx core.Context, key string, val interface{}) bool {
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

func (svc *MemoryCacheService) GetMulti(ctx core.Context, keys []string, val map[string]interface{}) bool {
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

func (ms *MemoryCacheService) Initialize(ctx core.ServerContext) error {
	return nil
}

func (ms *MemoryCacheService) Invoke(ctx core.RequestContext) error {
	return nil
}

func (ms *MemoryCacheService) GetConf() config.Config {
	return ms.conf
}
