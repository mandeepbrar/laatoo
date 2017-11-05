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

	factory := ctx.CreateConfig()
	factory.Set(ctx, CONF_SERVICEFACTORY, CONF_DATAADAPTER_SERVICES)
	factory.Set(ctx, CONF_DATAADAPTER_DATA_SVC, adapter.adapterdataSvcName)
	factory.Set(ctx, MIDDLEWARE, adapter.middleware)
	facs[adapter.adapterfacName] = factory

	log.Trace(ctx, "Returned factories", "facs", facs)
	return facs
}

func (adapter *DataAdapterModule) Services(ctx core.ServerContext) map[string]config.Config {
	ctx = ctx.SubContext("Getting services for module ")
	svcs := make(map[string]config.Config)

	dataService := ctx.CreateConfig()
	dataService.Set(ctx, CONF_SERVICEFACTORY, adapter.factory)
	dataService.Set(ctx, data.CONF_DATA_OBJECT, adapter.object)
	svcs[adapter.adapterdataSvcName] = dataService

	/*dataSvcName := adapter.createName(ctx, "dataservice")
	dataService := ctx.CreateConfig()
	dataService[CONF_SERVICEFACTORY] = adapter.factory
	svcs[dataSvcName] = dataService*/

	getSvcName := adapter.createName(ctx, "get")
	getService := ctx.CreateConfig()
	getService.Set(ctx, CONF_DATAADAPTER_DATA_SVC, adapter.adapterdataSvcName)
	getService.Set(ctx, CONF_SERVICEFACTORY, adapter.adapterfacName)
	getService.Set(ctx, SERVICE_METHOD, CONF_SVC_GET)
	svcs[getSvcName] = getService

	selectSvcName := adapter.createName(ctx, "select")
	selectService := ctx.CreateConfig()
	selectService.Set(ctx, CONF_DATAADAPTER_DATA_SVC, adapter.adapterdataSvcName)
	selectService.Set(ctx, CONF_SERVICEFACTORY, adapter.adapterfacName)
	selectService.Set(ctx, SERVICE_METHOD, CONF_SVC_SELECT)
	selectService.Set(ctx, "queryparams", []string{"pagesize", "pagenum"})
	selectService.Set(ctx, CHANNEL_DATAOBJECT, config.OBJECTTYPE_STRINGMAP)

	svcs[selectSvcName] = selectService

	saveSvcName := adapter.createName(ctx, "save")
	saveService := ctx.CreateConfig()
	saveService.Set(ctx, CONF_DATAADAPTER_DATA_SVC, adapter.adapterdataSvcName)
	saveService.Set(ctx, CONF_SERVICEFACTORY, adapter.adapterfacName)
	saveService.Set(ctx, SERVICE_METHOD, CONF_SVC_SAVE)
	saveService.Set(ctx, CHANNEL_DATAOBJECT, adapter.object)
	svcs[saveSvcName] = saveService

	updateSvcName := adapter.createName(ctx, "update")
	updateService := ctx.CreateConfig()
	updateService.Set(ctx, CONF_DATAADAPTER_DATA_SVC, adapter.adapterdataSvcName)
	updateService.Set(ctx, CONF_SERVICEFACTORY, adapter.adapterfacName)
	updateService.Set(ctx, SERVICE_METHOD, CONF_SVC_UPDATE)
	updateService.Set(ctx, CHANNEL_DATAOBJECT, adapter.object)
	svcs[updateSvcName] = updateService

	log.Trace(ctx, "Returned services", "svcs", svcs)
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

	objectChann := ctx.CreateConfig()
	objectChann.Set(ctx, CONF_PARENT_CHANNEL, adapter.parentChannel)
	objectChann.Set(ctx, REST_PATH, fmt.Sprintf("/%s", strings.ToLower(adapter.instance)))
	chans[adapter.instance] = objectChann

	getRestChannName := adapter.createName(ctx, "get")
	getRestChann := ctx.CreateConfig()
	getRestChann.Set(ctx, CHANNEL_SERVICE, adapter.createName(ctx, "get"))
	getRestChann.Set(ctx, CONF_PARENT_CHANNEL, adapter.instance)
	getRestChann.Set(ctx, REST_METHOD, REST_GET)
	getRestChann.Set(ctx, REST_PATH, "/:id")
	getparams := ctx.CreateConfig()
	getparams.Set(ctx, "id", "id")
	getRestChann.Set(ctx, REST_PARAMS, getparams)
	getstaticvals := ctx.CreateConfig()
	getstaticvals.Set(ctx, "permission", "View "+adapter.instance)
	getRestChann.Set(ctx, REST_STATIC, getstaticvals)
	chans[getRestChannName] = getRestChann

	selectRestChannName := adapter.createName(ctx, "select")
	selectRestChann := ctx.CreateConfig()
	selectRestChann.Set(ctx, CHANNEL_SERVICE, adapter.createName(ctx, "select"))
	selectRestChann.Set(ctx, CONF_PARENT_CHANNEL, adapter.instance)
	selectRestChann.Set(ctx, REST_METHOD, REST_POST)
	selectRestChann.Set(ctx, REST_PATH, "/view")
	selectstaticvals := ctx.CreateConfig()
	selectstaticvals.Set(ctx, "permission", "View "+adapter.instance)
	selectRestChann.Set(ctx, REST_STATIC, selectstaticvals)
	chans[selectRestChannName] = selectRestChann

	saveRestChannName := adapter.createName(ctx, "save")
	saveRestChann := ctx.CreateConfig()
	saveRestChann.Set(ctx, CHANNEL_SERVICE, adapter.createName(ctx, "save"))
	saveRestChann.Set(ctx, CONF_PARENT_CHANNEL, adapter.instance)
	saveRestChann.Set(ctx, REST_METHOD, REST_POST)
	saveRestChann.Set(ctx, REST_PATH, "")
	chans[saveRestChannName] = saveRestChann

	updateRestChannName := adapter.createName(ctx, "update")
	updateRestChann := ctx.CreateConfig()
	updateRestChann.Set(ctx, CHANNEL_SERVICE, adapter.createName(ctx, "update"))
	updateRestChann.Set(ctx, CONF_PARENT_CHANNEL, adapter.instance)
	updateRestChann.Set(ctx, REST_METHOD, REST_PUT)
	updateRestChann.Set(ctx, REST_PATH, "/:id")
	updateparams := ctx.CreateConfig()
	updateparams.Set(ctx, "id", "id")
	updateRestChann.Set(ctx, REST_PARAMS, updateparams)
	chans[updateRestChannName] = updateRestChann

	log.Trace(ctx, "Returned channels", "chans", chans)
	return chans
}
func (adapter *DataAdapterModule) Tasks(ctx core.ServerContext) map[string]config.Config {
	return nil
}
