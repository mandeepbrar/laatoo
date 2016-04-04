package data

import (
	"laatoo/core/registry"
	"laatoo/core/services"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/data"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"strings"
)

const (
	CONF_DATAADAPTER_SERVICES = "dataadapter"
	CONF_DATAADAPTER_DATA_SVC = "data_svc"
	CONF_DATA_ID              = "id"
	CONF_DATA_IDS             = "ids"
	CONF_SVC_GET              = "GET"
	CONF_SVC_PUT              = "PUT"
	CONF_SVC_PUTMULTIPLE      = "PUTMULTIPLE"
	CONF_SVC_GETMULTIPLE      = "GETMULTIPLE"
	CONF_SVC_SAVE             = "SAVE"
	CONF_SVC_DELETE           = "DELETE"
	CONF_SVC_SELECT           = "SELECT"
	CONF_SVC_UPDATE           = "UPDATE"
	CONF_SVC_UPDATEMULTIPLE   = "UPDATEMULTIPLE"
	CONF_FIELD_ORDERBY        = "orderby"
)

//Initialize service, register provider with laatoo
func init() {
	registry.RegisterServiceFactoryProvider(CONF_DATAADAPTER_SERVICES, NewDataAdapterFactory)
}

type DataAdapterFactory struct {
	DataStore       data.DataService
	dataServiceName string
}

//factory method returns the service object to the application
func NewDataAdapterFactory(ctx core.ServerContext, conf config.Config) (core.ServiceFactory, error) {
	svc := &DataAdapterFactory{}
	datasvc, ok := conf.GetString(CONF_DATAADAPTER_DATA_SVC)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Configuration", CONF_DATAADAPTER_DATA_SVC)
	}
	svc.dataServiceName = datasvc
	log.Logger.Debug(ctx, "Created Data Adapter factory", "Svc Name", datasvc)
	return svc, nil
}

//Create the services configured for factory.
func (es *DataAdapterFactory) CreateService(ctx core.ServerContext, name string, conf config.Config) (core.Service, error) {
	var svcfunc core.ServiceFunc
	//exported methods
	switch name {
	case CONF_SVC_GET:
		svcfunc = es.GETBYID
	case CONF_SVC_PUT:
		svcfunc = es.PUT
	case CONF_SVC_GETMULTIPLE:
		svcfunc = es.GETMULTI
	case CONF_SVC_SAVE:
		svcfunc = es.SAVE
	case CONF_SVC_DELETE:
		svcfunc = es.DELETE
	case CONF_SVC_SELECT:
		svcfunc = es.SELECT
	case CONF_SVC_UPDATE:
		svcfunc = es.UPDATE
	case CONF_SVC_PUTMULTIPLE:
		svcfunc = es.PUTMULTIPLE
	case CONF_SVC_UPDATEMULTIPLE:
		svcfunc = es.UPDATEMULTIPLE
	}
	if svcfunc != nil {
		log.Logger.Trace(ctx, "Created Data service", "Svc Name", es.dataServiceName, "Method", name)
		return services.NewService(ctx, svcfunc, conf), nil
	}
	return nil, nil
}

func (es *DataAdapterFactory) GETBYID(ctx core.RequestContext) error {
	id, ok := ctx.GetString(CONF_DATA_ID)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_DATA_ID)
	}
	result, err := es.DataStore.GetById(ctx, id)
	if err == nil {
		if result == nil {
			ctx.SetResponse(core.StatusNotFoundResponse)
		} else {
			ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, result, nil))
		}
	}
	return err
}

func (es *DataAdapterFactory) GETMULTI(ctx core.RequestContext) error {
	idsstr, ok := ctx.GetString(CONF_DATA_IDS)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_DATA_IDS)
	}
	ids := strings.Split(idsstr, ",")
	orderBy, _ := ctx.GetString(CONF_FIELD_ORDERBY)
	result, err := es.DataStore.GetMulti(ctx, ids, orderBy)
	if err == nil {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, result, nil))
	}
	return err
}

func (es *DataAdapterFactory) SELECT(ctx core.RequestContext) error {
	var err error
	pagesize, _ := ctx.GetInt(data.DATA_PAGESIZE)
	pagenum, _ := ctx.GetInt(data.DATA_PAGENUM)

	var argsMap map[string]interface{}

	body := ctx.GetRequestBody().(*map[string]interface{})
	argsMap = *body

	orderBy, _ := ctx.GetString(CONF_FIELD_ORDERBY)
	condition, err := es.DataStore.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	retdata, _, _, err := es.DataStore.Get(ctx, condition, pagesize, pagenum, "", orderBy)
	if err == nil {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, retdata, nil))
	}
	return err
}

func (es *DataAdapterFactory) SAVE(ctx core.RequestContext) error {
	ent := ctx.GetRequestBody()
	err := es.DataStore.Save(ctx, ent.(data.Storable))
	if err == nil {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, nil, nil))
	}
	return err
}

func (es *DataAdapterFactory) PUT(ctx core.RequestContext) error {
	id, ok := ctx.GetString(CONF_DATA_ID)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_DATA_ID)
	}
	ent := ctx.GetRequestBody()
	stor := ent.(data.Storable)
	err := es.DataStore.Put(ctx, id, stor)
	if err == nil {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, nil, nil))
	}
	return err
}

func (es *DataAdapterFactory) DELETE(ctx core.RequestContext) error {
	id, ok := ctx.GetString(CONF_DATA_ID)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_DATA_ID)
	}
	err := es.DataStore.Delete(ctx, id)
	if err == nil {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, nil, nil))
	}
	return err
}

func (es *DataAdapterFactory) UPDATE(ctx core.RequestContext) error {
	id, ok := ctx.GetString(CONF_DATA_ID)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_DATA_ID)
	}
	body := ctx.GetRequestBody().(*map[string]interface{})
	vals := *body
	err := es.DataStore.Update(ctx, id, vals)
	if err == nil {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, nil, nil))
	}
	return err
}

func (es *DataAdapterFactory) PUTMULTIPLE(ctx core.RequestContext) error {
	arr := ctx.GetRequestBody()
	storables, err := data.CastToStorableCollection(arr)
	if err != nil {
		return err
	}
	err = es.DataStore.PutMulti(ctx, storables)
	if err == nil {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, nil, nil))
	}
	return err
}

func (es *DataAdapterFactory) UPDATEMULTIPLE(ctx core.RequestContext) error {
	idsstr, ok := ctx.GetString(CONF_DATA_IDS)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_DATA_IDS)
	}
	ids := strings.Split(idsstr, ",")
	body := ctx.GetRequestBody().(*map[string]interface{})
	vals := *body
	err := es.DataStore.UpdateMulti(ctx, ids, vals)
	if err == nil {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, nil, nil))
	}
	return err
}

//The services start serving when this method is called
func (es *DataAdapterFactory) StartServices(ctx core.ServerContext) error {
	dataSvc, err := ctx.GetService(es.dataServiceName)
	if err != nil {
		return errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, err, "Name", es.dataServiceName)
	}
	es.DataStore = dataSvc.(data.DataService)
	return nil
}
