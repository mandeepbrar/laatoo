// +build !appengine

package log

import (
	"laatoo/sdk/core"
	slog "laatoo/sdk/log"
	"os"
)

var stderrWriteHandler WriteHandler

func NewStdLogger(appname string) slog.Logger {
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

func (jh *stdSimpleWriteHandler) Print(ctx core.Context, appname string, msg string, level int, strlevel string) {
	os.Stderr.WriteString(msg)
}
func (jh *stdSimpleWriteHandler) PrintBytes(ctx core.Context, appname string, msg []byte, level int, strlevel string) (int, error) {
	return os.Stderr.Write(msg)
}
