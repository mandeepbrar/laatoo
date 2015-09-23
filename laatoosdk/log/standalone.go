package log

import (
	//log "github.com/Sirupsen/logrus"
	logxi "github.com/mgutz/logxi/v1"
	"os"
)

func NewLogger() LoggerInterface {
	return &StandaloneLogger{logxi.NewLogger3(os.Stdout, "default", logxi.NewJSONFormatter("default"))}
}

type StandaloneLogger struct {
	logger logxi.Logger
}

func (log *StandaloneLogger) Trace(loggingCtx string, msg string, args ...interface{}) {
	log.logger.Trace(msg, "_ctx", loggingCtx, args)
}
func (log *StandaloneLogger) Debug(loggingCtx string, msg string, args ...interface{}) {
	log.logger.Debug(msg, "_ctx", loggingCtx, args)
}
func (log *StandaloneLogger) Info(loggingCtx string, msg string, args ...interface{}) {
	log.logger.Info(msg, "_ctx", loggingCtx, args)
}
func (log *StandaloneLogger) Warn(loggingCtx string, msg string, args ...interface{}) {
	log.logger.Warn(msg, "_ctx", loggingCtx, args)
}
func (log *StandaloneLogger) Error(loggingCtx string, msg string, args ...interface{}) {
	log.logger.Error(msg, "_ctx", loggingCtx, args)
}
func (log *StandaloneLogger) Fatal(loggingCtx string, msg string, args ...interface{}) {
	log.logger.Fatal(msg, "_ctx", loggingCtx, args)
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
