// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/pkg/storage/fileset/index/index.proto

package index

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	chunk "github.com/pachyderm/pachyderm/src/server/pkg/storage/chunk"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Op is the set of operations that can be associated with data.
type Op int32

const (
	Op_APPEND    Op = 0
	Op_OVERWRITE Op = 1
	Op_DELETE    Op = 2
)

var Op_name = map[int32]string{
	0: "APPEND",
	1: "OVERWRITE",
	2: "DELETE",
}

var Op_value = map[string]int32{
	"APPEND":    0,
	"OVERWRITE": 1,
	"DELETE":    2,
}

func (x Op) String() string {
	return proto.EnumName(Op_name, int32(x))
}

func (Op) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5610f63adbdd53a8, []int{0}
}

type Range struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	LastPath             string   `protobuf:"bytes,2,opt,name=last_path,json=lastPath,proto3" json:"last_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Range) Reset()         { *m = Range{} }
func (m *Range) String() string { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()    {}
func (*Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_5610f63adbdd53a8, []int{0}
}
func (m *Range) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Range.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Range.Merge(m, src)
}
func (m *Range) XXX_Size() int {
	return m.Size()
}
func (m *Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Range.DiscardUnknown(m)
}

var xxx_messageInfo_Range proto.InternalMessageInfo

func (m *Range) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Range) GetLastPath() string {
	if m != nil {
		return m.LastPath
	}
	return ""
}

// DataOp is a sequence of data references and an operation associated with the referenced data.
// Tags map identifiers to data in the resulting byte stream.
type DataOp struct {
	DataRefs             []*chunk.DataRef `protobuf:"bytes,1,rep,name=data_refs,json=dataRefs,proto3" json:"data_refs,omitempty"`
	Op                   Op               `protobuf:"varint,2,opt,name=op,proto3,enum=index.Op" json:"op,omitempty"`
	Tags                 []*Tag           `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DataOp) Reset()         { *m = DataOp{} }
func (m *DataOp) String() string { return proto.CompactTextString(m) }
func (*DataOp) ProtoMessage()    {}
func (*DataOp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5610f63adbdd53a8, []int{1}
}
func (m *DataOp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataOp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataOp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataOp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataOp.Merge(m, src)
}
func (m *DataOp) XXX_Size() int {
	return m.Size()
}
func (m *DataOp) XXX_DiscardUnknown() {
	xxx_messageInfo_DataOp.DiscardUnknown(m)
}

var xxx_messageInfo_DataOp proto.InternalMessageInfo

func (m *DataOp) GetDataRefs() []*chunk.DataRef {
	if m != nil {
		return m.DataRefs
	}
	return nil
}

func (m *DataOp) GetOp() Op {
	if m != nil {
		return m.Op
	}
	return Op_APPEND
}

