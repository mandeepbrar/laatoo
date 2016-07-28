package common

import "github.com/ugorji/go/codec"

var b []byte = make([]byte, 0, 64)
var h codec.Handle = new(codec.BincHandle)
var enc *codec.Encoder = codec.NewEncoderBytes(&b, h)
var dec *codec.Decoder = codec.NewDecoderBytes(b, h)

func Encode(val interface{}) ([]byte, error) {
	if val == nil {
		return nil, nil
	}
	var b []byte = make([]byte, 0, 500)
	enc.ResetBytes(&b)
	err := enc.Encode(val)
	if err != nil {
		return nil, err
	}
	return b, nil
	/*	var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		err := enc.Encode(val)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil*/

}

func Decode(arr []byte, val interface{}) error {
	if arr == nil {
		return nil
	}
	//	dec := gob.NewDecoder(bytes.NewReader(arr))
	dec.ResetBytes(arr)
	err := dec.Decode(val)

	//err := dec.Decode(val)
	if err != nil {
		return err
	}
	return nil
}
