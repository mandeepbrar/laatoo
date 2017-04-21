package common

import (
	"laatoo/sdk/core"

	"github.com/ugorji/go/codec"
)

type CacheEncoder struct {
	enc *codec.Encoder
	dec *codec.Decoder
}

func NewCacheEncoder(ctx core.ServerContext, encType string) *CacheEncoder {
	var h codec.Handle
	switch encType {
	case "json":
		{
			h = new(codec.JsonHandle)
		}
	case "binary":
		{
			h = new(codec.BincHandle)
		}
	default:
		{
			h = new(codec.BincHandle)
		}
	}
	var b []byte = make([]byte, 0, 64)
	cEncoder := &CacheEncoder{}
	cEncoder.enc = codec.NewEncoderBytes(&b, h)
	cEncoder.dec = codec.NewDecoderBytes(b, h)
	return cEncoder
}

func (enc *CacheEncoder) Encode(val interface{}) ([]byte, error) {
	if val == nil {
		return nil, nil
	}
	var b []byte = make([]byte, 0, 500)
	enc.enc.ResetBytes(&b)
	err := enc.enc.Encode(val)
	if err != nil {
		return nil, err
	}
	return b, nil

}

func (enc *CacheEncoder) Decode(arr []byte, val interface{}) error {
	if arr == nil {
		return nil
	}
	enc.dec.ResetBytes(arr)
	err := enc.dec.Decode(val)
	if err != nil {
		return err
	}
	return nil
}
