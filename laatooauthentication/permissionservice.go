package laatooauthentication

import (
	"github.com/labstack/echo"
	"laatoocore"
	//"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"net/http"
)

const (
	ENTITY_PERM_SERVICE_NAME = "perm_service"
)

type PermService struct {
	Context service.ServiceContext
}

func init() {
	laatoocore.RegisterObjectProvider(ENTITY_PERM_SERVICE_NAME, NewPermService)
}

func NewPermService(ctx interface{}, conf map[string]interface{}) (interface{}, error) {

	log.Logger.Info(ctx, LOGGING_CONTEXT, "Creating entity service", "ServiceName", ENTITY_PERM_SERVICE_NAME)
	serviceContext := ctx.(service.ServiceContext)
	svc := &PermService{Context: serviceContext}
	routerInt, _ := conf[laatoocore.CONF_ENV_ROUTER]
	router := routerInt.(*echo.Group)

	router.Get("", func(ctx *echo.Context) error {

		perms := svc.Context.ListAllPermissions()
		return ctx.JSON(http.StatusOK, perms)
	})
	return svc, nil
}

//Provides the name of the service
func (svc *PermService) GetName() string {
	return ENTITY_PERM_SERVICE_NAME
}

//Initialize the service. Consumer of a service passes the data
func (psvc *PermService) Initialize(ctx service.ServiceContext) error {
	return nil
}

//The service starts serving when this method is called
func (svc *PermService) Serve(reqContext interface{}) error {
	return nil
}

//Type of service
func (svc *PermService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

//Execute method
func (svc *PermService) Execute(ctx interface{}, name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}
