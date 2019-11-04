package grpc

import "laatoo/sdk/server/core"

func (eng *grpcEngine) addRoute(ctx core.ServerContext, route string, channel *grpcChannel) error {
	eng.routingTable[route] = channel
	return nil
}

func (eng *grpcEngine) getChannel(ctx core.ServerContext, route string) (*grpcChannel, bool) {
	channel, ok := eng.routingTable[route]
	return channel, ok
}
