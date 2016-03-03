package laatoocore

import (
	"laatoosdk/core"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/utils"
	"reflect"
)

var (
	//global provider register
	//objects factory register exists for every server
	ObjectsFactoryRegister = make(map[string]interface{}, 30)

	//register for collection types of objects
	//types of object collections are only registered once
	ObjectCollections = utils.NewMemoryStorer()
)

//every object that needs to be created by configuration should register a factory func
//this factory function provides a object if called
type ObjectFactory func(ctx core.Context, conf map[string]interface{}) (interface{}, error)

//register the object factory in the global register
func RegisterObjectProvider(objectName string, factory ObjectFactory) {
	_, ok := ObjectsFactoryRegister[objectName]
	if !ok {
		log.Logger.Info(nil, "core.objects", "Registering object factory ", "Object Name", objectName)
		ObjectsFactoryRegister[objectName] = factory
	}
}

//returns collection type for given object
func GetCollectionType(ctx core.Context, objectName string) (reflect.Type, error) {
	//find the collection type from memory
	typeInt, err := ObjectCollections.GetObject(objectName)

	if err == nil {
		//if present in registry return it
		return typeInt.(reflect.Type), nil
	} else {

		//if not present create type of collection by reflection
		//create an object of the type by factory func

		objectPtr, err := CreateObject(ctx, objectName, nil)
		if err != nil {
			return nil, err
		}
		// derive collection type from object type
		collectionType := reflect.SliceOf(reflect.TypeOf(reflect.ValueOf(objectPtr).Elem().Interface()))
		log.Logger.Info(ctx, "core.objects", "Creating collectionType ", "Collection Type", collectionType)
		ObjectCollections.PutObject(objectName, collectionType)
		return collectionType, nil
	}
}

//returns a collection of the object type
func CreateCollection(ctx core.Context, objectName string) (interface{}, error) {
	//get the collection type from registry
	collectionType, err := GetCollectionType(ctx, objectName)
	if err != nil {
		return nil, err
	}
	//create an instance of collection.
	//returns pointer to the collection
	return reflect.New(collectionType).Interface(), nil
}

//returns an object without any config
func CreateEmptyObject(ctx core.Context, objectName string) (interface{}, error) {
	return CreateObject(ctx, objectName, nil)
}

//Provides an object with a given name
func CreateObject(ctx core.Context, objectName string, confdata map[string]interface{}) (interface{}, error) {
	log.Logger.Trace(ctx, "core.objects", "Getting object ", "Object Name", objectName)

	//get the factory func from the register
	factoryInt, ok := ObjectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)

	}

	//cast to a creatory func
	factoryFunc, ok := factoryInt.(ObjectFactory)
	if !ok {
		return nil, errors.ThrowError(ctx, CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)
	}

	//call factory method for creating an object
	obj, err := factoryFunc(ctx, confdata)
	if err != nil {
		return nil, errors.RethrowError(ctx, CORE_ERROR_OBJECT_NOT_CREATED, err, "Object Name", objectName)
	}

	log.Logger.Trace(ctx, "core.objects", "Created object ", "Object Name", objectName)
	//return object by calling factory func
	return obj, nil
}
