package data

import (
	"laatoosdk/core"
)

//Object stored by data service
type Storable interface {
	GetId() string
	SetId(string)
	PreSave(ctx core.Context) error
	PostSave(ctx core.Context) error
	PostLoad(ctx core.Context) error
	GetIdField() string
}

//Factory function for creating storable
type StorableCreator func() interface{}
