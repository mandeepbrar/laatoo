package laatooentities

import (
	"fmt"
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"net/http"
	"reflect"
	"strings"
)

const (
	LOGGING_CONTEXT         = "entity_service"
	CONF_ENTITY_SERVICENAME = "entity_service"
	CONF_ENTITY             = "entity"

	CONF_ENTITY_DATA_SVC = "data_svc"
	CONF_ENTITY_TYPE     = "type"
	CONF_ENTITY_ID       = "id"
	CONF_ENTITY_METHODS  = "methods"
	CONF_ENTITY_METHOD   = "method"
	CONF_ENTITY_PATH     = "path"
)

//Environment hosting an application
type EntityService struct {
	serviceEnv      service.Environment
	EntityName      string
	dataServiceName string
	DataStore       data.DataService
	cache           data.Cache
	viewperm        string
	createperm      string
	editperm        string
	deleteperm      string
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_ENTITY_SERVICENAME, EntityServiceFactory)
}

//factory method returns the service object to the environment
func EntityServiceFactory(ctx *echo.Context, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Creating entity service")
	serviceEnv := ctx.Get(laatoocore.CONF_ENV_CONTEXT).(service.Environment)
	svc := &EntityService{serviceEnv: serviceEnv}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	router := routerInt.(*echo.Group)

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

	svc.viewperm = fmt.Sprintf("View %s", entityName)
	svc.createperm = fmt.Sprintf("Create %s", entityName)
	svc.editperm = fmt.Sprintf("Edit %s", entityName)
	svc.deleteperm = fmt.Sprintf("Delete %s", entityName)
	//serviceContext.RegisterPermissions(ctx, []string{viewperm, createperm, editperm, deleteperm})

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
				router.Get(path, func(ctx *echo.Context) error {
					id := ctx.P(0)
					ent, err := svc.getEntity(ctx, id)
					if err != nil {
						return err
					}
					return ctx.JSON(http.StatusOK, ent)
				})
			case "getbulk":
				router.Get(path, func(ctx *echo.Context) error {
					idsstr := ctx.P(0)
					ids := strings.Split(idsstr, ",")
					ents, err := svc.getEntities(ctx, ids)
					if err != nil {
						return err
					}
					return ctx.JSON(http.StatusOK, ents)
				})
			case "post":
				router.Post(path, func(ctx *echo.Context) error {
					if !serviceEnv.IsAllowed(ctx, svc.createperm) {
						return errors.ThrowError(ctx, laatoocore.AUTH_ERROR_SECURITY)
					}
					ent, err := laatoocore.CreateEmptyObject(ctx, entityName)
					if err != nil {
						return err
					}
					err = ctx.Bind(ent)
					if err != nil {
						return err
					}
					err = svc.DataStore.Save(ctx, entityName, ent)
					if err != nil {
						return err
					}
					if svc.cache != nil {
						stor := ent.(data.Storable)
						svc.invalidateCache(ctx, stor.GetId())
					}
					log.Logger.Trace(ctx, LOGGING_CONTEXT, "Saved entity", "Entity", ent)
					return nil
				})
			case "put":
				router.Put(path, func(ctx *echo.Context) error {
					if !serviceEnv.IsAllowed(ctx, svc.editperm) {
						return errors.ThrowError(ctx, laatoocore.AUTH_ERROR_SECURITY)
					}
					id := ctx.P(0)
					log.Logger.Trace(ctx, LOGGING_CONTEXT, "Updating entity", "ID", id)
					ent, err := laatoocore.CreateEmptyObject(ctx, entityName)
					if err != nil {
						return err
					}
					err = ctx.Bind(ent)
					if err != nil {
						return err
					}
					err = svc.DataStore.Put(ctx, entityName, id, ent)
					if err != nil {
						return err
					}
					if svc.cache != nil {
						svc.invalidateCache(ctx, id)
					}
					return nil
				})
			case "putbulk":
				router.Put(path, func(ctx *echo.Context) error {
					if !serviceEnv.IsAllowed(ctx, svc.editperm) {
						return errors.ThrowError(ctx, laatoocore.AUTH_ERROR_SECURITY)
					}
					typ, err := laatoocore.GetCollectionType(ctx, entityName)
					if err != nil {
						return err
					}
					arrPtr := reflect.New(typ)
					log.Logger.Trace(ctx, LOGGING_CONTEXT, "Binding entities with collection", "Entity", entityName)
					err = ctx.Bind(arrPtr.Interface())
					if err != nil {
						return err
					}
					arr := arrPtr.Elem()
					length := arr.Len()
					log.Logger.Trace(ctx, "Saving bulk entities", "Entity", entityName)
					for i := 0; i < length; i++ {
						entity := arr.Index(i).Addr().Interface().(data.Storable)
						err = svc.DataStore.Put(ctx, entityName, entity.GetId(), entity)
						if err != nil {
							return err
						}
						if svc.cache != nil {
							svc.invalidateCache(ctx, entity.GetId())
						}
					}
					return nil
				})
			case "delete":
				router.Delete(path, func(ctx *echo.Context) error {
					if !serviceEnv.IsAllowed(ctx, svc.deleteperm) {
						return errors.ThrowError(ctx, laatoocore.AUTH_ERROR_SECURITY)
					}
					id := ctx.P(0)
					log.Logger.Debug(ctx, LOGGING_CONTEXT, "Deleting entity", "ID", id)
					err := svc.DataStore.Delete(ctx, entityName, id)
					if err != nil {
						return err
					}
					if svc.cache != nil {
						svc.invalidateCache(ctx, id)
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
func (svc *EntityService) Initialize(ctx *echo.Context) error {
	svcenv := ctx.Get(laatoocore.CONF_ENV_CONTEXT).(service.Environment)
	dataSvc, err := svcenv.GetService(ctx, svc.dataServiceName)
	if err != nil {
		return errors.RethrowError(ctx, ENTITY_ERROR_MISSING_DATASVC, err)
	}

	svc.DataStore = dataSvc.(data.DataService)
	svc.cache = svcenv.GetCache()
	return nil
}

func (svc *EntityService) invalidateCache(ctx *echo.Context, id string) {
	if svc.cache != nil {
		svc.cache.Delete(ctx, svc.getCacheKey(id))
	}
}

func (svc *EntityService) getCacheKey(id string) string {
	return fmt.Sprintf("Entity_%s_%s", svc.EntityName, id)
}

//The service starts serving when this method is called
func (svc *EntityService) Serve(ctx *echo.Context) error {
	return nil
}

//Type of service
func (svc *EntityService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

//Execute method
func (svc *EntityService) Execute(ctx *echo.Context, name string, params map[string]interface{}) (interface{}, error) {
	switch name {
	case "Get":
		return svc.getEntity(ctx, params["id"].(string))
	case "GetBulk":
		return svc.getEntities(ctx, params["ids"].([]string))
	}
	return nil, nil
}

func (svc *EntityService) getEntity(ctx *echo.Context, id string) (interface{}, error) {
	if !svc.serviceEnv.IsAllowed(ctx, svc.viewperm) {
		return nil, errors.ThrowError(ctx, laatoocore.AUTH_ERROR_SECURITY)
	}
	if svc.cache != nil {
		ent, err := laatoocore.CreateEmptyObject(ctx, svc.EntityName)
		if err != nil {
			return nil, err
		}
		err = svc.cache.GetObject(ctx, svc.getCacheKey(id), ent)
		if err == nil {
			return ent, nil
		}
	}
	ent, err := svc.DataStore.GetById(ctx, svc.EntityName, id)
	if err != nil {
		return nil, err
	} else {
		if svc.cache != nil {
			svc.cache.PutObject(ctx, svc.getCacheKey(id), ent)
		}
	}
	return ent, nil
}

func (svc *EntityService) getEntities(ctx *echo.Context, ids []string) (map[string]interface{}, error) {
	if !svc.serviceEnv.IsAllowed(ctx, svc.viewperm) {
		return nil, errors.ThrowError(ctx, laatoocore.AUTH_ERROR_SECURITY)
	}
	ents, err := svc.DataStore.GetMulti(ctx, svc.EntityName, ids)
	if err != nil {
		return nil, err
	} else {
		for k, v := range ents {
			if v != nil {
				svc.cache.PutObject(ctx, svc.getCacheKey(k), v)
			}
		}
	}
	return ents, nil
}
