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

//configures logging level and returns true if its set to debug
func ConfigLogger(conf config.Config) bool {
	//read configuration
	loggingLevel, err := log.ParseLevel(conf.GetString(CONF_LOGGINGLEVEL))
	if err == nil {
		//set logger level
		Logger.Level = loggingLevel
		if loggingLevel == log.DebugLevel {
			return true
		}
		return false
	} else {
		Logger.Level = log.InfoLevel
		Logger.Errorf("Logger: Invalid logging level %s", err)
		return false
	}
}
