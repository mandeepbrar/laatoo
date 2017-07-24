package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
)

const (
	CONF_DATAADAPTER_SERVICES      = "dataadapter"
	CONF_DATAADAPTER_DATA_SVC      = "data_svc"
	CONF_DATA_ID                   = "id"
	CONF_DATA_IDS                  = "ids"
	CONF_SVC_GET                   = "GET"
	CONF_SVC_COUNT                 = "COUNT"
	CONF_SVC_PUT                   = "PUT"
	CONF_SVC_PUTMULTIPLE           = "PUTMULTIPLE"
	CONF_SVC_GETMULTIPLE           = "GETMULTIPLE"
	CONF_SVC_GETMULTIPLE_SELECTIDS = "GETMULTIPLE_SELECTIDS"
	CONF_SVC_SAVE                  = "SAVE"
	CONF_SVC_JOIN                  = "JOIN"
	CONF_SVC_DELETE                = "DELETE"
	CONF_SVC_DELETEALL             = "DELETEALL"
	CONF_SVC_SELECT                = "SELECT"
	CONF_SVC_UPSERT                = "UPSERT"
	CONF_SVC_UPDATE                = "UPDATE"
	CONF_SVC_UPDATEMULTIPLE        = "UPDATEMULTIPLE"
	CONF_SVC_UPDATE_WITH_STORABLE  = "UPDATE_WITH_STORABLE"
	CONF_FIELD_ORDERBY             = "orderby"
	CONF_DATA_TOTALRECS            = "totalrecords"
	CONF_DATA_RECSRETURNED         = "records"
	CONF_SVC_LOOKUPSVC             = "lookupsvc"
	CONF_SVC_LOOKUP_FIELD          = "lookupfield"
	CONF_SVC_UPDATE_FIELDS         = "updatefields"
)

func Manifest() []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_DATAADAPTER_SERVICES, Object: DataAdapterFactory{}}}
}

type DataAdapterFactory struct {
	DataStore       data.DataComponent
	dataServiceName string
}

func (es *DataAdapterFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	datasvc, ok := conf.GetString(CONF_DATAADAPTER_DATA_SVC)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Configuration", CONF_DATAADAPTER_DATA_SVC)
	}
	es.dataServiceName = datasvc
	return nil
}

//Create the services configured for factory.
func (es *DataAdapterFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	return newDataAdapterService(ctx, name, method, es)
}

//The services start serving when this method is called
func (es *DataAdapterFactory) Start(ctx core.ServerContext) error {
	dataSvc, err := ctx.GetService(es.dataServiceName)
	if err != nil {
		return errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, err, "Name", es.dataServiceName)
	}
	es.DataStore = dataSvc.(data.DataComponent)
	return nil
}

func newDataAdapterService(ctx core.ServerContext, name string, method string, fac *DataAdapterFactory) (core.Service, error) {
	switch method {
	case CONF_SVC_GET:
		return &getById{fac: fac, DataStore: fac.DataStore}, nil
	case CONF_SVC_PUT:
		return &put{fac: fac, DataStore: fac.DataStore}, nil
	case CONF_SVC_GETMULTIPLE:
		return &getMulti{fac: fac, DataStore: fac.DataStore}, nil
	case CONF_SVC_COUNT:
		return &count{fac: fac, DataStore: fac.DataStore}, nil
	case CONF_SVC_SAVE:
		return &save{fac: fac, DataStore: fac.DataStore}, nil
	case CONF_SVC_JOIN:
		return &join{fac: fac, DataStore: fac.DataStore}, nil
	case CONF_SVC_DELETE:
		return &deleteSvc{fac: fac, DataStore: fac.DataStore}, nil
	case CONF_SVC_DELETEALL:
		return &deleteAll{fac: fac, DataStore: fac.DataStore}, nil
	case CONF_SVC_SELECT:
		return &selectSvc{fac: fac, DataStore: fac.DataStore}, nil
	case CONF_SVC_UPSERT:
		return &upsert{fac: fac, DataStore: fac.DataStore}, nil
	case CONF_SVC_UPDATE:
		return &update{fac: fac, DataStore: fac.DataStore}, nil
	case CONF_SVC_PUTMULTIPLE:
		return &putmultiple{fac: fac, DataStore: fac.DataStore}, nil
	case CONF_SVC_UPDATEMULTIPLE:
		return &updatemultiple{fac: fac, DataStore: fac.DataStore}, nil
	case CONF_SVC_UPDATE_WITH_STORABLE:
		return &updateStorable{fac: fac, DataStore: fac.DataStore}, nil
	case CONF_SVC_GETMULTIPLE_SELECTIDS:
		return &getmulti_select{fac: fac, DataStore: fac.DataStore}, nil
	default:
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, "Wrong Service method", method)
	}
}
