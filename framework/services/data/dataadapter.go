package data

import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"reflect"
	"strings"
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
	CONF_FIELD_ORDERBY             = "orderby"
	CONF_DATA_TOTALRECS            = "totalrecords"
	CONF_DATA_RECSRETURNED         = "records"
	CONF_SVC_LOOKUPSVC             = "lookupsvc"
	CONF_SVC_LOOKUP_FIELD          = "lookupfield"
)

func init() {
	objects.RegisterObject(CONF_DATAADAPTER_SERVICES, createDataAdapterFactory, nil)
}

type DataAdapterFactory struct {
	DataStore       data.DataComponent
	dataServiceName string
}

func createDataAdapterFactory() interface{} {
	return &DataAdapterFactory{}
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

type dataAdapterService struct {
	name          string
	method        string
	svcfunc       core.ServiceFunc
	conf          config.Config
	fac           *DataAdapterFactory
	lookupSvcName string
	lookupSvc     data.DataComponent
	lookupField   string
	DataStore     data.DataComponent
}

func newDataAdapterService(ctx core.ServerContext, name string, method string, fac *DataAdapterFactory) (*dataAdapterService, error) {
	ds := &dataAdapterService{name: name, fac: fac, method: method}
	//exported methods
	switch method {
	case CONF_SVC_GET:
		ds.svcfunc = ds.GETBYID
	case CONF_SVC_PUT:
		ds.svcfunc = ds.PUT
	case CONF_SVC_GETMULTIPLE:
		ds.svcfunc = ds.GETMULTI
	case CONF_SVC_COUNT:
		ds.svcfunc = ds.COUNT
	case CONF_SVC_SAVE:
		ds.svcfunc = ds.SAVE
	case CONF_SVC_JOIN:
		ds.svcfunc = ds.JOIN
	case CONF_SVC_DELETE:
		ds.svcfunc = ds.DELETE
	case CONF_SVC_DELETEALL:
		ds.svcfunc = ds.DELETEALL
	case CONF_SVC_SELECT:
		ds.svcfunc = ds.SELECT
	case CONF_SVC_UPSERT:
		ds.svcfunc = ds.UPSERT
	case CONF_SVC_UPDATE:
		ds.svcfunc = ds.UPDATE
	case CONF_SVC_PUTMULTIPLE:
		ds.svcfunc = ds.PUTMULTIPLE
	case CONF_SVC_UPDATEMULTIPLE:
		ds.svcfunc = ds.UPDATEMULTIPLE
	case CONF_SVC_GETMULTIPLE_SELECTIDS:
		ds.svcfunc = ds.GETMULTI_SELECTIDS
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
	ctx = ctx.SubContext("Data adapter initialize")
	ds.conf = conf
	if ds.method == CONF_SVC_GETMULTIPLE_SELECTIDS {
		lookupSvcName, ok := conf.GetString(CONF_SVC_LOOKUPSVC)
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "arg", CONF_SVC_LOOKUPSVC)
		}
		lookupField, ok := conf.GetString(CONF_SVC_LOOKUP_FIELD)
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "arg", CONF_SVC_LOOKUP_FIELD)
		}
		ds.lookupSvcName = lookupSvcName
		ds.lookupField = lookupField
	}
	return nil
}

func (ds *dataAdapterService) Start(ctx core.ServerContext) error {
	ctx = ctx.SubContext("Data adapter start")
	ds.DataStore = ds.fac.DataStore
	if ds.method == CONF_SVC_GETMULTIPLE_SELECTIDS {
		lookupSvcInt, err := ctx.GetService(ds.lookupSvcName)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		lookupSvc, ok := lookupSvcInt.(data.DataComponent)
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG)
		}
		ds.lookupSvc = lookupSvc
	}
	return nil
}

func (ds *dataAdapterService) Invoke(ctx core.RequestContext) error {
	return ds.svcfunc(ctx)
}

func (es *dataAdapterService) GETBYID(ctx core.RequestContext) error {
	ctx = ctx.SubContext("GETBYID")
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
		return nil
	} else {
		return errors.WrapError(ctx, err)
	}
}

func (es *dataAdapterService) GETMULTI(ctx core.RequestContext) error {
	ctx = ctx.SubContext("GETMULTI")
	idsstr, ok := ctx.GetString(CONF_DATA_IDS)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_DATA_IDS)
	}
	ids := strings.Split(idsstr, ",")
	orderBy, _ := ctx.GetString(CONF_FIELD_ORDERBY)
	result, err := es.DataStore.GetMulti(ctx, ids, orderBy)
	if err == nil {
		log.Logger.Trace(ctx, "Returning results ", "result", result)
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, result, nil))
		return nil
	} else {
		return errors.WrapError(ctx, err)
	}
}

