package laatooauthentication

import (
	jwt "github.com/dgrijalva/jwt-go"
	"laatoocore"
	"laatoosdk/auth"
	"laatoosdk/core"
	"laatoosdk/data"
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
	InitializeType(ctx core.Context, authStart core.HandlerFunc, authCallback core.HandlerFunc) error
	//Called to validate the user by providing context
	ValidateUser(core.Context) error
	//Completes authentication
	CompleteAuthentication(core.Context) error
}

//setup local authentication
func (svc *SecurityService) SetupLocalAuth(ctx core.Context, conf map[string]interface{}) error {
	//create local authentication type
	localAuthType, err := NewLocalAuth(ctx, conf, svc)
	if err != nil {
		return err
	}
	//initialize local authentication
	err = svc.initializeAuthType(ctx, localAuthType)
	if err != nil {
		return err
	}
	return nil
}

//setup api authentication
func (svc *SecurityService) SetupKeyAuth(ctx core.Context, conf map[string]interface{}) error {
	//create local authentication type
	keyAuthType, err := NewKeyAuth(ctx, conf, svc)
	if err != nil {
		return err
	}
	//initialize local authentication
	err = svc.initializeAuthType(ctx, keyAuthType)
	if err != nil {
		return err
	}
	return nil
}

//setup local authentication
func (svc *SecurityService) SetupOAuth(ctx core.Context, conf map[string]interface{}) error {
	oAuthType, err := NewOAuth(ctx, conf, svc)
	if err != nil {
		return err
	}
	//initialize local authentication
	err = svc.initializeAuthType(ctx, oAuthType)
	if err != nil {
		return err
	}
	return nil
}

//The service starts serving when this method is called
func (svc *SecurityService) initializeAuthType(ctx core.Context, authType AuthType) error {
	//initialize auth type
	initializationErr := authType.InitializeType(ctx,
		func(ctx core.Context) error { ///  auth start method starts
			log.Logger.Trace(ctx, LOGGING_CONTEXT, "Validating user")
			err := authType.ValidateUser(ctx)
			if err != nil {
				return errors.RethrowError(ctx, AUTH_ERROR_USER_VALIDATION_FAILED, err)
			}
			return nil
		}, ///  auth start method ends
		func(ctx core.Context) error { ///  auth callback method starts
			err := authType.CompleteAuthentication(ctx)
			if err != nil {
				return errors.RethrowError(ctx, AUTH_ERROR_AUTH_COMPLETION_FAILED, err)
			}
			userInt := ctx.Get("User")
			if userInt == nil {
				return errors.ThrowError(ctx, AUTH_ERROR_INTERNAL_SERVER_ERROR_AUTH)
			}
			user := userInt.(auth.User)
			token := jwt.New(jwt.SigningMethodHS256)
			rbac, ok := user.(auth.RbacUser)
			if ok {
				svc.LoadPermissions(ctx, rbac, svc.UserDataService)
			}
			user.SetJWTClaims(token)
			token.Claims["UserId"] = user.GetId()
			token.Claims["AuthTypeName"] = authType.GetName()
			//token.Claims["IP"] = ctx.ClientIP()
			token.Claims["exp"] = time.Now().Add(time.Hour * 4).Unix()
			ctx.Set("JWT_Token", token)
			tokenString, err := token.SignedString([]byte(svc.JWTSecret))
			if err != nil {
				return errors.RethrowError(ctx, AUTH_ERROR_JWT_CREATION, err)
			}
			ctx.SetHeader(svc.AuthHeader, tokenString)
			log.Logger.Trace(ctx, LOGGING_CONTEXT, "Sending Token", svc.AuthHeader, tokenString)

			utils.FireEvent(&utils.Event{EVENT_AUTHSERVICE_LOGIN_COMPLETE, ctx})
			ctx.JSON(http.StatusOK, user)
			return nil
		}) ///  auth logout method ends
	if initializationErr != nil {
		return errors.RethrowError(ctx, AUTH_ERROR_INITIALIZING_TYPE, initializationErr)
	}
	svc.Router.Get(ctx, svc.LogoutPath, map[string]interface{}{}, svc.Logout)
	return nil
}

func (svc *SecurityService) Logout(ctx core.Context) error {
	ctx.SetHeader(svc.AuthHeader, "")
	ctx.Set("User", nil)
	utils.FireEvent(&utils.Event{EVENT_AUTHSERVICE_LOGOUT_COMPLETE, ctx})
	return nil
}

func (svc *SecurityService) CreateUser(ctx core.Context) (interface{}, error) {
	return laatoocore.CreateObject(ctx, svc.UserObject, nil)
}
func (svc *SecurityService) GetUserById(ctx core.Context, id string) (interface{}, error) {
	user, err := svc.UserDataService.GetById(ctx, svc.UserObject, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (svc *SecurityService) LoadPermissions(ctx core.Context, usr auth.RbacUser, roleStorer data.DataService) error {
	roles, _ := usr.GetRoles()
	permissions := utils.NewStringSet([]string{})
	for _, k := range roles {
		if k == svc.AdminRole {
			usr.SetPermissions(svc.Permissions)
			return nil
		}
		roleInt, err := roleStorer.GetById(ctx, laatoocore.DEFAULT_ROLE, k)
		if err == nil && roleInt != nil {
			role := roleInt.(auth.Role)
			permissions.Append(role.GetPermissions())
		}
	}
	usr.SetPermissions(permissions.Values())
	return nil
}
