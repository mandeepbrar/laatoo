package codecs

import (
	"encoding/json"
)

type JsonCodec struct {
}

func NewJsonCodec() *JsonCodec {
	return &JsonCodec{}
}

func (cdc *JsonCodec) Marshal(val interface{}) ([]byte, error) {
	if val == nil {
		return nil, nil
	}
	return json.Marshal(val)
}

func (cdc *JsonCodec) Unmarshal(arr []byte, val interface{}) error {
	if arr == nil {
		return nil
	}
	return json.Unmarshal(arr, val)
}
