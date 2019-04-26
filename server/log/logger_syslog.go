// +build !windows

package log

import (
	"fmt"
	"io"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"log/syslog"
)

var syslogWriteHandler WriteHandler

func NewSysLogger(ctx core.ServerContext, appname string, settings config.Config) components.Logger {
	if syslogWriteHandler == nil {
		protocol := "tcp"
		addr := "localhost:601"
		var ok bool
		if settings != nil {
			protocol, ok = settings.GetString(ctx, "protocol")
			if !ok {
				protocol = "tcp"
			}
			addr, ok = settings.GetString(ctx, "address")
			if !ok {
				addr = "localhost:601"
			}
		}
		logWriter, err := syslog.Dial(protocol, addr, syslog.LOG_ERR, appname)
		if err != nil {
			fmt.Println("********* Error in creating sys logger ************ ", err, protocol, addr)
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

func (jh *SyslogWriteHandler) Print(ctx ctx.Context, appname string, msg string, level int, strlevel string) {
	jh.writer.Write([]byte(msg))
}
func (jh *SyslogWriteHandler) PrintBytes(ctx ctx.Context, appname string, msg []byte, level int, strlevel string) (int, error) {
	return jh.writer.Write(msg)
}
