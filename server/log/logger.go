package log

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	slog "laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"laatoo/server/common"
	"laatoo/server/constants"
	"path"
)

const (
	CONF_LOGGINGLEVEL      = "level"
	CONF_LOGGER_TYPE       = "loggertype"
	CONF_LOGGER_SETTINGS   = "settings"
	CONF_APPLICATION       = "application"
	CONF_LOGGING_FORMAT    = "format"
	CONF_STDERR_LOGGER     = "stderr"
	CONF_SYS_LOGGER        = "sys"
	CONF_FMT_JSON          = "json"
	CONF_FMT_JSONMAX       = "jsonmax"
	CONF_FMT_HAPPY         = "happy"
	CONF_FMT_HAPPYMAX      = "happymax"
	CONF_FMT_HAPPYCOLOR    = "happycolor"
	CONF_FMT_HAPPYMAXCOLOR = "happymaxcolor"
	CONF_FMT_RFC5424       = "rfc5424"
)

type logger struct {
	proxy          elements.Logger
	loggerInstance components.Logger
	name           string
}

func (lgr *logger) Initialize(ctx core.ServerContext, conf config.Config) error {
	logconf, ok := conf.GetSubConfig(ctx, constants.CONF_LOGGING)
	if ok {
		loggerType, loggingFormat, loggingLevel, loggerSettings := processConf(ctx, logconf)
		lgr.loggerInstance = GetLogger(ctx, loggerType, loggingFormat, loggingLevel, lgr.name, loggerSettings)
		slog.Trace(ctx, "Logger initialized *************************************************")
	} else {
		baseDir, _ := ctx.GetString(config.BASEDIR)
		confFile := path.Join(baseDir, constants.CONF_LOGGING, constants.CONF_CONFIG_FILE)
		ok, _, _ = utils.FileExists(confFile)
		var err error
		if ok {
			if logconf, err = common.NewConfigFromFile(ctx, confFile, nil); err != nil {
				return errors.WrapError(ctx, err)
			}
			loggerType, loggingFormat, loggingLevel, loggerSettings := processConf(ctx, logconf)
			lgr.loggerInstance = GetLogger(ctx, loggerType, loggingFormat, loggingLevel, lgr.name, loggerSettings)
		}
	}
	return nil
}

func processConf(ctx core.ServerContext, logconf config.Config) (string, string, int, config.Config) {
	loggerType := CONF_STDERR_LOGGER
	loggingFormat := CONF_FMT_JSON
	loggingLevel := slog.TRACE
	val, ok := logconf.GetString(ctx, CONF_LOGGER_TYPE)
	if ok {
		loggerType = val
	}
	val, ok = logconf.GetString(ctx, CONF_LOGGING_FORMAT)
	if ok {
		loggingFormat = val
	}
	lLevel, ok := logconf.GetString(ctx, CONF_LOGGINGLEVEL)
	if ok {
		loggingLevel = GetLevel(lLevel)
	}
	loggerSettings, _ := logconf.GetSubConfig(ctx, CONF_LOGGER_SETTINGS)

	return loggerType, loggingFormat, loggingLevel, loggerSettings
}

func (lgr *logger) Start(ctx core.ServerContext) error {
	return nil
}
