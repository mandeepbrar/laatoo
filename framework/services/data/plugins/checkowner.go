package plugins

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	CHECK_READ = "CheckRead"
)

type checkOwnerService struct {
	*data.DataPlugin
	checkRead bool
}

func NewCheckOwnerService(ctx core.ServerContext) *checkOwnerService {
	return &checkOwnerService{DataPlugin: data.NewDataPlugin(ctx)}
}

func NewCheckOwnerServiceWithBase(ctx core.ServerContext, base data.DataComponent) *checkOwnerService {
	return &checkOwnerService{DataPlugin: data.NewDataPluginWithBase(ctx, base)}
}

func (svc *checkOwnerService) Initialize(ctx core.ServerContext, conf config.Config) error {
	err := svc.DataPlugin.Initialize(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	checkRead, ok := conf.GetBool(CHECK_READ)
	if ok {
		svc.checkRead = checkRead
	} else {
		svc.checkRead = false
	}
	return nil
}

func (svc *checkOwnerService) isOwned(ctx core.RequestContext, id string) (bool, error) {
	stor, err := svc.GetById(ctx, id)
	if err != nil {
		return false, err
	}
	if stor != nil {
		i, ok := stor.(data.Auditable)
		if ok {
			log.Logger.Trace(ctx, "checking owned", "created by", i.GetCreatedBy(), "user", ctx.GetUser().GetId())
			if i.GetCreatedBy() != ctx.GetUser().GetId() {
				ctx.SetResponse(core.StatusUnauthorizedResponse)
				return false, nil
			}
		}
	}
	return true, nil
}

func (svc *checkOwnerService) areOwned(ctx core.RequestContext, ids []string) (bool, error) {
	stors, err := svc.GetMulti(ctx, ids, "")
	if err != nil {
		return false, err
	}
	userId := ctx.GetUser().GetId()
	for _, item := range stors {
		i, ok := item.(data.Auditable)
		if ok && item.GetId() != "" {
			if i.GetCreatedBy() != userId {
				log.Logger.Info(ctx, "not owned", "item", item, "id", item.GetId(), "created by", i.GetCreatedBy(), "user id", userId)
				ctx.SetResponse(core.StatusUnauthorizedResponse)
				return false, nil
			}
		}
	}
	return true, nil
}

/*func (svc *dataCacheService) Save(ctx core.RequestContext, item data.Storable) error {
	return svc.DataComponent.Save(ctx, item)
}*/

func (svc *checkOwnerService) PutMulti(ctx core.RequestContext, items []data.Storable) error {
	ids := make([]string, len(items))
	for ind, item := range items {
		ids[ind] = item.GetId()
	}
	owned, err := svc.areOwned(ctx, ids)
	if owned {
		return svc.PluginDataComponent.PutMulti(ctx, items)
	}
	return err
}

func (svc *checkOwnerService) Put(ctx core.RequestContext, id string, item data.Storable) error {
	owned, err := svc.isOwned(ctx, id)
	if owned {
		return svc.PluginDataComponent.Put(ctx, id, item)
	}
	return err
}

//upsert an object ...insert if not there... update if there
func (svc *checkOwnerService) UpsertId(ctx core.RequestContext, id string, newVals map[string]interface{}) error {
	owned, err := svc.isOwned(ctx, id)
	if owned {
		return svc.PluginDataComponent.UpsertId(ctx, id, newVals)
	}
	return err
}

func (svc *checkOwnerService) Update(ctx core.RequestContext, id string, newVals map[string]interface{}) error {
	owned, err := svc.isOwned(ctx, id)
	if owned {
		return svc.PluginDataComponent.Update(ctx, id, newVals)
	}
	return err
}

func (svc *checkOwnerService) Upsert(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	return nil, errors.NotImplemented(ctx, "Upsert")
}

func (svc *checkOwnerService) UpdateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	return nil, errors.NotImplemented(ctx, "UpdateAll")
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (svc *checkOwnerService) UpdateMulti(ctx core.RequestContext, ids []string, newVals map[string]interface{}) error {
	owned, err := svc.areOwned(ctx, ids)
	if owned {
		return svc.PluginDataComponent.UpdateMulti(ctx, ids, newVals)
	}
	return err
}

//item must support Deleted field for soft deletes
func (svc *checkOwnerService) Delete(ctx core.RequestContext, id string) error {
	owned, err := svc.isOwned(ctx, id)
	log.Logger.Trace(ctx, "checking owned", "owned", owned)
	if owned {
		return svc.PluginDataComponent.Delete(ctx, id)
	}
	return err
}

//Delete object by ids
func (svc *checkOwnerService) DeleteMulti(ctx core.RequestContext, ids []string) error {
	owned, err := svc.areOwned(ctx, ids)
	if owned {
		return svc.PluginDataComponent.DeleteMulti(ctx, ids)
	}
	return err
}

//Delete object by condition
func (svc *checkOwnerService) DeleteAll(ctx core.RequestContext, queryCond interface{}) ([]string, error) {
	return nil, errors.NotImplemented(ctx, "DeleteAll")
}
