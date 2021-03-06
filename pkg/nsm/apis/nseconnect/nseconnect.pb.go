// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nseconnect.proto

package nseconnect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/ligato/networkservicemesh/pkg/nsm/apis/common"
import netmesh "github.com/ligato/networkservicemesh/pkg/nsm/apis/netmesh"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ConnectionRequest is sent by a NSM to NSE to build a connection.
type NSEConnectionRequest struct {
	RequestId            string                         `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Metadata             *common.Metadata               `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Channel              *netmesh.NetworkServiceChannel `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *NSEConnectionRequest) Reset()         { *m = NSEConnectionRequest{} }
func (m *NSEConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*NSEConnectionRequest) ProtoMessage()    {}
func (*NSEConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_nseconnect_e14735d60af7b65d, []int{0}
}
func (m *NSEConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NSEConnectionRequest.Unmarshal(m, b)
}
func (m *NSEConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NSEConnectionRequest.Marshal(b, m, deterministic)
}
func (dst *NSEConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NSEConnectionRequest.Merge(dst, src)
}
func (m *NSEConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_NSEConnectionRequest.Size(m)
}
func (m *NSEConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NSEConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NSEConnectionRequest proto.InternalMessageInfo

func (m *NSEConnectionRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *NSEConnectionRequest) GetMetadata() *common.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *NSEConnectionRequest) GetChannel() *netmesh.NetworkServiceChannel {
	if m != nil {
		return m.Channel
	}
	return nil
}

// NSEConnectionReply is sent back by NSE to NSM with information required for
// dataplane programming.
type NSEConnectionReply struct {
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Metadata contains NSE's POD name and namespace which will be needed for
	// dataplane programming
	Metadata             *common.Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	NetworkServiceName   string           `protobuf:"bytes,3,opt,name=network_service_name,json=networkServiceName,proto3" json:"network_service_name,omitempty"`
	LinuxNamespace       string           `protobuf:"bytes,4,opt,name=linux_namespace,json=linuxNamespace,proto3" json:"linux_namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *NSEConnectionReply) Reset()         { *m = NSEConnectionReply{} }
func (m *NSEConnectionReply) String() string { return proto.CompactTextString(m) }
func (*NSEConnectionReply) ProtoMessage()    {}
func (*NSEConnectionReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_nseconnect_e14735d60af7b65d, []int{1}
}
func (m *NSEConnectionReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NSEConnectionReply.Unmarshal(m, b)
}
func (m *NSEConnectionReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NSEConnectionReply.Marshal(b, m, deterministic)
}
func (dst *NSEConnectionReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NSEConnectionReply.Merge(dst, src)
}
func (m *NSEConnectionReply) XXX_Size() int {
	return xxx_messageInfo_NSEConnectionReply.Size(m)
}
func (m *NSEConnectionReply) XXX_DiscardUnknown() {
	xxx_messageInfo_NSEConnectionReply.DiscardUnknown(m)
}

var xxx_messageInfo_NSEConnectionReply proto.InternalMessageInfo

func (m *NSEConnectionReply) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *NSEConnectionReply) GetMetadata() *common.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *NSEConnectionReply) GetNetworkServiceName() string {
	if m != nil {
		return m.NetworkServiceName
	}
	return ""
}

func (m *NSEConnectionReply) GetLinuxNamespace() string {
	if m != nil {
		return m.LinuxNamespace
	}
	return ""
}

