package laatooentities

import (
	"fmt"
	"laatoocore"
	"laatoosdk/core"
	"laatoosdk/data"
	"laatoosdk/entities"
	"laatoosdk/errors"
	"laatoosdk/log"
	"net/http"
	"reflect"
	"strings"
)

const (
	LOGGING_CONTEXT         = "entity_service"
	CONF_ENTITY_SERVICENAME = "entity_service"
	CONF_ENTITY             = "entity"
	CONF_ENTITY_DATA_SVC    = "data_svc"
	CONF_ENTITY_TYPE        = "type"
	CONF_ENTITY_ID          = "id"
	CONF_ENTITY_METHODS     = "methods"
	CONF_ENTITY_METHOD      = "method"
	CONF_ENTITY_PATH        = "path"
	CONF_ENTITY_INVOKE      = "invoke"
)

//Environment hosting an application
type EntityService struct {
	EntityName      string
	dataServiceName string
	DataStore       data.DataService
	cache           bool
	updateOwnership bool
	viewOwnership   bool
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_ENTITY_SERVICENAME, EntityServiceFactory)
}

//factory method returns the service object to the environment
func EntityServiceFactory(ctx core.Context, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Creating entity service")
	svc := &EntityService{cache: true}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	router := routerInt.(core.Router)

	//get a map of all the entities
	entityInt, ok := conf[CONF_ENTITY]
	if !ok {
		return nil, errors.ThrowError(ctx, ENTITY_ERROR_MISSING_ENTITY)
	}
	entityName := entityInt.(string)
	svc.EntityName = entityName

	entitydatasvcInt, ok := conf[CONF_ENTITY_DATA_SVC]
	if !ok {
		return nil, errors.ThrowError(ctx, ENTITY_ERROR_MISSING_DATASVC)
	}
	svc.dataServiceName = entitydatasvcInt.(string)

	entitymethodsInt, ok := conf[CONF_ENTITY_METHODS]
	if !ok {
		return nil, errors.ThrowError(ctx, ENTITY_ERROR_MISSING_METHODS)
	}

	entityMethods, ok := entitymethodsInt.(map[string]interface{})
	if !ok {
		return nil, errors.ThrowError(ctx, ENTITY_ERROR_MISSING_METHODS)
	}

	for name, val := range entityMethods {

		methodConfig, ok := val.(map[string]interface{})
		if !ok {
			return nil, errors.ThrowError(ctx, ENTITY_ERROR_INCORRECT_METHOD_CONF, "Entity", entityName, "Method", name)
		}

		pathInt, ok := methodConfig[CONF_ENTITY_PATH]
		if !ok {
			return nil, errors.ThrowError(ctx, ENTITY_ERROR_MISSING_PATH, "Entity", entityName, "Method", name)
		} else {
			methodInt, ok := methodConfig[CONF_ENTITY_METHOD]
			if !ok {
				return nil, errors.ThrowError(ctx, ENTITY_ERROR_MISSING_METHOD, "Entity", entityName, "Method", name)
			}

			path := pathInt.(string)
			method := methodInt.(string)

			switch method {
			case "get":
				router.Get(ctx, path, methodConfig, func(ctx core.Context) error {
					id := ctx.ParamByIndex(0)
					ent, err := svc.getEntity(ctx, id)
					if err != nil {
						return err
					}
					return ctx.JSON(http.StatusOK, ent)
				})
			case "getbulk":
				router.Get(ctx, path, methodConfig, func(ctx core.Context) error {
					idsstr := ctx.ParamByIndex(0)
					ids := strings.Split(idsstr, ",")
					orderBy := ctx.Param(VIEW_ORDERBY)
					ents, err := svc.getEntities(ctx, ids, orderBy)
					if err != nil {
						return err
					}
					return ctx.JSON(http.StatusOK, ents)
				})
			case "post":
				router.Post(ctx, path, methodConfig, func(ctx core.Context) error {
					ent, err := laatoocore.CreateEmptyObject(ctx, entityName)
					if err != nil {
						return err
					}
					err = ctx.Bind(ent)
					if err != nil {
						return err
					}
					data.Audit(ctx, ent)
					err = svc.DataStore.Save(ctx, entityName, ent)
					if err != nil {
						return err
					}
					if svc.cache {
						stor := ent.(data.Storable)
						svc.invalidateCache(ctx, stor.GetId())
					}
					log.Logger.Trace(ctx, LOGGING_CONTEXT, "Saved entity")
					return nil
				})
			case "put":
				router.Put(ctx, path, methodConfig, func(ctx core.Context) error {
					ent, err := laatoocore.CreateEmptyObject(ctx, svc.EntityName)
					if err != nil {
						return err
					}
					err = ctx.Bind(ent)
					if err != nil {
						return err
					}
					stor := ent.(data.Storable)
					_, err = svc.putEntity(ctx, stor.GetId(), ent)
					return err
				})
			case "putbulk":
				router.Put(ctx, path, methodConfig, func(ctx core.Context) error {
					arr, err := laatoocore.CreateCollection(ctx, entityName)
					if err != nil {
						return err
					}
					log.Logger.Trace(ctx, LOGGING_CONTEXT, "Binding entities with collection", "Entity", entityName)
					err = ctx.Bind(arr)
					if err != nil {
						return err
					}
					log.Logger.Trace(ctx, LOGGING_CONTEXT, "Binding done")
					_, err = svc.putBulkEntity(ctx, reflect.ValueOf(arr).Elem().Interface())
					return err
				})
			case "delete":
				router.Delete(ctx, path, methodConfig, func(ctx core.Context) error {
					id := ctx.ParamByIndex(0)
					log.Logger.Debug(ctx, LOGGING_CONTEXT, "Deleting entity", "ID", id)
					err := svc.DataStore.Delete(ctx, entityName, id)
					if err != nil {
						return err
					}
					if svc.cache {
						svc.invalidateCache(ctx, id)
					}
					return nil
				})
			case "update":
				router.Post(ctx, path, methodConfig, func(ctx core.Context) error {
					id := ctx.ParamByIndex(0)
					log.Logger.Trace(ctx, LOGGING_CONTEXT, "Updating entity", "ID", id)
					vals := make(map[string]interface{}, 10)
					err := ctx.Bind(&vals)
					if err != nil {
						return err
					}
					_, err = svc.updateEntity(ctx, id, vals)
					if err != nil {
						return err
					}
					if svc.cache {
						svc.invalidateCache(ctx, id)
					}
					return nil
				})
			case "invokeonEntity":
				//incomplete
				router.Put(ctx, path, methodConfig, func(ctx core.Context) error {
					invokeReqInt, ok := methodConfig[CONF_ENTITY_INVOKE]
					if !ok {
						return errors.ThrowError(ctx, ENTITY_ERROR_MISSING_INV_METHOD, "Entity", entityName, "Invocation request", name)
					}
					id := ctx.ParamByIndex(0)
					obj, err := svc.getEntity(ctx, id)
					if err != nil {
						return err
					}
					ent, _ := obj.(entities.Entity)
					err = ent.Invoke(ctx, invokeReqInt.(string))
					if err != nil {
						return err
					}
					return nil
				})
			}
		}
	}
	return svc, nil
}

