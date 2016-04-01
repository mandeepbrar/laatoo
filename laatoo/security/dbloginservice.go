package security

import (
	"golang.org/x/crypto/bcrypt"
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
	CONF_LOGINSERVICE_USERDATASERVICE = "user_data_svc"
)

//service method for doing various tasks
func NewLoginService(ctx core.ServerContext, conf config.Config) (core.Service, error) {
	return &LoginService{conf: conf}, nil
}

type LoginService struct {
	conf        config.Config
	userCreator core.ObjectCreator
	adminRole   string
	//data service to use for users
	UserDataService data.DataService
}

func (ls *LoginService) Initialize(ctx core.ServerContext) error {
	userobject := ctx.GetServerVariable(core.USER)
	userCreator, err := registry.GetObjectCreator(ctx, userobject.(string))
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	ls.userCreator = userCreator
	userDataSvcName, ok := ls.conf.GetString(CONF_REGISTRATIONSERVICE_USERDATASERVICE)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_REGISTRATIONSERVICE_USERDATASERVICE)
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
	ent := ctx.GetRequestBody()
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
		resp, err := completeAuthentication(ctx, existingUser)
		if err != nil {
			ctx.SetResponse(core.StatusUnauthorizedResponse)
			return nil
		}
		ctx.SetResponse(resp)
	}
	return nil
}

func (ks *LoginService) GetConf() config.Config {
	return ks.conf
}
func (ks *LoginService) GetResponseHandler() core.ServiceResponseHandler {
	return nil
}
