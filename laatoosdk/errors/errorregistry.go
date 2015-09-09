package errors

import (
	"fmt"
	"laatoosdk/log"
	"runtime/debug"
)

//levels at which error messages can be logged
const (
	FATAL   = iota
	PANIC   = iota
	ERROR   = iota
	WARNING = iota
	INFO    = iota
	DEBUG   = iota
)

//error that is registered for an error code
type Error struct {
	InternalErrorCode string
	Loglevel          int
	Error             error
}

//Error handler for interrupting the error process
//Returns true if the error has been handled
type ErrorHandler func(err *Error, ctx map[string]interface{}, info ...string) bool

var (
	//errors register to store all errors in the process
	ErrorsRegister = make(map[string]*Error, 50)
	//registered handlers for errors
	ErrorsHandlersRegister = make(map[string][]ErrorHandler, 20)
)

//register error code
func RegisterCode(internalErrorCode string, loglevel int, err error) {
	ErrorsRegister[internalErrorCode] = &Error{internalErrorCode, loglevel, err}
}

//register error handler for an internal error code
//handler will be called before throwing an error
//nil will be retured if an error is handled
func RegisterErrorHandler(internalErrorCode string, eh ErrorHandler) {
	val := ErrorsHandlersRegister[internalErrorCode]
	//add a new array of handlers if it doesnt exist already
	if val == nil {
		val = []ErrorHandler{}
	}
	//append the handler to the existing list and add to the map
	val = append(val, eh)
	ErrorsHandlersRegister[internalErrorCode] = val
}

//rethrow an error with an internal error code
func RethrowHttpError(internalErrorCode string, ctx interface{}, err error, info ...string) error {
	return RethrowErrorWithContext(internalErrorCode, map[string]interface{}{"Context": ctx}, err, info...)
}

//rethrow an error with an internal error code
func RethrowError(internalErrorCode string, err error, info ...string) error {
	return RethrowErrorWithContext(internalErrorCode, nil, err, info...)
}

//rethrow an error with an internal error code
func RethrowErrorWithContext(internalErrorCode string, ctx map[string]interface{}, err error, info ...string) error {
	return ThrowErrorWithContext(internalErrorCode, ctx, append(info, err.Error())...)
}

//rethrow an error with an internal error code
func ThrowHttpError(internalErrorCode string, ctx interface{}, info ...string) error {
	return ThrowErrorWithContext(internalErrorCode, map[string]interface{}{"Context": ctx}, info...)
}

func ThrowError(internalErrorCode string, info ...string) error {
	return ThrowErrorWithContext(internalErrorCode, nil, info...)
}

//throw a registered error code
func ThrowErrorWithContext(internalErrorCode string, ctx map[string]interface{}, info ...string) error {
	err, ok := ErrorsRegister[internalErrorCode]
	if !ok {
		panic(fmt.Errorf("Invalid error code: %s", internalErrorCode))
	}
	switch err.Loglevel {
	case FATAL:
		log.Logger.Fatalf("Encountered error: %s\n, Internal Error Code: %s, Context: %s, Info: %s", err.Error, err.InternalErrorCode, ctx, info)
	case PANIC:
		log.Logger.Panicf("Encountered error: %s\n, Internal Error Code: %s, Context: %s, Info: %s", err.Error, err.InternalErrorCode, ctx, info)
	case ERROR:
		stack := ""
		if ctx == nil {
			stack = string(debug.Stack())
		}
		log.Logger.Errorf("Encountered error: %s\n, Internal Error Code: %s, Context: %s, Info: %s Stack: %s", err.Error, err.InternalErrorCode, ctx, info, stack)
	case WARNING:
		log.Logger.Warningf("Encountered error: %s\n, Internal Error Code: %s, Context: %s, Info: %s", err.Error, err.InternalErrorCode, ctx, info)
	case INFO:
		log.Logger.Infof("Encountered error: %s\n, Internal Error Code: %s, Context: %s, Info: %s", err.Error, err.InternalErrorCode, ctx, info)
	case DEBUG:
		log.Logger.Debugf("Encountered error: %s\n, Internal Error Code: %s, Context: %s, Info: %s", err.Error, err.InternalErrorCode, ctx, info)
	}
	//call the handlers while throwing an error
	handlers := ErrorsHandlersRegister[internalErrorCode]
	if handlers != nil {
		handled := false
		for _, val := range handlers {
			handled = val(err, ctx, info...) || handled
		}
		//if an error has been handled, dont throw it
		if handled {
			return nil
		}
	}
	//thwo the error
	return err.Error
}
