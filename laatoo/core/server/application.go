package server

import (
	"laatoo/core/engine/http"
	"laatoo/core/registry"
	"laatoo/core/security"
	"laatoo/sdk/auth"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/data"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	CONF_APP_ENGINE      = "engine"
	CONF_APP_ENGINE_NAME = "enginename"
	CONF_APP_SECURITY    = "security"
	CONF_APP_SETTINGS    = "settings"
	//header set by the service
	CONF_APP_COMMSVC  = "commsvc"
	CONF_APP_CACHESVC = "cachesvc"
	CONF_APP_USER     = "user_object"
	CONF_APP_ROLE     = "role_object"
	//header set by the service
	CONF_APP_AUTHHEADER = "auth_header"
	//secret key for jwt
	CONF_APP_JWTSECRETKEY = "jwtsecretkey"
)

type Application struct {
	//application config
	Config config.Config
	//store for service s in an application
	ServicesStore map[string]core.Service
	//store for service factory in an application
	ServiceFactoryStore map[string]core.ServiceFactory
	//store for service factory configuration
	ServiceFactoryConfig map[string]config.Config
	//pubsub service reference for publishing and receiving messages
	CommunicationService core.PubSub
	//header used for authentication tokens
	AuthHeader string
	//name of the application
	Name string
	//secret key for auth
	JWTSecret string
	//system user object
	SystemUser string
	//system role object
	SystemRole string
	//Name of the admin role
	AdminRole string
	//type of server application is hosted in
	ServerType string
	//caching service
	Cache data.Cache
	//security
	Security core.SecurityHandler
	/*	//application middleware
		middleware *Middleware
		//store for service factory middleware
		ServiceFactoryMiddleware map[string]*Middleware*/
	engineName string
	appEngine  core.Engine
	serverUser auth.User
	appContext core.ApplicationContext
}

//creates a new application
func newApplication(ctx *serverContext, appName string, appContext core.ApplicationContext, conf config.Config, serverType string) (*Application, error) {
	app := &Application{Name: appName, ServerType: serverType, appContext: appContext, Config: conf}
	ctx.application = app

	//default admin role
	app.AdminRole = core.CONF_SERVER_DEFAULT_ADMINROLE

	app.SystemUser = core.CONF_SERVER_DEFAULT_USEROBJ
	obj, _ := registry.CreateObject(ctx, app.SystemUser, nil)
	usr := obj.(auth.RbacUser)
	usr.SetId("__system__")
	usr.SetRoles([]string{"Admin"})
	app.serverUser = usr

	/////************TODO***********/
	//ctx.Set("Roles", []string{app.AdminRole})

	//store of all services
	app.ServicesStore = make(map[string]core.Service, 100)
	/*	//store of all factory middleware
		app.ServiceFactoryMiddleware = make(map[string]*Middleware, 30)*/
	//store of all services
	app.ServiceFactoryStore = make(map[string]core.ServiceFactory, 30)
	//store of all service configs
	app.ServiceFactoryConfig = make(map[string]config.Config, 30)
	//process application configuration
	if err := app.processConfiguration(ctx); err != nil {
		return nil, err
	}
	return app, nil
}

func (app *Application) processConfiguration(ctx *serverContext) error {

	secConfig, ok := app.Config.GetSubConfig(CONF_APP_SECURITY)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Config Name", CONF_APP_SECURITY)
	}
	app.processSecurityConfig(ctx, secConfig)

	/*	app.middleware = createMW(appconf, nil)*/

	if err := app.createServiceFactories(ctx); err != nil {
		return errors.RethrowError(ctx, CORE_APPLICATION_NOT_CREATED, err)
	}
	engineConf, ok := app.Config.GetSubConfig(CONF_APP_ENGINE)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Config Name", CONF_APP_ENGINE)
	}
	enginename, ok := engineConf.GetString(CONF_APP_ENGINE_NAME)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Config Name", CONF_APP_ENGINE_NAME)
	}
	app.engineName = enginename
	switch enginename {
	case core.CONF_ENGINE_HTTP:
		httpEngine, err := http.NewHttpEngine(ctx, engineConf)
		if err != nil {
			return errors.RethrowError(ctx, CORE_APPLICATION_NOT_CREATED, err)
		}
		app.appEngine = httpEngine //httpDelivery
	case core.CONF_ENGINE_TCP:
	default:
		return errors.ThrowError(ctx, CORE_APPLICATION_NOT_CREATED, "Wrong delivery mode", enginename)
	}
	return nil
}

