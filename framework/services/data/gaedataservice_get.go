package data

import (
	"fmt"
	"laatoo/sdk/components/data"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"reflect"

	glctx "golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type gaeDatastoreCondition struct {
	operation data.ConditionType
	arg1      interface{}
	arg2      interface{}
	arg3      interface{}
}

func (svc *gaeDataService) GetById(ctx core.RequestContext, id string) (data.Storable, error) {
	appEngineContext := ctx.GetAppengineContext()
	log.Logger.Trace(ctx, "Getting object by id ", "id", id, "object", svc.object)

	//try cache if the object is cacheable
	if svc.cacheable {
		ent, err := svc.objectCreator(ctx, nil)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		ok := getFromCache(ctx, svc.object, id, ent)
		if ok {
			return ent.(data.Storable), nil
		}
	}

	object, err := svc.objectCreator(ctx, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
	err = datastore.Get(appEngineContext, key, object)
	if err != nil {
		if err.Error() == "not found" {
			return nil, nil
		}
		return nil, errors.RethrowError(ctx, DATA_ERROR_OPERATION, err, "ID", id)
	}
	stor := object.(data.Storable)
	if svc.postload {
		stor.PostLoad(ctx)
	}
	if svc.cacheable {
		putInCache(ctx, svc.object, id, stor)
	}
	return stor, nil
}

//Get multiple objects by id
func (svc *gaeDataService) GetMulti(ctx core.RequestContext, ids []string, orderBy string) (map[string]data.Storable, error) {
	lenids := len(ids)
	retVal := make(map[string]data.Storable, lenids)
	if lenids == 0 {
		return retVal, nil
	}
	appEngineContext := ctx.GetAppengineContext()

	results, err := svc.objectCollectionCreator(ctx, lenids, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	keys := make([]*datastore.Key, lenids)
	for ind, id := range ids {
		key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
		keys[ind] = key
	}
	/*if len(orderBy) > 0 {
		query = query.Order(orderBy)
	}*/
	log.Logger.Debug(ctx, "Geting object", "results", results, "ids", ids, "object", svc.object, "keys", keys, "type", reflect.ValueOf(results).Type())
	err = datastore.GetMulti(appEngineContext, keys, reflect.ValueOf(results).Elem().Interface())
	if err != nil {
		if _, ok := err.(appengine.MultiError); !ok {
			log.Logger.Debug(ctx, "Geting object", "err", err)
			return nil, errors.WrapError(ctx, err)
		}
	}
	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	resultStor, err := data.CastToStorableCollection(results)
	for _, stor := range resultStor {
		id := stor.GetId()
		retVal[id] = stor
		if svc.postload {
			stor.PostLoad(ctx)
		}
		if svc.cacheable {
			putInCache(ctx, svc.object, stor.GetId(), stor)
		}
	}
	for _, id := range ids {
		_, ok := retVal[id]
		if !ok {
			retVal[id] = nil
		}
	}
	return retVal, nil
}

func (svc *gaeDataService) GetList(ctx core.RequestContext, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, totalrecs int, recsreturned int, err error) {
	return svc.Get(ctx, nil, pageSize, pageNum, mode, orderBy) // resultStor, totalrecs, recsreturned, nil
}

func (svc *gaeDataService) Get(ctx core.RequestContext, queryCond interface{}, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, totalrecs int, recsreturned int, err error) {
	appEngineContext := ctx.GetAppengineContext()
	totalrecs = -1
	recsreturned = -1
	query := datastore.NewQuery(svc.collection)
	query, err = svc.processCondition(ctx, appEngineContext, query, queryCond)
	if err != nil {
		return nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}
	if pageSize > 0 {
		totalrecs, err = query.Limit(500).Count(appEngineContext)
		if err != nil {
			return nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
		}
		recsToSkip := (pageNum - 1) * pageSize
		query = query.Limit(pageSize).Offset(recsToSkip)
	}
	if len(orderBy) > 0 {
		query = query.Order(orderBy)
	}
	results, err := svc.objectCollectionCreator(ctx, 0, nil)
	if err != nil {
		return nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}
	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	_, err = query.GetAll(appEngineContext, results)
	log.Logger.Trace(ctx, "Returning multiple objects ", "conditions", queryCond, "objectType", svc.object, "collection", svc.collection)
	resultStor, err := data.CastToStorableCollection(results)
	if err != nil {
		return nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}
	recsreturned = len(resultStor)
	for _, stor := range resultStor {
		if svc.postload {
			stor.PostLoad(ctx)
		}
		if svc.cacheable {
			putInCache(ctx, svc.object, stor.GetId(), stor)
		}
	}
	if recsreturned > totalrecs {
		totalrecs = recsreturned
	}
	log.Logger.Trace(ctx, "Returning multiple objects ", "conditions", queryCond, "objectType", svc.object, "recsreturned", recsreturned)
	return resultStor, totalrecs, recsreturned, nil
}

//create condition for passing to data service
func (svc *gaeDataService) CreateCondition(ctx core.RequestContext, operation data.ConditionType, args ...interface{}) (interface{}, error) {
	switch operation {
	case data.MATCHANCESTOR:
		{
			if len(args) < 2 {
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG)
			}
			return &gaeDatastoreCondition{operation: operation, arg1: args[0]}, nil
		}
	case data.MATCHMULTIPLEVALUES:
		{
			if len(args) < 2 {
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG)
			}
			return &gaeDatastoreCondition{operation: operation, arg1: args[0], arg2: args[1]}, nil
		}
	case data.FIELDVALUE:
		{
			if len(args) < 1 {
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG)
			}
			return &gaeDatastoreCondition{operation: operation, arg1: args[0]}, nil
		}
	}
	return nil, nil
}

func (svc *gaeDataService) processCondition(ctx core.RequestContext, appEngineContext glctx.Context, query *datastore.Query, condition interface{}) (*datastore.Query, error) {
	if condition == nil {
		return query, nil
	}
	dqCondition := condition.(*gaeDatastoreCondition)
	switch dqCondition.operation {
	case data.MATCHANCESTOR:
		id, ok := dqCondition.arg1.(string)
		if !ok {
			return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG)
		}
		key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
		return query.Ancestor(key), nil
	case data.MATCHMULTIPLEVALUES:
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_NOT_IMPLEMENTED)
	case data.FIELDVALUE:
		queryCondMap, ok := dqCondition.arg1.(map[string]interface{})
		if ok {
			if svc.softdelete {
				queryCondMap["Deleted"] = false
			}
			for k, v := range queryCondMap {
				query = query.Filter(fmt.Sprintf("%s =", k), v)
			}
			return query, nil
		} else {
			return nil, errors.ThrowError(ctx, errors.CORE_ERROR_TYPE_MISMATCH)
		}

	}
	return query, nil
}
