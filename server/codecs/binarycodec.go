package codecs

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"

	"github.com/ugorji/go/codec"
)

type BinaryCodec struct {
	enc *codec.Encoder
	dec *codec.Decoder
	h   codec.Handle
}

func NewBinaryCodec() *BinaryCodec {
	return &BinaryCodec{h: new(codec.BincHandle)}
}

func (cdc *BinaryCodec) Marshal(c ctx.Context, val interface{}) ([]byte, error) {
	if val == nil {
		return nil, nil
	}
	srl, ok := val.(core.Serializable)
	if ok {
		return cdc.MarshalSerializable(c, srl)
	}
	var b []byte = make([]byte, 0, 500)
	enc := codec.NewEncoderBytes(&b, cdc.h)
	err := enc.Encode(val)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (cdc *BinaryCodec) Unmarshal(c ctx.Context, arr []byte, val interface{}) error {
	if arr == nil {
		return nil
	}
	srl, ok := val.(core.Serializable)
	if ok {
		return cdc.UnmarshalSerializable(c, arr, srl)
	}
	dec := codec.NewDecoderBytes(arr, cdc.h)
	err := dec.Decode(val)
	if err != nil {
		return err
	}
	return nil
}

func (codec *BinaryCodec) UnmarshalSerializable(c ctx.Context, arr []byte, obj core.Serializable) error {
	return UnmarshalSerializable(c, codec, nil, arr, obj)
}

func (codec *BinaryCodec) MarshalSerializable(c ctx.Context, obj core.Serializable) ([]byte, error) {
	return MarshalSerializable(c, codec, nil, obj)
}

func (codec *BinaryCodec) UnmarshalSerializableProps(c ctx.Context, arr []byte, obj core.Serializable, props map[string]interface{}) error {
	return UnmarshalSerializableProps(c, codec, nil, arr, obj, props)
}

func (codec *BinaryCodec) MarshalSerializableProps(c ctx.Context, obj core.Serializable, props map[string]interface{}) ([]byte, error) {
	return MarshalSerializableProps(c, codec, nil, obj, props)
}
