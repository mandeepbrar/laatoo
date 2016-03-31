package security

import (
	"laatoo/core/registry"
	"laatoo/sdk/auth"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/data"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	//login path to be used for local and oauth authentication
	CONF_REGISTRATIONSERVICE_USERDATASERVICE = "user_data_svc"
	CONF_DEF_ROLE                            = "def_role"
)

// SecurityService contains a configuration and other details for running.
type RegistrationService struct {
	conf config.Config
	//authentication mode for service
	AuthMode string
	//admin role
	DefaultRole string
	userCreator core.ObjectCreator
	//user data service name
	userDataSvcName string
	//data service to use for users
	UserDataService data.DataService
}

//service method for doing various tasks
func NewRegistrationService(ctx core.ServerContext, conf config.Config) (core.Service, error) {
	rs := &RegistrationService{conf: conf}
	defrole, ok := rs.conf.GetString(CONF_DEF_ROLE)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_DEF_ROLE)
	}
	rs.DefaultRole = defrole
	userObject := ctx.GetServerVariable(core.USER).(string)
	userCreator, err := registry.GetObjectCreator(ctx, userObject)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	rs.userCreator = userCreator
	userDataSvcName, ok := rs.conf.GetString(CONF_REGISTRATIONSERVICE_USERDATASERVICE)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_DEF_ROLE)
	}
	rs.userDataSvcName = userDataSvcName
	return rs, nil
}

func (rs *RegistrationService) Initialize(ctx core.ServerContext) error {
	userService, err := ctx.GetService(rs.userDataSvcName)
	if err != nil {
		return errors.RethrowError(ctx, AUTH_ERROR_MISSING_USER_DATA_SERVICE, err)
	}
	userDataService, ok := userService.(data.DataService)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_MISSING_USER_DATA_SERVICE)
	}
	log.Logger.Debug(ctx, "User storer set for registration")
	//get and set the data service for accessing users
	rs.UserDataService = userDataService
	return nil
}

//Expects Rbac user to be provided inside the request
func (rs *RegistrationService) Invoke(ctx core.RequestContext) error {
	ent := ctx.GetRequestBody()
	user, ok := ent.(auth.RbacUser)
	id := user.GetId()
	existinguser, _ := rs.UserDataService.GetById(ctx, id)
	if existinguser != nil {
		return errors.ThrowError(ctx, AUTH_ERROR_USER_EXISTS)
	}
	user.SetRoles([]string{rs.DefaultRole})
	storable, ok := ent.(data.Storable)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_TYPE_MISMATCH)
	}
	err := rs.UserDataService.Save(ctx, storable)
	if err != nil {
		return err
	}
	log.Logger.Trace(ctx, "Saved user")
	return nil
}

func (rs *RegistrationService) GetConf() config.Config {
	return rs.conf
}
func (rs *RegistrationService) GetResponseHandler() core.ServiceResponseHandler {
	return nil
}
