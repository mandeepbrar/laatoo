package main

import (
	"laatoo/sdk/auth"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/services/security/common"

	"golang.org/x/crypto/bcrypt"
)

const (
	//login path to be used for local and oauth authentication
	CONF_SECURITYSERVICE_REGISTRATIONSERVICE = "REGISTRATION"
	CONF_SECURITYSERVICE_DB                  = "DB_LOGIN"
)

func Manifest() []core.PluginComponent {
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

func (ls *LoginService) Initialize(ctx core.ServerContext) error {

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

	ls.SetDescription("Db Login service")
	ls.SetRequestType(userObject.(string), false, false)
	ls.AddStringConfigurations([]string{common.CONF_LOGINSERVICE_USERDATASERVICE}, nil)

	return nil
}

//Expects Local user to be provided inside the request
func (ls *LoginService) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {

	ent := req.GetBody()
	usr, ok := ent.(auth.LocalAuthUser)
	if !ok {
		return core.StatusUnauthorizedResponse, nil
	}

	realm := usr.GetRealm()
	if realm != ls.realm {
		log.Trace(ctx, "Realm not found")
		return core.BadRequestResponse(common.AUTH_ERROR_REALM_MISMATCH), nil
	}

	username := usr.GetUserName()
	log.Trace(ctx, "getting user from service", "username", username)

	argsMap := map[string]interface{}{usr.GetUsernameField(): username, config.REALM: realm}

	cond, err := ls.UserDataService.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		return core.StatusUnauthorizedResponse, nil
	}

	usrs, _, _, recs, err := ls.UserDataService.Get(ctx, cond, -1, -1, "", "")
	if err != nil || recs <= 0 {
		log.Trace(ctx, "Tested user not found", "Err", err)
		return core.StatusUnauthorizedResponse, nil
	}

	//compare the user requested with the user from database
	existingUser := usrs[0].(auth.LocalAuthUser)
	log.Trace(ctx, "got user********", "existingUser", existingUser)
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.GetPassword()), []byte(usr.GetPassword()))
	log.Trace(ctx, "compared user********", "err", err)
	existingUser.ClearPassword()
	if err != nil {
		return core.StatusUnauthorizedResponse, nil
	} else {
		existingUser.ClearPassword()
		token, user, err := ls.tokenGenerator(existingUser, realm)
		if err != nil {
			return core.StatusUnauthorizedResponse, nil
		}
		return core.NewServiceResponse(core.StatusSuccess, user, map[string]interface{}{ls.authHeader: token}), nil
	}
}
func (ls *LoginService) SetTokenGenerator(ctx core.ServerContext, gen func(auth.User, string) (string, auth.User, error)) {
	ls.tokenGenerator = gen
}

func (ls *LoginService) Start(ctx core.ServerContext) error {
	userDataSvcName, _ := ls.GetConfiguration(common.CONF_LOGINSERVICE_USERDATASERVICE)
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
