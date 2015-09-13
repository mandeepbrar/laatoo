package laatoocore

import (
	"github.com/dghubble/sling"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"laatoosdk/auth"
	"laatoosdk/config"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"laatoosdk/utils"
	"net/http"
)

const (
	CONF_ENV_SERVICES    = "services"
	CONF_ENV_SERVICENAME = "servicename"
	CONF_ENV_USER        = "settings.user_object"
	CONF_ENV_ROLE        = "settings.role_object"
	//header set by the service
	CONF_ENV_AUTHHEADER = "settings.auth_header"
	CONF_ENV_COMMSVC    = "settings.commsvc"
	//secret key for jwt
	CONF_ENV_JWTSECRETKEY   = "settings.jwtsecretkey"
	DEFAULT_USER            = "User"
	DEFAULT_ROLE            = "Role"
	CONF_ENV_ROUTER         = "router"
	CONF_ENV_CONTEXT        = "context"
	CONF_SERVICE_BINDPATH   = "path"
	CONF_SERVICE_SERVERTYPE = "servertype"
	CONF_SERVICE_AUTHBYPASS = "bypassauth"
	CONF_AUTH_MODE          = "settings.authorization.mode"
	CONF_AUTH_MODE_LOCAL    = "local"
	CONF_AUTH_MODE_REMOTE   = "remote"
	CONF_API_AUTH           = "settings.authorization.apiauth"
	CONF_ROLES_API          = "settings.authorization.rolesapi"
	CONF_API_PUBKEY         = "settings.authorization.pubkey"
	CONF_API_DOMAIN         = "settings.authorization.domain"
)

//Environment hosting an application
type Environment struct {
	Router        *echo.Group
	Config        config.Config
	ServicesStore *utils.MemoryStorer
	pubSub        service.PubSub
	authHeader    string
	Name          string
}

//creates a new environment
func newEnvironment(envName string, conf string, router *echo.Group, serverType string) (*Environment, error) {
	env := &Environment{Name: envName, Router: router}
	env.ServicesStore = utils.NewMemoryStorer()
	//read config for standalone
	env.Config = config.NewConfigFromFile(conf)
	//create all services in the environment

	if err := env.createServices(serverType); err != nil {
		return nil, errors.RethrowError(CORE_ENVIRONMENT_NOT_CREATED, err, envName)
	}
	return env, nil
}

func (env *Environment) createServices(serverType string) error {

	//check if user service name to be used has been provided, otherwise set default name
	roleObject := env.Config.GetString(CONF_ENV_ROLE)
	if len(roleObject) == 0 {
		roleObject = DEFAULT_ROLE
	}
	SystemRole = roleObject

	//check if user service name to be used has been provided, otherwise set default name
	userObject := env.Config.GetString(CONF_ENV_USER)
	if len(userObject) == 0 {
		userObject = DEFAULT_USER
	}
	SystemUser = userObject
	auserInt, err := CreateEmptyObject(userObject)
	if err != nil {
		return err
	}
	anonymousUser := auserInt.(auth.RbacUser)
	anonymousUser.AddRole("Anonymous")

	jwtSecret := utils.RandomString(15)
	env.authHeader = "Auth-Token"

	//check if jwt secret key has been provided, otherwise create a key from random numbers
	jwtSecretInt := env.Config.GetString(CONF_ENV_JWTSECRETKEY)
	if len(jwtSecret) > 0 {
		jwtSecret = jwtSecretInt
	}

	//check if auth header to be set has been provided, otherwise set default token
	authTokenInt := env.Config.GetString(CONF_ENV_AUTHHEADER)
	if len(authTokenInt) > 0 {
		env.authHeader = authTokenInt
	}

	//get a map of all the services
	svcs := env.Config.GetMap(CONF_ENV_SERVICES)
	for alias, val := range svcs {
		//get the config for the service with given alias
		serviceConfig := val.(map[string]interface{})
		//get the service name to be created for the alias
		log.Logger.Infof("Creating service %s", alias)

		svcName, ok := serviceConfig[CONF_ENV_SERVICENAME].(string)
		if !ok {
			return errors.ThrowError(CORE_ERROR_MISSING_SERVICE_NAME, alias)
		}

		svcServerType, ok := serviceConfig[CONF_SERVICE_SERVERTYPE]

		if ok {
			if serverType != svcServerType.(string) {
				continue
			}
		}

		svcBindPath, ok := serviceConfig[CONF_SERVICE_BINDPATH]
		if ok {
			//router to be passed in the configuration
			router := env.Router.Group(svcBindPath.(string))
			router.Use(func(ctx *echo.Context) error {
				ctx.Set(CONF_ENV_CONTEXT, env)
				return nil
			})
			bypassauth := false
			//authentication required by default unless explicitly turned off
			bypassauthInt, ok := serviceConfig[CONF_SERVICE_AUTHBYPASS]
			if ok {
				bypassauth = (bypassauthInt == "true")
			}
			if !bypassauth {
				router.Use(func(ctx *echo.Context) error {
					headerVal := ctx.Request().Header.Get(env.authHeader)

					if headerVal != "" {
						token, err := jwt.Parse(headerVal, func(token *jwt.Token) (interface{}, error) {
							// Don't forget to validate the alg is what you expect:
							if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
								return nil, errors.ThrowError(AUTH_ERROR_WRONG_SIGNING_METHOD)
							}
							return []byte(jwtSecret), nil
						})
						if err == nil && token.Valid {
							userInt, err := CreateEmptyObject(userObject)
							if err != nil {
								return errors.RethrowHttpError(AUTH_ERROR_WRONG_SIGNING_METHOD, ctx, err)
							}
							user, ok := userInt.(auth.RbacUser)
							if !ok {
								return errors.ThrowHttpError(AUTH_ERROR_USEROBJECT_NOT_CREATED, ctx)
							}
							user.LoadJWTClaims(token)
							user.SetId(token.Claims["UserId"].(string))
							ctx.Set("User", userInt)
							roles, _ := user.GetRoles()
							ctx.Set("Roles", roles)
							ctx.Set("JWT_Token", token)
							utils.FireEvent(&utils.Event{EVENT_AUTHSERVICE_AUTH_COMPLETE, ctx})
							return nil
						} else {
							if token == nil || !token.Valid {
								return errors.RethrowHttpError(AUTH_ERROR_INVALID_TOKEN, ctx, err)
							}
							return err
						}
					} else {
						ctx.Set("User", anonymousUser)
					}
					return nil
				})
			}
			serviceConfig[CONF_ENV_ROUTER] = router
			serviceConfig[CONF_ENV_JWTSECRETKEY] = jwtSecret
			serviceConfig[CONF_ENV_AUTHHEADER] = env.authHeader
		}

		serviceConfig[CONF_ENV_CONTEXT] = env

		//get the service with a given name alias and config
		svcInt, err := CreateObject(svcName, serviceConfig)
		if err != nil {
			return errors.RethrowError(CORE_ERROR_SERVICE_CREATION, err, alias)
		}

		//put the created service in the store
		svc := svcInt.(service.Service)
		env.ServicesStore.PutObject(alias, svc)
	}
	return nil
}

