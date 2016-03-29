package entities

import (
	"encoding/json"
	"fmt"
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
	CONF_ENTITY_SERVICES            = "entity_services"
	CONF_ENTITY                     = "entity"
	CONF_ENTITY_DATA_SVC            = "data_svc"
	CONF_ENTITY_ID                  = "id"
	CONF_ENTITY_IDS                 = "ids"
	CONF_SOFT_DELETE                = "softdelete"
	CONF_SVC_GETENTITY              = "GETENTITY"
	CONF_SVC_PUTENTITY              = "PUTENTITY"
	CONF_SVC_PUTMULTIPLEENTITIES    = "PUTMULTIPLEENTITIES"
	CONF_SVC_GETMULTIPLEENTITIES    = "GETMULTIPLEENTITIES"
	CONF_SVC_SAVEENTITY             = "SAVEENTITY"
	CONF_SVC_DELETEENTITY           = "DELETEENTITY"
	CONF_SVC_SELECT                 = "SELECTENTITIES"
	CONF_SVC_UPDATE                 = "UPDATEENTITY"
	CONF_SVC_UPDATEMULTIPLEENTITIES = "UPDATEMULTIPLEENTITIES"
	CONF_GETENTITY                  = "getentity"
	CONF_PUTENTITY                  = "putentity"
	CONF_PUTMULTIPLEENTITIES        = "putmultipleentities"
	CONF_GETMULTIPLEENTITIES        = "getmultipleentities"
	CONF_SAVEENTITY                 = "saveentity"
	CONF_DELETEENTITY               = "deleteentity"
	CONF_SELECT                     = "selectentities"
	CONF_UPDATE                     = "updateentity"
	CONF_UPDATEMULTIPLEENTITIES     = "updatemultipleentities"
	CONF_FIELD_ORDERBY              = "orderby"
	DATA_SELECT_ARGS                = "args"
)

//Initialize service, register provider with laatoo
func init() {
	registry.RegisterServiceFactoryProvider(CONF_ENTITY_SERVICES, NewEntityServiceFactory)
}

//Environment hosting an application
type EntityServiceFactory struct {
	EntityName              string
	dataServiceName         string
	DataStore               data.DataService
	entityCreator           core.ObjectCreator
	entityCollectionCreator core.ObjectCollectionCreator
	cache                   bool
}

//factory method returns the service object to the environment
func NewEntityServiceFactory(ctx core.ServerContext, conf config.Config) (core.ServiceFactory, error) {
	log.Logger.Info(ctx, "Creating entity service")

	svc := &EntityServiceFactory{cache: true}

	//get a map of all the entities
	entity, ok := conf.GetString(CONF_ENTITY)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Configuration", CONF_ENTITY)
	}

	svc.EntityName = entity

	entitydatasvc, ok := conf.GetString(CONF_ENTITY_DATA_SVC)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Configuration", CONF_ENTITY_DATA_SVC)
	}
	svc.dataServiceName = entitydatasvc

	entityCreator, err := registry.GetObjectCreator(ctx, entity)
	if err != nil {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "Could not get Object creator for", entity)
	}

	entityCollectionCreator, err := registry.GetObjectCollectionCreator(ctx, entity)
	if err != nil {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "Could not get Object collection creator for", entity)
	}

	svc.entityCreator = entityCreator
	svc.entityCollectionCreator = entityCollectionCreator
	return svc, nil
}

//Create the services configured for factory.
func (es *EntityServiceFactory) CreateService(ctx core.ServerContext, name string, conf config.Config) (core.Service, error) {
	var svcfunc core.ServiceFunc
	//exported methods
	switch name {
	case CONF_SVC_GETENTITY:
		svcfunc = es.GETBYID
	case CONF_SVC_PUTENTITY:
		svcfunc = es.PUTENTITY
	case CONF_SVC_GETMULTIPLEENTITIES:
		svcfunc = es.GETMULTI
	case CONF_SVC_SAVEENTITY:
		svcfunc = es.SAVEENTITY
	case CONF_SVC_DELETEENTITY:
		svcfunc = es.DELETEENTITY
	case CONF_SVC_SELECT:
		svcfunc = es.SELECT
	case CONF_SVC_UPDATE:
		svcfunc = es.UPDATEENTITY
	case CONF_SVC_PUTMULTIPLEENTITIES:
		svcfunc = es.PUTMULTIPLEENTITIES
	case CONF_SVC_UPDATEMULTIPLEENTITIES:
		svcfunc = es.UPDATEMULTIPLEENTITIES
	}
	if svcfunc != nil {
		return services.NewService(ctx, svcfunc, conf), nil
	}
	/*//for internal consumption
	switch name {
	case CONF_GETENTITY:
		return es.GetById, nil
	case CONF_PUTENTITY:
		return es.PutEntity, nil
	case CONF_GETMULTIPLEENTITIES:
		return es.GetMulti, nil
	case CONF_SAVEENTITY:
		return es.SaveEntity, nil
	case CONF_DELETEENTITY:
		return es.DeleteEntity, nil
	case CONF_SELECT:
		return es.Select, nil
	case CONF_UPDATE:
		return es.UpdateEntity, nil
	case CONF_PUTMULTIPLEENTITIES:
		return es.PutMultipleEntities, nil
	case CONF_UPDATEMULTIPLEENTITIES:
		return es.UpdateMultipleEntities, nil
	}*/
	return nil, nil
}

