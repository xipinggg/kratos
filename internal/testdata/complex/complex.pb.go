// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: complex.proto

package complex

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Sex int32

const (
	Sex_man   Sex = 0
	Sex_woman Sex = 1
)

// Enum value maps for Sex.
var (
	Sex_name = map[int32]string{
		0: "man",
		1: "woman",
	}
	Sex_value = map[string]int32{
		"man":   0,
		"woman": 1,
	}
)

func (x Sex) Enum() *Sex {
	p := new(Sex)
	*p = x
	return p
}

func (x Sex) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sex) Descriptor() protoreflect.EnumDescriptor {
	return file_complex_proto_enumTypes[0].Descriptor()
}

func (Sex) Type() protoreflect.EnumType {
	return &file_complex_proto_enumTypes[0]
}

func (x Sex) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sex.Descriptor instead.
func (Sex) EnumDescriptor() ([]byte, []int) {
	return file_complex_proto_rawDescGZIP(), []int{0}
}

// SimpleMessage represents a simple message sent to the Echo service.
type Complex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id represents the message identifier.
	Id        int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NoOne     string                  `protobuf:"bytes,2,opt,name=no_one,json=numberOne,proto3" json:"no_one,omitempty"`
	Simple    *Simple                 `protobuf:"bytes,3,opt,name=simple,json=very_simple,proto3" json:"simple,omitempty"`
	Simples   []string                `protobuf:"bytes,4,rep,name=simples,proto3" json:"simples,omitempty"`
	B         bool                    `protobuf:"varint,5,opt,name=b,proto3" json:"b,omitempty"`
	Sex       Sex                     `protobuf:"varint,6,opt,name=sex,proto3,enum=testproto.Sex" json:"sex,omitempty"`
	Age       int32                   `protobuf:"varint,7,opt,name=age,proto3" json:"age,omitempty"`
	A         uint32                  `protobuf:"varint,8,opt,name=a,proto3" json:"a,omitempty"`
	Count     uint64                  `protobuf:"varint,9,opt,name=count,proto3" json:"count,omitempty"`
	Price     float32                 `protobuf:"fixed32,10,opt,name=price,proto3" json:"price,omitempty"`
	D         float64                 `protobuf:"fixed64,11,opt,name=d,proto3" json:"d,omitempty"`
	Byte      []byte                  `protobuf:"bytes,12,opt,name=byte,proto3" json:"byte,omitempty"`
	Timestamp *timestamppb.Timestamp  `protobuf:"bytes,13,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Duration  *durationpb.Duration    `protobuf:"bytes,14,opt,name=duration,proto3" json:"duration,omitempty"`
	Field     *fieldmaskpb.FieldMask  `protobuf:"bytes,15,opt,name=field,proto3" json:"field,omitempty"`
	Double    *wrapperspb.DoubleValue `protobuf:"bytes,16,opt,name=double,proto3" json:"double,omitempty"`
	Float     *wrapperspb.FloatValue  `protobuf:"bytes,17,opt,name=float,proto3" json:"float,omitempty"`
	Int64     *wrapperspb.Int64Value  `protobuf:"bytes,18,opt,name=int64,proto3" json:"int64,omitempty"`
	Int32     *wrapperspb.Int32Value  `protobuf:"bytes,19,opt,name=int32,proto3" json:"int32,omitempty"`
	Uint64    *wrapperspb.UInt64Value `protobuf:"bytes,20,opt,name=uint64,proto3" json:"uint64,omitempty"`
	Uint32    *wrapperspb.UInt32Value `protobuf:"bytes,21,opt,name=uint32,proto3" json:"uint32,omitempty"`
	Bool      *wrapperspb.BoolValue   `protobuf:"bytes,22,opt,name=bool,proto3" json:"bool,omitempty"`
	String_   *wrapperspb.StringValue `protobuf:"bytes,23,opt,name=string,proto3" json:"string,omitempty"`
	Bytes     *wrapperspb.BytesValue  `protobuf:"bytes,24,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Map       map[string]string       `protobuf:"bytes,25,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Complex) Reset() {
	*x = Complex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Complex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Complex) ProtoMessage() {}

