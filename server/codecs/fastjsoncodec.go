package codecs

import (
	"laatoo/sdk/ctx"

	"github.com/ugorji/go/codec"
)

type FastJsonCodec struct {
	enc *codec.Encoder
	dec *codec.Decoder
	h   codec.Handle
}

func NewFastJsonCodec() *FastJsonCodec {
	return &FastJsonCodec{h: new(codec.JsonHandle)}
}

func (cdc *FastJsonCodec) Marshal(c ctx.Context, val interface{}) ([]byte, error) {
	if val == nil {
		return nil, nil
	}
	var b []byte = make([]byte, 0, 500)
	enc := codec.NewEncoderBytes(&b, cdc.h)
	err := enc.Encode(val)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (cdc *FastJsonCodec) Unmarshal(c ctx.Context, arr []byte, val interface{}) error {
	if arr == nil {
		return nil
	}
	dec := codec.NewDecoderBytes(arr, cdc.h)
	err := dec.Decode(val)
	if err != nil {
		return err
	}
	return nil
}
