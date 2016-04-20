package data

import (
	"laatoo/sdk/core"
	"laatoo/sdk/services"
)

const (
	CONF_DATA_OBJECT        = "object"
	CONF_DATA_OBJECT_ID     = "id"
	CONF_DATA_POSTSAVE      = "postsave"
	CONF_DATA_POSTLOAD      = "postload"
	CONF_DATA_PRESAVE       = "presave"
	CONF_DATA_CACHEABLE     = "cacheable"
	CONF_DATA_AUDITABLE     = "auditable"
	CONF_DATA_SOFTDELETE    = "softdelete"
	CONF_DATA_COLLECTION    = "collection"
	CONF_DATA_NOTIFYUPDATES = "notifyupdates"
	CONF_DATA_NOTIFYNEW     = "notifynew"
)

func notifyUpdate(ctx core.RequestContext, objectType string, id string) {
	invalidateCache(ctx, objectType, id)
}

func notifyDelete(ctx core.RequestContext, objectType string, id string) {

}

func getFromCache(ctx core.RequestContext, objectType string, id string, object interface{}) bool {
	cachekey := services.GetCacheKey(objectType, id)
	return ctx.GetFromCache(cachekey, object)
}

func putInCache(ctx core.RequestContext, objectType string, id string, object interface{}) {
	ctx.PutInCache(services.GetCacheKey(objectType, id), object)
}

func invalidateCache(ctx core.RequestContext, objectType string, id string) {
	ctx.InvalidateCache(services.GetCacheKey(objectType, id))
}
