package log

import (
	"laatoosdk/config"
)

const (
	CONF_LOGGINGLEVEL = "logging.level"
	CONF_LOGGER       = "logging.loggertype"
)

type LoggerInterface interface {
	Trace(loggingCtx string, msg string, args ...interface{})
	Debug(loggingCtx string, msg string, args ...interface{})
	Info(loggingCtx string, msg string, args ...interface{})
	Warn(loggingCtx string, msg string, args ...interface{})
	Error(loggingCtx string, msg string, args ...interface{})
	Fatal(loggingCtx string, msg string, args ...interface{})

	SetLevel(string)
	IsTrace() bool
	IsDebug() bool
	IsInfo() bool
	IsWarn() bool
}

var (
	Logger LoggerInterface
)

//configures logging level and returns true if its set to debug
func ConfigLogger(conf config.Config) bool {
	Logger = NewLogger()
	loggingLevel := conf.GetString(CONF_LOGGINGLEVEL)
	Logger.SetLevel(loggingLevel)
	return Logger.IsDebug()
}