func (es *EntityServiceFactory) GETBYID(ctx core.RequestContext) error {
	id, ok := ctx.GetString(CONF_ENTITY_ID)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_ENTITY_ID)
	}
	result, err := es.GetById(ctx, id)
	if err == nil {
		ctx.SetResponse(result)
	}
	return err
}

func (es *EntityServiceFactory) GetById(ctx core.RequestContext, id string) (data.Storable, error) {
	log.Logger.Trace(ctx, "Getting Entity", "entity", es.EntityName, "id", id)
	cachekey := es.getCacheKey(id)
	if es.cache {
		ent, err := es.entityCreator(ctx, nil)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		ok := ctx.GetFromCache(cachekey, ent)
		if ok {
			return ent.(data.Storable), nil
		}
	}
	ent, err := es.DataStore.GetById(ctx, id)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	} else {
		if es.cache {
			ctx.PutInCache(cachekey, ent)
		}
	}
	return ent, nil
}

func (es *EntityServiceFactory) GETMULTI(ctx core.RequestContext) error {
	idsstr, ok := ctx.GetString(CONF_ENTITY_IDS)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_ENTITY_IDS)
	}
	ids := strings.Split(idsstr, ",")
	orderBy, _ := ctx.GetString(CONF_FIELD_ORDERBY)
	result, err := es.GetMulti(ctx, ids, orderBy)
	if err == nil {
		ctx.SetResponse(result)
	}
	return err
}

func (es *EntityServiceFactory) GetMulti(ctx core.RequestContext, ids []string, orderBy string) (map[string]data.Storable, error) {
	log.Logger.Trace(ctx, "Getting multiple entities", "entity", es.EntityName)
	ents, err := es.DataStore.GetMulti(ctx, ids, orderBy)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	} else {
		for k, v := range ents {
			if v != nil && es.cache {
				ctx.PutInCache(es.getCacheKey(k), v)
			}
		}
	}
	return ents, nil
}

func (es *EntityServiceFactory) SELECT(ctx core.RequestContext) error {
	var err error
	pagesize, _ := ctx.GetInt(data.DATA_PAGESIZE)
	pagenum, _ := ctx.GetInt(data.DATA_PAGENUM)

	args, argsok := ctx.GetString(DATA_SELECT_ARGS)
	var argsMap map[string]interface{}

	if argsok {
		argsbyt := []byte(args)
		if err := json.Unmarshal(argsbyt, &argsMap); err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	softDeletes, ok := ctx.GetString("Deleted")
	if ok {
		argsMap["Deleted"] = (softDeletes == "true")
	}
	orderBy, _ := ctx.GetString(CONF_FIELD_ORDERBY)
	retdata, _, _, err := es.Select(ctx, argsMap, pagesize, pagenum, orderBy)
	if err == nil {
		ctx.SetResponse(retdata)
	}
	return err
}

func (es *EntityServiceFactory) Select(ctx core.RequestContext, argsMap map[string]interface{}, pagesize int, pagenum int, orderBy string) (dataToReturn []data.Storable, totalrecs int, recsreturned int, err error) {
	condition, err := es.DataStore.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		return nil, -1, -1, errors.WrapError(ctx, err)
	}
	return es.DataStore.Get(ctx, condition, pagesize, pagenum, "", orderBy)
}

/*view, err := newEntitiesView(ctx, map[string]interface{}{"entity": es.EntityName})
	if err != nil {
		return nil, err
	}
	orderBy := ""
	orderByInt, ok := params[CONF_FIELD_ORDERBY]
	if ok {
		orderBy = orderByInt.(string)
	}
	entities, _, _, err := view.getData(ctx, es.DataStore, params, -1, -1, orderBy)
	return entities, err
}*/

func (es *EntityServiceFactory) SAVEENTITY(ctx core.RequestContext) error {
	ent := ctx.GetRequestBody()
	result, err := es.SaveEntity(ctx, ent.(data.Storable))
	if err == nil {
		ctx.SetResponse(result)
	}
	return err
}

func (es *EntityServiceFactory) SaveEntity(ctx core.RequestContext, ent data.Storable) (string, error) {
	data.Audit(ctx, ent)
	err := es.DataStore.Save(ctx, ent)
	if err != nil {
		return "", errors.WrapError(ctx, err)
	}
	id := ent.GetId()
	if es.cache {
		es.invalidateCache(ctx, id)
	}
	return id, nil
}

