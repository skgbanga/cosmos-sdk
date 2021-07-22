// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: testdata.proto

package testdata

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Dog struct {
	Size_ string `protobuf:"bytes,1,opt,name=size,proto3" json:"size,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Dog) Reset()         { *m = Dog{} }
func (m *Dog) String() string { return proto.CompactTextString(m) }
func (*Dog) ProtoMessage()    {}
func (*Dog) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c4782d007dfce9, []int{0}
}
func (m *Dog) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Dog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Dog.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Dog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dog.Merge(m, src)
}
func (m *Dog) XXX_Size() int {
	return m.Size()
}
func (m *Dog) XXX_DiscardUnknown() {
	xxx_messageInfo_Dog.DiscardUnknown(m)
}

var xxx_messageInfo_Dog proto.InternalMessageInfo

func (m *Dog) GetSize_() string {
	if m != nil {
		return m.Size_
	}
	return ""
}

func (m *Dog) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Cat struct {
	Moniker string `protobuf:"bytes,1,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Lives   int32  `protobuf:"varint,2,opt,name=lives,proto3" json:"lives,omitempty"`
}

func (m *Cat) Reset()         { *m = Cat{} }
func (m *Cat) String() string { return proto.CompactTextString(m) }
func (*Cat) ProtoMessage()    {}
func (*Cat) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c4782d007dfce9, []int{1}
}
func (m *Cat) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Cat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Cat.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Cat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cat.Merge(m, src)
}
func (m *Cat) XXX_Size() int {
	return m.Size()
}
func (m *Cat) XXX_DiscardUnknown() {
	xxx_messageInfo_Cat.DiscardUnknown(m)
}

var xxx_messageInfo_Cat proto.InternalMessageInfo

func (m *Cat) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *Cat) GetLives() int32 {
	if m != nil {
		return m.Lives
	}
	return 0
}

type HasAnimal struct {
	Animal *types.Any `protobuf:"bytes,1,opt,name=animal,proto3" json:"animal,omitempty"`
	X      int64      `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
}

func (m *HasAnimal) Reset()         { *m = HasAnimal{} }
func (m *HasAnimal) String() string { return proto.CompactTextString(m) }
func (*HasAnimal) ProtoMessage()    {}
func (*HasAnimal) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c4782d007dfce9, []int{2}
}
func (m *HasAnimal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HasAnimal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HasAnimal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HasAnimal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasAnimal.Merge(m, src)
}
func (m *HasAnimal) XXX_Size() int {
	return m.Size()
}
func (m *HasAnimal) XXX_DiscardUnknown() {
	xxx_messageInfo_HasAnimal.DiscardUnknown(m)
}

var xxx_messageInfo_HasAnimal proto.InternalMessageInfo

func (m *HasAnimal) GetAnimal() *types.Any {
	if m != nil {
		return m.Animal
	}
	return nil
}

func (m *HasAnimal) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

type HasHasAnimal struct {
	HasAnimal *types.Any `protobuf:"bytes,1,opt,name=has_animal,json=hasAnimal,proto3" json:"has_animal,omitempty"`
}

func (m *HasHasAnimal) Reset()         { *m = HasHasAnimal{} }
func (m *HasHasAnimal) String() string { return proto.CompactTextString(m) }
func (*HasHasAnimal) ProtoMessage()    {}
func (*HasHasAnimal) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c4782d007dfce9, []int{3}
}
func (m *HasHasAnimal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HasHasAnimal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HasHasAnimal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HasHasAnimal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasHasAnimal.Merge(m, src)
}
func (m *HasHasAnimal) XXX_Size() int {
	return m.Size()
}
func (m *HasHasAnimal) XXX_DiscardUnknown() {
	xxx_messageInfo_HasHasAnimal.DiscardUnknown(m)
}

var xxx_messageInfo_HasHasAnimal proto.InternalMessageInfo

func (m *HasHasAnimal) GetHasAnimal() *types.Any {
	if m != nil {
		return m.HasAnimal
	}
	return nil
}

type HasHasHasAnimal struct {
	HasHasAnimal *types.Any `protobuf:"bytes,1,opt,name=has_has_animal,json=hasHasAnimal,proto3" json:"has_has_animal,omitempty"`
}

func (m *HasHasHasAnimal) Reset()         { *m = HasHasHasAnimal{} }
func (m *HasHasHasAnimal) String() string { return proto.CompactTextString(m) }
func (*HasHasHasAnimal) ProtoMessage()    {}
func (*HasHasHasAnimal) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c4782d007dfce9, []int{4}
}
func (m *HasHasHasAnimal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HasHasHasAnimal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HasHasHasAnimal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HasHasHasAnimal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasHasHasAnimal.Merge(m, src)
}
func (m *HasHasHasAnimal) XXX_Size() int {
	return m.Size()
}
func (m *HasHasHasAnimal) XXX_DiscardUnknown() {
	xxx_messageInfo_HasHasHasAnimal.DiscardUnknown(m)
}

var xxx_messageInfo_HasHasHasAnimal proto.InternalMessageInfo

func (m *HasHasHasAnimal) GetHasHasAnimal() *types.Any {
	if m != nil {
		return m.HasHasAnimal
	}
	return nil
}

// bad MultiSignature with extra fields
type BadMultiSignature struct {
	Signatures       [][]byte `protobuf:"bytes,1,rep,name=signatures,proto3" json:"signatures,omitempty"`
	MaliciousField   []byte   `protobuf:"bytes,5,opt,name=malicious_field,json=maliciousField,proto3" json:"malicious_field,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BadMultiSignature) Reset()         { *m = BadMultiSignature{} }
func (m *BadMultiSignature) String() string { return proto.CompactTextString(m) }
func (*BadMultiSignature) ProtoMessage()    {}
func (*BadMultiSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c4782d007dfce9, []int{5}
}
func (m *BadMultiSignature) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BadMultiSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BadMultiSignature.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BadMultiSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BadMultiSignature.Merge(m, src)
}
func (m *BadMultiSignature) XXX_Size() int {
	return m.Size()
}
func (m *BadMultiSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_BadMultiSignature.DiscardUnknown(m)
}

