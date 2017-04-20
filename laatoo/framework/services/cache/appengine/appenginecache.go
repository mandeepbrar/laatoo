package appengine

import (
	"bytes"
	"encoding/gob"
	"laatoo/framework/core/objects"
	"laatoo/framework/services/cache/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"

	"google.golang.org/appengine/memcache"
	//	"laatoo/sdk/errors"
	//"laatoo/sdk/log"
	//"time"
)

type AppengineCacheFactory struct {
}

const (
	CONF_APPENGINECACHE_FACTORY = "appenginecache"
	CONF_APPENGINECACHE_SVC     = "cache"
)

func init() {
	objects.Register(CONF_APPENGINECACHE_FACTORY, AppengineCacheFactory{})
}

//Create the services configured for factory.
func (af *AppengineCacheFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	if method == CONF_APPENGINECACHE_SVC {
		return &AppengineCacheService{}, nil
	}
	return nil, nil
}

func (ds *AppengineCacheFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//The services start serving when this method is called
func (ds *AppengineCacheFactory) Start(ctx core.ServerContext) error {
	return nil
}

type AppengineCacheService struct {
}

func (svc *AppengineCacheService) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

func (svc *AppengineCacheService) Delete(ctx core.RequestContext, bucket string, key string) error {
	return memcache.Delete(ctx.GetAppengineContext(), key)
}

func (svc *AppengineCacheService) PutObject(ctx core.RequestContext, bucket string, key string, val interface{}) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(val)
	if err != nil {
		return err
	}
	return memcache.Set(ctx.GetAppengineContext(), &memcache.Item{Key: key, Value: buf.Bytes()})
}

func (svc *AppengineCacheService) GetObject(ctx core.RequestContext, bucket string, key string, val interface{}) error {
	item, err := memcache.Get(ctx.GetAppengineContext(), key)

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

func (svc *AppengineCacheService) GetMulti(ctx core.RequestContext, bucket string, keys []string, val map[string]interface{}) {
	_, err := memcache.GetMulti(ctx.GetAppengineContext(), keys)
	if err != nil {
		return
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
	return
}
func (svc *AppengineCacheService) Increment(ctx core.RequestContext, bucket string, key string) error {
	_, err := memcache.Increment(ctx.GetAppengineContext(), common.GetCacheKey(bucket, key), 1, 0)
	return err
}
func (svc *AppengineCacheService) Decrement(ctx core.RequestContext, bucket string, key string) error {
	_, err := memcache.Increment(ctx.GetAppengineContext(), common.GetCacheKey(bucket, key), -1, 0)
	return err
}
func (svc *AppengineCacheService) Invoke(ctx core.RequestContext) error {
	return nil
}

func (svc *AppengineCacheService) Start(ctx core.ServerContext) error {
	return nil
}