func (app *Application) processSecurityConfig(ctx *serverContext, conf config.Config) {
	//check if user service name to be used has been provided, otherwise set default name
	roleObject, _ := conf.GetString(CONF_APP_ROLE)
	if len(roleObject) == 0 {
		roleObject = core.CONF_SERVER_DEFAULT_ROLEOBJ
	}
	app.SystemRole = roleObject

	//check if user service name to be used has been provided, otherwise set default name
	userObject, _ := conf.GetString(CONF_APP_USER)
	if len(userObject) == 0 {
		userObject = core.CONF_SERVER_DEFAULT_USEROBJ
	}
	app.SystemUser = userObject

	//check if jwt secret key has been provided, otherwise create a key from random numbers
	jwtSecret, _ := conf.GetString(CONF_APP_JWTSECRETKEY)
	if len(jwtSecret) > 0 {
		jwtSecret = core.CONF_SERVER_DEFAULT_JWTSECRET
	}
	app.JWTSecret = jwtSecret

	//check if auth header to be set has been provided, otherwise set default token
	authToken, _ := conf.GetString(CONF_APP_AUTHHEADER)
	if len(authToken) > 0 {
		authToken = core.CONF_SERVER_DEFAULT_AUTHHEADER
	}
	app.AuthHeader = authToken

	app.Security = security.NewLocalSecurityHandler(ctx, conf)
	return
}

//Initialize an applicaiton
func (app *Application) InitializeApplication(ctx *serverContext) error {
	err := app.createServices(ctx)
	if err != nil {
		return errors.RethrowError(ctx, CORE_APPLICATION_NOT_INITIALIZED, err)
	}

	commSvc, ok := app.Config.GetString(CONF_APP_COMMSVC)
	if ok {
		svcInt, ok := app.ServicesStore[commSvc]
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, "Communication Service", commSvc)
		}
		app.CommunicationService = svcInt.(core.PubSub)
	}

	cacheSvc, ok := app.Config.GetString(CONF_APP_CACHESVC)
	if ok {
		svcInt, ok := app.ServicesStore[cacheSvc]
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, "Cache Service", cacheSvc)
		}
		app.Cache = svcInt.(data.Cache)
	} else {
		log.Logger.Warn(ctx, "Cache service has not been initialized for the application", "App Name", app.Name)
	}

	err = app.Security.Initialize(ctx, nil)
	if err != nil {
		return errors.RethrowError(ctx, CORE_APPLICATION_NOT_INITIALIZED, err)
	}

	log.Logger.Info(ctx, "Initializing Services", "App Name", app.Name)
	err = app.initializeServices(ctx)
	if err != nil {
		return errors.RethrowError(ctx, CORE_APPLICATION_NOT_INITIALIZED, err)
	}
	log.Logger.Info(ctx, "Initializing Engine", "App Name", app.Name)
	err = app.appEngine.InitializeEngine(ctx)
	if err != nil {
		return errors.RethrowError(ctx, CORE_APPLICATION_NOT_INITIALIZED, err)
	}
	if app.appContext != nil {
		settingsConf, _ := app.Config.GetSubConfig(CONF_APP_SETTINGS)
		newctx := ctx.subCtx("Init App Context:"+app.Name, settingsConf, app)
		log.Logger.Info(newctx, "Initializing App Context")
		err = app.appContext.Initialize(newctx, settingsConf)
		if err != nil {
			return errors.RethrowError(newctx, CORE_APPLICATION_NOT_INITIALIZED, err)
		}
	}
	return nil
}

//Provides the service reference by alias
func (app *Application) GetService(ctx core.Context, alias string) (core.Service, error) {
	//get the service for the alias
	svc, ok := app.ServicesStore[alias]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, "Service Alias", alias)
	}
	return svc, nil
}

func (app *Application) GetVariable(variable core.ServerVariable) interface{} {
	switch variable {
	case core.JWTSECRETKEY:
		return app.JWTSecret
	case core.AUTHHEADER:
		return app.AuthHeader
	case core.ADMINROLE:
		return app.AdminRole
	case core.USER:
		return app.SystemUser
	case core.ROLE:
		return app.SystemRole
	}
	return nil
}

func (app *Application) GetConfig() config.Config {
	return app.Config
}

//start services
func (app *Application) StartApplication(ctx *serverContext) error {
	log.Logger.Info(ctx, "Starting Application", "App Name", app.Name)
	/*err := app.subscribeTopics(ctx)
	if err != nil {
		return errors.RethrowError(ctx, CORE_ERROR_PUBSUB_INITIALIZATION, err)
	}

	//load role permissions
	err = app.loadRolePermissions(ctx)
	if err != nil {
		return errors.RethrowError(ctx, CORE_ROLES_INIT_ERROR, err)
	}*/
	/*//get list of all the services
	svcs := app.ServicesStore.GetList()
	*/
	//iterate through all the services
	for alias, svcFactory := range app.ServiceFactoryStore {
		if err := svcFactory.StartServices(ctx); err != nil {
			return errors.RethrowError(ctx, CORE_ERROR_SERVICES_NOT_STARTED, err, "Service factory", alias)
		}
	}
	if app.appContext != nil {
		err := app.appContext.Start(ctx)
		if err != nil {
			return err
		}
	}
	go app.appEngine.StartEngine(ctx)

	return nil
}

func (app *Application) GetCache() data.Cache {
	return app.Cache
}
