package sql

import (
	//	"github.com/jinzhu/gorm"
	"laatoo/framework/services/data/common"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	//	"laatoo/sdk/utils"
)

type sqlDataService struct {
	*data.BaseComponent
	conf                    config.Config
	objectConfig            *data.StorableConfig
	name                    string
	database                string
	auditable               bool
	softdelete              bool
	cacheable               bool
	refops                  bool
	presave                 bool
	postsave                bool
	postload                bool
	notifynew               bool
	notifyupdates           bool
	factory                 *sqlDataServicesFactory
	collection              string
	softDeleteField         string
	object                  string
	objectid                string
	objectCollectionCreator core.ObjectCollectionCreator
	objectCreator           core.ObjectCreator
	deleteRefOpers          []common.RefOperation
	getRefOpers             []common.RefOperation
	serviceType             string
}

func newSqlDataService(ctx core.ServerContext, name string, svc *sqlDataServicesFactory) (*sqlDataService, error) {
	sqlDataSvc := &sqlDataService{name: name, factory: svc, serviceType: "SQL"}
	return sqlDataSvc, nil
}

func (svc *sqlDataService) Initialize(ctx core.ServerContext, conf config.Config) error {
	ctx = ctx.SubContext("Initialize SQL Data Service")
	object, ok := conf.GetString(common.CONF_DATA_OBJECT)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", common.CONF_DATA_OBJECT)
	}
	objectCreator, err := ctx.GetObjectCreator(object)
	if err != nil {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "Could not get Object creator for", object)
	}
	objectCollectionCreator, err := ctx.GetObjectCollectionCreator(object)
	if err != nil {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "Could not get Object Collection creator for", object)
	}

	svc.conf = conf
	svc.object = object
	svc.objectCreator = objectCreator
	svc.objectCollectionCreator = objectCollectionCreator

	testObj, _ := objectCreator(ctx, nil)
	stor := testObj.(data.Storable)
	svc.objectConfig = stor.Config()

	svc.objectid = svc.objectConfig.IdField
	svc.softDeleteField = svc.objectConfig.SoftDeleteField

	if svc.softDeleteField == "" {
		svc.softdelete = false
	} else {
		svc.softdelete = true
	}

	collection, ok := conf.GetString(common.CONF_DATA_COLLECTION)
	if ok {
		svc.collection = collection
	} else {
		svc.collection = svc.objectConfig.Collection
	}

	if svc.collection == "" {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", common.CONF_DATA_COLLECTION)
	}

	cacheable, ok := conf.GetBool(common.CONF_DATA_CACHEABLE)
	if ok {
		svc.cacheable = cacheable
	} else {
		svc.cacheable = svc.objectConfig.Cacheable
	}

	auditable, ok := conf.GetBool(common.CONF_DATA_AUDITABLE)
	if ok {
		svc.auditable = auditable
	} else {
		svc.auditable = svc.objectConfig.Auditable
	}
	postsave, ok := conf.GetBool(common.CONF_DATA_POSTSAVE)
	if ok {
		svc.postsave = postsave
	} else {
		svc.postsave = svc.objectConfig.PostSave
	}
	presave, ok := conf.GetBool(common.CONF_DATA_PRESAVE)
	if ok {
		svc.presave = presave
	} else {
		svc.presave = svc.objectConfig.PreSave
	}
	postload, ok := conf.GetBool(common.CONF_DATA_POSTLOAD)
	if ok {
		svc.postload = postload
	} else {
		svc.postload = svc.objectConfig.PostLoad
	}

	refops, ok := conf.GetBool(common.CONF_DATA_REFOPS)
	if ok {
		svc.refops = refops
	} else {
		svc.refops = svc.objectConfig.RefOps
	}

	notifyupdates, ok := conf.GetBool(common.CONF_DATA_NOTIFYUPDATES)
	if ok {
		svc.notifyupdates = notifyupdates
	} else {
		svc.notifyupdates = svc.objectConfig.NotifyUpdates
	}
	notifynew, ok := conf.GetBool(common.CONF_DATA_NOTIFYNEW)
	if ok {
		svc.notifynew = notifynew
	} else {
		svc.notifynew = svc.objectConfig.NotifyNew
	}

	if svc.refops {
		deleteOps, _, _, getRefOps, err := common.BuildRefOps(ctx, conf)
		if err != nil {
			return err
		}
		err = common.InitialRefOps(ctx, deleteOps)
		if err != nil {
			return err
		}

		err = common.InitialRefOps(ctx, getRefOps)
		if err != nil {
			return err
		}
		svc.deleteRefOpers = deleteOps
		svc.getRefOpers = getRefOps
	}

	return nil
}

