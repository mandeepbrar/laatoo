package errors

import (
	"fmt"
	"laatoosdk/core"
	"laatoosdk/log"
	"runtime/debug"
)

//levels at which error messages can be logged
const (
	FATAL   = iota
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
	Context           string
}

var ShowStack = false

//Error handler for interrupting the error process
//Returns true if the error has been handled
type ErrorHandler func(ctx core.Context, err *Error, info ...interface{}) bool

var (
	//errors register to store all errors in the process
	ErrorsRegister = make(map[string]*Error, 50)
	//registered handlers for errors
	ErrorsHandlersRegister = make(map[string][]ErrorHandler, 20)
)

//register error code
func RegisterCode(internalErrorCode string, loglevel int, err error, ctx string) {
	ErrorsRegister[internalErrorCode] = &Error{internalErrorCode, loglevel, err, ctx}
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

func ThrowError(ctx core.Context, internalErrorCode string, info ...interface{}) error {
	return RethrowError(ctx, internalErrorCode, nil, info...)
}

func ThrowErrorInCtx(ctx core.Context, loggingCtx string, internalErrorCode string, info ...interface{}) error {
	return RethrowErrorInCtx(ctx, loggingCtx, internalErrorCode, nil, info...)
}

func RethrowError(ctx core.Context, internalErrorCode string, err error, info ...interface{}) error {
	registeredErr, ok := ErrorsRegister[internalErrorCode]
	if !ok {
		panic(fmt.Errorf("Invalid error code: %s", internalErrorCode))
	}
	return throwErrorInCtx(ctx, registeredErr.Context, registeredErr, err, info...)
}

func RethrowErrorInCtx(ctx core.Context, loggingCtx string, internalErrorCode string, err error, info ...interface{}) error {
	registeredErr, ok := ErrorsRegister[internalErrorCode]
	if !ok {
		panic(fmt.Errorf("Invalid error code: %s", internalErrorCode))
	}
	return throwErrorInCtx(ctx, loggingCtx, registeredErr, err, info...)
}

//throw a registered error code
//rethrow an error with an internal error code
func throwErrorInCtx(ctx core.Context, loggingCtx string, registeredError *Error, rethrownError error, info ...interface{}) error {
	var errDetails []interface{}
	if rethrownError == nil {
		errDetails = []interface{}{"Err", registeredError.Error.Error(), "Internal Error Code", registeredError.InternalErrorCode}
	} else {
		errDetails = []interface{}{"Err", registeredError.Error.Error(), "Internal Error Code", registeredError.InternalErrorCode, "Root Error", rethrownError}
	}
	infoArr := append(errDetails, info...)
	switch registeredError.Loglevel {
	case FATAL:
		log.Logger.Fatal(ctx, loggingCtx, "Encountered error", infoArr...)
	case ERROR:
		log.Logger.Error(ctx, loggingCtx, "Encountered error", infoArr...)
	case WARNING:
		log.Logger.Warn(ctx, loggingCtx, "Encountered warning", infoArr...)
	case INFO:
		log.Logger.Info(ctx, loggingCtx, "Info Error", infoArr...)
	case DEBUG:
		log.Logger.Debug(ctx, loggingCtx, "Debug Error", infoArr...)
	}
	if ShowStack {
		log.Logger.Debug(ctx, loggingCtx, string(debug.Stack()))
	}
	//call the handlers while throwing an error
	handlers := ErrorsHandlersRegister[registeredError.InternalErrorCode]
	if handlers != nil {
		handled := false
		for _, val := range handlers {
			handled = val(ctx, registeredError, info...) || handled
		}
		//if an error has been handled, dont throw it
		if handled {
			return nil
		}
	}
	//thwo the error
	return registeredError.Error
}
