package laatooauthentication

import (
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/data"
	"laatoosdk/entities"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"net/http"
)

const (
	ENTITY_ROLE_SERVICE_NAME = "role_service"
	PERM_ROLE_VIEW           = "View Role"
	PERM_ROLE_EDIT           = "Edit Role"
	PERM_ROLE_CREATE         = "Create Role"
	PERM_ROLE_DEL            = "Delete Role"
)

type RoleService struct {
	DataStore   data.DataService
	Router      *echo.Group
	dataSvcName string
	RoleObject  string
}

func init() {
	laatoocore.RegisterObjectProvider(ENTITY_ROLE_SERVICE_NAME, NewRoleService)
	laatoocore.RegisterPermissions([]string{PERM_ROLE_VIEW, PERM_ROLE_EDIT, PERM_ROLE_CREATE, PERM_ROLE_DEL})
}

func NewRoleService(conf map[string]interface{}) (interface{}, error) {

	log.Logger.Infof("Creating entity service", ENTITY_ROLE_SERVICE_NAME)

	svc := &RoleService{}

	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	router := routerInt.(*echo.Group)

	entitydatasvcInt, ok := conf[CONF_AUTHSERVICE_USERDATASERVICE]
	if !ok {
		return nil, errors.ThrowError(AUTH_ERROR_MISSING_USER_DATA_SERVICE, svc.GetName())
	}

	router.Get("/:id", func(ctx *echo.Context) error {
		id := ctx.P(0)
		log.Logger.Debugf("Getting entity %s", id)
		ent, err := svc.GetDataStore().GetById(svc.RoleObject, id)
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, ent)
	})
	router.Post("", func(ctx *echo.Context) error {
		ent, err := laatoocore.CreateEmptyObject(svc.RoleObject)
		if err != nil {
			return err
		}
		err = ctx.Bind(ent)
		if err != nil {
			return err
		}
		err = svc.GetDataStore().Save(svc.RoleObject, ent)
		if err != nil {
			return err
		}
		return nil
	})
	router.Put("/:id", func(ctx *echo.Context) error {
		id := ctx.P(0)
		log.Logger.Debugf("Updating entity %s", id)
		ent, err := laatoocore.CreateEmptyObject(svc.RoleObject)
		if err != nil {
			return err
		}
		err = ctx.Bind(ent)
		if err != nil {
			return err
		}
		err = svc.GetDataStore().Put(svc.RoleObject, id, ent)
		if err != nil {
			return err
		}
		return nil
	})
	router.Delete("/:id", func(ctx *echo.Context) error {
		id := ctx.P(0)
		log.Logger.Debugf("Deleting entity %s", id)
		err := svc.GetDataStore().Delete(svc.RoleObject, id)
		if err != nil {
			return err
		}
		return nil
	})

	svc.Router = router
	svc.dataSvcName = entitydatasvcInt.(string)
	return svc, nil
}

//Provides the name of the service
func (svc *RoleService) GetName() string {
	return ENTITY_ROLE_SERVICE_NAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *RoleService) Initialize(ctx service.ServiceContext) error {

	dataSvc, err := ctx.GetService(svc.dataSvcName)
	if err != nil {
		return errors.RethrowError(entities.ENTITY_ERROR_MISSING_DATASVC, err, ENTITY_ROLE_SERVICE_NAME)
	}

	//check if user service name to be used has been provided, otherwise set default name
	svc.RoleObject = ctx.GetConfig().GetString(laatoocore.CONF_ENV_ROLE)

	svc.DataStore = dataSvc.(data.DataService)

	return nil
}

//The service starts serving when this method is called
func (svc *RoleService) Serve() error {
	return nil
}

//Type of service
func (svc *RoleService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

func (svc *RoleService) GetDataStore() data.DataService {
	return svc.DataStore
}

//Execute method
func (svc *RoleService) Execute(name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}
