// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/query.proto

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

type RequestList struct {
	Index                int64             `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Limit                int64             `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Gt                   int64             `protobuf:"varint,3,opt,name=gt,proto3" json:"gt,omitempty"`
	Lt                   int64             `protobuf:"varint,4,opt,name=lt,proto3" json:"lt,omitempty"`
	Params               map[string]string `protobuf:"bytes,5,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RequestList) Reset()         { *m = RequestList{} }
func (m *RequestList) String() string { return proto.CompactTextString(m) }
func (*RequestList) ProtoMessage()    {}
func (*RequestList) Descriptor() ([]byte, []int) {
	return fileDescriptor_5958159b9fcae252, []int{0}
}

func (m *RequestList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestList.Unmarshal(m, b)
}
func (m *RequestList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestList.Marshal(b, m, deterministic)
}
func (m *RequestList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestList.Merge(m, src)
}
func (m *RequestList) XXX_Size() int {
	return xxx_messageInfo_RequestList.Size(m)
}
func (m *RequestList) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestList.DiscardUnknown(m)
}

var xxx_messageInfo_RequestList proto.InternalMessageInfo

func (m *RequestList) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *RequestList) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *RequestList) GetGt() int64 {
	if m != nil {
		return m.Gt
	}
	return 0
}

func (m *RequestList) GetLt() int64 {
	if m != nil {
		return m.Lt
	}
	return 0
}

func (m *RequestList) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

type RequestById struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestById) Reset()         { *m = RequestById{} }
func (m *RequestById) String() string { return proto.CompactTextString(m) }
func (*RequestById) ProtoMessage()    {}
func (*RequestById) Descriptor() ([]byte, []int) {
	return fileDescriptor_5958159b9fcae252, []int{1}
}

func (m *RequestById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestById.Unmarshal(m, b)
}
func (m *RequestById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestById.Marshal(b, m, deterministic)
}
func (m *RequestById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestById.Merge(m, src)
}
func (m *RequestById) XXX_Size() int {
	return xxx_messageInfo_RequestById.Size(m)
}
func (m *RequestById) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestById.DiscardUnknown(m)
}

var xxx_messageInfo_RequestById proto.InternalMessageInfo

func (m *RequestById) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RequestByName struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestByName) Reset()         { *m = RequestByName{} }
func (m *RequestByName) String() string { return proto.CompactTextString(m) }
func (*RequestByName) ProtoMessage()    {}
func (*RequestByName) Descriptor() ([]byte, []int) {
	return fileDescriptor_5958159b9fcae252, []int{2}
}

func (m *RequestByName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestByName.Unmarshal(m, b)
}
func (m *RequestByName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestByName.Marshal(b, m, deterministic)
}
func (m *RequestByName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestByName.Merge(m, src)
}
func (m *RequestByName) XXX_Size() int {
	return xxx_messageInfo_RequestByName.Size(m)
}
func (m *RequestByName) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestByName.DiscardUnknown(m)
}

var xxx_messageInfo_RequestByName proto.InternalMessageInfo

func (m *RequestByName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ResponseEffect struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Effect               int64    `protobuf:"varint,2,opt,name=effect,proto3" json:"effect,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseEffect) Reset()         { *m = ResponseEffect{} }
func (m *ResponseEffect) String() string { return proto.CompactTextString(m) }
func (*ResponseEffect) ProtoMessage()    {}
func (*ResponseEffect) Descriptor() ([]byte, []int) {
	return fileDescriptor_5958159b9fcae252, []int{3}
}

func (m *ResponseEffect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseEffect.Unmarshal(m, b)
}
func (m *ResponseEffect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseEffect.Marshal(b, m, deterministic)
}
func (m *ResponseEffect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseEffect.Merge(m, src)
}
func (m *ResponseEffect) XXX_Size() int {
	return xxx_messageInfo_ResponseEffect.Size(m)
}
func (m *ResponseEffect) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseEffect.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseEffect proto.InternalMessageInfo

func (m *ResponseEffect) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ResponseEffect) GetEffect() int64 {
	if m != nil {
		return m.Effect
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestList)(nil), "pb.requestList")
	proto.RegisterMapType((map[string]string)(nil), "pb.requestList.ParamsEntry")
	proto.RegisterType((*RequestById)(nil), "pb.requestById")
	proto.RegisterType((*RequestByName)(nil), "pb.requestByName")
	proto.RegisterType((*ResponseEffect)(nil), "pb.responseEffect")
}

