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
	serviceContext  service.ServiceContext
	EntityName      string
	dataServiceName string
	DataStore       data.DataService
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_ENTITY_SERVICENAME, EntityServiceFactory)
}

//factory method returns the service object to the environment
func EntityServiceFactory(ctx interface{}, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Creating entity service")
	serviceContext := ctx.(service.ServiceContext)
	svc := &EntityService{serviceContext: serviceContext}
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

	viewperm := fmt.Sprintf("View %s", entityName)
	createperm := fmt.Sprintf("Create %s", entityName)
	editperm := fmt.Sprintf("Edit %s", entityName)
	deleteperm := fmt.Sprintf("Delete %s", entityName)
	serviceContext.RegisterPermissions(ctx, []string{viewperm, createperm, editperm, deleteperm})

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
			/*			case "list":
						router.Get(path, svc.ListArticle)*/
			case "get":
				router.Get(path, func(ctx *echo.Context) error {
					if !serviceContext.IsAllowed(ctx, viewperm) {
						return errors.ThrowError(ctx, laatoocore.AUTH_ERROR_SECURITY)
					}
					id := ctx.P(0)
					ent, err := svc.DataStore.GetById(ctx, entityName, id)
					if err != nil {
						return err
					}
					return ctx.JSON(http.StatusOK, ent)
				})
			case "post":
				router.Post(path, func(ctx *echo.Context) error {
					if !serviceContext.IsAllowed(ctx, createperm) {
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
					log.Logger.Trace(ctx, LOGGING_CONTEXT, "Saved entity", "Entity", ent)
					return nil
				})
			case "put":
				router.Put(path, func(ctx *echo.Context) error {
					if !serviceContext.IsAllowed(ctx, editperm) {
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
					return nil
				})
			case "putbulk":
				router.Put(path, func(ctx *echo.Context) error {
					if !serviceContext.IsAllowed(ctx, editperm) {
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
					}
					return nil
				})
			case "delete":
				router.Delete(path, func(ctx *echo.Context) error {
					if !serviceContext.IsAllowed(ctx, deleteperm) {
						return errors.ThrowError(ctx, laatoocore.AUTH_ERROR_SECURITY)
					}
					id := ctx.P(0)
					log.Logger.Debug(ctx, LOGGING_CONTEXT, "Deleting entity", "ID", id)
					err := svc.DataStore.Delete(ctx, entityName, id)
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
func (svc *EntityService) Initialize(ctx service.ServiceContext) error {
	dataSvc, err := ctx.GetService(ctx, svc.dataServiceName)
	if err != nil {
		return errors.RethrowError(ctx, ENTITY_ERROR_MISSING_DATASVC, err)
	}

	svc.DataStore = dataSvc.(data.DataService)
	return nil
}

//The service starts serving when this method is called
func (svc *EntityService) Serve(ctx interface{}) error {
	return nil
}

//Type of service
func (svc *EntityService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

//Execute method
func (svc *EntityService) Execute(ctx interface{}, name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}
