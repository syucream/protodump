// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/protosl.proto

/*
Package protosl is a generated protocol buffer package.

It is generated from these files:
	internal/protosl.proto

It has these top-level messages:
	Example
	Child
*/
package protosl

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Example_Num int32

const (
	Example_ZERO Example_Num = 0
	Example_ONE  Example_Num = 1
)

var Example_Num_name = map[int32]string{
	0: "ZERO",
	1: "ONE",
}
var Example_Num_value = map[string]int32{
	"ZERO": 0,
	"ONE":  1,
}

func (x Example_Num) String() string {
	return proto.EnumName(Example_Num_name, int32(x))
}
func (Example_Num) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Example struct {
	Uint64Val   uint64        `protobuf:"varint,1,opt,name=uint64_val,json=uint64Val" json:"uint64_val,omitempty"`
	StringVal   string        `protobuf:"bytes,2,opt,name=string_val,json=stringVal" json:"string_val,omitempty"`
	Fixed64Val  uint64        `protobuf:"fixed64,3,opt,name=fixed64_val,json=fixed64Val" json:"fixed64_val,omitempty"`
	Fixed32Val  uint32        `protobuf:"fixed32,4,opt,name=fixed32_val,json=fixed32Val" json:"fixed32_val,omitempty"`
	EnumVal     Example_Num   `protobuf:"varint,5,opt,name=enum_val,json=enumVal,enum=Example_Num" json:"enum_val,omitempty"`
	ChildVal    *Child        `protobuf:"bytes,6,opt,name=child_val,json=childVal" json:"child_val,omitempty"`
	RUint64Val  []uint64      `protobuf:"varint,101,rep,packed,name=r_uint64_val,json=rUint64Val" json:"r_uint64_val,omitempty"`
	RStringVal  []string      `protobuf:"bytes,102,rep,name=r_string_val,json=rStringVal" json:"r_string_val,omitempty"`
	RFixed64Val []uint64      `protobuf:"fixed64,103,rep,packed,name=r_fixed64_val,json=rFixed64Val" json:"r_fixed64_val,omitempty"`
	RFixed32Val []uint32      `protobuf:"fixed32,104,rep,packed,name=r_fixed32_val,json=rFixed32Val" json:"r_fixed32_val,omitempty"`
	REnumVal    []Example_Num `protobuf:"varint,105,rep,packed,name=r_enum_val,json=rEnumVal,enum=Example_Num" json:"r_enum_val,omitempty"`
	RChildVal   []*Child      `protobuf:"bytes,106,rep,name=r_child_val,json=rChildVal" json:"r_child_val,omitempty"`
}

func (m *Example) Reset()                    { *m = Example{} }
func (m *Example) String() string            { return proto.CompactTextString(m) }
func (*Example) ProtoMessage()               {}
func (*Example) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Example) GetUint64Val() uint64 {
	if m != nil {
		return m.Uint64Val
	}
	return 0
}

func (m *Example) GetStringVal() string {
	if m != nil {
		return m.StringVal
	}
	return ""
}

func (m *Example) GetFixed64Val() uint64 {
	if m != nil {
		return m.Fixed64Val
	}
	return 0
}

func (m *Example) GetFixed32Val() uint32 {
	if m != nil {
		return m.Fixed32Val
	}
	return 0
}

func (m *Example) GetEnumVal() Example_Num {
	if m != nil {
		return m.EnumVal
	}
	return Example_ZERO
}

func (m *Example) GetChildVal() *Child {
	if m != nil {
		return m.ChildVal
	}
	return nil
}

func (m *Example) GetRUint64Val() []uint64 {
	if m != nil {
		return m.RUint64Val
	}
	return nil
}

func (m *Example) GetRStringVal() []string {
	if m != nil {
		return m.RStringVal
	}
	return nil
}

func (m *Example) GetRFixed64Val() []uint64 {
	if m != nil {
		return m.RFixed64Val
	}
	return nil
}

func (m *Example) GetRFixed32Val() []uint32 {
	if m != nil {
		return m.RFixed32Val
	}
	return nil
}

func (m *Example) GetREnumVal() []Example_Num {
	if m != nil {
		return m.REnumVal
	}
	return nil
}

func (m *Example) GetRChildVal() []*Child {
	if m != nil {
		return m.RChildVal
	}
	return nil
}

type Child struct {
	V uint64 `protobuf:"varint,1,opt,name=v" json:"v,omitempty"`
}

func (m *Child) Reset()                    { *m = Child{} }
func (m *Child) String() string            { return proto.CompactTextString(m) }
func (*Child) ProtoMessage()               {}
func (*Child) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Child) GetV() uint64 {
	if m != nil {
		return m.V
	}
	return 0
}

