package main

import (
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"
	"strings"
)

const (
	CONF_DATAADAPTER_FAC_NAME    = "dataadapter.DataAdapterFactory"
	DATA_ADAPTER_MODULE          = "DataAdapterModule"
	DATA_ADAPTER_INSTANCE        = "instance"
	CONF_DATASERVICE_FACTORY     = "dataservicefactory"
	CONF_EMBEDDED_SEARCH         = "embedded_doc_search"
	CONF_SERVICEFACTORY          = "factory"
	CONF_PARENT_CHANNEL          = "parent"
	CONF_CREATE_DATA_SERVICE     = "create_dataservice"
	CONF_CREATE_DATA_ADAPTER_FAC = "create_dataadapterfactory"
	CONF_CREATE_GET_SERVICE      = "create_getservice"
	CONF_CREATE_UPDATE_SERVICE   = "create_updateservice"
	CONF_CREATE_SAVE_SERVICE     = "create_saveservice"
	CONF_CREATE_SELECT_SERVICE   = "create_selectservice"
	CONF_CREATE_OBJECT_CHANNEL   = "create_objectchannel"
	BODY_PARAM_NAME              = "bodyparamname"
	CHANNEL_SERVICE              = "service"
	REST_METHOD                  = "method"
	REST_GET                     = "GET"
	REST_POST                    = "POST"
	REST_DELETE                  = "DELETE"
	REST_PUT                     = "PUT"
	REST_PATH                    = "path"
	REST_PARAMS                  = "paramvalues"
	CHANNEL_DATAOBJECT           = "dataobject"

	REST_STATIC    = "staticvalues"
	SERVICE_METHOD = "servicemethod"
	MIDDLEWARE     = "middleware"
)

type DataAdapterModule struct {
	core.Module
	object             string
	factory            string //underlying factory for creating data service
	instance           string
	adapterfacName     string
	adapterdataSvcName string
	middleware         string
	parentChannel      string
	embeddedDocSearch  bool

	dataservice        bool
	dataadapterfactory bool
	getservice         bool
	saveservice        bool
	updateservice      bool
	selectservice      bool
	objectchannel      bool
}

/*
func (adapter *DataAdapterModule) Describe(ctx core.ServerContext) {
	adapter.AddStringConfiguration(ctx, CONF_DATASERVICE_FACTORY)
	adapter.AddStringConfiguration(ctx, data.CONF_DATA_OBJECT)
	adapter.AddStringConfigurations(ctx, []string{DATA_ADAPTER_INSTANCE, MIDDLEWARE, CONF_PARENT_CHANNEL}, []string{"", "", "root"})
}*/
func (adapter *DataAdapterModule) MetaInfo(ctx core.ServerContext) map[string]interface{} {
	return map[string]interface{}{
		"services":  []string{},
		"factories": []string{},
		"channels":  []string{},
	}
}

