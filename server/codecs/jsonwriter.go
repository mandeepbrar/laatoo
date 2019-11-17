package codecs

import (
	"bytes"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"time"
)

func NewJsonWriter(c ctx.Context) (core.SerializableWriter, error) {
	var buf bytes.Buffer
	return &JsonWriter{&buf}, nil
}

type JsonWriter struct {
	buffer *bytes.Buffer
}

func (wtr *JsonWriter) Bytes() []byte {
	return wtr.buffer.Bytes()
}

func (wtr *JsonWriter) Write(ctx ctx.Context, cdc core.Codec, prop string) (core.SerializableWriter, error) {

}

func (wtr *JsonWriter) WriteByte(ctx ctx.Context, cdc core.Codec, prop string, val *byte) error {

}

func (wtr *JsonWriter) WriteBytes(ctx ctx.Context, cdc core.Codec, val *[]byte) error {

}

func (wtr *JsonWriter) WriteInt(ctx ctx.Context, cdc core.Codec, prop string, val *int) error {

}
func (wtr *JsonWriter) WriteInt64(ctx ctx.Context, cdc core.Codec, prop string, val *int64) error {

}
func (wtr *JsonWriter) WriteString(ctx ctx.Context, cdc core.Codec, prop string, val *string) error {

}
func (wtr *JsonWriter) WriteFloat32(ctx ctx.Context, cdc core.Codec, prop string, val *float32) error {

}
func (wtr *JsonWriter) WriteFloat64(ctx ctx.Context, cdc core.Codec, prop string, val *float64) error {

}
func (wtr *JsonWriter) WriteBool(ctx ctx.Context, cdc core.Codec, prop string, val *bool) error {

}
func (wtr *JsonWriter) WriteObject(ctx ctx.Context, cdc core.Codec, prop string) (core.SerializableWriter, error) {

}
func (wtr *JsonWriter) WriteMap(ctx ctx.Context, cdc core.Codec, prop string, val *map[string]interface{}) error {

}
func (wtr *JsonWriter) WriteArray(ctx ctx.Context, cdc core.Codec, prop string, val interface{}) error {

}
func (wtr *JsonWriter) WriteTime(ctx ctx.Context, cdc core.Codec, prop string, val *time.Time) error {

}
