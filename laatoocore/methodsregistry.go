package laatoocore

import (
	"laatoosdk/core"
	"laatoosdk/errors"
	"laatoosdk/log"
)

var (
	//global provider register
	//objects factory register exists for every server
	InvokableMethodsRegister = make(map[string]interface{}, 30)
)

//this method can be invoked remotely
type InvokableMethod func(ctx core.Context) error

//register the invokable method in the global register
func RegisterInvokableMethod(methodName string, method InvokableMethod) {
	_, ok := InvokableMethodsRegister[methodName]
	if !ok {
		log.Logger.Info(nil, "core.methods", "Registering invokable method ", "Method Name", methodName)
		InvokableMethodsRegister[methodName] = method
	}
}

//Provides an object with a given name
func GetMethod(ctx core.Context, methodName string) (InvokableMethod, error) {
	log.Logger.Trace(ctx, "core.methods", "Getting method ", "Method Name", methodName)

	//get the factory func from the register
	methodInt, ok := InvokableMethodsRegister[methodName]
	if !ok {
		return nil, errors.ThrowError(ctx, CORE_ERROR_PROVIDER_NOT_FOUND, "Method Name", methodName)

	}

	//cast to a creatory func
	method, ok := methodInt.(InvokableMethod)
	if !ok {
		return nil, errors.ThrowError(ctx, CORE_ERROR_PROVIDER_NOT_FOUND, "Method Name", methodName)
	}

	log.Logger.Trace(ctx, "core.methods", "Returned method ", "Method Name", methodName)
	//return method
	return method, nil
}
