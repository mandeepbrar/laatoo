package core

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/log"
	"reflect"
)

//service method for doing various tasks
func newObjectFactory(ctx ctx.Context, objLoader *objectLoader, objectName string, objectCreator core.ObjectCreator, objectCollectionCreator core.ObjectCollectionCreator, metadata core.Info) *objectFactory {
	if objectCreator != nil {
		fac := &objectFactory{objectCreator: objectCreator, objectCollectionCreator: objectCollectionCreator, metadata: metadata, objectName: objectName}
		fac.analyzeObject(ctx, objLoader)
		return fac
	} else {
		panic("Could not register object factory. Creator is nil.")
	}
}

func newObjectType(svrctx ctx.Context, objLoader *objectLoader, objectName string, obj interface{}, metadata core.Info) *objectFactory {
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
	fac.analyzeObject(svrctx, objLoader)
	return fac
}

type objectFactory struct {
	objectCreator           core.ObjectCreator
	objectCollectionCreator core.ObjectCollectionCreator
	metadata                core.Info
	fieldsToInit            map[string]*objectFactory
	constructor             reflect.Value
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
			entVal := reflect.ValueOf(obj).Elem()
			f := entVal.FieldByName(fieldName)
			fieldobj := fieldObjFac.CreateObject(ctx)
			f.Set(reflect.ValueOf(fieldobj))
		}
	}
	if factory.constructor.IsValid() {
		objVal := reflect.ValueOf(obj)
		factory.constructor.Call([]reflect.Value{objVal})
	}
}

//Creates collection
func (factory *objectFactory) CreateObjectCollection(ctx ctx.Context, length int) interface{} {
	return factory.objectCollectionCreator(ctx, length)
}

func (factory *objectFactory) Info() core.Info {
	return factory.metadata
}

func (factory *objectFactory) analyzeObject(ctx ctx.Context, objLoader *objectLoader) {
	log.Debug(ctx, "Analysing object", "Name", factory.objectName)

	obj := factory.objectCreator(ctx)
	objType := reflect.TypeOf(obj).Elem()
	objPtrType := reflect.TypeOf(obj)

	//typeToAnalyze := factory.objType
	for i := 0; i < objType.NumField(); i++ {
		fld := objType.Field(i)
		if fld.Type.Kind() == reflect.Ptr {
			objName, ok := fld.Tag.Lookup("initialize")
			if ok {
				objfac, ok := objLoader.objectsFactoryRegister[objName]
				if ok {
					if factory.fieldsToInit == nil {
						factory.fieldsToInit = make(map[string]*objectFactory)
					}
					factory.fieldsToInit[objName] = objfac.(*objectFactory)
					log.Trace(ctx, "assigning initialization to field ", "Object", factory.objectName, "field", fld.Name, "objName", objName, "fields to init", factory.fieldsToInit)
				}
			}
		}
	}
	constructor, ok := objPtrType.MethodByName("Constructor")
	if ok {
		log.Trace(ctx, "assigning constructor to object ", "Object", factory.objectName)
		factory.constructor = constructor.Func
	}
}
