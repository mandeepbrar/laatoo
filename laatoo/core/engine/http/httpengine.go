package http

import (
	"laatoo/core/common"
	"laatoo/core/engine/http/echo"
	"laatoo/core/engine/http/net"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"net/http"
)

const (
	CONF_ENGINE_NAME         = "http"
	CONF_SERVERTYPE_HOSTNAME = "hostname"
	CONF_APPPATH             = "path"
	CONF_ROUTECONF           = "routes"
	CONF_SERVER_SSL          = "ssl"
	CONF_SSLCERT             = "sslcert"
	CONF_SSLKEY              = "sslkey"
)

type HttpEngine struct {
	framework  net.Webframework
	router     *Router
	ssl        bool
	sslcert    string
	sslkey     string
	address    string
	path       string
	authHeader string
	config     config.Config
}

func NewHttpEngine(ctx core.ServerContext, conf config.Config) (*HttpEngine, error) {
	eng := &HttpEngine{config: conf, ssl: false}
	eng.framework = &echo.EchoWebFramework{}
	ssl, ok := conf.GetBool(CONF_SERVER_SSL)
	if ok && ssl {
		cert, ok := conf.GetString(CONF_SSLCERT)
		if !ok {
			return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Config Name", CONF_SSLCERT)
		}
		key, ok := conf.GetString(CONF_SSLKEY)
		if !ok {
			return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Config Name", CONF_SSLKEY)
		}
		eng.ssl = ssl
		eng.sslcert = cert
		eng.sslkey = key
	}
	appPath, ok := conf.GetString(CONF_APPPATH)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Config Name", CONF_APPPATH)
	}
	eng.path = appPath
	address, ok := conf.GetString(CONF_SERVERTYPE_HOSTNAME)
	if !ok {
		if ctx.GetServerType() == core.CONF_SERVERTYPE_STANDALONE {
			return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Config name", CONF_SERVERTYPE_HOSTNAME)
		}
	} else {
		eng.address = address
	}
	eng.framework.Initialize()
	return eng, nil
}

func (eng *HttpEngine) GetContext() core.EngineServerContext {
	return &HttpEngineContext{eng}
}

func (eng *HttpEngine) InitializeEngine(ctx core.ServerContext) error {
	eng.authHeader = ctx.GetServerVariable(core.AUTHHEADER).(string)
	routesConf, err := common.ConfigFileAdapter(eng.config, CONF_ROUTECONF)
	if err != nil {
		return errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_CONF, err, "Router config ", CONF_ROUTECONF)
	}
	router := &Router{name: "Root", Router: eng.framework.GetParentRouter(), config: routesConf, engine: eng}
	eng.router = router
	engCtx := ctx.SubContext("Configuring engine", routesConf)
	if err = eng.router.ConfigureRoutes(engCtx); err != nil {
		return errors.RethrowError(engCtx, CORE_ERROR_INCORRECT_DELIVERY_CONF, err)
	}
	if ctx.GetServerType() == core.CONF_SERVERTYPE_GOOGLEAPP {
		http.Handle(eng.path, eng.framework.GetRootHandler())
	}
	return nil
}

func (eng *HttpEngine) StartEngine(ctx core.ServerContext) error {
	if ctx.GetServerType() == core.CONF_SERVERTYPE_STANDALONE {
		log.Logger.Info(ctx, "Starting server", "address", eng.address, "ssl", eng.ssl)
		if eng.ssl {
			//start listening
			err := eng.framework.StartSSLServer(eng.address, eng.sslcert, eng.sslkey)
			if err != nil {
				panic("Failed to start application" + err.Error())
			}
			return nil
		} else {
			//start listening
			err := eng.framework.StartServer(eng.address)
			if err != nil {
				panic("Failed to start application" + err.Error())
			}
			return nil
		}
	}
	return nil
}
