package grpc

import (
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type defaultResponseHandler struct {
	codec      core.Codec
	svrContext core.ServerContext
}

func DefaultResponseHandler(ctx core.ServerContext, codec core.Codec) *defaultResponseHandler {
	return &defaultResponseHandler{codec: codec}
}

func (rh *defaultResponseHandler) Initialize(ctx core.ServerContext, conf config.Config) error {
	rh.svrContext = ctx
	return nil
}

func (rh *defaultResponseHandler) HandleResponse(ctx core.RequestContext, resp *core.Response, handlingError error) error {
	serverStream := ctx.EngineRequestContext().(grpc.ServerStream)
	log.Trace(ctx, "Returning request with default response handler")
	if resp == nil {
		resp = core.StatusSuccessResponse
	}
	switch resp.Status {
	case core.StatusSuccess:
		rh.sendResponse(ctx, serverStream, resp)
	case core.StatusInternalError:
		rh.sendResponse(ctx, serverStream, resp)
	case core.StatusFunctionalError:
		rh.sendResponse(ctx, serverStream, resp)
	default:
		log.Error(ctx, "HandleResponse status not implemented", "resp", resp)
		rh.sendResponse(ctx, serverStream, core.NewServiceResponse(core.StatusInternalError, nil))
	}
	return nil
}

func (rh *defaultResponseHandler) sendResponse(ctx core.RequestContext, serverStream grpc.ServerStream, resp *core.Response) error {
	var err error
	if resp.MetaInfo != nil {
		for k, v := range resp.MetaInfo {
			header := metadata.Pairs(k, fmt.Sprint(v))
			err = serverStream.SendHeader(header)
		}
	}
	if err != nil {
		log.Error(ctx, "Error in sending header on GRPC", "Error", err)
		return err
	}

	if resp.Data != nil {
		byts, err := rh.codec.Marshal(ctx, resp.Data)
		if err != nil {
			log.Error(ctx, "Error in marshalling response", "Error", err)
			return fmt.Errorf("Error marshalling response")
		}
		err = serverStream.SendMsg(byts)
	} else {
		err = serverStream.SendMsg([]byte{})
	}
	if err != nil {
		log.Error(ctx, "Error in sending response on GRPC stream", "Error", err)
		return err
	}
	log.Trace(ctx, "Sent response from grpc", "resp", resp)
	return nil
}

func (rh *defaultResponseHandler) Reference() core.ServerElement {
	anotherref := rh
	return anotherref
}

func (rh *defaultResponseHandler) GetProperty(name string) interface{} {
	return nil
}

func (rh *defaultResponseHandler) GetName() string {
	return "GRPCResponseHandler"
}
func (rh *defaultResponseHandler) GetType() core.ServerElementType {
	return core.ServerElementServiceResponseHandler
}
func (rh *defaultResponseHandler) GetContext() core.ServerContext {
	return rh.svrContext
}

/*
type Void struct {
	proto.Message
}

func (Void) Reset() {
}
func (Void) String() string {
	return ""
}
func (Void) ProtoMessage() {}
*/