var xxx_messageInfo_BadMultiSignature proto.InternalMessageInfo

func (m *BadMultiSignature) GetSignatures() [][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *BadMultiSignature) GetMaliciousField() []byte {
	if m != nil {
		return m.MaliciousField
	}
	return nil
}

type TableModel struct {
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *TableModel) Reset()         { *m = TableModel{} }
func (m *TableModel) String() string { return proto.CompactTextString(m) }
func (*TableModel) ProtoMessage()    {}
func (*TableModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c4782d007dfce9, []int{6}
}
func (m *TableModel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TableModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TableModel.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TableModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableModel.Merge(m, src)
}
func (m *TableModel) XXX_Size() int {
	return m.Size()
}
func (m *TableModel) XXX_DiscardUnknown() {
	xxx_messageInfo_TableModel.DiscardUnknown(m)
}

var xxx_messageInfo_TableModel proto.InternalMessageInfo

func (m *TableModel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TableModel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Dog)(nil), "testdata.Dog")
	proto.RegisterType((*Cat)(nil), "testdata.Cat")
	proto.RegisterType((*HasAnimal)(nil), "testdata.HasAnimal")
	proto.RegisterType((*HasHasAnimal)(nil), "testdata.HasHasAnimal")
	proto.RegisterType((*HasHasHasAnimal)(nil), "testdata.HasHasHasAnimal")
	proto.RegisterType((*BadMultiSignature)(nil), "testdata.BadMultiSignature")
	proto.RegisterType((*TableModel)(nil), "testdata.TableModel")
}

func init() { proto.RegisterFile("testdata.proto", fileDescriptor_40c4782d007dfce9) }

