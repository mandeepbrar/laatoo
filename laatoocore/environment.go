package laatoocore

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/rs/cors"
	"io/ioutil"
	"laatoosdk/auth"
	"laatoosdk/config"
	"laatoosdk/context"
	"laatoosdk/data"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"laatoosdk/utils"
	"net/http"
	"reflect"
)

const (
	CONF_ENV_SERVICES    = "services"
	CONF_ENV_SERVICENAME = "servicename"
	CONF_ENVPATH         = "settings.path"
	CONF_ENV_USER        = "settings.user_object"
	CONF_ENV_CORSHOSTS   = "settings.corshosts"
	CONF_ENV_ROLE        = "settings.role_object"
	//header set by the service
	CONF_ENV_AUTHHEADER = "settings.auth_header"
	CONF_ENV_COMMSVC    = "settings.commsvc"
	CONF_ENV_CACHESVC   = "settings.cachesvc"
	//secret key for jwt
	CONF_ENV_JWTSECRETKEY   = "settings.jwtsecretkey"
	DEFAULT_USER            = "User"
	DEFAULT_ROLE            = "Role"
	CONF_ENV_ADMINROLE      = "AdminRole"
	CONF_ENV_ROUTER         = "router"
	CONF_ENV_CONTEXT        = "context"
	CONF_SERVICE_BINDPATH   = "path"
	CONF_SERVICE_SERVERTYPE = "servertype"
	CONF_SERVICE_AUTHBYPASS = "bypassauth"
	CONF_SERVICE_USECORS    = "usecors"
	CONF_AUTH_MODE          = "settings.authorization.mode"
	CONF_AUTH_MODE_LOCAL    = "local"
	CONF_AUTH_MODE_REMOTE   = "remote"
	CONF_API_AUTH           = "settings.authorization.apiauth"
	CONF_ROLES_API          = "settings.authorization.rolesapi"
	CONF_SECURITY_SVC       = "settings.authorization.securitysvc"
	CONF_PERMISSIONS_API    = "settings.authorization.permissionsapi"
	CONF_API_PUBKEY         = "settings.authorization.pubkey"
	CONF_API_DOMAIN         = "settings.authorization.domain"
)

//Environment hosting an application
type Environment struct {
	//router used by the environment
	Router *echo.Group
	//environment config
	Config config.Config
	//store for services in an environment
	ServicesStore map[string]service.Service
	//pubsub service reference for publishing and receiving messages
	CommunicationService service.PubSub
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
	//permissions set for the environment
	Permissions utils.StringSet
	//permissions assigned to a role
	RolePermissions map[string]bool
	//type of server environment is hosted in
	ServerType string
	//caching service
	Cache data.Cache
}

//creates a new environment
func newEnvironment(ctx *echo.Context, envName string, conf string, router *echo.Echo, serverType string) (*Environment, error) {
	env := &Environment{Name: envName, ServerType: serverType}
	ctx.Set(CONF_ENV_CONTEXT, env)
	//default admin role
	env.AdminRole = "Admin"

	ctx.Set("Roles", []string{env.AdminRole})
	//construct permissions set
	env.Permissions = utils.NewStringSet([]string{})
	//map containing roles and permissions
	env.RolePermissions = make(map[string]bool)
	//store of all services
	env.ServicesStore = make(map[string]service.Service, 30)
	//read config for standalone
	env.Config = config.NewConfigFromFile(conf)

	envPath := env.Config.GetString(CONF_ENVPATH)

	env.AuthHeader = "X-Auth-Token"
	//check if auth header to be set has been provided, otherwise set default token
	authTokenInt := env.Config.GetString(CONF_ENV_AUTHHEADER)
	if len(authTokenInt) > 0 {
		env.AuthHeader = authTokenInt
	}

	env.Router = router.Group(envPath)
	//create all services in the environment
	if err := env.createServices(ctx); err != nil {
		return nil, errors.RethrowError(ctx, CORE_ENVIRONMENT_NOT_CREATED, err, envName)
	}
	return env, nil
}

