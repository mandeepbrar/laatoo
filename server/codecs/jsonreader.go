package codecs

import (
	"io"
	"io/ioutil"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"reflect"
	"time"

	"github.com/valyala/fastjson"
)

func NewJsonReader(c ctx.Context, arr []byte) (core.SerializableReader, error) {
	var p fastjson.Parser
	v, err := p.ParseBytes(arr)
	if err != nil {
		return nil, err
	}
	return &JsonReader{root: v}, nil
}

//Not full support for reader
func NewJsonStreamReader(c ctx.Context, rdr io.Reader) (core.SerializableReader, error) {
	arr, err := ioutil.ReadAll(rdr)
	if err != nil {
		return nil, err
	}
	var p fastjson.Parser
	v, err := p.ParseBytes(arr)
	if err != nil {
		return nil, err
	}
	return &JsonReader{root: v}, nil
}

func internalJsonReader(c ctx.Context, val core.SerializableReader) (core.SerializableReader, error) {
	jsonref := val.(*JsonReader)
	return &JsonReader{root: jsonref.root}, nil
}

type JsonReader struct {
	io.Reader
	//	keys []string
	root *fastjson.Value
}

func (rdr *JsonReader) Start() error {
	return nil
}

func (rdr *JsonReader) Bytes() []byte {
	return rdr.root.MarshalTo(nil)
}

func (rdr *JsonReader) ReadBytes(ctx ctx.Context, cdc core.Codec, prop string) ([]byte, bool, error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return nil, false, nil
	}
	return v.MarshalTo(nil), true, nil
}

func (rdr *JsonReader) ReadProp(ctx ctx.Context, cdc core.Codec, prop string) (core.SerializableReader, bool, error) {
	//keys := append(rdr.keys, prop)
	v := rdr.root.Get(prop)
	if v == nil {
		return nil, false, nil
	}
	return &JsonReader{root: v}, true, nil
}

func (rdr *JsonReader) ReadInt(ctx ctx.Context, cdc core.Codec, prop string, val *int) (b bool, err error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return false, nil
	}
	*val, err = v.Int()
	return true, err
}

func (rdr *JsonReader) ReadInt64(ctx ctx.Context, cdc core.Codec, prop string, val *int64) (b bool, err error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return false, nil
	}
	*val, err = v.Int64()
	return true, nil
}

func (rdr *JsonReader) ReadString(ctx ctx.Context, cdc core.Codec, prop string, val *string) (bool, error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return false, nil
	}
	byts, err := v.StringBytes()

	if byts != nil {
		*val = string(byts)
	}
	return true, err
}

func (rdr *JsonReader) ReadFloat32(ctx ctx.Context, cdc core.Codec, prop string, val *float32) (b bool, err error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return false, nil
	}
	flt, err := v.Float64()
	*val = float32(flt)
	return true, err
}

func (rdr *JsonReader) ReadFloat64(ctx ctx.Context, cdc core.Codec, prop string, val *float64) (b bool, err error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return false, nil
	}
	*val, err = v.Float64()
	return true, err
}

func (rdr *JsonReader) ReadBool(ctx ctx.Context, cdc core.Codec, prop string, val *bool) (b bool, err error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return false, nil
	}
	*val, err = v.Bool()
	return true, err
}

func (rdr *JsonReader) ReadObject(ctx ctx.Context, cdc core.Codec, prop string, val interface{}) (b bool, err error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return false, err
	}
	obj, _, err := rdr.createObject(ctx, cdc, val)
	if err != nil {
		return true, err
	}
	err = rdr.readObject(ctx, cdc, v, obj)
	return true, err
}

func (rdr *JsonReader) ReadMap(ctx ctx.Context, cdc core.Codec, prop string, val *map[string]interface{}) (bool, error) {
	//valMap := *val
	v := rdr.root.Get(prop)
	if v == nil {
		return false, nil
	}
	bys := v.MarshalTo(nil)
	return true, cdc.Unmarshal(ctx, bys, val)
}

func (rdr *JsonReader) ReadArray(ctx ctx.Context, cdc core.Codec, prop string, val interface{}) (bool, error) {

	arrV := rdr.root.Get(prop)
	if arrV == nil {
		return false, nil
	}
	return true, rdr.readArray(ctx, cdc, arrV, val)
}

func (rdr *JsonReader) ReadTime(ctx ctx.Context, cdc core.Codec, prop string, val *time.Time) (bool, error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return false, nil
	}
	bys := v.MarshalTo(nil)
	tim, err := time.Parse(time.RFC3339, string(bys))
	if err != nil {
		return true, err
	}
	*val = tim
	return true, nil
}

func (rdr *JsonReader) readObject(ctx ctx.Context, cdc core.Codec, val *fastjson.Value, obj interface{}) (err error) {

	objVal := reflect.ValueOf(obj)
	if objVal.IsZero() {
		return nil
	}
	if objVal.Kind() == reflect.Array {
		return rdr.readArray(ctx, cdc, val, obj)
	}
	srl, ok := obj.(core.Serializable)
	if ok {
		newrdr := &JsonReader{root: val}
		err = srl.ReadAll(ctx, cdc, newrdr)
	} else {
		byts := val.MarshalTo(nil)
		err = json.Unmarshal(byts, obj)
	}
	return
}

func (rdr *JsonReader) createObject(ctx ctx.Context, cdc core.Codec, val interface{}) (obj interface{}, objType string, err error) {
	objVal := reflect.ValueOf(val)
	if objVal.Kind() != reflect.Ptr {
		return nil, "", errWrongObject
	}

	reqCtx := ctx.(core.RequestContext)
	objType = reqCtx.GetRegName(val)

	if !objVal.Elem().IsNil() {
		return
	}

	obj, err = reqCtx.CreateObject(objType)
	if err == nil {
		objVal.Elem().Set(reflect.ValueOf(obj))
	} else {
		obj := reflect.New(objVal.Type())
		objVal.Elem().Set(reflect.ValueOf(obj))
	}
	return
}

func (rdr *JsonReader) createCollection(ctx ctx.Context, cdc core.Codec, val interface{}, len int) (objType string, reg bool, err error) {
	collVal := reflect.ValueOf(val)
	if collVal.Kind() != reflect.Ptr && collVal.Elem().Kind() != reflect.Array {
		return "", reg, errWrongObject
	}

	reqCtx := ctx.(core.RequestContext)
	objType = reqCtx.GetRegName(val)

	collection, err := reqCtx.CreateCollection(objType, len)
	if err == nil {
		collVal.Elem().Set(reflect.ValueOf(collection))
		reg = true
	} else {
		err = nil
		return objType, reg, nil
	}
	return
}

func (rdr *JsonReader) readArray(ctx ctx.Context, cdc core.Codec, arrV *fastjson.Value, val interface{}) error {
	vArr, err := arrV.Array()
	if err != nil {
		return err
	}
	objType, reg, err := rdr.createCollection(ctx, cdc, val, len(vArr))
	if err != nil {
		return err
	}
	if !reg {
		byts := arrV.MarshalTo(nil)
		err = json.Unmarshal(byts, val)
		return err
	}
	collVal := reflect.ValueOf(val)
	for i, v := range vArr {
		obj, err := ctx.(core.RequestContext).CreateObject(objType)
		if err != nil {
			return err
		}
		err = rdr.readObject(ctx, cdc, v, obj)
		if err != nil {
			return err
		}
		iVal := collVal.Index(i)
		iVal.Set(reflect.ValueOf(obj))
	}
	return nil
}