func (x *Complex) ProtoReflect() protoreflect.Message {
	mi := &file_complex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Complex.ProtoReflect.Descriptor instead.
func (*Complex) Descriptor() ([]byte, []int) {
	return file_complex_proto_rawDescGZIP(), []int{0}
}

func (x *Complex) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Complex) GetNoOne() string {
	if x != nil {
		return x.NoOne
	}
	return ""
}

func (x *Complex) GetSimple() *Simple {
	if x != nil {
		return x.Simple
	}
	return nil
}

func (x *Complex) GetSimples() []string {
	if x != nil {
		return x.Simples
	}
	return nil
}

func (x *Complex) GetB() bool {
	if x != nil {
		return x.B
	}
	return false
}

func (x *Complex) GetSex() Sex {
	if x != nil {
		return x.Sex
	}
	return Sex_man
}

func (x *Complex) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Complex) GetA() uint32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Complex) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Complex) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Complex) GetD() float64 {
	if x != nil {
		return x.D
	}
	return 0
}

func (x *Complex) GetByte() []byte {
	if x != nil {
		return x.Byte
	}
	return nil
}

func (x *Complex) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Complex) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *Complex) GetField() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.Field
	}
	return nil
}

func (x *Complex) GetDouble() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Double
	}
	return nil
}

func (x *Complex) GetFloat() *wrapperspb.FloatValue {
	if x != nil {
		return x.Float
	}
	return nil
}

func (x *Complex) GetInt64() *wrapperspb.Int64Value {
	if x != nil {
		return x.Int64
	}
	return nil
}

func (x *Complex) GetInt32() *wrapperspb.Int32Value {
	if x != nil {
		return x.Int32
	}
	return nil
}

func (x *Complex) GetUint64() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Uint64
	}
	return nil
}

func (x *Complex) GetUint32() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Uint32
	}
	return nil
}

func (x *Complex) GetBool() *wrapperspb.BoolValue {
	if x != nil {
		return x.Bool
	}
	return nil
}

func (x *Complex) GetString_() *wrapperspb.StringValue {
	if x != nil {
		return x.String_
	}
	return nil
}

func (x *Complex) GetBytes() *wrapperspb.BytesValue {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Complex) GetMap() map[string]string {
	if x != nil {
		return x.Map
	}
	return nil
}

type Simple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Component string `protobuf:"bytes,1,opt,name=component,proto3" json:"component,omitempty"`
}

func (x *Simple) Reset() {
	*x = Simple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Simple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Simple) ProtoMessage() {}

func (x *Simple) ProtoReflect() protoreflect.Message {
	mi := &file_complex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Simple.ProtoReflect.Descriptor instead.
func (*Simple) Descriptor() ([]byte, []int) {
	return file_complex_proto_rawDescGZIP(), []int{1}
}

func (x *Simple) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

var File_complex_proto protoreflect.FileDescriptor

var file_complex_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x07,
	0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x06, 0x6e, 0x6f, 0x5f,
	0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x4f, 0x6e, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x0c,
	0x0a, 0x01, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x01, 0x62, 0x12, 0x20, 0x0a, 0x03,
	0x73, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x78, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x79, 0x74, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x79, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x34, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x31, 0x0a, 0x05, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x34,
	0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x75, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x12, 0x34, 0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x2e, 0x0a, 0x04, 0x62, 0x6f,
	0x6f, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x12, 0x31, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x19, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6d,
	0x61, 0x70, 0x1a, 0x36, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x26, 0x0a, 0x06, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x2a, 0x19, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x12, 0x07, 0x0a, 0x03, 0x6d, 0x61, 0x6e,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x77, 0x6f, 0x6d, 0x61, 0x6e, 0x10, 0x01, 0x42, 0x57, 0x5a,
	0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6d, 0x64,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x68,
	0x74, 0x74, 0x70, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x2f, 0x3b, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_complex_proto_rawDescOnce sync.Once
	file_complex_proto_rawDescData = file_complex_proto_rawDesc
)

