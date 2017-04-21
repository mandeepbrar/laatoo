package sql

import (
	//	"github.com/jinzhu/gorm"
	"fmt"
	"laatoo/framework/services/data/common"
	"laatoo/sdk/components/data"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/utils"
	"strings"

	"github.com/jinzhu/gorm"
)

func (svc *sqlDataService) GetById(ctx core.RequestContext, id string) (data.Storable, error) {
	ctx = ctx.SubContext("GetById")
	log.Logger.Trace(ctx, "Getting object by id ", "id", id, "object", svc.Object)

	object := svc.ObjectCreator()

	err := svc.db.First(object, id).Error

	if err != nil {
		return nil, errors.RethrowError(ctx, common.DATA_ERROR_OPERATION, err, "ID", id)
	}
	stor := object.(data.Storable)
	if stor.IsDeleted() {
		return nil, nil
	}
	if svc.postload {
		stor.PostLoad(ctx)
	}
	return stor, nil
}

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

func (svc *sqlDataService) GetMultiHash(ctx core.RequestContext, ids []string) (map[string]data.Storable, error) {
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

func (svc *sqlDataService) postArrayGet(ctx core.RequestContext, results interface{}) ([]data.Storable, []string, error) {
	resultStor, ids, err := data.CastToStorableCollection(results)
	if err != nil {
		return nil, nil, errors.WrapError(ctx, err)
	}
	log.Logger.Trace(ctx, "Processing results in postArrayGet ", "number", len(resultStor))
	for _, stor := range resultStor {
		svc.postLoad(ctx, stor)
	}
	return resultStor, ids, nil
}

func (svc *sqlDataService) postLoad(ctx core.RequestContext, stor data.Storable) error {
	if svc.postload {
		stor.PostLoad(ctx)
	}
	return nil
}

//Get multiple objects by id
func (svc *sqlDataService) getMulti(ctx core.RequestContext, ids []string, orderBy string) (interface{}, error) {
	lenids := len(ids)
	if lenids == 0 {
		return nil, nil
	}
	results := svc.ObjectCollectionCreator(lenids)

	query := svc.db.Where(ids)

	if len(orderBy) > 0 {
		query = query.Order(orderBy)
	}

	err := query.Find(results).Error
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	log.Logger.Trace(ctx, "Got multiple objects by Ids", "len Ids", lenids, "collection", svc.collection)
	return results, nil
}

func (svc *sqlDataService) Count(ctx core.RequestContext, queryCond interface{}) (count int, err error) {
	ctx = ctx.SubContext("Count")
	query, err := svc.processCondition(ctx, queryCond, svc.db)
	if err != nil {
		return -1, err
	}
	err = query.Count(&count).Error
	if err != nil {
		return -1, errors.WrapError(ctx, err)
	}

	return count, nil
}

func (svc *sqlDataService) CountGroups(ctx core.RequestContext, queryCond interface{}, groupids []string, group string) (res map[string]interface{}, err error) {
	ctx = ctx.SubContext("CountGroups")
	groupIdsCond, err := svc.CreateCondition(ctx, data.MATCHMULTIPLEVALUES, group, groupids)
	if err != nil {
		return nil, err
	}
	countcond, err := svc.CreateCondition(ctx, data.COMBINECONDTITIONS, groupIdsCond, queryCond)
	if err != nil {
		return nil, err
	}
	query, err := svc.processCondition(ctx, countcond, svc.db.Select(group+" , count(*) "))
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	rows, err := query.Group(group).Rows()
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	res = make(map[string]interface{})
	for rows.Next() {
		var postId string
		var count int
		rows.Scan(&postId, &count)
		res[postId] = count
	}
	return
}

func (svc *sqlDataService) GetList(ctx core.RequestContext, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	ctx = ctx.SubContext("GetList")
	return svc.Get(ctx, map[string]interface{}{}, pageSize, pageNum, mode, orderBy) // resultStor, totalrecs, recsreturned, nil
}

func (svc *sqlDataService) Get(ctx core.RequestContext, queryCond interface{}, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	ctx = ctx.SubContext("Get")
	totalrecs = -1
	recsreturned = -1
	//0 is just a placeholder... mongo provides results of its own
	results := svc.ObjectCollectionCreator(0)

	query, err := svc.processCondition(ctx, queryCond, svc.db)
	if err != nil {
		return nil, nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
	}

	if pageSize > 0 {
		err := query.Count(&totalrecs).Error
		if err != nil {
			return nil, nil, totalrecs, recsreturned, errors.WrapError(ctx, err)
		}
		recsToSkip := (pageNum - 1) * pageSize
		query = query.Limit(pageSize).Offset(recsToSkip)
	}

	if len(orderBy) > 0 {
		query = query.Order(orderBy)
	}

	err = query.Find(results).Error

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
	log.Logger.Trace(ctx, "Returning multiple objects ", "conditions", queryCond, "objectType", svc.Object, "recsreturned", recsreturned)
	return resultStor, ids, totalrecs, recsreturned, nil
}

func (svc *sqlDataService) processCondition(ctx core.RequestContext, condition interface{}, input *gorm.DB) (*gorm.DB, error) {
	cond := condition.(*sqlCondition)

	switch cond.operation {
	case data.MATCHMULTIPLEVALUES:
		{
			if len(cond.args) < 2 {
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG)
			}
			arr := utils.CastToInterfaceArray(cond.args[1])
			lenarr := len(arr)
			if lenarr == 0 {
				return input, nil
			}
			placeholders := fmt.Sprintf("%s in( %s)", cond.args[0], strings.Trim(strings.Repeat("?,", lenarr), ","))
			return input.Where(placeholders, arr...), nil
			//append([]interface{}{placeholders}, arr...), nil //"name in (?)", []string{"jinzhu", "jinzhu 2"}
		}
	case data.COMBINECONDTITIONS:
		{
			retval := input
			for _, subcond := range cond.args {
				v, err := svc.processCondition(ctx, subcond, retval)
				if err != nil {
					return nil, err
				}
				retval = v
			}
			return retval, nil
		}
	case data.FIELDVALUE:
		{
			if len(cond.args) < 1 {
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_ARG)
			}
			argsMap := cond.args[0].(map[string]interface{})
			if svc.softdelete {
				argsMap[svc.softDeleteField] = false
			}
			return input.Where(argsMap), nil
		}
	default:
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_NOT_IMPLEMENTED)
	}
}

//create condition for passing to data service
func (svc *sqlDataService) CreateCondition(ctx core.RequestContext, operation data.ConditionType, args ...interface{}) (interface{}, error) {
	return &sqlCondition{operation, args}, nil
}

type sqlCondition struct {
	operation data.ConditionType
	args      []interface{}
}
