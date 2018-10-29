package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"time"
)

type objectDataType int

const (
	__stringmap objectDataType = iota
	__stringsmap
	__maptype
	__datetime
	__config
	__bytes
	__files
	__inttype
	__stringtype
	__stringarr
	__booltype
	__custom
	__none
)

type param struct {
	name                        string
	ptype                       string
	oDataType                   objectDataType
	dataObjectCreator           core.ObjectCreator
	dataObjectCollectionCreator core.ObjectCollectionCreator
	collection                  bool
	isStream                    bool
	required                    bool
	value                       interface{}
}

func newParam(ctx core.ServerContext, name, ptype string, collection, stream, required bool) (*param, error) {
	dataObjectType := convertDataType(ptype)
	p := &param{name, ptype, dataObjectType, nil, nil, collection, stream, required, nil}
	if dataObjectType == __custom {
		if collection {
			dataObjectCollectionCreator, err := ctx.GetObjectCollectionCreator(ptype)
			if err != nil {
				return nil, errors.RethrowError(ctx, errors.CORE_ERROR_BAD_CONF, err, "No such object", ptype)
			}
			p.dataObjectCollectionCreator = dataObjectCollectionCreator
		} else {
			dataObjectCreator, err := ctx.GetObjectCreator(ptype)
			if err != nil {
				return nil, errors.RethrowError(ctx, errors.CORE_ERROR_BAD_CONF, err, "No such object", ptype)
			}
			p.dataObjectCreator = dataObjectCreator
		}
	}

	return p, nil
}

func (p *param) clone() *param {
	return &param{p.name, p.ptype, p.oDataType, p.dataObjectCreator, p.dataObjectCollectionCreator, p.collection, p.isStream, p.required, p.value}
}

func (p *param) GetName() string {
	return p.name
}

func (p *param) IsCollection() bool {
	return p.collection
}

func (p *param) IsRequired() bool {
	return p.required
}

func (p *param) GetDataType() string {
	return p.ptype
}

func (p *param) IsStream() bool {
	return p.isStream
}

func (p *param) GetValue() interface{} {
	return p.value
}

func cloneParamsMap(params map[string]core.Param) map[string]core.Param {
	cparams := make(map[string]core.Param, len(params))
	for k, v := range params {
		cparams[k] = v.(*param).clone()
	}
	return cparams
}

func readParamsConf(ctx core.ServerContext, paramsConf config.Config) (map[string]core.Param, error) {
	params := make(map[string]core.Param)
	paramNames := paramsConf.AllConfigurations(ctx)
	for _, paramName := range paramNames {
		paramDesc, _ := paramsConf.GetSubConfig(ctx, paramName)
		collection, _ := paramDesc.GetBool(ctx, SVCPARAMCOLLECTION)
		paramtype, _ := paramDesc.GetString(ctx, SVCPARAMTYPE)
		stream, _ := paramDesc.GetBool(ctx, SVCSTREAM)
		paramreqd, _ := paramDesc.GetBool(ctx, SVCPARAMREQD)
		p, err := newParam(ctx, paramName, paramtype, collection, stream, paramreqd)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		} else {
			params[paramName] = p
		}
	}
	return params, nil
}

func convertDataType(dtype string) objectDataType {
	switch dtype {
	case "":
		return __none
	case config.OBJECTTYPE_STRINGMAP:
		return __stringmap
	case config.OBJECTTYPE_STRINGSMAP:
		return __stringsmap
	case config.OBJECTTYPE_BYTES:
		return __bytes
	case config.OBJECTTYPE_STRING:
		return __stringtype
	case config.OBJECTTYPE_STRINGARR:
		return __stringarr
	case config.OBJECTTYPE_BOOL:
		return __booltype
	case config.OBJECTTYPE_FILES:
		return __files
	case config.OBJECTTYPE_DATETIME:
		return __datetime
	case config.OBJECTTYPE_INT:
		return __inttype
	case config.OBJECTTYPE_CONFIG:
		return __config
	case config.OBJECTTYPE_MAP:
		return __maptype
	default:
		return __custom
	}
}

func (p *param) setValue(ctx ctx.Context, val interface{}, codec core.Codec, encoded bool) error {
	var reqData interface{}
	resPtr := false
	ok := false
	switch p.oDataType {
	case __stringmap:
		if encoded {
			reqData = make(map[string]interface{}, 10)
		} else {
			p.value, ok = val.(map[string]interface{})
		}
	case __stringsmap:
		if encoded {
			reqData = make(map[string]string, 10)
		} else {
			p.value, ok = val.(map[string]string)
		}
	case __bytes:
		p.value, ok = val.([]byte)
	case __inttype:
		p.value, ok = val.(int)
	case __datetime:
		if encoded {
			t, ok := val.(string)
			if ok {
				tvalue, err := time.Parse(time.RFC1123, t)
				if err != nil {
					return err
				} else {
					p.value = tvalue
				}
			}
		} else {
			p.value, ok = val.(time.Time)
		}
	case __stringtype:
		p.value, ok = val.(string)
		log.Error(ctx, "set values", "p.value", p.value, "ok", ok)
	case __stringarr:
		p.value, ok = val.([]string)
	case __booltype:
		p.value, ok = val.(bool)
	default:
		if encoded {
			if p.IsCollection() {
				reqData = p.dataObjectCollectionCreator(5)
			} else {
				reqData = p.dataObjectCreator()
			}
			resPtr = true
		} else {
			p.value = val
			ok = true
		}
	}

	/** decode encoded objects**/
	if encoded && !ok {
		var reqBytes []byte
		reqBytes, ok = val.([]byte)
		if ok {

			var err error
			if resPtr {
				err = codec.Unmarshal(ctx, reqBytes, reqData)
			} else {
				err = codec.Unmarshal(ctx, reqBytes, &reqData)
			}
			log.Trace(ctx, "unmarshalling bytes", "val", reqData, " err", err)
			if err != nil {
				return err
			} else {
				p.value = reqData
			}
		}
	}
	if !ok {
		return errors.BadArg(ctx, p.name)
	}

	return nil
}
