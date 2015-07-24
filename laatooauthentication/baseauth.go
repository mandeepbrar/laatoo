package laatooauthentication

import (
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"laatoosdk/user"
	"laatoosdk/utils"
)

const (
	CONF_AUTHSERVICE_SERVICENAME = "auth_service"
	//header set by the service
	CONF_AUTHSERVICE_AUTHHEADER = "auth_header"
	//secret key for jwt
	CONF_AUTHSERVICE_JWTSECRETKEY = "jwtsecretkey"
	//name of the user object. if not provided, default user is used
	CONF_AUTHSERVICE_USEROBJECT = "user_object"
	//alias of rthe user data service
	CONF_AUTHSERVICE_USERDATASERVICE = "user_data_svc"
	//authentication mode for the service
	CONF_AUTHSERVICE_AUTHMODE      = "auth_mode"
	CONF_AUTHSERVICE_AUTHMODELOCAL = "local"
	CONF_AUTHSERVICE_AUTHMODEOAUTH = "oauth"
	CONF_AUTHSERVICE_AUTHMODALL    = "all"
	//providers to enable for oauth authentication
	CONF_AUTHSERVICE_OAUTHPROVIDERS = "oauth_providers"
	//logout path for the service
	CONF_AUTHSERVICE_LOGOUTPATH = "logout_path"
)

// AuthService contains a configuration and other details for running.
type AuthService struct {
	//authentication mode for service
	AuthMode string
	//mailer to use for reminders
	Mailer Mailer
	//user object for the service
	UserObject string
	//data service to use for users
	UserDataService data.DataService
	//context for the service for getting other services like data service and objects
	Context service.ServiceContext
	//router to be used by the service. provided by the environment
	Router *echo.Group
	//login path for the applicaton
	LoginPath string
	//logout path for the application
	LogoutPath string
	//configuration provided by the environment
	Configuration map[string]interface{}
	//header to be set with jwt token
	AuthHeader string
	//secret key for jwt token
	JWTSecret string
	//Local auth object
	AuthTypes map[string]AuthType
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_AUTHSERVICE_SERVICENAME, AuthServiceFactory)
}

//factory method returns the service object to the environment
func AuthServiceFactory(conf map[string]interface{}) (interface{}, error) {
	log.Logger.Infof("Creating auth service with alias")
	svc := &AuthService{}
	//store configuration object
	svc.Configuration = conf
	svc.AuthTypes = make(map[string]AuthType, 5)
	//check if auth mode has been provided, otherwise allow auth by all modes
	authModeInt, ok := svc.Configuration[CONF_AUTHSERVICE_AUTHMODE]
	if ok {
		svc.AuthMode = authModeInt.(string)
	} else {
		svc.AuthMode = CONF_AUTHSERVICE_AUTHMODALL
	}

	//check if jwt secret key has been provided, otherwise create a key from random numbers
	jwtSecretInt, ok := svc.Configuration[CONF_AUTHSERVICE_JWTSECRETKEY]
	if ok {
		svc.JWTSecret = jwtSecretInt.(string)
	} else {
		svc.JWTSecret = utils.RandomString(15)
	}

	//check if auth header to be set has been provided, otherwise set default token
	authTokenInt, ok := svc.Configuration[CONF_AUTHSERVICE_AUTHHEADER]
	if ok {
		svc.AuthHeader = authTokenInt.(string)
	} else {
		svc.AuthHeader = "Auth-Token"
	}

	//check if auth header to be set has been provided, otherwise set default token
	userObjectInt, ok := svc.Configuration[CONF_AUTHSERVICE_USEROBJECT]
	if ok {
		svc.UserObject = userObjectInt.(string)
	} else {
		svc.UserObject = user.CONF_DEFAULT_USER
	}

	//set the router object from configuration
	routerInt, ok := svc.Configuration[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(AUTH_ERROR_MISSING_ROUTER)
	}
	svc.Router = routerInt.(*echo.Group)

	//local auth is enabled, set local auth type
	if svc.AuthMode == CONF_AUTHSERVICE_AUTHMODALL || svc.AuthMode == CONF_AUTHSERVICE_AUTHMODELOCAL {
		svc.SetupLocalAuth(conf)
	}
	//if oauth mode has been enabled, set oauth mode
	if svc.AuthMode == CONF_AUTHSERVICE_AUTHMODALL || svc.AuthMode == CONF_AUTHSERVICE_AUTHMODEOAUTH {

	}

	//get the logout path for the application
	logoutPathInt, ok := svc.Configuration[CONF_AUTHSERVICE_LOGOUTPATH]
	if ok {
		svc.LogoutPath = logoutPathInt.(string)
	} else {
		svc.LogoutPath = "/logout"
	}

	//register logout route
	svc.Router.Get(svc.LogoutPath, svc.Logout)

	//return the service
	return svc, nil
}

//Provides the name of the service
func (svc *AuthService) GetName() string {
	return CONF_AUTHSERVICE_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *AuthService) Initialize(ctx service.ServiceContext) error {
	//store the context for the service
	svc.Context = ctx
	//setup the user data service from the context
	userDataInt, ok := svc.Configuration[CONF_AUTHSERVICE_USERDATASERVICE]

	if ok {
		//get the name of the data service to be used for accessing users database
		svcAlias := userDataInt.(string)
		userService, err := ctx.GetService(svcAlias)
		if err != nil {
			return errors.RethrowError(AUTH_ERROR_MISSING_USER_DATA_SERVICE, err)
		}
		userDataService, ok := userService.(data.DataService)
		if !ok {
			return errors.ThrowError(AUTH_ERROR_MISSING_USER_DATA_SERVICE)
		}
		log.Logger.Debug("User storer set for authenticaion")
		//get and set the data service for accessing users
		svc.UserDataService = userDataService
	} else {
		return errors.ThrowError(AUTH_ERROR_MISSING_USER_DATA_SERVICE)
	}
	return nil
}

//The service starts serving when this method is called
func (svc *AuthService) Serve() error {
	return nil
}

//Type of service
func (svc *AuthService) GetServiceType() string {
	return service.SERVICE_TYPE_APP
}
