package plugins

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/utils"
)

const (
	CACHE_BUCKET = "bucket"
)

type dataCacheService struct {
	*data.DataPlugin
	bucket string
}

func NewDataCacheService(ctx core.ServerContext) *dataCacheService {
	return &dataCacheService{DataPlugin: data.NewDataPlugin(ctx)}
}

func NewCacheServiceWithBase(ctx core.ServerContext, base data.DataComponent) *dataCacheService {
	return &dataCacheService{DataPlugin: data.NewDataPluginWithBase(ctx, base)}
}

func (svc *dataCacheService) Initialize(ctx core.ServerContext, conf config.Config) error {
	err := svc.DataPlugin.Initialize(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	bucket, ok := conf.GetString(CACHE_BUCKET)
	if !ok {
		svc.bucket = svc.Object
	} else {
		svc.bucket = bucket
	}
	return nil
}

/*func (svc *dataCacheService) Save(ctx core.RequestContext, item data.Storable) error {
	return svc.DataComponent.Save(ctx, item)
}*/

func (svc *dataCacheService) PutMulti(ctx core.RequestContext, items []data.Storable) error {
	err := svc.PluginDataComponent.PutMulti(ctx, items)
	if err != nil {
		for _, item := range items {
			id := item.GetId()
			ctx.InvalidateCache(svc.bucket, id)
		}
	}
	return err
}

func (svc *dataCacheService) Put(ctx core.RequestContext, id string, item data.Storable) error {
	err := svc.PluginDataComponent.Put(ctx, id, item)
	if err != nil {
		ctx.InvalidateCache(svc.bucket, id)
	}
	return err
}

//upsert an object ...insert if not there... update if there
func (svc *dataCacheService) UpsertId(ctx core.RequestContext, id string, newVals map[string]interface{}) error {
	err := svc.PluginDataComponent.UpsertId(ctx, id, newVals)
	if err != nil {
		ctx.InvalidateCache(svc.bucket, id)
	}
	return err
}

func (svc *dataCacheService) Update(ctx core.RequestContext, id string, newVals map[string]interface{}) error {
	err := svc.PluginDataComponent.Update(ctx, id, newVals)
	if err != nil {
		ctx.InvalidateCache(svc.bucket, id)
	}
	return err
}

func (svc *dataCacheService) Upsert(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	ids, err := svc.PluginDataComponent.Upsert(ctx, queryCond, newVals)
	if err != nil {
		for _, id := range ids {
			ctx.InvalidateCache(svc.bucket, id)
		}
	}
	return ids, err
}

func (svc *dataCacheService) UpdateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error) {
	ids, err := svc.PluginDataComponent.UpdateAll(ctx, queryCond, newVals)
	if err != nil {
		for _, id := range ids {
			ctx.InvalidateCache(svc.bucket, id)
		}
	}
	return ids, err
}

//update objects by ids, fields to be updated should be provided as key value pairs
func (svc *dataCacheService) UpdateMulti(ctx core.RequestContext, ids []string, newVals map[string]interface{}) error {
	err := svc.PluginDataComponent.UpdateMulti(ctx, ids, newVals)
	if err != nil {
		for _, id := range ids {
			ctx.InvalidateCache(svc.bucket, id)
		}
	}
	return err
}

//item must support Deleted field for soft deletes
func (svc *dataCacheService) Delete(ctx core.RequestContext, id string) error {
	err := svc.PluginDataComponent.Delete(ctx, id)
	if err != nil {
		ctx.InvalidateCache(svc.bucket, id)
	}
	return err
}

//Delete object by ids
func (svc *dataCacheService) DeleteMulti(ctx core.RequestContext, ids []string) error {
	err := svc.PluginDataComponent.DeleteMulti(ctx, ids)
	if err != nil {
		for _, id := range ids {
			ctx.InvalidateCache(svc.bucket, id)
		}
	}
	return err
}

//Delete object by condition
func (svc *dataCacheService) DeleteAll(ctx core.RequestContext, queryCond interface{}) ([]string, error) {
	ids, err := svc.PluginDataComponent.DeleteAll(ctx, queryCond)
	if err != nil {
		for _, id := range ids {
			ctx.InvalidateCache(svc.bucket, id)
		}
	}
	return ids, err
}

func (svc *dataCacheService) GetById(ctx core.RequestContext, id string) (data.Storable, error) {
	ctx = ctx.SubContext("Cache_GetById")

	ent, ok := ctx.GetObjectFromCache(svc.bucket, id, svc.Object)
	if ok {
		return ent.(data.Storable), nil
	}
	stor, err := svc.PluginDataComponent.GetById(ctx, id)
	if err == nil {
		ctx.PutInCache(svc.bucket, id, stor)
	}
	return stor, err
}

//Get multiple objects by id
func (svc *dataCacheService) GetMulti(ctx core.RequestContext, ids []string, orderBy string) ([]data.Storable, error) {
	ctx = ctx.SubContext("Cache_GetMulti")
	res := make([]data.Storable, len(ids))
	cachedItems := ctx.GetObjectsFromCache(svc.bucket, ids, svc.Object)
	idsNotCached := make([]string, 0, 10)
	for index, id := range ids {
		item, ok := cachedItems[id]
		if !ok || item == nil {
			idsNotCached = append(idsNotCached, id)
		} else {
			res[index] = item.(data.Storable)
		}
	}
	if len(idsNotCached) == 0 {
		return res, nil
	}
	stormap, err := svc.PluginDataComponent.GetMultiHash(ctx, idsNotCached)
	if err == nil {
		for index, id := range ids {
			if res[index] == nil {
				item, ok := stormap[id]
				if ok {
					ctx.PutInCache(svc.bucket, id, item)
					res[index] = item
				}
			}
		}
	}
	return res, err
}

func (svc *dataCacheService) GetMultiHash(ctx core.RequestContext, ids []string) (map[string]data.Storable, error) {
	ctx = ctx.SubContext("Cache_GetMultiHash")
	res, err := svc.PluginDataComponent.GetMultiHash(ctx, ids)
	if err == nil {
		ctx.PutMultiInCache(svc.bucket, utils.CastToStringMap(res))
	}
	return res, err
}

func (svc *dataCacheService) Count(ctx core.RequestContext, queryCond interface{}) (count int, err error) {
	return svc.PluginDataComponent.Count(ctx, queryCond)
}

func (svc *dataCacheService) CountGroups(ctx core.RequestContext, queryCond interface{}, groupids []string, group string) (res map[string]interface{}, err error) {
	return svc.PluginDataComponent.CountGroups(ctx, queryCond, groupids, group)
}

func (svc *dataCacheService) GetList(ctx core.RequestContext, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	return svc.PluginDataComponent.GetList(ctx, pageSize, pageNum, mode, orderBy)
}

func (svc *dataCacheService) Get(ctx core.RequestContext, queryCond interface{}, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []data.Storable, ids []string, totalrecs int, recsreturned int, err error) {
	return svc.PluginDataComponent.Get(ctx, queryCond, pageSize, pageNum, mode, orderBy)
}
