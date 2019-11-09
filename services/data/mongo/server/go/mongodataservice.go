package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
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
	conn := ms.factory.getConnection(ctx)
	id := item.GetId()
	if id == "" {
		item.Initialize(ctx, nil)
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
		item.SetTenant(ctx.GetUser().GetTenant())
	}
	if ms.Auditable {
		data.Audit(ctx, item)
	}
	_, err := conn.Database(ms.database).Collection(ms.collection).InsertOne(ctx, item)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	//	ctx.Set("NewId", res.InsertedID)
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
			if item.GetTenant() != ctx.GetUser().GetTenant() {
				return errors.ThrowError(ctx, errors.CORE_ERROR_TENANT_MISMATCH, "Provided tenant", item.GetTenant(), "Item", item.GetId())
			}
		}
	}
	log.Trace(ctx, "Saving multiple objects", "ObjectType", ms.Object)
	conn := ms.factory.getConnection(ctx)

	var operations []mongo.WriteModel

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

			condition := bson.M{}
			condition[ms.ObjectId] = item.GetId()
			if ms.Multitenant {
				condition["Tenant"] = ctx.GetUser().GetTenant()
			}

			operation := mongo.NewUpdateOneModel()
			operation.SetFilter(condition)
			operation.SetUpdate(item)
			operation.SetUpsert(true)
			operations = append(operations, operation)

		}
	}
	_, err := conn.Database(ms.database).Collection(ms.collection).BulkWrite(ctx, operations)
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
	if ms.Multitenant && (item.GetTenant() != ctx.GetUser().GetTenant()) {
		return errors.ThrowError(ctx, errors.CORE_ERROR_TENANT_MISMATCH, "Provided tenant", item.GetTenant(), "Item", item.GetId())
	}
	log.Trace(ctx, "Putting object", "ObjectType", ms.Object, "id", id)
	conn := ms.factory.getConnection(ctx)

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

	_, err := conn.Database(ms.database).Collection(ms.collection).ReplaceOne(ctx, condition, item, options.Replace().SetUpsert(true))

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
	var err error
	if ms.PreSave {
		err = ms.updateWithPresave(ctx, id, newVals, upsert)
	} else {
		err = ms.updateWithoutPresave(ctx, id, newVals, upsert)
	}
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	if ms.PostUpdate {
		err = ctx.SendSynchronousMessage(data.CONF_POSTUPDATE_MSG, map[string]interface{}{"id": id, "type": ms.Object, "data": newVals})
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}

func (ms *mongoDataService) updateWithoutPresave(ctx core.RequestContext, id string, newVals map[string]interface{}, upsert bool) error {
	if upsert {
		newVals[ms.ObjectId] = id
	}
	condition := bson.M{}
	condition[ms.ObjectId] = id
	if ms.Multitenant {
		condition["Tenant"] = ctx.GetUser().GetTenant()
	}
	conn := ms.factory.getConnection(ctx)

	var err error
	if upsert {
		_, err = conn.Database(ms.database).Collection(ms.collection).UpdateOne(ctx, condition, newVals, options.Update().SetUpsert(true))

	} else {
		updateInterface := map[string]interface{}{"$set": newVals}
		_, err = conn.Database(ms.database).Collection(ms.collection).UpdateMany(ctx, condition, updateInterface)
	}
	return errors.WrapError(ctx, err)
}

func (ms *mongoDataService) updateWithPresave(ctx core.RequestContext, id string, newVals map[string]interface{}, upsert bool) error {

	err := ctx.SendSynchronousMessage(data.CONF_PREUPDATE_MSG, map[string]interface{}{"id": id, "type": ms.Object, "data": newVals})
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if upsert {
		newVals[ms.ObjectId] = id
	}
	condition := bson.M{}
	condition[ms.ObjectId] = id
	if ms.Multitenant {
		condition["Tenant"] = ctx.GetUser().GetTenant()
	}

	conn := ms.factory.getConnection(ctx)
	object, _ := ctx.CreateObject(ms.Object)

	err = conn.Database(ms.database).Collection(ms.collection).FindOne(ctx, condition).Decode(object)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	stor := object.(data.Storable)
	if ms.Multitenant && (stor.GetTenant() != ctx.GetUser().GetTenant()) {
		return errors.ThrowError(ctx, errors.CORE_ERROR_TENANT_MISMATCH, "Provided tenant", ctx.GetUser().GetTenant(), "Item", id)
	}

	log.Info(ctx, "Going to set values", "stor", stor, "newVals", newVals)
	stor.SetValues(object, newVals)
	if upsert {
		return ms.Save(ctx, stor)
	} else {
		return ms.Put(ctx, stor.GetId(), stor)
	}
}

