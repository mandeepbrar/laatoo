package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	//"io"
	"laatoo/sdk/core"
	"strings"
	"time"
)

const (
	STR_TRACE = "Trace"
	STR_DEBUG = "Debug"
	STR_INFO  = "Info"
	STR_WARN  = "Warn"
	STR_ERROR = "Error"
	STR_FATAL = "Fatal"
)

type SimpleWriteHandler interface {
	Print(msg string)
	PrintBytes(msg []byte) (int, error)
}

func NewSimpleLogger(wh SimpleWriteHandler) LoggerInterface {
	return &SimpleLogger{format: "json", level: INFO, wh: wh}
}

type SimpleLogger struct {
	wh SimpleWriteHandler
	//buffer bytes.Buffer
	format string
	level  int
}

func (log *SimpleLogger) Trace(reqContext core.Context, msg string, args ...interface{}) {
	if log.level > DEBUG {
		log.wh.Print(log.buildMessage(STR_TRACE, reqContext, msg, args...))
	}
}
func (log *SimpleLogger) Debug(reqContext core.Context, msg string, args ...interface{}) {
	if log.level > INFO {
		log.wh.Print(log.buildMessage(STR_DEBUG, reqContext, msg, args...))
	}
}
func (log *SimpleLogger) Info(reqContext core.Context, msg string, args ...interface{}) {
	if log.level > WARN {
		log.wh.Print(log.buildMessage(STR_INFO, reqContext, msg, args...))
	}
}
func (log *SimpleLogger) Warn(reqContext core.Context, msg string, args ...interface{}) {
	if log.level > ERROR {
		log.wh.Print(log.buildMessage(STR_WARN, reqContext, msg, args...))
	}
}
func (log *SimpleLogger) Error(reqContext core.Context, msg string, args ...interface{}) {
	if log.level > FATAL {
		log.wh.Print(log.buildMessage(STR_ERROR, reqContext, msg, args...))
	}
}
func (log *SimpleLogger) Fatal(reqContext core.Context, msg string, args ...interface{}) {
	log.wh.Print(log.buildMessage(STR_FATAL, reqContext, msg, args...))
}

func (log *SimpleLogger) SetFormat(format string) {
	log.format = format
}

func (log *SimpleLogger) SetType(loggertype string) {
}

func (log *SimpleLogger) SetLevel(level int) {
	log.level = level
}
func (log *SimpleLogger) IsTrace() bool {
	return log.level == TRACE
}
func (log *SimpleLogger) IsDebug() bool {
	return log.level == DEBUG
}
func (log *SimpleLogger) IsInfo() bool {
	return log.level == INFO
}
func (log *SimpleLogger) IsWarn() bool {
	return log.level == WARN
}

func (log *SimpleLogger) Write(p []byte) (int, error) {
	return log.wh.PrintBytes(p)
}

