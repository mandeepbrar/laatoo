package errors

import (
	"fmt"
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

//Error handler for interrupting the error process
//Returns true if the error has been handled
type ErrorHandler func(err *Error, ctx interface{}, info ...interface{}) bool

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

//rethrow an error with an internal error code
func RethrowHttpError(internalErrorCode string, ctx interface{}, err error, info ...interface{}) error {
	return RethrowErrorWithContext(internalErrorCode, ctx, err, info...)
}

//rethrow an error with an internal error code
func RethrowError(internalErrorCode string, err error, info ...interface{}) error {
	return RethrowErrorWithContext(internalErrorCode, nil, err, info...)
}

//rethrow an error with an internal error code
func RethrowErrorWithContext(internalErrorCode string, ctx interface{}, err error, info ...interface{}) error {
	return ThrowErrorWithContext(internalErrorCode, ctx, append([]interface{}{"Root Error", err.Error()}, info...)...)
}

//rethrow an error with an internal error code
func ThrowHttpError(internalErrorCode string, ctx interface{}, info ...interface{}) error {
	return ThrowErrorWithContext(internalErrorCode, ctx, info...)
}

func ThrowError(internalErrorCode string, info ...interface{}) error {
	return ThrowErrorWithContext(internalErrorCode, nil, info...)
}

//throw a registered error code
func ThrowErrorWithContext(internalErrorCode string, ctx interface{}, info ...interface{}) error {
	err, ok := ErrorsRegister[internalErrorCode]
	if !ok {
		panic(fmt.Errorf("Invalid error code: %s", internalErrorCode))
	}
	stack := ""
	if ctx == nil {
		stack = string(debug.Stack())
	}
	infoArr := append([]interface{}{"Err", err.Error.Error(), "Internal Error Code", err.InternalErrorCode, "Context", fmt.Sprint(ctx), "Stack", stack}, info...)
	switch err.Loglevel {
	case FATAL:
		log.Logger.Fatal(err.Context, "Encountered error", infoArr...)
	case ERROR:
		log.Logger.Error(err.Context, "Encountered error", infoArr...)
	case WARNING:
		log.Logger.Warn(err.Context, "Encountered warning", infoArr)
	case INFO:
		log.Logger.Info(err.Context, "Info Error", infoArr)
	case DEBUG:
		log.Logger.Debug(err.Context, "Debug Error", infoArr)
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
