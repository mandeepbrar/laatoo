package http

import (
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

/*"github.com/vulcand/oxy/forward"
"github.com/vulcand/oxy/buffer"
"github.com/vulcand/oxy/roundrobin"*/

//"laatoo/engine/http/goji"
type httpEngine struct {
	framework    net.Webframework
	name         string
	ssl          bool
	sslcert      string
	sslkey       string
	address      string
	path         string
	authHeader   string
	loadBalanced bool
	leader       bool
	//	lb *roundrobin.RoundRobin
	proxy       server.Engine
	rootChannel *httpChannel
	conf        config.Config
	fwname      string
}

func (eng *httpEngine) Initialize(ctx core.ServerContext, conf config.Config) error {
	initCtx := ctx.SubContext("InitializeEngine: " + eng.name)
	eng.fwname = "Echo"
	fw, ok := eng.conf.GetString(constants.CONF_HTTP_FRAMEWORK)
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
	ssl, ok := eng.conf.GetBool(constants.CONF_ENG_SSL)
	if ok && ssl {
		cert, ok := eng.conf.GetString(constants.CONF_ENG_SSLCERT)
		if !ok {
			return errors.ThrowError(initCtx, errors.CORE_ERROR_MISSING_CONF, "Config Name", constants.CONF_ENG_SSLCERT)
		}
		key, ok := eng.conf.GetString(constants.CONF_ENG_SSLKEY)
		if !ok {
			return errors.ThrowError(initCtx, errors.CORE_ERROR_MISSING_CONF, "Config Name", constants.CONF_ENG_SSLKEY)
		}
		eng.ssl = ssl
		eng.sslcert = cert
		eng.sslkey = key
	}
	address, ok := eng.conf.GetString(constants.CONF_SERVER_ADDRESS)
	if !ok {
		return errors.ThrowError(initCtx, errors.CORE_ERROR_MISSING_CONF, "Config name", constants.CONF_SERVER_ADDRESS)
	} else {
		eng.address = address
	}
	log.Trace(initCtx, "Initializing framework")
	eng.framework.Initialize()

	//eng.authHeader = ctx.GetServerVariable(core.AUTHHEADER).(string)

	eng.rootChannel = &httpChannel{name: eng.name, Router: eng.framework.GetParentRouter(""), config: eng.conf, engine: eng}
	err := eng.rootChannel.configure(ctx)
	if err != nil {
		return err
	}

	log.Trace(ctx, "Setting root channel", "root", eng.rootChannel)

	/*eng.loadBalanced, _ := eng.conf.GetBool(constants.CONF_LOAD_BALANCED)

	if(eng.loadBalanced) {
		eng.startGossping(ctx, conf)
	}*/

	//engCtx := ctx.SubContext("Configuring engine")
	/*if err = eng.router.ConfigureRoutes(engCtx); err != nil {
		return errors.RethrowError(engCtx, CORE_ERROR_INCORRECT_DELIVERY_CONF, err)
	}*/
	/*loaderCtx := ctx.GetElement(core.ServerElementLoader)
	return facMgr.createServiceFactories(ctx, conf, loaderCtx.(server.ObjectLoader))*/
	log.Debug(initCtx, "Initialized engine")
	return nil
}

/*
func (eng *httpEngine) startGossping(ctx core.ServerContext, conf config.Config) error {
	if(eng.leader) {
		if err := eng.startBalancer(ctx, eng.conf); err!=nil {
			return errors.ThrowError(ctx, err)
		}
	}
	if eng.lb != nil {
		eng.lb.UpsertServer(url)
	}
}


func (eng *httpEngine) insertClusterNode(ctx core.ServerContext, url string) error {
	if eng.lb != nil {
		eng.lb.UpsertServer(url)
	}
}

func (eng *httpEngine) startBalancer(ctx core.ServerContext, conf config.Config) error {
	fwd, _ := forward.New()
	eng.lb, _ := roundrobin.New(fwd)
	// buffer will read the request body and will replay the request again in case if forward returned status
	// corresponding to nework error (e.g. Gateway Timeout)
	buffer, _ := buffer.New(eng.lb, buffer.Retry(`IsNetworkError() && Attempts() < 2`))

	// that's it! our reverse proxy is ready!
	s := &http.Server{
		Addr:           ":8080",
		Handler:        buffer,
	}
	s.ListenAndServe()
}*/

func (eng *httpEngine) Start(ctx core.ServerContext) error {
	startCtx := ctx.SubContext("Start Engine: " + eng.name)
	log.Info(startCtx, "Starting http engine", "address", eng.address, "ssl", eng.ssl)
	if eng.ssl {
		//start listening
		err := eng.framework.StartSSLServer(eng.address, eng.sslcert, eng.sslkey)
		if err != nil {
			panic("Failed to start application" + err.Error())
		}
	} else {
		//start listening
		err := eng.framework.StartServer(eng.address)
		if err != nil {
			panic("Failed to start application" + err.Error())
		}
	}
	log.Info(startCtx, "Started engine*********************************")
	return nil
}
