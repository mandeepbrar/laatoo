package security

import (
	"laatoo/sdk/auth"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"

	"golang.org/x/crypto/bcrypt"
)

const (
	//login path to be used for local and oauth authentication
	CONF_LOGINSERVICE_USERDATASERVICE = "user_data_svc"
)

type LoginService struct {
	name           string
	authHeader     string
	adminRole      string
	tokenGenerator func(auth.User, string) (string, auth.User, error)
	//data service to use for users
	UserDataService data.DataComponent
}

func (ls *LoginService) Initialize(ctx core.ServerContext, conf config.Config) error {
	sechandler := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sechandler == nil {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	authHeader, ok := sechandler.GetString(config.AUTHHEADER)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ls.authHeader = authHeader
	userDataSvcName, ok := conf.GetString(CONF_LOGINSERVICE_USERDATASERVICE)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_LOGINSERVICE_USERDATASERVICE)
	}
	userService, err := ctx.GetService(userDataSvcName)
	if err != nil {
		return errors.RethrowError(ctx, AUTH_ERROR_MISSING_USER_DATA_SERVICE, err)
	}
	userDataService, ok := userService.(data.DataComponent)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_MISSING_USER_DATA_SERVICE)
	}
	log.Logger.Debug(ctx, "User storer set for registration")
	//get and set the data service for accessing users
	ls.UserDataService = userDataService
	return nil
}

//Expects Local user to be provided inside the request
func (ls *LoginService) Invoke(ctx core.RequestContext) error {

	ent := ctx.GetRequest()
	usr, ok := ent.(auth.LocalAuthUser)
	if !ok {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}

	realm := usr.GetRealm()

	username := usr.GetUserName()
	log.Logger.Trace(ctx, "getting user from service", "username", username)

	argsMap := map[string]interface{}{usr.GetUsernameField(): username, config.REALM: realm}

	cond, err := ls.UserDataService.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}

	usrs, _, _, recs, err := ls.UserDataService.Get(ctx, cond, -1, -1, "", "")
	if err != nil || recs <= 0 {
		log.Logger.Trace(ctx, "Tested user not found", "Err", err)
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}

	//compare the user requested with the user from database
	existingUser := usrs[0].(auth.LocalAuthUser)
	log.Logger.Trace(ctx, "got user********", "existingUser", existingUser)
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.GetPassword()), []byte(usr.GetPassword()))
	log.Logger.Trace(ctx, "compared user********", "err", err)
	existingUser.ClearPassword()
	if err != nil {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	} else {
		existingUser.ClearPassword()
		token, user, err := ls.tokenGenerator(existingUser, realm)
		if err != nil {
			ctx.SetResponse(core.StatusUnauthorizedResponse)
			return nil
		}
		resp := core.NewServiceResponse(core.StatusSuccess, user, map[string]interface{}{ls.authHeader: token})
		ctx.SetResponse(resp)
	}
	return nil
}
func (ls *LoginService) SetTokenGenerator(ctx core.ServerContext, gen func(auth.User, string) (string, auth.User, error)) {
	ls.tokenGenerator = gen
}

func (ls *LoginService) Start(ctx core.ServerContext) error {
	return nil
}