//create services within an environment
func (env *Environment) createServices(ctx *echo.Context) error {

	//check if user service name to be used has been provided, otherwise set default name
	roleObject := env.Config.GetString(CONF_ENV_ROLE)
	if len(roleObject) == 0 {
		roleObject = DEFAULT_ROLE
	}
	env.SystemRole = roleObject

	//check if user service name to be used has been provided, otherwise set default name
	userObject := env.Config.GetString(CONF_ENV_USER)
	if len(userObject) == 0 {
		userObject = DEFAULT_USER
	}
	env.SystemUser = userObject

	env.JWTSecret = utils.RandomString(15)

	//check if jwt secret key has been provided, otherwise create a key from random numbers
	jwtSecretInt := env.Config.GetString(CONF_ENV_JWTSECRETKEY)
	if len(jwtSecretInt) > 0 {
		env.JWTSecret = jwtSecretInt
	}

	//get a map of all the services
	svcs := env.Config.GetMap(CONF_ENV_SERVICES)
	for alias, val := range svcs {
		//create the service
		svc, err := env.createService(ctx, alias, val)
		if err != nil {
			return err
		}
		if svc != nil {
			//add the service to the environment
			env.ServicesStore[alias] = svc
		}
	}
	return nil
}

//this method creates a service with a given configuration
func (env *Environment) createService(ctx *echo.Context, alias string, conf interface{}) (service.Service, error) {

	//get the config for the service with given alias
	serviceConfig := conf.(map[string]interface{})
	//get the service name to be created for the alias
	log.Logger.Info(ctx, "core.env.createservice", "Creating service ", "service alias", alias)
	//get the name of the service to be constructed for the alias
	//services can be created multiple times
	svcName, ok := serviceConfig[CONF_ENV_SERVICENAME].(string)
	if !ok {
		return nil, errors.ThrowError(ctx, CORE_ERROR_MISSING_SERVICE_NAME, "service alias", alias)
	}

	//get the type of the server for the service
	svcServerType, ok := serviceConfig[CONF_SERVICE_SERVERTYPE]
	//if the type of the server on which service is hosted
	//is not suitable, skip the service
	if ok {
		if env.ServerType != svcServerType.(string) {
			return nil, nil
		}
	}

	//get the bind path for the service
	svcBindPath, ok := serviceConfig[CONF_SERVICE_BINDPATH]
	if ok {
		//router to be passed in the configuration
		router := env.Router.Group(svcBindPath.(string))
		_, ok := serviceConfig[CONF_SERVICE_USECORS]
		if ok {
			corsHostsInt := env.Config.GetArray(CONF_ENV_CORSHOSTS)
			if corsHostsInt != nil {
				allowedOrigins := make([]string, len(corsHostsInt))
				i := 0
				for _, k := range corsHostsInt {
					allowedOrigins[i] = k.(string)
					i++
				}
				corsMw := cors.New(cors.Options{
					AllowedOrigins:   allowedOrigins,
					AllowedHeaders:   []string{"*"},
					ExposedHeaders:   []string{env.AuthHeader},
					AllowCredentials: true,
				}).Handler
				log.Logger.Info(ctx, "core.env", "CORS enabled for hosts ", "hosts", allowedOrigins)
				router.Use(corsMw)
			}
		}

		//provide environment context to every request using middleware
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
			//use authentication middleware for the service unless explicitly bypassed
			env.setupAuthMiddleware(ctx, router)
		}
		serviceConfig[CONF_ENV_ROUTER] = router
	}

	//get the service with a given name alias and config
	svcInt, err := CreateObject(ctx, svcName, serviceConfig)
	if err != nil {
		return nil, errors.RethrowError(ctx, CORE_ERROR_SERVICE_CREATION, err, "Alias", alias)
	}
	//put the created service in the store
	svc := svcInt.(service.Service)
	return svc, nil
}

