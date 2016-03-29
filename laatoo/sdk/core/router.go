package core

import (
	"laatoo/sdk/config"
)

type HandlerFunc func(ctx RequestContext) error

type Router interface {
	//Get a sub router
	Group(ctx ServerContext, path string, conf config.Config) Router
	//Use middleware
	Use(ctx ServerContext, handler interface{})
	Get(ctx ServerContext, path string, conf config.Config, handler HandlerFunc) error
	Post(ctx ServerContext, path string, conf config.Config, handler HandlerFunc) error
	Put(ctx ServerContext, path string, conf config.Config, handler HandlerFunc) error
	Delete(ctx ServerContext, path string, conf config.Config, handler HandlerFunc) error
	Static(ctx ServerContext, path string, conf config.Config, dir string) error
	ServeFile(ctx ServerContext, pagePath string, conf config.Config, dest string)
}
