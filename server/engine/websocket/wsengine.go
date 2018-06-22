package websocket

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/codecs"
	"laatoo/server/constants"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/twinj/uuid"
)

type wsEngine struct {
	name        string
	address     string
	path        string
	proxy       server.Engine
	rtr         *router
	rootChannel *wsChannel
	conf        config.Config
	codec       core.Codec
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

	codec, ok := conf.GetString(ctx, constants.CONF_ENGINE_CODEC)

	if codec == "fastjson" {
		eng.codec = codecs.NewFastJsonCodec()
	} else {
		eng.codec = codecs.NewJsonCodec()
	}

	eng.rtr, _ = newRouter(ctx, eng)

	eng.rootChannel = &wsChannel{name: eng.name, config: eng.conf, engine: eng}
	err := eng.rootChannel.configure(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	notifysvcName, ok := conf.GetString(ctx, "NotifierService")
	if !ok {
		notifysvcName = "WebClientNotifier"
	}
	eng.setupNotifierService(ctx, notifysvcName)

	log.Debug(initCtx, "Initialized engine")
	return nil
}

func (eng *wsEngine) Start(ctx core.ServerContext) error {
	startCtx := ctx.SubContext("Start Engine: " + eng.name)
	log.Error(startCtx, "Starting websocket engine", "address", eng.address)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := eng.upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		sessionId := uuid.NewV1().String()
		eng.rtr.addConnection(ctx, sessionId, conn)
		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			jsonRpcMsg := &rpcRequest{}
			err = eng.codec.Unmarshal(ctx, msg, jsonRpcMsg)
			if err != nil {
				conn.WriteMessage(msgType, []byte(err.Error()))
				return
			}
			go eng.rtr.routeMessage(ctx, jsonRpcMsg, conn)
		}
	})

	//	origin := "http://localhost/"
	//	url := "ws://localhost:12345/ws"
	err := http.ListenAndServe(eng.address, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
	return nil
}

func (eng *wsEngine) setupNotifierService(ctx core.ServerContext, objName string) error {
	objLoader := ctx.GetServerElement(core.ServerElementLoader).(server.ObjectLoader)
	objLoader.RegisterObject(ctx, objName, func() interface{} {
		return &NotifierService{engine: eng}
	}, nil, nil)
	return nil
}
