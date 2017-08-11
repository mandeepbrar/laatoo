package core

import (
	"laatoo/sdk/core"
	"reflect"
)

//service method for doing various tasks
func newObjectFactory(objectCreator core.ObjectCreator, objectCollectionCreator core.ObjectCollectionCreator, metadata core.Info) *objectFactory {
	if objectCreator != nil {
		return &objectFactory{objectCreator: objectCreator, objectCollectionCreator: objectCollectionCreator, metadata: metadata}
	} else {
		panic("Could not register object factory. Creator is nil.")
	}
}

func newObjectType(obj interface{}, metadata core.Info) *objectFactory {
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
	return &objectFactory{objectCreator: objectCreator, objectCollectionCreator: collectCreator, metadata: metadata}
}

type objectFactory struct {
	objectCreator           core.ObjectCreator
	objectCollectionCreator core.ObjectCollectionCreator
	metadata                core.Info
}

//Creates object
func (fac *objectFactory) CreateObject() interface{} {
	return fac.objectCreator()
}

//Creates collection
func (fac *objectFactory) CreateObjectCollection(length int) interface{} {
	return fac.objectCollectionCreator(length)
}

func (fac *objectFactory) Info() core.Info {
	return fac.metadata
}
