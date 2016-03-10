// +build !appengine

package log

import (
	"github.com/Sirupsen/logrus"
	"laatoosdk/core"
)

func NewLogrus() LoggerInterface {
	return &LogrusLogger{logrus.New()}
}

type LogrusLogger struct {
	logger *logrus.Logger
}

func (log *LogrusLogger) Trace(reqContext core.Context, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Debug([]interface{}{loggingCtx, msg, args})
}
func (log *LogrusLogger) Debug(reqContext core.Context, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Debug([]interface{}{loggingCtx, msg, args})
}
func (log *LogrusLogger) Info(reqContext core.Context, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Info([]interface{}{loggingCtx, msg, args})
}
func (log *LogrusLogger) Warn(reqContext core.Context, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Warn([]interface{}{loggingCtx, msg, args})
}
func (log *LogrusLogger) Error(reqContext core.Context, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Error([]interface{}{loggingCtx, msg, args})
}
func (log *LogrusLogger) Fatal(reqContext core.Context, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Fatal([]interface{}{loggingCtx, msg, args})
}
func (log *LogrusLogger) SetFormat(format string) {
	switch format {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "happy":
	}
}
func (log *LogrusLogger) SetType(logtype string) {

}

func (log *LogrusLogger) SetLevel(level string) {
	switch level {
	case "all":
		log.logger.Level = logrus.DebugLevel
	case "trace":
		log.logger.Level = logrus.DebugLevel
	case "debug":
		log.logger.Level = logrus.DebugLevel
	case "info":
		log.logger.Level = logrus.InfoLevel
	case "warn":
		log.logger.Level = logrus.WarnLevel
	default:
		log.logger.Level = logrus.ErrorLevel
	}
}
func (log *LogrusLogger) IsTrace() bool {
	return log.logger.Level == logrus.DebugLevel
}
func (log *LogrusLogger) IsDebug() bool {
	return log.logger.Level == logrus.DebugLevel
}
func (log *LogrusLogger) IsInfo() bool {
	return log.logger.Level == logrus.InfoLevel
}
func (log *LogrusLogger) IsWarn() bool {
	return log.logger.Level == logrus.WarnLevel
}
