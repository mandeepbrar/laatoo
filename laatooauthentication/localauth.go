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
	//encryption cost for local password encryption, if not provided, default is used
	CONF_AUTHSERVICE_BCRYPTCOST = "bcrypt_cost"
)

type localAuthType struct {
	// BCryptCost is the cost of the bcrypt password hashing function.
	bCryptCost int
	//login path to register for local authentication
	loginpath string
	//method called in case of callback
	authCallback echo.HandlerFunc
	//reference to the main auth service
	authService *AuthService
}

//method called for creating new auth type
func NewLocalAuth(conf map[string]interface{}, svc *AuthService) (*localAuthType, error) {
	//create the new auth type
	localauth := &localAuthType{}
	//store the reference to the parent
	localauth.authService = svc
	log.Logger.Debug("localAuthProvider: Initializing")

	//get the login path
	localauth.loginpath = "/login"
	loginpath, ok := conf[CONF_AUTHSERVICE_LOGINPATH]
	if ok {
		localauth.loginpath = loginpath.(string)
	}

	//get the bcryptcost from conf
	bcryptcost, ok := conf[CONF_AUTHSERVICE_BCRYPTCOST]
	if ok {
		localauth.bCryptCost = bcryptcost.(int)
	} else {
		localauth.bCryptCost = bcrypt.DefaultCost
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
	localauth.authService.Router.Post(localauth.loginpath, authStart)
	localauth.authCallback = authCallback
	return nil
}

//validate the local user
//derive the data from context object
func (localauth *localAuthType) ValidateUser(ctx *echo.Context) error {
	log.Logger.Debug("localAuthProvider: Validating Credentials")

	//create the user
	usrInt, err := localauth.authService.CreateUser()
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
	testedUser, err := localauth.authService.GetUserById(id)
	if err != nil {
		return errors.RethrowHttpError(AUTH_ERROR_USER_NOT_FOUND, ctx, err)
	}

	//compare the user requested with the user from database
	existingUser := testedUser.(auth.LocalAuthUser)
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.GetPassword()), []byte(usr.GetPassword()))
	if err != nil {
		return errors.RethrowHttpError(AUTH_ERROR_WRONG_PASSWORD, ctx, err)
	} else {
		ctx.Set("User", testedUser)
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
