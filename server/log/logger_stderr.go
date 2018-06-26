// +build !appengine

package log

import (
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/ctx"
	"os"
)

var stderrWriteHandler WriteHandler

func NewStdLogger(appname string) components.Logger {
	if stderrWriteHandler == nil {
		stderrWriteHandler = stdSimpleLogsHandler()
	}
	return NewSimpleLogger(appname, stderrWriteHandler)
}

func stdSimpleLogsHandler() WriteHandler {
	wh := &stdSimpleWriteHandler{}
	return wh
}

type stdSimpleWriteHandler struct {
}

func (jh *stdSimpleWriteHandler) Print(ctx ctx.Context, appname string, msg string, level int, strlevel string) {
	os.Stderr.WriteString(msg)
}
func (jh *stdSimpleWriteHandler) PrintBytes(ctx ctx.Context, appname string, msg []byte, level int, strlevel string) (int, error) {
	return os.Stderr.Write(msg)
}
