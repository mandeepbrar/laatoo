package main

import (
	"fmt"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
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
	ctx = ctx.SubContext("GetById")
	appEngineContext := ctx.GetAppengineContext()
	log.Trace(ctx, "Getting object by id ", "id", id, "object", svc.Object)

	object, _ := ctx.CreateObject(svc.Object)

	key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
	err := datastore.Get(appEngineContext, key, object)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, nil
		}
		return nil, errors.RethrowError(ctx, data.DATA_ERROR_OPERATION, err, "ID", id)
	}
	stor := object.(data.Storable)
	if stor.IsDeleted() {
		return nil, nil
	}
	if svc.PostLoad {
		stor.PostLoad(ctx)
	}
	return stor, nil
}

//Get multiple objects by id
func (svc *gaeDataService) GetMulti(ctx core.RequestContext, ids []string, orderBy string) ([]data.Storable, error) {
	ctx = ctx.SubContext("GetMulti")
	results, err := svc.getMulti(ctx, ids, orderBy)
	if err != nil {
		return nil, err
	}
	if results == nil {
		return []data.Storable{}, nil
	}
	res, _, err := svc.postArrayGet(ctx, results)
	return res, err
}

func (svc *gaeDataService) GetMultiHash(ctx core.RequestContext, ids []string) (map[string]data.Storable, error) {
	ctx = ctx.SubContext("GetMultiHash")
	results, err := svc.getMulti(ctx, ids, "")
	if err != nil {
		return nil, err
	}
	if results == nil {
		return map[string]data.Storable{}, nil
	}
	resultStor, err := data.CastToStorableHash(results)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	for _, stor := range resultStor {
		svc.postLoad(ctx, stor)
	}
	return resultStor, nil
}

func (svc *gaeDataService) postArrayGet(ctx core.RequestContext, results interface{}) ([]data.Storable, []string, error) {
	resultStor, ids, err := data.CastToStorableCollection(ctx, results)
	if err != nil {
		return nil, nil, errors.WrapError(ctx, err)
	}

	for _, stor := range resultStor {
		svc.postLoad(ctx, stor)
	}
	return resultStor, ids, nil
}

func (svc *gaeDataService) postLoad(ctx core.RequestContext, stor data.Storable) error {
	if svc.PostLoad {
		stor.PostLoad(ctx)
	}
	return nil
}

func (svc *gaeDataService) getMulti(ctx core.RequestContext, ids []string, orderBy string) (interface{}, error) {
	lenids := len(ids)
	if lenids == 0 {
		return nil, nil
	}
	appEngineContext := ctx.GetAppengineContext()

	results, _ := ctx.CreateCollection(svc.Object, lenids)

	keys := make([]*datastore.Key, lenids)
	for ind, id := range ids {
		key := datastore.NewKey(appEngineContext, svc.collection, id, 0, nil)
		keys[ind] = key
	}
	/*if len(orderBy) > 0 {
		query = query.Order(orderBy)
	}*/
	err := datastore.GetMulti(appEngineContext, keys, reflect.ValueOf(results).Elem().Interface())
	if err != nil {
		if _, ok := err.(appengine.MultiError); !ok {
			log.Debug(ctx, "Geting object", "err", err)
			return nil, errors.WrapError(ctx, err)
		}
	}
	return results, nil
}

func (svc *gaeDataService) Count(ctx core.RequestContext, queryCond interface{}) (count int, err error) {
	ctx = ctx.SubContext("Count")
	appEngineContext := ctx.GetAppengineContext()
	query := datastore.NewQuery(svc.collection)
	query, err = svc.processCondition(ctx, appEngineContext, query, queryCond)
	if err != nil {
		return -1, errors.WrapError(ctx, err)
	}
	return query.Count(appEngineContext)
}

func (svc *gaeDataService) GetList(ctx core.RequestContext, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	ctx = ctx.SubContext("GetList")
	return svc.Get(ctx, nil, pageSize, pageNum, mode, orderBy) // resultStor, totalrecs, recsreturned, nil
}

func (svc *gaeDataService) Get(ctx core.RequestContext, queryCond interface{}, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	ctx = ctx.SubContext("Get")
	appEngineContext := ctx.GetAppengineContext()
	totalrecs = -1
	recsreturned = -1
	query := datastore.NewQuery(svc.collection)
	query, err = svc.processCondition(ctx, appEngineContext, query, queryCond)
	if err != nil {
		return nil, nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}
	if pageSize > 0 {
		totalrecs, err = query.Limit(500).Count(appEngineContext)
		if err != nil {
			return nil, nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
		}
		recsToSkip := (pageNum - 1) * pageSize
		query = query.Limit(pageSize).Offset(recsToSkip)
	}
	if len(orderBy) > 0 {
		query = query.Order(orderBy)
	}
	results, _ := ctx.CreateCollection(svc.Object, 0)

	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	_, err = query.GetAll(appEngineContext, results)
	if err != nil {
		return nil, nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}

	resultStor, ids, err := svc.postArrayGet(ctx, results)
	recsreturned = len(ids)
	if recsreturned > totalrecs {
		totalrecs = recsreturned
	}
	log.Trace(ctx, "Returning multiple objects ", "conditions", queryCond, "objectType", svc.Object, "recsreturned", recsreturned)
	return resultStor, ids, totalrecs, recsreturned, nil
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
	default:
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_NOT_IMPLEMENTED)
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
	/*case data.COMBINECONDTITIONS:
	{
		retval := query
		for _, subcond := range dqCondition.args {
			v, err := svc.processCondition(ctx, subcond, retval)
			if err != nil {
				return nil, err
			}
			retval = v
		}
	}*/
	case data.FIELDVALUE:
		queryCondMap, ok := dqCondition.arg1.(map[string]interface{})
		if svc.EmbeddedSearch {
			retMap := make(map[string]interface{})
			utils.FlattenMap(queryCondMap, retMap, "")
			queryCondMap = retMap
			log.Error(ctx, "creating condition embedded search", "args", queryCondMap)
		}
		if ok {
			queryCondMap = svc.PreProcessConditionMap(ctx, data.FIELDVALUE, queryCondMap)
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
