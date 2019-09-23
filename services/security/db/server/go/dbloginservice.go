package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/auth"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	common "securitycommon"
	"golang.org/x/crypto/bcrypt"
)

const (
	//login path to be used for local and oauth authentication
	CONF_SECURITYSERVICE_REGISTRATIONSERVICE = "REGISTRATION"
	CONF_SECURITYSERVICE_DB                  = "DB_LOGIN"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_SECURITYSERVICE_REGISTRATIONSERVICE, Object: RegistrationService{}},
		core.PluginComponent{Name: CONF_SECURITYSERVICE_DB, Object: LoginService{}}}
}

type LoginService struct {
	core.Service
	name           string
	authHeader     string
	adminRole      string
	realm          string
	tokenGenerator func(auth.User, string) (string, auth.User, error)
	//data service to use for users
	UserDataService data.DataComponent
}

func (ls *LoginService) Initialize(ctx core.ServerContext, conf config.Config) error {
	sechandler := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sechandler == nil {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	authHeader := sechandler.GetProperty(config.AUTHHEADER)
	if authHeader == nil {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ls.authHeader = authHeader.(string)
	realm := sechandler.GetProperty(config.REALM)
	if realm == nil {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ls.realm = realm.(string)

	userObject := sechandler.GetProperty(config.USER)

	log.Error(ctx, "***********************Got user object", "userObject", userObject)

	err := ls.AddParamWithType(ctx, "credentials", userObject.(string))
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	return nil
}

//Expects Local user to be provided inside the request
func (ls *LoginService) Invoke(ctx core.RequestContext) error {

	ent, _ := ctx.GetParamValue("credentials")
	usr, ok := ent.(auth.LocalAuthUser)
	if !ok {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}

	realm := usr.GetRealm()
	if realm != ls.realm {
		log.Trace(ctx, "Realm not found")
		ctx.SetResponse(core.BadRequestResponse(common.AUTH_ERROR_REALM_MISMATCH))
		return nil
	}

	username := usr.GetUserName()
	log.Trace(ctx, "getting user from service", "username", username)

	argsMap := map[string]interface{}{usr.GetUsernameField(): username, config.REALM: realm}

	cond, err := ls.UserDataService.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}

	usrs, _, _, recs, err := ls.UserDataService.Get(ctx, cond, -1, -1, "", "")
	if err != nil || recs <= 0 {
		log.Trace(ctx, "Tested user not found", "Err", err)
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}

	//compare the user requested with the user from database
	existingUser := usrs[0].(auth.LocalAuthUser)
	log.Trace(ctx, "got user********", "existingUser", existingUser, "tried password", usr.GetPassword())
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.GetPassword()), []byte(usr.GetPassword()))
	log.Trace(ctx, "compared user********", "err", err)
	existingUser.ClearPassword()
	if err != nil {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	} else {
		existingUser.ClearPassword()

		if ls.tokenGenerator == nil {
			ctx.SetResponse(core.StatusUnauthorizedResponse)
			return nil

		}
		token, user, err := ls.tokenGenerator(existingUser, realm)
		if err != nil {
			ctx.SetResponse(core.StatusUnauthorizedResponse)
			return nil
		}

		info := map[string]interface{}{ls.authHeader: token}

		err = ctx.SendSynchronousMessage(common.EVT_LOGIN_SUCCESS, map[string]interface{}{"Data": user, "info": info})
		if err != nil {
			log.Error(ctx, "Encountered Error in sending event", "error", err)
		}
		ctx.SetResponse(core.SuccessResponseWithInfo(user, info))
	}
	return nil
}
func (ls *LoginService) SetTokenGenerator(ctx core.ServerContext, gen func(auth.User, string) (string, auth.User, error)) {
	ls.tokenGenerator = gen
}

func (ls *LoginService) Start(ctx core.ServerContext) error {
	userDataSvcName, _ := ls.GetConfiguration(ctx, common.CONF_LOGINSERVICE_USERDATASERVICE)
	userService, err := ctx.GetService(userDataSvcName.(string))
	if err != nil {
		return errors.RethrowError(ctx, common.AUTH_ERROR_MISSING_USER_DATA_SERVICE, err)
	}
	userDataService, ok := userService.(data.DataComponent)
	if !ok {
		return errors.ThrowError(ctx, common.AUTH_ERROR_MISSING_USER_DATA_SERVICE)
	}
	log.Debug(ctx, "User storer set for login")

	//get and set the data service for accessing users
	ls.UserDataService = userDataService
	return nil
}