func (svc *sqlDataService) Start(ctx core.ServerContext) error {
	return nil
}

func (svc *sqlDataService) Invoke(ctx core.RequestContext) error {
	return nil
}

func (svc *sqlDataService) GetName() string {
	return svc.name
}

func (svc *sqlDataService) GetDataServiceType() string {
	return data.DATASERVICE_TYPE_NOSQL
}

func (svc *sqlDataService) GetObject() string {
	return svc.object
}

func (svc *sqlDataService) Supports(feature data.Feature) bool {
	switch feature {
	case data.InQueries:
		return true
	case data.Ancestors:
		return false
	}
	return false
}

func (svc *sqlDataService) Save(ctx core.RequestContext, item data.Storable) error {
	ctx = ctx.SubContext("Save")
	log.Logger.Trace(ctx, "Saving object", "Object", svc.object)
	if svc.presave {
		err := ctx.SendSynchronousMessage(common.CONF_PRESAVE_MSG, item)
		if err != nil {
			return err
		}
		err = item.PreSave(ctx)
		if err != nil {
			return err
		}
	}
	if svc.auditable {
		data.Audit(ctx, item)
	}
	id := item.GetId()
	if id == "" {
		return errors.ThrowError(ctx, common.DATA_ERROR_ID_NOT_FOUND, "ObjectType", svc.object)
	}
	err := svc.factory.connection.Create(item).Error
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if svc.postsave {
		err = item.PostSave(ctx)
		if err != nil {
			errors.WrapError(ctx, err)
		}
		err := ctx.SendSynchronousMessage(common.CONF_NEWOBJ_MSG, item)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
func (svc *sqlDataService) PutMulti(ctx core.RequestContext, items []data.Storable) error {
	ctx = ctx.SubContext("PutMulti")
	log.Logger.Trace(ctx, "Saving multiple objects", "ObjectType", svc.object)
	connCopy := svc.factory.connection.Copy()
	defer connCopy.Close()
	bulk := connCopy.DB(svc.database).C(svc.collection).Bulk()
	for _, item := range items {
		id := item.GetId()
		common.InvalidateCache(ctx, svc.object, id)
		if svc.presave {
			err := ctx.SendSynchronousMessage(common.CONF_PRESAVE_MSG, item)
			if err != nil {
				return err
			}
			err = item.PreSave(ctx)
			if err != nil {
				return err
			}
		}
		if svc.auditable {
			data.Audit(ctx, item)
		}
		bulk.Upsert(bson.M{svc.objectid: id}, item)
	}
	_, err := bulk.Run()
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if svc.postsave || svc.notifyupdates {
		for _, item := range items {
			if svc.postsave {
				err = item.PostSave(ctx)
				if err != nil {
					errors.WrapError(ctx, err)
				}
			}
			if svc.notifyupdates {
				common.NotifyUpdate(ctx, svc.object, item.GetId())
			}
		}
	}
	log.Logger.Trace(ctx, "Saved multiple objects")
	return nil
}

func (svc *sqlDataService) Put(ctx core.RequestContext, id string, item data.Storable) error {
	ctx = ctx.SubContext("Put")
	common.InvalidateCache(ctx, svc.object, id)
	log.Logger.Trace(ctx, "Putting object", "ObjectType", svc.object, "id", id)
	connCopy := svc.factory.connection.Copy()
	defer connCopy.Close()
	condition := bson.M{}
	condition[svc.objectid] = id
	item.SetId(id)
	if svc.presave {
		err := ctx.SendSynchronousMessage(common.CONF_PRESAVE_MSG, item)
		if err != nil {
			return err
		}
		err = item.PreSave(ctx)
		log.Logger.Trace(ctx, "Putting object", "err", err)
		if err != nil {
			return err
		}
	}
	if svc.auditable {
		data.Audit(ctx, item)
	}
	err := connCopy.DB(svc.database).C(svc.collection).Update(condition, item)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if svc.postsave {
		err = item.PostSave(ctx)
		if err != nil {
			errors.WrapError(ctx, err)
		}
	}
	if svc.notifyupdates {
		common.NotifyUpdate(ctx, svc.object, id)
	}
	return nil
}

//upsert an object ...insert if not there... update if there
func (svc *sqlDataService) UpsertId(ctx core.RequestContext, id string, newVals map[string]interface{}) error {
	ctx = ctx.SubContext("UpsertId")
	return svc.update(ctx, id, newVals, true)
}

func (svc *sqlDataService) Update(ctx core.RequestContext, id string, newVals map[string]interface{}) error {
	ctx = ctx.SubContext("Update")
	return svc.update(ctx, id, newVals, false)
}

func (svc *sqlDataService) update(ctx core.RequestContext, id string, newVals map[string]interface{}, upsert bool) error {
	if svc.auditable {
		data.Audit(ctx, newVals)
	}
	if svc.presave {
		err := ctx.SendSynchronousMessage(common.CONF_PREUPDATE_MSG, map[string]interface{}{"id": id, "type": svc.object, "data": newVals})
		if err != nil {
			return err
		}
	}
	if upsert {
		newVals[svc.objectid] = id
	}
	condition := bson.M{}
	condition[svc.objectid] = id
	connCopy := svc.factory.connection.Copy()
	defer connCopy.Close()
	var err error
	if upsert {
		_, err = connCopy.DB(svc.database).C(svc.collection).UpsertId(condition, newVals)

	} else {
		updateInterface := map[string]interface{}{"$set": newVals}
		err = connCopy.DB(svc.database).C(svc.collection).Update(condition, updateInterface)
	}
	if err != nil {
		return err
	}
	common.InvalidateCache(ctx, svc.object, id)
	if svc.notifyupdates {
		common.NotifyUpdate(ctx, svc.object, id)
	}
	return nil
}

func (svc *sqlDataService) Upsert(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	ctx = ctx.SubContext("Upsert")
	return svc.updateAll(ctx, queryCond, newVals, true)
}

func (svc *sqlDataService) UpdateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	ctx = ctx.SubContext("UpdateAll")
	return svc.updateAll(ctx, queryCond, newVals, false)
}

func (svc *sqlDataService) updateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}, upsert bool) ([]string, error) {
	results, err := svc.objectCollectionCreator(ctx, 0, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	if svc.presave {
		err := ctx.SendSynchronousMessage(common.CONF_PREUPDATE_MSG, map[string]interface{}{"cond": queryCond, "type": svc.object, "data": newVals})
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
	}
	connCopy := svc.factory.connection.Copy()
	defer connCopy.Close()
	query := connCopy.DB(svc.database).C(svc.collection).Find(queryCond)
	err = query.All(results)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	resultStor, ids, err := data.CastToStorableCollection(results)
	length := len(resultStor)
	if length == 0 {
		if upsert {
			object, err := svc.objectCreator(ctx, nil)
			if err != nil {
				return nil, err
			}
			utils.SetObjectFields(object, newVals)
			stor := object.(data.Storable)
			err = svc.Save(ctx, stor)
			if err != nil {
				return nil, err
			}
			return []string{stor.GetId()}, nil
		}
		return []string{}, nil
	}
	if svc.auditable {
		data.Audit(ctx, newVals)
	}
	_, err = connCopy.DB(svc.database).C(svc.collection).UpdateAll(queryCond, map[string]interface{}{"$set": newVals})
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	for _, id := range ids {
		common.InvalidateCache(ctx, svc.object, id)
		if svc.notifyupdates {
			common.NotifyUpdate(ctx, svc.object, id)
		}
	}
	return ids, err
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (svc *sqlDataService) UpdateMulti(ctx core.RequestContext, ids []string, newVals map[string]interface{}) error {
	ctx = ctx.SubContext("UpdateMulti")
	if svc.auditable {
		data.Audit(ctx, newVals)
	}
	if svc.presave {
		err := ctx.SendSynchronousMessage(common.CONF_PREUPDATE_MSG, map[string]interface{}{"ids": ids, "type": svc.object, "data": newVals})
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	updateInterface := map[string]interface{}{"$set": newVals}
	condition, _ := svc.CreateCondition(ctx, data.MATCHMULTIPLEVALUES, svc.objectid, ids)
	connCopy := svc.factory.connection.Copy()
	defer connCopy.Close()
	_, err := connCopy.DB(svc.database).C(svc.collection).UpdateAll(condition, updateInterface)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	for _, id := range ids {
		common.InvalidateCache(ctx, svc.object, id)
		if svc.notifyupdates {
			common.NotifyUpdate(ctx, svc.object, id)
		}
	}
	return nil
}

//item must support Deleted field for soft deletes
func (svc *sqlDataService) Delete(ctx core.RequestContext, id string) error {
	ctx = ctx.SubContext("Delete")
	if svc.softdelete {
		err := svc.Update(ctx, id, map[string]interface{}{svc.softDeleteField: true})
		if err == nil {
			err = common.DeleteRefOps(ctx, svc.deleteRefOpers, []string{id})
		}
		return err
	}
	condition := bson.M{}
	condition[svc.objectid] = id
	connCopy := svc.factory.connection.Copy()
	defer connCopy.Close()
	err := connCopy.DB(svc.database).C(svc.collection).Remove(condition)
	if err == nil {
		err = common.DeleteRefOps(ctx, svc.deleteRefOpers, []string{id})
	}
	common.InvalidateCache(ctx, svc.object, id)
	return err
}

//Delete object by ids
func (svc *sqlDataService) DeleteMulti(ctx core.RequestContext, ids []string) error {
	ctx = ctx.SubContext("DeleteMulti")
	if svc.softdelete {
		err := svc.UpdateMulti(ctx, ids, map[string]interface{}{svc.softDeleteField: true})
		if err == nil {
			err = common.DeleteRefOps(ctx, svc.deleteRefOpers, ids)
		}
		return err
	}
	conditionVal := bson.M{}
	conditionVal["$in"] = ids
	condition := bson.M{}
	condition[svc.objectid] = conditionVal
	connCopy := svc.factory.connection.Copy()
	defer connCopy.Close()
	_, err := connCopy.DB(svc.database).C(svc.collection).RemoveAll(condition)
	if err == nil {
		err = common.DeleteRefOps(ctx, svc.deleteRefOpers, ids)
	}
	for _, id := range ids {
		common.InvalidateCache(ctx, svc.object, id)
	}
	return err
}

//Delete object by condition
func (svc *sqlDataService) DeleteAll(ctx core.RequestContext, queryCond interface{}) ([]string, error) {
	ctx = ctx.SubContext("DeleteAll")
	if svc.softdelete {
		ids, err := svc.UpdateAll(ctx, queryCond, map[string]interface{}{svc.softDeleteField: true})
		if err == nil {
			err = common.DeleteRefOps(ctx, svc.deleteRefOpers, ids)
		}
		return ids, err
	}
	results, err := svc.objectCollectionCreator(ctx, 0, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	connCopy := svc.factory.connection.Copy()
	defer connCopy.Close()
	query := connCopy.DB(svc.database).C(svc.collection).Find(queryCond)
	err = query.All(results)
	resultStor, ids, err := data.CastToStorableCollection(results)
	length := len(resultStor)
	if length == 0 {
		return nil, nil
	}
	_, err = connCopy.DB(svc.database).C(svc.collection).RemoveAll(queryCond)
	if err == nil {
		err = common.DeleteRefOps(ctx, svc.deleteRefOpers, ids)
		if err != nil {
			return ids, err
		}
		for _, id := range ids {
			common.InvalidateCache(ctx, svc.object, id)
		}
	} else {
		return nil, err
	}
	return ids, err
}*/