//Provides the name of the service
func (svc *EntityService) GetName() string {
	return CONF_ENTITY_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *EntityService) Initialize(ctx core.Context) error {
	dataSvc, err := ctx.GetService(svc.dataServiceName)
	if err != nil {
		return errors.RethrowError(ctx, ENTITY_ERROR_MISSING_DATASVC, err)
	}

	svc.DataStore = dataSvc.(data.DataService)
	return nil
}

func (svc *EntityService) invalidateCache(ctx core.Context, id string) {
	if svc.cache {
		ctx.DeleteFromCache(svc.getCacheKey(id))
	}
}

func (svc *EntityService) getCacheKey(id string) string {
	return fmt.Sprintf("Entity_%s_%s", svc.EntityName, id)
}

//The service starts serving when this method is called
func (svc *EntityService) Serve(ctx core.Context) error {
	return nil
}

//Type of service
func (svc *EntityService) GetServiceType() string {
	return core.SERVICE_TYPE_WEB
}

//Execute method
func (svc *EntityService) Execute(ctx core.Context, name string, params map[string]interface{}) (interface{}, error) {
	switch name {
	case "Get":
		return svc.getEntity(ctx, params["id"].(string))
	case "Put":
		return svc.putEntity(ctx, params["id"].(string), params["entity"])
	case "Update":
		return svc.updateEntity(ctx, params["id"].(string), params["data"].(map[string]interface{}))
	case "PutBulk":
		return svc.putBulkEntity(ctx, params["entities"])
	case "GetBulk":
		orderBy := ""
		orderByInt, ok := params[VIEW_ORDERBY]
		if ok {
			orderBy = orderByInt.(string)
		}
		return svc.getEntities(ctx, params["ids"].([]string), orderBy)
	case "Select":
		view, err := newEntitiesView(ctx, map[string]interface{}{"entity": svc.EntityName})
		if err != nil {
			return nil, err
		}
		orderBy := ""
		orderByInt, ok := params[VIEW_ORDERBY]
		if ok {
			orderBy = orderByInt.(string)
		}
		entities, _, _, err := view.getData(ctx, svc.DataStore, params, -1, -1, orderBy)
		return entities, err
	}
	return nil, nil
}

