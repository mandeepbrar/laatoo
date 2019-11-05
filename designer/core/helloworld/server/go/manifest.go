package main

import (
	"fmt"
	"laatoo/sdk/server/core"
	"reflect"

	"github.com/gogo/protobuf/proto"
)

type Name struct {
	//data.SerializableBase `json:",inline" protobuf:"group,51,name=SerializableBase,proto3" bson:",inline" `
	Name     string `protobuf:"bytes,5,name=name,proto3" json:"name,omitempty"`
	Newfield string `protobuf:"bytes,6,name=newfield,proto3"`
}

func (b *Name) Reset() {
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(b).Elem())
	*b = reflect.New(reflect.TypeOf(b).Elem()).Elem().Interface().(Name)
	fmt.Println(b)
	fmt.Println(*b)
}

func (m *Name) String() string { return proto.CompactTextString(m) }

func (*Name) ProtoMessage() {}

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
		core.PluginComponent{Name: "Name", Object: Name{}},
		core.PluginComponent{Name: "helloworld", Object: HelloWorld{}},
	}
}
