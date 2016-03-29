package laatooauthentication

import (
	"laatoocore"
	"laatoosdk/auth"
	"laatoosdk/core"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
	"net/http"
)

const (
	CONF_REGISTRATIONSERVICE_SERVICENAME = "registration_service"
	//alias of rthe user data service
	CONF_REGISTRATIONSERVICE_USERDATASERVICE = "user_data_svc"
	CONF_DEF_ROLE                            = "def_role"
)

// SecurityService contains a configuration and other details for running.
type RegistrationService struct {
	//authentication mode for service
	AuthMode string
	//admin role
	DefaultRole string
	//user object
	UserObject string
	//user data service name
	userDataSvcName string
	//data service to use for users
	UserDataService data.DataService
	//router to be used by the service. provided by the environment
	Router core.Router
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_REGISTRATIONSERVICE_SERVICENAME, RegistrationServiceFactory)
}

//factory method returns the service object to the environment
func RegistrationServiceFactory(ctx core.Context, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Creating registration service ")
	svc := &RegistrationService{}

	svc.DefaultRole = conf[CONF_DEF_ROLE].(string)
	svc.UserObject = ctx.GetVariable(laatoocore.CONF_ENV_USER).(string)
	svc.userDataSvcName = conf[CONF_REGISTRATIONSERVICE_USERDATASERVICE].(string)

	//set the router object from configuration
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(ctx, AUTH_ERROR_MISSING_ROUTER)
	}
	svc.Router = routerInt.(core.Router)
	svc.Router.Post(ctx, "", conf, func(ctx core.Context) error {
		ent, err := laatoocore.CreateEmptyObject(ctx, svc.UserObject)
		if err != nil {
			return err
		}
		err = ctx.Bind(ent)
		if err != nil {
			return err
		}
		user := ent.(auth.RbacUser)
		id := user.GetId()
		existinguser, _ := svc.UserDataService.GetById(ctx, svc.UserObject, id)
		if existinguser != nil {
			return errors.ThrowError(ctx, AUTH_ERROR_USER_EXISTS)
		}
		user.SetRoles([]string{svc.DefaultRole})
		err = svc.UserDataService.Save(ctx, svc.UserObject, ent)
		if err != nil {
			return err
		}
		log.Logger.Trace(ctx, LOGGING_CONTEXT, "Saved user")
		return ctx.NoContent(http.StatusOK)
	})

	//return the service
	return svc, nil
}

//Provides the name of the service
func (svc *RegistrationService) GetName() string {
	return CONF_REGISTRATIONSERVICE_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *RegistrationService) Initialize(ctx core.Context) error {
	//setup the user data service from the context

	if svc.userDataSvcName != "" {
		userService, err := ctx.GetService(svc.userDataSvcName)
		if err != nil {
			return errors.RethrowError(ctx, AUTH_ERROR_MISSING_USER_DATA_SERVICE, err)
		}
		userDataService, ok := userService.(data.DataService)
		if !ok {
			return errors.ThrowError(ctx, AUTH_ERROR_MISSING_USER_DATA_SERVICE)
		}
		log.Logger.Debug(ctx, LOGGING_CONTEXT, "User storer set for registration")
		//get and set the data service for accessing users
		svc.UserDataService = userDataService

	} else {
		return errors.ThrowError(ctx, AUTH_ERROR_MISSING_USER_DATA_SERVICE)
	}
	return nil
}

//The service starts serving when this method is called
func (svc *RegistrationService) Serve(ctx core.Context) error {
	return nil
}

//Type of service
func (svc *RegistrationService) GetServiceType() string {
	return core.SERVICE_TYPE_APP
}

//Execute method
func (svc *RegistrationService) Execute(reqContext core.Context, name string, params map[string]interface{}) (interface{}, error) {
	return nil, nil
}