var fileDescriptor_40c4782d007dfce9 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0xae, 0xda, 0x30,
	0x18, 0x85, 0x31, 0x01, 0x5a, 0xfe, 0x46, 0x41, 0xb5, 0x18, 0x52, 0x86, 0x14, 0x65, 0x29, 0x43,
	0x49, 0xaa, 0xa2, 0x2e, 0x6c, 0x40, 0xd5, 0xb2, 0xb0, 0xa4, 0x9d, 0xba, 0x20, 0x87, 0x98, 0xc4,
	0xc2, 0x89, 0x2b, 0xec, 0x54, 0xd0, 0xa7, 0xe8, 0x2b, 0xf4, 0x6d, 0x3a, 0x32, 0x76, 0xbc, 0x82,
	0x17, 0xb9, 0x8a, 0x93, 0x5c, 0x18, 0xee, 0xc0, 0x94, 0x73, 0x8e, 0x73, 0x3e, 0xfd, 0x89, 0x7f,
	0xb0, 0x14, 0x95, 0x2a, 0x22, 0x8a, 0x78, 0x3f, 0xf7, 0x42, 0x09, 0xfc, 0xb2, 0xf6, 0x83, 0x7e,
	0x2c, 0x62, 0xa1, 0x43, 0xbf, 0x50, 0xe5, 0xf9, 0xe0, 0x4d, 0x2c, 0x44, 0xcc, 0xa9, 0xaf, 0x5d,
	0x98, 0x6f, 0x7d, 0x92, 0x1d, 0xcb, 0x23, 0x77, 0x0c, 0xc6, 0x67, 0x11, 0x63, 0x0c, 0x2d, 0xc9,
	0x7e, 0x53, 0x1b, 0x0d, 0xd1, 0xa8, 0x1b, 0x68, 0x5d, 0x64, 0x19, 0x49, 0xa9, 0xdd, 0x2c, 0xb3,
	0x42, 0xbb, 0x9f, 0xc0, 0x58, 0x10, 0x85, 0x6d, 0x78, 0x91, 0x8a, 0x8c, 0xed, 0xe8, 0xbe, 0x6a,
	0xd4, 0x16, 0xf7, 0xa1, 0xcd, 0xd9, 0x2f, 0x2a, 0x75, 0xab, 0x1d, 0x94, 0xc6, 0xfd, 0x0a, 0xdd,
	0x25, 0x91, 0xb3, 0x8c, 0xa5, 0x84, 0xe3, 0xf7, 0xd0, 0x21, 0x5a, 0xe9, 0xee, 0xab, 0x8f, 0x7d,
	0xaf, 0x1c, 0xcf, 0xab, 0xc7, 0xf3, 0x66, 0xd9, 0x31, 0xa8, 0xde, 0xc1, 0x26, 0xa0, 0x83, 0x86,
	0x19, 0x01, 0x3a, 0xb8, 0x0b, 0x30, 0x97, 0x44, 0x5e, 0x59, 0x13, 0x80, 0x84, 0xc8, 0xf5, 0x1d,
	0xbc, 0x6e, 0x52, 0x97, 0xdc, 0x15, 0xf4, 0x4a, 0xc8, 0x95, 0x33, 0x05, 0xab, 0xe0, 0xdc, 0xc9,
	0x32, 0x93, 0x9b, 0xae, 0x1b, 0xc2, 0xeb, 0x39, 0x89, 0x56, 0x39, 0x57, 0xec, 0x1b, 0x8b, 0x33,
	0xa2, 0xf2, 0x3d, 0xc5, 0x0e, 0x80, 0xac, 0x8d, 0xb4, 0xd1, 0xd0, 0x18, 0x99, 0xc1, 0x4d, 0x82,
	0xdf, 0x41, 0x2f, 0x25, 0x9c, 0x6d, 0x98, 0xc8, 0xe5, 0x7a, 0xcb, 0x28, 0x8f, 0xec, 0xf6, 0x10,
	0x8d, 0xcc, 0xc0, 0x7a, 0x8a, 0xbf, 0x14, 0xe9, 0xb4, 0x75, 0xfa, 0xfb, 0x16, 0xb9, 0x1f, 0x00,
	0xbe, 0x93, 0x90, 0xd3, 0x95, 0x88, 0x28, 0xc7, 0x16, 0x34, 0x59, 0x54, 0xfd, 0xf9, 0x26, 0x8b,
	0x9e, 0xbb, 0xa9, 0xf9, 0xf2, 0xdf, 0xd9, 0x41, 0xa7, 0xb3, 0x83, 0x1e, 0xce, 0x0e, 0xfa, 0x73,
	0x71, 0x1a, 0xa7, 0x8b, 0xd3, 0xf8, 0x7f, 0x71, 0x1a, 0x3f, 0xbc, 0x98, 0xa9, 0x24, 0x0f, 0xbd,
	0x8d, 0x48, 0xfd, 0x8d, 0x90, 0xa9, 0x90, 0xd5, 0x63, 0x2c, 0xa3, 0x9d, 0x5f, 0xac, 0x52, 0xae,
	0x18, 0xf7, 0xeb, 0x9d, 0x0a, 0x3b, 0xfa, 0xdb, 0x27, 0x8f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfa,
	0xe5, 0x40, 0xb7, 0x76, 0x02, 0x00, 0x00,
}

