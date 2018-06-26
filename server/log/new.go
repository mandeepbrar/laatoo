package log

import (
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	slog "laatoo/sdk/server/log"
)

func NewLogger(ctx core.ServerContext, name string) (*logger, *loggerProxy) {
	logger := &logger{name: name, loggerInstance: GetLogger(CONF_STDERR_LOGGER, CONF_FMT_JSON, slog.TRACE, "")}
	loggerElem := &loggerProxy{logger: logger}
	logger.proxy = loggerElem
	return logger, loggerElem
}

/*
func ChildLogger(ctx core.ServerContext, name string, parentLogger core.ServerElement, parentElem core.ServerElement, conf config.Config, filters ...elements.Filter) (*logger, *loggerProxy) {
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
*/

func ChildLogger(ctx core.ServerContext, name string, parentLogger core.ServerElement) (elements.ServerElementHandle, elements.Logger) {
	var loggerInstance components.Logger
	if parentLogger != nil {
		loggerInstance = parentLogger.(*loggerProxy).logger.loggerInstance
	}
	childLogger := &logger{name: name, loggerInstance: loggerInstance}
	childLoggerElem := &loggerProxy{logger: childLogger}
	childLogger.proxy = childLoggerElem
	return childLogger, childLoggerElem
}
