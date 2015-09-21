package laatooauthentication

import (
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"laatoocore"
	"laatoosdk/auth"
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
	authCallback echo.HandlerFunc
	//reference to the main auth service
	securityService *SecurityService
}

//method called for creating new auth type
func NewLocalAuth(conf map[string]interface{}, svc *SecurityService) (*localAuthType, error) {
	//create the new auth type
	localauth := &localAuthType{}
	//store the reference to the parent
	localauth.securityService = svc
	log.Logger.Debug("localAuthProvider: Initializing")

	//get the login path
	localauth.loginpath = "/login"
	loginpath, ok := conf[CONF_AUTHSERVICE_LOGINPATH]
	if ok {
		localauth.loginpath = loginpath.(string)
	}
	return localauth, nil
}

//method called for service
func (localauth *localAuthType) Serve() error {
	return nil
}

//initialize auth type called by base auth for initializing
func (localauth *localAuthType) InitializeType(authStart echo.HandlerFunc, authCallback echo.HandlerFunc) error {
	//setup path for listening to login post request
	localauth.securityService.Router.Post(localauth.loginpath, authStart)
	localauth.authCallback = authCallback
	return nil
}

//validate the local user
//derive the data from context object
func (localauth *localAuthType) ValidateUser(ctx *echo.Context) error {
	log.Logger.Debug("localAuthProvider: Validating Credentials")

	//create the user
	usrInt, err := localauth.securityService.CreateUser()
	if err != nil {
		return errors.RethrowHttpError(laatoocore.AUTH_ERROR_USEROBJECT_NOT_CREATED, ctx, err)
	}

	//ctx.Request().Body
	err = ctx.Bind(usrInt)
	if err != nil {
		return errors.RethrowHttpError(AUTH_ERROR_INCORRECT_REQ_FORMAT, ctx, err)
	}

	//get the ide of the user to be tested
	usr := usrInt.(auth.LocalAuthUser)
	id := usr.GetId()

	//get the tested user from database
	testedUser, err := localauth.securityService.GetUserById(id)
	if err != nil {
		return errors.RethrowHttpError(AUTH_ERROR_USER_NOT_FOUND, ctx, err)
	}
	if testedUser == nil {
		return errors.ThrowHttpError(AUTH_ERROR_USER_NOT_FOUND, ctx)
	}

	//compare the user requested with the user from database
	existingUser := testedUser.(auth.LocalAuthUser)
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.GetPassword()), []byte(usr.GetPassword()))
	existingUser.SetPassword("")
	if err != nil {
		return errors.RethrowHttpError(AUTH_ERROR_WRONG_PASSWORD, ctx, err)
	} else {
		existingUser.SetPassword("")
		ctx.Set("User", testedUser)
		log.Logger.Debugf("**********Auth user", testedUser)
		return localauth.authCallback(ctx)
	}
}

func (localauth *localAuthType) GetName() string {
	return "local"
}

//complete authentication
func (localauth *localAuthType) CompleteAuthentication(ctx *echo.Context) error {
	log.Logger.Info("localAuthProvider: Authentication Successful")
	return nil
}
