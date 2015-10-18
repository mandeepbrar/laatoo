package log

import (
	"github.com/labstack/echo"
	"laatoosdk/config"
)

const (
	CONF_LOGGINGLEVEL = "logging.level"
	CONF_LOGGER       = "logging.loggertype"
)

type LoggerInterface interface {
	Trace(reqContext *echo.Context, loggingCtx string, msg string, args ...interface{})
	Debug(reqContext *echo.Context, loggingCtx string, msg string, args ...interface{})
	Info(reqContext *echo.Context, loggingCtx string, msg string, args ...interface{})
	Warn(reqContext *echo.Context, loggingCtx string, msg string, args ...interface{})
	Error(reqContext *echo.Context, loggingCtx string, msg string, args ...interface{})
	Fatal(reqContext *echo.Context, loggingCtx string, msg string, args ...interface{})

	SetLevel(string)
	IsTrace() bool
	IsDebug() bool
	IsInfo() bool
	IsWarn() bool
}

var (
	Logger LoggerInterface
)

func init() {
	Logger = NewLogger()
}

//configures logging level and returns true if its set to debug
func ConfigLogger(conf config.Config) bool {
	loggingLevel := conf.GetString(CONF_LOGGINGLEVEL)
	Logger.SetLevel(loggingLevel)
	return Logger.IsDebug()
}
