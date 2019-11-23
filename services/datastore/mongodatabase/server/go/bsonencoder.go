package main

import (
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
)

type StorableSerializer struct {
	defaultReg  *bsoncodec.Registry
	serverCtx   core.ServerContext
	timeEncoder bsoncodec.ValueEncoder
	timeDecoder bsoncodec.ValueDecoder
}

func (enc *StorableSerializer) EncodeValue(ctx bsoncodec.EncodeContext, writer bsonrw.ValueWriter, val reflect.Value) error {
	if val.Kind() == reflect.Struct {
		storfld := val.FieldByName("Storable")
		if storfld.IsValid() {
			stor, ok := storfld.Interface().(data.Storable)
			if ok {
				valType := val.Type()
				docWriter, err := writer.WriteDocument()
				if err != nil {
					return err
				}
				for i := 0; i < val.NumField(); i++ {
					fieldName := valType.Field(i).Name
					field := val.Field(i)
					if field == storfld {
						err = enc.writeStorValues(ctx, docWriter, stor)
						if err != nil {
							return err
						}
					} else {
						valWriter, err := docWriter.WriteDocumentElement(fieldName)
						if err != nil {
							return err
						}
						enc.writeDefault(ctx, valWriter, field)
					}
				}
				err = docWriter.WriteDocumentEnd()
				if err != nil {
					return err
				}
				return nil
			}
		}
	}

	return enc.writeDefault(ctx, writer, val)
}

func (enc *StorableSerializer) writeStorValues(ctx bsoncodec.EncodeContext, docWriter bsonrw.DocumentWriter, stor data.Storable) error {
	vw, err := docWriter.WriteDocumentElement("Id")
	if err != nil {
		return err
	}
	err = vw.WriteString(stor.GetId())
	if err != nil {
		return err
	} //deleted, is new???
	auditable, ok := stor.(data.Auditable)
	if ok {
		vw, err = docWriter.WriteDocumentElement("CreatedBy")
		if err != nil {
			return err
		}
		err = vw.WriteString(auditable.GetCreatedBy())
		if err != nil {
			return err
		}
		vw, err = docWriter.WriteDocumentElement("UpdatedBy")
		if err != nil {
			return err
		}
		err = vw.WriteString(auditable.GetUpdatedBy())
		if err != nil {
			return err
		}
		vw, err = docWriter.WriteDocumentElement("CreatedAt")
		if err != nil {
			return err
		}
		err = enc.timeEncoder.EncodeValue(ctx, vw, reflect.ValueOf(auditable.GetCreatedAt()))
		if err != nil {
			return err
		}

		vw, err = docWriter.WriteDocumentElement("UpdatedAt")
		if err != nil {
			return err
		}
		err = enc.timeEncoder.EncodeValue(ctx, vw, reflect.ValueOf(auditable.GetUpdatedAt()))
		if err != nil {
			return err
		}
	}
	softdeletable, ok := stor.(data.SoftDeletable)
	if ok {
		vw, err = docWriter.WriteDocumentElement(softdeletable.SoftDeleteField())
		if err != nil {
			return err
		}
		err = vw.WriteBoolean(softdeletable.IsDeleted())
		if err != nil {
			return err
		}
	}
	return nil
}

func (enc *StorableSerializer) writeDefault(ctx bsoncodec.EncodeContext, writer bsonrw.ValueWriter, val reflect.Value) error {
	encoder, err := enc.defaultReg.LookupEncoder(val.Type())
	if err != nil {
		return err
	}
	return encoder.EncodeValue(ctx, writer, val)
}

func (enc *StorableSerializer) DecodeValue(ctx bsoncodec.DecodeContext, reader bsonrw.ValueReader, val reflect.Value) error {
	decoder, err := enc.defaultReg.LookupDecoder(val.Type())
	if err != nil {
		return err
	}

	if val.Kind() == reflect.Ptr {
		_, ok := val.Interface().(data.Storable)
		if ok {
			objType := enc.serverCtx.GetRegName(val.Interface())
			obj, _ := enc.serverCtx.CreateObject(objType)
			val.Set(reflect.ValueOf(obj))
		}
	}

	if val.Kind() == reflect.Struct {
		storfld := val.FieldByName("Storable")
		if storfld.IsValid() {
			stor, ok := storfld.Interface().(data.Storable)
			audit, aok := stor.(data.Auditable)
			if ok {
				dr, err := reader.ReadDocument()
				if err != nil {
					return err
				}

				for {
					name, vr, err := dr.ReadElement()
					if err == bsonrw.ErrEOD {
						return nil
					}
					if err != nil {
						return err
					}

					switch name {
					case "Id":
						{
							strval, err := vr.ReadString()
							if err != nil {
								return err
							}
							log.Println("Id", strval)
							stor.SetId(strval)
						}
					case "CreatedBy":
						{
							if aok {
								strval, err := vr.ReadString()
								if err != nil {
									return err
								}
								audit.SetCreatedBy(strval)
							}
						}
					case "UpdatedBy":
						{
							if aok {
								strval, err := vr.ReadString()
								if err != nil {
									return err
								}
								audit.SetUpdatedBy(strval)
							}
						}
					/*case "CreatedAt":
						{
							if aok {
								t := time.Now()
								err := enc.timeDecoder.DecodeValue(ctx, vr, reflect.ValueOf(t))
								if err != nil {
									return err
								}
								audit.SetCreatedAt(t)
							}
						}
					case "UpdatedAt":
						{
							if aok {
								t := time.Now()
								err := enc.timeDecoder.DecodeValue(ctx, vr, reflect.ValueOf(t))
								if err != nil {
									return err
								}
								audit.SetCreatedAt(t)
							}
						}*/
					default:
						{
							log.Println("decoding field", name)
							fldToSet := val.FieldByName(name)
							if fldToSet.IsValid() {
								err = enc.DecodeValue(ctx, vr, fldToSet)
								if err != nil {
									return err
								}
							} else {
								err = vr.Skip()
								log.Println("skipping field", name)
							}
						}
					}
				}
			}
		}
	}

	return decoder.DecodeValue(ctx, reader, val)
}
