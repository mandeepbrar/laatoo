package main

import (
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"

	"gopkg.in/mgo.v2/bson"
)

func (ms *mongoDataService) GetById(ctx core.RequestContext, id string) (data.Storable, error) {
	ctx = ctx.SubContext("GetById")
	log.Info(ctx, "Getting object by id ", "id", id, "object", ms.Object)

	object := ms.ObjectCreator()

	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	condition := bson.M{}
	condition[ms.ObjectId] = id
	err := connCopy.DB(ms.database).C(ms.collection).Find(condition).One(object)
	if err != nil {
		if err.Error() == "not found" {
			return nil, nil
		}
		return nil, errors.RethrowError(ctx, data.DATA_ERROR_OPERATION, err, "ID", id)
	}
	stor := object.(data.Storable)
	if stor.IsDeleted() {
		return nil, nil
	}
	log.Trace(ctx, "Result", "id", id, "stor", stor)
	if ms.PostLoad {
		stor.PostLoad(ctx)
	}
	return stor, nil
}

//Get multiple objects by id
func (ms *mongoDataService) GetMulti(ctx core.RequestContext, ids []string, orderBy string) ([]data.Storable, error) {
	ctx = ctx.SubContext("GetMulti")
	results, err := ms.getMulti(ctx, ids, orderBy)
	if err != nil {
		return nil, err
	}
	if results == nil {
		return []data.Storable{}, nil
	}
	res, _, err := ms.postArrayGet(ctx, results)
	return res, err
}

func (ms *mongoDataService) GetMultiHash(ctx core.RequestContext, ids []string) (map[string]data.Storable, error) {
	ctx = ctx.SubContext("GetMultiHash")
	results, err := ms.getMulti(ctx, ids, "")
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
		ms.postLoad(ctx, stor)
	}
	return resultStor, nil
}

func (ms *mongoDataService) postArrayGet(ctx core.RequestContext, results interface{}) ([]data.Storable, []string, error) {
	resultStor, ids, err := data.CastToStorableCollection(results)
	if err != nil {
		return nil, nil, errors.WrapError(ctx, err)
	}
	for _, stor := range resultStor {
		ms.postLoad(ctx, stor)
	}
	log.Trace(ctx, "Returning results in postArrayGet ", "resultStor", resultStor)
	return resultStor, ids, nil
}

func (ms *mongoDataService) postLoad(ctx core.RequestContext, stor data.Storable) error {
	if ms.PostLoad {
		stor.PostLoad(ctx)
	}
	return nil
}

//Get multiple objects by id
func (ms *mongoDataService) getMulti(ctx core.RequestContext, ids []string, orderBy string) (interface{}, error) {
	lenids := len(ids)
	if lenids == 0 {
		return nil, nil
	}
	results := ms.ObjectCollectionCreator(lenids)

	log.Trace(ctx, "Getting multiple objects ", "Ids", ids)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	condition := bson.M{}
	operatorCond := bson.M{}
	operatorCond["$in"] = ids
	condition[ms.ObjectId] = operatorCond
	query := connCopy.DB(ms.database).C(ms.collection).Find(condition)
	if len(orderBy) > 0 {
		query = query.Sort(orderBy)
	}
	err := query.All(results)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	log.Trace(ctx, "Getting multiple objects by Ids", "len Ids", lenids, "collection", ms.collection)
	return results, nil
}

func (ms *mongoDataService) Count(ctx core.RequestContext, queryCond interface{}) (count int, err error) {
	ctx = ctx.SubContext("Count")
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	query := connCopy.DB(ms.database).C(ms.collection).Find(queryCond)
	return query.Count()
}

func (ms *mongoDataService) GetList(ctx core.RequestContext, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	ctx = ctx.SubContext("GetList")
	return ms.Get(ctx, bson.M{}, pageSize, pageNum, mode, orderBy) // resultStor, totalrecs, recsreturned, nil
}

func (ms *mongoDataService) Get(ctx core.RequestContext, queryCond interface{}, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	ctx = ctx.SubContext("Get")
	totalrecs = -1
	recsreturned = -1
	//0 is just a placeholder... mongo provides results of its own
	results := ms.ObjectCollectionCreator(0)

	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	query := connCopy.DB(ms.database).C(ms.collection).Find(queryCond)
	if pageSize > 0 {
		totalrecs, err = query.Count()
		if err != nil {
			return nil, nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
		}
		recsToSkip := (pageNum - 1) * pageSize
		query = query.Limit(pageSize).Skip(recsToSkip)
	}
	if len(orderBy) > 0 {
		query = query.Sort(orderBy)
	}
	err = query.All(results)
	if err != nil {
		return nil, nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}
	resultStor, ids, err := ms.postArrayGet(ctx, results)
	if err != nil {
		return nil, nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}
	recsreturned = len(ids)
	if recsreturned > totalrecs {
		totalrecs = recsreturned
	}
	log.Trace(ctx, "Returning multiple objects ", "conditions", queryCond, "objectType", ms.Object, "recsreturned", recsreturned)
	return resultStor, ids, totalrecs, recsreturned, nil
}

//create condition for passing to data service
func (ms *mongoDataService) CreateCondition(ctx core.RequestContext, operation data.ConditionType, args ...interface{}) (interface{}, error) {
	switch operation {
	case data.MATCHMULTIPLEVALUES:
		{
			if len(args) < 2 {
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG)
			}
			return bson.M{args[0].(string): bson.M{"$in": args[1]}}, nil
		}
	case data.FIELDVALUE:
		{
			if len(args) < 1 {
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG)
			}
			argsMap := args[0].(map[string]interface{})
			if ms.SoftDelete {
				argsMap[ms.SoftDeleteField] = false
			}
			return argsMap, nil
		}
	default:
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_NOT_IMPLEMENTED)
	}
	return nil, nil
}
