// +build appengine

package server

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	//	"laatoo/sdk/errors"
	//	"laatoo/sdk/log"
	"log"
	"net/http"
	"sync"
)

const (
	SERVER_TYPE = core.CONF_SERVERTYPE_GOOGLEAPP
)

func Main(configFile string) error {
	var once sync.Once
	var request *http.Request
	warmupFunc := func() {
		rootctx := newServerContext()
		rootctx.SetGaeReq(request)
		err := main(rootctx, configFile)
		if err == nil {
			log.Println("**********Listening")
		} else {
			log.Fatal(err)
			panic(err)
		}
	}
	http.HandleFunc("/_ah/warmup", func(w http.ResponseWriter, req *http.Request) {
		request = req
		once.Do(warmupFunc)
	})
	/*http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		request = req
		once.Do(warmupFunc)
	})*/
	return nil
}

func startListening(ctx core.ServerContext, conf config.Config) error {
	/*//find the address to bind from the server
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
	*/
	return nil
}

/*
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

//Create a new server
func NewServer(configName string) (*Server, error) {
	ctx := &Context{Context: echo.NewContext(nil, nil, router)}
	server := &Server{ServerType: CONF_SERVERTYPE_GOOGLEAPP}
	err := server.InitServer(ctx, configName, router)
	if err != nil {
		return nil, errors.WrapError(err)
	}
	http.Handle("/", router)
	log.Logger.Error(ctx, "core.appengine.warmup", "setting up router for warmup")
	var req *Context
	var once sync.Once
	warmupFunc := func() {
		log.Logger.Error(req, "core.appengine.warmup", "starting server")
		server.Start(req)
	}
	router.Use(func(ctx *echo.Context) error {
		req = &Context{Context: ctx}
		once.Do(warmupFunc)
		return nil
	})
	router.Get("/_ah/warmup", func(ctx *echo.Context) error {
		req = &Context{Context: ctx}
		once.Do(warmupFunc)
		return nil
	})
	return server, nil
}
*/