//Initialize an environment
func (env *Environment) InitializeEnvironment() error {
	commSvc := env.Config.GetString(CONF_ENV_COMMSVC)
	if len(commSvc) > 0 {
		svcInt, err := env.ServicesStore.GetObject(commSvc)
		if err != nil {
			return err
		}
		env.pubSub = svcInt.(service.PubSub)
	}

	//go through list of all the services
	svcs := env.ServicesStore.GetList()
	//iterate through all the services
	for _, svcInt := range svcs {
		svc := svcInt.(service.Service)
		//initialize service
		err := svc.Initialize(env)
		if err != nil {
			return errors.RethrowError(CORE_ERROR_SERVICE_INITIALIZATION, err, svc.GetName())
		}
	}
	return nil
}

//Provides the service reference by alias
func (env *Environment) GetService(alias string) (service.Service, error) {
	svcInt, err := env.ServicesStore.GetObject(alias)
	if err != nil {
		return nil, err
	}
	svc, ok := svcInt.(service.Service)
	if !ok {
		return nil, errors.RethrowError(CORE_ERROR_SERVICE_NOT_FOUND, err, alias)
	}
	return svc, nil
}

func (env *Environment) GetAllServices() []interface{} {
	return env.ServicesStore.GetList()
}

//creates a named object if the factory has been registered with environment
func (env *Environment) CreateObject(objName string, confData map[string]interface{}) (interface{}, error) {
	return CreateObject(objName, confData)
}

func (env *Environment) CreateEmptyObject(objName string) (interface{}, error) {
	return CreateEmptyObject(objName)
}

func (env *Environment) CreateCollection(objName string) (interface{}, error) {
	return CreateCollection(objName)
}

func (env *Environment) GetConfig() config.Config {
	return env.Config
}

//start services
func (env *Environment) StartEnvironment() error {
	log.Logger.Infof("Starting environment %s", env.Name)
	err := env.loadRemoteRolePermissions()
	if err != nil {
		return errors.RethrowError(CORE_ROLES_INIT_ERROR, err)
	}
	err = env.subscribeTopics()
	if err != nil {
		return errors.RethrowError(CORE_ERROR_PUBSUB_INITIALIZATION, err)
	}
	//go through list of all the services
	svcs := env.ServicesStore.GetList()
	//iterate through all the services
	for _, svcInt := range svcs {
		svc := svcInt.(service.Service)
		//start service
		if err := svc.Serve(); err != nil {
			return errors.RethrowError(CORE_ERROR_SERVICE_NOT_STARTED, err, svc.GetName())
		}
	}
	return nil
}

func (env *Environment) loadRemoteRolePermissions() error {
	mode := env.Config.GetString(CONF_AUTH_MODE)
	if mode == CONF_AUTH_MODE_REMOTE {
		apiauth := env.Config.GetString(CONF_API_AUTH)
		if len(apiauth) == 0 {
			errors.ThrowError(AUTH_MISSING_API)
		}
		pubkey := env.Config.GetString(CONF_API_PUBKEY)
		domain := env.Config.GetString(CONF_API_DOMAIN)
		log.Logger.Infof("pubkey", pubkey)
		key, err := EncryptWithKey(pubkey, domain)
		if err != nil {
			return err
		}
		form := &KeyAuth{Key: key}
		req, err := sling.New().Post(apiauth).BodyJSON(form).Request()
		if err != nil {
			return err
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			errors.ThrowError(AUTH_APISEC_NOTALLOWED)
		} else {
			token := resp.Header.Get(env.authHeader)
			rolesurl := env.Config.GetString(CONF_ROLES_API)
			if len(rolesurl) == 0 {
				return errors.ThrowError(CORE_ROLESAPI_NOT_FOUND)
			}
			roles, err := CreateCollection(SystemRole)
			if err != nil {
				return err
			}
			base := sling.New().Set(env.authHeader, token)
			//req, err := base.New().Get("gophergram/list").Request()
			resp, err = base.New().Get(rolesurl).ReceiveSuccess(roles)
			if err != nil {
				return err
			}
			if resp.StatusCode != 200 {
				return errors.ThrowError(CORE_ROLESAPI_NOT_FOUND)
			}
			RegisterRoles(roles)
		}

	}
	return nil
}
