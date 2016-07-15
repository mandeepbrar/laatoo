package data

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"

	"gopkg.in/mgo.v2/bson"
)

func (ms *mongoDataService) GetById(ctx core.RequestContext, id string) (data.Storable, error) {
	log.Logger.Trace(ctx, "Getting object by id ", "id", id, "object", ms.object)

	//try cache if the object is cacheable
	if ms.cacheable {
		ent, err := ms.objectCreator(ctx, nil)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		ok := getFromCache(ctx, ms.object, id, ent)
		if ok {
			return ent.(data.Storable), nil
		}
	}

	object, err := ms.objectCreator(ctx, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	condition := bson.M{}
	condition[ms.objectid] = id
	err = connCopy.DB(ms.database).C(ms.collection).Find(condition).One(object)
	if err != nil {
		if err.Error() == "not found" {
			return nil, nil
		}
		return nil, errors.RethrowError(ctx, DATA_ERROR_OPERATION, err, "ID", id)
	}
	stor := object.(data.Storable)
	if ms.presave && stor.IsDeleted() {
		return nil, nil
	}
	if ms.postload {
		stor.PostLoad(ctx)
	}
	if ms.cacheable {
		putInCache(ctx, ms.object, stor.GetId(), stor)
	}
	return stor, nil
}

//Get multiple objects by id
func (ms *mongoDataService) GetMulti(ctx core.RequestContext, ids []string, orderBy string) (map[string]data.Storable, error) {
	lenids := len(ids)
	results, err := ms.objectCollectionCreator(ctx, lenids, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	log.Logger.Trace(ctx, "Getting multiple objects ", "Ids", ids)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	condition := bson.M{}
	operatorCond := bson.M{}
	operatorCond["$in"] = ids
	condition[ms.objectid] = operatorCond
	query := connCopy.DB(ms.database).C(ms.collection).Find(condition)
	if len(orderBy) > 0 {
		query = query.Sort(orderBy)
	}
	err = query.All(results)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	log.Logger.Trace(ctx, "Getting multiple objects ", "len Ids", lenids, "collection", ms.collection, "condition", condition)
	retVal := make(map[string]data.Storable, lenids)
	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	resultStor, err := data.CastToStorableCollection(results)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	for _, stor := range resultStor {
		if ms.presave && stor.IsDeleted() {
			continue
		}
		id := stor.GetId()
		retVal[id] = stor
		if ms.postload {
			stor.PostLoad(ctx)
		}
		if ms.cacheable {
			putInCache(ctx, ms.object, stor.GetId(), stor)
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

func (ms *mongoDataService) GetList(ctx core.RequestContext, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, totalrecs int, recsreturned int, err error) {
	return ms.Get(ctx, bson.M{}, pageSize, pageNum, mode, orderBy) // resultStor, totalrecs, recsreturned, nil
}

func (ms *mongoDataService) Get(ctx core.RequestContext, queryCond interface{}, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, totalrecs int, recsreturned int, err error) {
	totalrecs = -1
	recsreturned = -1
	//0 is just a placeholder... mongo provides results of its own
	results, err := ms.objectCollectionCreator(ctx, 0, nil)
	if err != nil {
		return nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	query := connCopy.DB(ms.database).C(ms.collection).Find(queryCond)
	if pageSize > 0 {
		totalrecs, err = query.Count()
		if err != nil {
			return nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
		}
		recsToSkip := (pageNum - 1) * pageSize
		query = query.Limit(pageSize).Skip(recsToSkip)
	}
	if len(orderBy) > 0 {
		query = query.Sort(orderBy)
	}
	err = query.All(results)
	log.Logger.Trace(ctx, "Returning multiple objects ", "conditions", queryCond, "objectType", ms.object, "collection", ms.collection)
	resultStor, err := data.CastToStorableCollection(results)
	if err != nil {
		return nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}
	recsreturned = len(resultStor)
	for _, stor := range resultStor {
		log.Logger.Trace(ctx, "Returning multiple objects ", "stor", stor)
		if ms.postload {
			stor.PostLoad(ctx)
		}
		if ms.cacheable {
			putInCache(ctx, ms.object, stor.GetId(), stor)
		}
	}
	if recsreturned > totalrecs {
		totalrecs = recsreturned
	}
	log.Logger.Trace(ctx, "Returning multiple objects ", "conditions", queryCond, "objectType", ms.object, "recsreturned", recsreturned)
	return resultStor, totalrecs, recsreturned, nil
}

//create condition for passing to data service
func (ms *mongoDataService) CreateCondition(ctx core.RequestContext, operation data.ConditionType, args ...interface{}) (interface{}, error) {
	switch operation {
	case data.MATCHANCESTOR:
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_NOT_IMPLEMENTED)
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
			if ms.softdelete {
				argsMap[ms.softDeleteField] = false
			}
			return argsMap, nil
		}
	}
	return nil, nil
}
