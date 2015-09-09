// +build !appengine

package laatoocore

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"laatoosdk/errors"
	"net/http"
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
		server.Start()
		//find the address to bind from the server
		address := server.Config.GetString(CONF_SERVERTYPE_HOSTNAME)
		if address == "" {
			return nil, errors.ThrowError(CORE_SERVERADD_NOT_FOUND)
		}
		http.Handle("/", router)
		//start listening
		http.ListenAndServe(address, nil)
	}
	return server, nil
}
