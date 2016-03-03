package laatooauthentication

import (
	"golang.org/x/crypto/bcrypt"
	"laatoocore"
	"laatoosdk/auth"
	"laatoosdk/core"
	"laatoosdk/errors"
	"laatoosdk/log"
)

const (
	//login path to be used for local and oauth authentication
	CONF_AUTHSERVICE_LOGINPATH = "login_path"
)

type localAuthType struct {
	//login path to register for local authentication
	loginpath string
	//method called in case of callback
	authCallback core.HandlerFunc
	//reference to the main auth service
	securityService *SecurityService
}

//method called for creating new auth type
func NewLocalAuth(ctx core.Context, conf map[string]interface{}, svc *SecurityService) (*localAuthType, error) {
	//create the new auth type
	localauth := &localAuthType{}
	//store the reference to the parent
	localauth.securityService = svc
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "localAuthProvider: Initializing")

	//get the login path
	localauth.loginpath = "/login"
	loginpath, ok := conf[CONF_AUTHSERVICE_LOGINPATH]
	if ok {
		localauth.loginpath = loginpath.(string)
	}
	return localauth, nil
}

//initialize auth type called by base auth for initializing
func (localauth *localAuthType) InitializeType(ctx core.Context, authStart core.HandlerFunc, authCallback core.HandlerFunc) error {
	//setup path for listening to login post request
	localauth.securityService.Router.Post(ctx, localauth.loginpath, map[string]interface{}{}, authStart)
	localauth.authCallback = authCallback
	return nil
}

//validate the local user
//derive the data from context object
func (localauth *localAuthType) ValidateUser(ctx core.Context) error {
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "localAuthProvider: Validating Credentials")

	//create the user
	usrInt, err := localauth.securityService.CreateUser(ctx)
	if err != nil {
		return errors.RethrowError(ctx, laatoocore.AUTH_ERROR_USEROBJECT_NOT_CREATED, err)
	}

	//ctx.Request().Body
	err = ctx.Bind(usrInt)
	if err != nil {
		return errors.RethrowError(ctx, AUTH_ERROR_INCORRECT_REQ_FORMAT, err)
	}

	//get the ide of the user to be tested
	usr := usrInt.(auth.LocalAuthUser)
	id := usr.GetId()
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Binding complete for id", "Id", id)

	//get the tested user from database
	testedUser, err := localauth.securityService.GetUserById(ctx, id)
	if err != nil {
		log.Logger.Info(ctx, LOGGING_CONTEXT, "Tested user not found", "Err", err)
		return errors.RethrowError(ctx, AUTH_ERROR_USER_NOT_FOUND, err)
	}
	if testedUser == nil {
		log.Logger.Info(ctx, LOGGING_CONTEXT, "Tested user not found")
		return errors.ThrowError(ctx, AUTH_ERROR_USER_NOT_FOUND)
	}

	//compare the user requested with the user from database
	existingUser := testedUser.(auth.LocalAuthUser)
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Comparing passwords", "existing", existingUser.GetPassword(), "to test", usr.GetPassword())
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.GetPassword()), []byte(usr.GetPassword()))
	existingUser.SetPassword("")
	if err != nil {
		return errors.RethrowError(ctx, AUTH_ERROR_WRONG_PASSWORD, err)
	} else {
		existingUser.SetPassword("")
		ctx.Set("User", testedUser)
		return localauth.authCallback(ctx)
	}
}

func (localauth *localAuthType) GetName() string {
	return "local"
}

//complete authentication
func (localauth *localAuthType) CompleteAuthentication(ctx core.Context) error {
	log.Logger.Info(ctx, LOGGING_CONTEXT, "localAuthProvider: Authentication Successful")
	return nil
}
