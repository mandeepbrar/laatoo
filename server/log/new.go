package log

import (
	"laatoo/sdk/components"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	slog "laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/constants"
)

func NewLogger(ctx core.ServerContext, name string) (*logger, *loggerProxy) {
	logger := &logger{name: name, loggerInstance: GetLogger(CONF_STDERR_LOGGER, CONF_FMT_JSON, slog.TRACE, "")}
	loggerElem := &loggerProxy{logger: logger}
	logger.proxy = loggerElem
	return logger, loggerElem
}

func ChildLoggerWithConf(ctx core.ServerContext, name string, parentLogger core.ServerElement, parentElem core.ServerElement, conf config.Config) (*logger, *loggerProxy) {
	var loggerInstance components.Logger
	logconf, ok := conf.GetSubConfig(constants.CONF_LOGGING)
	if ok {
		loggerType, loggingFormat, loggingLevel := processConf(ctx, logconf)
		loggerInstance = GetLogger(loggerType, loggingFormat, loggingLevel, name)
	} else {
		if parentLogger != nil {
			loggerInstance = parentLogger.(*loggerProxy).logger.loggerInstance
		}
	}
	logger := &logger{name: name, loggerInstance: loggerInstance}
	loggerElem := &loggerProxy{logger: logger}
	logger.proxy = loggerElem
	return logger, loggerElem
}

func ChildLogger(ctx core.ServerContext, name string, parentLogger core.ServerElement, filters ...server.Filter) (server.ServerElementHandle, server.Logger) {
	var loggerInstance components.Logger
	if parentLogger != nil {
		loggerInstance = parentLogger.(*loggerProxy).logger.loggerInstance
	}
	childLogger := &logger{name: name, loggerInstance: loggerInstance}
	childLoggerElem := &loggerProxy{logger: childLogger}
	childLogger.proxy = childLoggerElem
	return childLogger, childLoggerElem
}
