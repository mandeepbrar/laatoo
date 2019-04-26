package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"laatoo/sdk/server/ctx"
	"strings"
	"time"

	rfc5424 "github.com/influxdata/go-syslog/rfc5424"
	//"github.com/crewjam/rfc5424"
)

func init() {
	logFormats[CONF_FMT_JSON] = printJSON
	logFormats[CONF_FMT_JSONMAX] = printJSONMax
	logFormats[CONF_FMT_HAPPY] = printHappy
	logFormats[CONF_FMT_HAPPYMAX] = printHappyMax
	logFormats[CONF_FMT_RFC5424] = printRFC5424
}

func printJSON(ctx ctx.Context, app string, strlevel string, wh WriteHandler, level int, msg string, args ...interface{}) {
	if len(args)%2 > 0 {
		panic("wrong logging")
	}
	var buffer bytes.Buffer
	enc := json.NewEncoder(&buffer)
	mapToPrint := map[string]string{"MESSAGE": msg, "LEVEL": strlevel}
	if ctx != nil {
		mapToPrint["CONTEXT"] = ctx.GetName()
		mapToPrint["LEVEL"] = strlevel
	}
	argslen := len(args)
	for i := 0; (i + 1) < argslen; i = i + 2 {
		mapToPrint[args[i].(string)] = fmt.Sprint(args[i+1])
	}
	err := enc.Encode(mapToPrint)
	if err != nil {
		fmt.Println(err)
	}
	wh.Print(ctx, app, buffer.String(), level, strlevel)
}
func printJSONMax(ctx ctx.Context, app string, strlevel string, wh WriteHandler, level int, msg string, args ...interface{}) {
	if len(args)%2 > 0 {
		panic("wrong logging")
	}
	var buffer bytes.Buffer
	enc := json.NewEncoder(&buffer)
	mapToPrint := map[string]string{"TIME": time.Now().String(), "MESSAGE": msg, "LEVEL": strlevel}
	if ctx != nil {
		mapToPrint["CONTEXT"] = ctx.GetName()
		mapToPrint["PATH"] = ctx.GetPath()
		mapToPrint["ID"] = ctx.GetId()
	}
	argslen := len(args)
	for i := 0; (i + 1) < argslen; i = i + 2 {
		mapToPrint[args[i].(string)] = fmt.Sprint(args[i+1])
	}
	err := enc.Encode(mapToPrint)
	if err != nil {
		fmt.Println(err)
	}
	wh.Print(ctx, app, buffer.String(), level, strlevel)
}
func printHappy(ctx ctx.Context, app string, strlevel string, wh WriteHandler, level int, msg string, args ...interface{}) {
	if len(args)%2 > 0 {
		panic("wrong logging")
	}
	var buffer bytes.Buffer
	firstline := msg
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
	if ctx != nil {
		buffer.WriteString(fmt.Sprintln("		", ctx.GetName()))
	}
	wh.Print(ctx, app, buffer.String(), level, strlevel)
}
func printHappyMax(ctx ctx.Context, app string, strlevel string, wh WriteHandler, level int, msg string, args ...interface{}) {
	if len(args)%2 > 0 {
		panic("wrong logging")
	}
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintln("MESSAGE ", msg))
	buffer.WriteString(fmt.Sprintln("		TIME ", time.Now().String()))
	if ctx != nil {
		buffer.WriteString(fmt.Sprintln("		CONTEXT ", ctx.GetName()))
		buffer.WriteString(fmt.Sprintln("		PATH ", ctx.GetPath()))
		buffer.WriteString(fmt.Sprintln("		ID ", ctx.GetId()))
	}
	argslen := len(args)
	for i := 0; (i + 1) < argslen; i = i + 2 {
		buffer.WriteString(fmt.Sprintln("		", args[i], " ", args[i+1]))
	}
	wh.Print(ctx, app, buffer.String(), level, strlevel)
}

/*
//"github.com/crewjam/rfc5424"
func printRFC5424(ctx ctx.Context, app string, strlevel string, wh WriteHandler, level int, msg string, args ...interface{}) {
	if len(args)%2 > 0 {
		panic("wrong logging")
	}

	m := rfc5424.Message{
		Priority:  rfc5424.Daemon | rfc5424.Info,
		Timestamp: time.Now(),
		AppName:   app,
		Message:   []byte(msg),
	}

	if ctx != nil {
		m.AddDatum("Params", "CONTEXT", ctx.GetName())
		m.AddDatum("Params", "PATH", ctx.GetPath())
		m.AddDatum("Params", "ID", ctx.GetId())
	}
	argslen := len(args)
	for i := 0; (i + 1) < argslen; i = i + 2 {
		if args[i] != nil {
			pname := args[i].(string)
			m.AddDatum("Params", pname, fmt.Sprint(args[i+1]))
		}
	}

	rfcmsgtxt, err := m.MarshalBinary()
	if err != nil {
		fmt.Println(err)
	}

	wh.PrintBytes(ctx, app, rfcmsgtxt, level, strlevel)
}*/

//	rfc5424 "github.com/influxdata/go-syslog/rfc5424"

func printRFC5424(ctx ctx.Context, app string, strlevel string, wh WriteHandler, level int, msg string, args ...interface{}) {
	if len(args)%2 > 0 {
		panic("wrong logging")
	}

	rfcmsg := &rfc5424.SyslogMessage{}
	rfcmsg.SetTimestamp(time.Now().Format(time.RFC3339))
	rfcmsg.SetPriority(191)
	rfcmsg.SetVersion(1)
	rfcmsg.SetMessage(msg)
	rfcmsg.SetAppname(app)
	rfcmsg.SetElementID("PARAMS")
	if ctx != nil {
		rfcmsg.SetParameter("PARAMS", "CONTEXT", ctx.GetName())
		rfcmsg.SetParameter("PARAMS", "PATH", ctx.GetPath())
		rfcmsg.SetParameter("PARAMS", "ID", ctx.GetId())
	}
	argslen := len(args)
	for i := 0; (i + 1) < argslen; i = i + 2 {
		if args[i] != nil {
			pname := args[i].(string)
			rfcmsg.SetParameter("PARAMS", pname, fmt.Sprint(args[i+1]))
		}
	}

	rfcmsgtxt, err := rfcmsg.String()
	if err != nil {
		fmt.Println(err)
	}
	wh.Print(ctx, app, rfcmsgtxt+"\n", level, strlevel)
}
