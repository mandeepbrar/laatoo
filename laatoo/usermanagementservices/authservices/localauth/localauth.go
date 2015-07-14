package ginauth_local

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ginauth"
	"golang.org/x/crypto/bcrypt"
	//jwt_lib "github.com/dgrijalva/jwt-go"
)

const (
	BCRYPTCOST          = "BCryptCost"
	LOGINPATH           = "LoginPath"
	REGISTRATIONALLOWED = "RegistrationAllowed"
	REGISTRATIONPATH    = "RegistrationPath"
)

func init() {
	localauth := &localAuth{}
	ginauth.RegisterModule("localauth", localauth)
	ginauth.RegisterAuthType("localauth", localauth)
}

type localAuth struct {
	app *ginauth.App
	// BCryptCost is the cost of the bcrypt password hashing function.
	bCryptCost          int
	loginpath           string
	registrationAllowed bool
	registrationPath    string
	authSuccessful      gin.HandlerFunc
	authFailure         gin.HandlerFunc
}

func (localauth *localAuth) Initialize(app *ginauth.App) error {
	localauth.app = app
	localauth.app.Logger.Debug("localAuthProvider: Initializing")
	err := localauth.configPaths()
	if err != nil {
		return err
	}
	err = localauth.configBCryptCost()
	if err != nil {
		return err
	}
	return nil
}

func (localauth *localAuth) Serve() error {
	return nil
}

func (localauth *localAuth) InitializeType(authStart gin.HandlerFunc, authFailed gin.HandlerFunc, authSuccessful gin.HandlerFunc) error {
	mountpath := fmt.Sprintf("%s/local", localauth.app.AppMountPath)
	groupRouter := localauth.app.Router.Group(mountpath)
	groupRouter.POST(localauth.loginpath, authStart)
	if localauth.registrationAllowed {
		groupRouter.POST(localauth.registrationPath, localauth.register)
	}
	localauth.authFailure = authFailed
	localauth.authSuccessful = authSuccessful
	return nil
}

func (localauth *localAuth) ValidateUser(ctx *gin.Context) error {
	localauth.app.Logger.Debug("localAuthProvider: Validating Credentials")
	usrInt := localauth.app.UserCreator()
	var err error
	if ctx.ContentType() == "application/json" {
		err = ctx.BindJSON(usrInt)
		//err = usr.BindJSON(ctx)
	} else if ctx.ContentType() == "multipart/form-data" || ctx.ContentType() == "application/x-www-form-urlencoded" {
		err = ctx.BindWith(usrInt, binding.Form)
		//		err = usr.BindForm(ctx)
	}
	if err != nil {
		ctx.Set("AuthError", err)
		localauth.authFailure(ctx)
		return nil
	}
	usr := usrInt.(LocalAuthUser)
	id := usr.GetId()
	testedUser, err := localauth.app.Storer.GetById(id)
	if err != nil {
		ctx.Set("AuthError", err)
		localauth.authFailure(ctx)
		return nil
	}
	existingUser := testedUser.(LocalAuthUser)
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.GetPassword()), []byte(usr.GetPassword()))
	if err != nil {
		ctx.Set("AuthError", err)
		localauth.authFailure(ctx)
		return nil
	} else {
		ctx.Set("User", testedUser)
		localauth.authSuccessful(ctx)
		return nil
	}
}
func (localauth *localAuth) CompleteAuthentication(ctx *gin.Context) error {
	localauth.app.Logger.Info("localAuthProvider: Authentication Successful")
	return nil
}