func (es *dataAdapterService) JOIN(ctx core.RequestContext) error {
	ctx = ctx.SubContext("JOIN")
	retdata, _, totalrecs, recsreturned, err := es.selectMethod(ctx, es.DataStore)
	lookupids := make([]string, len(retdata))
	for ind, item := range retdata {
		entVal := reflect.ValueOf(item).Elem()
		f := entVal.FieldByName(es.lookupField)
		if !f.IsValid() {
			return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG)
		}
		lookupids[ind] = f.String()
	}
	log.Logger.Trace(ctx, "JOIN: Looking up ids", "ids", lookupids)
	result, err := es.DataStore.GetMultiHash(ctx, lookupids)
	if err == nil {
		for i, id := range lookupids {
			stor := retdata[i]
			item, ok := result[id]
			if ok {
				stor.Join(item)
			}
		}
		requestinfo := make(map[string]interface{}, 2)
		requestinfo[CONF_DATA_RECSRETURNED] = recsreturned
		requestinfo[CONF_DATA_TOTALRECS] = totalrecs
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, result, requestinfo))
		return nil
	} else {
		return errors.WrapError(ctx, err)
	}
}

func (es *dataAdapterService) GETMULTI_SELECTIDS(ctx core.RequestContext) error {
	ctx = ctx.SubContext("GETMULTI_SELECTIDS")
	retdata, _, totalrecs, recsreturned, err := es.selectMethod(ctx, es.lookupSvc)
	lookupids := make([]string, len(retdata))
	for ind, item := range retdata {
		entVal := reflect.ValueOf(item).Elem()
		f := entVal.FieldByName(es.lookupField)
		if !f.IsValid() {
			return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG)
		}
		lookupids[ind] = f.String()
	}
	log.Logger.Trace(ctx, "GETMULTI_SELECTIDS: Looking up ids", "ids", lookupids)
	hashmap, _ := ctx.GetBool("hashmap")
	var result interface{}
	if hashmap {
		result, err = es.DataStore.GetMultiHash(ctx, lookupids)
	} else {
		result, err = es.DataStore.GetMulti(ctx, lookupids, "")
	}
	if err == nil {
		requestinfo := make(map[string]interface{}, 2)
		requestinfo[CONF_DATA_RECSRETURNED] = recsreturned
		requestinfo[CONF_DATA_TOTALRECS] = totalrecs
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, result, requestinfo))
		return nil
	} else {
		return errors.WrapError(ctx, err)
	}

}

func (es *dataAdapterService) selectMethod(ctx core.RequestContext, datastore data.DataComponent) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	pagesize, _ := ctx.GetInt(data.DATA_PAGESIZE)
	pagenum, _ := ctx.GetInt(data.DATA_PAGENUM)

	var argsMap map[string]interface{}

	body := ctx.GetRequest().(*map[string]interface{})
	argsMap = *body
	log.Logger.Trace(ctx, "select", "argsMap", argsMap, "pagesize", pagesize, "pagenum", pagenum)
	orderBy, _ := ctx.GetString(CONF_FIELD_ORDERBY)
	condition, err := datastore.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		return nil, nil, -1, -1, errors.WrapError(ctx, err)
	}
	return datastore.Get(ctx, condition, pagesize, pagenum, "", orderBy)
}

func (es *dataAdapterService) SELECT(ctx core.RequestContext) error {
	ctx = ctx.SubContext("SELECT")
	retdata, _, totalrecs, recsreturned, err := es.selectMethod(ctx, es.DataStore)
	if err == nil {
		requestinfo := make(map[string]interface{}, 2)
		requestinfo[CONF_DATA_RECSRETURNED] = recsreturned
		requestinfo[CONF_DATA_TOTALRECS] = totalrecs
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, retdata, requestinfo))
		return nil
	} else {
		return errors.WrapError(ctx, err)
	}
}

func (es *dataAdapterService) COUNT(ctx core.RequestContext) error {
	ctx = ctx.SubContext("COUNT")
	body := ctx.GetRequest().(*map[string]interface{})
	argsMap := *body
	condition, err := es.DataStore.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	count, err := es.DataStore.Count(ctx, condition)
	if err == nil {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, count, nil))
		return nil
	} else {
		return errors.WrapError(ctx, err)
	}
}