func (m *DataOp) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Tag struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SizeBytes            int64    `protobuf:"varint,2,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_5610f63adbdd53a8, []int{2}
}
func (m *Tag) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return m.Size()
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tag) GetSizeBytes() int64 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

// Index stores an index to and metadata about a tar stream.
type Index struct {
	Range  *Range  `protobuf:"bytes,1,opt,name=range,proto3" json:"range,omitempty"`
	DataOp *DataOp `protobuf:"bytes,2,opt,name=data_op,json=dataOp,proto3" json:"data_op,omitempty"`
	// Size of the content being indexed (does not include headers).
	SizeBytes            int64    `protobuf:"varint,3,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_5610f63adbdd53a8, []int{3}
}
func (m *Index) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Index.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(m, src)
}
func (m *Index) XXX_Size() int {
	return m.Size()
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetRange() *Range {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *Index) GetDataOp() *DataOp {
	if m != nil {
		return m.DataOp
	}
	return nil
}

func (m *Index) GetSizeBytes() int64 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

func init() {
	proto.RegisterEnum("index.Op", Op_name, Op_value)
	proto.RegisterType((*Range)(nil), "index.Range")
	proto.RegisterType((*DataOp)(nil), "index.DataOp")
	proto.RegisterType((*Tag)(nil), "index.Tag")
	proto.RegisterType((*Index)(nil), "index.Index")
}

func init() {
	proto.RegisterFile("server/pkg/storage/fileset/index/index.proto", fileDescriptor_5610f63adbdd53a8)
}

var fileDescriptor_5610f63adbdd53a8 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0xea, 0xd3, 0x40,
	0x10, 0xfe, 0x6d, 0x62, 0x62, 0x33, 0xb5, 0xa5, 0xec, 0x41, 0xa2, 0x62, 0x28, 0x41, 0xa4, 0x58,
	0x49, 0x20, 0x7a, 0xf4, 0x62, 0x6d, 0x0e, 0x05, 0x31, 0x65, 0x89, 0x0a, 0x5e, 0xca, 0x36, 0xd9,
	0xfc, 0xa1, 0xb5, 0x59, 0x76, 0xb7, 0x62, 0x7d, 0x12, 0x1f, 0xc9, 0xa3, 0x8f, 0x20, 0xf5, 0x45,
	0x24, 0xbb, 0x39, 0xa8, 0x08, 0x5e, 0x86, 0x99, 0x6f, 0x66, 0xbe, 0xf9, 0xbe, 0x64, 0xe1, 0xa9,
	0x64, 0xe2, 0x13, 0x13, 0x31, 0x3f, 0xd4, 0xb1, 0x54, 0x9d, 0xa0, 0x35, 0x8b, 0xab, 0xf6, 0xc8,
	0x24, 0x53, 0x71, 0x7b, 0x2a, 0xd9, 0x67, 0x13, 0x23, 0x2e, 0x3a, 0xd5, 0x61, 0x47, 0x17, 0xf7,
	0x1f, 0xfd, 0x63, 0xa9, 0x68, 0xce, 0xa7, 0x83, 0x89, 0x66, 0x38, 0x7c, 0x01, 0x0e, 0xa1, 0xa7,
	0x9a, 0xe1, 0xbb, 0xe0, 0x76, 0x55, 0x25, 0x99, 0xf2, 0xd1, 0x1c, 0x2d, 0x6c, 0x32, 0x54, 0xf8,
	0x01, 0x78, 0x47, 0x2a, 0xd5, 0x8e, 0x53, 0xd5, 0xf8, 0xd6, 0x1c, 0x2d, 0x3c, 0x32, 0xea, 0x81,
	0x2d, 0x55, 0x4d, 0xc8, 0xc1, 0x5d, 0x53, 0x45, 0x33, 0x8e, 0x97, 0xe0, 0x95, 0x54, 0xd1, 0x9d,
	0x60, 0x95, 0xf4, 0xd1, 0xdc, 0x5e, 0x8c, 0x93, 0x69, 0x64, 0x0e, 0xf5, 0x13, 0x84, 0x55, 0x64,
	0x54, 0x9a, 0x44, 0xe2, 0x7b, 0x60, 0x75, 0x5c, 0x93, 0x4d, 0x13, 0x2f, 0x32, 0xda, 0x33, 0x4e,
	0xac, 0x8e, 0xe3, 0x00, 0x6e, 0x29, 0x5a, 0x4b, 0xdf, 0xd6, 0x14, 0x30, 0x34, 0x73, 0x5a, 0x13,
	0x8d, 0x87, 0xcf, 0xc1, 0xce, 0x69, 0x8d, 0xa7, 0x60, 0xb5, 0xa5, 0x56, 0xea, 0x11, 0xab, 0x2d,
	0xf1, 0x43, 0x00, 0xd9, 0x7e, 0x61, 0xbb, 0xfd, 0x45, 0x31, 0xa9, 0x99, 0x6d, 0xe2, 0xf5, 0xc8,
	0xaa, 0x07, 0x42, 0x01, 0xce, 0xa6, 0x27, 0xc2, 0x21, 0x38, 0xa2, 0xb7, 0xab, 0x57, 0xc7, 0xc9,
	0x9d, 0x81, 0x5f, 0x7f, 0x02, 0x62, 0x5a, 0xf8, 0x31, 0xdc, 0xd6, 0x56, 0x06, 0x89, 0xe3, 0x64,
	0x32, 0x4c, 0x19, 0xab, 0xc4, 0x2d, 0x8d, 0xe5, 0x3f, 0x6f, 0xda, 0x7f, 0xdd, 0x7c, 0xb2, 0x04,
	0x2b, 0xe3, 0x18, 0xc0, 0x7d, 0xb9, 0xdd, 0xa6, 0x6f, 0xd6, 0xb3, 0x1b, 0x3c, 0x01, 0x2f, 0x7b,
	0x97, 0x92, 0xf7, 0x64, 0x93, 0xa7, 0x33, 0xd4, 0xb7, 0xd6, 0xe9, 0xeb, 0x34, 0x4f, 0x67, 0xd6,
	0xea, 0xed, 0xb7, 0x6b, 0x80, 0xbe, 0x5f, 0x03, 0xf4, 0xe3, 0x1a, 0xa0, 0xaf, 0x3f, 0x83, 0x9b,
	0x0f, 0xaf, 0xea, 0x56, 0x35, 0xe7, 0x7d, 0x54, 0x74, 0x1f, 0x63, 0x4e, 0x8b, 0xe6, 0x52, 0x32,
	0xf1, 0x7b, 0x26, 0x45, 0x11, 0xff, 0xef, 0x59, 0xec, 0x5d, 0xfd, 0x93, 0x9f, 0xfd, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0x42, 0x77, 0x67, 0x93, 0x41, 0x02, 0x00, 0x00,
}