func init() {
	proto.RegisterType((*NSEConnectionRequest)(nil), "nseconnect.NSEConnectionRequest")
	proto.RegisterType((*NSEConnectionReply)(nil), "nseconnect.NSEConnectionReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NSEConnectionClient is the client API for NSEConnection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NSEConnectionClient interface {
	RequestNSEConnection(ctx context.Context, in *NSEConnectionRequest, opts ...grpc.CallOption) (*NSEConnectionReply, error)
}

type nSEConnectionClient struct {
	cc *grpc.ClientConn
}

func NewNSEConnectionClient(cc *grpc.ClientConn) NSEConnectionClient {
	return &nSEConnectionClient{cc}
}

func (c *nSEConnectionClient) RequestNSEConnection(ctx context.Context, in *NSEConnectionRequest, opts ...grpc.CallOption) (*NSEConnectionReply, error) {
	out := new(NSEConnectionReply)
	err := c.cc.Invoke(ctx, "/nseconnect.NSEConnection/RequestNSEConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NSEConnectionServer is the server API for NSEConnection service.
type NSEConnectionServer interface {
	RequestNSEConnection(context.Context, *NSEConnectionRequest) (*NSEConnectionReply, error)
}

func RegisterNSEConnectionServer(s *grpc.Server, srv NSEConnectionServer) {
	s.RegisterService(&_NSEConnection_serviceDesc, srv)
}

func _NSEConnection_RequestNSEConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NSEConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSEConnectionServer).RequestNSEConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nseconnect.NSEConnection/RequestNSEConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSEConnectionServer).RequestNSEConnection(ctx, req.(*NSEConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NSEConnection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nseconnect.NSEConnection",
	HandlerType: (*NSEConnectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestNSEConnection",
			Handler:    _NSEConnection_RequestNSEConnection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nseconnect.proto",
}

func init() { proto.RegisterFile("nseconnect.proto", fileDescriptor_nseconnect_e14735d60af7b65d) }

var fileDescriptor_nseconnect_e14735d60af7b65d = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0x8a, 0xda, 0x11, 0xb5, 0x2c, 0x3d, 0x84, 0x82, 0xa5, 0xf4, 0x62, 0x0f, 0x92,
	0x48, 0xbd, 0x78, 0x2f, 0x45, 0x3c, 0x98, 0x43, 0x7a, 0xf1, 0x56, 0xb6, 0xdb, 0xa1, 0x5d, 0x9a,
	0x9d, 0x5d, 0xb3, 0x5b, 0xb5, 0x0f, 0xe3, 0x9b, 0xf8, 0x70, 0xe2, 0xee, 0x56, 0xad, 0x88, 0x20,
	0x78, 0x9a, 0xf0, 0xe7, 0xe3, 0xe7, 0xcb, 0x4c, 0xa0, 0x49, 0x16, 0x85, 0x26, 0x42, 0xe1, 0x32,
	0x53, 0x6b, 0xa7, 0x19, 0x7c, 0x26, 0xed, 0x9b, 0xb9, 0x74, 0x8b, 0xd5, 0x34, 0x13, 0x5a, 0xe5,
	0x95, 0x9c, 0x73, 0xa7, 0x73, 0x42, 0xf7, 0xa4, 0xeb, 0xa5, 0xc5, 0xfa, 0x51, 0x0a, 0x54, 0x68,
	0x17, 0xb9, 0x59, 0xce, 0x73, 0xb2, 0x2a, 0xe7, 0x46, 0xda, 0xf7, 0xf7, 0x3e, 0x8c, 0x33, 0x94,
	0xb6, 0x47, 0x7f, 0x2f, 0x12, 0x5a, 0x29, 0x4d, 0x71, 0x84, 0x9a, 0xde, 0x4b, 0x02, 0xad, 0x62,
	0x3c, 0x1a, 0x06, 0x3d, 0xa9, 0xa9, 0xc4, 0x87, 0x15, 0x5a, 0xc7, 0xce, 0x00, 0xea, 0xf0, 0x38,
	0x91, 0xb3, 0x34, 0xe9, 0x26, 0xfd, 0x46, 0xd9, 0x88, 0xc9, 0xed, 0x8c, 0x5d, 0xc0, 0xa1, 0x42,
	0xc7, 0x67, 0xdc, 0xf1, 0x74, 0xa7, 0x9b, 0xf4, 0x8f, 0x06, 0xcd, 0x2c, 0x16, 0xdf, 0xc5, 0xbc,
	0xfc, 0x20, 0xd8, 0x35, 0x1c, 0x88, 0x05, 0x27, 0xc2, 0x2a, 0xdd, 0xf5, 0x70, 0x27, 0xdb, 0x7c,
	0x4d, 0x11, 0xa4, 0xc7, 0x41, 0x7a, 0x18, 0xa8, 0x72, 0x83, 0xf7, 0x5e, 0x13, 0x60, 0xdf, 0xfc,
	0x4c, 0xb5, 0xfe, 0x5f, 0xbb, 0x4b, 0x68, 0xc5, 0xd5, 0x4d, 0xe2, 0xee, 0x26, 0xc4, 0x15, 0x7a,
	0xd5, 0x46, 0xc9, 0x68, 0xcb, 0xb0, 0xe0, 0x0a, 0xd9, 0x39, 0x9c, 0x56, 0x92, 0x56, 0xcf, 0x9e,
	0xb3, 0x86, 0x0b, 0x4c, 0xf7, 0x3c, 0x7c, 0xe2, 0xe3, 0x62, 0x93, 0x0e, 0x24, 0x1c, 0x6f, 0xd9,
	0xb3, 0x7b, 0x68, 0xc5, 0x0d, 0x6f, 0xe7, 0xdd, 0xec, 0xcb, 0x6f, 0xf3, 0xd3, 0x41, 0xda, 0x9d,
	0x5f, 0x08, 0x53, 0xad, 0xa7, 0xfb, 0xfe, 0xa0, 0x57, 0x6f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd0,
	0x63, 0x31, 0xa2, 0x80, 0x02, 0x00, 0x00,
}
