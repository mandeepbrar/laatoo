// +build !appengine

package log

import (
	"laatoo/sdk/core"
	"os"
)

func NewLogger() LoggerInterface {
	return &StandaloneLogger{logger: NewSimpleLogger(stdSimpleLogsHandler()), level: TRACE}
}

type StandaloneLogger struct {
	logger LoggerInterface
	level  int
}

func (log *StandaloneLogger) Trace(reqContext core.Context, msg string, args ...interface{}) {
	log.logger.Trace(reqContext, msg, args...)
}
func (log *StandaloneLogger) Debug(reqContext core.Context, msg string, args ...interface{}) {
	log.logger.Debug(reqContext, msg, args...)
}
func (log *StandaloneLogger) Info(reqContext core.Context, msg string, args ...interface{}) {
	log.logger.Info(reqContext, msg, args...)
}
func (log *StandaloneLogger) Warn(reqContext core.Context, msg string, args ...interface{}) {
	log.logger.Warn(reqContext, msg, args...)
}
func (log *StandaloneLogger) Error(reqContext core.Context, msg string, args ...interface{}) {
	log.logger.Error(reqContext, msg, args...)
}
func (log *StandaloneLogger) Fatal(reqContext core.Context, msg string, args ...interface{}) {
	log.logger.Fatal(reqContext, msg, args...)
}
func (log *StandaloneLogger) SetFormat(format string) {
	log.logger.SetFormat(format)
}

func (log *StandaloneLogger) Write(p []byte) (n int, err error) {
	return log.logger.Write(p)
}

func (log *StandaloneLogger) SetType(loggertype string) {
	/*if loggertype == "logrus" {
		log.logger = NewLogrus()
	}
	if loggertype == "logxi" {
		log.logger = NewLogxiLogger()
	}*/
}

func (log *StandaloneLogger) SetLevel(level int) {
	log.level = level
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

func stdSimpleLogsHandler() SimpleWriteHandler {
	wh := &StdSimpleWriteHandler{}
	return wh
}

type StdSimpleWriteHandler struct {
}

func (jh *StdSimpleWriteHandler) Print(msg string) {
	os.Stderr.WriteString(msg)
}
func (jh *StdSimpleWriteHandler) PrintBytes(msg []byte) (int, error) {
	return os.Stderr.Write(msg)
}
