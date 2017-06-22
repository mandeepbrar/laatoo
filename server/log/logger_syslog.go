// +build !windows

package log

import (
	"io"
	"laatoo/sdk/components"
	"laatoo/sdk/core"
	"log/syslog"
)

var syslogWriteHandler WriteHandler

func NewSysLogger(appname string) components.Logger {
	if syslogWriteHandler == nil {
		logWriter, err := syslog.Dial("", "", syslog.LOG_ERR, appname)
		if err != nil {
			syslogWriteHandler = stdSimpleLogsHandler()
		} else {
			syslogWriteHandler = sysLogsHandler(logWriter)
		}
	}
	return NewSimpleLogger(appname, syslogWriteHandler)
}

func sysLogsHandler(writer io.Writer) WriteHandler {
	wh := &SyslogWriteHandler{writer}
	return wh
}

type SyslogWriteHandler struct {
	writer io.Writer
}

func (jh *SyslogWriteHandler) Print(ctx core.Context, appname string, msg string, level int, strlevel string) {
	jh.writer.Write([]byte(msg))
}
func (jh *SyslogWriteHandler) PrintBytes(ctx core.Context, appname string, msg []byte, level int, strlevel string) (int, error) {
	return jh.writer.Write(msg)
}
