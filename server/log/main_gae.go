// +build appengine

package log

import slog "laatoo/sdk/log"

func GetLogger(loggertype string, logformat string, loglevel int, name string) slog.Logger {
	logger := NewGaeLogger(name)
	logger.SetFormat(logformat)
	logger.SetLevel(loglevel)
	return logger
}
