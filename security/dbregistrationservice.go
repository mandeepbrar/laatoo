package security

import (
	"laatoo/sdk/auth"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
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
	//authentication mode for service
	AuthMode string
	//admin role
	DefaultRole string
	jwtSecret   string
	authHeader  string
	userObject  string
	name        string
	userCreator core.ObjectCreator
	//user data service name
	userDataSvcName string
	//data service to use for users
	UserDataService data.DataComponent
}

func (rs *RegistrationService) Initialize(ctx core.ServerContext, conf config.Config) error {
	sechandler := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sechandler == nil {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	jwtSecret, ok := sechandler.GetString(config.JWTSECRET)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	rs.jwtSecret = jwtSecret
	authHeader, ok := sechandler.GetString(config.AUTHHEADER)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	rs.authHeader = authHeader
	userObject, ok := sechandler.GetString(config.USER)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	rs.userObject = userObject
	defrole, ok := conf.GetString(CONF_DEF_ROLE)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_DEF_ROLE)
	}
	rs.DefaultRole = defrole
	userCreator, err := ctx.GetObjectCreator(userObject)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	rs.userCreator = userCreator
	userDataSvcName, ok := conf.GetString(CONF_REGISTRATIONSERVICE_USERDATASERVICE)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_DEF_ROLE)
	}
	rs.userDataSvcName = userDataSvcName
	return nil
}

//Expects Rbac user to be provided inside the request
func (rs *RegistrationService) Invoke(ctx core.RequestContext) error {
	ent := ctx.GetRequest()
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

func (rs *RegistrationService) Start(ctx core.ServerContext) error {
	userService, err := ctx.GetService(rs.userDataSvcName)
	if err != nil {
		return errors.RethrowError(ctx, AUTH_ERROR_MISSING_USER_DATA_SERVICE, err)
	}
	userDataService, ok := userService.(data.DataComponent)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_MISSING_USER_DATA_SERVICE)
	}
	log.Logger.Debug(ctx, "User storer set for registration")
	//get and set the data service for accessing users
	rs.UserDataService = userDataService
	return nil
}
