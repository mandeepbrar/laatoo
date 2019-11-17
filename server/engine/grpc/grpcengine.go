package grpc

import (
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/codecs"
	"laatoo/server/constants"
	"net"

	"google.golang.org/grpc"
)

type grpcEngine struct {
	name         string
	address      string
	proxy        elements.Engine
	rootChannel  *grpcChannel
	conf         config.Config
	codec        core.Codec
	svrContext   core.ServerContext
	routingTable map[string]*grpcChannel
	ssl          bool
}

func (eng *grpcEngine) Initialize(ctx core.ServerContext, conf config.Config) error {
	initCtx := ctx.SubContext("InitializeEngine: " + eng.name)

	address, ok := eng.conf.GetString(ctx, constants.CONF_SERVER_ADDRESS)
	if !ok {
		return errors.ThrowError(initCtx, errors.CORE_ERROR_MISSING_CONF, "Config name", constants.CONF_SERVER_ADDRESS)
	} else {
		eng.address = address
	}

	eng.ssl, _ = conf.GetBool(ctx, constants.CONF_ENG_SSL)

	eng.routingTable = make(map[string]*grpcChannel)

	codec, ok := conf.GetString(ctx, constants.CONF_ENGINE_CODEC)

	eng.codec, ok = ctx.GetCodec(codec)
	if !ok {
		eng.codec = codecs.NewProtobufCodec()
	}

	path, ok := conf.GetString(ctx, constants.CONF_GRPCENGINE_PATH)
	if !ok {
		errors.BadConf(ctx, constants.CONF_GRPCENGINE_PATH)
	}

	eng.rootChannel = &grpcChannel{name: eng.name, config: eng.conf, engine: eng, path: "/" + path}
	err := eng.rootChannel.configure(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	log.Debug(initCtx, "Initialized engine")
	return nil
}

func (eng *grpcEngine) Start(ctx core.ServerContext) error {
	startCtx := ctx.SubContext("Start GRPC Engine: " + eng.name)

	log.Error(startCtx, "Starting grpc engine", "address", eng.address)

	listener, err := net.Listen("tcp", eng.address)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

	serverMaxRecvMsgSize := 100000
	serverMaxSendMsgSize := 100000

	serverOptions := []grpc.ServerOption{
		grpc.CustomCodec(bytesCodec{}),
		grpc.UnknownServiceHandler(eng.getHandler(ctx)),
		grpc.MaxRecvMsgSize(serverMaxRecvMsgSize),
		grpc.MaxSendMsgSize(serverMaxSendMsgSize),
	}
	// Create new gRPC server with (blank) options
	grpcServer := grpc.NewServer(serverOptions...)

	// start the server
	if err := grpcServer.Serve(listener); err != nil {
		panic("ListenAndServe: " + err.Error())
	}

	return nil
}

func (eng *grpcEngine) getHandler(serverctx core.ServerContext) grpc.StreamHandler {
	return func(srv interface{}, serverStream grpc.ServerStream) error {
		streamMethod, ok := grpc.MethodFromServerStream(serverStream)
		if !ok {
			return fmt.Errorf("Invalid GRPC stream")
		}
		log.Info(serverctx, "Got grpc request for method", "method", streamMethod)

		channel, ok := eng.getChannel(serverctx, streamMethod)
		if !ok {
			return errors.ThrowError(serverctx, errors.CORE_ERROR_RES_NOT_FOUND, "Method", streamMethod)
		}
		log.Debug(serverctx, "Found channel for method", "route", streamMethod)

		if channel.disabled {
			return errors.ThrowError(serverctx, errors.CORE_ERROR_RES_NOT_FOUND, "Method", streamMethod)
		}

		if channel.handleRequest != nil {
			err := channel.handleRequest(serverctx, serverStream)
			grpcErr := eng.processErrors(serverctx, serverStream, err)
			if grpcErr != nil {
				return grpcErr
			}
		} else {
			return errors.InternalError(serverctx, "Error", "Request handler not configured")
		}
		return nil
	}
}

func (eng *grpcEngine) processErrors(serverctx core.ServerContext, serverStream grpc.ServerStream, err error) error {
	log.Debug(serverctx, "Processing errors ")
	return err
}

// customCodec pass bytes to/from the wire without modification.

type bytesCodec struct{}

// Marshal takes a []byte and passes it through as a []byte.
func (bytesCodec) Marshal(obj interface{}) ([]byte, error) {
	switch value := obj.(type) {
	case []byte:
		return value, nil
	default:
		return nil, fmt.Errorf("Do not know the type of object received")
	}
}

// Unmarshal takes a []byte pointer as obj and points it to data.
func (bytesCodec) Unmarshal(data []byte, obj interface{}) error {
	switch value := obj.(type) {
	case *[]byte:
		*value = data
		return nil
	default:
		return fmt.Errorf("Do not know the type of object received")
	}
}

func (bytesCodec) String() string {
	return "bytes codec"
}
