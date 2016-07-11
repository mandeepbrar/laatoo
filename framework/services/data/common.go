package data

import (
	"laatoo/sdk/components"
	"laatoo/sdk/core"
)

const (
	CONF_DATA_SVCS          = "dataservices"
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
	CONF_PRESAVE_MSG        = "storable_presave"
	CONF_PREUPDATE_MSG      = "storable_preupdate"
)

func notifyUpdate(ctx core.RequestContext, objectType string, id string) {
	invalidateCache(ctx, objectType, id)
}

func notifyDelete(ctx core.RequestContext, objectType string, id string) {

}

func getFromCache(ctx core.RequestContext, objectType string, id string, object interface{}) bool {
	cachekey := components.GetCacheKey(objectType, id)
	return ctx.GetFromCache(cachekey, object)
}

func putInCache(ctx core.RequestContext, objectType string, id string, object interface{}) {
	ctx.PutInCache(components.GetCacheKey(objectType, id), object)
}

func invalidateCache(ctx core.RequestContext, objectType string, id string) {
	ctx.InvalidateCache(components.GetCacheKey(objectType, id))
}
