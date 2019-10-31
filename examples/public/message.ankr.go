// Code generated by protoc-gen-ankr. DO NOT EDIT.
// source: message.proto

package pulibc_message

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

type RspCodeMsg struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RspCodeMsg) Reset()         { *m = RspCodeMsg{} }
func (m *RspCodeMsg) String() string { return proto.CompactTextString(m) }
func (*RspCodeMsg) ProtoMessage()    {}
func (*RspCodeMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *RspCodeMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RspCodeMsg.Unmarshal(m, b)
}
func (m *RspCodeMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RspCodeMsg.Marshal(b, m, deterministic)
}
func (m *RspCodeMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RspCodeMsg.Merge(m, src)
}
func (m *RspCodeMsg) XXX_Size() int {
	return xxx_messageInfo_RspCodeMsg.Size(m)
}
func (m *RspCodeMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_RspCodeMsg.DiscardUnknown(m)
}

var xxx_messageInfo_RspCodeMsg proto.InternalMessageInfo

func (m *RspCodeMsg) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *RspCodeMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*RspCodeMsg)(nil), "pulibc.message.RspCodeMsg")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 98 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2b, 0x28, 0xcd, 0xc9, 0x4c,
	0x4a, 0xd6, 0x83, 0x8a, 0x2a, 0x19, 0x71, 0x71, 0x05, 0x15, 0x17, 0x38, 0xe7, 0xa7, 0xa4, 0xfa,
	0x16, 0xa7, 0x0b, 0x09, 0x71, 0xb1, 0x24, 0xe7, 0xa7, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0x81, 0xd9, 0x42, 0x02, 0x5c, 0xcc, 0xb9, 0xc5, 0xe9, 0x12, 0x4c, 0x60, 0x21, 0x10, 0x33,
	0x89, 0x0d, 0x6c, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x60, 0x12, 0x30, 0x2e, 0x5b, 0x00,
	0x00, 0x00,
}
