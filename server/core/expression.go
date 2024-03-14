package core

import "laatoo.io/sdk/datatypes"

type Expression struct {
	Value      interface{}
	Expression string
	Type       string
	DType      datatypes.DataType
}
