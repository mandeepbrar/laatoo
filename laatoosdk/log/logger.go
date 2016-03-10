package log

import (
	"laatoosdk/config"
	"laatoosdk/core"
)

const (
	CONF_LOGGINGLEVEL   = "logging.level"
	CONF_LOGGER         = "logging.loggertype"
	CONF_LOGGING_FORMAT = "logging.format"
)

type LoggerInterface interface {
	Trace(reqContext core.Context, loggingCtx string, msg string, args ...interface{})
	Debug(reqContext core.Context, loggingCtx string, msg string, args ...interface{})
	Info(reqContext core.Context, loggingCtx string, msg string, args ...interface{})
	Warn(reqContext core.Context, loggingCtx string, msg string, args ...interface{})
	Error(reqContext core.Context, loggingCtx string, msg string, args ...interface{})
	Fatal(reqContext core.Context, loggingCtx string, msg string, args ...interface{})

	SetLevel(string)
	SetType(string)
	SetFormat(string)
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
	loggingFormat := conf.GetString(CONF_LOGGING_FORMAT)
	loggerType := conf.GetString(CONF_LOGGER)
	if loggerType != "" {
		Logger.SetType(loggerType)
	}
	Logger.SetLevel(loggingLevel)
	Logger.SetFormat(loggingFormat)
	return Logger.IsDebug()
}