func file_complex_proto_rawDescGZIP() []byte {
	file_complex_proto_rawDescOnce.Do(func() {
		file_complex_proto_rawDescData = protoimpl.X.CompressGZIP(file_complex_proto_rawDescData)
	})
	return file_complex_proto_rawDescData
}

var file_complex_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_complex_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_complex_proto_goTypes = []any{
	(Sex)(0),                       // 0: testproto.sex
	(*Complex)(nil),                // 1: testproto.Complex
	(*Simple)(nil),                 // 2: testproto.Simple
	nil,                            // 3: testproto.Complex.MapEntry
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),    // 5: google.protobuf.Duration
	(*fieldmaskpb.FieldMask)(nil),  // 6: google.protobuf.FieldMask
	(*wrapperspb.DoubleValue)(nil), // 7: google.protobuf.DoubleValue
	(*wrapperspb.FloatValue)(nil),  // 8: google.protobuf.FloatValue
	(*wrapperspb.Int64Value)(nil),  // 9: google.protobuf.Int64Value
	(*wrapperspb.Int32Value)(nil),  // 10: google.protobuf.Int32Value
	(*wrapperspb.UInt64Value)(nil), // 11: google.protobuf.UInt64Value
	(*wrapperspb.UInt32Value)(nil), // 12: google.protobuf.UInt32Value
	(*wrapperspb.BoolValue)(nil),   // 13: google.protobuf.BoolValue
	(*wrapperspb.StringValue)(nil), // 14: google.protobuf.StringValue
	(*wrapperspb.BytesValue)(nil),  // 15: google.protobuf.BytesValue
}
var file_complex_proto_depIdxs = []int32{
	2,  // 0: testproto.Complex.simple:type_name -> testproto.Simple
	0,  // 1: testproto.Complex.sex:type_name -> testproto.sex
	4,  // 2: testproto.Complex.timestamp:type_name -> google.protobuf.Timestamp
	5,  // 3: testproto.Complex.duration:type_name -> google.protobuf.Duration
	6,  // 4: testproto.Complex.field:type_name -> google.protobuf.FieldMask
	7,  // 5: testproto.Complex.double:type_name -> google.protobuf.DoubleValue
	8,  // 6: testproto.Complex.float:type_name -> google.protobuf.FloatValue
	9,  // 7: testproto.Complex.int64:type_name -> google.protobuf.Int64Value
	10, // 8: testproto.Complex.int32:type_name -> google.protobuf.Int32Value
	11, // 9: testproto.Complex.uint64:type_name -> google.protobuf.UInt64Value
	12, // 10: testproto.Complex.uint32:type_name -> google.protobuf.UInt32Value
	13, // 11: testproto.Complex.bool:type_name -> google.protobuf.BoolValue
	14, // 12: testproto.Complex.string:type_name -> google.protobuf.StringValue
	15, // 13: testproto.Complex.bytes:type_name -> google.protobuf.BytesValue
	3,  // 14: testproto.Complex.map:type_name -> testproto.Complex.MapEntry
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_complex_proto_init() }
func file_complex_proto_init() {
	if File_complex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_complex_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Complex); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_complex_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Simple); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_complex_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_complex_proto_goTypes,
		DependencyIndexes: file_complex_proto_depIdxs,
		EnumInfos:         file_complex_proto_enumTypes,
		MessageInfos:      file_complex_proto_msgTypes,
	}.Build()
	File_complex_proto = out.File
	file_complex_proto_rawDesc = nil
	file_complex_proto_goTypes = nil
	file_complex_proto_depIdxs = nil
}
