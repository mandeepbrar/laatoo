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
	userObject  string
	name        string
	userCreator core.ObjectCreator
	//user data service name
	userDataSvcName string
	//data service to use for users
	UserDataService data.DataComponent
	realm           string
}

func (rs *RegistrationService) Initialize(ctx core.ServerContext, conf config.Config) error {
	sechandler := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sechandler == nil {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	realm, ok := sechandler.GetString(config.REALM)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	rs.realm = realm

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
	body, ok := ent.(*map[string]interface{})
	if !ok {
		log.Logger.Trace(ctx, "Not map")
		ctx.SetResponse(core.StatusBadRequestResponse)
		return nil
	}
	fieldMap := *body
	fieldMap["Roles"] = []string{rs.DefaultRole}
	log.Logger.Trace(ctx, "data", " map", fieldMap)

	obj, nil := rs.userCreator(ctx, fieldMap)
	user := obj.(auth.User)

	username := user.GetUserName()
	if username == "" {
		log.Logger.Trace(ctx, "Username not found")
		ctx.SetResponse(core.BadRequestResponse(AUTH_ERROR_MISSING_USER))
		return nil
	}

	realm := user.GetRealm()
	if realm != rs.realm {
		log.Logger.Trace(ctx, "Realm not found")
		ctx.SetResponse(core.BadRequestResponse(AUTH_ERROR_REALM_MISMATCH))
		return nil
	}

	argsMap := map[string]interface{}{user.GetUsernameField(): username, config.REALM: realm}

	cond, err := rs.UserDataService.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		return err
	}

	_, _, _, recs, err := rs.UserDataService.Get(ctx, cond, -1, -1, "", "")
	if err == nil && recs > 0 {
		log.Logger.Trace(ctx, "Tested user found")
		ctx.SetResponse(core.BadRequestResponse(AUTH_ERROR_USER_EXISTS))
		return nil
	}
	if err != nil {
		return err
	}

	err = rs.UserDataService.Save(ctx, obj.(data.Storable))
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
