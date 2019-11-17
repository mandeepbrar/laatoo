package codecs

import (
	"bytes"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"reflect"
	"strconv"
	"time"
	"unicode/utf8"
)

func NewJsonWriter(c ctx.Context) (core.SerializableWriter, error) {
	var buf bytes.Buffer
	return &JsonWriter{&buf}, nil
}

type JsonWriter struct {
	buffer *bytes.Buffer
}

func (wtr *JsonWriter) Bytes() []byte {
	wtr.buffer.Write(endObject)
	bys := wtr.buffer.Bytes()
	if len(bys) > 0 && bys[0] == byte(',') {
		bys[0] = byte('{')
	}
	return bys
}

func (wtr *JsonWriter) WriteByte(ctx ctx.Context, cdc core.Codec, prop string, val *byte) error {
	wtr.key(ctx, prop)
	_, err := wtr.buffer.Write([]byte{*val})
	return err
}

func (wtr *JsonWriter) WriteBytes(ctx ctx.Context, cdc core.Codec, prop string, val *[]byte) error {
	wtr.key(ctx, prop)
	_, err := wtr.buffer.Write(*val)
	return err
}

func (wtr *JsonWriter) WriteInt(ctx ctx.Context, cdc core.Codec, prop string, val *int) error {
	wtr.key(ctx, prop)
	_, err := wtr.buffer.Write([]byte(strconv.Itoa(*val)))
	return err
}

func (wtr *JsonWriter) WriteInt64(ctx ctx.Context, cdc core.Codec, prop string, val *int64) error {
	wtr.key(ctx, prop)
	_, err := wtr.buffer.Write([]byte(strconv.FormatInt(*val, 10)))
	return err
}

func (wtr *JsonWriter) WriteString(ctx ctx.Context, cdc core.Codec, prop string, val *string) error {
	wtr.key(ctx, prop)
	return wtr.escapedString(ctx, *val)
}

func (wtr *JsonWriter) WriteFloat32(ctx ctx.Context, cdc core.Codec, prop string, val *float32) error {
	wtr.key(ctx, prop)
	_, err := wtr.buffer.Write([]byte(strconv.FormatFloat(float64(*val), 'g', -1, 32)))
	return err
}

func (wtr *JsonWriter) WriteFloat64(ctx ctx.Context, cdc core.Codec, prop string, val *float64) error {
	wtr.key(ctx, prop)
	_, err := wtr.buffer.Write([]byte(strconv.FormatFloat(*val, 'g', -1, 64)))
	return err
}

func (wtr *JsonWriter) WriteBool(ctx ctx.Context, cdc core.Codec, prop string, val *bool) error {
	wtr.key(ctx, prop)
	var err error
	if *val {
		_, err = wtr.buffer.Write(_true)
	} else {
		_, err = wtr.buffer.Write(_false)
	}
	return err
}

func (wtr *JsonWriter) WriteObject(ctx ctx.Context, cdc core.Codec, prop string, val interface{}) error {
	wtr.key(ctx, prop)
	var bys []byte
	var err error
	srl, ok := val.(core.Serializable)
	if ok {
		bys, err = cdc.MarshalSerializable(ctx, srl)
	} else {
		bys, err = cdc.Marshal(ctx, val)
	}
	if err != nil {
		return err
	}
	_, err = wtr.buffer.Write(bys)
	return err
}

func (wtr *JsonWriter) WriteMap(ctx ctx.Context, cdc core.Codec, prop string, val *map[string]interface{}) error {
	wtr.key(ctx, prop)
	bys, err := cdc.Marshal(ctx, *val)
	_, err = wtr.buffer.Write(bys)
	return err
}

func (wtr *JsonWriter) WriteArray(ctx ctx.Context, cdc core.Codec, prop string, val interface{}) error {
	var bys []byte
	var err error
	var obj interface{}
	wtr.key(ctx, prop)
	wtr.buffer.Write(startArray)
	collVal := reflect.ValueOf(val).Elem()
	for i := 0; i < collVal.Len(); i++ {
		if i != 0 {
			wtr.buffer.Write(comma)
		}
		obj = collVal.Index(i).Interface()
		srl, ok := obj.(core.Serializable)
		if ok {
			bys, err = cdc.MarshalSerializable(ctx, srl)
		} else {
			bys, err = cdc.Marshal(ctx, obj)
		}
		if err != nil {
			return err
		}
		_, err = wtr.buffer.Write(bys)
	}
	_, err = wtr.buffer.Write(endArray)
	return err
}

func (wtr *JsonWriter) WriteTime(ctx ctx.Context, cdc core.Codec, prop string, val *time.Time) error {
	wtr.key(ctx, prop)
	wtr.buffer.Write(quote)
	wtr.buffer.Write([]byte(val.Format(time.RFC3339Nano)))
	_, err := wtr.buffer.Write(quote)
	return err
}

