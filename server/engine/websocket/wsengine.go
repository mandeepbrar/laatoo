package websocket

import (
	"fmt"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/server"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type wsEngine struct {
	name        string
	address     string
	path        string
	proxy       server.Engine
	rootChannel *wsChannel
	conf        config.Config
	upgrader    websocket.Upgrader
}

func (eng *wsEngine) Initialize(ctx core.ServerContext, conf config.Config) error {
	initCtx := ctx.SubContext("InitializeEngine: " + eng.name)

	address, ok := eng.conf.GetString(ctx, constants.CONF_SERVER_ADDRESS)
	if !ok {
		return errors.ThrowError(initCtx, errors.CORE_ERROR_MISSING_CONF, "Config name", constants.CONF_SERVER_ADDRESS)
	} else {
		eng.address = address
	}

	eng.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	eng.origin = eng.address

	eng.rootChannel = &wsChannel{name: eng.name, config: eng.conf, engine: eng}
	err := eng.rootChannel.configure(ctx)
	if err != nil {
		return err
	}

	log.Debug(initCtx, "Initialized engine")
	return nil
}

func (eng *wsEngine) Start(ctx core.ServerContext) error {
	startCtx := ctx.SubContext("Start Engine: " + eng.name)
	log.Info(startCtx, "Starting websocket engine", "address", eng.address)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := eng.upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	//	origin := "http://localhost/"
	//	url := "ws://localhost:12345/ws"
	err := http.ListenAndServe(eng.address, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

	log.Info(startCtx, "Started engine*********************************")
	return nil
}
