package mongo

import (
	"gopkg.in/mgo.v2/bson"
	"laatoo/framework/services/data/common"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/utils"
)

type mongoDataService struct {
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
	factory                 *mongoDataServicesFactory
	collection              string
	softDeleteField         string
	object                  string
	objectid                string
	objectCollectionCreator core.ObjectCollectionCreator
	objectCreator           core.ObjectCreator
	deleteRefOpers          []common.RefOperation
	getRefOpers             []common.RefOperation
	serviceType             string
	/*getRefOpers    map[string][]*refKeyOperation
	putRefOpers    map[string][]*refKeyOperation
	updateRefOpers map[string][]*refKeyOperation*/
}

const (
	CONF_MONGO_DATABASE = "database"
)

func newMongoDataService(ctx core.ServerContext, name string, ms *mongoDataServicesFactory) (*mongoDataService, error) {
	mongoSvc := &mongoDataService{name: name, factory: ms, serviceType: "Mongo"}
	return mongoSvc, nil
}

func (ms *mongoDataService) Initialize(ctx core.ServerContext, conf config.Config) error {
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
	ms.database = ms.factory.database

	ms.conf = conf
	ms.object = object
	ms.objectCreator = objectCreator
	ms.objectCollectionCreator = objectCollectionCreator

	testObj, _ := objectCreator(ctx, nil)
	stor := testObj.(data.Storable)
	ms.objectConfig = stor.Config()

	ms.objectid = ms.objectConfig.IdField
	ms.softDeleteField = ms.objectConfig.SoftDeleteField

	if ms.softDeleteField == "" {
		ms.softdelete = false
	} else {
		ms.softdelete = true
	}

	collection, ok := conf.GetString(common.CONF_DATA_COLLECTION)
	if ok {
		ms.collection = collection
	} else {
		ms.collection = ms.objectConfig.Collection
	}

	if ms.collection == "" {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", common.CONF_DATA_COLLECTION)
	}

	cacheable, ok := conf.GetBool(common.CONF_DATA_CACHEABLE)
	if ok {
		ms.cacheable = cacheable
	} else {
		ms.cacheable = ms.objectConfig.Cacheable
	}

	auditable, ok := conf.GetBool(common.CONF_DATA_AUDITABLE)
	if ok {
		ms.auditable = auditable
	} else {
		ms.auditable = ms.objectConfig.Auditable
	}
	postsave, ok := conf.GetBool(common.CONF_DATA_POSTSAVE)
	if ok {
		ms.postsave = postsave
	} else {
		ms.postsave = ms.objectConfig.PostSave
	}
	presave, ok := conf.GetBool(common.CONF_DATA_PRESAVE)
	if ok {
		ms.presave = presave
	} else {
		ms.presave = ms.objectConfig.PreSave
	}
	postload, ok := conf.GetBool(common.CONF_DATA_POSTLOAD)
	if ok {
		ms.postload = postload
	} else {
		ms.postload = ms.objectConfig.PostLoad
	}

	refops, ok := conf.GetBool(common.CONF_DATA_REFOPS)
	if ok {
		ms.refops = refops
	} else {
		ms.refops = ms.objectConfig.RefOps
	}

	notifyupdates, ok := conf.GetBool(common.CONF_DATA_NOTIFYUPDATES)
	if ok {
		ms.notifyupdates = notifyupdates
	} else {
		ms.notifyupdates = ms.objectConfig.NotifyUpdates
	}
	notifynew, ok := conf.GetBool(common.CONF_DATA_NOTIFYNEW)
	if ok {
		ms.notifynew = notifynew
	} else {
		ms.notifynew = ms.objectConfig.NotifyNew
	}

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
	ms.deleteRefOpers = deleteOps
	ms.getRefOpers = getRefOps

	return nil
}

func (ms *mongoDataService) Start(ctx core.ServerContext) error {
	return nil
}

func (ms *mongoDataService) Invoke(ctx core.RequestContext) error {
	return nil
}

func (fs *mongoDataService) GetName() string {
	return fs.name
}

func (ms *mongoDataService) GetDataServiceType() string {
	return data.DATASERVICE_TYPE_NOSQL
}

func (ms *mongoDataService) GetObject() string {
	return ms.object
}

func (ms *mongoDataService) Supports(feature data.Feature) bool {
	switch feature {
	case data.InQueries:
		return true
	case data.Ancestors:
		return false
	}
	return false
}

