// +build appengine

package log

import (
	"bytes"
	glog "google.golang.org/appengine/log"
	"laatoosdk/context"
	logxi "logxi/v1"
	"os"
)

func NewLogger() LoggerInterface {
	return &StandaloneLogger{logxi.NewLoggerWithHandler(&AppEngineHandler{}, "default", logxi.NewJSONFormatter("default"))}
}

type StandaloneLogger struct {
	logger logxi.Logger
}

func (log *StandaloneLogger) Trace(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Trace(reqContext, msg, loggingCtx, args...)
}
func (log *StandaloneLogger) Debug(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Debug(reqContext, msg, loggingCtx, args...)
}
func (log *StandaloneLogger) Info(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Info(reqContext, msg, loggingCtx, args...)
}
func (log *StandaloneLogger) Warn(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Warn(reqContext, msg, loggingCtx, args...)
}
func (log *StandaloneLogger) Error(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Error(reqContext, msg, loggingCtx, args...)
}
func (log *StandaloneLogger) Fatal(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	log.logger.Fatal(reqContext, msg, loggingCtx, args...)
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

type AppEngineHandler struct {
}

func (ah *AppEngineHandler) WriteLog(ctx interface{}, loggingCtx string, buf *bytes.Buffer, level int, msg string, args []interface{}) {
	appengineContext := context.GetAppengineContext(ctx)
	if appengineContext != nil {
		switch level {
		case logxi.LevelTrace:
			glog.Debugf(appengineContext, buf.String())
		case logxi.LevelDebug:
			glog.Debugf(appengineContext, buf.String())
		case logxi.LevelInfo:
			glog.Infof(appengineContext, buf.String())
		case logxi.LevelWarn:
			glog.Warningf(appengineContext, buf.String())
		default:
			glog.Errorf(appengineContext, buf.String())
		}
	} else {
		buf.WriteTo(os.Stderr)
	}
}