func (log *SimpleLogger) buildMessage(level string, reqContext core.Context, msg string, args ...interface{}) string {
	if len(args)%2 > 0 {
		panic("wrong logging")
	}
	switch log.format {
	case "json":
		{
			var buffer bytes.Buffer
			enc := json.NewEncoder(&buffer)
			mapToPrint := map[string]string{"MESSAGE": msg}
			argslen := len(args)
			for i := 0; (i + 1) < argslen; i = i + 2 {
				mapToPrint[args[i].(string)] = fmt.Sprint(args[i+1])
			}
			err := enc.Encode(mapToPrint)
			if err != nil {
				fmt.Println(err)
			}
			return buffer.String()
		}
	case "jsonmax":
		{
			var buffer bytes.Buffer
			enc := json.NewEncoder(&buffer)
			mapToPrint := map[string]string{"TIME": time.Now().String(), "LEVEL": level, "CONTEXT": reqContext.GetName(), "ID": reqContext.GetId(), "MESSAGE": msg}
			argslen := len(args)
			for i := 0; (i + 1) < argslen; i = i + 2 {
				mapToPrint[args[i].(string)] = fmt.Sprint(args[i+1])
			}
			err := enc.Encode(mapToPrint)
			if err != nil {
				fmt.Println(err)
			}
			return buffer.String()
		}
	case "happymaxcolor":
		{
			var buffer bytes.Buffer
			firstline := ""
			if level == STR_ERROR {
				firstline = color.RedString("%s: %s", strings.ToUpper(level), msg)
			} else {
				firstline = color.BlueString("%s: %s", strings.ToUpper(level), msg)
			}
			argslen := len(args)
			if argslen > 0 {
				firstline = fmt.Sprintf("%s    %s", firstline, color.MagentaString("%s:%s", strings.ToUpper(args[0].(string)), fmt.Sprint(args[1])))
			}
			if argslen > 2 {
				firstline = fmt.Sprintf("%s    %s", firstline, color.CyanString("%s:%s", strings.ToUpper(args[2].(string)), fmt.Sprint(args[3])))
			}
			buffer.WriteString(fmt.Sprintln(firstline))
			for i := 4; (i + 1) < argslen; i = i + 2 {
				buffer.WriteString(fmt.Sprintln("		", args[i], ":", args[i+1]))
			}
			buffer.WriteString(fmt.Sprintln("		", reqContext.GetName()))
			buffer.WriteString(fmt.Sprintln("		TIME: ", time.Now().String()))
			buffer.WriteString(fmt.Sprintln("		ID: ", reqContext.GetId()))
			return buffer.String()
		}
	case "happy":
		{
			var buffer bytes.Buffer
			firstline := fmt.Sprintf("%s: %s", strings.ToUpper(level), msg)
			argslen := len(args)
			if argslen > 0 {
				firstline = fmt.Sprintf("%s    %s:%s", firstline, strings.ToUpper(args[0].(string)), fmt.Sprint(args[1]))
			}
			if argslen > 2 {
				firstline = fmt.Sprintf("%s    %s:%s", firstline, strings.ToUpper(args[2].(string)), fmt.Sprint(args[3]))
			}
			buffer.WriteString(fmt.Sprintln(firstline))
			for i := 4; (i + 1) < argslen; i = i + 2 {
				buffer.WriteString(fmt.Sprintln("		", args[i], ":", args[i+1]))
			}
			buffer.WriteString(fmt.Sprintln("		", reqContext.GetName()))
			return buffer.String()
		}
	case "happycolor":
		{
			var buffer bytes.Buffer
			firstline := ""
			if level == STR_ERROR {
				firstline = color.RedString("%s: %s", strings.ToUpper(level), msg)
			} else {
				firstline = color.BlueString("%s: %s", strings.ToUpper(level), msg)
			}
			argslen := len(args)
			if argslen > 0 {
				firstline = fmt.Sprintf("%s    %s", firstline, color.MagentaString("%s:%s", strings.ToUpper(args[0].(string)), fmt.Sprint(args[1])))
			}
			if argslen > 2 {
				firstline = fmt.Sprintf("%s    %s", firstline, color.CyanString("%s:%s", strings.ToUpper(args[2].(string)), fmt.Sprint(args[3])))
			}
			buffer.WriteString(fmt.Sprintln(firstline))
			for i := 4; (i + 1) < argslen; i = i + 2 {
				buffer.WriteString(fmt.Sprintln("		", args[i], ":", args[i+1]))
			}
			buffer.WriteString(fmt.Sprintln("		", reqContext.GetName()))
			return buffer.String()
		}
	case "happymax":
		{
			var buffer bytes.Buffer
			buffer.WriteString(fmt.Sprintln("MESSAGE ", msg))
			buffer.WriteString(fmt.Sprintln("		TIME ", time.Now().String()))
			buffer.WriteString(fmt.Sprintln("		LEVEL ", level))
			buffer.WriteString(fmt.Sprintln("		CONTEXT ", reqContext.GetName()))
			buffer.WriteString(fmt.Sprintln("		ID ", reqContext.GetId()))
			argslen := len(args)
			for i := 0; (i + 1) < argslen; i = i + 2 {
				buffer.WriteString(fmt.Sprintln("		", args[i], " ", args[i+1]))
			}
			return buffer.String()
		}
	default:
		{
			var buffer bytes.Buffer
			buffer.WriteString(fmt.Sprintln("MESSAGE ", msg))
			argslen := len(args)
			for i := 0; (i + 1) < argslen; i = i + 2 {
				buffer.WriteString(fmt.Sprintln("		", args[i], " ", args[i+1]))
			}
			return buffer.String()
		}
	}
}
