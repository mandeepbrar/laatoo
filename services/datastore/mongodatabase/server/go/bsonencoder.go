package main

import (
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/ctx"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
)

type StorableSerializer struct {
	defaultReg  *bsoncodec.Registry
	logCtx      ctx.Context
	timeEncoder bsoncodec.ValueEncoder
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
	}
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
		vw, err = docWriter.WriteDocumentElement("CreatedBy")
		if err != nil {
			return err
		}
		err = vw.WriteString(auditable.GetCreatedBy())
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
	return decoder.DecodeValue(ctx, reader, val)
}
