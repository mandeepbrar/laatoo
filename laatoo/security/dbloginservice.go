package security

import (
	"golang.org/x/crypto/bcrypt"
	"laatoo/sdk/auth"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/data"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	//login path to be used for local and oauth authentication
	CONF_LOGINSERVICE_USERDATASERVICE = "user_data_svc"
)

type LoginService struct {
	name        string
	jwtSecret   string
	authHeader  string
	userObject  string
	userCreator core.ObjectCreator
	adminRole   string
	//data service to use for users
	UserDataService data.DataService
}

func (ls *LoginService) Initialize(ctx core.ServerContext, conf config.Config) error {
	sechandler := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sechandler == nil {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	jwtSecret, ok := sechandler.GetString(config.JWTSECRET)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ls.jwtSecret = jwtSecret
	authHeader, ok := sechandler.GetString(config.AUTHHEADER)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ls.authHeader = authHeader
	userObject, ok := sechandler.GetString(config.USER)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	ls.userObject = userObject
	userCreator, err := ctx.GetObjectCreator(userObject)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	ls.userCreator = userCreator
	userDataSvcName, ok := conf.GetString(CONF_LOGINSERVICE_USERDATASERVICE)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_LOGINSERVICE_USERDATASERVICE)
	}
	userService, err := ctx.GetService(userDataSvcName)
	if err != nil {
		return errors.RethrowError(ctx, AUTH_ERROR_MISSING_USER_DATA_SERVICE, err)
	}
	userDataService, ok := userService.(data.DataService)
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
	id := usr.GetId()
	log.Logger.Debug(ctx, "getting user from service", "ls", ls, "id", id)
	//get the tested user from database
	testedUser, err := ls.UserDataService.GetById(ctx, id)
	if err != nil {
		log.Logger.Trace(ctx, "Tested user not found", "Err", err)
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	if testedUser == nil {
		log.Logger.Info(ctx, "Tested user not found")
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}

	//compare the user requested with the user from database
	existingUser := testedUser.(auth.LocalAuthUser)
	log.Logger.Info(ctx, "Comparing passwords", "existing", existingUser.GetPassword(), "to test", usr.GetPassword())
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.GetPassword()), []byte(usr.GetPassword()))
	existingUser.SetPassword("")
	if err != nil {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	} else {
		existingUser.SetPassword("")
		resp, err := completeAuthentication(ctx, existingUser, ls.jwtSecret, ls.authHeader)
		if err != nil {
			ctx.SetResponse(core.StatusUnauthorizedResponse)
			return nil
		}
		ctx.SetResponse(resp)
	}
	return nil
}

func (ls *LoginService) Start(ctx core.ServerContext) error {
	return nil
}
