package codecs

import (
	"encoding/json"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
)

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

func (cdc *JsonCodec) Unmarshal(c ctx.Context, arr []byte, val interface{}) error {
	if arr == nil {
		return nil
	}
	srl, ok := val.(core.Serializable)
	if ok {
		return cdc.UnmarshalSerializable(c, arr, srl)
	}
	return json.Unmarshal(arr, val)
}

func (codec *JsonCodec) UnmarshalSerializable(c ctx.Context, arr []byte, obj core.Serializable) error {
	return UnmarshalSerializable(c, codec, NewJsonReader, arr, obj)
}

func (codec *JsonCodec) MarshalSerializable(c ctx.Context, obj core.Serializable) ([]byte, error) {
	return MarshalSerializable(c, codec, NewJsonWriter, obj)
}

func (codec *JsonCodec) UnmarshalSerializableProps(c ctx.Context, arr []byte, obj core.Serializable, props map[string]interface{}) error {
	return UnmarshalSerializableProps(c, codec, NewJsonReader, arr, obj, props)
}

func (codec *JsonCodec) MarshalSerializableProps(c ctx.Context, obj core.Serializable, props map[string]interface{}) ([]byte, error) {
	return MarshalSerializableProps(c, codec, NewJsonWriter, obj, props)
}

func (codec *JsonCodec) UnmarshalReader(c ctx.Context, rdr core.SerializableReader, obj core.Serializable) error {
	return UnmarshalReader(c, codec, internalJsonReader, rdr, obj)
}

func (codec *JsonCodec) MarshalWriter(c ctx.Context, wtr core.SerializableWriter) ([]byte, error) {
	return nil, nil
}
