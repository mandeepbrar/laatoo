package sql

import (
	"laatoo/framework/services/data/common"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"

	"github.com/jinzhu/gorm"

	//	"laatoo/sdk/utils"
)

type sqlDataService struct {
	*data.BaseComponent
	conf            config.Config
	connection      *gorm.DB
	db              *gorm.DB
	name            string
	database        string
	auditable       bool
	softdelete      bool
	presave         bool
	postsave        bool
	postload        bool
	notifynew       bool
	notifyupdates   bool
	factory         *sqlDataServicesFactory
	collection      string
	softDeleteField string
	objectid        string
	serviceType     string
}

func newSqlDataService(ctx core.ServerContext, name string, svc *sqlDataServicesFactory) (*sqlDataService, error) {
	sqlDataSvc := &sqlDataService{BaseComponent: &data.BaseComponent{}, name: name, factory: svc, serviceType: "SQL"}
	return sqlDataSvc, nil
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
		return errors.RethrowError(ctx, common.DATA_ERROR_CONNECTION, err, "Connection String", svc.factory.connectionString)
	}
	svc.connection = sess
	svc.collection = svc.connection.NewScope(object).GetModelStruct().TableName(svc.connection)
	if svc.collection == "" {
		return errors.MissingConf(ctx, common.CONF_DATA_COLLECTION)
	}
	svc.db = svc.connection.Table(svc.collection)
	svc.db.LogMode(true)
	svc.objectid = svc.ObjectConfig.IdField
	svc.softDeleteField = svc.ObjectConfig.SoftDeleteField

	if svc.softDeleteField == "" {
		svc.softdelete = false
	} else {
		svc.softdelete = true
	}

	auditable, ok := conf.GetBool(common.CONF_DATA_AUDITABLE)
	if ok {
		svc.auditable = auditable
	} else {
		svc.auditable = svc.ObjectConfig.Auditable
	}
	postsave, ok := conf.GetBool(common.CONF_DATA_POSTSAVE)
	if ok {
		svc.postsave = postsave
	} else {
		svc.postsave = svc.ObjectConfig.PostSave
	}
	presave, ok := conf.GetBool(common.CONF_DATA_PRESAVE)
	if ok {
		svc.presave = presave
	} else {
		svc.presave = svc.ObjectConfig.PreSave
	}
	postload, ok := conf.GetBool(common.CONF_DATA_POSTLOAD)
	if ok {
		svc.postload = postload
	} else {
		svc.postload = svc.ObjectConfig.PostLoad
	}

	notifyupdates, ok := conf.GetBool(common.CONF_DATA_NOTIFYUPDATES)
	if ok {
		svc.notifyupdates = notifyupdates
	} else {
		svc.notifyupdates = svc.ObjectConfig.NotifyUpdates
	}
	notifynew, ok := conf.GetBool(common.CONF_DATA_NOTIFYNEW)
	if ok {
		svc.notifynew = notifynew
	} else {
		svc.notifynew = svc.ObjectConfig.NotifyNew
	}
	return nil
}

func (svc *sqlDataService) Start(ctx core.ServerContext) error {
	return nil
}

func (svc *sqlDataService) Invoke(ctx core.RequestContext) error {
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
	log.Logger.Trace(ctx, "Saving object", "Object", svc.Object)
	id := item.GetId()
	if id == "" {
		item.Init(ctx, nil)
	}
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
	err := svc.db.Create(item).Error
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

func (svc *sqlDataService) CreateMulti(ctx core.RequestContext, items []data.Storable) error {
	return svc.putMulti(ctx, items, true)
}
func (svc *sqlDataService) PutMulti(ctx core.RequestContext, items []data.Storable) error {
	return svc.putMulti(ctx, items, false)
}
func (svc *sqlDataService) putMulti(ctx core.RequestContext, items []data.Storable, createNew bool) error {
	ctx = ctx.SubContext("PutMulti")
	for _, item := range items {
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
	for _, item := range items {
		id := item.GetId()
		if svc.postsave || svc.notifyupdates {
			if svc.postsave {
				err := item.PostSave(ctx)
				if err != nil {
					errors.WrapError(ctx, err)
				}
			}
			if svc.notifyupdates {
				common.NotifyUpdate(ctx, svc.Object, id)
			}
		}
	}
	log.Logger.Trace(ctx, "Saved multiple objects")
	return nil
}

func (svc *sqlDataService) Put(ctx core.RequestContext, id string, item data.Storable) error {
	ctx = ctx.SubContext("Put")
	log.Logger.Trace(ctx, "Putting object", "ObjectType", svc.Object, "id", id)
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

	err := svc.db.Save(item).Error
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
		common.NotifyUpdate(ctx, svc.Object, id)
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
		err := ctx.SendSynchronousMessage(common.CONF_PREUPDATE_MSG, map[string]interface{}{"id": id, "type": svc.Object, "data": newVals})
		if err != nil {
			return err
		}
	}
	if upsert {
		newVals[svc.objectid] = id
	}
	var err error
	if upsert {
		err = svc.db.Where([]string{id}).FirstOrCreate(newVals).Error
	} else {
		err = svc.db.Where([]string{id}).Updates(newVals).Error
	}
	if err != nil {
		return err
	}
	if svc.notifyupdates {
		common.NotifyUpdate(ctx, svc.Object, id)
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
	if svc.presave {
		err := ctx.SendSynchronousMessage(common.CONF_PREUPDATE_MSG, map[string]interface{}{"cond": queryCond, "type": svc.Object, "data": newVals})
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
	if svc.auditable {
		data.Audit(ctx, newVals)
	}
	log.Logger.Error(ctx, "Field details", "newVals", newVals)
	err = query.Updates(newVals).Error
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	if svc.notifyupdates {
		for _, id := range ids {
			common.NotifyUpdate(ctx, svc.Object, id)
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
		err := ctx.SendSynchronousMessage(common.CONF_PREUPDATE_MSG, map[string]interface{}{"ids": ids, "type": svc.Object, "data": newVals})
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	err := svc.db.Where(ids).Updates(newVals).Error
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if svc.notifyupdates {
		for _, id := range ids {
			common.NotifyUpdate(ctx, svc.Object, id)
		}
	}
	return nil
}

//item must support Deleted field for soft deletes
func (svc *sqlDataService) Delete(ctx core.RequestContext, id string) error {
	ctx = ctx.SubContext("Delete")
	if svc.softdelete {
		return svc.Update(ctx, id, map[string]interface{}{svc.softDeleteField: true})
	}
	return svc.db.Delete(id).Error
}

//Delete object by ids
func (svc *sqlDataService) DeleteMulti(ctx core.RequestContext, ids []string) error {
	ctx = ctx.SubContext("DeleteMulti")
	if svc.softdelete {
		return svc.UpdateMulti(ctx, ids, map[string]interface{}{svc.softDeleteField: true})
	}
	return svc.db.Delete(ids).Error
}

//Delete object by condition
func (svc *sqlDataService) DeleteAll(ctx core.RequestContext, queryCond interface{}) ([]string, error) {
	ctx = ctx.SubContext("DeleteAll")
	if svc.softdelete {
		return svc.UpdateAll(ctx, queryCond, map[string]interface{}{svc.softDeleteField: true})
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