func init() { proto.RegisterFile("pb/query.proto", fileDescriptor_5958159b9fcae252) }

var fileDescriptor_5958159b9fcae252 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0x69, 0xba, 0x55, 0x77, 0x8b, 0x45, 0x82, 0x48, 0x50, 0x84, 0x51, 0x5f, 0xf6, 0x54,
	0xc1, 0xbd, 0xa8, 0x8f, 0x83, 0x3d, 0x08, 0x22, 0xd2, 0x6f, 0x90, 0xb6, 0x77, 0x23, 0xd8, 0x3f,
	0x59, 0x92, 0x8a, 0xfd, 0x88, 0x7e, 0x2b, 0xe9, 0x6d, 0x36, 0xf6, 0x76, 0x7f, 0xe7, 0x1e, 0x92,
	0x73, 0x0f, 0x24, 0xba, 0x78, 0x3a, 0xf4, 0x68, 0x86, 0x4c, 0x9b, 0xce, 0x75, 0x9c, 0xe9, 0x22,
	0xfd, 0x0b, 0x20, 0x36, 0x78, 0xe8, 0xd1, 0xba, 0x0f, 0x65, 0x1d, 0xbf, 0x81, 0xb9, 0x6a, 0x2b,
	0xfc, 0x15, 0xc1, 0x32, 0x58, 0x85, 0xf9, 0x04, 0xa3, 0x5a, 0xab, 0x46, 0x39, 0xc1, 0x26, 0x95,
	0x80, 0x27, 0xc0, 0xf6, 0x4e, 0x84, 0x24, 0xb1, 0x3d, 0x71, 0xed, 0xc4, 0x6c, 0xe2, 0xda, 0xf1,
	0x35, 0x44, 0x5a, 0x1a, 0xd9, 0x58, 0x31, 0x5f, 0x86, 0xab, 0xf8, 0xf9, 0x3e, 0xd3, 0x45, 0x76,
	0xf6, 0x59, 0xf6, 0x45, 0xdb, 0x6d, 0xeb, 0xcc, 0x90, 0x7b, 0xeb, 0xdd, 0x2b, 0xc4, 0x67, 0x32,
	0xbf, 0x86, 0xf0, 0x1b, 0x07, 0x4a, 0xb3, 0xc8, 0xc7, 0x71, 0xcc, 0xf2, 0x23, 0xeb, 0x1e, 0x29,
	0xcb, 0x22, 0x9f, 0xe0, 0x8d, 0xbd, 0x04, 0xe9, 0xc3, 0xe9, 0x94, 0xcd, 0xf0, 0x5e, 0x8d, 0x71,
	0x54, 0xe5, 0xef, 0x60, 0xaa, 0x4a, 0x1f, 0xe1, 0xea, 0xb4, 0xfe, 0x94, 0x0d, 0x72, 0x0e, 0xb3,
	0x56, 0x36, 0xe8, 0x1f, 0xa7, 0x39, 0xdd, 0x40, 0x62, 0xd0, 0xea, 0xae, 0xb5, 0xb8, 0xdd, 0xed,
	0xb0, 0x74, 0x5c, 0xc0, 0x85, 0xed, 0xcb, 0x12, 0xad, 0x25, 0xe3, 0x65, 0x7e, 0x44, 0x7e, 0x0b,
	0x11, 0x92, 0xc7, 0xd7, 0xe2, 0xa9, 0x88, 0xa8, 0xde, 0xf5, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xb6, 0x21, 0x95, 0x1f, 0x70, 0x01, 0x00, 0x00,
}