func (wtr *JsonWriter) key(ctx ctx.Context, key string) error {
	//wtr.buffer.Separator()
	wtr.buffer.Write(comma)
	wtr.buffer.Write(keyStart)
	wtr.buffer.Write([]byte(key))
	_, err := wtr.buffer.Write(keyEnd)
	return err
}

func (wtr *JsonWriter) escapedString(ctx ctx.Context, val string) error {
	//wtr.buffer.Separator()
	wtr.buffer.Write(quote)
	start, end := 0, 0
	var special []byte
L:
	for i, r := range val {
		switch r {
		case '"':
			special = escapedQuote
		case '\\':
			special = escapedSlash
		case '\b':
			special = escapedBS
		case '\f':
			special = escapedFF
		case '\n':
			special = escapedNL
		case '\r':
			special = escapedLF
		case '\t':
			special = escapedTab
		case 0x00:
			special = escapedNull
		case 0x01:
			special = escapedSOH
		case 0x02:
			special = escapedSTX
		case 0x03:
			special = escapedETX
		case 0x04:
			special = escapedEOT
		case 0x05:
			special = escapedENQ
		case 0x06:
			special = escapedACK
		case 0x07:
			special = escapedBEL
		case 0x0b:
			special = escapedVT
		case 0x0e:
			special = escapedSO
		case 0x0f:
			special = escapedSI
		case 0x10:
			special = escapedDLE
		case 0x11:
			special = escapedDC1
		case 0x12:
			special = escapedDC2
		case 0x13:
			special = escapedDC3
		case 0x14:
			special = escapedDC4
		case 0x15:
			special = escapedNAK
		case 0x16:
			special = escapedSYN
		case 0x17:
			special = escapedETB
		case 0x18:
			special = escapedCAN
		case 0x19:
			special = escapedEM
		case 0x1a:
			special = escapedSUB
		case 0x1b:
			special = escapedESC
		case 0x1c:
			special = escapedFS
		case 0x1d:
			special = escapedGS
		case 0x1e:
			special = escapedRS
		case 0x1f:
			special = escapedUS
		default:
			end += utf8.RuneLen(r)
			continue L
		}

		if end > start {
			wtr.buffer.Write([]byte(val[start:end]))
		}
		wtr.buffer.Write(special)
		start = i + 1
		end = start
	}
	if end > start {
		wtr.buffer.Write([]byte(val[start:end]))
	}

	_, err := wtr.buffer.Write(quote)
	return err
}

var (
	quote        = []byte(`"`)
	keyStart     = quote
	null         = []byte("null")
	_true        = []byte("true")
	_false       = []byte("false")
	comma        = []byte(",")
	keyEnd       = []byte(`":`)
	startObject  = []byte("{")
	endObject    = []byte("}")
	startArray   = []byte("[")
	endArray     = []byte("]")
	escapedQuote = []byte(`\"`)
	escapedSlash = []byte(`\\`)
	escapedBS    = []byte(`\b`)
	escapedFF    = []byte(`\f`)
	escapedNL    = []byte(`\n`)
	escapedLF    = []byte(`\r`)
	escapedTab   = []byte(`\t`)
	escapedNull  = []byte(`\u0000`)
	escapedSOH   = []byte(`\u0001`)
	escapedSTX   = []byte(`\u0002`)
	escapedETX   = []byte(`\u0003`)
	escapedEOT   = []byte(`\u0004`)
	escapedENQ   = []byte(`\u0005`)
	escapedACK   = []byte(`\u0006`)
	escapedBEL   = []byte(`\u0007`)
	escapedVT    = []byte(`\u000b`)
	escapedSO    = []byte(`\u000e`)
	escapedSI    = []byte(`\u000f`)
	escapedDLE   = []byte(`\u0010`)
	escapedDC1   = []byte(`\u0011`)
	escapedDC2   = []byte(`\u0012`)
	escapedDC3   = []byte(`\u0013`)
	escapedDC4   = []byte(`\u0014`)
	escapedNAK   = []byte(`\u0015`)
	escapedSYN   = []byte(`\u0016`)
	escapedETB   = []byte(`\u0017`)
	escapedCAN   = []byte(`\u0018`)
	escapedEM    = []byte(`\u0019`)
	escapedSUB   = []byte(`\u001a`)
	escapedESC   = []byte(`\u001b`)
	escapedFS    = []byte(`\u001c`)
	escapedGS    = []byte(`\u001d`)
	escapedRS    = []byte(`\u001e`)
	escapedUS    = []byte(`\u001f`)
)
