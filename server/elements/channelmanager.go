package elements

import "laatoo.io/sdk/server/core"

type ChannelManager interface {
	core.ServerElement
	GetChannel(ctx core.ServerContext, name string) (Channel, bool)
	List(ctx core.ServerContext) core.StringsMap
	ChangeLogger(ctx core.ServerContext, chanName string, logLevel string, logFormat string) error
	Describe(ctx core.ServerContext, chanName string) (core.StringMap, error)
}
