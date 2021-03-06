// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/decoderbufs.proto

/*
Package decoderbufs is a generated protocol buffer package.

It is generated from these files:
	proto/decoderbufs.proto

It has these top-level messages:
	Point
	DatumMessage
	TypeInfo
	RowMessage
*/
package decoderbufs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Op int32

const (
	Op_INSERT Op = 0
	Op_UPDATE Op = 1
	Op_DELETE Op = 2
)

var Op_name = map[int32]string{
	0: "INSERT",
	1: "UPDATE",
	2: "DELETE",
}
var Op_value = map[string]int32{
	"INSERT": 0,
	"UPDATE": 1,
	"DELETE": 2,
}

func (x Op) Enum() *Op {
	p := new(Op)
	*p = x
	return p
}
func (x Op) String() string {
	return proto.EnumName(Op_name, int32(x))
}
func (x *Op) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Op_value, data, "Op")
	if err != nil {
		return err
	}
	*x = Op(value)
	return nil
}
func (Op) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Point struct {
	X                *float64 `protobuf:"fixed64,1,req,name=x" json:"x,omitempty"`
	Y                *float64 `protobuf:"fixed64,2,req,name=y" json:"y,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Point) GetX() float64 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Point) GetY() float64 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

type DatumMessage struct {
	ColumnName       *string  `protobuf:"bytes,1,opt,name=column_name,json=columnName" json:"column_name,omitempty"`
	ColumnType       *int64   `protobuf:"varint,2,opt,name=column_type,json=columnType" json:"column_type,omitempty"`
	DatumInt32       *int32   `protobuf:"varint,3,opt,name=datum_int32,json=datumInt32" json:"datum_int32,omitempty"`
	DatumInt64       *int64   `protobuf:"varint,4,opt,name=datum_int64,json=datumInt64" json:"datum_int64,omitempty"`
	DatumFloat       *float32 `protobuf:"fixed32,5,opt,name=datum_float,json=datumFloat" json:"datum_float,omitempty"`
	DatumDouble      *float64 `protobuf:"fixed64,6,opt,name=datum_double,json=datumDouble" json:"datum_double,omitempty"`
	DatumBool        *bool    `protobuf:"varint,7,opt,name=datum_bool,json=datumBool" json:"datum_bool,omitempty"`
	DatumString      *string  `protobuf:"bytes,8,opt,name=datum_string,json=datumString" json:"datum_string,omitempty"`
	DatumBytes       []byte   `protobuf:"bytes,9,opt,name=datum_bytes,json=datumBytes" json:"datum_bytes,omitempty"`
	DatumPoint       *Point   `protobuf:"bytes,10,opt,name=datum_point,json=datumPoint" json:"datum_point,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DatumMessage) Reset()                    { *m = DatumMessage{} }
func (m *DatumMessage) String() string            { return proto.CompactTextString(m) }
func (*DatumMessage) ProtoMessage()               {}
func (*DatumMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DatumMessage) GetColumnName() string {
	if m != nil && m.ColumnName != nil {
		return *m.ColumnName
	}
	return ""
}

func (m *DatumMessage) GetColumnType() int64 {
	if m != nil && m.ColumnType != nil {
		return *m.ColumnType
	}
	return 0
}

func (m *DatumMessage) GetDatumInt32() int32 {
	if m != nil && m.DatumInt32 != nil {
		return *m.DatumInt32
	}
	return 0
}

func (m *DatumMessage) GetDatumInt64() int64 {
	if m != nil && m.DatumInt64 != nil {
		return *m.DatumInt64
	}
	return 0
}

func (m *DatumMessage) GetDatumFloat() float32 {
	if m != nil && m.DatumFloat != nil {
		return *m.DatumFloat
	}
	return 0
}

func (m *DatumMessage) GetDatumDouble() float64 {
	if m != nil && m.DatumDouble != nil {
		return *m.DatumDouble
	}
	return 0
}

func (m *DatumMessage) GetDatumBool() bool {
	if m != nil && m.DatumBool != nil {
		return *m.DatumBool
	}
	return false
}

func (m *DatumMessage) GetDatumString() string {
	if m != nil && m.DatumString != nil {
		return *m.DatumString
	}
	return ""
}

func (m *DatumMessage) GetDatumBytes() []byte {
	if m != nil {
		return m.DatumBytes
	}
	return nil
}

func (m *DatumMessage) GetDatumPoint() *Point {
	if m != nil {
		return m.DatumPoint
	}
	return nil
}

type TypeInfo struct {
	Modifier         *string `protobuf:"bytes,1,req,name=modifier" json:"modifier,omitempty"`
	ValueOptional    *bool   `protobuf:"varint,2,req,name=value_optional,json=valueOptional" json:"value_optional,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TypeInfo) Reset()                    { *m = TypeInfo{} }
func (m *TypeInfo) String() string            { return proto.CompactTextString(m) }
func (*TypeInfo) ProtoMessage()               {}
func (*TypeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TypeInfo) GetModifier() string {
	if m != nil && m.Modifier != nil {
		return *m.Modifier
	}
	return ""
}

func (m *TypeInfo) GetValueOptional() bool {
	if m != nil && m.ValueOptional != nil {
		return *m.ValueOptional
	}
	return false
}

type RowMessage struct {
	TransactionId    *uint32         `protobuf:"varint,1,opt,name=transaction_id,json=transactionId" json:"transaction_id,omitempty"`
	CommitTime       *uint64         `protobuf:"varint,2,opt,name=commit_time,json=commitTime" json:"commit_time,omitempty"`
	Table            *string         `protobuf:"bytes,3,opt,name=table" json:"table,omitempty"`
	Op               *Op             `protobuf:"varint,4,opt,name=op,enum=decoderbufs.Op" json:"op,omitempty"`
	NewTuple         []*DatumMessage `protobuf:"bytes,5,rep,name=new_tuple,json=newTuple" json:"new_tuple,omitempty"`
	OldTuple         []*DatumMessage `protobuf:"bytes,6,rep,name=old_tuple,json=oldTuple" json:"old_tuple,omitempty"`
	NewTypeinfo      []*TypeInfo     `protobuf:"bytes,7,rep,name=new_typeinfo,json=newTypeinfo" json:"new_typeinfo,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RowMessage) Reset()                    { *m = RowMessage{} }
func (m *RowMessage) String() string            { return proto.CompactTextString(m) }
func (*RowMessage) ProtoMessage()               {}
func (*RowMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RowMessage) GetTransactionId() uint32 {
	if m != nil && m.TransactionId != nil {
		return *m.TransactionId
	}
	return 0
}

func (m *RowMessage) GetCommitTime() uint64 {
	if m != nil && m.CommitTime != nil {
		return *m.CommitTime
	}
	return 0
}

func (m *RowMessage) GetTable() string {
	if m != nil && m.Table != nil {
		return *m.Table
	}
	return ""
}

func (m *RowMessage) GetOp() Op {
	if m != nil && m.Op != nil {
		return *m.Op
	}
	return Op_INSERT
}

func (m *RowMessage) GetNewTuple() []*DatumMessage {
	if m != nil {
		return m.NewTuple
	}
	return nil
}

func (m *RowMessage) GetOldTuple() []*DatumMessage {
	if m != nil {
		return m.OldTuple
	}
	return nil
}

func (m *RowMessage) GetNewTypeinfo() []*TypeInfo {
	if m != nil {
		return m.NewTypeinfo
	}
	return nil
}

func init() {
	proto.RegisterType((*Point)(nil), "decoderbufs.Point")
	proto.RegisterType((*DatumMessage)(nil), "decoderbufs.DatumMessage")
	proto.RegisterType((*TypeInfo)(nil), "decoderbufs.TypeInfo")
	proto.RegisterType((*RowMessage)(nil), "decoderbufs.RowMessage")
	proto.RegisterEnum("decoderbufs.Op", Op_name, Op_value)
}

func init() { proto.RegisterFile("proto/decoderbufs.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xd1, 0x6e, 0xd3, 0x3e,
	0x14, 0xc6, 0xff, 0x4e, 0xd7, 0xad, 0x3d, 0xed, 0xf6, 0x9f, 0x2c, 0x10, 0x06, 0x09, 0x2d, 0x14,
	0x0d, 0x45, 0x5c, 0x14, 0xa9, 0x9b, 0x26, 0x6e, 0xa9, 0x5a, 0x44, 0x25, 0xb6, 0x55, 0x5e, 0xb9,
	0x8e, 0xd2, 0xc6, 0xad, 0x2c, 0x25, 0x3e, 0x26, 0x71, 0xd8, 0xc2, 0x7b, 0xf1, 0x5a, 0x3c, 0x03,
	0xb2, 0x9d, 0xd2, 0xf6, 0x8a, 0xbb, 0x73, 0x7e, 0xfd, 0xbe, 0x63, 0xfb, 0x7c, 0x0d, 0xbc, 0xd0,
	0x05, 0x1a, 0xfc, 0x90, 0x8a, 0x15, 0xa6, 0xa2, 0x58, 0x56, 0xeb, 0x72, 0xe8, 0x08, 0xed, 0xed,
	0xa1, 0xc1, 0x5b, 0x68, 0xcf, 0x51, 0x2a, 0x43, 0xfb, 0x40, 0x9e, 0x18, 0x09, 0x83, 0x88, 0x70,
	0xf2, 0x64, 0xbb, 0x9a, 0x05, 0xbe, 0xab, 0x07, 0xbf, 0x03, 0xe8, 0x4f, 0x12, 0x53, 0xe5, 0xb7,
	0xa2, 0x2c, 0x93, 0x8d, 0xa0, 0x17, 0xd0, 0x5b, 0x61, 0x56, 0xe5, 0x2a, 0x56, 0x49, 0x2e, 0x18,
	0x09, 0x49, 0xd4, 0xe5, 0xe0, 0xd1, 0x5d, 0x92, 0xef, 0x0b, 0x4c, 0xad, 0x05, 0x0b, 0x42, 0x12,
	0xb5, 0xb6, 0x82, 0x45, 0xad, 0x9d, 0x20, 0xb5, 0x13, 0x63, 0xa9, 0xcc, 0xd5, 0x88, 0xb5, 0x42,
	0x12, 0xb5, 0x39, 0x38, 0x34, 0xb3, 0xe4, 0x40, 0x70, 0x73, 0xcd, 0x8e, 0xfc, 0x84, 0xad, 0xe0,
	0xe6, 0x7a, 0x27, 0x58, 0x67, 0x98, 0x18, 0xd6, 0x0e, 0x49, 0x14, 0x34, 0x82, 0xcf, 0x96, 0xd0,
	0x37, 0xd0, 0xf7, 0x82, 0x14, 0xab, 0x65, 0x26, 0xd8, 0x71, 0x48, 0x22, 0xc2, 0xbd, 0x69, 0xe2,
	0x10, 0x7d, 0x0d, 0xde, 0x10, 0x2f, 0x11, 0x33, 0x76, 0x12, 0x92, 0xa8, 0xc3, 0xbb, 0x8e, 0x8c,
	0x11, 0xb3, 0xdd, 0x84, 0xd2, 0x14, 0x52, 0x6d, 0x58, 0xc7, 0xbd, 0xd3, 0x4f, 0x78, 0x70, 0x68,
	0x77, 0x8b, 0x65, 0x6d, 0x44, 0xc9, 0xba, 0x21, 0x89, 0xfa, 0xcd, 0x2d, 0xc6, 0x96, 0xd0, 0xab,
	0xad, 0x40, 0xdb, 0x35, 0x33, 0x08, 0x49, 0xd4, 0x1b, 0xd1, 0xe1, 0x7e, 0x2c, 0x2e, 0x80, 0xc6,
	0xe4, 0xea, 0xc1, 0x2d, 0x74, 0xec, 0x96, 0x66, 0x6a, 0x8d, 0xf4, 0x15, 0x74, 0x72, 0x4c, 0xe5,
	0x5a, 0x8a, 0xc2, 0xe5, 0xd3, 0xe5, 0x7f, 0x7b, 0x7a, 0x09, 0x67, 0x3f, 0x92, 0xac, 0x12, 0x31,
	0x6a, 0x23, 0x51, 0x25, 0x99, 0xcb, 0xac, 0xc3, 0x4f, 0x1d, 0xbd, 0x6f, 0xe0, 0xe0, 0x57, 0x00,
	0xc0, 0xf1, 0x71, 0x9b, 0xde, 0x25, 0x9c, 0x99, 0x22, 0x51, 0x65, 0xb2, 0xb2, 0xbf, 0xc7, 0x32,
	0x75, 0x01, 0x9e, 0xf2, 0xd3, 0x3d, 0x3a, 0x4b, 0x7d, 0x86, 0x79, 0x2e, 0x4d, 0x6c, 0x64, 0xee,
	0x33, 0x3c, 0xb2, 0x19, 0x5a, 0xb4, 0x90, 0xb9, 0xa0, 0xcf, 0xa0, 0x6d, 0x12, 0xbb, 0xd9, 0x96,
	0xdb, 0x8b, 0x6f, 0xe8, 0x05, 0x04, 0xa8, 0x5d, 0x5e, 0x67, 0xa3, 0xff, 0x0f, 0xde, 0x79, 0xaf,
	0x79, 0x80, 0x9a, 0xde, 0x40, 0x57, 0x89, 0xc7, 0xd8, 0x54, 0x3a, 0x13, 0xac, 0x1d, 0xb6, 0xa2,
	0xde, 0xe8, 0xe5, 0x81, 0x6e, 0xff, 0xaf, 0xc6, 0x3b, 0x4a, 0x3c, 0x2e, 0xac, 0xd4, 0xfa, 0x30,
	0x4b, 0x1b, 0xdf, 0xf1, 0x3f, 0x7d, 0x98, 0xa5, 0xde, 0xf7, 0x11, 0xfa, 0xee, 0xbc, 0x5a, 0x0b,
	0xa9, 0xd6, 0xc8, 0x4e, 0x9c, 0xf5, 0xf9, 0x81, 0x75, 0xbb, 0x6d, 0xde, 0xb3, 0xc7, 0x35, 0xca,
	0xf7, 0x11, 0x04, 0xf7, 0x9a, 0x02, 0x1c, 0xcf, 0xee, 0x1e, 0xa6, 0x7c, 0x71, 0xfe, 0x9f, 0xad,
	0xbf, 0xcd, 0x27, 0x9f, 0x16, 0xd3, 0x73, 0x62, 0xeb, 0xc9, 0xf4, 0xeb, 0x74, 0x31, 0x3d, 0x0f,
	0xc6, 0x57, 0xf0, 0x4e, 0xe2, 0x30, 0x15, 0x4b, 0xf1, 0x53, 0x56, 0xf9, 0x70, 0x85, 0x4a, 0x89,
	0x95, 0xc1, 0x62, 0xa8, 0xb1, 0x34, 0x9b, 0x42, 0x94, 0xdf, 0x33, 0xff, 0xf5, 0x8d, 0x4f, 0xe6,
	0x9b, 0xb9, 0x2d, 0xbe, 0x90, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xb3, 0x5a, 0x22, 0xa2,
	0x03, 0x00, 0x00,
}