//Initialize an environment
func (env *Environment) InitializeEnvironment(ctx *echo.Context) error {

	commSvc := env.Config.GetString(CONF_ENV_COMMSVC)
	if len(commSvc) > 0 {
		svcInt, ok := env.ServicesStore[commSvc]
		if !ok {
			return errors.ThrowError(ctx, CORE_ERROR_SERVICE_NOT_FOUND, "Communication Service", commSvc)
		}
		env.CommunicationService = svcInt.(service.PubSub)
	}

	cacheSvc := env.Config.GetString(CONF_ENV_CACHESVC)
	if len(cacheSvc) > 0 {
		svcInt, ok := env.ServicesStore[cacheSvc]
		if !ok {
			return errors.ThrowError(ctx, CORE_ERROR_SERVICE_NOT_FOUND, "Cache Service", cacheSvc)
		}
		env.Cache = svcInt.(data.Cache)
	}

	//get list of all the services
	/*svcs := env.ServicesStore.GetList()*/
	//iterate through all the servicesand initialize them
	for alias, svc := range env.ServicesStore {
		//initialize service
		err := svc.Initialize(ctx)
		if err != nil {
			return errors.RethrowError(ctx, CORE_ERROR_SERVICE_INITIALIZATION, err, "Service Alias", alias)
		}
	}
	return nil
}

//Provides the service reference by alias
func (env *Environment) GetService(ctx *echo.Context, alias string) (service.Service, error) {
	//get the service for the alias
	svcInt, ok := env.ServicesStore[alias]
	if !ok {
		return nil, errors.ThrowError(ctx, CORE_ERROR_SERVICE_NOT_FOUND, "Service Alias", alias)
	}
	svc, _ := svcInt.(service.Service)
	return svc, nil
}

func (env *Environment) GetVariable(variable string) interface{} {
	switch variable {
	case CONF_ENV_JWTSECRETKEY:
		return env.JWTSecret
	case CONF_ENV_AUTHHEADER:
		return env.AuthHeader
	case CONF_ENV_ADMINROLE:
		return env.AdminRole
	case CONF_ENV_USER:
		return env.SystemUser
	case CONF_ENV_ROLE:
		return env.SystemRole
	}
	return nil
}

func (env *Environment) GetConfig() config.Config {
	return env.Config
}

//start services
func (env *Environment) StartEnvironment(ctx *echo.Context) error {
	log.Logger.Info(ctx, "core.env", "Starting environment", "Env Name", env.Name)

	err := env.subscribeTopics(ctx)
	if err != nil {
		return errors.RethrowError(ctx, CORE_ERROR_PUBSUB_INITIALIZATION, err)
	}

	//load role permissions
	err = env.loadRolePermissions(ctx)
	if err != nil {
		return errors.RethrowError(ctx, CORE_ROLES_INIT_ERROR, err)
	}
	/*//get list of all the services
	svcs := env.ServicesStore.GetList()
	*/
	//iterate through all the services
	for alias, svc := range env.ServicesStore {
		//svc := svcInt.(service.Service)
		//start service
		if err := svc.Serve(ctx); err != nil {
			return errors.RethrowError(ctx, CORE_ERROR_SERVICE_NOT_STARTED, err, "Service Alias", alias)
		}
	}
	return nil
}

func (env *Environment) GetCache() data.Cache {
	return env.Cache
}