func (ms *mongoDataService) Save(ctx core.RequestContext, item data.Storable) error {
	log.Logger.Trace(ctx, "Saving object", "Object", ms.object)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	if ms.presave {
		err := ctx.SendSynchronousMessage(common.CONF_PRESAVE_MSG, item)
		if err != nil {
			return err
		}
		err = item.PreSave(ctx)
		if err != nil {
			return err
		}
	}
	if ms.auditable {
		data.Audit(ctx, item)
	}
	id := item.GetId()
	if id == "" {
		return errors.ThrowError(ctx, common.DATA_ERROR_ID_NOT_FOUND, "ObjectType", ms.object)
	}
	err := connCopy.DB(ms.database).C(ms.collection).Insert(item)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if ms.postsave {
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

func (ms *mongoDataService) PutMulti(ctx core.RequestContext, items []data.Storable) error {
	log.Logger.Trace(ctx, "Saving multiple objects", "ObjectType", ms.object)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	bulk := connCopy.DB(ms.database).C(ms.collection).Bulk()
	for _, item := range items {
		id := item.GetId()
		common.InvalidateCache(ctx, ms.object, id)
		if ms.presave {
			err := ctx.SendSynchronousMessage(common.CONF_PRESAVE_MSG, item)
			if err != nil {
				return err
			}
			err = item.PreSave(ctx)
			if err != nil {
				return err
			}
		}
		if ms.auditable {
			data.Audit(ctx, item)
		}
		bulk.Upsert(bson.M{ms.objectid: id}, item)
	}
	_, err := bulk.Run()
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if ms.postsave || ms.notifyupdates {
		for _, item := range items {
			if ms.postsave {
				err = item.PostSave(ctx)
				if err != nil {
					errors.WrapError(ctx, err)
				}
			}
			if ms.notifyupdates {
				common.NotifyUpdate(ctx, ms.object, item.GetId())
			}
		}
	}
	log.Logger.Trace(ctx, "Saved multiple objects")
	return nil
}

func (ms *mongoDataService) Put(ctx core.RequestContext, id string, item data.Storable) error {
	common.InvalidateCache(ctx, ms.object, id)
	log.Logger.Trace(ctx, "Putting object", "ObjectType", ms.object, "id", id)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	condition := bson.M{}
	condition[ms.objectid] = id
	item.SetId(id)
	if ms.presave {
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
	if ms.auditable {
		data.Audit(ctx, item)
	}
	err := connCopy.DB(ms.database).C(ms.collection).Update(condition, item)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if ms.postsave {
		err = item.PostSave(ctx)
		if err != nil {
			errors.WrapError(ctx, err)
		}
	}
	if ms.notifyupdates {
		common.NotifyUpdate(ctx, ms.object, id)
	}
	return nil
}

//upsert an object ...insert if not there... update if there
func (ms *mongoDataService) UpsertId(ctx core.RequestContext, id string, newVals map[string]interface{}) error {
	return ms.update(ctx, id, newVals, true)
}

func (ms *mongoDataService) Update(ctx core.RequestContext, id string, newVals map[string]interface{}) error {
	return ms.update(ctx, id, newVals, false)
}

func (ms *mongoDataService) update(ctx core.RequestContext, id string, newVals map[string]interface{}, upsert bool) error {
	if ms.auditable {
		data.Audit(ctx, newVals)
	}
	if ms.presave {
		err := ctx.SendSynchronousMessage(common.CONF_PREUPDATE_MSG, map[string]interface{}{"id": id, "type": ms.object, "data": newVals})
		if err != nil {
			return err
		}
	}
	if upsert {
		newVals[ms.objectid] = id
	}
	condition := bson.M{}
	condition[ms.objectid] = id
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	var err error
	if upsert {
		_, err = connCopy.DB(ms.database).C(ms.collection).UpsertId(condition, newVals)

	} else {
		updateInterface := map[string]interface{}{"$set": newVals}
		err = connCopy.DB(ms.database).C(ms.collection).Update(condition, updateInterface)
	}
	if err != nil {
		return err
	}
	common.InvalidateCache(ctx, ms.object, id)
	if ms.notifyupdates {
		common.NotifyUpdate(ctx, ms.object, id)
	}
	return nil
}

func (ms *mongoDataService) Upsert(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	return ms.updateAll(ctx, queryCond, newVals, true)
}

func (ms *mongoDataService) UpdateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	return ms.updateAll(ctx, queryCond, newVals, false)
}

func (ms *mongoDataService) updateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}, upsert bool) ([]string, error) {
	results, err := ms.objectCollectionCreator(ctx, 0, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	if ms.presave {
		err := ctx.SendSynchronousMessage(common.CONF_PREUPDATE_MSG, map[string]interface{}{"cond": queryCond, "type": ms.object, "data": newVals})
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
	}
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	query := connCopy.DB(ms.database).C(ms.collection).Find(queryCond)
	err = query.All(results)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	resultStor, ids, err := data.CastToStorableCollection(results)
	length := len(resultStor)
	if length == 0 {
		if upsert {
			object, err := ms.objectCreator(ctx, nil)
			if err != nil {
				return nil, err
			}
			utils.SetObjectFields(object, newVals)
			stor := object.(data.Storable)
			err = ms.Save(ctx, stor)
			if err != nil {
				return nil, err
			}
			return []string{stor.GetId()}, nil
		}
		return []string{}, nil
	}
	if ms.auditable {
		data.Audit(ctx, newVals)
	}
	_, err = connCopy.DB(ms.database).C(ms.collection).UpdateAll(queryCond, map[string]interface{}{"$set": newVals})
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	for _, id := range ids {
		common.InvalidateCache(ctx, ms.object, id)
		if ms.notifyupdates {
			common.NotifyUpdate(ctx, ms.object, id)
		}
	}
	return ids, err
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (ms *mongoDataService) UpdateMulti(ctx core.RequestContext, ids []string, newVals map[string]interface{}) error {
	if ms.auditable {
		data.Audit(ctx, newVals)
	}
	if ms.presave {
		err := ctx.SendSynchronousMessage(common.CONF_PREUPDATE_MSG, map[string]interface{}{"ids": ids, "type": ms.object, "data": newVals})
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	updateInterface := map[string]interface{}{"$set": newVals}
	condition, _ := ms.CreateCondition(ctx, data.MATCHMULTIPLEVALUES, ms.objectid, ids)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	_, err := connCopy.DB(ms.database).C(ms.collection).UpdateAll(condition, updateInterface)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	for _, id := range ids {
		common.InvalidateCache(ctx, ms.object, id)
		if ms.notifyupdates {
			common.NotifyUpdate(ctx, ms.object, id)
		}
	}
	return nil
}

//item must support Deleted field for soft deletes
func (ms *mongoDataService) Delete(ctx core.RequestContext, id string) error {
	if ms.softdelete {
		err := ms.Update(ctx, id, map[string]interface{}{ms.softDeleteField: true})
		if err == nil {
			err = common.DeleteRefOps(ctx, ms.deleteRefOpers, []string{id})
		}
		return err
	}
	condition := bson.M{}
	condition[ms.objectid] = id
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	err := connCopy.DB(ms.database).C(ms.collection).Remove(condition)
	if err == nil {
		err = common.DeleteRefOps(ctx, ms.deleteRefOpers, []string{id})
	}
	common.InvalidateCache(ctx, ms.object, id)
	return err
}

//Delete object by ids
func (ms *mongoDataService) DeleteMulti(ctx core.RequestContext, ids []string) error {
	if ms.softdelete {
		err := ms.UpdateMulti(ctx, ids, map[string]interface{}{ms.softDeleteField: true})
		if err == nil {
			err = common.DeleteRefOps(ctx, ms.deleteRefOpers, ids)
		}
		return err
	}
	conditionVal := bson.M{}
	conditionVal["$in"] = ids
	condition := bson.M{}
	condition[ms.objectid] = conditionVal
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	_, err := connCopy.DB(ms.database).C(ms.collection).RemoveAll(condition)
	if err == nil {
		err = common.DeleteRefOps(ctx, ms.deleteRefOpers, ids)
	}
	for _, id := range ids {
		common.InvalidateCache(ctx, ms.object, id)
	}
	return err
}

//Delete object by condition
func (ms *mongoDataService) DeleteAll(ctx core.RequestContext, queryCond interface{}) ([]string, error) {
	if ms.softdelete {
		ids, err := ms.UpdateAll(ctx, queryCond, map[string]interface{}{ms.softDeleteField: true})
		if err == nil {
			err = common.DeleteRefOps(ctx, ms.deleteRefOpers, ids)
		}
		return ids, err
	}
	results, err := ms.objectCollectionCreator(ctx, 0, nil)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	query := connCopy.DB(ms.database).C(ms.collection).Find(queryCond)
	err = query.All(results)
	resultStor, ids, err := data.CastToStorableCollection(results)
	length := len(resultStor)
	if length == 0 {
		return nil, nil
	}
	_, err = connCopy.DB(ms.database).C(ms.collection).RemoveAll(queryCond)
	if err == nil {
		err = common.DeleteRefOps(ctx, ms.deleteRefOpers, ids)
		if err != nil {
			return ids, err
		}
		for _, id := range ids {
			common.InvalidateCache(ctx, ms.object, id)
		}
	} else {
		return nil, err
	}
	return ids, err
}
