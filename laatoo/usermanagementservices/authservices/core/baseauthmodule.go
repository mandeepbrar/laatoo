package ginauth

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/storageutils"
	"net/http"
	"time"
)

var authTypes = make(map[string]AuthType)

func init() {
	RegisterModule("BaseAuth", &BaseAuth{})
}

type BaseAuth struct {
	app *App
}

func RegisterAuthType(name string, authType AuthType) {
	Logger.Info("BaseAuth: Registering Auth Provider ", name)
	authTypes[name] = authType
}

func (baseAuth *BaseAuth) Initialize(app *App) error {
	baseAuth.app = app
	return nil
}

func (baseAuth *BaseAuth) Serve() error {
	for _, authType := range authTypes {
		initializationErr := authType.InitializeType(
			func(ctx *gin.Context) { ///  auth start method starts
				err := authType.ValidateUser(ctx)
				if err != nil {
					err := fmt.Errorf("Internal Server Error starting authentication")
					ctx.AbortWithError(http.StatusInternalServerError, err)
					baseAuth.app.Logger.Error("BaseAuthProvider: Authentication error ", err)
				}
			}, ///  auth start method ends
			func(ctx *gin.Context) { ///  auth failed method starts
				delete(ctx.Keys, "User")
				ctx.Header(AuthToken, "")
				storageutils.FireEvent(&storageutils.Event{EVENT_LOGIN_FAIL, ctx})
				return
			}, ///  auth failed method ends
			func(ctx *gin.Context) { ///  auth successful method starts
				err := authType.CompleteAuthentication(ctx)
				if err != nil {
					err := fmt.Errorf("Internal Server Error completing authentication")
					ctx.AbortWithError(http.StatusInternalServerError, err)
					baseAuth.app.Logger.Error("BaseAuthProvider: Authentication error ", err)
				}
				userInt, exists := ctx.Get("User")
				if !exists {
					err := fmt.Errorf("Internal Server Error completing authentication")
					ctx.AbortWithError(http.StatusInternalServerError, err)
					baseAuth.app.Logger.Error("BaseAuthProvider: Authentication error ", err)
					return
				}
				user := userInt.(storageutils.Storable)
				token := jwt.New(jwt.SigningMethodHS256)
				token.Claims["UserId"] = user.GetId()
				token.Claims["IP"] = ctx.ClientIP()
				token.Claims["exp"] = time.Now().Add(time.Hour * 4).Unix()
				ctx.Set("JWT_Token", token)
				storageutils.FireEvent(&storageutils.Event{EVENT_JWT_TOKEN_PREPARED, ctx})
				tokenString, err := token.SignedString([]byte(JWTSecret))
				if err != nil {
					ctx.AbortWithError(http.StatusInternalServerError, err)
					baseAuth.app.Logger.Error("BaseAuthProvider: Authentication error ", err)
					return
				}
				ctx.Header(AuthToken, tokenString)
				storageutils.FireEvent(&storageutils.Event{EVENT_LOGIN_COMPLETE, ctx})
				ctx.JSON(http.StatusOK, user)
				return
			}) ///  auth logout method ends
		if initializationErr != nil {
			return initializationErr
		}
	}
	if baseAuth.app.Context.SingleApp {
		baseAuth.app.Router.GET(baseAuth.app.logoutPath, baseAuth.Logout)
	} else {
		logoutpath := fmt.Sprintf("%s%s", baseAuth.app.AppMountPath, baseAuth.app.logoutPath)
		baseAuth.app.Router.GET(logoutpath, baseAuth.Logout)
	}
	return nil
}
func (baseAuth *BaseAuth) Logout(ctx *gin.Context) {
	ctx.Header(AuthToken, "")
	delete(ctx.Keys, "User")
	storageutils.FireEvent(&storageutils.Event{EVENT_LOGOUT_COMPLETE, ctx})
}
