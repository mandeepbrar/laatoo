package server

import (
	"laatoo/core/engine/http"
	"laatoo/core/security"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/data"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	CONF_ENV_ENGINE      = "engine"
	CONF_ENV_ENGINE_NAME = "enginename"
	CONF_ENV_SECURITY    = "security"
	//header set by the service
	CONF_ENV_COMMSVC  = "commsvc"
	CONF_ENV_CACHESVC = "cachesvc"
	CONF_ENV_USER     = "user_object"
	CONF_ENV_ROLE     = "role_object"
	//header set by the service
	CONF_ENV_AUTHHEADER = "auth_header"
	//secret key for jwt
	CONF_ENV_JWTSECRETKEY = "jwtsecretkey"
)

//Environment hosting an application
type Environment struct {
	//environment config
	Config config.Config
	//store for service s in an environment
	ServicesStore map[string]core.Service
	//store for service factory in an environment
	ServiceFactoryStore map[string]core.ServiceFactory
	//store for service factory configuration
	ServiceFactoryConfig map[string]config.Config
	//pubsub service reference for publishing and receiving messages
	CommunicationService core.PubSub
	//header used for authentication tokens
	AuthHeader string
	//name of the environment
	Name string
	//secret key for auth
	JWTSecret string
	//system user object
	SystemUser string
	//system role object
	SystemRole string
	//Name of the admin role
	AdminRole string
	//type of server environment is hosted in
	ServerType string
	//caching service
	Cache data.Cache
	//security
	Security core.SecurityHandler
	/*	//environment middleware
		middleware *Middleware
		//store for service factory middleware
		ServiceFactoryMiddleware map[string]*Middleware*/
	engineName string
	envEngine  core.Engine
}

//creates a new environment
func newEnvironment(ctx *serverContext, envName string, conf string, serverType string) (*Environment, error) {
	env := &Environment{Name: envName, ServerType: serverType}
	ctx.environment = env

	//default admin role
	env.AdminRole = core.CONF_SERVER_DEFAULT_ADMINROLE

	env.SystemUser = core.CONF_SERVER_DEFAULT_USEROBJ

	/////************TODO***********/
	//ctx.Set("Roles", []string{env.AdminRole})

	//store of all services
	env.ServicesStore = make(map[string]core.Service, 100)
	/*	//store of all factory middleware
		env.ServiceFactoryMiddleware = make(map[string]*Middleware, 30)*/
	//store of all services
	env.ServiceFactoryStore = make(map[string]core.ServiceFactory, 30)
	//store of all service configs
	env.ServiceFactoryConfig = make(map[string]config.Config, 30)
	//process environment configuration
	if err := env.processConfiguration(ctx, conf); err != nil {
		return nil, err
	}
	return env, nil
}

func (env *Environment) processConfiguration(ctx *serverContext, conf string) error {
	//read config for standalone
	envconf, err := config.NewConfigFromFile(conf)
	if err != nil {
		return errors.RethrowError(ctx, CORE_ENVIRONMENT_NOT_CREATED, err)
	}
	env.Config = envconf

	secConfig, ok := envconf.GetSubConfig(CONF_ENV_SECURITY)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Config Name", CONF_ENV_SECURITY)
	}
	env.processSecurityConfig(ctx, secConfig)

	/*	env.middleware = createMW(envconf, nil)*/

	if err = env.createServiceFactories(ctx); err != nil {
		return errors.RethrowError(ctx, CORE_ENVIRONMENT_NOT_CREATED, err)
	}
	engineConf, ok := envconf.GetSubConfig(CONF_ENV_ENGINE)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Config Name", CONF_ENV_ENGINE)
	}
	enginename, ok := engineConf.GetString(CONF_ENV_ENGINE_NAME)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Config Name", CONF_ENV_ENGINE_NAME)
	}
	env.engineName = enginename
	switch enginename {
	case core.CONF_ENGINE_HTTP:
		httpEngine, err := http.NewHttpEngine(ctx, engineConf)
		if err != nil {
			return errors.RethrowError(ctx, CORE_ENVIRONMENT_NOT_CREATED, err)
		}
		env.envEngine = httpEngine //httpDelivery
	case core.CONF_ENGINE_TCP:
	default:
		return errors.ThrowError(ctx, CORE_ENVIRONMENT_NOT_CREATED, "Wrong delivery mode", enginename)
	}
	return nil
}

