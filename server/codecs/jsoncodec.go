package codecs

import (
	"fmt"
	"io"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/log"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type JsonCodec struct {
}

func NewJsonCodec() *JsonCodec {
	return &JsonCodec{}
}

func (cdc *JsonCodec) Marshal(c ctx.Context, val interface{}) ([]byte, error) {
	if val == nil {
		return nil, nil
	}
	srl, ok := val.(core.Serializable)
	if ok {
		return cdc.MarshalSerializable(c, srl)
	}
	return json.Marshal(val)
}

func (cdc *JsonCodec) Unmarshal(c ctx.Context, arr []byte, val interface{}) (err error) {
	c = c.SubCtx("unmarshalling")
	defer c.CompleteContext()
	if arr == nil {
		return nil
	}

	srl, ok := val.(core.Serializable)
	if ok {
		err = cdc.UnmarshalSerializable(c, arr, srl)
	}
	err = json.Unmarshal(arr, val)
	log.Error(c, "Unmarshalling error values", "input", string(arr), "output", fmt.Sprint("%#v", val), "time taken", c.GetElapsedTime())
	return
}

func (cdc *JsonCodec) Encode(c ctx.Context, outStream io.Writer, val interface{}) error {
	if val == nil {
		return nil
	}
	srl, ok := val.(core.Serializable)
	if ok {
		err := cdc.MarshalSerializableToStream(c, outStream, srl)
		log.Error(c, "Encoding ends. Sending error", "err", err)
		if err != nil {
			return err
		}
		return nil
	}
	enc := json.NewEncoder(outStream)
	return enc.Encode(val)
}

func (cdc *JsonCodec) Decode(c ctx.Context, inpStream io.Reader, val interface{}) error {
	if val == nil {
		return nil
	}
	srl, ok := val.(core.Serializable)
	if ok {
		return cdc.UnmarshalSerializableFromStream(c, inpStream, srl)
	}
	dec := json.NewDecoder(inpStream)
	return dec.Decode(val)
}

func (codec *JsonCodec) UnmarshalSerializable(c ctx.Context, arr []byte, obj core.Serializable) error {
	return UnmarshalSerializable(c, codec, NewJsonReader, arr, obj)
}

func (codec *JsonCodec) UnmarshalSerializableProps(c ctx.Context, arr []byte, obj core.Serializable, props map[string]interface{}) error {
	return UnmarshalSerializableProps(c, codec, NewJsonReader, arr, obj, props)
}

func (codec *JsonCodec) UnmarshalSerializableFromStream(c ctx.Context, rdr io.Reader, obj core.Serializable) error {
	return UnmarshalSerializableFromStream(c, codec, NewJsonStreamReader, rdr, obj)
}

func (codec *JsonCodec) UnmarshalSerializableFromStreamProps(c ctx.Context, rdr io.Reader, obj core.Serializable, props map[string]interface{}) error {
	return UnmarshalStreamProps(c, codec, NewJsonStreamReader, rdr, obj, props)
}

func (codec *JsonCodec) MarshalSerializable(c ctx.Context, obj core.Serializable) ([]byte, error) {
	return MarshalSerializable(c, codec, NewJsonWriter, obj)
}

func (codec *JsonCodec) MarshalSerializableProps(c ctx.Context, obj core.Serializable, props map[string]interface{}) ([]byte, error) {
	return MarshalSerializableProps(c, codec, NewJsonWriter, obj, props)
}

func (codec *JsonCodec) MarshalSerializableToStream(c ctx.Context, wtr io.Writer, obj core.Serializable) error {
	return MarshalSerializableToStream(c, codec, NewJsonStreamWriter, obj, wtr)
}

func (codec *JsonCodec) MarshalSerializableToStreamProps(c ctx.Context, wtr io.Writer, obj core.Serializable, props map[string]interface{}) error {
	return MarshalStreamProps(c, codec, NewJsonStreamWriter, obj, wtr, props)
}

func (codec *JsonCodec) UnmarshalReader(c ctx.Context, rdr core.SerializableReader, obj core.Serializable) error {
	return UnmarshalReader(c, codec, internalJsonReader, rdr, obj)
}

func (codec *JsonCodec) MarshalWriter(c ctx.Context, wtr core.SerializableWriter, obj core.Serializable) ([]byte, error) {
	obj.WriteAll(c, codec, wtr)
	byts := wtr.Bytes()
	return byts, nil
}
