// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sumFunc.proto

package model

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SumRequest struct {
	A                    int64    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int64    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fdb574872ef000e, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *SumRequest) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

type SumReply struct {
	V                    int64    `protobuf:"varint,1,opt,name=v,proto3" json:"v,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumReply) Reset()         { *m = SumReply{} }
func (m *SumReply) String() string { return proto.CompactTextString(m) }
func (*SumReply) ProtoMessage()    {}
func (*SumReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fdb574872ef000e, []int{1}
}

func (m *SumReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumReply.Unmarshal(m, b)
}
func (m *SumReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumReply.Marshal(b, m, deterministic)
}
func (m *SumReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumReply.Merge(m, src)
}
func (m *SumReply) XXX_Size() int {
	return xxx_messageInfo_SumReply.Size(m)
}
func (m *SumReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SumReply.DiscardUnknown(m)
}

var xxx_messageInfo_SumReply proto.InternalMessageInfo

func (m *SumReply) GetV() int64 {
	if m != nil {
		return m.V
	}
	return 0
}

func (m *SumReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type ConcatRequest struct {
	Data1                string   `protobuf:"bytes,1,opt,name=data1,proto3" json:"data1,omitempty"`
	Data2                string   `protobuf:"bytes,2,opt,name=data2,proto3" json:"data2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConcatRequest) Reset()         { *m = ConcatRequest{} }
func (m *ConcatRequest) String() string { return proto.CompactTextString(m) }
func (*ConcatRequest) ProtoMessage()    {}
func (*ConcatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fdb574872ef000e, []int{2}
}

func (m *ConcatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConcatRequest.Unmarshal(m, b)
}
func (m *ConcatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConcatRequest.Marshal(b, m, deterministic)
}
func (m *ConcatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConcatRequest.Merge(m, src)
}
func (m *ConcatRequest) XXX_Size() int {
	return xxx_messageInfo_ConcatRequest.Size(m)
}
func (m *ConcatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConcatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConcatRequest proto.InternalMessageInfo

func (m *ConcatRequest) GetData1() string {
	if m != nil {
		return m.Data1
	}
	return ""
}

func (m *ConcatRequest) GetData2() string {
	if m != nil {
		return m.Data2
	}
	return ""
}

type ConcatReply struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConcatReply) Reset()         { *m = ConcatReply{} }
func (m *ConcatReply) String() string { return proto.CompactTextString(m) }
func (*ConcatReply) ProtoMessage()    {}
func (*ConcatReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fdb574872ef000e, []int{3}
}

func (m *ConcatReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConcatReply.Unmarshal(m, b)
}
func (m *ConcatReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConcatReply.Marshal(b, m, deterministic)
}
func (m *ConcatReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConcatReply.Merge(m, src)
}
func (m *ConcatReply) XXX_Size() int {
	return xxx_messageInfo_ConcatReply.Size(m)
}
func (m *ConcatReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ConcatReply.DiscardUnknown(m)
}

var xxx_messageInfo_ConcatReply proto.InternalMessageInfo

func (m *ConcatReply) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *ConcatReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "model.SumRequest")
	proto.RegisterType((*SumReply)(nil), "model.SumReply")
	proto.RegisterType((*ConcatRequest)(nil), "model.ConcatRequest")
	proto.RegisterType((*ConcatReply)(nil), "model.ConcatReply")
}

func init() { proto.RegisterFile("sumFunc.proto", fileDescriptor_7fdb574872ef000e) }

var fileDescriptor_7fdb574872ef000e = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0xcd, 0x75,
	0x2b, 0xcd, 0x4b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcd, 0x4f, 0x49, 0xcd,
	0x51, 0xd2, 0xe0, 0xe2, 0x0a, 0x2e, 0xcd, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2,
	0xe1, 0x62, 0x4c, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x62, 0x4c, 0x04, 0xf1, 0x92, 0x24,
	0x98, 0x20, 0xbc, 0x24, 0x25, 0x2d, 0x2e, 0x0e, 0xb0, 0xca, 0x82, 0x9c, 0x4a, 0x90, 0x4c, 0x19,
	0x4c, 0x5d, 0x99, 0x90, 0x00, 0x17, 0x73, 0x6a, 0x51, 0x11, 0x58, 0x25, 0x67, 0x10, 0x88, 0xa9,
	0x64, 0xcd, 0xc5, 0xeb, 0x9c, 0x9f, 0x97, 0x9c, 0x58, 0x02, 0x33, 0x58, 0x84, 0x8b, 0x35, 0x25,
	0xb1, 0x24, 0xd1, 0x10, 0xac, 0x89, 0x33, 0x08, 0xc2, 0x81, 0x89, 0x1a, 0x41, 0xb5, 0x42, 0x38,
	0x4a, 0xe6, 0x5c, 0xdc, 0x30, 0xcd, 0x20, 0xbb, 0xc4, 0xb8, 0xd8, 0x8a, 0x52, 0x8b, 0x4b, 0x73,
	0x4a, 0xa0, 0x7a, 0xa1, 0x3c, 0x4c, 0x5b, 0x8d, 0x32, 0xb8, 0x98, 0x1d, 0x53, 0x52, 0x84, 0xb4,
	0xb9, 0x98, 0x83, 0x4b, 0x73, 0x85, 0x04, 0xf5, 0xc0, 0x3e, 0xd4, 0x43, 0x78, 0x4f, 0x8a, 0x1f,
	0x59, 0xa8, 0x20, 0xa7, 0x52, 0x89, 0x41, 0xc8, 0x84, 0x8b, 0x0d, 0x62, 0x99, 0x90, 0x08, 0x54,
	0x12, 0xc5, 0xe1, 0x52, 0x42, 0x68, 0xa2, 0x60, 0x5d, 0x49, 0x6c, 0xe0, 0x30, 0x34, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xcf, 0x39, 0x7d, 0x3a, 0x54, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddClient is the client API for Add service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error)
	Concat(ctx context.Context, in *ConcatRequest, opts ...grpc.CallOption) (*ConcatReply, error)
}

type addClient struct {
	cc *grpc.ClientConn
}

func NewAddClient(cc *grpc.ClientConn) AddClient {
	return &addClient{cc}
}

func (c *addClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error) {
	out := new(SumReply)
	err := c.cc.Invoke(ctx, "/model.Add/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addClient) Concat(ctx context.Context, in *ConcatRequest, opts ...grpc.CallOption) (*ConcatReply, error) {
	out := new(ConcatReply)
	err := c.cc.Invoke(ctx, "/model.Add/Concat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServer is the server API for Add service.
type AddServer interface {
	Sum(context.Context, *SumRequest) (*SumReply, error)
	Concat(context.Context, *ConcatRequest) (*ConcatReply, error)
}

// UnimplementedAddServer can be embedded to have forward compatible implementations.
type UnimplementedAddServer struct {
}

func (*UnimplementedAddServer) Sum(ctx context.Context, req *SumRequest) (*SumReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedAddServer) Concat(ctx context.Context, req *ConcatRequest) (*ConcatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Concat not implemented")
}

func RegisterAddServer(s *grpc.Server, srv AddServer) {
	s.RegisterService(&_Add_serviceDesc, srv)
}

func _Add_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Add/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Add_Concat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConcatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServer).Concat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Add/Concat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServer).Concat(ctx, req.(*ConcatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Add_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Add",
	HandlerType: (*AddServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _Add_Sum_Handler,
		},
		{
			MethodName: "Concat",
			Handler:    _Add_Concat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sumFunc.proto",
}