package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"

	"github.com/jinzhu/gorm"
)

type sqlDataService struct {
	*data.BaseComponent
	conf        config.Config
	connection  *gorm.DB
	db          *gorm.DB
	name        string
	database    string
	factory     *sqlDataServicesFactory
	collection  string
	serviceType string
}

func newSqlDataService(ctx core.ServerContext, name string, svc *sqlDataServicesFactory) (*sqlDataService, error) {
	sqlDataSvc := &sqlDataService{BaseComponent: &data.BaseComponent{}, name: name, factory: svc, serviceType: "SQL"}
	return sqlDataSvc, nil
}

func (svc *sqlDataService) Describe(ctx core.ServerContext) error {
	err := svc.BaseComponent.Describe(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	svc.SetDescription(ctx, "SQL data component")
	return nil
}

func (svc *sqlDataService) Initialize(ctx core.ServerContext, conf config.Config) error {
	ctx = ctx.SubContext("Initialize SQL Data Service")

	err := svc.BaseComponent.Initialize(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	object := svc.ObjectCreator()
	sess, err := gorm.Open(svc.factory.vendor, svc.factory.connectionString)
	if err != nil {
		return errors.RethrowError(ctx, data.DATA_ERROR_CONNECTION, err, "Connection String", svc.factory.connectionString)
	}
	svc.connection = sess
	svc.collection = svc.connection.NewScope(object).GetModelStruct().TableName(svc.connection)
	if svc.collection == "" {
		return errors.MissingConf(ctx, data.CONF_DATA_COLLECTION)
	}
	svc.db = svc.connection.Table(svc.collection)
	svc.db.LogMode(true)

	return nil
}

func (svc *sqlDataService) Start(ctx core.ServerContext) error {
	err := svc.BaseComponent.Start(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	return nil
}

func (svc *sqlDataService) CreateDBCollection(ctx core.RequestContext) error {
	object := svc.ObjectCreator()
	return svc.connection.CreateTable(object).Error
}

func (svc *sqlDataService) DropDBCollection(ctx core.RequestContext) error {
	object := svc.ObjectCreator()
	return svc.connection.DropTable(object).Error
}

func (svc *sqlDataService) DBCollectionExists(ctx core.RequestContext) (bool, error) {
	object := svc.ObjectCreator()
	return svc.connection.HasTable(object), nil
}

func (svc *sqlDataService) GetName() string {
	return svc.name
}

func (svc *sqlDataService) GetDataServiceType() string {
	return data.DATASERVICE_TYPE_NOSQL
}

func (svc *sqlDataService) GetObject() string {
	return svc.Object
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
	log.Trace(ctx, "Saving object", "Object", svc.Object)
	id := item.GetId()
	if id == "" {
		item.Initialize(ctx, nil)
	}
	if svc.PreSave {
		err := ctx.SendSynchronousMessage(data.CONF_PRESAVE_MSG, item)
		if err != nil {
			return err
		}
		err = item.PreSave(ctx)
		if err != nil {
			return err
		}
	}
	if svc.Multitenant {
		item.(data.StorableMT).SetTenant(ctx.GetUser().GetTenant())
	}
	if svc.Auditable {
		data.Audit(ctx, item)
	}
	err := svc.db.Create(item).Error
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if svc.PostSave {
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

func (svc *sqlDataService) CreateMulti(ctx core.RequestContext, items []data.Storable) error {
	return svc.putMulti(ctx, items, true)
}
func (svc *sqlDataService) PutMulti(ctx core.RequestContext, items []data.Storable) error {
	return svc.putMulti(ctx, items, false)
}
func (svc *sqlDataService) putMulti(ctx core.RequestContext, items []data.Storable, createNew bool) error {
	ctx = ctx.SubContext("PutMulti")
	if svc.Multitenant {
		for _, item := range items {
			if item.(data.StorableMT).GetTenant() != ctx.GetUser().GetTenant() {
				return errors.ThrowError(ctx, errors.CORE_ERROR_TENANT_MISMATCH, "Provided tenant", item.(data.StorableMT).GetTenant(), "Item", item.GetId())
			}
		}
	}
	for _, item := range items {
		if svc.PreSave {
			err := ctx.SendSynchronousMessage(data.CONF_PRESAVE_MSG, item)
			if err != nil {
				return err
			}
			err = item.PreSave(ctx)
			if err != nil {
				return err
			}
		}
		if svc.Auditable {
			data.Audit(ctx, item)
		}
	}
	for _, item := range items {
		var err error
		if createNew {
			err = svc.db.Create(item).Error
		} else {
			err = svc.db.Save(item).Error
		}
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	if svc.PostSave {
		for _, item := range items {
			err := item.PostSave(ctx)
			if err != nil {
				errors.WrapError(ctx, err)
			}
		}
	}
	log.Trace(ctx, "Saved multiple objects")
	return nil
}

func (svc *sqlDataService) Put(ctx core.RequestContext, id string, item data.Storable) error {
	ctx = ctx.SubContext("Put")
	if svc.Multitenant && (item.(data.StorableMT).GetTenant() != ctx.GetUser().GetTenant()) {
		return errors.ThrowError(ctx, errors.CORE_ERROR_TENANT_MISMATCH, "Provided tenant", item.(data.StorableMT).GetTenant(), "Item", item.GetId())
	}
	log.Trace(ctx, "Putting object", "ObjectType", svc.Object, "id", id)
	item.SetId(id)
	if svc.PreSave {
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
	if svc.Auditable {
		data.Audit(ctx, item)
	}

	err := svc.db.Save(item).Error
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if svc.PostSave {
		err = item.PostSave(ctx)
		if err != nil {
			errors.WrapError(ctx, err)
		}
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
	if svc.Auditable {
		data.Audit(ctx, newVals)
	}

	var err error
	if svc.PreSave {
		err = svc.updateWithPresave(ctx, id, newVals, upsert)
	} else {
		err = svc.updateWithoutPresave(ctx, id, newVals, upsert)
	}
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	if svc.PostUpdate {
		err = ctx.SendSynchronousMessage(data.CONF_POSTUPDATE_MSG, map[string]interface{}{"id": id, "type": svc.Object, "data": newVals})
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}

func (svc *sqlDataService) updateWithoutPresave(ctx core.RequestContext, id string, newVals map[string]interface{}, upsert bool) error {
	if upsert {
		newVals[svc.ObjectId] = id
	}
	query := svc.getMultitenantQuery(ctx, svc.db)
	var err error
	if upsert {
		err = query.Where([]string{id}).FirstOrCreate(newVals).Error
	} else {
		err = query.Where([]string{id}).Updates(newVals).Error
	}
	if err != nil {
		errors.WrapError(ctx, err)
	}
	return nil
}

func (svc *sqlDataService) updateWithPresave(ctx core.RequestContext, id string, newVals map[string]interface{}, upsert bool) error {
	err := ctx.SendSynchronousMessage(data.CONF_PREUPDATE_MSG, map[string]interface{}{"id": id, "type": svc.Object, "data": newVals})
	if err != nil {
		return err
	}
	if upsert {
		newVals[svc.ObjectId] = id
	}
	query := svc.getMultitenantQuery(ctx, svc.db)

	object := svc.ObjectCreator()

	err = query.First(object, id).Error

	if err != nil {
		return errors.WrapError(ctx, err)
	}
	stor := object.(data.Storable)

	if svc.Multitenant && (stor.(data.StorableMT).GetTenant() != ctx.GetUser().GetTenant()) {
		return errors.ThrowError(ctx, errors.CORE_ERROR_TENANT_MISMATCH, "Provided tenant", ctx.GetUser().GetTenant(), "Item", id)
	}

	log.Info(ctx, "Going to set values", "stor", stor, "newVals", newVals)
	stor.SetValues(object, newVals)
	if upsert {
		err = svc.Save(ctx, stor)
	} else {
		err = svc.Put(ctx, stor.GetId(), stor)
	}
	if err != nil {
		return errors.WrapError(ctx, err)
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
	results := svc.ObjectCollectionCreator(0)
	if svc.PreSave {
		err := ctx.SendSynchronousMessage(data.CONF_PREUPDATE_MSG, map[string]interface{}{"cond": queryCond, "type": svc.Object, "data": newVals})
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
	}
	query, err := svc.processCondition(ctx, queryCond, svc.db)
	if err != nil {
		return nil, err
	}
	err = query.Find(results).Error
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	resultStor, ids, err := data.CastToStorableCollection(results)
	length := len(resultStor)
	if length == 0 {
		if upsert {
			object := svc.ObjectCreator()
			stor := object.(data.Storable)
			stor.SetValues(object, newVals)
			err := svc.Save(ctx, stor)
			if err != nil {
				return nil, err
			}
			return []string{stor.GetId()}, nil
		}
		return []string{}, nil
	}
	if svc.Auditable {
		data.Audit(ctx, newVals)
	}
	log.Error(ctx, "Field details", "newVals", newVals)
	err = query.Updates(newVals).Error
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	if svc.PostUpdate {
		for _, id := range ids {
			err := ctx.SendSynchronousMessage(data.CONF_POSTUPDATE_MSG, map[string]interface{}{"id": id, "type": svc.Object, "data": newVals})
			if err != nil {
				return nil, errors.WrapError(ctx, err)
			}
		}
	}
	return ids, err
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (svc *sqlDataService) UpdateMulti(ctx core.RequestContext, ids []string, newVals map[string]interface{}) error {
	ctx = ctx.SubContext("UpdateMulti")
	if svc.Auditable {
		data.Audit(ctx, newVals)
	}
	if svc.PreSave {
		err := ctx.SendSynchronousMessage(data.CONF_PREUPDATE_MSG, map[string]interface{}{"ids": ids, "type": svc.Object, "data": newVals})
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	query := svc.getMultitenantQuery(ctx, svc.db)
	err := query.Where(ids).Updates(newVals).Error
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if svc.PostUpdate {
		for _, id := range ids {
			err := ctx.SendSynchronousMessage(data.CONF_POSTUPDATE_MSG, map[string]interface{}{"id": id, "type": svc.Object, "data": newVals})
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	}
	return nil
}

//item must support Deleted field for soft deletes
func (svc *sqlDataService) Delete(ctx core.RequestContext, id string) error {
	ctx = ctx.SubContext("Delete")
	if svc.SoftDelete {
		return svc.Update(ctx, id, map[string]interface{}{svc.SoftDeleteField: true})
	}
	query := svc.getMultitenantQuery(ctx, svc.db)
	return query.Delete(id).Error
}

//Delete object by ids
func (svc *sqlDataService) DeleteMulti(ctx core.RequestContext, ids []string) error {
	ctx = ctx.SubContext("DeleteMulti")
	if svc.SoftDelete {
		return svc.UpdateMulti(ctx, ids, map[string]interface{}{svc.SoftDeleteField: true})
	}
	query := svc.getMultitenantQuery(ctx, svc.db)
	return query.Delete(ids).Error
}

//Delete object by condition
func (svc *sqlDataService) DeleteAll(ctx core.RequestContext, queryCond interface{}) ([]string, error) {
	ctx = ctx.SubContext("DeleteAll")
	if svc.SoftDelete {
		return svc.UpdateAll(ctx, queryCond, map[string]interface{}{svc.SoftDeleteField: true})
	}
	results := svc.ObjectCollectionCreator(0)
	query, err := svc.processCondition(ctx, queryCond, svc.db)
	if err != nil {
		return nil, err
	}
	err = query.Find(results).Error
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	resultStor, ids, err := data.CastToStorableCollection(results)
	length := len(resultStor)
	if length == 0 {
		return nil, nil
	}
	err = query.Delete(svc.ObjectCreator()).Error
	return ids, err
}
