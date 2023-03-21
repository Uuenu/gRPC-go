// Code generated by protoc-gen-go. DO NOT EDIT.
// source: laptop_service.proto

package pb

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

type CreateLaptopRequest struct {
	Laptop               *Laptop  `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLaptopRequest) Reset()         { *m = CreateLaptopRequest{} }
func (m *CreateLaptopRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLaptopRequest) ProtoMessage()    {}
func (*CreateLaptopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{0}
}

func (m *CreateLaptopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLaptopRequest.Unmarshal(m, b)
}
func (m *CreateLaptopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLaptopRequest.Marshal(b, m, deterministic)
}
func (m *CreateLaptopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLaptopRequest.Merge(m, src)
}
func (m *CreateLaptopRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLaptopRequest.Size(m)
}
func (m *CreateLaptopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLaptopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLaptopRequest proto.InternalMessageInfo

func (m *CreateLaptopRequest) GetLaptop() *Laptop {
	if m != nil {
		return m.Laptop
	}
	return nil
}

type CreateLaptopResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLaptopResponse) Reset()         { *m = CreateLaptopResponse{} }
func (m *CreateLaptopResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLaptopResponse) ProtoMessage()    {}
func (*CreateLaptopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{1}
}

func (m *CreateLaptopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLaptopResponse.Unmarshal(m, b)
}
func (m *CreateLaptopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLaptopResponse.Marshal(b, m, deterministic)
}
func (m *CreateLaptopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLaptopResponse.Merge(m, src)
}
func (m *CreateLaptopResponse) XXX_Size() int {
	return xxx_messageInfo_CreateLaptopResponse.Size(m)
}
func (m *CreateLaptopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLaptopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLaptopResponse proto.InternalMessageInfo

func (m *CreateLaptopResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateLaptopRequest)(nil), "lab.pcbook.CreateLaptopRequest")
	proto.RegisterType((*CreateLaptopResponse)(nil), "lab.pcbook.CreateLaptopResponse")
}

func init() {
	proto.RegisterFile("laptop_service.proto", fileDescriptor_240c60d9fb227e71)
}

var fileDescriptor_240c60d9fb227e71 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x49, 0x2c, 0x28,
	0xc9, 0x2f, 0x88, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0xca, 0x49, 0x4c, 0xd2, 0x2b, 0x48, 0x4e, 0xca, 0xcf, 0xcf, 0x96, 0x82, 0xa9, 0xc8,
	0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x87, 0xaa, 0x50, 0x72, 0xe4, 0x12, 0x76, 0x2e, 0x4a, 0x4d, 0x2c,
	0x49, 0xf5, 0x01, 0xcb, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x69, 0x71, 0xb1, 0x41,
	0x94, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x09, 0xe9, 0x21, 0x4c, 0xd2, 0x83, 0x2a, 0x85,
	0xaa, 0x50, 0x52, 0xe3, 0x12, 0x41, 0x35, 0xa2, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x88, 0x8f,
	0x8b, 0x29, 0x33, 0x05, 0xac, 0x9f, 0x33, 0x88, 0x29, 0x33, 0xc5, 0x28, 0x85, 0x8b, 0x17, 0xa2,
	0x22, 0x18, 0xe2, 0x46, 0xa1, 0x60, 0x2e, 0x1e, 0x64, 0x8d, 0x42, 0xf2, 0xc8, 0x96, 0x60, 0x71,
	0x95, 0x94, 0x02, 0x6e, 0x05, 0x10, 0x3b, 0x95, 0x18, 0x9c, 0x58, 0xa3, 0x98, 0xf5, 0x0b, 0x92,
	0x92, 0xd8, 0xc0, 0xde, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x67, 0xe4, 0xe0, 0xd0, 0x18,
	0x01, 0x00, 0x00,
}