func (m *Range) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Range) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Range) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.LastPath) > 0 {
		i -= len(m.LastPath)
		copy(dAtA[i:], m.LastPath)
		i = encodeVarintIndex(dAtA, i, uint64(len(m.LastPath)))
		i--
		dAtA[i] = 0x12
	}
	if m.Offset != 0 {
		i = encodeVarintIndex(dAtA, i, uint64(m.Offset))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DataOp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataOp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataOp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Tags) > 0 {
		for iNdEx := len(m.Tags) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tags[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIndex(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Op != 0 {
		i = encodeVarintIndex(dAtA, i, uint64(m.Op))
		i--
		dAtA[i] = 0x10
	}
	if len(m.DataRefs) > 0 {
		for iNdEx := len(m.DataRefs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DataRefs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIndex(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Tag) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tag) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Tag) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.SizeBytes != 0 {
		i = encodeVarintIndex(dAtA, i, uint64(m.SizeBytes))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintIndex(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Index) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Index) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Index) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.SizeBytes != 0 {
		i = encodeVarintIndex(dAtA, i, uint64(m.SizeBytes))
		i--
		dAtA[i] = 0x18
	}
	if m.DataOp != nil {
		{
			size, err := m.DataOp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIndex(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Range != nil {
		{
			size, err := m.Range.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIndex(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIndex(dAtA []byte, offset int, v uint64) int {
	offset -= sovIndex(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Range) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Offset != 0 {
		n += 1 + sovIndex(uint64(m.Offset))
	}
	l = len(m.LastPath)
	if l > 0 {
		n += 1 + l + sovIndex(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DataOp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DataRefs) > 0 {
		for _, e := range m.DataRefs {
			l = e.Size()
			n += 1 + l + sovIndex(uint64(l))
		}
	}
	if m.Op != 0 {
		n += 1 + sovIndex(uint64(m.Op))
	}
	if len(m.Tags) > 0 {
		for _, e := range m.Tags {
			l = e.Size()
			n += 1 + l + sovIndex(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Tag) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovIndex(uint64(l))
	}
	if m.SizeBytes != 0 {
		n += 1 + sovIndex(uint64(m.SizeBytes))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Index) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Range != nil {
		l = m.Range.Size()
		n += 1 + l + sovIndex(uint64(l))
	}
	if m.DataOp != nil {
		l = m.DataOp.Size()
		n += 1 + l + sovIndex(uint64(l))
	}
	if m.SizeBytes != 0 {
		n += 1 + sovIndex(uint64(m.SizeBytes))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIndex(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIndex(x uint64) (n int) {
	return sovIndex(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Range) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndex
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
			return fmt.Errorf("proto: Range: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Range: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
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
				return ErrInvalidLengthIndex
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIndex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIndex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIndex
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIndex
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
func (m *DataOp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndex
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
			return fmt.Errorf("proto: DataOp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataOp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataRefs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
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
				return ErrInvalidLengthIndex
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIndex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataRefs = append(m.DataRefs, &chunk.DataRef{})
			if err := m.DataRefs[len(m.DataRefs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= Op(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
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
				return ErrInvalidLengthIndex
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIndex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, &Tag{})
			if err := m.Tags[len(m.Tags)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIndex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIndex
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIndex
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
func (m *Tag) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndex
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
			return fmt.Errorf("proto: Tag: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tag: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
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
				return ErrInvalidLengthIndex
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIndex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SizeBytes", wireType)
			}
			m.SizeBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SizeBytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIndex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIndex
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIndex
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
func (m *Index) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndex
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
			return fmt.Errorf("proto: Index: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Index: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Range", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
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
				return ErrInvalidLengthIndex
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIndex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Range == nil {
				m.Range = &Range{}
			}
			if err := m.Range.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataOp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
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
				return ErrInvalidLengthIndex
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIndex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DataOp == nil {
				m.DataOp = &DataOp{}
			}
			if err := m.DataOp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SizeBytes", wireType)
			}
			m.SizeBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SizeBytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIndex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIndex
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIndex
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
func skipIndex(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIndex
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
					return 0, ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIndex
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
				return 0, ErrInvalidLengthIndex
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthIndex
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIndex
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipIndex(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthIndex
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthIndex = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIndex   = fmt.Errorf("proto: integer overflow")
)