func init() {
	proto.RegisterType((*Example)(nil), "Example")
	proto.RegisterType((*Child)(nil), "Child")
	proto.RegisterEnum("Example_Num", Example_Num_name, Example_Num_value)
}

func init() { proto.RegisterFile("internal/protosl.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0x47, 0xbf, 0xf9, 0x26, 0xcd, 0x9f, 0x9b, 0x2a, 0x25, 0xa0, 0x74, 0x23, 0x5e, 0x22, 0xe8,
	0xe0, 0x22, 0x42, 0x2a, 0xbe, 0x40, 0x89, 0xcb, 0x16, 0x46, 0x74, 0xe1, 0x26, 0x44, 0x9b, 0xb6,
	0x91, 0x24, 0x2d, 0xb7, 0x49, 0xe9, 0x6b, 0xfa, 0x46, 0x32, 0xd3, 0xb1, 0x09, 0xb8, 0x0a, 0xfc,
	0xce, 0x09, 0xcc, 0xe1, 0xc2, 0x65, 0x51, 0x37, 0x39, 0xd5, 0x59, 0xf9, 0xb0, 0xa5, 0x4d, 0xb3,
	0xd9, 0x95, 0x91, 0xfe, 0x86, 0xdf, 0x1c, 0x9c, 0xe4, 0x90, 0x55, 0xdb, 0x32, 0x0f, 0xae, 0x00,
	0xda, 0xa2, 0x6e, 0x9e, 0x1e, 0xd3, 0x7d, 0x56, 0x8e, 0x19, 0x32, 0x61, 0x49, 0xef, 0xb8, 0xbc,
	0x65, 0xa5, 0xc2, 0xbb, 0x86, 0x8a, 0x7a, 0xa5, 0xf1, 0x7f, 0x64, 0xc2, 0x93, 0xde, 0x71, 0x51,
	0xf8, 0x1a, 0xfc, 0x65, 0x71, 0xc8, 0x17, 0xe6, 0x77, 0x8e, 0x4c, 0xd8, 0x12, 0xcc, 0xd4, 0x17,
	0x26, 0xb1, 0x16, 0x2c, 0x64, 0xc2, 0x31, 0xc2, 0x24, 0x56, 0xc2, 0x1d, 0xb8, 0x79, 0xdd, 0x56,
	0x9a, 0x0e, 0x90, 0x89, 0xf3, 0x78, 0x18, 0x99, 0xb7, 0x45, 0xb3, 0xb6, 0x92, 0x8e, 0xa2, 0x4a,
	0xbc, 0x01, 0xef, 0x73, 0x5d, 0x94, 0x0b, 0x6d, 0xda, 0xc8, 0x84, 0x1f, 0xdb, 0xd1, 0x54, 0x2d,
	0xd2, 0xd5, 0x40, 0x49, 0x08, 0x43, 0x4a, 0x7b, 0x3d, 0x39, 0x72, 0x61, 0x49, 0xa0, 0xd7, 0x53,
	0x90, 0x36, 0x7a, 0x49, 0x4b, 0xe4, 0xc2, 0x93, 0x40, 0x2f, 0xa7, 0xa6, 0x10, 0xce, 0x28, 0xed,
	0x57, 0xad, 0x90, 0x0b, 0x5b, 0xfa, 0xf4, 0xdc, 0x65, 0x75, 0x8e, 0x09, 0x5b, 0x23, 0x17, 0xce,
	0xaf, 0x73, 0x2c, 0xbb, 0x07, 0xa0, 0xf4, 0xd4, 0x56, 0x20, 0xff, 0xd3, 0xe6, 0x52, 0x62, 0xe2,
	0x6e, 0xc1, 0xa7, 0xb4, 0xcb, 0xfb, 0x42, 0xde, 0xcb, 0xf3, 0x68, 0x6a, 0xfa, 0xc2, 0x31, 0xf0,
	0x59, 0x5b, 0x05, 0x2e, 0x58, 0xef, 0x89, 0x9c, 0x8f, 0xfe, 0x05, 0x0e, 0xf0, 0xf9, 0x2c, 0x19,
	0xb1, 0xf0, 0x02, 0x06, 0xda, 0x0a, 0x86, 0xc0, 0xf6, 0xe6, 0x8e, 0x6c, 0xff, 0x61, 0xeb, 0x8b,
	0x4f, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x1b, 0xe1, 0x7e, 0x0b, 0x02, 0x00, 0x00,
}
