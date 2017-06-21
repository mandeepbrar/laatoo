// +build !appengine

package log

import slog "laatoo/sdk/log"

func GetLogger(loggertype string, logformat string, loglevel int, name string) slog.Logger {
	var logger slog.Logger
	if loggertype == CONF_STDERR_LOGGER {
		logger = NewStdLogger(name)
	} else {
		logger = NewSysLogger(name)
	}
	logger.SetFormat(logformat)
	logger.SetLevel(loglevel)
	return logger
}
