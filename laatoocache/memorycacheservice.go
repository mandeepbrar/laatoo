package laatoocache

import (
	"laatoocore"
	"laatoosdk/core"
	"laatoosdk/utils"
	//	"laatoosdk/errors"
	"laatoosdk/log"
	"reflect"
)

type MemoryCacheService struct {
	memoryStorer *utils.MemoryStorer
}

const (
	CONF_MEMORYCACHE_NAME = "memory_cache"
)

func init() {
	laatoocore.RegisterObjectProvider(CONF_MEMORYCACHE_NAME, MemoryCacheServiceFactory)
}

func MemoryCacheServiceFactory(ctx core.Context, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Creating redis cache service ")
	memorySvc := &MemoryCacheService{memoryStorer: utils.NewMemoryStorer()}

	return memorySvc, nil
}

func (svc *MemoryCacheService) GetServiceType() string {
	return core.SERVICE_TYPE_DATA
}

//name of the service
func (svc *MemoryCacheService) GetName() string {
	return CONF_MEMORYCACHE_NAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *MemoryCacheService) Initialize(ctx core.Context) error {
	return nil
}

//The service starts serving when this method is called
func (svc *MemoryCacheService) Serve(ctx core.Context) error {
	return nil
}

func (svc *MemoryCacheService) Delete(ctx core.Context, key string) error {
	return svc.memoryStorer.Delete(key)
}

func (svc *MemoryCacheService) PutObject(ctx core.Context, key string, val interface{}) error {
	svc.memoryStorer.PutObject(key, val)
	return nil
}

func (svc *MemoryCacheService) GetObject(ctx core.Context, key string, val interface{}) error {
	obj, err := svc.memoryStorer.GetObject(key)
	if err != nil {
		return err
	}
	reflect.ValueOf(val).Elem().Set(reflect.ValueOf(obj).Elem())
	return nil
}

func (svc *MemoryCacheService) GetMulti(ctx core.Context, keys []string, val map[string]interface{}) error {
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
	return nil
}

//Execute method
func (svc *MemoryCacheService) Execute(ctx core.Context, name string, params map[string]interface{}) (interface{}, error) {
	return nil, nil
}
