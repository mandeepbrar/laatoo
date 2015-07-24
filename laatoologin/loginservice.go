package laatoologin

import (
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"net/http"
)

const (
	CONF_LOGIN_SERVICENAME = "login_service"
	CONF_LOGIN_PUBLICDIR   = "publicdir"
)

//Environment hosting an application
type LoginService struct {
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_LOGIN_SERVICENAME, LoginServiceFactory)
}

//factory method returns the service object to the environment
func LoginServiceFactory(conf map[string]interface{}) (interface{}, error) {
	log.Logger.Infof("Creating login service")
	svc := &LoginService{}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(LOGIN_ERROR_MISSING_ROUTER)
	}
	publicdir, ok := conf[CONF_LOGIN_PUBLICDIR]
	if !ok {
		return nil, errors.ThrowError(LOGIN_ERROR_MISSING_PUBLICDIR)
	}
	router := routerInt.(*echo.Group)
	log.Logger.Infof("Router %s", router)

	log.Logger.Infof("Designer service starting with page path %s", publicdir)
	router.Static("/", publicdir.(string))
	router.Get("/login", func(ctx *echo.Context) error {
		ctx.Redirect(http.StatusMovedPermanently, "/index.html")
		return nil
	})
	return svc, nil
}

//Provides the name of the service
func (svc *LoginService) GetName() string {
	return CONF_LOGIN_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *LoginService) Initialize(ctx service.ServiceContext) error {
	return nil
}

//The service starts serving when this method is called
func (svc *LoginService) Serve() error {
	return nil
}

//Type of service
func (svc *LoginService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}
