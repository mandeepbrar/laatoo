// +build !appengine

package log

import (
	"laatoosdk/core"
)

func NewLogger() LoggerInterface {
	return &StandaloneLogger{NewLogxiLogger()}
}

type StandaloneLogger struct {
	logger LoggerInterface
}

func (log *StandaloneLogger) Trace(reqContext core.Context, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Trace(reqContext, loggingCtx, msg, args...)
}
func (log *StandaloneLogger) Debug(reqContext core.Context, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Debug(reqContext, loggingCtx, msg, args...)
}
func (log *StandaloneLogger) Info(reqContext core.Context, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Info(reqContext, loggingCtx, msg, args...)
}
func (log *StandaloneLogger) Warn(reqContext core.Context, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Warn(reqContext, loggingCtx, msg, args...)
}
func (log *StandaloneLogger) Error(reqContext core.Context, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Error(reqContext, loggingCtx, msg, args...)
}
func (log *StandaloneLogger) Fatal(reqContext core.Context, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Fatal(reqContext, loggingCtx, msg, args...)
}
func (log *StandaloneLogger) SetFormat(format string) {
	log.logger.SetFormat(format)
}

func (log *StandaloneLogger) SetType(loggertype string) {
	if loggertype == "logrus" {
		log.logger = NewLogrus()
	}
}

func (log *StandaloneLogger) SetLevel(level string) {
	log.logger.SetLevel(level)
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
