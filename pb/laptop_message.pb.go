// Code generated by protoc-gen-go. DO NOT EDIT.
// source: laptop_message.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Laptop struct {
	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand    string     `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Name     string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Cpu      *CPU       `protobuf:"bytes,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Ram      *Memory    `protobuf:"bytes,5,opt,name=ram,proto3" json:"ram,omitempty"`
	Gpus     []*GPU     `protobuf:"bytes,6,rep,name=gpus,proto3" json:"gpus,omitempty"`
	Storages []*Storage `protobuf:"bytes,7,rep,name=storages,proto3" json:"storages,omitempty"`
	Screen   *Screen    `protobuf:"bytes,8,opt,name=screen,proto3" json:"screen,omitempty"`
	Keyboard *Keyboard  `protobuf:"bytes,9,opt,name=keyboard,proto3" json:"keyboard,omitempty"`
	// Types that are valid to be assigned to Weight:
	//	*Laptop_WeightKg
	//	*Laptop_WeightLb
	Weight               isLaptop_Weight      `protobuf_oneof:"weight"`
	PriceUsd             float64              `protobuf:"fixed64,12,opt,name=price_usd,json=priceUsd,proto3" json:"price_usd,omitempty"`
	ReleaseYear          uint32               `protobuf:"varint,13,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Laptop) Reset()         { *m = Laptop{} }
func (m *Laptop) String() string { return proto.CompactTextString(m) }
func (*Laptop) ProtoMessage()    {}
func (*Laptop) Descriptor() ([]byte, []int) {
	return fileDescriptor_07a3824d46f4b731, []int{0}
}

func (m *Laptop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Laptop.Unmarshal(m, b)
}
func (m *Laptop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Laptop.Marshal(b, m, deterministic)
}
func (m *Laptop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Laptop.Merge(m, src)
}
func (m *Laptop) XXX_Size() int {
	return xxx_messageInfo_Laptop.Size(m)
}
func (m *Laptop) XXX_DiscardUnknown() {
	xxx_messageInfo_Laptop.DiscardUnknown(m)
}

var xxx_messageInfo_Laptop proto.InternalMessageInfo

func (m *Laptop) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Laptop) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Laptop) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Laptop) GetCpu() *CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Laptop) GetRam() *Memory {
	if m != nil {
		return m.Ram
	}
	return nil
}

func (m *Laptop) GetGpus() []*GPU {
	if m != nil {
		return m.Gpus
	}
	return nil
}

func (m *Laptop) GetStorages() []*Storage {
	if m != nil {
		return m.Storages
	}
	return nil
}

func (m *Laptop) GetScreen() *Screen {
	if m != nil {
		return m.Screen
	}
	return nil
}

func (m *Laptop) GetKeyboard() *Keyboard {
	if m != nil {
		return m.Keyboard
	}
	return nil
}

type isLaptop_Weight interface {
	isLaptop_Weight()
}

type Laptop_WeightKg struct {
	WeightKg float64 `protobuf:"fixed64,10,opt,name=weight_kg,json=weightKg,proto3,oneof"`
}

type Laptop_WeightLb struct {
	WeightLb float64 `protobuf:"fixed64,11,opt,name=weight_lb,json=weightLb,proto3,oneof"`
}

func (*Laptop_WeightKg) isLaptop_Weight() {}

func (*Laptop_WeightLb) isLaptop_Weight() {}

func (m *Laptop) GetWeight() isLaptop_Weight {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *Laptop) GetWeightKg() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightKg); ok {
		return x.WeightKg
	}
	return 0
}

func (m *Laptop) GetWeightLb() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightLb); ok {
		return x.WeightLb
	}
	return 0
}

func (m *Laptop) GetPriceUsd() float64 {
	if m != nil {
		return m.PriceUsd
	}
	return 0
}

func (m *Laptop) GetReleaseYear() uint32 {
	if m != nil {
		return m.ReleaseYear
	}
	return 0
}

func (m *Laptop) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Laptop) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Laptop_WeightKg)(nil),
		(*Laptop_WeightLb)(nil),
	}
}

func init() {
	proto.RegisterType((*Laptop)(nil), "lab.pcbook.Laptop")
}

func init() {
	proto.RegisterFile("laptop_message.proto", fileDescriptor_07a3824d46f4b731)
}

var fileDescriptor_07a3824d46f4b731 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x49, 0xd3, 0x86, 0xe4, 0x75, 0x1b, 0x92, 0x29, 0x60, 0x15, 0x21, 0x32, 0xe0, 0x50,
	0x71, 0x48, 0x11, 0x9c, 0x38, 0x32, 0x0e, 0x20, 0x6d, 0x48, 0x93, 0x47, 0x0f, 0x70, 0x89, 0xec,
	0xe4, 0x61, 0xa2, 0x26, 0xb5, 0x65, 0x3b, 0x42, 0xbd, 0xf2, 0x97, 0xa3, 0x3a, 0xe9, 0xc8, 0xb2,
	0x5b, 0xde, 0xf7, 0xf3, 0x51, 0xfc, 0x7e, 0xc0, 0xa2, 0xe6, 0xda, 0x29, 0x9d, 0x37, 0x68, 0x2d,
	0x97, 0x98, 0x69, 0xa3, 0x9c, 0x22, 0x50, 0x73, 0x91, 0xe9, 0x42, 0x28, 0xb5, 0x5d, 0x3e, 0xd3,
	0x46, 0x15, 0x68, 0xad, 0x32, 0x77, 0xa5, 0xe5, 0xa2, 0xc1, 0x46, 0x99, 0xfd, 0x28, 0x7d, 0x62,
	0x9d, 0x32, 0x5c, 0xe2, 0x58, 0xb6, 0x85, 0x41, 0xdc, 0x8d, 0xd2, 0xa7, 0x5b, 0xdc, 0x0b, 0xc5,
	0x4d, 0x39, 0xca, 0x5f, 0x4a, 0xa5, 0x64, 0x8d, 0x6b, 0x5f, 0x89, 0xf6, 0xd7, 0xda, 0x55, 0x0d,
	0x5a, 0xc7, 0x1b, 0xdd, 0x09, 0xaf, 0xfe, 0x4e, 0x21, 0xba, 0xf2, 0x9d, 0x93, 0x33, 0x98, 0x54,
	0x25, 0x0d, 0xd2, 0x60, 0x95, 0xb0, 0x49, 0x55, 0x92, 0x05, 0xcc, 0x84, 0xe1, 0xbb, 0x92, 0x4e,
	0x7c, 0xd4, 0x15, 0x84, 0xc0, 0x74, 0xc7, 0x1b, 0xa4, 0xa1, 0x0f, 0xfd, 0x37, 0x39, 0x87, 0xb0,
	0xd0, 0x2d, 0x9d, 0xa6, 0xc1, 0x6a, 0xfe, 0xfe, 0x51, 0xf6, 0x7f, 0xe6, 0xec, 0xf3, 0xf5, 0x86,
	0x1d, 0x18, 0x79, 0x03, 0xa1, 0xe1, 0x0d, 0x9d, 0x79, 0x85, 0x0c, 0x95, 0x6f, 0x7e, 0x78, 0x76,
	0xc0, 0xe4, 0x35, 0x4c, 0xa5, 0x6e, 0x2d, 0x8d, 0xd2, 0x70, 0xfc, 0xa7, 0x2f, 0xd7, 0x1b, 0xe6,
	0x21, 0x59, 0x43, 0xdc, 0xaf, 0xc6, 0xd2, 0x87, 0x5e, 0x7c, 0x3c, 0x14, 0x6f, 0x3a, 0xc6, 0x6e,
	0x25, 0xf2, 0x16, 0xa2, 0x6e, 0x69, 0x34, 0xbe, 0xff, 0xfc, 0x8d, 0x27, 0xac, 0x37, 0xc8, 0x3b,
	0x88, 0x8f, 0xab, 0xa4, 0x89, 0xb7, 0x17, 0x43, 0xfb, 0xb2, 0x67, 0xec, 0xd6, 0x22, 0x2f, 0x20,
	0xf9, 0x83, 0x95, 0xfc, 0xed, 0xf2, 0xad, 0xa4, 0x90, 0x06, 0xab, 0xe0, 0xeb, 0x03, 0x16, 0x77,
	0xd1, 0xa5, 0x1c, 0xe0, 0x5a, 0xd0, 0xf9, 0x5d, 0x7c, 0x25, 0xc8, 0x73, 0x48, 0xb4, 0xa9, 0x0a,
	0xcc, 0x5b, 0x5b, 0xd2, 0x93, 0x03, 0x66, 0xb1, 0x0f, 0x36, 0xb6, 0x24, 0xe7, 0x70, 0x62, 0xb0,
	0x46, 0x6e, 0x31, 0xdf, 0x23, 0x37, 0xf4, 0x34, 0x0d, 0x56, 0xa7, 0x6c, 0xde, 0x67, 0x3f, 0x90,
	0x1b, 0xf2, 0x11, 0xa0, 0xd5, 0x25, 0x77, 0x58, 0xe6, 0xdc, 0xd1, 0x33, 0xdf, 0xf1, 0x32, 0xeb,
	0xae, 0x9e, 0x1d, 0xaf, 0x9e, 0x7d, 0x3f, 0x5e, 0x9d, 0x25, 0xbd, 0xfd, 0xc9, 0x5d, 0xc4, 0x10,
	0x75, 0x6d, 0x5c, 0xcc, 0x7e, 0x86, 0x6b, 0x2d, 0x44, 0xe4, 0xfd, 0x0f, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x47, 0x0d, 0xf2, 0x15, 0xcb, 0x02, 0x00, 0x00,
}
