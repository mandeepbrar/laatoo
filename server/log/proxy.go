package log

import (
	"laatoo/sdk/core"
	"laatoo/server/common"
)

type loggerProxy struct {
	*common.Context
	logger *logger
}

func (lgr *loggerProxy) Trace(reqContext core.Context, msg string, args ...interface{}) {
	lgr.logger.loggerInstance.Trace(reqContext, msg, args...)
}
func (lgr *loggerProxy) Debug(reqContext core.Context, msg string, args ...interface{}) {
	lgr.logger.loggerInstance.Trace(reqContext, msg, args...)
}
func (lgr *loggerProxy) Info(reqContext core.Context, msg string, args ...interface{}) {
	lgr.logger.loggerInstance.Trace(reqContext, msg, args...)
}
func (lgr *loggerProxy) Warn(reqContext core.Context, msg string, args ...interface{}) {
	lgr.logger.loggerInstance.Trace(reqContext, msg, args...)
}
func (lgr *loggerProxy) Error(reqContext core.Context, msg string, args ...interface{}) {
	lgr.logger.loggerInstance.Trace(reqContext, msg, args...)
}
func (lgr *loggerProxy) Fatal(reqContext core.Context, msg string, args ...interface{}) {
	lgr.logger.loggerInstance.Trace(reqContext, msg, args...)
}

func (lgr *loggerProxy) SetLevel(lvl int) {
	lgr.logger.loggerInstance.SetLevel(lvl)
}

func (lgr *loggerProxy) SetFormat(fmt string) {
	lgr.logger.loggerInstance.SetFormat(fmt)
}
func (lgr *loggerProxy) IsTrace() bool {
	return lgr.logger.loggerInstance.IsTrace()
}
func (lgr *loggerProxy) IsDebug() bool {
	return lgr.logger.loggerInstance.IsDebug()
}
func (lgr *loggerProxy) IsInfo() bool {
	return lgr.logger.loggerInstance.IsInfo()
}
func (lgr *loggerProxy) IsWarn() bool {
	return lgr.logger.loggerInstance.IsWarn()
}
