// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: verify.proto

package verify

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConAddr []byte `protobuf:"bytes,1,opt,name=conAddr,proto3" json:"conAddr,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_verify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_verify_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetConAddr() []byte {
	if x != nil {
		return x.ConAddr
	}
	return nil
}

func (x *Request) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_verify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_verify_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_verify_proto protoreflect.FileDescriptor

var file_verify_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x22, 0x3d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x32, 0x3b, 0x0a, 0x08, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x54, 0x78, 0x12, 0x2f, 0x0a, 0x08, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x12,
	0x0f, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6c, 0x69, 0x6f, 0x6e, 0x65, 0x2f, 0x50, 0x61, 0x72, 0x62, 0x6c,
	0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_verify_proto_rawDescOnce sync.Once
	file_verify_proto_rawDescData = file_verify_proto_rawDesc
)

func file_verify_proto_rawDescGZIP() []byte {
	file_verify_proto_rawDescOnce.Do(func() {
		file_verify_proto_rawDescData = protoimpl.X.CompressGZIP(file_verify_proto_rawDescData)
	})
	return file_verify_proto_rawDescData
}

var file_verify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_verify_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: verify.Request
	(*Response)(nil), // 1: verify.Response
}
var file_verify_proto_depIdxs = []int32{
	0, // 0: verify.VerifyTx.Evaluate:input_type -> verify.Request
	1, // 1: verify.VerifyTx.Evaluate:output_type -> verify.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_verify_proto_init() }
func file_verify_proto_init() {
	if File_verify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_verify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_verify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_verify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_verify_proto_goTypes,
		DependencyIndexes: file_verify_proto_depIdxs,
		MessageInfos:      file_verify_proto_msgTypes,
	}.Build()
	File_verify_proto = out.File
	file_verify_proto_rawDesc = nil
	file_verify_proto_goTypes = nil
	file_verify_proto_depIdxs = nil
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

func (*UnimplementedVerifyTxServer) Evaluate(context.Context, *Request) (*Response, error) {
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
