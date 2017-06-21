package log

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	slog "laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/constants"
)

const (
	CONF_LOGGINGLEVEL      = "level"
	CONF_LOGGER_TYPE       = "loggertype"
	CONF_APPLICATION       = "application"
	CONF_LOGGING_FORMAT    = "format"
	CONF_STDERR_LOGGER     = "stderr"
	CONF_SYS_LOGGER        = "stderr"
	CONF_FMT_JSON          = "json"
	CONF_FMT_JSONMAX       = "jsonmax"
	CONF_FMT_HAPPY         = "happy"
	CONF_FMT_HAPPYMAX      = "happymax"
	CONF_FMT_HAPPYCOLOR    = "happycolor"
	CONF_FMT_HAPPYMAXCOLOR = "happymaxcolor"
)

type logger struct {
	parent         core.ServerElement
	proxy          server.Logger
	loggerInstance slog.Logger
	name           string
}

func (lgr *logger) Initialize(ctx core.ServerContext, conf config.Config) error {
	logconf, ok := conf.GetSubConfig(constants.CONF_LOGGING)
	if ok {
		loggerType := CONF_STDERR_LOGGER
		loggingFormat := CONF_FMT_JSON
		loggingLevel := slog.INFO
		val, ok := logconf.GetString(CONF_LOGGER_TYPE)
		if ok {
			loggerType = val
		}
		val, ok = logconf.GetString(CONF_LOGGING_FORMAT)
		if ok {
			loggingFormat = val
		}
		lLevel, ok := logconf.GetString(CONF_LOGGINGLEVEL)
		if ok {
			loggingLevel = GetLevel(lLevel)
		}
		lgr.loggerInstance = GetLogger(loggerType, loggingFormat, loggingLevel, lgr.name)
	}
	return nil
}

func (lgr *logger) Start(ctx core.ServerContext) error {
	return nil
}
