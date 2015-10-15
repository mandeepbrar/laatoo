// +build appengine

package laatoocache

import (
	"laatoocore"
	"laatoosdk/context"
	//	"laatoosdk/errors"
	"bytes"
	"encoding/gob"
	"google.golang.org/appengine/memcache"
	"laatoosdk/log"
	"laatoosdk/service"
)

type AppEngineCacheService struct {
	name    string
	context service.ServiceContext
}

const (
	APPNEGINE_LOGGING_CONTEXT       = "appenginecache"
	CONF_APPENGINECACHE_NAME        = "appengine_cache"
	CONF_APPENGINE_CONNECTIONSTRING = "server"
	CONF_APPENGINE_DATABASE         = "db"
)

func init() {
	laatoocore.RegisterObjectProvider(CONF_APPENGINECACHE_NAME, AppEngineCacheServiceFactory)
}

func AppEngineCacheServiceFactory(ctx interface{}, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(ctx, APPNEGINE_LOGGING_CONTEXT, "Creating appengine cache service ")
	appengineSvc := &AppEngineCacheService{name: CONF_APPENGINECACHE_NAME}

	return appengineSvc, nil
}

func (svc *AppEngineCacheService) GetServiceType() string {
	return service.SERVICE_TYPE_DATA
}

//name of the service
func (svc *AppEngineCacheService) GetName() string {
	return svc.name
}

//Initialize the service. Consumer of a service passes the data
func (svc *AppEngineCacheService) Initialize(ctx service.ServiceContext) error {
	svc.context = ctx
	return nil
}

//The service starts serving when this method is called
func (svc *AppEngineCacheService) Serve(ctx interface{}) error {
	return nil
}

func (svc *AppEngineCacheService) Delete(ctx interface{}, key string) error {
	return memcache.Delete(context.GetAppengineContext(ctx), key)
}

func (svc *AppEngineCacheService) PutObject(ctx interface{}, key string, val interface{}) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(val)
	if err != nil {
		return err
	}
	return memcache.Set(context.GetAppengineContext(ctx), &memcache.Item{Key: key, Value: buf.Bytes()})
}

func (svc *AppEngineCacheService) GetObject(ctx interface{}, key string, val interface{}) error {
	item, err := memcache.Get(context.GetAppengineContext(ctx), key)
	if err != nil {
		return err
	} else {
		dec := gob.NewDecoder(bytes.NewReader(item.Value))
		err := dec.Decode(val)
		if err != nil {
			return err
		}
		return nil
	}
}

func (svc *AppEngineCacheService) GetMulti(ctx interface{}, keys []string, val map[string]interface{}) error {
	_, err := memcache.GetMulti(context.GetAppengineContext(ctx), keys)
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
		}*/
	}
	return nil
}

//Execute method
func (svc *AppEngineCacheService) Execute(ctx interface{}, name string, params map[string]interface{}) (interface{}, error) {
	return nil, nil
}