func (es *dataAdapterService) SAVE(ctx core.RequestContext) error {
	ctx = ctx.SubContext("SAVE")
	ent := ctx.GetRequest()
	stor := ent.(data.Storable)
	err := es.DataStore.Save(ctx, stor)
	if err == nil {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, stor.GetId(), nil))
		return nil
	} else {
		return errors.WrapError(ctx, err)
	}
}

func (es *dataAdapterService) PUT(ctx core.RequestContext) error {
	ctx = ctx.SubContext("PUT")
	id, ok := ctx.GetString(CONF_DATA_ID)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_DATA_ID)
	}
	ent := ctx.GetRequest()
	stor := ent.(data.Storable)
	err := es.DataStore.Put(ctx, id, stor)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (es *dataAdapterService) DELETE(ctx core.RequestContext) error {
	ctx = ctx.SubContext("DELETE")
	id, ok := ctx.GetString(CONF_DATA_ID)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_DATA_ID)
	}
	err := es.DataStore.Delete(ctx, id)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (es *dataAdapterService) DELETEALL(ctx core.RequestContext) error {
	ctx = ctx.SubContext("DELETEALL")
	body := ctx.GetRequest().(*map[string]interface{})
	argsMap := *body
	condition, err := es.DataStore.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	retval, err := es.DataStore.DeleteAll(ctx, condition)
	if err == nil {
		ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, retval, nil))
		return nil
	} else {
		return errors.WrapError(ctx, err)
	}
}

func (es *dataAdapterService) UPSERT(ctx core.RequestContext) error {
	ctx = ctx.SubContext("UPSERT")
	body := ctx.GetRequest().(*map[string]interface{})
	vals := *body
	condition, err := es.DataStore.CreateCondition(ctx, data.FIELDVALUE, vals["condition"].(map[string]interface{}))
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	_, err = es.DataStore.Upsert(ctx, condition, vals["value"].(map[string]interface{}))
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (es *dataAdapterService) UPDATE(ctx core.RequestContext) error {
	ctx = ctx.SubContext("UPDATE")
	id, ok := ctx.GetString(CONF_DATA_ID)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_DATA_ID)
	}
	body := ctx.GetRequest().(*map[string]interface{})
	vals := *body
	err := es.DataStore.Update(ctx, id, vals)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (es *dataAdapterService) PUTMULTIPLE(ctx core.RequestContext) error {
	ctx = ctx.SubContext("PUTMULTIPLE")
	arr := ctx.GetRequest()
	log.Logger.Trace(ctx, "Collection ", "arr", arr)
	storables, _, err := data.CastToStorableCollection(arr)
	if err != nil {
		return err
	}
	err = es.DataStore.PutMulti(ctx, storables)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (es *dataAdapterService) UPDATEMULTIPLE(ctx core.RequestContext) error {
	ctx = ctx.SubContext("UPDATEMULTIPLE")
	body := ctx.GetRequest().(*map[string]interface{})
	vals := *body
	ids, ok := vals["ids"]
	if !ok {
		log.Logger.Error(ctx, "Missing argument", "Name", "ids")
		ctx.SetResponse(core.StatusBadRequestResponse)
		return nil
	}
	log.Logger.Info(ctx, "Value of ids", "IDs", ids, "Type", reflect.TypeOf(ids))
	idsArr, ok := ids.([]interface{})
	if !ok {
		log.Logger.Error(ctx, "Bad argument", "Name", "ids")
		ctx.SetResponse(core.StatusBadRequestResponse)
		return nil
	}
	stringIds := make([]string, len(idsArr))
	for i, val := range idsArr {
		stringIds[i], ok = val.(string)
		if !ok {
			log.Logger.Error(ctx, "Bad argument")
			ctx.SetResponse(core.StatusBadRequestResponse)
			return nil
		}
	}
	data, ok := vals["data"]
	if !ok {
		log.Logger.Error(ctx, "Missing argument", "Name", "data")
		ctx.SetResponse(core.StatusBadRequestResponse)
		return nil
	}
	updatesMap, ok := data.(map[string]interface{})
	if !ok {
		log.Logger.Error(ctx, "Bad argument", "Name", "data")
		ctx.SetResponse(core.StatusBadRequestResponse)
		return nil
	}
	err := es.DataStore.UpdateMulti(ctx, stringIds, updatesMap)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}
