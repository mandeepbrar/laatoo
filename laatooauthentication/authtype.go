package laatooauthentication

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"laatoosdk/auth"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/utils"
	"net/http"
	"time"
)

//Interface that has to be implemented by every auth type interface
type AuthType interface {
	GetName() string
	//Initializes the authentication type module
	InitializeType(authStart echo.HandlerFunc, authCallback echo.HandlerFunc) error
	//Called to validate the user by providing context
	ValidateUser(*echo.Context) error
	//Completes authentication
	CompleteAuthentication(*echo.Context) error
}

//setup local authentication
func (svc *AuthService) SetupLocalAuth(conf map[string]interface{}) error {
	//create local authentication type
	localAuthType, err := NewLocalAuth(conf, svc)
	if err != nil {
		return err
	}
	//initialize local authentication
	err = svc.initializeAuthType(localAuthType)
	if err != nil {
		return err
	}
	return nil
}

//setup local authentication
func (svc *AuthService) SetupOAuth(conf map[string]interface{}) error {
	/*localAuthType, err : = localAuth.NewAuthType(conf)
	if(err !=nil) {
		return err
	}
	svc.Router.POST(svc.LoginPath, authStart)
	localauth.authFailure = authFailed
	localauth.authSuccessful = authSuccessful*/
	return nil
}

//The service starts serving when this method is called
func (svc *AuthService) initializeAuthType(authType AuthType) error {
	//initialize auth type
	initializationErr := authType.InitializeType(
		func(ctx *echo.Context) error { ///  auth start method starts
			log.Logger.Debug("Validating user")
			err := authType.ValidateUser(ctx)
			if err != nil {
				return errors.RethrowHttpError(AUTH_ERROR_USER_VALIDATION_FAILED, ctx, err)
			}
			return nil
		}, ///  auth start method ends
		func(ctx *echo.Context) error { ///  auth callback method starts
			err := authType.CompleteAuthentication(ctx)
			if err != nil {
				return errors.RethrowHttpError(AUTH_ERROR_AUTH_COMPLETION_FAILED, ctx, err)
			}
			userInt := ctx.Get("User")
			if userInt == nil {
				return errors.ThrowHttpError(AUTH_ERROR_INTERNAL_SERVER_ERROR_AUTH, ctx)
			}
			user := userInt.(auth.User)
			token := jwt.New(jwt.SigningMethodHS256)
			user.SetJWTClaims(token)
			token.Claims["UserId"] = user.GetId()
			token.Claims["AuthTypeName"] = authType.GetName()
			//token.Claims["IP"] = ctx.ClientIP()
			token.Claims["exp"] = time.Now().Add(time.Hour * 4).Unix()
			ctx.Set("JWT_Token", token)
			tokenString, err := token.SignedString([]byte(svc.JWTSecret))
			if err != nil {
				return errors.RethrowHttpError(AUTH_ERROR_JWT_CREATION, ctx, err)
			}
			ctx.Response().Header().Set(svc.AuthHeader, tokenString)

			utils.FireEvent(&utils.Event{EVENT_AUTHSERVICE_LOGIN_COMPLETE, ctx})
			ctx.JSON(http.StatusOK, user)
			return nil
		}) ///  auth logout method ends
	if initializationErr != nil {
		return errors.RethrowError(AUTH_ERROR_INITIALIZING_TYPE, initializationErr)
	}
	svc.Router.Get(svc.LogoutPath, svc.Logout)
	return nil
}

func (svc *AuthService) Logout(ctx *echo.Context) error {
	ctx.Response().Header().Set(svc.AuthHeader, "")
	ctx.Set("User", nil)
	utils.FireEvent(&utils.Event{EVENT_AUTHSERVICE_LOGOUT_COMPLETE, ctx})
	return nil
}

func (svc *AuthService) CreateUser() (interface{}, error) {
	return svc.Context.CreateObject(svc.UserObject, nil)
}

func (svc *AuthService) GetUserById(id string) (interface{}, error) {
	user, err := svc.UserDataService.GetById(svc.UserObject, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
