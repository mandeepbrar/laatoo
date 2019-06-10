package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"

	"gopkg.in/mgo.v2/bson"
)

type mongoDataService struct {
	*data.BaseComponent
	conf        config.Config
	name        string
	database    string
	factory     *mongoDataServicesFactory
	collection  string
	serviceType string
}

const (
	CONF_MONGO_DATABASE = "mongodatabase"
)

func newMongoDataService(ctx core.ServerContext, name string, ms *mongoDataServicesFactory) (*mongoDataService, error) {
	mongoSvc := &mongoDataService{BaseComponent: &data.BaseComponent{}, name: name, factory: ms, serviceType: "Mongo"}
	return mongoSvc, nil
}

func (svc *mongoDataService) Describe(ctx core.ServerContext) error {
	err := svc.BaseComponent.Describe(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	svc.AddOptionalConfigurations(ctx, map[string]string{data.CONF_DATA_COLLECTION: config.OBJECTTYPE_STRING}, nil)
	svc.SetDescription(ctx, "Mongo data component")
	return nil
}

func (svc *mongoDataService) Initialize(ctx core.ServerContext, conf config.Config) error {
	ctx = ctx.SubContext("Initialize Mongo Service")

	err := svc.BaseComponent.Initialize(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	collection, ok := svc.GetConfiguration(ctx, data.CONF_DATA_COLLECTION)
	if ok {
		svc.collection = collection.(string)
	} else {
		svc.collection = svc.ObjectConfig.Collection
	}

	if svc.collection == "" {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Missing Conf", data.CONF_DATA_COLLECTION)
	}
	svc.database = svc.factory.database
	return nil
}

func (svc *mongoDataService) Start(ctx core.ServerContext) error {
	err := svc.BaseComponent.Start(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	return nil
}

func (fs *mongoDataService) GetName() string {
	return fs.name
}

func (ms *mongoDataService) GetDataServiceType() string {
	return data.DATASERVICE_TYPE_NOSQL
}

func (ms *mongoDataService) GetObject() string {
	return ms.Object
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
	ctx = ctx.SubContext("Save")
	log.Trace(ctx, "Saving object", "Object", ms.Object)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	id := item.GetId()
	if id == "" {
		item.Init(ctx, nil)
	}
	if ms.PreSave {
		err := ctx.SendSynchronousMessage(data.CONF_PRESAVE_MSG, item)
		if err != nil {
			return err
		}
		err = item.PreSave(ctx)
		if err != nil {
			return err
		}
	}
	if ms.Multitenant {
		item.(data.StorableMT).SetTenant(ctx.GetUser().GetTenant())
	}
	if ms.Auditable {
		data.Audit(ctx, item)
	}
	err := connCopy.DB(ms.database).C(ms.collection).Insert(item)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if ms.PostSave {
		err = item.PostSave(ctx)
		if err != nil {
			errors.WrapError(ctx, err)
		}
		err := ctx.SendSynchronousMessage(data.CONF_NEWOBJ_MSG, item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ms *mongoDataService) CreateMulti(ctx core.RequestContext, items []data.Storable) error {
	return ms.PutMulti(ctx, items)
}

func (ms *mongoDataService) PutMulti(ctx core.RequestContext, items []data.Storable) error {
	ctx = ctx.SubContext("PutMulti")
	if ms.Multitenant {
		for _, item := range items {
			if item.(data.StorableMT).GetTenant() != ctx.GetUser().GetTenant() {
				return errors.ThrowError(ctx, errors.CORE_ERROR_TENANT_MISMATCH, "Provided tenant", item.(data.StorableMT).GetTenant(), "Item", item.GetId())
			}
		}
	}
	log.Trace(ctx, "Saving multiple objects", "ObjectType", ms.Object)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	bulk := connCopy.DB(ms.database).C(ms.collection).Bulk()
	for _, item := range items {
		if item != nil {
			if ms.PreSave {
				err := ctx.SendSynchronousMessage(data.CONF_PRESAVE_MSG, item)
				if err != nil {
					return err
				}
				err = item.PreSave(ctx)
				if err != nil {
					return err
				}
			}
			if ms.Auditable {
				data.Audit(ctx, item)
			}
			bulk.Upsert(bson.M{ms.ObjectId: item.GetId()}, item)
		}
	}
	_, err := bulk.Run()
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if ms.PostSave {
		for _, item := range items {
			if ms.PostSave {
				err = item.PostSave(ctx)
				if err != nil {
					errors.WrapError(ctx, err)
				}
			}
		}
	}
	log.Trace(ctx, "Saved multiple objects")
	return nil
}

func (ms *mongoDataService) Put(ctx core.RequestContext, id string, item data.Storable) error {
	ctx = ctx.SubContext("Put")
	if ms.Multitenant && (item.(data.StorableMT).GetTenant() != ctx.GetUser().GetTenant()) {
		return errors.ThrowError(ctx, errors.CORE_ERROR_TENANT_MISMATCH, "Provided tenant", item.(data.StorableMT).GetTenant(), "Item", item.GetId())
	}
	log.Trace(ctx, "Putting object", "ObjectType", ms.Object, "id", id)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	condition := bson.M{}
	condition[ms.ObjectId] = id
	item.SetId(id)
	if ms.PreSave {
		err := ctx.SendSynchronousMessage(data.CONF_PRESAVE_MSG, item)
		if err != nil {
			return err
		}
		err = item.PreSave(ctx)
		log.Trace(ctx, "Putting object", "err", err)
		if err != nil {
			return err
		}
	}
	if ms.Auditable {
		data.Audit(ctx, item)
	}
	_, err := connCopy.DB(ms.database).C(ms.collection).Upsert(condition, item)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if ms.PostSave {
		err = item.PostSave(ctx)
		if err != nil {
			errors.WrapError(ctx, err)
		}
	}
	return nil
}

//upsert an object ...insert if not there... update if there
func (ms *mongoDataService) UpsertId(ctx core.RequestContext, id string, newVals map[string]interface{}) error {
	ctx = ctx.SubContext("UpsertId")
	return ms.update(ctx, id, newVals, true)
}

func (ms *mongoDataService) Update(ctx core.RequestContext, id string, newVals map[string]interface{}) error {
	ctx = ctx.SubContext("Update")
	return ms.update(ctx, id, newVals, false)
}

func (ms *mongoDataService) update(ctx core.RequestContext, id string, newVals map[string]interface{}, upsert bool) error {
	if ms.Auditable {
		data.Audit(ctx, newVals)
	}
	if ms.PreSave {
		err := ctx.SendSynchronousMessage(data.CONF_PREUPDATE_MSG, map[string]interface{}{"id": id, "type": ms.Object, "data": newVals})
		if err != nil {
			return err
		}
	}
	if upsert {
		newVals[ms.ObjectId] = id
	}
	condition := bson.M{}
	condition[ms.ObjectId] = id
	if ms.Multitenant {
		condition["Tenant"] = ctx.GetUser().GetTenant()
	}
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	var err error
	if upsert {
		_, err = connCopy.DB(ms.database).C(ms.collection).Upsert(condition, newVals)

	} else {
		updateInterface := map[string]interface{}{"$set": newVals}
		err = connCopy.DB(ms.database).C(ms.collection).Update(condition, updateInterface)
	}
	if err != nil {
		return err
	}
	if ms.PostUpdate {
		err = ctx.SendSynchronousMessage(data.CONF_POSTUPDATE_MSG, map[string]interface{}{"id": id, "type": ms.Object, "data": newVals})
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}

func (ms *mongoDataService) Upsert(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	ctx = ctx.SubContext("Upsert")
	return ms.updateAll(ctx, queryCond, newVals, true)
}

func (ms *mongoDataService) UpdateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	ctx = ctx.SubContext("UpdateAll")
	return ms.updateAll(ctx, queryCond, newVals, false)
}

func (ms *mongoDataService) updateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}, upsert bool) ([]string, error) {
	results := ms.ObjectCollectionCreator(0)

	if ms.PreSave {
		err := ctx.SendSynchronousMessage(data.CONF_PREUPDATE_MSG, map[string]interface{}{"cond": queryCond, "type": ms.Object, "data": newVals})
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
	}
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	query := connCopy.DB(ms.database).C(ms.collection).Find(queryCond)
	err := query.All(results)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	resultStor, ids, err := data.CastToStorableCollection(results)
	length := len(resultStor)
	if length == 0 {
		if upsert {
			object := ms.ObjectCreator()
			stor := object.(data.Storable)
			stor.SetValues(object, newVals)
			err := ms.Save(ctx, stor)
			if err != nil {
				return nil, err
			}
			return []string{stor.GetId()}, nil
		}
		return []string{}, nil
	}
	if ms.Auditable {
		data.Audit(ctx, newVals)
	}
	_, err = connCopy.DB(ms.database).C(ms.collection).UpdateAll(queryCond, map[string]interface{}{"$set": newVals})
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	if ms.PostUpdate {
		for _, id := range ids {
			err = ctx.SendSynchronousMessage(data.CONF_POSTUPDATE_MSG, map[string]interface{}{"id": id, "type": ms.Object, "data": newVals})
			if err != nil {
				return nil, errors.WrapError(ctx, err)
			}
		}
	}
	return ids, err
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (ms *mongoDataService) UpdateMulti(ctx core.RequestContext, ids []string, newVals map[string]interface{}) error {
	ctx = ctx.SubContext("UpdateMulti")
	if ms.Auditable {
		data.Audit(ctx, newVals)
	}
	if ms.PreSave {
		err := ctx.SendSynchronousMessage(data.CONF_PREUPDATE_MSG, map[string]interface{}{"ids": ids, "type": ms.Object, "data": newVals})
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	updateInterface := map[string]interface{}{"$set": newVals}
	condition, _ := ms.CreateCondition(ctx, data.MATCHMULTIPLEVALUES, ms.ObjectId, ids)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	_, err := connCopy.DB(ms.database).C(ms.collection).UpdateAll(condition, updateInterface)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if ms.PostUpdate {
		for _, id := range ids {
			err = ctx.SendSynchronousMessage(data.CONF_POSTUPDATE_MSG, map[string]interface{}{"id": id, "type": ms.Object, "data": newVals})
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	}
	return nil
}

//item must support Deleted field for soft deletes
func (ms *mongoDataService) Delete(ctx core.RequestContext, id string) error {
	ctx = ctx.SubContext("Delete")
	if ms.SoftDelete {
		return ms.Update(ctx, id, map[string]interface{}{ms.SoftDeleteField: true})
	}
	condition := bson.M{}
	if ms.Multitenant {
		condition["Tenant"] = ctx.GetUser().GetTenant()
	}
	condition[ms.ObjectId] = id
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	err := connCopy.DB(ms.database).C(ms.collection).Remove(condition)
	return err
}

//Delete object by ids
func (ms *mongoDataService) DeleteMulti(ctx core.RequestContext, ids []string) error {
	ctx = ctx.SubContext("DeleteMulti")
	if ms.SoftDelete {
		return ms.UpdateMulti(ctx, ids, map[string]interface{}{ms.SoftDeleteField: true})
	}
	conditionVal := bson.M{}
	conditionVal["$in"] = ids
	condition := bson.M{}
	if ms.Multitenant {
		condition["Tenant"] = ctx.GetUser().GetTenant()
	}
	condition[ms.ObjectId] = conditionVal
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	_, err := connCopy.DB(ms.database).C(ms.collection).RemoveAll(condition)
	return err
}

//Delete object by condition
func (ms *mongoDataService) DeleteAll(ctx core.RequestContext, queryCond interface{}) ([]string, error) {
	ctx = ctx.SubContext("DeleteAll")
	if ms.SoftDelete {
		return ms.UpdateAll(ctx, queryCond, map[string]interface{}{ms.SoftDeleteField: true})
	}
	results := ms.ObjectCollectionCreator(0)
	connCopy := ms.factory.connection.Copy()
	defer connCopy.Close()
	query := connCopy.DB(ms.database).C(ms.collection).Find(queryCond)
	err := query.All(results)
	resultStor, ids, err := data.CastToStorableCollection(results)
	length := len(resultStor)
	if length == 0 {
		return nil, nil
	}
	_, err = connCopy.DB(ms.database).C(ms.collection).RemoveAll(queryCond)
	return ids, err
}
