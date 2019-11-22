package core

import (
	"fmt"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"reflect"
	"strings"
)

//service method for doing various tasks
func newObjectFactory(ctx ctx.Context, objLoader *objectLoader, objectName string, objectCreator core.ObjectCreator, objectCollectionCreator core.ObjectCollectionCreator, metadata core.Info) (*objectFactory, error) {
	if objectCreator != nil {
		fac := &objectFactory{objectCreator: objectCreator, objectCollectionCreator: objectCollectionCreator, metadata: metadata, objectName: objectName}
		err := fac.analyzeObject(ctx, objLoader)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		return fac, nil
	} else {
		return nil, fmt.Errorf("Nil value of Creator")
	}
}

func newObjectType(svrctx ctx.Context, objLoader *objectLoader, objectName string, obj interface{}, metadata core.Info) (*objectFactory, error) {
	typ := reflect.TypeOf(obj)
	slice := reflect.MakeSlice(reflect.SliceOf(typ), 0, 0)
	slicTyp := slice.Type()
	objectCreator := func(cx ctx.Context) interface{} {
		return reflect.New(typ).Interface()
	}
	collectCreator := func(cx ctx.Context, length int) interface{} {
		k := reflect.MakeSlice(reflect.SliceOf(typ), length, length)
		// Create a pointer to a slice value and set it to the slice
		x := reflect.New(slicTyp)
		x.Elem().Set(k)
		return x.Interface()
	}
	fac := &objectFactory{objectCreator: objectCreator, objectCollectionCreator: collectCreator, metadata: metadata, objectName: objectName}
	err := fac.analyzeObject(svrctx, objLoader)
	if err != nil {
		return nil, errors.WrapError(svrctx, err)
	}
	return fac, nil
}

type objectFactory struct {
	objectCreator           core.ObjectCreator
	objectCollectionCreator core.ObjectCollectionCreator
	metadata                core.Info
	fieldsToInit            map[string]*objectFactory
	objectName              string
}

//Creates object
func (factory *objectFactory) CreateObject(ctx ctx.Context) interface{} {
	obj := factory.objectCreator(ctx)
	factory.initObject(ctx, obj)
	return obj
}

func (factory *objectFactory) initObject(ctx ctx.Context, obj interface{}) {
	if factory.fieldsToInit != nil {
		for fieldName, fieldObjFac := range factory.fieldsToInit {
			entVal := reflect.ValueOf(obj)
			f := entVal.Elem().FieldByName(fieldName)
			fieldobj := fieldObjFac.CreateObject(ctx)
			fobj := reflect.ValueOf(fieldobj)
			pref := fobj.Elem().FieldByName("P_ref")
			if pref.IsValid() {
				pref.Set(entVal.Convert(pref.Type()))
			}
			f.Set(fobj.Convert(f.Type()))
		}
	}
	constructible, ok := obj.(constructible)
	if (constructible != nil) && ok {
		constructible.Constructor()
	}
}

//Creates collection
func (factory *objectFactory) CreateObjectCollection(ctx ctx.Context, length int) interface{} {
	return factory.objectCollectionCreator(ctx, length)
}

func (factory *objectFactory) Info() core.Info {
	return factory.metadata
}

func (factory *objectFactory) analyzeObject(ctx ctx.Context, objLoader *objectLoader) error {
	log.Debug(ctx, "Analysing object", "Name", factory.objectName)

	obj := factory.objectCreator(ctx)
	objType := reflect.TypeOf(obj).Elem()
	tagName := "laatoo"
	//typeToAnalyze := factory.objType
	for i := 0; i < objType.NumField(); i++ {
		fld := objType.Field(i)
		log.Trace(ctx, "Analysing field", "Field", fld.Name)
		laatooTagVal, ok := fld.Tag.Lookup(tagName)
		if ok {
			tags := strings.Split(laatooTagVal, ",")
			var initobj string
			var audit bool
			var multitenant bool
			var softdelete bool
			for _, tag := range tags {
				tag = strings.TrimSpace(tag)
				if strings.HasPrefix(tag, "initialize") {
					_, err := fmt.Sscanf(tag, "initialize=%s", &initobj)
					if err != nil {
						log.Warn(ctx, "Laatoo tag not correctly formatted", "object", factory.objectName, "tag", tag)
						return err
					}
				} else {
					switch tag {
					case "auditable":
						{
							audit = true
						}
					case "softdelete":
						{
							softdelete = true
						}
					case "multitenant":
						{
							multitenant = true
						}
					}
				}
			}

			if fld.Name == "Storable" {
				var suffix string
				if multitenant {
					suffix = "MT"
				}
				if initobj == "" {
					if audit {
						if softdelete {
							initobj = "laatoo/server/data.SoftDeleteAuditable" + suffix
						} else {
							initobj = "laatoo/server/data.HardDeleteAuditable" + suffix
						}
					} else {
						if softdelete {
							initobj = "laatoo/server/data.SoftDeleteStorable" + suffix
						} else {
							initobj = "laatoo/server/data.AbstractStorable" + suffix
						}
					}
				}
			}

			if initobj != "" {
				objfac, ok := objLoader.objectsFactoryRegister[initobj]
				log.Debug(ctx, "assigning initialize initialization to field ", "initobj", initobj, "ok", ok)
				if ok {
					if factory.fieldsToInit == nil {
						factory.fieldsToInit = make(map[string]*objectFactory)
					}
					factory.fieldsToInit[fld.Name] = objfac.(*objectFactory)
					log.Trace(ctx, "assigning initialize initialization to field ", "Object", factory.objectName, "field", fld.Name, "assignment", initobj)
				}
			}
		}
	}
	return nil
}

type constructible interface {
	Constructor()
}
