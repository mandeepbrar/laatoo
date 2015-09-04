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
	EntityName      string
	dataServiceName string
	DataStore       data.DataService
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_ENTITY_SERVICENAME, EntityServiceFactory)
}

//factory method returns the service object to the environment
func EntityServiceFactory(conf map[string]interface{}) (interface{}, error) {
	log.Logger.Infof("Creating entity service")
	svc := &EntityService{}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	router := routerInt.(*echo.Group)

	//get a map of all the entities
	entityInt, ok := conf[CONF_ENTITY]
	if !ok {
		return nil, errors.ThrowError(ENTITY_ERROR_MISSING_ENTITY)
	}
	entityName := entityInt.(string)
	svc.EntityName = entityName

	entitydatasvcInt, ok := conf[CONF_ENTITY_DATA_SVC]
	if !ok {
		return nil, errors.ThrowError(ENTITY_ERROR_MISSING_DATASVC, svc.GetName())
	}
	svc.dataServiceName = entitydatasvcInt.(string)

	entitymethodsInt, ok := conf[CONF_ENTITY_METHODS]
	if !ok {
		return nil, errors.ThrowError(ENTITY_ERROR_MISSING_METHODS, svc.GetName())
	}

	entityMethods, ok := entitymethodsInt.(map[string]interface{})
	if !ok {
		return nil, errors.ThrowError(ENTITY_ERROR_MISSING_METHODS, svc.GetName())
	}

	viewperm := fmt.Sprintf("View %s", entityName)
	createperm := fmt.Sprintf("Create %s", entityName)
	editperm := fmt.Sprintf("Edit %s", entityName)
	deleteperm := fmt.Sprintf("Delete %s", entityName)
	laatoocore.RegisterPermissions([]string{viewperm, createperm, editperm, deleteperm})

	for name, val := range entityMethods {

		methodConfig, ok := val.(map[string]interface{})
		if !ok {
			return nil, errors.ThrowError(ENTITY_ERROR_INCORRECT_METHOD_CONF, entityName, name)
		}

		pathInt, ok := methodConfig[CONF_ENTITY_PATH]
		if !ok {
			return nil, errors.ThrowError(ENTITY_ERROR_MISSING_PATH, entityName, name)
		} else {
			methodInt, ok := methodConfig[CONF_ENTITY_METHOD]
			if !ok {
				return nil, errors.ThrowError(ENTITY_ERROR_MISSING_METHOD, entityName, name)
			}

			path := pathInt.(string)
			method := methodInt.(string)

			switch method {
			/*			case "list":
						router.Get(path, svc.ListArticle)*/
			case "get":
				router.Get(path, func(ctx *echo.Context) error {
					if !laatoocore.IsAllowed(ctx, viewperm) {
						return errors.ThrowHttpError(laatoocore.AUTH_ERROR_SECURITY, ctx)
					}
					id := ctx.P(0)
					log.Logger.Debugf("Getting entity %s", id)
					ent, err := svc.DataStore.GetById(entityName, id)
					if err != nil {
						return err
					}
					return ctx.JSON(http.StatusOK, ent)
				})
			case "post":
				router.Post(path, func(ctx *echo.Context) error {
					if !laatoocore.IsAllowed(ctx, createperm) {
						return errors.ThrowHttpError(laatoocore.AUTH_ERROR_SECURITY, ctx)
					}
					ent, err := laatoocore.CreateEmptyObject(entityName)
					if err != nil {
						return err
					}
					err = ctx.Bind(ent)
					if err != nil {
						return err
					}
					err = svc.DataStore.Save(entityName, ent)
					if err != nil {
						return err
					}
					return nil
				})
			case "put":
				router.Put(path, func(ctx *echo.Context) error {
					if !laatoocore.IsAllowed(ctx, editperm) {
						return errors.ThrowHttpError(laatoocore.AUTH_ERROR_SECURITY, ctx)
					}
					id := ctx.P(0)
					log.Logger.Debugf("Updating entity %s", id)
					ent, err := laatoocore.CreateEmptyObject(entityName)
					if err != nil {
						return err
					}
					err = ctx.Bind(ent)
					if err != nil {
						return err
					}
					err = svc.DataStore.Put(entityName, id, ent)
					if err != nil {
						return err
					}
					return nil
				})
			case "putbulk":
				router.Put(path, func(ctx *echo.Context) error {
					if !laatoocore.IsAllowed(ctx, editperm) {
						return errors.ThrowHttpError(laatoocore.AUTH_ERROR_SECURITY, ctx)
					}
					typ, err := laatoocore.GetCollectionType(entityName)
					if err != nil {
						return err
					}
					arrPtr := reflect.New(typ)
					log.Logger.Debugf("Binding entities with collection %s", entityName)
					err = ctx.Bind(arrPtr.Interface())
					if err != nil {
						return err
					}
					arr := arrPtr.Elem()
					length := arr.Len()
					log.Logger.Debugf("Saving bulk entities %s", entityName)
					for i := 0; i < length; i++ {
						entity := arr.Index(i).Addr().Interface().(data.Storable)
						err = svc.DataStore.Put(entityName, entity.GetId(), entity)
						if err != nil {
							return err
						}
					}
					return nil
				})
			case "delete":
				router.Delete(path, func(ctx *echo.Context) error {
					if !laatoocore.IsAllowed(ctx, deleteperm) {
						return errors.ThrowHttpError(laatoocore.AUTH_ERROR_SECURITY, ctx)
					}
					id := ctx.P(0)
					log.Logger.Debugf("Deleting entity %s", id)
					err := svc.DataStore.Delete(entityName, id)
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
	dataSvc, err := ctx.GetService(svc.dataServiceName)
	if err != nil {
		return errors.RethrowError(ENTITY_ERROR_MISSING_DATASVC, err, svc.EntityName)
	}

	svc.DataStore = dataSvc.(data.DataService)
	return nil
}

//The service starts serving when this method is called
func (svc *EntityService) Serve() error {
	return nil
}

//Type of service
func (svc *EntityService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

//Execute method
func (svc *EntityService) Execute(name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}