func (env *Environment) processSecurityConfig(ctx *serverContext, conf config.Config) {
	//check if user service name to be used has been provided, otherwise set default name
	roleObject, _ := conf.GetString(CONF_ENV_ROLE)
	if len(roleObject) == 0 {
		roleObject = core.CONF_SERVER_DEFAULT_ROLEOBJ
	}
	env.SystemRole = roleObject

	//check if user service name to be used has been provided, otherwise set default name
	userObject, _ := conf.GetString(CONF_ENV_USER)
	if len(userObject) == 0 {
		userObject = core.CONF_SERVER_DEFAULT_USEROBJ
	}
	env.SystemUser = userObject

	//check if jwt secret key has been provided, otherwise create a key from random numbers
	jwtSecret, _ := conf.GetString(CONF_ENV_JWTSECRETKEY)
	if len(jwtSecret) > 0 {
		jwtSecret = core.CONF_SERVER_DEFAULT_JWTSECRET
	}
	env.JWTSecret = jwtSecret

	//check if auth header to be set has been provided, otherwise set default token
	authToken, _ := conf.GetString(CONF_ENV_AUTHHEADER)
	if len(authToken) > 0 {
		authToken = core.CONF_SERVER_DEFAULT_AUTHHEADER
	}
	env.AuthHeader = authToken

	env.Security = security.NewLocalSecurityHandler(ctx)
	return
}

//Initialize an environment
func (env *Environment) InitializeEnvironment(ctx *serverContext) error {
	err := env.createServices(ctx)
	if err != nil {
		return errors.RethrowError(ctx, CORE_ENVIRONMENT_NOT_INITIALIZED, err)
	}

	commSvc, ok := env.Config.GetString(CONF_ENV_COMMSVC)
	if ok {
		svcInt, ok := env.ServicesStore[commSvc]
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, "Communication Service", commSvc)
		}
		env.CommunicationService = svcInt.(core.PubSub)
	}

	cacheSvc, ok := env.Config.GetString(CONF_ENV_CACHESVC)
	if ok {
		svcInt, ok := env.ServicesStore[cacheSvc]
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, "Cache Service", cacheSvc)
		}
		env.Cache = svcInt.(data.Cache)
	} else {
		log.Logger.Warn(ctx, "Cache service has not been initialized for the environment", "Env Name", env.Name)
	}

	log.Logger.Info(ctx, "Initializing Services", "Env Name", env.Name)
	err = env.initializeServices(ctx)
	if err != nil {
		return errors.RethrowError(ctx, CORE_ENVIRONMENT_NOT_INITIALIZED, err)
	}
	log.Logger.Info(ctx, "Initializing Engine", "Env Name", env.Name)
	err = env.envEngine.InitializeEngine(ctx)
	if err != nil {
		return errors.RethrowError(ctx, CORE_ENVIRONMENT_NOT_INITIALIZED, err)
	}
	return nil
}

//Provides the service reference by alias
func (env *Environment) GetService(ctx core.Context, alias string) (core.Service, error) {
	//get the service for the alias
	svc, ok := env.ServicesStore[alias]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, "Service Alias", alias)
	}
	return svc, nil
}

func (env *Environment) GetVariable(variable core.ServerVariable) interface{} {
	switch variable {
	case core.JWTSECRETKEY:
		return env.JWTSecret
	case core.AUTHHEADER:
		return env.AuthHeader
	case core.ADMINROLE:
		return env.AdminRole
	case core.USER:
		return env.SystemUser
	case core.ROLE:
		return env.SystemRole
	}
	return nil
}

func (env *Environment) GetConfig() config.Config {
	return env.Config
}

//start services
func (env *Environment) StartEnvironment(ctx *serverContext) error {
	log.Logger.Info(ctx, "Starting environment", "Env Name", env.Name)

	/*err := env.subscribeTopics(ctx)
	if err != nil {
		return errors.RethrowError(ctx, CORE_ERROR_PUBSUB_INITIALIZATION, err)
	}

	//load role permissions
	err = env.loadRolePermissions(ctx)
	if err != nil {
		return errors.RethrowError(ctx, CORE_ROLES_INIT_ERROR, err)
	}*/
	/*//get list of all the services
	svcs := env.ServicesStore.GetList()
	*/
	//iterate through all the services
	for alias, svcFactory := range env.ServiceFactoryStore {
		if err := svcFactory.StartServices(ctx); err != nil {
			return errors.RethrowError(ctx, CORE_ERROR_SERVICES_NOT_STARTED, err, "Service factory", alias)
		}
	}

	go env.envEngine.StartEngine(ctx)

	return nil
}

func (env *Environment) GetCache() data.Cache {
	return env.Cache
}
