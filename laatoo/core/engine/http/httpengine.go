package http

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	mw "github.com/labstack/echo/middleware"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"net/http"
)

const (
	CONF_SERVERTYPE_HOSTNAME = "hostname"
	CONF_ENVPATH             = "path"
	CONF_ROUTECONF           = "routes"
	CONF_SERVER_SSL          = "ssl"
	CONF_SSLCERT             = "sslcert"
	CONF_SSLKEY              = "sslkey"
)

type HttpEngine struct {
	rootRouter *echo.Echo
	router     *Router
	ssl        bool
	sslcert    string
	sslkey     string
	address    string
	path       string
	config     config.Config
}

func NewHttpEngine(ctx core.ServerContext, conf config.Config) (*HttpEngine, error) {
	eng := &HttpEngine{config: conf, ssl: false}
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
	envPath, ok := conf.GetString(CONF_ENVPATH)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Config Name", CONF_ENVPATH)
	}
	eng.path = envPath
	address, ok := conf.GetString(CONF_SERVERTYPE_HOSTNAME)
	if !ok {
		if ctx.GetServerType() == core.CONF_SERVERTYPE_STANDALONE {
			return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Config name", CONF_SERVERTYPE_HOSTNAME)
		}
	} else {
		eng.address = address
	}

	//create all service factories in the environment
	//initialize router
	router := echo.New()
	// Middleware
	router.Use(mw.Logger())
	router.Use(mw.Recover())
	router.Use(mw.Gzip())
	eng.rootRouter = router

	return eng, nil
}

func (eng *HttpEngine) GetContext() core.EngineContext {
	return &HttpEngineContext{eng}
}

func (eng *HttpEngine) InitializeEngine(ctx core.ServerContext) error {
	routesConf, ok := eng.config.GetString(CONF_ROUTECONF)
	if ok {
		routerFileName := fmt.Sprintf("%s/%s", ctx.GetServerName(), routesConf)
		config, err := config.NewConfigFromFile(routerFileName)
		if err != nil {
			return errors.RethrowError(ctx, errors.CORE_ERROR_MISSING_CONF, err, "Router config name", routesConf)
		}
		eng.router = &Router{eRouter: eng.rootRouter.Group(""), config: config, engine: eng}
		if err = eng.router.ConfigureRoutes(ctx); err != nil {
			return errors.RethrowError(ctx, CORE_ERROR_INCORRECT_DELIVERY_CONF, err)
		}
	}
	if ctx.GetServerType() == core.CONF_SERVERTYPE_GOOGLEAPP {
		s := standard.New("")
		s.SetHandler(eng.rootRouter)
		http.Handle(eng.path, s)
	}
	return nil
}

func (eng *HttpEngine) StartEngine(ctx core.ServerContext) error {
	if ctx.GetServerType() == core.CONF_SERVERTYPE_STANDALONE {
		log.Logger.Info(ctx, "Starting server", "address", eng.address, "ssl", eng.ssl)
		if eng.ssl {
			//start listening
			eng.rootRouter.Run(standard.NewFromTLS(eng.address, eng.sslcert, eng.sslkey)) //http.ListenAndServeTLS(delivery.address, delivery.sslcert, delivery.sslkey, delivery.rootRouter)
			/*if err != nil {
				panic("Failed to start environment" + err.Error())
			}*/
			return nil
		} else {
			//start listening
			eng.rootRouter.Run(standard.New(eng.address)) //http.ListenAndServe(delivery.address, delivery.rootRouter)
			/*if err != nil {
				panic("Failed to start environment" + err.Error())
			}*/
			return nil
		}
	}
	return nil
}
