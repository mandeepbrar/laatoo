package laatooauthentication

import (
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"laatoocore"
	"laatoosdk/auth"
	"laatoosdk/data"
	"laatoosdk/entities"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"net/http"
)

const (
	ENTITY_USER_SERVICE_NAME = "user_service"
	//encryption cost for local password encryption, if not provided, default is used
	CONF_AUTHSERVICE_BCRYPTCOST = "bcrypt_cost"
	PERM_USER_VIEW              = "View User"
	PERM_USER_EDIT              = "Edit User"
	PERM_USER_CREATE            = "Create User"
	PERM_USER_DEL               = "Delete User"
)

type UserService struct {
	DataStore   data.DataService
	Router      *echo.Group
	dataSvcName string
	UserObject  string
	bCryptCost  int
}

func init() {
	laatoocore.RegisterObjectProvider(ENTITY_USER_SERVICE_NAME, NewUserService)
	laatoocore.RegisterPermissions([]string{PERM_USER_VIEW, PERM_USER_EDIT, PERM_USER_CREATE, PERM_USER_DEL})
}

func NewUserService(conf map[string]interface{}) (interface{}, error) {

	log.Logger.Infof("Creating entity service", ENTITY_USER_SERVICE_NAME)

	svc := &UserService{}

	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	router := routerInt.(*echo.Group)

	//get the bcryptcost from conf
	bcryptcost, ok := conf[CONF_AUTHSERVICE_BCRYPTCOST]
	if ok {
		svc.bCryptCost = bcryptcost.(int)
	} else {
		svc.bCryptCost = bcrypt.DefaultCost
	}

	entitydatasvcInt, ok := conf[CONF_AUTHSERVICE_USERDATASERVICE]
	if !ok {
		return nil, errors.ThrowError(AUTH_ERROR_MISSING_USER_DATA_SERVICE, svc.GetName())
	}

	router.Get("/:id", func(ctx *echo.Context) error {
		id := ctx.P(0)
		log.Logger.Debugf("Getting entity %s", id)
		ent, err := svc.GetDataStore().GetById(svc.UserObject, id)
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, ent)
	})
	router.Post("", func(ctx *echo.Context) error {
		ent, err := laatoocore.CreateEmptyObject(svc.UserObject)
		if err != nil {
			return err
		}
		err = ctx.Bind(ent)
		if err != nil {
			return err
		}
		err = svc.encryptPassword(ent.(auth.LocalAuthUser))
		if err != nil {
			return err
		}
		err = svc.GetDataStore().Save(svc.UserObject, ent)
		if err != nil {
			return err
		}
		return nil
	})
	router.Put("/:id", func(ctx *echo.Context) error {
		id := ctx.P(0)
		log.Logger.Debugf("Updating entity %s", id)
		ent, err := laatoocore.CreateEmptyObject(svc.UserObject)
		if err != nil {
			return err
		}
		err = ctx.Bind(ent)
		if err != nil {
			return err
		}
		err = svc.encryptPassword(ent.(auth.LocalAuthUser))
		if err != nil {
			return err
		}
		err = svc.GetDataStore().Put(svc.UserObject, id, ent)
		if err != nil {
			return err
		}
		return nil
	})
	router.Delete("/:id", func(ctx *echo.Context) error {
		id := ctx.P(0)
		log.Logger.Debugf("Deleting entity %s", id)
		err := svc.GetDataStore().Delete(svc.UserObject, id)
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
func (svc *UserService) GetName() string {
	return ENTITY_USER_SERVICE_NAME
}

func (svc *UserService) encryptPassword(usr auth.LocalAuthUser) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(usr.GetPassword()), svc.bCryptCost)
	if err != nil {
		return err
	}
	usr.SetPassword(string(hash))
	return nil
}

//Initialize the service. Consumer of a service passes the data
func (svc *UserService) Initialize(ctx service.ServiceContext) error {

	dataSvc, err := ctx.GetService(svc.dataSvcName)
	if err != nil {
		return errors.RethrowError(entities.ENTITY_ERROR_MISSING_DATASVC, err, ENTITY_USER_SERVICE_NAME)
	}

	//check if user service name to be used has been provided, otherwise set default name
	svc.UserObject = ctx.GetConfig().GetString(laatoocore.CONF_ENV_USER)

	svc.DataStore = dataSvc.(data.DataService)

	return nil
}

//The service starts serving when this method is called
func (svc *UserService) Serve() error {
	return nil
}

//Type of service
func (svc *UserService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

func (svc *UserService) GetDataStore() data.DataService {
	return svc.DataStore
}

//Execute method
func (svc *UserService) Execute(name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}
