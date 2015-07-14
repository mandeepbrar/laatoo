package log

import (
	log "github.com/Sirupsen/logrus"
	"laatoo/config"
)

var Logger *log.Logger

func init() {
	Logger = log.New()
}

func ConfigLogger(conf config.Config) {
	//read configuration
	loggingLevel, err := log.ParseLevel(conf.GetConfig(config.LOGGINGLEVEL))
	if err == nil {
		//set logger level
		Logger.Level = loggingLevel
	} else {
		Logger.Errorf("Logger: Invalid logging level %s", err)
	}
}
