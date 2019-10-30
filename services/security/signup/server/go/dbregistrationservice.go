package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/auth"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"reflect"
	common "securitycommon"
)

const (
	//login path to be used for local and oauth authentication
	CONF_SECURITYSERVICE_REGISTRATIONSERVICE = "REGISTRATION"
	CONF_EMAIL_TASKPROCESSOR                 = "SignupEmailTask"
	CONF_REGISTRATIONSERVICE_USERDATASERVICE = "user_data_svc"
	CONF_DEF_ROLE                            = "def_role"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_SECURITYSERVICE_REGISTRATIONSERVICE, Object: RegistrationService{}},
		core.PluginComponent{Name: CONF_EMAIL_TASKPROCESSOR, Object: SignupEmailTask{}},
	}
}

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
	UserDataService   data.DataComponent
	WorkflowName      string
	WorkflowInitiator components.WorkflowInitiator
	WorkflowSupport   bool

	realm string
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

	if rs.WorkflowSupport {
		err = rs.AddParamWithType(ctx, "credentials", config.OBJECTTYPE_STRINGSMAP)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	} else {
		err = rs.AddParamWithType(ctx, "credentials", rs.userObject)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	log.Error(ctx, "registration service", "WorkflowSupport", rs.WorkflowSupport, "Workflow initiator service", rs.WorkflowInitiator)
	/*
		rs.SetDescription("Db Registration service")
		rs.SetRequestType(config.CONF_OBJECT_STRINGMAP, false, false)
		rs.AddStringConfigurations([]string{CONF_REGISTRATIONSERVICE_USERDATASERVICE, CONF_DEF_ROLE}, nil)
	*/
	return nil
}

//Expects Rbac user to be provided inside the request
func (rs *RegistrationService) Invoke(ctx core.RequestContext) (err error) {
	if rs.WorkflowSupport {
		err = rs.registerWithWorkflowSupport(ctx)
	} else {
		err = rs.registerWithoutWorkflowSupport(ctx)
	}
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (rs *RegistrationService) registerWithWorkflowSupport(ctx core.RequestContext) error {
	emailMap, _ := ctx.GetStringsMapParam("credentials")
	log.Trace(ctx, "param value ", "ent", emailMap)
	email := emailMap["email"]
	if email != "" {
		go rs.WorkflowInitiator.StartWorkflow(ctx, rs.WorkflowName, emailMap)
	} else {
		return errors.BadRequest(ctx, "Missing email in request map", "email")
	}
	return nil
}
func (rs *RegistrationService) registerWithoutWorkflowSupport(ctx core.RequestContext) error {
	ent, _ := ctx.GetParamValue("credentials")
	log.Trace(ctx, "param value ", "ent", ent)
	//ent := ctx.GetBody()
	user, ok := ent.(auth.RbacUser)
	if !ok {
		log.Trace(ctx, "Not user", "type", reflect.TypeOf(ent))
		ctx.SetResponse(core.BadRequestResponse("Credentials user is not an RBAC User "))
		return nil
	}
	//fieldMap := *body
	//fieldMap["Roles"] = []string{rs.DefaultRole}
	user.SetRoles([]string{rs.DefaultRole})
	log.Trace(ctx, "data", " user", user)

	//obj := rs.userCreator()
	init := ent.(core.Initializable)
	init.Initialize(ctx, nil)
	//user := obj.(auth.User)

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
		ctx.SetResponse(core.InternalErrorResponse("Could not create condition for comparison"))
		return err
	}

	_, _, _, recs, err := rs.UserDataService.Get(ctx, cond, -1, -1, "", "")
	if err == nil && recs > 0 {
		log.Trace(ctx, "Tested user found")
		ctx.SetResponse(core.BadRequestResponse(common.AUTH_ERROR_USER_EXISTS))
		return nil
	}
	if err != nil {
		ctx.SetResponse(core.InternalErrorResponse("could not get the user" + err.Error()))
		return err
	}

	err = rs.UserDataService.Save(ctx, ent.(data.Storable))
	if err != nil {
		ctx.SetResponse(core.InternalErrorResponse("Could not save user"))
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
