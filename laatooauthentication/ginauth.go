package authentication

import (
	//	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"laatooutils/commonobjects"
)

var (
	Logger = log.New()
)

func init() {
	Logger.Level = log.InfoLevel
}

// GinAuth contains a configuration and other details for running.
type GinAuth struct {
	Applications map[string]*App
	SingleApp    bool
}
type App struct {
	Configuration Config
	Mailer        Mailer
	logoutPath    string
	AppMountPath  string
	Logger        *log.Logger
	Router        *gin.Engine
	Storer        storageutils.Storer
	UserCreator   func() interface{}
	AppName       string
	Context       *GinAuth
}

var AuthToken string
var JWTSecret string

func New(singleApp bool) *GinAuth {
	ginAuth := &GinAuth{SingleApp: singleApp}
	if !singleApp {
		ginAuth.Applications = make(map[string]*App)
	}
	return ginAuth
}

// New makes a new instance of authboss with a default
// configuration.
func (ginAuth *GinAuth) Register(appName string, appConfig Config) (*App, error) {
	app := &App{AppName: appName, Logger: Logger, Configuration: appConfig, Context: ginAuth}
	if !ginAuth.SingleApp {
		ginAuth.Applications[appName] = app
	}
	var err error
	err = app.configUser()
	if err != nil {
		return nil, err
	}
	err = app.configStorer()
	if err != nil {
		return nil, err
	}
	err = app.configJWT()
	if err != nil {
		return nil, err
	}
	err = app.configPaths()
	if err != nil {
		return nil, err
	}
	if !ginAuth.SingleApp {
		err = app.configMountPath()
		if err != nil {
			return nil, err
		}
	}
	err = app.configRouter()
	if err != nil {
		return nil, err
	}
	err = app.configModules()
	if err != nil {
		return nil, err
	}
	return app, nil
}
