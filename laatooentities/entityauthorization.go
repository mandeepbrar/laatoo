package laatooentities

import (
	"laatoocore"
	"laatoosdk/core"
	"laatoosdk/errors"
	"laatoosdk/log"
	"reflect"
)

const (
	CHECK_ENTITY_OWNERSHIP_AUTH_METHOD = "EnforceEntityOwnership"
	AUTHORIZED_ENTITY                  = "AuthorizedEntity"
)

func init() {
	laatoocore.RegisterInvokableMethod(CHECK_ENTITY_OWNERSHIP_AUTH_METHOD, EnforceEntityOwnership)
}

func EnforceEntityOwnership(ctx core.Context, conf map[string]interface{}) error {
	id := ctx.ParamByIndex(0)
	if id == "" {
		return errors.ThrowError(ctx, ENTITY_ERROR_NOT_FOUND)
	}
	usr := ctx.GetUser()
	if usr == nil {
		log.Logger.Trace(ctx, LOGGING_CONTEXT, "Entity not accessible by anonymous user", "entity", id)
		return errors.ThrowError(ctx, ENTITY_ERROR_NOT_ALLOWED)
	}
	field := conf["ownerfield"]
	if field == nil {
		return errors.ThrowError(ctx, ENTITY_ERROR_INCORRECT_METHOD_CONF)
	}
	entsvc := conf["entitysvc"]
	if entsvc == nil {
		return errors.ThrowError(ctx, ENTITY_ERROR_INCORRECT_METHOD_CONF)
	}
	svc, err := ctx.GetService(entsvc.(string))
	if err != nil {
		return errors.ThrowError(ctx, ENTITY_ERROR_INCORRECT_METHOD_CONF)
	}
	entSvc := svc.(*EntityService)
	if entSvc == nil {
		return errors.ThrowError(ctx, ENTITY_ERROR_INCORRECT_METHOD_CONF)
	}
	ent, err := entSvc.getEntity(ctx, id)
	if err != nil {
		return errors.ThrowError(ctx, ENTITY_ERROR_NOT_FOUND)
	}
	entVal := reflect.ValueOf(ent).Elem()
	f := entVal.FieldByName(field.(string))
	ownerVal := f.Interface().(string)
	if ownerVal != usr.GetId() {
		log.Logger.Trace(ctx, LOGGING_CONTEXT, "Entity accessible only by owner", "entity", id, "user", usr.GetId(), "ownerVal", ownerVal)
		return errors.ThrowError(ctx, ENTITY_ERROR_NOT_ALLOWED)
	}
	ctx.Set(AUTHORIZED_ENTITY, ent)
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "accessible owner", "entity", id, "user", usr.GetId(), "ownerVal", ownerVal)
	return nil
}
