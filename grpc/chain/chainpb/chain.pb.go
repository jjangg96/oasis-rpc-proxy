// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/chain/chainpb/chain.proto

package chainpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Chain struct {
	Id                   string               `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	GenesisTime          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=GenesisTime,proto3" json:"GenesisTime,omitempty"`
	Height               int64                `protobuf:"varint,4,opt,name=Height,proto3" json:"Height,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Chain) Reset()         { *m = Chain{} }
func (m *Chain) String() string { return proto.CompactTextString(m) }
func (*Chain) ProtoMessage()    {}
func (*Chain) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce694d07bea9bb01, []int{0}
}

func (m *Chain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chain.Unmarshal(m, b)
}
func (m *Chain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chain.Marshal(b, m, deterministic)
}
func (m *Chain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chain.Merge(m, src)
}
func (m *Chain) XXX_Size() int {
	return xxx_messageInfo_Chain.Size(m)
}
func (m *Chain) XXX_DiscardUnknown() {
	xxx_messageInfo_Chain.DiscardUnknown(m)
}

var xxx_messageInfo_Chain proto.InternalMessageInfo

func (m *Chain) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Chain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Chain) GetGenesisTime() *timestamp.Timestamp {
	if m != nil {
		return m.GenesisTime
	}
	return nil
}

func (m *Chain) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type GetCurrentRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentRequest) Reset()         { *m = GetCurrentRequest{} }
func (m *GetCurrentRequest) String() string { return proto.CompactTextString(m) }
func (*GetCurrentRequest) ProtoMessage()    {}
func (*GetCurrentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce694d07bea9bb01, []int{1}
}

func (m *GetCurrentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentRequest.Unmarshal(m, b)
}
func (m *GetCurrentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentRequest.Marshal(b, m, deterministic)
}
func (m *GetCurrentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentRequest.Merge(m, src)
}
func (m *GetCurrentRequest) XXX_Size() int {
	return xxx_messageInfo_GetCurrentRequest.Size(m)
}
func (m *GetCurrentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentRequest proto.InternalMessageInfo

type GetCurrentResponse struct {
	Chain                *Chain   `protobuf:"bytes,1,opt,name=Chain,proto3" json:"Chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentResponse) Reset()         { *m = GetCurrentResponse{} }
func (m *GetCurrentResponse) String() string { return proto.CompactTextString(m) }
func (*GetCurrentResponse) ProtoMessage()    {}
func (*GetCurrentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce694d07bea9bb01, []int{2}
}

func (m *GetCurrentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentResponse.Unmarshal(m, b)
}
func (m *GetCurrentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentResponse.Marshal(b, m, deterministic)
}
func (m *GetCurrentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentResponse.Merge(m, src)
}
func (m *GetCurrentResponse) XXX_Size() int {
	return xxx_messageInfo_GetCurrentResponse.Size(m)
}
func (m *GetCurrentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentResponse proto.InternalMessageInfo

func (m *GetCurrentResponse) GetChain() *Chain {
	if m != nil {
		return m.Chain
	}
	return nil
}

func init() {
	proto.RegisterType((*Chain)(nil), "chain.Chain")
	proto.RegisterType((*GetCurrentRequest)(nil), "chain.GetCurrentRequest")
	proto.RegisterType((*GetCurrentResponse)(nil), "chain.GetCurrentResponse")
}

func init() { proto.RegisterFile("grpc/chain/chainpb/chain.proto", fileDescriptor_ce694d07bea9bb01) }

var fileDescriptor_ce694d07bea9bb01 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x49, 0xff, 0xa1, 0x5e, 0x2a, 0x24, 0x0e, 0x09, 0x85, 0x0c, 0x10, 0x65, 0xca, 0xe4,
	0x48, 0x61, 0x61, 0x60, 0x22, 0x43, 0xe9, 0xc2, 0xe0, 0x76, 0x62, 0x4b, 0xd2, 0x23, 0xb5, 0x44,
	0x62, 0x63, 0x3b, 0xec, 0x7c, 0x73, 0x84, 0x9d, 0x8a, 0x4a, 0xb0, 0xd8, 0x77, 0xef, 0x9e, 0xfd,
	0xbb, 0x3b, 0xb8, 0x6d, 0xb5, 0x6a, 0xf2, 0xe6, 0x50, 0x89, 0xde, 0x9f, 0xaa, 0xf6, 0x37, 0x53,
	0x5a, 0x5a, 0x89, 0x73, 0x97, 0xc4, 0x77, 0xad, 0x94, 0xed, 0x3b, 0xe5, 0x4e, 0xac, 0x87, 0xb7,
	0xdc, 0x8a, 0x8e, 0x8c, 0xad, 0x3a, 0xe5, 0x7d, 0xe9, 0x57, 0x00, 0xf3, 0xf2, 0xc7, 0x8a, 0x17,
	0x30, 0xd9, 0xec, 0xa3, 0x20, 0x09, 0xb2, 0x25, 0x9f, 0x6c, 0xf6, 0x88, 0x30, 0x7b, 0xa9, 0x3a,
	0x8a, 0x26, 0x4e, 0x71, 0x31, 0x3e, 0x42, 0xb8, 0xa6, 0x9e, 0x8c, 0x30, 0x3b, 0xd1, 0x51, 0x34,
	0x4d, 0x82, 0x2c, 0x2c, 0x62, 0xe6, 0x21, 0xec, 0x08, 0x61, 0xbb, 0x23, 0x84, 0x9f, 0xda, 0xf1,
	0x1a, 0x16, 0xcf, 0x24, 0xda, 0x83, 0x8d, 0x66, 0x49, 0x90, 0x4d, 0xf9, 0x98, 0xa5, 0x57, 0x70,
	0xb9, 0x26, 0x5b, 0x0e, 0x5a, 0x53, 0x6f, 0x39, 0x7d, 0x0c, 0x64, 0x6c, 0xfa, 0x00, 0x78, 0x2a,
	0x1a, 0x25, 0x7b, 0x43, 0x98, 0x8e, 0xdd, 0xba, 0x3e, 0xc3, 0x62, 0xc5, 0xfc, 0xcc, 0x4e, 0xe3,
	0xbe, 0x54, 0x6c, 0x61, 0xe5, 0x82, 0x2d, 0xe9, 0x4f, 0xd1, 0x10, 0x96, 0x00, 0xbf, 0x3f, 0x61,
	0x34, 0x3e, 0xf9, 0x43, 0x8c, 0x6f, 0xfe, 0xa9, 0x78, 0x6c, 0x7a, 0xf6, 0xb4, 0x7c, 0x3d, 0x1f,
	0xd7, 0x5c, 0x2f, 0xdc, 0x9c, 0xf7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x39, 0x4c, 0x2b,
	0x83, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChainServiceClient is the client API for ChainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChainServiceClient interface {
	GetCurrent(ctx context.Context, in *GetCurrentRequest, opts ...grpc.CallOption) (*GetCurrentResponse, error)
}

type chainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChainServiceClient(cc grpc.ClientConnInterface) ChainServiceClient {
	return &chainServiceClient{cc}
}

func (c *chainServiceClient) GetCurrent(ctx context.Context, in *GetCurrentRequest, opts ...grpc.CallOption) (*GetCurrentResponse, error) {
	out := new(GetCurrentResponse)
	err := c.cc.Invoke(ctx, "/chain.ChainService/GetCurrent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChainServiceServer is the server API for ChainService service.
type ChainServiceServer interface {
	GetCurrent(context.Context, *GetCurrentRequest) (*GetCurrentResponse, error)
}

// UnimplementedChainServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChainServiceServer struct {
}

func (*UnimplementedChainServiceServer) GetCurrent(ctx context.Context, req *GetCurrentRequest) (*GetCurrentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrent not implemented")
}

func RegisterChainServiceServer(s *grpc.Server, srv ChainServiceServer) {
	s.RegisterService(&_ChainService_serviceDesc, srv)
}

func _ChainService_GetCurrent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).GetCurrent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.ChainService/GetCurrent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).GetCurrent(ctx, req.(*GetCurrentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chain.ChainService",
	HandlerType: (*ChainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrent",
			Handler:    _ChainService_GetCurrent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/chain/chainpb/chain.proto",
}