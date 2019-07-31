// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/todo/v1beta1/todo.proto

package todov1beta1

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Todo is a note describing a task to be done.
type Todo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Done                 bool     `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_411429d5661e536a, []int{0}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Todo) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func init() {
	proto.RegisterType((*Todo)(nil), "todo.v1beta1.Todo")
}

func init() { proto.RegisterFile("api/proto/todo/v1beta1/todo.proto", fileDescriptor_411429d5661e536a) }

var fileDescriptor_411429d5661e536a = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0xc9, 0x4f, 0xc9, 0xd7, 0x2f, 0x33, 0x4c, 0x4a, 0x2d,
	0x49, 0x34, 0x04, 0x73, 0xf4, 0xc0, 0xe2, 0x42, 0x3c, 0x60, 0x36, 0x54, 0x42, 0xc9, 0x8e, 0x8b,
	0x25, 0x24, 0x3f, 0x25, 0x5f, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x33, 0x88, 0x29, 0x33, 0x45, 0x48, 0x88, 0x8b, 0xa5, 0x24, 0xb5, 0xa2, 0x44, 0x82, 0x09, 0x2c,
	0x02, 0x66, 0x83, 0xc4, 0x52, 0xf2, 0xf3, 0x52, 0x25, 0x98, 0x15, 0x18, 0x35, 0x38, 0x82, 0xc0,
	0x6c, 0xa7, 0x00, 0x2e, 0x81, 0xe4, 0xfc, 0x5c, 0x3d, 0x64, 0x33, 0x9d, 0x38, 0x41, 0x26, 0x06,
	0x80, 0x2c, 0x0b, 0x60, 0x8c, 0xe2, 0x06, 0x49, 0x41, 0x65, 0x16, 0x31, 0x31, 0x87, 0x44, 0x44,
	0xac, 0x62, 0xe2, 0x01, 0x29, 0xd0, 0x0b, 0x33, 0x74, 0x02, 0x09, 0x9e, 0x82, 0x70, 0x63, 0xa0,
	0xdc, 0x24, 0x36, 0xb0, 0x33, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x2a, 0xb3, 0x02,
	0xcb, 0x00, 0x00, 0x00,
}
