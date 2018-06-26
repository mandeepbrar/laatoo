package log

import slog "laatoo/sdk/server/log"

func GetLevel(level string) int {
	switch level {
	case "all":
		return slog.TRACE
	case "trace":
		return slog.TRACE
	case "debug":
		return slog.DEBUG
	case "info":
		return slog.INFO
	case "warn":
		return slog.WARN
	default:
		return slog.ERROR
	}
}
