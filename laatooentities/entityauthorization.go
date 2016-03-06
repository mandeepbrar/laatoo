package laatooentities

import (
	"laatoocore"
	"laatoosdk/core"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
	"reflect"
)

const (
	CHECK_ENTITY_OWNERSHIP_AUTH_METHOD = "EnforceEntityOwnership"
)

func init() {
	laatoocore.RegisterInvokableMethod(CHECK_ENTITY_OWNERSHIP_AUTH_METHOD, EnforceEntityOwnership)
}

func EnforceEntityOwnership(ctx core.Context, conf map[string]interface{}) error {
	id := ctx.ParamByIndex(0)
	if id == "" {
		return errors.ThrowError(ctx, ENTITY_ERROR_NOT_FOUND)
	}
	entsvc := conf["entitysvc"]
	if entsvc == nil {
		return errors.ThrowError(ctx, ENTITY_ERROR_INCORRECT_METHOD_CONF)
	}
	svc, err := ctx.GetService(entsvc.(string))
	if err != nil {
		return errors.ThrowError(ctx, ENTITY_ERROR_INCORRECT_METHOD_CONF)
	}
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Got Service", "entity", id, "entsvc", entsvc, "svc", svc, "type", reflect.TypeOf(svc))
	entSvc := svc.(*EntityService)
	if entSvc == nil {
		return errors.ThrowError(ctx, ENTITY_ERROR_INCORRECT_METHOD_CONF)
	}
	ent, err := entSvc.getEntity(ctx, id)
	if err != nil {
		return errors.ThrowError(ctx, ENTITY_ERROR_NOT_FOUND)
	}
	auditable := ent.(data.Auditable)
	owner := auditable.GetCreatedBy()
	usr := ctx.GetUser()
	if usr == nil {
		log.Logger.Trace(ctx, LOGGING_CONTEXT, "Entity not accessible by anonymous user", "entity", id)
		return errors.ThrowError(ctx, ENTITY_ERROR_NOT_ALLOWED)
	}
	if owner != usr.GetId() {
		log.Logger.Trace(ctx, LOGGING_CONTEXT, "Entity accessible only by owner", "entity", id, "user", usr.GetId(), "owner", owner)
		return errors.ThrowError(ctx, ENTITY_ERROR_NOT_ALLOWED)
	}
	return nil
}
