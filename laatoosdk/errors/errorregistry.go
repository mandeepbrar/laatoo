package errors

import (
	"fmt"
	"github.com/labstack/echo"
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

var ShowStack = true

//Error handler for interrupting the error process
//Returns true if the error has been handled
type ErrorHandler func(ctx *echo.Context, err *Error, info ...interface{}) bool

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

func ThrowError(ctx *echo.Context, internalErrorCode string, info ...interface{}) error {
	return RethrowError(ctx, internalErrorCode, nil, info...)
}

//throw a registered error code
//rethrow an error with an internal error code
func RethrowError(ctx *echo.Context, internalErrorCode string, err error, info ...interface{}) error {
	registeredErr, ok := ErrorsRegister[internalErrorCode]
	if !ok {
		panic(fmt.Errorf("Invalid error code: %s", internalErrorCode))
	}
	stack := ""
	if ShowStack {
		stack = string(debug.Stack())
	}
	var errDetails []interface{}
	if err == nil {
		errDetails = []interface{}{"Err", registeredErr.Error.Error(), "Internal Error Code", internalErrorCode, "Stack", stack}
	} else {
		errDetails = []interface{}{"Err", registeredErr.Error.Error(), "Internal Error Code", internalErrorCode, "Stack", stack, "Root Error", err}
	}
	infoArr := append(errDetails, info)
	switch registeredErr.Loglevel {
	case FATAL:
		log.Logger.Fatal(ctx, registeredErr.Context, "Encountered error", infoArr...)
	case ERROR:
		log.Logger.Error(ctx, registeredErr.Context, "Encountered error", infoArr...)
	case WARNING:
		log.Logger.Warn(ctx, registeredErr.Context, "Encountered warning", infoArr...)
	case INFO:
		log.Logger.Info(ctx, registeredErr.Context, "Info Error", infoArr...)
	case DEBUG:
		log.Logger.Debug(ctx, registeredErr.Context, "Debug Error", infoArr...)
	}
	//call the handlers while throwing an error
	handlers := ErrorsHandlersRegister[internalErrorCode]
	if handlers != nil {
		handled := false
		for _, val := range handlers {
			handled = val(ctx, registeredErr, info...) || handled
		}
		//if an error has been handled, dont throw it
		if handled {
			return nil
		}
	}
	//thwo the error
	return registeredErr.Error
}
