package core

import (
	"laatoo/sdk/config"
)

type MethodArgs map[string]interface{}

//Creates object
type ObjectCreator func(ctx Context, args MethodArgs) (interface{}, error)

//Creates collection
type ObjectCollectionCreator func(ctx Context, args MethodArgs) (interface{}, error)

//interface that needs to be implemented by any object provider in a system
type ObjectFactory interface {
	//Initialize the object factory
	Initialize(ctx ServerContext, config config.Config) error
	//Creates object
	CreateObject(ctx Context, args MethodArgs) (interface{}, error)
	//Creates collection
	CreateObjectCollection(ctx Context, args MethodArgs) (interface{}, error)
}
