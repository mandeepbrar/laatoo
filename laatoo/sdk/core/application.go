package core

import (
	"laatoo/sdk/config"
)

type ApplicationContext interface {
	Initialize(ctx ServerContext, conf config.Config) error
	Start(ctx ServerContext) error
}

type ApplicationContextProvider func(name string) (ApplicationContext, error)
