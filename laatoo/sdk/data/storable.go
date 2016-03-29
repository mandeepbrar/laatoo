package data

import (
	"fmt"
	"laatoo/sdk/core"
	"reflect"
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
//type StorableCreator func() interface{}

func CastToStorableCollection(items interface{}) ([]Storable, error) {
	arr := reflect.ValueOf(items).Elem()
	if arr.Kind() != reflect.Array {
		return nil, fmt.Errorf("Invalid cast")
	}
	length := arr.Len()
	retVal := make([]Storable, length)
	for i := 0; i < length; i++ {
		valPtr := arr.Index(i).Addr().Interface()
		stor, ok := valPtr.(Storable)
		if !ok {
			return nil, fmt.Errorf("Invalid cast")
		}
		retVal[i] = stor
	}
	return retVal, nil
}
