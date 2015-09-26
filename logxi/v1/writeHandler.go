package log

import (
	"bytes"
)

type WriteHandler interface {
	WriteLog(reqContext interface{}, loggingCtx string, buf *bytes.Buffer, level int, msg string, args []interface{})
}
