// +build !appengine

package log

import "laatoo/sdk/components"

func GetLogger(loggertype string, logformat string, loglevel int, name string) components.Logger {
	var logger components.Logger
	if loggertype == CONF_STDERR_LOGGER {
		logger = NewStdLogger(name)
	} else {
		logger = NewSysLogger(name)
	}
	logger.SetFormat(logformat)
	logger.SetLevel(loglevel)
	return logger
}
