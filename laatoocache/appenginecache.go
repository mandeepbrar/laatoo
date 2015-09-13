// +build appengine

package laatoopubsub

import (
	"laatoocore"
	//	"laatoosdk/errors"
	"appengine/memcache"
	"laatoosdk/log"
	"laatoosdk/service"
	"time"
)

type AppEngineCacheService struct {
	name    string
	context service.ServiceContext
}

const (
	CONF_APPENGINECACHE_NAME        = "appengine_cache"
	CONF_APPENGINE_CONNECTIONSTRING = "server"
	CONF_APPENGINE_DATABASE         = "db"
)

func init() {
	laatoocore.RegisterObjectProvider(CONF_APPENGINECACHE_NAME, AppEngineCacheServiceFactory)
}

func AppEngineCacheServiceFactory(conf map[string]interface{}) (interface{}, error) {
	log.Logger.Infof("Creating redis cache service ")
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
func (svc *AppEngineCacheService) Serve() error {
	return nil
}

func (svc *AppEngineCacheService) Delete(key string) error {
	return nil
}

func (svc *AppEngineCacheService) PutObject(key string, val interface{}) error {
	return memcache.Set(c, &memcache.Item{ Key: key, Value: []byte(val))
}

func (svc *AppEngineCacheService) GetObject(key string) (interface{}, error) {
	item, err := memcache.Get(c, key)
	if(err != nil ) {
		return nil, err
	} else {
		return item.Value, nil
	}
}

//Execute method
func (svc *AppEngineCacheService) Execute(name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}
