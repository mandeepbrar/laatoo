// +build !appengine

package laatoocore

import (
	"crypto/tls"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	glctx "golang.org/x/net/context"
	"laatoosdk/errors"
	"laatoosdk/log"
	"net"
	"net/http"
	"time"
)

const (
	CONF_SERVER_SSL = "ssl"
	CONF_SSLCERT    = "sslcert"
	CONF_SSLKEY     = "sslkey"
)

//Create a new server
func NewServer(configName string, serverType string) (*Server, error) {
	//initialize router
	router := echo.New()
	// Middleware
	router.Use(mw.Logger())
	router.Use(mw.Recover())
	router.Use(mw.Gzip())
	ctx := &Context{Context: echo.NewContext(nil, nil, router)}
	server := &Server{ServerType: serverType}
	server.InitServer(ctx, configName, router)
	//listen if server type is standalone
	if serverType == CONF_SERVERTYPE_STANDALONE {
		//find the address to bind from the server
		address := server.Config.GetString(CONF_SERVERTYPE_HOSTNAME)
		if address == "" {
			return nil, errors.ThrowError(ctx, CORE_SERVERADD_NOT_FOUND)
		}
		http.Handle("/", router)
		go startServer(ctx, address, server)
		ssl := server.Config.GetBool(CONF_SERVER_SSL)
		log.Logger.Info(ctx, "core.server", "Starting server", "address", address, "ssl", ssl)
		var err error
		if ssl {
			cert := server.Config.GetString(CONF_SSLCERT)
			key := server.Config.GetString(CONF_SSLKEY)
			//start listening
			err = http.ListenAndServeTLS(address, cert, key, nil)
		} else {
			//start listening
			err = http.ListenAndServe(address, nil)
		}

		if err != nil {
			log.Logger.Error(ctx, "core.server", "Error in listening", "address", address, "Error", err)
		}
	}
	return server, nil
}

func startServer(ctx *Context, address string, server *Server) {
	for i := 0; i < 10; i++ {
		_, err := net.Dial("tcp", address)
		if err != nil {
			time.Sleep(100 * time.Millisecond)
		} else {
			server.Start(ctx)
			return
		}
	}
	panic("Server could not be started")
}

func GetAppengineContext(ctx *Context) glctx.Context {
	return nil
}

func GetCloudContext(ctx *Context, scope string) glctx.Context {
	return nil
}
func HttpClient(ctx *Context) *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}
