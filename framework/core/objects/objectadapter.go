package objects

import (
	"laatoo/sdk/core"
	"reflect"
)

//service method for doing various tasks
func NewObjectFactory(objectCreator core.ObjectCreator, objectCollectionCreator core.ObjectCollectionCreator) core.ObjectFactory {
	if objectCreator != nil {
		return &objectFactory{objectCreator: objectCreator, objectCollectionCreator: objectCollectionCreator}
	} else {
		panic("Could not register object factory. Creator is nil.")
	}
	return nil
}

func NewObjectType(obj interface{}) *objectFactory {
	typ := reflect.TypeOf(obj)
	slice := reflect.MakeSlice(reflect.SliceOf(typ), 0, 0)
	slicTyp := slice.Type()
	objectCreator := func() interface{} {
		return reflect.New(typ).Interface()
	}
	collectCreator := func(length int) interface{} {
		k := reflect.MakeSlice(reflect.SliceOf(typ), length, length)
		// Create a pointer to a slice value and set it to the slice
		x := reflect.New(slicTyp)
		x.Elem().Set(k)
		return x.Interface()
	}
	return &objectFactory{objectCreator: objectCreator, objectCollectionCreator: collectCreator}
}

type objectFactory struct {
	objectCreator           core.ObjectCreator
	objectCollectionCreator core.ObjectCollectionCreator
}

//Creates object
func (fac *objectFactory) CreateObject() interface{} {
	return fac.objectCreator()
}

//Creates collection
func (fac *objectFactory) CreateObjectCollection(length int) interface{} {
	return fac.objectCollectionCreator(length)
}
