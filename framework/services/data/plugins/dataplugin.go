package plugins

/*
import (
	"laatoo/framework/services/data/common"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
)

type dataPlugin struct {
	*data.BaseComponent
	dataServiceName string
	DataComponent   data.DataComponent
}

func NewDataPlugin(ctx core.ServerContext) *dataPlugin {
	return &dataPlugin{BaseComponent: &data.BaseComponent{}}
}

func (svc *dataPlugin) Initialize(ctx core.ServerContext, conf config.Config) error {
	err := svc.BaseComponent.Initialize(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	bsSvc, ok := conf.GetString(common.CONF_BASE_SVC)
	if !ok {
		return errors.MissingConf(ctx, common.CONF_BASE_SVC)
	}
	svc.dataServiceName = bsSvc
	return nil
}

func (svc *dataPlugin) Start(ctx core.ServerContext) error {
	s, err := ctx.GetService(svc.dataServiceName)
	if err != nil {
		return errors.BadConf(ctx, common.CONF_BASE_SVC)
	}
	DataComponent, ok := s.(data.DataComponent)
	if !ok {
		return errors.BadConf(ctx, common.CONF_BASE_SVC)
	}
	svc.DataComponent = DataComponent
	return nil
}

func (svc *dataPlugin) Invoke(ctx core.RequestContext) error {
	return nil
}

func (svc *dataPlugin) CreateDBCollection(ctx core.RequestContext) error {
	return svc.DataComponent.CreateDBCollection(ctx)
}

func (svc *dataPlugin) DropDBCollection(ctx core.RequestContext) error {
	return svc.DataComponent.DropDBCollection(ctx)
}

func (svc *dataPlugin) DBCollectionExists(ctx core.RequestContext) (bool, error) {
	return svc.DataComponent.DBCollectionExists(ctx)
}

func (svc *dataPlugin) GetDataServiceType() string {
	return svc.DataComponent.GetDataServiceType()
}

func (svc *dataPlugin) Supports(feature data.Feature) bool {
	return svc.DataComponent.Supports(feature)
}

func (svc *dataPlugin) Save(ctx core.RequestContext, item data.Storable) error {
	return svc.DataComponent.Save(ctx, item)
}

func (svc *dataPlugin) PutMulti(ctx core.RequestContext, items []data.Storable) error {
	return svc.DataComponent.PutMulti(ctx, items)
}

func (svc *dataPlugin) Put(ctx core.RequestContext, id string, item data.Storable) error {
	return svc.DataComponent.Put(ctx, id, item)
}

//upsert an object ...insert if not there... update if there
func (svc *dataPlugin) UpsertId(ctx core.RequestContext, id string, newVals map[string]interface{}) error {
	return svc.DataComponent.UpsertId(ctx, id, newVals)
}

func (svc *dataPlugin) Update(ctx core.RequestContext, id string, newVals map[string]interface{}) error {
	return svc.DataComponent.Update(ctx, id, newVals)
}

func (svc *dataPlugin) Upsert(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	return svc.DataComponent.Upsert(ctx, queryCond, newVals)
}

func (svc *dataPlugin) UpdateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	return svc.DataComponent.UpdateAll(ctx, queryCond, newVals)
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (svc *dataPlugin) UpdateMulti(ctx core.RequestContext, ids []string, newVals map[string]interface{}) error {
	return svc.DataComponent.UpdateMulti(ctx, ids, newVals)
}

//item must support Deleted field for soft deletes
func (svc *dataPlugin) Delete(ctx core.RequestContext, id string) error {
	return svc.DataComponent.Delete(ctx, id)
}

//Delete object by ids
func (svc *dataPlugin) DeleteMulti(ctx core.RequestContext, ids []string) error {
	return svc.DataComponent.DeleteMulti(ctx, ids)
}

//Delete object by condition
func (svc *dataPlugin) DeleteAll(ctx core.RequestContext, queryCond interface{}) ([]string, error) {
	return svc.DataComponent.DeleteAll(ctx, queryCond)
}

func (svc *dataPlugin) GetById(ctx core.RequestContext, id string) (data.Storable, error) {
	return svc.DataComponent.GetById(ctx, id)
}

//Get multiple objects by id
func (svc *dataPlugin) GetMulti(ctx core.RequestContext, ids []string, orderBy string) ([]data.Storable, error) {
	return svc.DataComponent.GetMulti(ctx, ids, orderBy)
}

func (svc *dataPlugin) GetMultiHash(ctx core.RequestContext, ids []string) (map[string]data.Storable, error) {
	return svc.DataComponent.GetMultiHash(ctx, ids)
}

func (svc *dataPlugin) Count(ctx core.RequestContext, queryCond interface{}) (count int, err error) {
	return svc.DataComponent.Count(ctx, queryCond)
}

func (svc *dataPlugin) CountGroups(ctx core.RequestContext, queryCond interface{}, group string) (res map[string]interface{}, err error) {
	return svc.DataComponent.CountGroups(ctx, queryCond, group)
}

func (svc *dataPlugin) GetList(ctx core.RequestContext, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	return svc.DataComponent.GetList(ctx, pageSize, pageNum, mode, orderBy)
}

func (svc *dataPlugin) Get(ctx core.RequestContext, queryCond interface{}, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	return svc.DataComponent.Get(ctx, queryCond, pageSize, pageNum, mode, orderBy)
}

//create condition for passing to data service
func (svc *dataPlugin) CreateCondition(ctx core.RequestContext, operation data.ConditionType, args ...interface{}) (interface{}, error) {
	return svc.DataComponent.CreateCondition(ctx, operation, args...)
}
*/
