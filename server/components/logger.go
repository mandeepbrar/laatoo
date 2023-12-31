package components

import "laatoo.io/sdk/ctx"

type Logger interface {
	Trace(reqContext ctx.Context, msg string, args ...interface{})
	Debug(reqContext ctx.Context, msg string, args ...interface{})
	Info(reqContext ctx.Context, msg string, args ...interface{})
	Warn(reqContext ctx.Context, msg string, args ...interface{})
	Error(reqContext ctx.Context, msg string, args ...interface{})
	Fatal(reqContext ctx.Context, msg string, args ...interface{})

	SetLevel(int)
	SetFormat(string)
	GetLevel() int
	GetFormat() string
	IsTrace() bool
	IsDebug() bool
	IsInfo() bool
	IsWarn() bool
}
