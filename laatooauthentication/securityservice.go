package laatooauthentication

import (
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/auth"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"reflect"
)

const (
	LOGGING_CONTEXT = "securityservice"

	CONF_SECURITYSERVICE_SERVICENAME = "security_service"
	//alias of rthe user data service
	CONF_SECURITYSERVICE_USERDATASERVICE = "user_data_svc"
	//authentication mode for the service
	CONF_SECURITYSERVICE_AUTHMODE      = "auth_mode"
	CONF_SECURITYSERVICE_AUTHMODELOCAL = "local"
	CONF_SECURITYSERVICE_AUTHMODEOAUTH = "oauth"
	CONF_SECURITYSERVICE_SEEDUSER      = "seeduser"
	CONF_SECURITYSERVICE_SEEDUSER_ID   = "user"
	CONF_SECURITYSERVICE_SEEDUSER_PASS = "password"
	CONF_SECURITYSERVICE_AUTHMODALL    = "all"
	CONF_SECURITYSERVICE_AUTHMODEKEY   = "key"
	//providers to enable for oauth authentication
	CONF_SECURITYSERVICE_OAUTHPROVIDERS = "oauth_providers"
	//logout path for the service
	CONF_SECURITYSERVICE_LOGOUTPATH = "logout_path"
)

// SecurityService contains a configuration and other details for running.
type SecurityService struct {
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
	laatoocore.RegisterObjectProvider(CONF_SECURITYSERVICE_SERVICENAME, SecurityServiceFactory)
}

//factory method returns the service object to the environment
func SecurityServiceFactory(conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(LOGGING_CONTEXT, "Creating auth service with alias")
	svc := &SecurityService{}
	//store configuration object
	svc.Configuration = conf
	svc.AuthTypes = make(map[string]AuthType, 5)
	//check if auth mode has been provided, otherwise allow auth by all modes
	authModeInt, ok := svc.Configuration[CONF_SECURITYSERVICE_AUTHMODE]
	if ok {
		svc.AuthMode = authModeInt.(string)
	} else {
		svc.AuthMode = CONF_SECURITYSERVICE_AUTHMODALL
	}

	//check if jwt secret key has been provided, otherwise create a key from random numbers
	svc.JWTSecret, _ = conf[laatoocore.CONF_ENV_JWTSECRETKEY].(string)

	//check if auth header to be set has been provided, otherwise set default token
	svc.AuthHeader, _ = conf[laatoocore.CONF_ENV_AUTHHEADER].(string)

	//set the router object from configuration
	routerInt, ok := svc.Configuration[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(AUTH_ERROR_MISSING_ROUTER)
	}
	svc.Router = routerInt.(*echo.Group)

	//local auth is enabled, set local auth type
	if svc.AuthMode == CONF_SECURITYSERVICE_AUTHMODALL || svc.AuthMode == CONF_SECURITYSERVICE_AUTHMODELOCAL {
		svc.SetupLocalAuth(conf)
	}
	//if oauth mode has been enabled, set oauth mode
	if svc.AuthMode == CONF_SECURITYSERVICE_AUTHMODALL || svc.AuthMode == CONF_SECURITYSERVICE_AUTHMODEOAUTH {

	}
	//if oauth mode has been enabled, set oauth mode
	if svc.AuthMode == CONF_SECURITYSERVICE_AUTHMODALL || svc.AuthMode == CONF_SECURITYSERVICE_AUTHMODEKEY {
		svc.SetupKeyAuth(conf)
	}

	//get the logout path for the application
	logoutPathInt, ok := svc.Configuration[CONF_SECURITYSERVICE_LOGOUTPATH]
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
func (svc *SecurityService) GetName() string {
	return CONF_SECURITYSERVICE_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *SecurityService) Initialize(ctx service.ServiceContext) error {
	//store the context for the service
	svc.Context = ctx
	//setup the user data service from the context
	userDataInt, ok := svc.Configuration[CONF_SECURITYSERVICE_USERDATASERVICE]

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
		log.Logger.Debug(LOGGING_CONTEXT, "User storer set for authenticaion")
		//get and set the data service for accessing users
		svc.UserDataService = userDataService

		//check if user service name to be used has been provided, otherwise set default name
		svc.UserObject = laatoocore.SystemUser

	} else {
		return errors.ThrowError(AUTH_ERROR_MISSING_USER_DATA_SERVICE)
	}
	return nil
}

//The service starts serving when this method is called
func (svc *SecurityService) Serve() error {
	seedUserInt, ok := svc.Configuration[CONF_SECURITYSERVICE_SEEDUSER]
	if ok {
		seedUserConf := seedUserInt.(map[string]interface{})
		userId := seedUserConf[CONF_SECURITYSERVICE_SEEDUSER_ID].(string)
		user, _ := svc.UserDataService.GetById(laatoocore.SystemUser, userId)
		if user == nil {
			log.Logger.Info(LOGGING_CONTEXT, "Creating seed user", "ID", userId)
			suserInt, err := laatoocore.CreateEmptyObject(laatoocore.SystemUser)
			sUser := suserInt.(auth.RbacUser)
			sLocalUser := suserInt.(auth.LocalAuthUser)
			sUser.SetId(userId)
			sUser.SetRoles([]string{laatoocore.AdminRole})
			sLocalUser.SetPassword(seedUserConf[CONF_SECURITYSERVICE_SEEDUSER_PASS].(string))
			err = svc.UserDataService.Save(laatoocore.SystemUser, sUser)
			if err != nil {
				return err
			}
		}
	}
	rolesInt, _, _, err := svc.UserDataService.Get(laatoocore.SystemRole, nil, -1, -1, "")
	if err != nil {
		return err
	}
	adminExists := false
	anonExists := false
	if rolesInt != nil {
		arr := reflect.ValueOf(rolesInt).Elem()
		length := arr.Len()
		for i := 0; i < length; i++ {
			role := arr.Index(i).Addr().Interface().(auth.Role)
			if role.GetId() == "Anonymous" {
				anonExists = true
			}
			if role.GetId() == laatoocore.AdminRole {
				adminExists = true
			}
			laatoocore.RegisterRolePermissions(role)
		}
	}

	if !anonExists {
		aroleInt, err := laatoocore.CreateEmptyObject(laatoocore.SystemRole)
		anonymousRole := aroleInt.(auth.Role)
		anonymousRole.SetId("Anonymous")
		err = svc.UserDataService.Save(laatoocore.SystemRole, anonymousRole)
		if err != nil {
			return err
		}
	}
	if !adminExists {
		aroleInt, err := laatoocore.CreateEmptyObject(laatoocore.SystemRole)
		adminRole := aroleInt.(auth.Role)
		adminRole.SetId(laatoocore.AdminRole)
		err = svc.UserDataService.Save(laatoocore.SystemRole, adminRole)
		if err != nil {
			return err
		}
	}
	return nil
}

//Type of service
func (svc *SecurityService) GetServiceType() string {
	return service.SERVICE_TYPE_APP
}

//Execute method
func (svc *SecurityService) Execute(name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}
