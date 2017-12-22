package log

import (
	"laatoo/sdk/components"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	slog "laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/sdk/utils"
	"laatoo/server/common"
	"laatoo/server/constants"
	"path"
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
	proxy          server.Logger
	loggerInstance components.Logger
	name           string
}

func (lgr *logger) Initialize(ctx core.ServerContext, conf config.Config) error {
	logconf, ok := conf.GetSubConfig(ctx, constants.CONF_LOGGING)
	if ok {
		loggerType, loggingFormat, loggingLevel := processConf(ctx, logconf)
		lgr.loggerInstance = GetLogger(loggerType, loggingFormat, loggingLevel, lgr.name)
	} else {
		baseDir, _ := ctx.GetString(config.BASEDIR)
		confFile := path.Join(baseDir, constants.CONF_LOGGING, constants.CONF_CONFIG_FILE)
		ok, _, _ = utils.FileExists(confFile)
		var err error
		if ok {
			if logconf, err = common.NewConfigFromFile(ctx, confFile, nil); err != nil {
				return errors.WrapError(ctx, err)
			}
			loggerType, loggingFormat, loggingLevel := processConf(ctx, logconf)
			lgr.loggerInstance = GetLogger(loggerType, loggingFormat, loggingLevel, lgr.name)
		}
	}
	return nil
}

func processConf(ctx core.ServerContext, logconf config.Config) (string, string, int) {
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
	return loggerType, loggingFormat, loggingLevel
}

func (lgr *logger) Start(ctx core.ServerContext) error {
	return nil
}
