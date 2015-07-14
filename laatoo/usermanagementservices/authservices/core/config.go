package ginauth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/storageutils"
)

type Config map[string]interface{}

const (
	ROUTER       = "Router"
	AUTHTOKEN    = "AuthToken"
	JWTSECRETKEY = "JWTSecretKey"
	USERSTORER   = "UserStorer"
	USERCREATOR  = "UserCreator"
	LOGOUTPATH   = "LogoutPath"
	APPMOUNTPATH = "AppMountPath"
)

func (app *App) configRouter() error {
	router, ok := app.Configuration[ROUTER]
	if ok {
		app.Router = router.(*gin.Engine)
	} else {
		return fmt.Errorf("GinAuth: Gin router not provided")
	}
	return nil
}

func (app *App) configUser() error {
	userCreatorInt, ok := app.Configuration[USERCREATOR]
	if ok {
		app.UserCreator = userCreatorInt.(func() interface{})
	} else {
		return fmt.Errorf("GinAuth: User Creation Function not provided")
	}
	return nil
}

func (app *App) configJWT() error {
	jwtSecretInt, ok := app.Configuration[JWTSECRETKEY]
	if ok {
		JWTSecret = jwtSecretInt.(string)
	} else {
		JWTSecret = storageutils.RandomString(15)
	}

	authTokenInt, tokenok := app.Configuration[AUTHTOKEN]
	if tokenok {
		AuthToken = authTokenInt.(string)
	} else {
		AuthToken = "Auth-Token"
	}
	return nil
}

func (app *App) configStorer() error {
	storer, ok := app.Configuration[USERSTORER]
	if ok {
		app.Storer = storer.(storageutils.Storer)
	} else {
		app.Storer = storageutils.NewMemoryStorer("User")
	}
	app.Logger.Info("Using storer %s", app.Storer.GetName())
	return nil
}
func (app *App) configPaths() error {
	logoutpath, ok := app.Configuration[LOGOUTPATH]
	fmt.Println("configuring path")
	if ok {
		app.logoutPath = logoutpath.(string)
	} else {
		app.logoutPath = "/logout"
	}
	return nil
}

func (app *App) configMountPath() error {
	appmountpath, ok := app.Configuration[APPMOUNTPATH]
	if ok {
		app.AppMountPath = appmountpath.(string)
	} else {
		return fmt.Errorf("Mount Path not provided for app %s", app.AppName)
	}
	return nil
}
