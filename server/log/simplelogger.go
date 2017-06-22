package log

import (
	"laatoo/sdk/components"
	"laatoo/sdk/core"
	slog "laatoo/sdk/log"
)

const (
	STR_TRACE = "Trace"
	STR_DEBUG = "Debug"
	STR_INFO  = "Info"
	STR_WARN  = "Warn"
	STR_ERROR = "Error"
	STR_FATAL = "Fatal"
)

type logPrinter func(ctx core.Context, app string, strlevel string, wh WriteHandler, level int, msg string, args ...interface{})

type WriteHandler interface {
	Print(ctx core.Context, app string, msg string, level int, strlevel string)
	PrintBytes(ctx core.Context, app string, msg []byte, level int, strlevel string) (int, error)
}

var (
	logFormats = make(map[string]logPrinter, 6)
)

func NewSimpleLogger(appname string, wh WriteHandler) components.Logger {
	logger := &SimpleLogger{format: "json", level: slog.INFO, wh: wh, app: "Laatoo"}
	logger.printer = printJSON
	return logger
}

type SimpleLogger struct {
	wh  WriteHandler
	app string
	//buffer bytes.Buffer
	format  string
	printer logPrinter
	level   int
}

func (log *SimpleLogger) Trace(ctx core.Context, msg string, args ...interface{}) {
	if log.level > slog.DEBUG {
		log.printer(ctx, log.app, STR_TRACE, log.wh, slog.TRACE, msg, args...)
	}
}
func (log *SimpleLogger) Debug(ctx core.Context, msg string, args ...interface{}) {
	if log.level > slog.INFO {
		log.printer(ctx, log.app, STR_DEBUG, log.wh, slog.DEBUG, msg, args...)
	}
}
func (log *SimpleLogger) Info(ctx core.Context, msg string, args ...interface{}) {
	if log.level > slog.WARN {
		log.printer(ctx, log.app, STR_INFO, log.wh, slog.INFO, msg, args...)
	}
}
func (log *SimpleLogger) Warn(ctx core.Context, msg string, args ...interface{}) {
	if log.level > slog.ERROR {
		log.printer(ctx, log.app, STR_WARN, log.wh, slog.WARN, msg, args...)
	}
}
func (log *SimpleLogger) Error(ctx core.Context, msg string, args ...interface{}) {
	if log.level > slog.FATAL {
		log.printer(ctx, log.app, STR_ERROR, log.wh, slog.ERROR, msg, args...)
	}
}
func (log *SimpleLogger) Fatal(ctx core.Context, msg string, args ...interface{}) {
	log.printer(ctx, log.app, STR_FATAL, log.wh, slog.FATAL, msg, args...)
}

func (log *SimpleLogger) SetFormat(format string) {
	log.format = format
	printer, ok := logFormats[format]
	if ok {
		log.printer = printer
	}
}

func (log *SimpleLogger) SetLevel(level int) {
	log.level = level
}
func (log *SimpleLogger) IsTrace() bool {
	return log.level == slog.TRACE
}
func (log *SimpleLogger) IsDebug() bool {
	return log.level == slog.DEBUG
}
func (log *SimpleLogger) IsInfo() bool {
	return log.level == slog.INFO
}
func (log *SimpleLogger) IsWarn() bool {
	return log.level == slog.WARN
}

func (log *SimpleLogger) Write(p []byte) (int, error) {
	return log.wh.PrintBytes(nil, log.app, p, slog.INFO, STR_INFO)
}
