// +build !appengine

package laatoocore

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"laatoosdk/errors"
	"laatoosdk/log"
	"net"
	"net/http"
	"time"
)

//Create a new server
func NewServer(configName string, serverType string) (*Server, error) {
	//initialize router
	router := echo.New()
	// Middleware
	router.Use(mw.Logger())
	router.Use(mw.Recover())
	server := &Server{ServerType: serverType}
	server.InitServer(configName, router)
	//listen if server type is standalone
	if serverType == CONF_SERVERTYPE_STANDALONE {
		//find the address to bind from the server
		address := server.Config.GetString(CONF_SERVERTYPE_HOSTNAME)
		if address == "" {
			return nil, errors.ThrowError(nil, CORE_SERVERADD_NOT_FOUND)
		}
		http.Handle("/", router)
		go startServer(nil, address, server)
		log.Logger.Info(nil, "core.server", "Starting server", "address", address)
		//start listening
		err := http.ListenAndServe(address, nil)
		if err != nil {
			log.Logger.Error(nil, "core.server", "Error in listening", "address", address, "Error", err)
		}
	}
	return server, nil
}

func startServer(ctx interface{}, address string, server *Server) {
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
