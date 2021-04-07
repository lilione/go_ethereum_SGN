// Code generated by protoc-gen-go. DO NOT EDIT.
// source: verify.proto

package verify

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

type Request struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87721bac73e05a3, []int{0}
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

func (m *Request) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Response struct {
	Valid                bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87721bac73e05a3, []int{1}
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

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*Request)(nil), "verify.Request")
	proto.RegisterType((*Response)(nil), "verify.Response")
}

func init() { proto.RegisterFile("verify.proto", fileDescriptor_a87721bac73e05a3) }

var fileDescriptor_a87721bac73e05a3 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4b, 0x2d, 0xca,
	0x4c, 0xab, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x94, 0xb9, 0xd8,
	0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x0b, 0x12, 0x2b, 0x73, 0xf2,
	0x13, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0x60, 0x5c, 0x25, 0x05, 0x2e, 0x8e, 0xa0,
	0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0x4c, 0x88,
	0x1a, 0x8e, 0x20, 0x08, 0xc7, 0xc8, 0x9a, 0x8b, 0x23, 0x0c, 0x6c, 0x60, 0x48, 0x85, 0x90, 0x3e,
	0x17, 0x87, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0x62, 0x49, 0xaa, 0x10, 0xbf, 0x1e, 0xd4, 0x56, 0xa8,
	0x25, 0x52, 0x02, 0x08, 0x01, 0x88, 0x81, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x27, 0x19, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x28, 0xf5, 0x40, 0xdc, 0xa2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VerifyTxClient is the client API for VerifyTx service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VerifyTxClient interface {
	Evaluate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type verifyTxClient struct {
	cc grpc.ClientConnInterface
}

func NewVerifyTxClient(cc grpc.ClientConnInterface) VerifyTxClient {
	return &verifyTxClient{cc}
}

func (c *verifyTxClient) Evaluate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/verify.VerifyTx/Evaluate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerifyTxServer is the server API for VerifyTx service.
type VerifyTxServer interface {
	Evaluate(context.Context, *Request) (*Response, error)
}

// UnimplementedVerifyTxServer can be embedded to have forward compatible implementations.
type UnimplementedVerifyTxServer struct {
}

func (*UnimplementedVerifyTxServer) Evaluate(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Evaluate not implemented")
}

func RegisterVerifyTxServer(s *grpc.Server, srv VerifyTxServer) {
	s.RegisterService(&_VerifyTx_serviceDesc, srv)
}

func _VerifyTx_Evaluate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyTxServer).Evaluate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verify.VerifyTx/Evaluate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyTxServer).Evaluate(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _VerifyTx_serviceDesc = grpc.ServiceDesc{
	ServiceName: "verify.VerifyTx",
	HandlerType: (*VerifyTxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Evaluate",
			Handler:    _VerifyTx_Evaluate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "verify.proto",
}