func (ms *mongoDataService) Upsert(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}, getids bool) ([]string, error) {
	ctx = ctx.SubContext("Upsert")
	return ms.updateAll(ctx, queryCond, newVals, true, getids)
}

func (ms *mongoDataService) UpdateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}, getids bool) ([]string, error) {
	ctx = ctx.SubContext("UpdateAll")
	return ms.updateAll(ctx, queryCond, newVals, false, getids)
}

func (ms *mongoDataService) updateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}, upsert, getids bool) ([]string, error) {
	var ids []string
	var err error
	if ms.PreSave {
		err := ctx.SendSynchronousMessage(data.CONF_PREUPDATE_MSG, map[string]interface{}{"cond": queryCond, "type": ms.Object, "data": newVals})
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
	}

	conn := ms.factory.getConnection(ctx)
	if getids || ms.PostUpdate || upsert {

		cursor, err := conn.Database(ms.database).Collection(ms.collection).Find(ctx, queryCond)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		defer cursor.Close(ctx)

		countonly := !(getids || ms.PostUpdate)

		var count int

		_, ids, count, err = ms.getResultsFromCursor(ctx, cursor, countonly)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}

		if count == 0 {
			if upsert {
				object := ms.ObjectFactory.CreateObject(ctx)
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

	}

	/****** TODO audit****/
	if ms.Auditable {
		data.Audit(ctx, newVals)
	}

	_, err = conn.Database(ms.database).Collection(ms.collection).UpdateMany(ctx, queryCond, map[string]interface{}{"$set": newVals})

	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	if ms.PostUpdate {
		if ids != nil {
			for _, id := range ids {
				err = ctx.SendSynchronousMessage(data.CONF_POSTUPDATE_MSG, map[string]interface{}{"id": id, "type": ms.Object, "data": newVals})
				if err != nil {
					return nil, errors.WrapError(ctx, err)
				}
			}
		}
	}

	return ids, err
}

/*
	var operations []mongo.WriteModel

operation := mongo.NewUpdateOneModel()
operation.Filter(bson.NewDocument(
    bson.EC.SubDocumentFromElements("qty",
        bson.EC.Int32("$lt", 50),
    ),
))
operation.Update(bson.NewDocument(
    bson.EC.SubDocumentFromElements("$set",
        bson.EC.String("size.uom", "cm"),
        bson.EC.String("status", "P"),
    ),
    bson.EC.SubDocumentFromElements("$currentDate",
        bson.EC.Boolean("lastModified", true),
    ),
))

operations = append(operations, operation)

result, err := coll.BulkWrite(
    context.Background(),
    operations,
)*/

//update objects by ids, fields to be updated should be provided as key value pairs
func (ms *mongoDataService) UpdateMulti(ctx core.RequestContext, ids []string, newVals map[string]interface{}) error {
	ctx = ctx.SubContext("UpdateMulti")
	if ms.Auditable {
		//data.Audit(ctx, newVals)
	}
	if ms.PreSave {
		err := ctx.SendSynchronousMessage(data.CONF_PREUPDATE_MSG, map[string]interface{}{"ids": ids, "type": ms.Object, "data": newVals})
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	updateInterface := map[string]interface{}{"$set": newVals}
	condition, _ := ms.CreateCondition(ctx, data.MATCHMULTIPLEVALUES, ms.ObjectId, ids)

	conn := ms.factory.getConnection(ctx)
	_, err := conn.Database(ms.database).Collection(ms.collection).UpdateMany(ctx, condition, updateInterface)

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
	conn := ms.factory.getConnection(ctx)
	_, err := conn.Database(ms.database).Collection(ms.collection).DeleteOne(ctx, condition)
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

	conn := ms.factory.getConnection(ctx)
	_, err := conn.Database(ms.database).Collection(ms.collection).DeleteMany(ctx, condition)
	return err
}

//Delete object by condition
func (ms *mongoDataService) DeleteAll(ctx core.RequestContext, queryCond interface{}, getids bool) ([]string, error) {
	ctx = ctx.SubContext("DeleteAll")
	if ms.SoftDelete {
		return ms.UpdateAll(ctx, queryCond, map[string]interface{}{ms.SoftDeleteField: true}, getids)
	}
	conn := ms.factory.getConnection(ctx)
	var ids []string
	var count int
	var err error
	if getids {
		cursor, err := conn.Database(ms.database).Collection(ms.collection).Find(ctx, queryCond)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		defer cursor.Close(ctx)

		_, ids, count, err = ms.getResultsFromCursor(ctx, cursor, false)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		if count == 0 {
			return nil, nil
		}
	}

	_, err = conn.Database(ms.database).Collection(ms.collection).DeleteMany(ctx, queryCond)

	return ids, err
}
