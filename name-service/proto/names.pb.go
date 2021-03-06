// Code generated by protoc-gen-go. DO NOT EDIT.
// source: names.proto

package names

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

type Request struct {
	UserIds              []int64  `protobuf:"varint,1,rep,packed,name=userIds,proto3" json:"userIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4268625867c617c, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUserIds() []int64 {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type Response struct {
	Users                []*User  `protobuf:"bytes,101,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4268625867c617c, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type User struct {
	Id                   int64    `protobuf:"varint,201,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,202,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,203,opt,name=lastName,proto3" json:"lastName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4268625867c617c, []int{2}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*User)(nil), "User")
}

func init() {
	proto.RegisterFile("names.proto", fileDescriptor_f4268625867c617c)
}

var fileDescriptor_f4268625867c617c = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x4b, 0xcc, 0x4d,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe6, 0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x92, 0xe0, 0x62, 0x2f, 0x2d, 0x4e, 0x2d, 0xf2, 0x4c, 0x29, 0x96, 0x60, 0x54,
	0x60, 0xd6, 0x60, 0x0e, 0x82, 0x71, 0x95, 0xd4, 0xb9, 0x38, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3,
	0x8a, 0x53, 0x85, 0xa4, 0xb9, 0x58, 0x41, 0xc2, 0xc5, 0x12, 0xa9, 0x0a, 0xcc, 0x1a, 0xdc, 0x46,
	0xac, 0x7a, 0xa1, 0xc5, 0xa9, 0x45, 0x41, 0x10, 0x31, 0xa5, 0x60, 0x2e, 0x16, 0x10, 0x57, 0x88,
	0x9f, 0x8b, 0x29, 0x33, 0x45, 0xe2, 0x24, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x53, 0x66, 0x8a,
	0x90, 0x2c, 0x17, 0x67, 0x5a, 0x66, 0x51, 0x71, 0x89, 0x5f, 0x62, 0x6e, 0xaa, 0xc4, 0x29, 0x90,
	0x38, 0x67, 0x10, 0x42, 0x44, 0x48, 0x9a, 0x8b, 0x23, 0x27, 0x11, 0x2a, 0x7b, 0x1a, 0x22, 0x0b,
	0x17, 0x30, 0xd2, 0xe2, 0x62, 0x05, 0xd1, 0xc5, 0x42, 0x8a, 0x5c, 0x1c, 0xee, 0xa9, 0x25, 0x10,
	0x36, 0x87, 0x1e, 0xd4, 0xd9, 0x52, 0x9c, 0x7a, 0x30, 0xb7, 0x29, 0x31, 0x24, 0xb1, 0x81, 0x7d,
	0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x96, 0xd6, 0x6c, 0xe4, 0x00, 0x00, 0x00,
}
