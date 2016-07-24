package sql

import (
	//	"github.com/jinzhu/gorm"
	"laatoo/framework/services/data/common"
	"laatoo/sdk/components/data"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

func (svc *sqlDataService) GetById(ctx core.RequestContext, id string) (data.Storable, error) {
	ctx = ctx.SubContext("GetById")
	log.Logger.Trace(ctx, "Getting object by id ", "id", id, "object", svc.object)

	//try cache if the object is cacheable
	if svc.cacheable {
		ent, err := svc.objectCreator(ctx, nil)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		ok := common.GetFromCache(ctx, svc.object, id, ent)
		if ok {
			return ent.(data.Storable), nil
		}
	}

	object, err := svc.objectCreator(ctx, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	err = svc.factory.connection.First(object, id).Error

	if err != nil {
		return nil, errors.RethrowError(ctx, common.DATA_ERROR_OPERATION, err, "ID", id)
	}
	stor := object.(data.Storable)
	if stor.IsDeleted() {
		return nil, nil
	}
	if svc.refops {
		res, err := common.GetRefOps(ctx, svc.getRefOpers, []string{id}, []data.Storable{stor})
		if err != nil {
			return nil, err
		}
		r := res.([]data.Storable)
		stor = r[0]
	}
	if svc.postload {
		stor.PostLoad(ctx)
	}
	if svc.cacheable {
		common.PutInCache(ctx, svc.object, stor.GetId(), stor)
	}
	return stor, nil
}

/*
//Get multiple objects by id
func (svc *sqlDataService) GetMulti(ctx core.RequestContext, ids []string, orderBy string) ([]data.Storable, error) {
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

func (svc *sqlDataService) GetMultiHash(ctx core.RequestContext, ids []string, orderBy string) (map[string]data.Storable, error) {
	ctx = ctx.SubContext("GetMultiHash")
	results, err := svc.getMulti(ctx, ids, orderBy)
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

	if svc.refops {
		res, err := common.GetRefOps(ctx, svc.getRefOpers, ids, resultStor)
		if err != nil {
			return nil, err
		}
		resultStor = res.(map[string]data.Storable)
	}

	for _, stor := range resultStor {
		svc.postLoad(ctx, stor)
	}
	return resultStor, nil
}

func (svc *sqlDataService) postArrayGet(ctx core.RequestContext, results interface{}) ([]data.Storable, []string, error) {
	resultStor, ids, err := data.CastToStorableCollection(results)
	if err != nil {
		return nil, nil, errors.WrapError(ctx, err)
	}
	log.Logger.Trace(ctx, "Processing results in postArrayGet ", "number", len(resultStor))
	if svc.refops {
		res, err := common.GetRefOps(ctx, svc.getRefOpers, ids, resultStor)
		if err != nil {
			return nil, nil, err
		}
		resultStor = res.([]data.Storable)
		log.Logger.Trace(ctx, "Assigned results", "resultstor", resultStor)
	}
	for _, stor := range resultStor {
		svc.postLoad(ctx, stor)
	}
	log.Logger.Trace(ctx, "Returning results in postArrayGet ", "resultStor", resultStor)
	return resultStor, ids, nil
}

func (svc *sqlDataService) postLoad(ctx core.RequestContext, stor data.Storable) error {
	if svc.postload {
		stor.PostLoad(ctx)
	}
	if svc.cacheable {
		common.PutInCache(ctx, svc.object, stor.GetId(), stor)
	}
	return nil
}

//Get multiple objects by id
func (svc *sqlDataService) getMulti(ctx core.RequestContext, ids []string, orderBy string) (interface{}, error) {
	lenids := len(ids)
	if lenids == 0 {
		return nil, nil
	}
	results, err := svc.objectCollectionCreator(ctx, lenids, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	log.Logger.Trace(ctx, "Getting multiple objects ", "Ids", ids)
	connCopy := svc.factory.connection.Copy()
	defer connCopy.Close()
	condition := bson.M{}
	operatorCond := bson.M{}
	operatorCond["$in"] = ids
	condition[svc.objectid] = operatorCond
	query := connCopy.DB(svc.database).C(svc.collection).Find(condition)
	if len(orderBy) > 0 {
		query = query.Sort(orderBy)
	}
	err = query.All(results)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	log.Logger.Trace(ctx, "Getting multiple objects by Ids", "len Ids", lenids, "collection", svc.collection)
	return results, nil
}

func (svc *sqlDataService) Count(ctx core.RequestContext, queryCond interface{}) (count int, err error) {
	ctx = ctx.SubContext("Count")
	connCopy := svc.factory.connection.Copy()
	defer connCopy.Close()
	query := connCopy.DB(svc.database).C(svc.collection).Find(queryCond)
	return query.Count()
}

func (svc *sqlDataService) GetList(ctx core.RequestContext, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	ctx = ctx.SubContext("GetList")
	return svc.Get(ctx, bson.M{}, pageSize, pageNum, mode, orderBy) // resultStor, totalrecs, recsreturned, nil
}

func (svc *sqlDataService) Get(ctx core.RequestContext, queryCond interface{}, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	ctx = ctx.SubContext("Get")
	totalrecs = -1
	recsreturned = -1
	//0 is just a placeholder... mongo provides results of its own
	results, err := svc.objectCollectionCreator(ctx, 0, nil)
	if err != nil {
		return nil, nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}
	connCopy := svc.factory.connection.Copy()
	defer connCopy.Close()
	query := connCopy.DB(svc.database).C(svc.collection).Find(queryCond)
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
	resultStor, ids, err := svc.postArrayGet(ctx, results)
	if err != nil {
		return nil, nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}
	recsreturned = len(ids)
	if recsreturned > totalrecs {
		totalrecs = recsreturned
	}
	log.Logger.Trace(ctx, "Returning multiple objects ", "conditions", queryCond, "objectType", svc.object, "recsreturned", recsreturned)
	return resultStor, ids, totalrecs, recsreturned, nil
}

//create condition for passing to data service
func (svc *sqlDataService) CreateCondition(ctx core.RequestContext, operation data.ConditionType, args ...interface{}) (interface{}, error) {
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
			if svc.softdelete {
				argsMap[svc.softDeleteField] = false
			}
			return argsMap, nil
		}
	}
	return nil, nil
}
*/
