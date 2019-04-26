// +build !appengine

package log

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	slog "laatoo/sdk/server/log"
)

func GetLogger(ctx core.ServerContext, loggertype string, logformat string, loglevel int, name string, settings config.Config) components.Logger {
	slog.Error(ctx, "Initializing Logger  *************************************************", "loggertype", loggertype, "logformat", logformat, "loglevel", loglevel, "name", name, "settings", settings)
	var logger components.Logger
	if loggertype == CONF_STDERR_LOGGER {
		logger = NewStdLogger(ctx, name, settings)
	} else {
		logger = NewSysLogger(ctx, name, settings)
	}
	logger.SetFormat(logformat)
	logger.SetLevel(loglevel)
	return logger
}
