package server

import (
	"laatoo/sdk/core"
)

type ObjectLoader interface {
	core.ServerElement
	RegisterObjectFactory(ctx core.Context, objectName string, factory core.ObjectFactory)
	RegisterObject(ctx core.Context, objectName string, objectCreator core.ObjectCreator, objectCollectionCreator core.ObjectCollectionCreator)
	CreateCollection(ctx core.Context, objectName string, args core.MethodArgs) (interface{}, error)
	CreateObject(ctx core.Context, objectName string, args core.MethodArgs) (interface{}, error)
	GetObjectCollectionCreator(ctx core.Context, objectName string) (core.ObjectCollectionCreator, error)
	GetObjectCreator(ctx core.Context, objectName string) (core.ObjectCreator, error)
}