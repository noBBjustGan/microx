// Code generated by protoc-gen-go. DO NOT EDIT.
// source: microx/common/proto/topic/user.proto

package microx_topic

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

type UserInfo struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Mobile               string   `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1973e2cca5f4aea0, []int{0}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserInfo) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

type UserCreated struct {
	Id                   int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic                string    `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Info                 *UserInfo `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserCreated) Reset()         { *m = UserCreated{} }
func (m *UserCreated) String() string { return proto.CompactTextString(m) }
func (*UserCreated) ProtoMessage()    {}
func (*UserCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_1973e2cca5f4aea0, []int{1}
}

func (m *UserCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCreated.Unmarshal(m, b)
}
func (m *UserCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCreated.Marshal(b, m, deterministic)
}
func (m *UserCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCreated.Merge(m, src)
}
func (m *UserCreated) XXX_Size() int {
	return xxx_messageInfo_UserCreated.Size(m)
}
func (m *UserCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCreated.DiscardUnknown(m)
}

var xxx_messageInfo_UserCreated proto.InternalMessageInfo

func (m *UserCreated) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserCreated) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *UserCreated) GetInfo() *UserInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "microx.topic.UserInfo")
	proto.RegisterType((*UserCreated)(nil), "microx.topic.UserCreated")
}

func init() {
	proto.RegisterFile("microx/common/proto/topic/user.proto", fileDescriptor_1973e2cca5f4aea0)
}

var fileDescriptor_1973e2cca5f4aea0 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xc1, 0x0b, 0x82, 0x30,
	0x14, 0x87, 0x99, 0x96, 0xd5, 0x33, 0x3a, 0x8c, 0x30, 0x8f, 0x22, 0x1d, 0xa4, 0xc3, 0x06, 0x75,
	0xec, 0xd8, 0xc9, 0xab, 0xd0, 0x59, 0xd2, 0x4d, 0x78, 0xd0, 0x7c, 0x32, 0x0d, 0xfa, 0xf3, 0xc3,
	0x69, 0xd1, 0xf1, 0xb7, 0x7d, 0x1f, 0xef, 0x83, 0xa3, 0xc1, 0xda, 0xd2, 0x5b, 0xd6, 0x64, 0x0c,
	0xb5, 0xb2, 0xb3, 0x34, 0x90, 0x1c, 0xa8, 0xc3, 0x5a, 0xbe, 0x7a, 0x6d, 0x85, 0x7b, 0xe0, 0xdb,
	0x89, 0x12, 0xee, 0x23, 0xbd, 0xc2, 0xfa, 0xde, 0x6b, 0x9b, 0xb7, 0x0d, 0xf1, 0x03, 0xac, 0x46,
	0xae, 0x44, 0x15, 0xb3, 0x84, 0x65, 0x7e, 0x11, 0x8c, 0x33, 0x57, 0x3c, 0x82, 0xc0, 0x50, 0x85,
	0x4f, 0x1d, 0x7b, 0x09, 0xcb, 0x36, 0xc5, 0xbc, 0xd2, 0x12, 0xc2, 0x51, 0xbe, 0x59, 0xfd, 0x18,
	0xb4, 0xe2, 0x3b, 0xf0, 0x7e, 0xaa, 0x87, 0x8a, 0xef, 0x61, 0xe9, 0x8e, 0xcc, 0xd6, 0x34, 0xf8,
	0x09, 0x16, 0xd8, 0x36, 0x14, 0xfb, 0x09, 0xcb, 0xc2, 0x73, 0x24, 0xfe, 0x73, 0xc4, 0xb7, 0xa5,
	0x70, 0x4c, 0x15, 0xb8, 0xe4, 0xcb, 0x27, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x8a, 0xa2, 0x64, 0xda,
	0x00, 0x00, 0x00,
}
