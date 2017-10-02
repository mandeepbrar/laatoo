package main

import (
	"fmt"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"strings"
)

const (
	DATA_ADAPTER_MODULE      = "DataAdapterModule"
	DATA_ADAPTER_INSTANCE    = "instance"
	CONF_DATASERVICE_FACTORY = "dataservicefactory"
	CONF_SERVICEFACTORY      = "factory"
	CONF_PARENT_CHANNEL      = "parent"
	CHANNEL_SERVICE          = "service"
	REST_METHOD              = "method"
	REST_GET                 = "GET"
	REST_POST                = "POST"
	REST_DELETE              = "DELETE"
	REST_PUT                 = "PUT"
	REST_PATH                = "path"
	REST_PARAMS              = "paramvalues"
	CHANNEL_DATAOBJECT       = "dataobject"

	REST_STATIC    = "staticvalues"
	SERVICE_METHOD = "servicemethod"
	MIDDLEWARE     = "middleware"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_DATAADAPTER_SERVICES, Object: DataAdapterFactory{}},
		core.PluginComponent{Name: DATA_ADAPTER_MODULE, Object: DataAdapterModule{}}}
}

type DataAdapterModule struct {
	core.Module
	object             string
	factory            string //underlying factory for creating data service
	instance           string
	adapterfacName     string
	adapterdataSvcName string
	middleware         string
	parentChannel      string
}

/*
func (adapter *DataAdapterModule) Describe(ctx core.ServerContext) {
	adapter.AddStringConfiguration(ctx, CONF_DATASERVICE_FACTORY)
	adapter.AddStringConfiguration(ctx, data.CONF_DATA_OBJECT)
	adapter.AddStringConfigurations(ctx, []string{DATA_ADAPTER_INSTANCE, MIDDLEWARE, CONF_PARENT_CHANNEL}, []string{"", "", "root"})
}*/

func (adapter *DataAdapterModule) Initialize(ctx core.ServerContext, conf config.Config) error {
	ctx = ctx.SubContext("Starting data adapter module")
	adapter.factory, _ = adapter.GetStringConfiguration(ctx, CONF_DATASERVICE_FACTORY)
	adapter.object, _ = adapter.GetStringConfiguration(ctx, data.CONF_DATA_OBJECT)
	adapter.instance, _ = adapter.GetStringConfiguration(ctx, DATA_ADAPTER_INSTANCE)
	adapter.middleware, _ = adapter.GetStringConfiguration(ctx, MIDDLEWARE)
	adapter.parentChannel, _ = adapter.GetStringConfiguration(ctx, CONF_PARENT_CHANNEL)

	adapter.adapterfacName = adapter.createName(ctx, "factory")
	adapter.adapterdataSvcName = adapter.createName(ctx, "dataservice")

	if adapter.instance == "" {
		adapter.instance = adapter.object
	}
	return nil
}

func (adapter *DataAdapterModule) Factories(ctx core.ServerContext) map[string]config.Config {
	ctx = ctx.SubContext("Getting factories for module ")
	facs := make(map[string]config.Config)

	factory := make(config.GenericConfig)
	factory[CONF_SERVICEFACTORY] = CONF_DATAADAPTER_SERVICES
	factory[CONF_DATAADAPTER_DATA_SVC] = adapter.adapterdataSvcName
	factory[MIDDLEWARE] = adapter.middleware
	facs[adapter.adapterfacName] = factory

	log.Error(ctx, "Returned factories", "facs", facs)
	return facs
}

func (adapter *DataAdapterModule) Services(ctx core.ServerContext) map[string]config.Config {
	ctx = ctx.SubContext("Getting services for module ")
	svcs := make(map[string]config.Config)

	dataService := make(config.GenericConfig)
	dataService[CONF_SERVICEFACTORY] = adapter.factory
	dataService[data.CONF_DATA_OBJECT] = adapter.object
	svcs[adapter.adapterdataSvcName] = dataService

	/*dataSvcName := adapter.createName(ctx, "dataservice")
	dataService := make(config.GenericConfig)
	dataService[CONF_SERVICEFACTORY] = adapter.factory
	svcs[dataSvcName] = dataService*/

	getSvcName := adapter.createName(ctx, "get")
	getService := make(config.GenericConfig)
	getService[CONF_DATAADAPTER_DATA_SVC] = adapter.adapterdataSvcName
	getService[CONF_SERVICEFACTORY] = adapter.adapterfacName
	getService[SERVICE_METHOD] = CONF_SVC_GET
	svcs[getSvcName] = getService

	selectSvcName := adapter.createName(ctx, "select")
	selectService := make(config.GenericConfig)
	selectService[CONF_DATAADAPTER_DATA_SVC] = adapter.adapterdataSvcName
	selectService[CONF_SERVICEFACTORY] = adapter.adapterfacName
	selectService[SERVICE_METHOD] = CONF_SVC_SELECT
	selectService[CHANNEL_DATAOBJECT] = config.OBJECTTYPE_STRINGMAP
	selectService["queryparams"] = []string{"pagesize", "pagenum"}

	svcs[selectSvcName] = selectService

	log.Error(ctx, "Returned services", "svcs", svcs)
	return svcs
}

func (adapter *DataAdapterModule) createName(ctx core.ServerContext, svc string) string {
	if adapter.instance != "" {
		return fmt.Sprintf("dataadapter.%s.%s", svc, adapter.instance)
	} else {
		return fmt.Sprintf("dataadapter.%s.%s", svc, adapter.object)
	}
}

func (adapter *DataAdapterModule) Rules(ctx core.ServerContext) map[string]config.Config {
	return nil
}
func (adapter *DataAdapterModule) Channels(ctx core.ServerContext) map[string]config.Config {

	ctx = ctx.SubContext("Getting channels for module ")
	chans := make(map[string]config.Config)

	objectChann := make(config.GenericConfig)
	objectChann[CONF_PARENT_CHANNEL] = adapter.parentChannel
	objectChann[REST_PATH] = fmt.Sprintf("/%s", strings.ToLower(adapter.instance))
	chans[adapter.instance] = objectChann

	getRestChannName := adapter.createName(ctx, "get")
	getRestChann := make(config.GenericConfig)
	getRestChann[CHANNEL_SERVICE] = adapter.createName(ctx, "get")
	getRestChann[CONF_PARENT_CHANNEL] = adapter.instance
	getRestChann[REST_METHOD] = REST_GET
	getRestChann[REST_PATH] = "/:id"
	getparams := make(config.GenericConfig)
	getparams["id"] = "id"
	getRestChann[REST_PARAMS] = getparams
	getstaticvals := make(config.GenericConfig)
	getstaticvals["permission"] = "View " + adapter.instance
	getRestChann[REST_STATIC] = getstaticvals
	chans[getRestChannName] = getRestChann

	selectRestChannName := adapter.createName(ctx, "select")
	selectRestChann := make(config.GenericConfig)
	selectRestChann[CHANNEL_SERVICE] = adapter.createName(ctx, "select")
	selectRestChann[CONF_PARENT_CHANNEL] = adapter.instance
	selectRestChann[REST_METHOD] = REST_POST
	selectRestChann[REST_PATH] = "/view"
	selectstaticvals := make(config.GenericConfig)
	selectstaticvals["permission"] = "View " + adapter.instance
	selectRestChann[REST_STATIC] = selectstaticvals
	chans[selectRestChannName] = selectRestChann

	log.Error(ctx, "Returned channels", "chans", chans)
	return chans
}
func (adapter *DataAdapterModule) Tasks(ctx core.ServerContext) map[string]config.Config {
	return nil
}
