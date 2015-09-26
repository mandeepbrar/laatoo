// +build !appengine

package log

import (
	//log "github.com/Sirupsen/logrus"
	logxi "logxi/v1"
	"os"
)

func NewLogger() LoggerInterface {
	return &StandaloneLogger{logxi.NewLogger3(os.Stdout, "default", logxi.NewJSONFormatter("default"))}
}

type StandaloneLogger struct {
	logger logxi.Logger
}

func (log *StandaloneLogger) Trace(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Trace(reqContext, loggingCtx, msg, args...)
}
func (log *StandaloneLogger) Debug(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Debug(reqContext, loggingCtx, msg, args...)
}
func (log *StandaloneLogger) Info(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Info(reqContext, loggingCtx, msg, args...)
}
func (log *StandaloneLogger) Warn(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Warn(reqContext, loggingCtx, msg, args...)
}
func (log *StandaloneLogger) Error(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Error(reqContext, loggingCtx, msg, args...)
}
func (log *StandaloneLogger) Fatal(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Fatal(reqContext, loggingCtx, msg, args...)
}

func (log *StandaloneLogger) SetLevel(level string) {
	switch level {
	case "all":
		log.logger.SetLevel(logxi.LevelAll)
	case "debug":
		log.logger.SetLevel(logxi.LevelDebug)
	case "info":
		log.logger.SetLevel(logxi.LevelInfo)
	case "warn":
		log.logger.SetLevel(logxi.LevelWarn)
	case "error":
		log.logger.SetLevel(logxi.LevelError)
	}
}
func (log *StandaloneLogger) IsTrace() bool {
	return log.logger.IsTrace()
}
func (log *StandaloneLogger) IsDebug() bool {
	return log.logger.IsDebug()
}
func (log *StandaloneLogger) IsInfo() bool {
	return log.logger.IsInfo()
}
func (log *StandaloneLogger) IsWarn() bool {
	return log.logger.IsWarn()
}