func (m *Dog) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Dog) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Dog) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTestdata(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Size_) > 0 {
		i -= len(m.Size_)
		copy(dAtA[i:], m.Size_)
		i = encodeVarintTestdata(dAtA, i, uint64(len(m.Size_)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Cat) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cat) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Cat) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Lives != 0 {
		i = encodeVarintTestdata(dAtA, i, uint64(m.Lives))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintTestdata(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HasAnimal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HasAnimal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HasAnimal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.X != 0 {
		i = encodeVarintTestdata(dAtA, i, uint64(m.X))
		i--
		dAtA[i] = 0x10
	}
	if m.Animal != nil {
		{
			size, err := m.Animal.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTestdata(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HasHasAnimal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HasHasAnimal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HasHasAnimal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HasAnimal != nil {
		{
			size, err := m.HasAnimal.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTestdata(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HasHasHasAnimal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HasHasHasAnimal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HasHasHasAnimal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HasHasAnimal != nil {
		{
			size, err := m.HasHasAnimal.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTestdata(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BadMultiSignature) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BadMultiSignature) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BadMultiSignature) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.MaliciousField) > 0 {
		i -= len(m.MaliciousField)
		copy(dAtA[i:], m.MaliciousField)
		i = encodeVarintTestdata(dAtA, i, uint64(len(m.MaliciousField)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Signatures) > 0 {
		for iNdEx := len(m.Signatures) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signatures[iNdEx])
			copy(dAtA[i:], m.Signatures[iNdEx])
			i = encodeVarintTestdata(dAtA, i, uint64(len(m.Signatures[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TableModel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TableModel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TableModel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTestdata(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTestdata(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTestdata(dAtA []byte, offset int, v uint64) int {
	offset -= sovTestdata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Dog) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Size_)
	if l > 0 {
		n += 1 + l + sovTestdata(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTestdata(uint64(l))
	}
	return n
}

func (m *Cat) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovTestdata(uint64(l))
	}
	if m.Lives != 0 {
		n += 1 + sovTestdata(uint64(m.Lives))
	}
	return n
}

func (m *HasAnimal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Animal != nil {
		l = m.Animal.Size()
		n += 1 + l + sovTestdata(uint64(l))
	}
	if m.X != 0 {
		n += 1 + sovTestdata(uint64(m.X))
	}
	return n
}

func (m *HasHasAnimal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HasAnimal != nil {
		l = m.HasAnimal.Size()
		n += 1 + l + sovTestdata(uint64(l))
	}
	return n
}

func (m *HasHasHasAnimal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HasHasAnimal != nil {
		l = m.HasHasAnimal.Size()
		n += 1 + l + sovTestdata(uint64(l))
	}
	return n
}

func (m *BadMultiSignature) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Signatures) > 0 {
		for _, b := range m.Signatures {
			l = len(b)
			n += 1 + l + sovTestdata(uint64(l))
		}
	}
	l = len(m.MaliciousField)
	if l > 0 {
		n += 1 + l + sovTestdata(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TableModel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTestdata(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTestdata(uint64(l))
	}
	return n
}

func sovTestdata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTestdata(x uint64) (n int) {
	return sovTestdata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Dog) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestdata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Dog: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Dog: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestdata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTestdata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTestdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Size_ = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestdata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTestdata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTestdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTestdata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTestdata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Cat) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestdata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Cat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestdata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTestdata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTestdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lives", wireType)
			}
			m.Lives = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestdata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lives |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTestdata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTestdata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HasAnimal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestdata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HasAnimal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HasAnimal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Animal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestdata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTestdata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTestdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Animal == nil {
				m.Animal = &types.Any{}
			}
			if err := m.Animal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestdata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.X |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTestdata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTestdata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HasHasAnimal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestdata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HasHasAnimal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HasHasAnimal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasAnimal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestdata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTestdata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTestdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HasAnimal == nil {
				m.HasAnimal = &types.Any{}
			}
			if err := m.HasAnimal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTestdata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTestdata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HasHasHasAnimal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestdata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HasHasHasAnimal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HasHasHasAnimal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasHasAnimal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestdata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTestdata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTestdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HasHasAnimal == nil {
				m.HasHasAnimal = &types.Any{}
			}
			if err := m.HasHasAnimal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTestdata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTestdata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BadMultiSignature) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestdata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BadMultiSignature: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BadMultiSignature: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestdata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTestdata
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTestdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatures = append(m.Signatures, make([]byte, postIndex-iNdEx))
			copy(m.Signatures[len(m.Signatures)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaliciousField", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestdata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTestdata
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTestdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaliciousField = append(m.MaliciousField[:0], dAtA[iNdEx:postIndex]...)
			if m.MaliciousField == nil {
				m.MaliciousField = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTestdata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTestdata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TableModel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestdata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TableModel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TableModel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestdata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTestdata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTestdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestdata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTestdata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTestdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTestdata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTestdata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTestdata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTestdata
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTestdata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTestdata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTestdata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTestdata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTestdata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTestdata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTestdata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTestdata = fmt.Errorf("proto: unexpected end of group")
)
