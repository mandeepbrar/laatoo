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
	ObjectsFactoryRegister = make(map[string]interface{}, 30)
	ObjectCollections      = utils.NewMemoryStorer()
)

//every object that needs to be created by configuration should register a factory func
//this factory function provides a object if called
type ObjectFactory func(conf map[string]interface{}) (interface{}, error)

//register the object factory in the global register
func RegisterObjectProvider(objectName string, factory ObjectFactory) {
	_, ok := ObjectsFactoryRegister[objectName]
	if !ok {
		log.Logger.Info("core.objects", "Registering object factory ", "Object Name", objectName)
		ObjectsFactoryRegister[objectName] = factory
	}
}

//returns collection type for given object
func GetCollectionType(objectName string) (reflect.Type, error) {
	//find the collection type from memory
	typeInt, err := ObjectCollections.GetObject(objectName)
	if err == nil {
		return typeInt.(reflect.Type), nil
	} else {
		objectPtr, err := CreateObject(objectName, nil)
		if err != nil {
			return nil, err
		}
		collectionType := reflect.SliceOf(reflect.TypeOf(reflect.ValueOf(objectPtr).Elem().Interface()))
		log.Logger.Info("core.objects", "Creating collectionType ", "Collection Type", collectionType)
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

func CreateEmptyObject(objectName string) (interface{}, error) {
	return CreateObject(objectName, nil)
}

//Provides a object with a given name
func CreateObject(objectName string, confdata map[string]interface{}) (interface{}, error) {
	log.Logger.Trace("core.objects", "Getting object ", "Object Name", objectName)

	//get the factory from the register
	factoryInt, ok := ObjectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)

	}
	//cast to a creatory func
	factoryFunc, ok := factoryInt.(ObjectFactory)
	if !ok {
		return nil, errors.ThrowError(CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)
	}

	//call factory method for creating an object
	obj, err := factoryFunc(confdata)
	if err != nil {
		return nil, errors.RethrowError(CORE_ERROR_OBJECT_NOT_CREATED, err, "Object Name", objectName)
	}
	log.Logger.Trace("core.objects", "Created object ", "Object Name", objectName)
	//return object by calling factory func
	return obj, nil
}
