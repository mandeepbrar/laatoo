// +build !appengine

package server

import (
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"log"
	"net"
	"time"
)

//Create a new server
func NewServer(configName string) (*Server, error) {
	server := &Server{ServerType: core.CONF_SERVERTYPE_STANDALONE}
	ctx := NewServerContext("ServerInit", nil, nil) // //Context: echo.NewContext(nil, nil, router)}
	err := server.InitServer(ctx, configName)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	//find the address to bind from the server
	address, ok := server.Config.GetString(CONF_SERVERTYPE_HOSTNAME)
	if !ok {
		panic("Host name not provided for standalone server")
	}

	go startServer(ctx, address, server)

	//start the standalone tcp loop
	// Listen on TCP port 2000 on all interfaces.
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			//io.Copy(c, c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
	return server, nil
}

func startServer(ctx *serverContext, address string, server *Server) error {
	for i := 0; i < 10; i++ {
		_, err := net.Dial("tcp", address)
		if err != nil {
			time.Sleep(100 * time.Millisecond)
		} else {
			return server.Start(ctx)
		}
	}
	panic("Server could not be started")
}
