// Code generated by protoc-gen-go. DO NOT EDIT.
// source: execution/record.proto

package execution

import (
	fmt "fmt"
	octosql "github.com/cube2222/octosql"
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

type RecordID struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordID) Reset()         { *m = RecordID{} }
func (m *RecordID) String() string { return proto.CompactTextString(m) }
func (*RecordID) ProtoMessage()    {}
func (*RecordID) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd7ef896dd4ca095, []int{0}
}

func (m *RecordID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordID.Unmarshal(m, b)
}
func (m *RecordID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordID.Marshal(b, m, deterministic)
}
func (m *RecordID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordID.Merge(m, src)
}
func (m *RecordID) XXX_Size() int {
	return xxx_messageInfo_RecordID.Size(m)
}
func (m *RecordID) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordID.DiscardUnknown(m)
}

var xxx_messageInfo_RecordID proto.InternalMessageInfo

func (m *RecordID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Metadata struct {
	Id                   *RecordID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Undo                 bool      `protobuf:"varint,2,opt,name=undo,proto3" json:"undo,omitempty"`
	EventTimeField       string    `protobuf:"bytes,3,opt,name=eventTimeField,proto3" json:"eventTimeField,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd7ef896dd4ca095, []int{1}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetId() *RecordID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Metadata) GetUndo() bool {
	if m != nil {
		return m.Undo
	}
	return false
}

func (m *Metadata) GetEventTimeField() string {
	if m != nil {
		return m.EventTimeField
	}
	return ""
}

type Record struct {
	Metadata             *Metadata        `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	FieldNames           []string         `protobuf:"bytes,2,rep,name=fieldNames,proto3" json:"fieldNames,omitempty"`
	Data                 []*octosql.Value `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd7ef896dd4ca095, []int{2}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Record) GetFieldNames() []string {
	if m != nil {
		return m.FieldNames
	}
	return nil
}

func (m *Record) GetData() []*octosql.Value {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*RecordID)(nil), "execution.RecordID")
	proto.RegisterType((*Metadata)(nil), "execution.Metadata")
	proto.RegisterType((*Record)(nil), "execution.Record")
}

func init() { proto.RegisterFile("execution/record.proto", fileDescriptor_fd7ef896dd4ca095) }

var fileDescriptor_fd7ef896dd4ca095 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x14, 0x84, 0xd9, 0x74, 0x59, 0xda, 0xb7, 0xd2, 0x43, 0x04, 0x29, 0x7b, 0x90, 0x52, 0x51, 0x7b,
	0x4a, 0x21, 0xfe, 0x03, 0x29, 0x42, 0x0f, 0x7a, 0x08, 0xe2, 0xc1, 0x5b, 0x9a, 0x3c, 0x35, 0xd8,
	0x36, 0xda, 0x26, 0x8b, 0x17, 0xff, 0xbb, 0x6c, 0x76, 0xb7, 0x88, 0x78, 0x0b, 0x33, 0x93, 0xf9,
	0x1e, 0x03, 0x67, 0xf8, 0x85, 0xca, 0x3b, 0x63, 0x87, 0x6a, 0x44, 0x65, 0x47, 0xcd, 0x3e, 0x46,
	0xeb, 0x2c, 0x4d, 0x66, 0x7d, 0x73, 0xb2, 0x95, 0x9d, 0xc7, 0x69, 0x6f, 0x14, 0x1b, 0x88, 0x45,
	0x08, 0x36, 0x35, 0x4d, 0x81, 0x34, 0x75, 0xb6, 0xc8, 0x17, 0x65, 0x22, 0x48, 0x53, 0x17, 0xef,
	0x10, 0xdf, 0xa3, 0x93, 0x5a, 0x3a, 0x49, 0x2f, 0x80, 0x18, 0x1d, 0xbc, 0x35, 0x3f, 0x65, 0x73,
	0x1b, 0x3b, 0x7e, 0x16, 0xc4, 0x68, 0x4a, 0x61, 0xe9, 0x07, 0x6d, 0x33, 0x92, 0x2f, 0xca, 0x58,
	0x84, 0x37, 0xbd, 0x82, 0x14, 0xb7, 0x38, 0xb8, 0x47, 0xd3, 0xe3, 0x9d, 0xc1, 0x4e, 0x67, 0x51,
	0x00, 0xfc, 0x51, 0x8b, 0x6f, 0x58, 0xed, 0xbb, 0x68, 0x05, 0x71, 0x7f, 0xc0, 0xfe, 0x03, 0x3c,
	0x5e, 0x24, 0xe6, 0x10, 0x3d, 0x07, 0x78, 0xd9, 0x75, 0x3c, 0xc8, 0x1e, 0xa7, 0x8c, 0xe4, 0x51,
	0x99, 0x88, 0x5f, 0x0a, 0x2d, 0x60, 0x19, 0xca, 0xa2, 0x3c, 0x2a, 0xd7, 0x3c, 0x65, 0x56, 0x39,
	0x3b, 0x7d, 0x76, 0xec, 0x69, 0x37, 0x84, 0x08, 0xde, 0xed, 0xf5, 0xf3, 0xe5, 0xab, 0x71, 0x6f,
	0xbe, 0x65, 0xca, 0xf6, 0x95, 0xf2, 0x2d, 0x72, 0xce, 0x79, 0x75, 0x88, 0x56, 0x33, 0xbf, 0x5d,
	0x85, 0xdd, 0x6e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xde, 0x23, 0xe8, 0x56, 0x6a, 0x01, 0x00,
	0x00,
}