//load role permissions if needed from another environment
func (env *Environment) loadRolePermissions(ctx *echo.Context) error {
	//check the authenticatino mode
	mode := env.Config.GetString(CONF_AUTH_MODE)
	if mode == CONF_AUTH_MODE_REMOTE {
		//load permissions from remote system
		apiauth := env.Config.GetString(CONF_API_AUTH)
		if len(apiauth) == 0 {
			errors.ThrowError(ctx, AUTH_MISSING_API)
		}
		//authenticate to the remote system using public key
		pubkey := env.Config.GetString(CONF_API_PUBKEY)
		domain := env.Config.GetString(CONF_API_DOMAIN)
		//encrypt system domain and send
		key, err := EncryptWithKey(pubkey, domain)
		if err != nil {
			return err
		}
		client := context.HttpClient(ctx)
		form := &KeyAuth{Key: key}
		load, _ := json.Marshal(form)
		resp, err := client.Post(apiauth, "application/json", bytes.NewBuffer(load))
		if err != nil {
			log.Logger.Trace(ctx, "core.env.remoteroles", "Error in key auth", "err", err)
			return err
		}
		log.Logger.Trace(ctx, "core.env.remoteroles", "Got Response for api key", "Response", resp.StatusCode)
		if resp.StatusCode != 200 {
			//if the remote system did not allow auth
			errors.ThrowError(ctx, AUTH_APISEC_NOTALLOWED)
		} else {

			//get token from remote system
			token := resp.Header.Get(env.AuthHeader)
			log.Logger.Trace(ctx, "core.env.remoteroles", "Auth token for api key", "Token", token)

			//get the url for remote system
			rolesurl := env.Config.GetString(CONF_ROLES_API)
			if len(rolesurl) == 0 {
				return errors.ThrowError(ctx, CORE_ROLESAPI_NOT_FOUND)
			}
			//create remote system role
			roles, err := CreateCollection(ctx, env.SystemRole)
			if err != nil {
				return err
			}
			req, err := http.NewRequest("GET", rolesurl, nil)
			req.Header.Add(env.AuthHeader, token)
			res, err := client.Do(req)
			if err != nil {
				log.Logger.Trace(ctx, "core.env.remoteroles", "Error in roles query", "err", err)
				return err
			}
			body, err := ioutil.ReadAll(res.Body)
			log.Logger.Trace(ctx, "core.env.remoteroles", "result for roles query", "body", body)
			err = json.Unmarshal(body, &roles)
			/*base := sling.New().Set(env.AuthHeader, token)
			//req, err := base.New().Get("gophergram/list").Request()
			resp, err = base.New().Client(client).Get(rolesurl).ReceiveSuccess(roles)*/
			log.Logger.Trace(ctx, "core.env.remoteroles", "result for roles query", "err", err)
			if err != nil {
				return err
			}
			log.Logger.Trace(ctx, "core.env.remoteroles", "result for roles query", "Status code", resp.StatusCode)
			//get the response
			if resp.StatusCode != 200 {
				return errors.ThrowError(ctx, CORE_ROLESAPI_NOT_FOUND)
			}
			//register the roles and permissions received from auth system
			env.RegisterRoles(ctx, roles)
		}
	} else {
		//load permissions from remote system
		secsvcname := env.Config.GetString(CONF_SECURITY_SVC)
		if len(secsvcname) == 0 {
			errors.ThrowError(ctx, AUTH_MISSING_API)
		}
		secsvc, err := env.GetService(ctx, secsvcname)
		if err != nil {
			errors.RethrowError(ctx, AUTH_MISSING_API, err)
		}
		rolesInt, err := secsvc.Execute(ctx, "GetRoles", nil)
		if err != nil {
			return err
		}
		log.Logger.Trace(ctx, "core.env.localroles", "Got Roles")
		adminExists := false
		anonExists := false
		if rolesInt != nil {
			arr := reflect.ValueOf(rolesInt).Elem()
			length := arr.Len()
			for i := 0; i < length; i++ {
				role := arr.Index(i).Addr().Interface().(auth.Role)
				if role.GetId() == "Anonymous" {
					anonExists = true
				}
				if role.GetId() == env.AdminRole {
					adminExists = true
				}
				env.RegisterRolePermissions(ctx, role)
			}
			log.Logger.Trace(ctx, "core.env.localroles", "Registered Roles")
		}

		if !anonExists {
			aroleInt, err := CreateEmptyObject(ctx, env.SystemRole)
			anonymousRole := aroleInt.(auth.Role)
			anonymousRole.SetId("Anonymous")
			data := make(map[string]interface{}, 1)
			data["data"] = anonymousRole
			_, err = secsvc.Execute(ctx, "SaveRole", data)
			if err != nil {
				return err
			}
		}
		if !adminExists {
			aroleInt, err := CreateEmptyObject(ctx, env.SystemRole)
			adminRole := aroleInt.(auth.Role)
			adminRole.SetId(env.AdminRole)
			permissionsInt, err := secsvc.Execute(ctx, "GetPermissions", nil)
			if err != nil {
				return err
			}
			adminRole.SetPermissions(permissionsInt.([]string))
			data := make(map[string]interface{}, 1)
			data["data"] = adminRole
			_, err = secsvc.Execute(ctx, "SaveRole", data)
			if err != nil {
				return err
			}
		}
		//log.Logger.Trace(ctx, "core.env.localroles", "Got Registering roles")
		//env.RegisterRoles(ctx, rolesInt)
	}
	return nil
}