func (es *EntityServiceFactory) PUTENTITY(ctx core.RequestContext) error {
	ent := ctx.GetRequestBody()
	stor := ent.(data.Storable)
	result, err := es.PutEntity(ctx, stor.GetId(), stor)
	if err == nil {
		ctx.SetResponse(result)
	}
	return err
}

func (es *EntityServiceFactory) PutEntity(ctx core.RequestContext, id string, ent data.Storable) (interface{}, error) {
	data.Audit(ctx, ent)
	err := es.DataStore.Put(ctx, id, ent)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	if es.cache {
		es.invalidateCache(ctx, ent.GetId())
	}
	log.Logger.Trace(ctx, "Saved entity")
	return nil, nil
}

func (es *EntityServiceFactory) DELETEENTITY(ctx core.RequestContext) error {
	id, ok := ctx.GetString(CONF_ENTITY_ID)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_ENTITY_ID)
	}
	softdelete := false
	sdconfig, sok := ctx.GetString(CONF_SOFT_DELETE)
	if sok {
		softdelete = (sdconfig == "true")
	}
	result, err := es.DeleteEntity(ctx, id, softdelete)
	if err == nil {
		ctx.SetResponse(result)
	}
	return err
}

func (es *EntityServiceFactory) DeleteEntity(ctx core.RequestContext, id string, softdelete bool) (interface{}, error) {
	log.Logger.Debug(ctx, "Deleting entity", "ID", id)
	err := es.DataStore.Delete(ctx, id, softdelete)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	if es.cache {
		es.invalidateCache(ctx, id)
	}
	return nil, nil
}

func (es *EntityServiceFactory) UPDATEENTITY(ctx core.RequestContext) error {
	id, ok := ctx.GetString(CONF_ENTITY_ID)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_ENTITY_ID)
	}
	vals := ctx.GetRequestBody().(map[string]interface{})
	result, err := es.UpdateEntity(ctx, id, vals)
	if err == nil {
		ctx.SetResponse(result)
	}
	return err
}

func (es *EntityServiceFactory) UpdateEntity(ctx core.RequestContext, id string, newVals map[string]interface{}) (interface{}, error) {
	log.Logger.Trace(ctx, "Updating Entity", "entity", es.EntityName, "id", id)
	data.Audit(ctx, newVals)
	err := es.DataStore.Update(ctx, id, newVals)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	if es.cache {
		es.invalidateCache(ctx, id)
	}
	return nil, nil
}

func (es *EntityServiceFactory) PUTMULTIPLEENTITIES(ctx core.RequestContext) error {
	arr := ctx.GetRequestBody()
	storables, err := data.CastToStorableCollection(arr)
	result, err := es.PutMultipleEntities(ctx, storables)
	if err == nil {
		ctx.SetResponse(result)
	}
	return err
}

func (es *EntityServiceFactory) PutMultipleEntities(ctx core.RequestContext, arr []data.Storable) (interface{}, error) {
	log.Logger.Trace(ctx, "Saving bulk entities")
	length := len(arr)
	ids := make([]string, length)
	for i := 0; i < length; i++ {
		ids[i] = arr[i].GetId()
		data.Audit(ctx, arr[i])
	}
	err := es.DataStore.PutMulti(ctx, ids, arr)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	for i := 0; i < length; i++ {
		if es.cache {
			es.invalidateCache(ctx, ids[i])
		}
	}
	return nil, err
}

func (es *EntityServiceFactory) UPDATEMULTIPLEENTITIES(ctx core.RequestContext) error {
	idsstr, ok := ctx.GetString(CONF_ENTITY_IDS)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG, "argument", CONF_ENTITY_IDS)
	}
	ids := strings.Split(idsstr, ",")
	vals := ctx.GetRequestBody().(map[string]interface{})
	result, err := es.UpdateMultipleEntities(ctx, ids, vals)
	if err == nil {
		ctx.SetResponse(result)
	}
	return err
}

func (es *EntityServiceFactory) UpdateMultipleEntities(ctx core.RequestContext, ids []string, newVals map[string]interface{}) (interface{}, error) {
	log.Logger.Trace(ctx, "Updating multiple Entities", "entity", es.EntityName)
	data.Audit(ctx, newVals)
	err := es.DataStore.UpdateMulti(ctx, ids, newVals)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	if es.cache {
		for _, id := range ids {
			es.invalidateCache(ctx, id)
		}
	}
	return nil, nil
}

//The services start serving when this method is called
func (es *EntityServiceFactory) StartServices(ctx core.ServerContext) error {
	dataSvc, err := ctx.GetService(es.dataServiceName)
	if err != nil {
		return errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, err, "Name", es.dataServiceName)
	}
	es.DataStore = dataSvc.(data.DataService)
	return nil
}

func (es *EntityServiceFactory) invalidateCache(ctx core.RequestContext, id string) {
	if es.cache {
		ctx.DeleteFromCache(es.getCacheKey(id))
	}
}

func (es *EntityServiceFactory) getCacheKey(id string) string {
	return fmt.Sprintf("Entity_%s_%s", es.EntityName, id)
}