func (adapter *DataAdapterModule) Initialize(ctx core.ServerContext, conf config.Config) error {
	ctx = ctx.SubContext("Starting data adapter module")
	adapter.factory, _ = adapter.GetStringConfiguration(ctx, CONF_DATASERVICE_FACTORY)
	adapter.object, _ = adapter.GetStringConfiguration(ctx, data.CONF_DATA_OBJECT)
	adapter.instance, _ = adapter.GetStringConfiguration(ctx, DATA_ADAPTER_INSTANCE)
	adapter.middleware, _ = adapter.GetStringConfiguration(ctx, MIDDLEWARE)
	adapter.parentChannel, _ = adapter.GetStringConfiguration(ctx, CONF_PARENT_CHANNEL)
	adapter.embeddedDocSearch, _ = adapter.GetBoolConfiguration(ctx, CONF_EMBEDDED_SEARCH)
	adapter.dataservice, _ = adapter.GetBoolConfiguration(ctx, CONF_CREATE_DATA_SERVICE)
	adapter.getservice, _ = adapter.GetBoolConfiguration(ctx, CONF_CREATE_GET_SERVICE)
	adapter.saveservice, _ = adapter.GetBoolConfiguration(ctx, CONF_CREATE_SAVE_SERVICE)
	adapter.updateservice, _ = adapter.GetBoolConfiguration(ctx, CONF_CREATE_UPDATE_SERVICE)
	adapter.selectservice, _ = adapter.GetBoolConfiguration(ctx, CONF_CREATE_SELECT_SERVICE)
	adapter.dataadapterfactory, _ = adapter.GetBoolConfiguration(ctx, CONF_CREATE_DATA_ADAPTER_FAC)
	adapter.objectchannel, _ = adapter.GetBoolConfiguration(ctx, CONF_CREATE_OBJECT_CHANNEL)

	adapter.adapterfacName = adapter.createName(ctx, "factory")
	adapter.adapterdataSvcName = adapter.createName(ctx, "dataservice")

	if adapter.instance == "" {
		adapter.instance = adapter.object
	}
	log.Error(ctx, "Data adapter. Services to create", "data service", adapter.dataservice, " adapter factory", adapter.dataadapterfactory, "select service", adapter.selectservice,
		"object channel", adapter.objectchannel, "get service", adapter.getservice, "save service ", adapter.saveservice, "update service", adapter.updateservice, "conf", conf)
	return nil
}

func (adapter *DataAdapterModule) Factories(ctx core.ServerContext) map[string]config.Config {
	ctx = ctx.SubContext("Getting factories for module ")
	facs := make(map[string]config.Config)

	if adapter.dataadapterfactory {
		factory := ctx.CreateConfig()
		factory.Set(ctx, CONF_SERVICEFACTORY, CONF_DATAADAPTER_FAC_NAME)
		factory.Set(ctx, CONF_DATAADAPTER_DATA_SVC, adapter.adapterdataSvcName)
		factory.Set(ctx, MIDDLEWARE, adapter.middleware)
		facs[adapter.adapterfacName] = factory
	}

	log.Trace(ctx, "Returned factories", "facs", facs)
	return facs
}

