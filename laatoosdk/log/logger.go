package log

import (
	log "github.com/Sirupsen/logrus"
	"laatoosdk/config"
)

const (
	CONF_LOGGINGLEVEL = "logging.level"
)

var Logger *log.Logger

func init() {
	Logger = log.New()
}

func ConfigLogger(conf config.Config) {
	//read configuration
	loggingLevel, err := log.ParseLevel(conf.GetConfig(CONF_LOGGINGLEVEL))
	if err == nil {
		//set logger level
		Logger.Level = loggingLevel
	} else {
		Logger.Level = log.InfoLevel
		Logger.Errorf("Logger: Invalid logging level %s", err)
	}
}
