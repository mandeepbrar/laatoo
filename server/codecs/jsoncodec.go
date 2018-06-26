package codecs

import (
	"encoding/json"
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
	return json.Marshal(val)
}

func (cdc *JsonCodec) Unmarshal(c ctx.Context, arr []byte, val interface{}) error {
	if arr == nil {
		return nil
	}
	return json.Unmarshal(arr, val)
}
