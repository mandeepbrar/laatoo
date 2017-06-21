package http

import (
	"laatoo/server/common"
	"laatoo/server/constants"
	"laatoo/server/engine/http/echo"
	"laatoo/server/engine/http/gin"

	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/engine/http/net"
)

//"laatoo/engine/http/goji"
type httpEngine struct {
	framework   net.Webframework
	name        string
	ssl         bool
	sslcert     string
	sslkey      string
	address     string
	path        string
	authHeader  string
	proxy       server.Engine
	rootChannel *httpChannel
	conf        config.Config
	fwname      string
}

func (eng *httpEngine) Initialize(ctx core.ServerContext, conf config.Config) error {
	initCtx := eng.createContext(ctx, "InitializeEngine: "+eng.name)
	eng.fwname = "Echo"
	fw, ok := conf.GetString(constants.CONF_HTTP_FRAMEWORK)
	if ok {
		eng.fwname = fw
	}
	switch eng.fwname {
	case "Echo":
		eng.framework = &echo.EchoWebFramework{}
	default:
		eng.framework = &gin.GinWebFramework{Name: eng.name}
		/*	case "Goji":
			eng.framework = &goji.GojiWebFramework{}*/
	}
	ssl, ok := conf.GetBool(constants.CONF_ENG_SSL)
	if ok && ssl {
		cert, ok := conf.GetString(constants.CONF_ENG_SSLCERT)
		if !ok {
			return errors.ThrowError(initCtx, errors.CORE_ERROR_MISSING_CONF, "Config Name", constants.CONF_ENG_SSLCERT)
		}
		key, ok := conf.GetString(constants.CONF_ENG_SSLKEY)
		if !ok {
			return errors.ThrowError(initCtx, errors.CORE_ERROR_MISSING_CONF, "Config Name", constants.CONF_ENG_SSLKEY)
		}
		eng.ssl = ssl
		eng.sslcert = cert
		eng.sslkey = key
	}
	if initCtx.GetServerType() == core.CONF_SERVERTYPE_STANDALONE {
		address, ok := conf.GetString(constants.CONF_SERVER_ADDRESS)
		if !ok {
			return errors.ThrowError(initCtx, errors.CORE_ERROR_MISSING_CONF, "Config name", constants.CONF_SERVER_ADDRESS)
		} else {
			eng.address = address
		}
	} else {
		rootPath, ok := conf.GetString(constants.CONF_HTTPENGINE_PATH)
		if !ok {
			return errors.ThrowError(initCtx, errors.CORE_ERROR_MISSING_CONF, "Config Name", constants.CONF_HTTPENGINE_PATH)
		}
		eng.path = rootPath
	}
	log.Trace(initCtx, "Initializing framework")
	eng.framework.Initialize()

	//eng.authHeader = ctx.GetServerVariable(core.AUTHHEADER).(string)
	eng.conf = conf

	eng.rootChannel = newHttpChannel(ctx, eng.name, eng.conf, eng, nil)

	//engCtx := ctx.SubContext("Configuring engine")
	/*if err = eng.router.ConfigureRoutes(engCtx); err != nil {
		return errors.RethrowError(engCtx, CORE_ERROR_INCORRECT_DELIVERY_CONF, err)
	}*/
	/*loaderCtx := ctx.GetElement(core.ServerElementLoader)
	return facMgr.createServiceFactories(ctx, conf, loaderCtx.(server.ObjectLoader))*/
	log.Debug(initCtx, "Initialized engine")
	return nil
}

func (eng *httpEngine) Start(ctx core.ServerContext) error {
	startCtx := eng.createContext(ctx, "Start Engine: "+eng.name)
	if startCtx.GetServerType() == core.CONF_SERVERTYPE_STANDALONE {
		log.Info(startCtx, "Starting http engine", "address", eng.address, "ssl", eng.ssl)
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
	if startCtx.GetServerType() == core.CONF_SERVERTYPE_GOOGLEAPP {
		common.GaeHandle(eng.path, eng.framework.GetRootHandler())
	}
	log.Info(startCtx, "Started engine*********************************")
	return nil
}

//creates a context specific to environment
func (eng *httpEngine) createContext(ctx core.ServerContext, name string) core.ServerContext {
	return ctx.NewContextWithElements(name,
		core.ContextMap{core.ServerElementEngine: eng.proxy}, core.ServerElementEngine)
}
