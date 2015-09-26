package log

import (
	"bytes"
	"fmt"
	"io"
)

// DefaultLogger is the default logger for this package.
type DefaultLogger struct {
	name         string
	level        int
	formatter    Formatter
	writeHandler WriteHandler
}

// NewLogger creates a new default logger. If writer is not concurrent
// safe, wrap it with NewConcurrentWriter.
func NewLogger(writer io.Writer, name string) Logger {
	formatter, err := createFormatter(name, logxiFormat)
	if err != nil {
		panic("Could not create formatter")
	}
	return NewLogger3(writer, name, formatter)
}

// NewLogger3 creates a new logger with a writer, name and formatter. If writer is not concurrent
// safe, wrap it with NewConcurrentWriter.
func NewLoggerWithHandler(writeHandler WriteHandler, name string, formatter Formatter) Logger {
	var level int
	if name != "__logxi" {
		// if err is returned, then it means the log is disabled
		level = getLogLevel(name)
		if level == LevelOff {
			return NullLog
		}
	}

	log := &DefaultLogger{
		formatter:    formatter,
		writeHandler: writeHandler,
		name:         name,
		level:        level,
	}

	// TODO loggers will be used when watching changes to configuration such
	// as in consul, etcd
	loggers.Lock()
	loggers.loggers[name] = log
	loggers.Unlock()
	return log
}

func NewLogger3(writer io.Writer, name string, formatter Formatter) Logger {
	return NewLoggerWithHandler(&DefaultWriteHandler{writer}, name, formatter)
}

// New creates a colorable default logger.
func New(name string) Logger {
	return NewLogger(colorableStdout, name)
}

// Trace logs a debug entry.
func (l *DefaultLogger) Trace(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	l.Log(reqContext, loggingCtx, LevelTrace, msg, args)
}

// Debug logs a debug entry.
func (l *DefaultLogger) Debug(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	l.Log(reqContext, loggingCtx, LevelDebug, msg, args)
}

// Info logs an info entry.
func (l *DefaultLogger) Info(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	l.Log(reqContext, loggingCtx, LevelInfo, msg, args)
}

// Warn logs a warn entry.
func (l *DefaultLogger) Warn(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) error {
	if l.IsWarn() {
		defer l.Log(reqContext, loggingCtx, LevelWarn, msg, args)

		for _, arg := range args {
			if err, ok := arg.(error); ok {
				return err
			}
		}

		return nil
	}
	return nil
}

func (l *DefaultLogger) extractLogError(reqContext interface{}, loggingCtx string, level int, msg string, args []interface{}) error {
	defer l.Log(reqContext, loggingCtx, level, msg, args)

	for _, arg := range args {
		if err, ok := arg.(error); ok {
			return err
		}
	}
	return fmt.Errorf(msg)
}

// Error logs an error entry.
func (l *DefaultLogger) Error(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) error {
	return l.extractLogError(reqContext, loggingCtx, LevelError, msg, args)
}

// Fatal logs a fatal entry then panics.
func (l *DefaultLogger) Fatal(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	l.extractLogError(reqContext, loggingCtx, LevelFatal, msg, args)
	//defer panic("Exit due to fatal error: ")
}

// Log logs a leveled entry.
func (l *DefaultLogger) Log(reqContext interface{}, loggingCtx string, level int, msg string, args []interface{}) {
	// log if the log level (warn=4) >= level of message (err=3)
	if l.level < level || silent {
		return
	}
	l.formatter.Format(reqContext, loggingCtx, l.writeHandler, level, msg, args)
}

// IsTrace determines if this logger logs a debug statement.
func (l *DefaultLogger) IsTrace() bool {
	// DEBUG(7) >= TRACE(10)
	return l.level >= LevelTrace
}

// IsDebug determines if this logger logs a debug statement.
func (l *DefaultLogger) IsDebug() bool {
	return l.level >= LevelDebug
}

// IsInfo determines if this logger logs an info statement.
func (l *DefaultLogger) IsInfo() bool {
	return l.level >= LevelInfo
}

// IsWarn determines if this logger logs a warning statement.
func (l *DefaultLogger) IsWarn() bool {
	return l.level >= LevelWarn
}

// SetLevel sets the level of this logger.
func (l *DefaultLogger) SetLevel(level int) {
	l.level = level
}

// SetFormatter set the formatter for this logger.
func (l *DefaultLogger) SetFormatter(formatter Formatter) {
	l.formatter = formatter
}

type DefaultWriteHandler struct {
	writer io.Writer
}

func (dwh *DefaultWriteHandler) WriteLog(reqContext interface{}, loggingCtx string, buf *bytes.Buffer, level int, msg string, args []interface{}) {
	buf.WriteTo(dwh.writer)
}
