package core

import (
	"laatoo/sdk/config"
)

type HandlerFunc func(ctx RequestContext) error

type Router interface {
	//Get a sub router
	Group(ctx ServerContext, path string, name string, conf config.Config) Router
	//Use middleware
	Use(ctx ServerContext, handler interface{})
	Get(ctx ServerContext, path string, conf config.Config, handler HandlerFunc) error
	Post(ctx ServerContext, path string, conf config.Config, handler HandlerFunc) error
	Put(ctx ServerContext, path string, conf config.Config, handler HandlerFunc) error
	Delete(ctx ServerContext, path string, conf config.Config, handler HandlerFunc) error
}
