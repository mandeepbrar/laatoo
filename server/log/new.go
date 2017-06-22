package log

import (
	"laatoo/sdk/components"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	slog "laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/common"
	"laatoo/server/constants"
)

func NewLogger(ctx core.ServerContext, name string, parentElem core.ServerElement) (*logger, *loggerProxy) {
	logger := &logger{parent: parentElem, name: name, loggerInstance: GetLogger(CONF_STDERR_LOGGER, CONF_FMT_JSON, slog.TRACE, "")}
	loggerCtx := parentElem.NewCtx(name)
	loggerElem := &loggerProxy{Context: loggerCtx.(*common.Context), logger: logger}
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
	logger := &logger{parent: parentElem, name: name, loggerInstance: loggerInstance}
	loggerCtx := parentElem.NewCtx(name)
	loggerElem := &loggerProxy{Context: loggerCtx.(*common.Context), logger: logger}
	logger.proxy = loggerElem
	return logger, loggerElem
}

func ChildLogger(ctx core.ServerContext, name string, parentLogger core.ServerElement, parent core.ServerElement, filters ...server.Filter) (server.ServerElementHandle, server.Logger) {
	var loggerInstance components.Logger
	if parentLogger != nil {
		loggerInstance = parentLogger.(*loggerProxy).logger.loggerInstance
	}
	childLogger := &logger{parent: parent, name: name, loggerInstance: loggerInstance}
	childLoggerCtx := parentLogger.NewCtx(name)
	childLoggerElem := &loggerProxy{Context: childLoggerCtx.(*common.Context), logger: childLogger}
	childLogger.proxy = childLoggerElem
	return childLogger, childLoggerElem
}
