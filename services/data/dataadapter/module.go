package main

import (
	"fmt"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
)

const (
	DATA_ADAPTER_MODULE      = "DataAdapterModule"
	DATA_ADAPTER_INSTANCE    = "instance"
	CONF_DATASERVICE_FACTORY = "dataservicefactory"
	CONF_SERVICEFACTORY      = "factory"
	SERVICE_METHOD           = "servicemethod"
)

func Manifest() []core.PluginComponent {
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
}

func (adapter *DataAdapterModule) Initialize(ctx core.ServerContext) error {
	adapter.AddStringConfiguration(ctx, CONF_DATASERVICE_FACTORY)
	adapter.AddStringConfiguration(ctx, data.CONF_DATA_OBJECT)
	adapter.AddStringConfigurations(ctx, []string{DATA_ADAPTER_INSTANCE}, []string{""})
	return nil
}

func (adapter *DataAdapterModule) Start(ctx core.ServerContext) error {
	adapter.factory, _ = adapter.GetStringConfiguration(ctx, CONF_DATASERVICE_FACTORY)
	adapter.object, _ = adapter.GetStringConfiguration(ctx, data.CONF_DATA_OBJECT)
	adapter.instance, _ = adapter.GetStringConfiguration(ctx, DATA_ADAPTER_INSTANCE)

	adapter.adapterfacName = adapter.createName(ctx, "factory")
	adapter.adapterdataSvcName = adapter.createName(ctx, "dataservice")

	return nil
}

func (adapter *DataAdapterModule) Factories(ctx core.ServerContext) map[string]config.Config {
	ctx = ctx.SubContext("Getting factories for module ")
	facs := make(map[string]config.Config)

	factory := make(config.GenericConfig)
	factory[CONF_SERVICEFACTORY] = CONF_DATAADAPTER_SERVICES
	factory[CONF_DATAADAPTER_DATA_SVC] = adapter.adapterdataSvcName
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

	log.Error(ctx, "Returned services", "svcs", svcs)
	return svcs
}

func (adapter *DataAdapterModule) createName(ctx core.ServerContext, svc string) string {
	svcName := fmt.Sprintf("dataadapter.%s.%s", svc, adapter.object)
	if adapter.instance != "" {
		svcName = fmt.Sprint(svcName, ".", adapter.instance)
	}
	return svcName
}

func (adapter *DataAdapterModule) Rules(ctx core.ServerContext) map[string]config.Config {
	return nil
}
func (adapter *DataAdapterModule) Channels(ctx core.ServerContext) map[string]config.Config {
	return nil
}
func (adapter *DataAdapterModule) Tasks(ctx core.ServerContext) map[string]config.Config {
	return nil
}