func (svc *EntityService) putEntity(ctx core.Context, id string, ent interface{}) (interface{}, error) {
	data.Audit(ctx, ent)
	err := svc.DataStore.Put(ctx, svc.EntityName, id, ent)
	if err != nil {
		return nil, err
	}
	if svc.cache {
		stor := ent.(data.Storable)
		svc.invalidateCache(ctx, stor.GetId())
	}
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Saved entity")
	return nil, nil
}

func (svc *EntityService) putBulkEntity(ctx core.Context, arrInt interface{}) (interface{}, error) {
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Saving bulk entities")
	arr := reflect.ValueOf(arrInt)
	length := arr.Len()
	ids := make([]string, length)
	for i := 0; i < length; i++ {
		entinter := arr.Index(i).Addr().Interface()
		entity := entinter.(data.Storable)
		ids[i] = entity.GetId()
		data.Audit(ctx, entinter)
	}
	err := svc.DataStore.PutMulti(ctx, svc.EntityName, ids, arrInt)
	for i := 0; i < length; i++ {
		if svc.cache {
			svc.invalidateCache(ctx, ids[i])
		}
	}
	return nil, err
}

func (svc *EntityService) getEntity(ctx core.Context, id string) (interface{}, error) {
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Getting Entity", "entity", svc.EntityName, "id", id)
	if svc.cache {
		ent, err := laatoocore.CreateEmptyObject(ctx, svc.EntityName)
		if err != nil {
			return nil, err
		}
		err = ctx.GetFromCache(svc.getCacheKey(id), ent)
		if err == nil {
			return ent, nil
		}
	}
	ent, err := svc.DataStore.GetById(ctx, svc.EntityName, id)
	if err != nil {
		return nil, err
	} else {
		if svc.cache {
			ctx.PutInCache(svc.getCacheKey(id), ent)
		}
	}
	return ent, nil
}

func (svc *EntityService) updateEntity(ctx core.Context, id string, newVals map[string]interface{}) (interface{}, error) {
	ent, err := svc.getEntity(ctx, id)
	if err != nil {
		return nil, err
	}
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Updating Entity", "entity", svc.EntityName, "id", id)
	entVal := reflect.ValueOf(ent).Elem()
	for k := range newVals {
		v := newVals[k]
		f := entVal.FieldByName(k)
		if f.IsValid() {
			// A Value can be changed only if it is
			// addressable and was not obtained by
			// the use of unexported struct fields.
			if f.CanSet() {
				// change value of N
				if f.Kind() == reflect.Int {
					f.SetInt(v.(int64))
				}
				// change value of N
				if f.Kind() == reflect.String {
					f.SetString(v.(string))
				}
				// change value of N
				if f.Kind() == reflect.Bool {
					f.SetBool(v.(bool))
				}
			}
		}
	}
	data.Audit(ctx, ent)
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Updating Entity", "ent", ent, "id", id)
	_, err = svc.putEntity(ctx, id, ent)
	if err != nil {
		return nil, err
	}
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "updated entity", "id", id)
	return nil, nil
}

func (svc *EntityService) getEntities(ctx core.Context, ids []string, orderBy string) (map[string]interface{}, error) {
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Getting multiple entities", "entity", svc.EntityName)
	ents, err := svc.DataStore.GetMulti(ctx, svc.EntityName, ids, orderBy)
	if err != nil {
		return nil, err
	} else {
		for k, v := range ents {
			if v != nil && svc.cache {
				ctx.PutInCache(svc.getCacheKey(k), v)
			}
		}
	}
	return ents, nil
}
