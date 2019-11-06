package main

import (
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
)

func (ms *mongoDataService) GetById(ctx core.RequestContext, id string) (data.Storable, error) {
	ctx = ctx.SubContext("GetById")
	log.Info(ctx, "Getting object by id ", "id", id, "object", ms.Object)

	object := ms.ObjectFactory.CreateObject(ctx)

	conn := ms.factory.getConnection(ctx)
	condition := bson.M{}
	if ms.Multitenant {
		condition["Tenant"] = ctx.GetUser().GetTenant()
	}
	if ms.SoftDelete {
		condition[ms.SoftDeleteField] = true
	}
	condition[ms.ObjectId] = id

	err := conn.Database(ms.database).Collection(ms.collection).FindOne(ctx, condition).Decode(object)
	if err != nil {
		//ErrNoDocuments
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.RethrowError(ctx, data.DATA_ERROR_OPERATION, err, "ID", id)
	}
	stor := object.(data.Storable)
	log.Trace(ctx, "Result", "id", id, "stor", stor)
	if ms.PostLoad {
		stor.PostLoad(ctx)
	}
	return stor, nil
}

//Get multiple objects by id
func (ms *mongoDataService) GetMulti(ctx core.RequestContext, ids []string, orderBy string) ([]data.Storable, error) {
	ctx = ctx.SubContext("GetMulti")
	results, _, err := ms.getMulti(ctx, ids, orderBy)
	if err != nil {
		return nil, err
	}
	if results == nil {
		return []data.Storable{}, nil
	}

	if ms.PostLoad {
		for _, stor := range results {
			stor.PostLoad(ctx)
		}
	}
	return results, err
}

func (ms *mongoDataService) GetMultiHash(ctx core.RequestContext, ids []string) (map[string]data.Storable, error) {
	ctx = ctx.SubContext("GetMultiHash")
	results, _, err := ms.getMulti(ctx, ids, "")
	if err != nil {
		return nil, err
	}
	if results == nil {
		return map[string]data.Storable{}, nil
	}

	if ms.PostLoad {
		for _, stor := range results {
			stor.PostLoad(ctx)
		}
	}
	resMap := make(map[string]data.Storable)
	for _, stor := range results {
		resMap[stor.GetId()] = stor
	}
	return resMap, nil
}

//Get multiple objects by id
func (ms *mongoDataService) getMulti(ctx core.RequestContext, ids []string, orderBy string) ([]data.Storable, []string, error) {

	log.Trace(ctx, "Getting multiple objects ", "Ids", ids)
	conn := ms.factory.getConnection(ctx)
	condition := bson.M{}
	if ms.Multitenant {
		condition["Tenant"] = ctx.GetUser().GetTenant()
	}
	operatorCond := bson.M{}
	operatorCond["$in"] = ids
	condition[ms.ObjectId] = operatorCond

	queryctx, _ := ctx.WithTimeout(10 * time.Second)

	/*	if len(orderBy) > 0 {
			query = query.Sort(orderBy)
		}
	*/

	cur, err := conn.Database(ms.database).Collection(ms.collection).Find(queryctx, condition, options.Find().SetSort(orderBy))

	if err != nil {
		return nil, nil, errors.WrapError(ctx, err)
	}
	defer cur.Close(queryctx)

	results, ids, count, err := ms.getResultsFromCursor(ctx, cur, false)

	if err != nil {
		return nil, nil, errors.WrapError(ctx, err)
	}

	log.Trace(ctx, "Getting multiple objects by Ids", "count", count, "collection", ms.collection)
	return results, ids, nil
}

func (ms *mongoDataService) Count(ctx core.RequestContext, queryCond interface{}) (count int, err error) {
	ctx = ctx.SubContext("Count")
	conn := ms.factory.getConnection(ctx)
	cur, err := conn.Database(ms.database).Collection(ms.collection).Find(ctx, queryCond)
	if err != nil {
		return -1, errors.WrapError(ctx, err)
	}
	defer cur.Close(ctx)
	_, _, count, err = ms.getResultsFromCursor(ctx, cur, true)

	return count, err
}

func (ms *mongoDataService) GetList(ctx core.RequestContext, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	ctx = ctx.SubContext("GetList")
	return ms.Get(ctx, bson.M{}, pageSize, pageNum, mode, orderBy) // resultStor, totalrecs, recsreturned, nil
}

func (ms *mongoDataService) Get(ctx core.RequestContext, queryCond interface{}, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	ctx = ctx.SubContext("Get")
	totalrecs = -1
	recsreturned = -1

	conn := ms.factory.getConnection(ctx)

	queryctx, _ := ctx.WithTimeout(10 * time.Second)
	findoptions := options.Find()
	findoptions.SetSort(orderBy)
	findoptions.SetBatchSize(int32(pageSize))
	recsToSkip := 0
	if pageSize > 0 {
		recsToSkip = (pageNum - 1) * pageSize
	}
	findoptions.SetSkip(int64(recsToSkip))

	cur, err := conn.Database(ms.database).Collection(ms.collection).Find(queryctx, queryCond, findoptions)

	if err != nil {
		return nil, nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}
	defer cur.Close(queryctx)

	results, ids, recsreturned, err := ms.getResultsFromCursor(ctx, cur, false)

	if err != nil {
		return nil, nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}

	if ms.PostLoad {
		for _, stor := range results {
			stor.PostLoad(ctx)
		}
	}
	log.Trace(ctx, "Returning multiple objects ", "conditions", queryCond, "objectType", ms.Object, "recsreturned", recsreturned)
	return results, ids, totalrecs, recsreturned, nil
}

func (ms *mongoDataService) getResultsFromCursor(ctx core.RequestContext, cursor *mongo.Cursor, countOnly bool) ([]data.Storable, []string, int, error) {
	count := 0
	var results []data.Storable
	var ids []string

	for cursor.Next(ctx) {
		count++
		if !countOnly {
			obj := ms.ObjectFactory.CreateObject(ctx)
			err := cursor.Decode(obj)
			if err != nil {
				return results, ids, count, errors.WrapError(ctx, err)
			}
			stor := obj.(data.Storable)
			results = append(results, stor)
			ids = append(ids, stor.GetId())
		}
	}
	return results, ids, count, nil
}

//create condition for passing to data service
func (ms *mongoDataService) CreateCondition(ctx core.RequestContext, operation data.ConditionType, args ...interface{}) (interface{}, error) {
	switch operation {
	case data.MATCHMULTIPLEVALUES:
		{
			if len(args) < 2 {
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG)
			}
			retVal := bson.M{args[0].(string): bson.M{"$in": args[1]}}
			if ms.Multitenant {
				retVal["Tenant"] = ctx.GetUser().GetTenant()
			}
			return retVal, nil
		}
	case data.FIELDVALUE:
		{
			if len(args) < 1 {
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG)
			}
			argsMap := args[0].(map[string]interface{})
			if ms.EmbeddedSearch {
				retMap := make(map[string]interface{})
				utils.FlattenMap(argsMap, retMap, "")
				argsMap = retMap
				log.Error(ctx, "creating condition embedded search", "args", argsMap)
			}
			log.Error(ctx, "creating condition search", "args", argsMap)
			argsMap = ms.PreProcessConditionMap(ctx, data.FIELDVALUE, argsMap)
			return argsMap, nil
		}
	default:
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_NOT_IMPLEMENTED)
	}
	return nil, nil
}
