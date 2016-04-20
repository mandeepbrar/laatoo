// +build !appengine

package server

import (
	"crypto/tls"
	"fmt"
	glctx "golang.org/x/net/context"
	"golang.org/x/oauth2"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"net"
	"net/http"
	"time"
)

const (
	//if this file is built the server type will be standalone
	SERVER_TYPE = core.CONF_SERVERTYPE_STANDALONE
)

func startListening(ctx core.ServerContext, conf config.Config) error {
	//find the address to bind from the server
	address, ok := conf.GetString(config.CONF_SERVER_ADDRESS)
	if !ok {
		panic("Host name not provided for standalone server")
	}

	go dialServer(ctx, address)

	//start the standalone tcp loop
	// Listen on TCP port 2000 on all interfaces.
	//this is an admin port and not environment address
	// more functionality to be built on the admin port
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer l.Close()
	log.Logger.Info(ctx, "Listening...")
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Shut down the connection.
			c.Close()
		}(conn)
	}

}

//try tcp connection to the server... it will allow connection if the server is bound to the address
func dialServer(ctx core.ServerContext, address string) error {
	for i := 0; i < 10; i++ {
		_, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Print(err)
			time.Sleep(1000 * time.Millisecond)
		} else {
			return nil
		}
	}
	panic("Server could not be started")
}

func GetAppengineContext(ctx core.RequestContext) glctx.Context {
	return nil
}

func GetCloudContext(ctx core.RequestContext, scope string) glctx.Context {
	return nil
}
func HttpClient(ctx core.RequestContext) *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

func GetOAuthContext(ctx core.Context) glctx.Context {
	return oauth2.NoContext
}
