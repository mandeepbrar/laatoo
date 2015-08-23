package laatoocore

import (
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/utils"
	"reflect"
)

var (
	//global provider register
	//objects factory register exists for every server
	ObjectsFactoryRegister = utils.NewMemoryStorer()
	ObjectCollections      = utils.NewMemoryStorer()
	EmptyObjects           = utils.NewMemoryStorer()
)

//every object that needs to be created by configuration should register a factory func
//this factory function provides a object if called
type ObjectFactory func(conf map[string]interface{}) (interface{}, error)

//register the object factory in the global register
func RegisterObjectProvider(objectName string, factory ObjectFactory) {

	log.Logger.Infof("Registering object factory for %s", objectName)
	ObjectsFactoryRegister.PutObject(objectName, factory)
}

func CreateEmptyObject(objectName string) (interface{}, error) {
	typeInt, err := EmptyObjects.GetObject(objectName)
	if err == nil {
		log.Logger.Infof("Returning object ", typeInt)
		return reflect.New(typeInt.(reflect.Type)).Interface(), nil
	} else {
		objectPtr, err := createObject(objectName, nil)
		if err != nil {
			return nil, err
		}
		typeVal := reflect.TypeOf(reflect.ValueOf(objectPtr).Elem().Interface())
		log.Logger.Infof("Creating object   %s", typeVal)
		EmptyObjects.PutObject(objectName, typeVal)
		return objectPtr, nil
	}
}

//returns collection type for given object
func GetCollectionType(objectName string) (reflect.Type, error) {
	//find the collection type from memory
	typeInt, err := ObjectCollections.GetObject(objectName)
	if err == nil {
		return typeInt.(reflect.Type), nil
	} else {
		objectPtr, err := createObject(objectName, nil)
		if err != nil {
			return nil, err
		}
		collectionType := reflect.SliceOf(reflect.TypeOf(reflect.ValueOf(objectPtr).Elem().Interface()))
		log.Logger.Infof("Creating collectionType   %s", collectionType)
		ObjectCollections.PutObject(objectName, collectionType)
		return collectionType, nil
	}
}

func CreateCollection(objectName string) (interface{}, error) {
	collectionType, err := GetCollectionType(objectName)
	if err != nil {
		return nil, err
	}
	return reflect.New(collectionType).Interface(), nil
}

//Provides a object with a given name
func CreateObject(objectName string, confdata map[string]interface{}) (interface{}, error) {
	if confdata == nil {
		return CreateEmptyObject(objectName)
	} else {
		return createObject(objectName, confdata)
	}
}

//Provides a object with a given name
func createObject(objectName string, confdata map[string]interface{}) (interface{}, error) {
	log.Logger.Debugf("Getting object %s", objectName)

	//get the factory from the register
	factoryInt, err := ObjectsFactoryRegister.GetObject(objectName)
	if err != nil {
		return nil, errors.RethrowError(CORE_ERROR_PROVIDER_NOT_FOUND, err, objectName)

	}
	//cast to a creatory func
	factoryFunc, ok := factoryInt.(ObjectFactory)
	if !ok {
		return nil, errors.RethrowError(CORE_ERROR_PROVIDER_NOT_FOUND, err, objectName)
	}

	log.Logger.Debugf("Creating object %s from factory", objectName)
	//call factory method for creating an object
	obj, err := factoryFunc(confdata)
	if err != nil {
		return nil, errors.RethrowError(CORE_ERROR_OBJECT_NOT_CREATED, err, objectName)
	}
	log.Logger.Debugf("Created object %s", objectName)
	//return object by calling factory func
	return obj, nil
}
