package log

// Trace logs a trace statement. On terminals file and line number are logged.
func Trace(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	DefaultLog.Trace(reqContext, loggingCtx, msg, args...)
}

// Debug logs a debug statement.
func Debug(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	DefaultLog.Debug(reqContext, loggingCtx, msg, args...)
}

// Info logs an info statement.
func Info(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	DefaultLog.Info(reqContext, loggingCtx, msg, args...)
}

// Warn logs a warning statement. On terminals it logs file and line number.
func Warn(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	DefaultLog.Warn(reqContext, loggingCtx, msg, args...)
}

// Error logs an error statement with callstack.
func Error(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	DefaultLog.Error(reqContext, loggingCtx, msg, args...)
}

// Fatal logs a fatal statement.
func Fatal(reqContext interface{}, loggingCtx string, msg string, args ...interface{}) {
	DefaultLog.Fatal(reqContext, loggingCtx, msg, args...)
}

// IsTrace determines if this logger logs a trace statement.
func IsTrace() bool {
	return DefaultLog.IsTrace()
}

// IsDebug determines if this logger logs a debug statement.
func IsDebug() bool {
	return DefaultLog.IsDebug()
}

// IsInfo determines if this logger logs an info statement.
func IsInfo() bool {
	return DefaultLog.IsInfo()
}

// IsWarn determines if this logger logs a warning statement.
func IsWarn() bool {
	return DefaultLog.IsWarn()
}
