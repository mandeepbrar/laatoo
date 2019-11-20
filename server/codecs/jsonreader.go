package codecs

import (
	"fmt"
	"io"
	"io/ioutil"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/log"
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

func (rdr *JsonReader) ReadBytes(ctx ctx.Context, cdc core.Codec, prop string) ([]byte, error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return nil, nil
	}
	return v.MarshalTo(nil), nil
}

func (rdr *JsonReader) ReadProp(ctx ctx.Context, cdc core.Codec, prop string) (core.SerializableReader, error) {
	//keys := append(rdr.keys, prop)
	v := rdr.root.Get(prop)
	return &JsonReader{root: v}, nil
}

func (rdr *JsonReader) ReadInt(ctx ctx.Context, cdc core.Codec, prop string, val *int) error {
	intV := rdr.root.GetInt(prop)
	*val = intV
	return nil
}

func (rdr *JsonReader) ReadInt64(ctx ctx.Context, cdc core.Codec, prop string, val *int64) error {
	intV := rdr.root.GetInt64(prop)
	*val = intV
	return nil
}

func (rdr *JsonReader) ReadString(ctx ctx.Context, cdc core.Codec, prop string, val *string) error {
	byts := rdr.root.GetStringBytes(prop)
	if byts != nil {
		*val = string(byts)
	}
	return nil
}

func (rdr *JsonReader) ReadFloat32(ctx ctx.Context, cdc core.Codec, prop string, val *float32) error {
	fltV := rdr.root.GetFloat64(prop)
	*val = float32(fltV)
	return nil
}

func (rdr *JsonReader) ReadFloat64(ctx ctx.Context, cdc core.Codec, prop string, val *float64) error {
	fltV := rdr.root.GetFloat64(prop)
	*val = fltV
	return nil
}

func (rdr *JsonReader) ReadBool(ctx ctx.Context, cdc core.Codec, prop string, val *bool) error {
	boolV := rdr.root.GetBool(prop)
	*val = boolV
	return nil
}

func (rdr *JsonReader) ReadObject(ctx ctx.Context, cdc core.Codec, prop string, val interface{}) (err error) {
	v := rdr.root.Get(prop)
	if v == nil {
		return nil
	}
	err = rdr.readObject(ctx, cdc, v, val)
	return
}

func (rdr *JsonReader) ReadMap(ctx ctx.Context, cdc core.Codec, prop string, val *map[string]interface{}) error {
	//valMap := *val
	v := rdr.root.Get(prop)
	if v == nil {
		return nil
	}
	bys := v.MarshalTo(nil)
	return cdc.Unmarshal(ctx, bys, val)
}

func (rdr *JsonReader) ReadArray(ctx ctx.Context, cdc core.Codec, prop string, val interface{}) error {

	arrV := rdr.root.Get(prop)
	if arrV == nil {
		return nil
	}
	return rdr.readArray(ctx, cdc, arrV, val)
}

func (rdr *JsonReader) ReadTime(ctx ctx.Context, cdc core.Codec, prop string, val *time.Time) error {
	v := rdr.root.Get(prop)
	if v == nil {
		return nil
	}
	bys := v.MarshalTo(nil)
	tim, err := time.Parse(time.RFC3339, string(bys))
	if err != nil {
		return err
	}
	*val = tim
	return nil
}

func (rdr *JsonReader) readObject(ctx ctx.Context, cdc core.Codec, val *fastjson.Value, obj interface{}) (err error) {

	objVal := reflect.ValueOf(obj)
	if objVal.Kind() == reflect.Array {
		return rdr.readArray(ctx, cdc, val, obj)
	}

	srl, ok := obj.(core.Serializable)
	if ok {
		newrdr := &JsonReader{root: val}
		err = srl.ReadAll(ctx, cdc, newrdr)
	} else {
		byts := val.MarshalTo(nil)
		log.Error(ctx, "json ignitor unmarshalling", "bys", string(byts))
		err = json.Unmarshal(byts, obj)
		log.Error(ctx, "json ignitor unmarshalling", "err", err, " obj", fmt.Sprint(obj))
	}
	return
}

func (rdr *JsonReader) createCollection(ctx ctx.Context, cdc core.Codec, val interface{}, len int) (objType string, err error) {
	collVal := reflect.ValueOf(val)
	if collVal.Kind() != reflect.Array {
		return "", errWrongObject
	}

	reqCtx := ctx.(core.RequestContext)
	objType = reqCtx.GetRegName(val)

	collection, err := reqCtx.CreateCollection(objType, len)
	if err == nil {
		reflect.ValueOf(val).Set(reflect.ValueOf(collection))
	}
	return
}

func (rdr *JsonReader) readArray(ctx ctx.Context, cdc core.Codec, arrV *fastjson.Value, val interface{}) error {
	vArr, err := arrV.Array()
	if err != nil {
		return err
	}
	objType, err := rdr.createCollection(ctx, cdc, val, len(vArr))
	if err != nil {
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
