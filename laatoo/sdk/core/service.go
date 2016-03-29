package core

import (
	"laatoo/sdk/config"
)

type ServiceFunc func(ctx RequestContext) error

type Service interface {
	GetConf() config.Config
	Initialize(ServerContext) error
	Invoke(RequestContext) error
}
