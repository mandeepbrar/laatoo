package data

import (
	"laatoo/core/objects"
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
	CONF_DATA_TOTALRECS       = "totalrecords"
	CONF_DATA_RECSRETURNED    = "records"
)

func init() {
	objects.RegisterObject(CONF_DATAADAPTER_SERVICES, createDataAdapterFactory, nil)
}

type DataAdapterFactory struct {
	DataStore       data.DataService
	dataServiceName string
}

func createDataAdapterFactory(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	return &DataAdapterFactory{}, nil
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
func (es *DataAdapterFactory) CreateService(ctx core.ServerContext, name string, method string) (core.Service, error) {
	return newDataAdapterService(ctx, name, method, es)
}

//The services start serving when this method is called
func (es *DataAdapterFactory) Start(ctx core.ServerContext) error {
	dataSvc, err := ctx.GetService(es.dataServiceName)
	if err != nil {
		return errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, err, "Name", es.dataServiceName)
	}
	es.DataStore = dataSvc.(data.DataService)
	return nil
}

type dataAdapterService struct {
	name      string
	svcfunc   core.ServiceFunc
	conf      config.Config
	fac       *DataAdapterFactory
	DataStore data.DataService
}

func newDataAdapterService(ctx core.ServerContext, name string, method string, fac *DataAdapterFactory) (*dataAdapterService, error) {
	ds := &dataAdapterService{name: name, fac: fac}
	//exported methods
	switch method {
	case CONF_SVC_GET:
		ds.svcfunc = ds.GETBYID
	case CONF_SVC_PUT:
		ds.svcfunc = ds.PUT
	case CONF_SVC_GETMULTIPLE:
		ds.svcfunc = ds.GETMULTI
	case CONF_SVC_SAVE:
		ds.svcfunc = ds.SAVE
	case CONF_SVC_DELETE:
		ds.svcfunc = ds.DELETE
	case CONF_SVC_SELECT:
		ds.svcfunc = ds.SELECT
	case CONF_SVC_UPDATE:
		ds.svcfunc = ds.UPDATE
	case CONF_SVC_PUTMULTIPLE:
		ds.svcfunc = ds.PUTMULTIPLE
	case CONF_SVC_UPDATEMULTIPLE:
		ds.svcfunc = ds.UPDATEMULTIPLE
	default:
		return nil, nil
	}
	//cache, _ := conf.GetBool(CONF_DATA_CACHEABLE)
	//ds.cache = cache
	log.Logger.Trace(ctx, "Created Data Adapter service", "Svc Name", name, "Method", method)
	return ds, nil
}

func (ds *dataAdapterService) GetName() string {
	return ds.name
}

func (ds *dataAdapterService) Initialize(ctx core.ServerContext, conf config.Config) error {
	ds.conf = conf
	ds.DataStore = ds.fac.DataStore
	return nil
}

func (ds *dataAdapterService) Start(ctx core.ServerContext) error {
	return nil
}

func (ds *dataAdapterService) Invoke(ctx core.RequestContext) error {
	return ds.svcfunc(ctx)
}

func (es *dataAdapterService) GETBYID(ctx core.RequestContext) error {
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

func (es *dataAdapterService) GETMULTI(ctx core.RequestContext) error {
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

func (es *dataAdapterService) SELECT(ctx core.RequestContext) error {
	var err error
	pagesize, _ := ctx.GetInt(data.DATA_PAGESIZE)
	pagenum, _ := ctx.GetInt(data.DATA_PAGENUM)

	var argsMap map[string]interface{}

	body := ctx.GetRequest().(*map[string]interface{})
	argsMap = *body
	orderBy, _ := ctx.GetString(CONF_FIELD_ORDERBY)
	condition, err := es.DataStore.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	retdata, totalrecs, recsreturned, err := es.DataStore.Get(ctx, condition, pagesize, pagenum, "", orderBy)
	if err == nil {
		requestinfo := make(map[string]interface{}, 2)
		requestinfo[CONF_DATA_RECSRETURNED] = recsreturned
		requestinfo[CONF_DATA_TOTALRECS] = totalrecs
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, retdata, requestinfo))
	}
	return err
}

func (es *dataAdapterService) SAVE(ctx core.RequestContext) error {
	ent := ctx.GetRequest()
	err := es.DataStore.Save(ctx, ent.(data.Storable))
	if err == nil {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, nil, nil))
	}
	return err
}

func (es *dataAdapterService) PUT(ctx core.RequestContext) error {
	id, ok := ctx.GetString(CONF_DATA_ID)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_DATA_ID)
	}
	ent := ctx.GetRequest()
	stor := ent.(data.Storable)
	err := es.DataStore.Put(ctx, id, stor)
	if err == nil {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, nil, nil))
	}
	return err
}

func (es *dataAdapterService) DELETE(ctx core.RequestContext) error {
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

func (es *dataAdapterService) UPDATE(ctx core.RequestContext) error {
	id, ok := ctx.GetString(CONF_DATA_ID)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_DATA_ID)
	}
	body := ctx.GetRequest().(*map[string]interface{})
	vals := *body
	err := es.DataStore.Update(ctx, id, vals)
	if err == nil {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, nil, nil))
	}
	return err
}

func (es *dataAdapterService) PUTMULTIPLE(ctx core.RequestContext) error {
	arr := ctx.GetRequest()
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

func (es *dataAdapterService) UPDATEMULTIPLE(ctx core.RequestContext) error {
	idsstr, ok := ctx.GetString(CONF_DATA_IDS)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_DATA_IDS)
	}
	ids := strings.Split(idsstr, ",")
	body := ctx.GetRequest().(*map[string]interface{})
	vals := *body
	err := es.DataStore.UpdateMulti(ctx, ids, vals)
	if err == nil {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, nil, nil))
	}
	return err
}