func (adapter *DataAdapterModule) Services(ctx core.ServerContext) map[string]config.Config {
	ctx = ctx.SubContext("Getting services for module ")
	svcs := make(map[string]config.Config)

	if adapter.dataservice {
		dataService := ctx.CreateConfig()
		dataService.Set(ctx, CONF_SERVICEFACTORY, adapter.factory)
		dataService.Set(ctx, data.CONF_DATA_OBJECT, adapter.object)
		dataService.Set(ctx, CONF_EMBEDDED_SEARCH, adapter.embeddedDocSearch)
		svcs[adapter.adapterdataSvcName] = dataService
	}

	/*dataSvcName := adapter.createName(ctx, "dataservice")
	dataService := ctx.CreateConfig()
	dataService[CONF_SERVICEFACTORY] = adapter.factory
	svcs[dataSvcName] = dataService*/

	if adapter.getservice {
		getSvcName := adapter.createName(ctx, "get")
		getService := ctx.CreateConfig()
		getService.Set(ctx, CONF_DATAADAPTER_DATA_SVC, adapter.adapterdataSvcName)
		getService.Set(ctx, CONF_SERVICEFACTORY, adapter.adapterfacName)
		getService.Set(ctx, SERVICE_METHOD, CONF_SVC_GET)
		svcs[getSvcName] = getService
	}

	if adapter.selectservice {
		selectSvcName := adapter.createName(ctx, "select")
		selectService := ctx.CreateConfig()
		selectService.Set(ctx, CONF_DATAADAPTER_DATA_SVC, adapter.adapterdataSvcName)
		selectService.Set(ctx, CONF_SERVICEFACTORY, adapter.adapterfacName)
		selectService.Set(ctx, SERVICE_METHOD, CONF_SVC_SELECT)
		selectService.Set(ctx, "queryparams", []string{"pagesize", "pagenum"})
		selectService.Set(ctx, CHANNEL_DATAOBJECT, config.OBJECTTYPE_STRINGMAP)

		svcs[selectSvcName] = selectService
	}

	if adapter.saveservice {
		saveSvcName := adapter.createName(ctx, "save")
		saveService := ctx.CreateConfig()
		saveService.Set(ctx, CONF_DATAADAPTER_DATA_SVC, adapter.adapterdataSvcName)
		saveService.Set(ctx, CONF_SERVICEFACTORY, adapter.adapterfacName)
		saveService.Set(ctx, SERVICE_METHOD, CONF_SVC_SAVE)
		saveService.Set(ctx, CHANNEL_DATAOBJECT, adapter.object)
		svcs[saveSvcName] = saveService
	}

	if adapter.updateservice {
		updateSvcName := adapter.createName(ctx, "update")
		updateService := ctx.CreateConfig()
		updateService.Set(ctx, CONF_DATAADAPTER_DATA_SVC, adapter.adapterdataSvcName)
		updateService.Set(ctx, CONF_SERVICEFACTORY, adapter.adapterfacName)
		updateService.Set(ctx, SERVICE_METHOD, CONF_SVC_UPDATE)
		updateService.Set(ctx, CHANNEL_DATAOBJECT, adapter.object)
		svcs[updateSvcName] = updateService
	}

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

	if adapter.objectchannel {
		objectChann := ctx.CreateConfig()
		objectChann.Set(ctx, CONF_PARENT_CHANNEL, adapter.parentChannel)
		objectChann.Set(ctx, REST_PATH, fmt.Sprintf("/%s", strings.ToLower(adapter.instance)))
		chans[adapter.instance] = objectChann
	}

	if adapter.getservice {
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
	}

	if adapter.selectservice {
		selectRestChannName := adapter.createName(ctx, "select")
		selectRestChann := ctx.CreateConfig()
		selectRestChann.Set(ctx, CHANNEL_SERVICE, adapter.createName(ctx, "select"))
		selectRestChann.Set(ctx, CONF_PARENT_CHANNEL, adapter.instance)
		selectRestChann.Set(ctx, REST_METHOD, REST_POST)
		selectRestChann.Set(ctx, REST_PATH, "/view")
		selectRestChann.Set(ctx, BODY_PARAM_NAME, "argsMap")
		selectstaticvals := ctx.CreateConfig()
		selectstaticvals.Set(ctx, "permission", "View "+adapter.instance)
		selectRestChann.Set(ctx, REST_STATIC, selectstaticvals)
		chans[selectRestChannName] = selectRestChann
	}

	if adapter.saveservice {
		saveRestChannName := adapter.createName(ctx, "save")
		saveRestChann := ctx.CreateConfig()
		saveRestChann.Set(ctx, CHANNEL_SERVICE, adapter.createName(ctx, "save"))
		saveRestChann.Set(ctx, CONF_PARENT_CHANNEL, adapter.instance)
		saveRestChann.Set(ctx, REST_METHOD, REST_POST)
		saveRestChann.Set(ctx, BODY_PARAM_NAME, "object")
		saveRestChann.Set(ctx, REST_PATH, "")
		chans[saveRestChannName] = saveRestChann
	}

	if adapter.updateservice {
		updateRestChannName := adapter.createName(ctx, "update")
		updateRestChann := ctx.CreateConfig()
		updateRestChann.Set(ctx, CHANNEL_SERVICE, adapter.createName(ctx, "update"))
		updateRestChann.Set(ctx, CONF_PARENT_CHANNEL, adapter.instance)
		updateRestChann.Set(ctx, REST_METHOD, REST_PUT)
		updateRestChann.Set(ctx, REST_PATH, "/:id")
		updateRestChann.Set(ctx, BODY_PARAM_NAME, "argsMap")
		updateparams := ctx.CreateConfig()
		updateparams.Set(ctx, "id", "id")
		updateRestChann.Set(ctx, REST_PARAMS, updateparams)
		chans[updateRestChannName] = updateRestChann
	}

	log.Trace(ctx, "Returned channels", "chans", chans)
	return chans
}
func (adapter *DataAdapterModule) Tasks(ctx core.ServerContext) map[string]config.Config {
	return nil
}
