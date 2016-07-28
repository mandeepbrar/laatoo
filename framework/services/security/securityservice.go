package security

import (
	"laatoo/framework/core/objects"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

const (
	CONF_SECURITYSERVICE_SERVICEPROVIDER     = "security_service"
	CONF_SECURITYSERVICE_REGISTRATIONSERVICE = "REGISTRATION"
	CONF_SECURITYSERVICE_DB                  = "DB_LOGIN"
	CONF_SECURITYSERVICE_OAUTH               = "OAUTH"
	CONF_SECURITYSERVICE_KEYAUTH             = "KEYAUTH"
)

func init() {
	objects.Register(CONF_SECURITYSERVICE_SERVICEPROVIDER, SecurityServiceFactory{})
}

type SecurityServiceFactory struct {
}

//Create the services configured for factory.
func (sf *SecurityServiceFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	switch method {
	case CONF_SECURITYSERVICE_DB:
		{
			return &LoginService{name: name}, nil
		}
	case CONF_SECURITYSERVICE_OAUTH:
		{
			return &OAuthLoginService{}, nil
		}
	case CONF_SECURITYSERVICE_KEYAUTH:
		{
			return &KeyAuthService{name: name}, nil
		}
	case CONF_SECURITYSERVICE_REGISTRATIONSERVICE:
		{
			return &RegistrationService{name: name}, nil
		}
	}
	return nil, nil
}

//The services start serving when this method is called
func (sf *SecurityServiceFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//The services start serving when this method is called
func (sf *SecurityServiceFactory) Start(ctx core.ServerContext) error {
	return nil
}

/*

const (
	CONF_SECURITYSERVICE_SERVICEPROVIDER = "security_service"
	CONF_SECURITYSERVICE_PERM            = "permissions"
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
type SecurityServiceProvider struct {
	//authentication mode for service
	AuthMode string
	//mailer to use for reminders
	Mailer Mailer
	//router to be used by the service. provided by the environment
	Router core.Router
	//login path for the applicaton
	LoginPath string
	//logout path for the application
	LogoutPath string
	//Local auth object
	AuthTypes map[string]AuthType
	//permissions array for security
	Permissions []string
}

//factory method returns the service object to the environment
func SecurityServiceFactory(ctx core.Context, conf map[string]interface{}) (interface{}, error) {

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

	//set the router object from configuration
	routerInt, ok := svc.Configuration[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(ctx, AUTH_ERROR_MISSING_ROUTER)
	}
	svc.Router = routerInt.(core.Router)

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
	svc.Router.Get(ctx, svc.LogoutPath, nil, svc.Logout)

	permissionsPath := "/permissions"
	//get the permissions path for the application
	permissionsPathInt, ok := svc.Configuration[CONF_SECURITYSERVICE_PERMPATH]
	if ok {
		permissionsPath = permissionsPathInt.(string)
	}

	svc.Router.Get(ctx, permissionsPath, svc.Configuration, func(ctx core.Context) error {
		return ctx.JSON(http.StatusOK, svc.Permissions)
	})

	//return the service
	return svc, nil
}


//Initialize the service. Consumer of a service passes the data
func (svc *SecurityService) Initialize(ctx core.Context) error {
	//setup the user data service from the context
	userDataInt, ok := svc.Configuration[CONF_SECURITYSERVICE_USERDATASERVICE]

	if ok {
		//get the name of the data service to be used for accessing users database
		svcAlias := userDataInt.(string)
		userService, err := ctx.GetService(svcAlias)
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
func (svc *SecurityService) Serve(ctx core.Context) error {
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

//Execute method
func (svc *SecurityService) Execute(reqContext core.Context, name string, params map[string]interface{}) (interface{}, error) {
	switch name {
	case "GetPermissions":
		return svc.Permissions, nil
	case "GetRoles":
		orderBy := "UpdatedOn"
		rolesInt, _, _, err := svc.UserDataService.Get(reqContext, svc.RoleObject, nil, -1, -1, "", orderBy)
		return rolesInt, err
	case "SaveRole":
		return nil, svc.UserDataService.Save(reqContext, svc.RoleObject, params["data"])
	}
	return nil, nil
}
*/
