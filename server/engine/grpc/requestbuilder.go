package grpc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type RequestBuilder func(stream grpc.ServerStream) (core.RequestContext, map[string]interface{}, error)

func (channel *grpcChannel) getRequestBuilder(ctx core.ServerContext, bodyParamName string) RequestBuilder {

	return func(serverStream grpc.ServerStream) (core.RequestContext, map[string]interface{}, error) {

		grpcctx := serverStream.Context()

		grpcMetadata, ok := metadata.FromIncomingContext(grpcctx)
		if grpcMetadata == nil || !ok {
			return nil, nil, fmt.Errorf("Metadata Not found")
		}
		log.Info(ctx, "got metadata for method", "grpcMetadata", grpcMetadata)

		reqCtx, err := ctx.CreateNewRequest(channel.svcName, channel.engine.proxy, serverStream, "")
		if err != nil {
			return nil, nil, errors.WrapError(ctx, err)
		}
		vals := make(map[string]interface{})

		vals["encoding"] = "protobuf"

		msg := []byte{}

		if err := serverStream.RecvMsg(&msg); err != nil {
			log.Error(ctx, "GRPC streaming of message failed")
			return reqCtx, nil, fmt.Errorf("Could not receive message from stream")
		}

		allBytes, err := ioutil.ReadAll(bytes.NewReader(msg))
		if err != nil {
			log.Error(ctx, "GRPC streaming of message failed")
			return reqCtx, nil, fmt.Errorf("Could not receive message from stream")
		}

		//use ioutil. Noop readcloser for streams

		vals[bodyParamName] = allBytes

		return reqCtx, vals, nil
	}

}
