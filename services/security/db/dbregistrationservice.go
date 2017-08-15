package main

import (
	"laatoo/sdk/auth"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/services/security/common"
)

const (
	//login path to be used for local and oauth authentication
	CONF_REGISTRATIONSERVICE_USERDATASERVICE = "user_data_svc"
	CONF_DEF_ROLE                            = "def_role"
)

// SecurityService contains a configuration and other details for running.
type RegistrationService struct {
	core.Service
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
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	realm := sechandler.GetProperty(config.REALM)
	if realm == nil {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	rs.realm = realm.(string)

	userObject := sechandler.GetProperty(config.USER)
	if userObject == nil {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	rs.userObject = userObject.(string)

	userCreator, err := ctx.GetObjectCreator(rs.userObject)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	rs.userCreator = userCreator
	/*
		rs.SetDescription("Db Registration service")
		rs.SetRequestType(config.CONF_OBJECT_STRINGMAP, false, false)
		rs.AddStringConfigurations([]string{CONF_REGISTRATIONSERVICE_USERDATASERVICE, CONF_DEF_ROLE}, nil)
	*/
	return nil
}

//Expects Rbac user to be provided inside the request
func (rs *RegistrationService) Invoke(ctx core.RequestContext) error {
	ent := ctx.GetBody()
	body, ok := ent.(*map[string]interface{})
	if !ok {
		log.Trace(ctx, "Not map")
		ctx.SetResponse(core.StatusBadRequestResponse)
		return nil
	}
	fieldMap := *body
	fieldMap["Roles"] = []string{rs.DefaultRole}
	log.Trace(ctx, "data", " map", fieldMap)

	obj := rs.userCreator()
	init := obj.(core.Initializable)
	init.Init(ctx, fieldMap)
	user := obj.(auth.User)

	username := user.GetUserName()
	if username == "" {
		log.Trace(ctx, "Username not found")
		ctx.SetResponse(core.BadRequestResponse(common.AUTH_ERROR_MISSING_USER))
		return nil
	}

	realm := user.GetRealm()
	if realm != rs.realm {
		log.Trace(ctx, "Realm not found")
		ctx.SetResponse(core.BadRequestResponse(common.AUTH_ERROR_REALM_MISMATCH))
		return nil
	}

	argsMap := map[string]interface{}{user.GetUsernameField(): username, config.REALM: realm}

	cond, err := rs.UserDataService.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		ctx.SetResponse(core.StatusInternalErrorResponse)
		return err
	}

	_, _, _, recs, err := rs.UserDataService.Get(ctx, cond, -1, -1, "", "")
	if err == nil && recs > 0 {
		log.Trace(ctx, "Tested user found")
		ctx.SetResponse(core.BadRequestResponse(common.AUTH_ERROR_USER_EXISTS))
		return nil
	}
	if err != nil {
		ctx.SetResponse(core.StatusInternalErrorResponse)
		return err
	}

	err = rs.UserDataService.Save(ctx, obj.(data.Storable))
	if err != nil {
		ctx.SetResponse(core.StatusInternalErrorResponse)
		return err
	}
	log.Trace(ctx, "Saved user")
	ctx.SetResponse(core.StatusSuccessResponse)
	return nil
}

func (rs *RegistrationService) Start(ctx core.ServerContext) error {

	userDataSvcName, _ := rs.GetConfiguration(ctx, CONF_REGISTRATIONSERVICE_USERDATASERVICE)
	rs.userDataSvcName = userDataSvcName.(string)

	defrole, ok := rs.GetConfiguration(ctx, CONF_DEF_ROLE)
	rs.DefaultRole = defrole.(string)

	userService, err := ctx.GetService(rs.userDataSvcName)
	if err != nil {
		return errors.RethrowError(ctx, common.AUTH_ERROR_MISSING_USER_DATA_SERVICE, err)
	}
	userDataService, ok := userService.(data.DataComponent)
	if !ok {
		return errors.ThrowError(ctx, common.AUTH_ERROR_MISSING_USER_DATA_SERVICE)
	}
	log.Debug(ctx, "User storer set for registration")
	//get and set the data service for accessing users
	rs.UserDataService = userDataService
	return nil
}
