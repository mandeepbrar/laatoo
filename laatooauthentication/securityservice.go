package laatooauthentication

import (
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/auth"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"net/http"
)

const (
	LOGGING_CONTEXT = "securityservice"

	CONF_SECURITYSERVICE_SERVICENAME = "security_service"
	CONF_SECURITYSERVICE_PERM        = "permissions"
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
	CONF_SECURITYSERVICE_PERMPATH   = "permissions_path"
)

// SecurityService contains a configuration and other details for running.
type SecurityService struct {
	//authentication mode for service
	AuthMode string
	//mailer to use for reminders
	Mailer Mailer
	//user object for the service
	UserObject string
	//role object for the service
	RoleObject string
	//admin role
	AdminRole string
	//data service to use for users
	UserDataService data.DataService
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
	//permissions array for security
	Permissions []string
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_SECURITYSERVICE_SERVICENAME, SecurityServiceFactory)
}

//factory method returns the service object to the environment
func SecurityServiceFactory(ctx *echo.Context, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Creating auth service with alias")
	svc := &SecurityService{}
	svcenv := ctx.Get(laatoocore.CONF_ENV_CONTEXT).(service.Environment)
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

	permInt, ok := conf[CONF_SECURITYSERVICE_PERM]
	if ok {
		perms := permInt.([]interface{})
		svc.Permissions = make([]string, len(perms))
		i := 0
		for _, k := range perms {
			svc.Permissions[i] = k.(string)
			i++
		}
	} else {
		svc.Permissions = []string{}
	}

	//check if jwt secret key has been provided, otherwise create a key from random numbers
	svc.JWTSecret, _ = svcenv.GetVariable(laatoocore.CONF_ENV_JWTSECRETKEY).(string)

	//check if auth header to be set has been provided, otherwise set default token
	svc.AuthHeader, _ = svcenv.GetVariable(laatoocore.CONF_ENV_AUTHHEADER).(string)

	svc.AdminRole = svcenv.GetVariable(laatoocore.CONF_ENV_ADMINROLE).(string)
	svc.UserObject = svcenv.GetVariable(laatoocore.CONF_ENV_USER).(string)
	svc.RoleObject = svcenv.GetVariable(laatoocore.CONF_ENV_ROLE).(string)

	//set the router object from configuration
	routerInt, ok := svc.Configuration[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(ctx, AUTH_ERROR_MISSING_ROUTER)
	}
	svc.Router = routerInt.(*echo.Group)

	//local auth is enabled, set local auth type
	if svc.AuthMode == CONF_SECURITYSERVICE_AUTHMODALL || svc.AuthMode == CONF_SECURITYSERVICE_AUTHMODELOCAL {
		svc.SetupLocalAuth(ctx, conf)
	}
	//if oauth mode has been enabled, set oauth mode
	if svc.AuthMode == CONF_SECURITYSERVICE_AUTHMODALL || svc.AuthMode == CONF_SECURITYSERVICE_AUTHMODEOAUTH {
		svc.SetupOAuth(ctx, conf)
	}
	//if oauth mode has been enabled, set oauth mode
	if svc.AuthMode == CONF_SECURITYSERVICE_AUTHMODALL || svc.AuthMode == CONF_SECURITYSERVICE_AUTHMODEKEY {
		svc.SetupKeyAuth(ctx, conf)
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

	permissionsPath := "/permissions"
	//get the permissions path for the application
	permissionsPathInt, ok := svc.Configuration[CONF_SECURITYSERVICE_PERMPATH]
	if ok {
		permissionsPath = permissionsPathInt.(string)
	}

	svc.Router.Get(permissionsPath, func(ctx *echo.Context) error {
		return ctx.JSON(http.StatusOK, svc.Permissions)
	})

	//return the service
	return svc, nil
}

//Provides the name of the service
func (svc *SecurityService) GetName() string {
	return CONF_SECURITYSERVICE_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *SecurityService) Initialize(ctx *echo.Context) error {
	//setup the user data service from the context
	userDataInt, ok := svc.Configuration[CONF_SECURITYSERVICE_USERDATASERVICE]
	svcenv := ctx.Get(laatoocore.CONF_ENV_CONTEXT).(service.Environment)

	if ok {
		//get the name of the data service to be used for accessing users database
		svcAlias := userDataInt.(string)
		userService, err := svcenv.GetService(ctx, svcAlias)
		if err != nil {
			return errors.RethrowError(ctx, AUTH_ERROR_MISSING_USER_DATA_SERVICE, err)
		}
		userDataService, ok := userService.(data.DataService)
		if !ok {
			return errors.ThrowError(ctx, AUTH_ERROR_MISSING_USER_DATA_SERVICE)
		}
		log.Logger.Debug(ctx, LOGGING_CONTEXT, "User storer set for authenticaion")
		//get and set the data service for accessing users
		svc.UserDataService = userDataService

	} else {
		return errors.ThrowError(ctx, AUTH_ERROR_MISSING_USER_DATA_SERVICE)
	}
	return nil
}

//The service starts serving when this method is called
func (svc *SecurityService) Serve(ctx *echo.Context) error {
	seedUserInt, ok := svc.Configuration[CONF_SECURITYSERVICE_SEEDUSER]
	if ok {
		seedUserConf := seedUserInt.(map[string]interface{})
		userId := seedUserConf[CONF_SECURITYSERVICE_SEEDUSER_ID].(string)
		user, _ := svc.UserDataService.GetById(ctx, svc.UserObject, userId)
		if user == nil {
			log.Logger.Info(ctx, LOGGING_CONTEXT, "Creating seed user", "ID", userId)
			suserInt, err := laatoocore.CreateEmptyObject(ctx, svc.UserObject)
			sUser := suserInt.(auth.RbacUser)
			sLocalUser := suserInt.(auth.LocalAuthUser)
			sUser.SetId(userId)
			sUser.SetRoles([]string{svc.AdminRole})
			sLocalUser.SetPassword(seedUserConf[CONF_SECURITYSERVICE_SEEDUSER_PASS].(string))
			err = svc.UserDataService.Save(ctx, svc.UserObject, sUser)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//Type of service
func (svc *SecurityService) GetServiceType() string {
	return service.SERVICE_TYPE_APP
}

//Execute method
func (svc *SecurityService) Execute(reqContext *echo.Context, name string, params map[string]interface{}) (interface{}, error) {
	switch name {
	case "GetPermissions":
		return svc.Permissions, nil
	case "GetRoles":
		rolesInt, _, _, err := svc.UserDataService.Get(reqContext, svc.RoleObject, nil, -1, -1, "")
		return rolesInt, err
	case "SaveRole":
		return nil, svc.UserDataService.Save(reqContext, svc.RoleObject, params["data"])
	}
	return nil, nil
}
