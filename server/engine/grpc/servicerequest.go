package grpc

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/constants"

	"google.golang.org/grpc"
)

func (channel *grpcChannel) serve(ctx core.ServerContext) error {
	ctx = ctx.SubContext("Serve")
	if !channel.disabled {
		log.Trace(ctx, "Channel config", "name", channel.name, "config", channel.config)

		svcManager := ctx.GetServerElement(core.ServerElementServiceManager).(elements.ServiceManager)
		svc, err := svcManager.GetService(ctx, channel.svcName)
		if err != nil {
			return err
		}
		channel.svc = svc

		handler := ctx.GetServerElement(core.ServerElementServiceResponseHandler)
		if handler != nil {
			channel.respHandler = handler.(elements.ServiceResponseHandler)
		} else {
			channel.respHandler = DefaultResponseHandler(ctx, channel.engine.codec)
		}

		channel.handleRequest = channel.getHandler(ctx)

		log.Error(ctx, "Assigning handle request", "name", channel.name, "config", channel.handleRequest)

	}
	return nil
}

func (channel *grpcChannel) getHandler(ctx core.ServerContext) func(core.ServerContext, grpc.ServerStream) error {

	bodyParamName := "Data"
	bodyParam, ok := channel.config.GetString(ctx, constants.CONF_HTTPENGINE_BODYPARAMNAME)
	if ok {
		bodyParamName = bodyParam
	}

	reqBuilder := channel.getRequestBuilder(ctx, bodyParamName)

	return func(ctx core.ServerContext, serverStream grpc.ServerStream) error {

		reqCtx, vals, err := reqBuilder(serverStream)
		if err != nil {
			err = channel.respHandler.HandleResponse(reqCtx, core.BadRequestResponse(err.Error()))
			return err
		}
		defer reqCtx.CompleteRequest()

		/*err = authRequest(reqCtx, vals)
		if err != nil {
			err = respHandler.HandleResponse(reqCtx, core.StatusUnauthorizedResponse)
			errChannel <- err
			return
		}
		*/

		log.Trace(reqCtx, "Invoking service ", "vals", vals)
		resp, err := channel.svc.HandleRequest(reqCtx, vals)
		log.Trace(reqCtx, "Completed request for service. Handling Response")

		if err == nil {
			err = channel.respHandler.HandleResponse(reqCtx, resp)
			return err
		} else {
			return errors.BadRequest(reqCtx, err.Error())
		}
		return nil
	}
}

/*



		name := &Name{}
		//		var msg []byte
		if err := serverStream.RecvMsg(name); err != nil {
			log.Error(serverctx, "GRPC streaming of message failed")
			return fmt.Errorf("Could not receive message from stream")
		}
		log.Error(serverctx, "Received the following bytes", "name", name)


type Name struct {
	Name     string `protobuf:"bytes,5,name=name,proto3" json:"name,omitempty"`
	Newfield string `protobuf:"bytes,6,name=newfield,proto3"`
}

func (m *Name) Reset()         { *m = Name{} }
func (m *Name) String() string { return proto.CompactTextString(m) }
func (*Name) ProtoMessage()    {}

*/
