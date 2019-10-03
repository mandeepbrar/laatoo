package log

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

type loggerProxy struct {
	logger *logger
}

func (proxy *loggerProxy) Reference() core.ServerElement {
	return &loggerProxy{logger: proxy.logger}
}
func (proxy *loggerProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *loggerProxy) GetName() string {
	return proxy.logger.name
}
func (proxy *loggerProxy) GetType() core.ServerElementType {
	return core.ServerElementLogger
}
func (proxy *loggerProxy) GetContext() core.ServerContext {
	return proxy.logger.svrContext
}

func (lgr *loggerProxy) Trace(reqContext ctx.Context, msg string, args ...interface{}) {
	lgr.logger.loggerInstance.Trace(reqContext, msg, args...)
}
func (lgr *loggerProxy) Debug(reqContext ctx.Context, msg string, args ...interface{}) {
	lgr.logger.loggerInstance.Debug(reqContext, msg, args...)
}
func (lgr *loggerProxy) Info(reqContext ctx.Context, msg string, args ...interface{}) {
	lgr.logger.loggerInstance.Info(reqContext, msg, args...)
}
func (lgr *loggerProxy) Warn(reqContext ctx.Context, msg string, args ...interface{}) {
	lgr.logger.loggerInstance.Warn(reqContext, msg, args...)
}
func (lgr *loggerProxy) Error(reqContext ctx.Context, msg string, args ...interface{}) {
	lgr.logger.loggerInstance.Error(reqContext, msg, args...)
}
func (lgr *loggerProxy) Fatal(reqContext ctx.Context, msg string, args ...interface{}) {
	lgr.logger.loggerInstance.Fatal(reqContext, msg, args...)
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
