// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testdataplane.proto

package testdataplane

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/ligato/networkservicemesh/pkg/nsm/apis/common"

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

type NSMPodType int32

const (
	NSMPodType_DEFAULT_POD NSMPodType = 0
	NSMPodType_NSMCLIENT   NSMPodType = 1
	NSMPodType_NSE         NSMPodType = 2
)

var NSMPodType_name = map[int32]string{
	0: "DEFAULT_POD",
	1: "NSMCLIENT",
	2: "NSE",
}
var NSMPodType_value = map[string]int32{
	"DEFAULT_POD": 0,
	"NSMCLIENT":   1,
	"NSE":         2,
}

func (x NSMPodType) String() string {
	return proto.EnumName(NSMPodType_name, int32(x))
}
func (NSMPodType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_testdataplane_70255cbe4608ea4b, []int{0}
}

type Pod struct {
	Metadata             *common.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	IpAddress            string           `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Pod) Reset()         { *m = Pod{} }
func (m *Pod) String() string { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()    {}
func (*Pod) Descriptor() ([]byte, []int) {
	return fileDescriptor_testdataplane_70255cbe4608ea4b, []int{0}
}
func (m *Pod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pod.Unmarshal(m, b)
}
func (m *Pod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pod.Marshal(b, m, deterministic)
}
func (dst *Pod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pod.Merge(dst, src)
}
func (m *Pod) XXX_Size() int {
	return xxx_messageInfo_Pod.Size(m)
}
func (m *Pod) XXX_DiscardUnknown() {
	xxx_messageInfo_Pod.DiscardUnknown(m)
}

var xxx_messageInfo_Pod proto.InternalMessageInfo

func (m *Pod) GetMetadata() *common.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Pod) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

type BuildConnectRequest struct {
	SourcePod            *Pod     `protobuf:"bytes,1,opt,name=source_pod,json=sourcePod,proto3" json:"source_pod,omitempty"`
	DestinationPod       *Pod     `protobuf:"bytes,2,opt,name=destination_pod,json=destinationPod,proto3" json:"destination_pod,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildConnectRequest) Reset()         { *m = BuildConnectRequest{} }
func (m *BuildConnectRequest) String() string { return proto.CompactTextString(m) }
func (*BuildConnectRequest) ProtoMessage()    {}
func (*BuildConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_testdataplane_70255cbe4608ea4b, []int{1}
}
func (m *BuildConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildConnectRequest.Unmarshal(m, b)
}
func (m *BuildConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildConnectRequest.Marshal(b, m, deterministic)
}
func (dst *BuildConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildConnectRequest.Merge(dst, src)
}
func (m *BuildConnectRequest) XXX_Size() int {
	return xxx_messageInfo_BuildConnectRequest.Size(m)
}
func (m *BuildConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuildConnectRequest proto.InternalMessageInfo

func (m *BuildConnectRequest) GetSourcePod() *Pod {
	if m != nil {
		return m.SourcePod
	}
	return nil
}

func (m *BuildConnectRequest) GetDestinationPod() *Pod {
	if m != nil {
		return m.DestinationPod
	}
	return nil
}

type BuildConnectReply struct {
	Built                bool     `protobuf:"varint,1,opt,name=built,proto3" json:"built,omitempty"`
	BuildError           string   `protobuf:"bytes,2,opt,name=build_error,json=buildError,proto3" json:"build_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildConnectReply) Reset()         { *m = BuildConnectReply{} }
func (m *BuildConnectReply) String() string { return proto.CompactTextString(m) }
func (*BuildConnectReply) ProtoMessage()    {}
func (*BuildConnectReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_testdataplane_70255cbe4608ea4b, []int{2}
}
func (m *BuildConnectReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildConnectReply.Unmarshal(m, b)
}
func (m *BuildConnectReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildConnectReply.Marshal(b, m, deterministic)
}
func (dst *BuildConnectReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildConnectReply.Merge(dst, src)
}
func (m *BuildConnectReply) XXX_Size() int {
	return xxx_messageInfo_BuildConnectReply.Size(m)
}
func (m *BuildConnectReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildConnectReply.DiscardUnknown(m)
}

var xxx_messageInfo_BuildConnectReply proto.InternalMessageInfo

func (m *BuildConnectReply) GetBuilt() bool {
	if m != nil {
		return m.Built
	}
	return false
}

func (m *BuildConnectReply) GetBuildError() string {
	if m != nil {
		return m.BuildError
	}
	return ""
}

type DeleteConnectRequest struct {
	Pod                  *Pod       `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	PodType              NSMPodType `protobuf:"varint,2,opt,name=pod_type,json=podType,proto3,enum=testdataplane.NSMPodType" json:"pod_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DeleteConnectRequest) Reset()         { *m = DeleteConnectRequest{} }
func (m *DeleteConnectRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteConnectRequest) ProtoMessage()    {}
func (*DeleteConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_testdataplane_70255cbe4608ea4b, []int{3}
}
func (m *DeleteConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteConnectRequest.Unmarshal(m, b)
}
func (m *DeleteConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteConnectRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteConnectRequest.Merge(dst, src)
}
func (m *DeleteConnectRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteConnectRequest.Size(m)
}
func (m *DeleteConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteConnectRequest proto.InternalMessageInfo

func (m *DeleteConnectRequest) GetPod() *Pod {
	if m != nil {
		return m.Pod
	}
	return nil
}

func (m *DeleteConnectRequest) GetPodType() NSMPodType {
	if m != nil {
		return m.PodType
	}
	return NSMPodType_DEFAULT_POD
}

type DeleteConnectReply struct {
	Deleted              bool     `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
	DeleteError          string   `protobuf:"bytes,2,opt,name=delete_error,json=deleteError,proto3" json:"delete_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteConnectReply) Reset()         { *m = DeleteConnectReply{} }
func (m *DeleteConnectReply) String() string { return proto.CompactTextString(m) }
func (*DeleteConnectReply) ProtoMessage()    {}
func (*DeleteConnectReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_testdataplane_70255cbe4608ea4b, []int{4}
}
func (m *DeleteConnectReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteConnectReply.Unmarshal(m, b)
}
func (m *DeleteConnectReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteConnectReply.Marshal(b, m, deterministic)
}
func (dst *DeleteConnectReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteConnectReply.Merge(dst, src)
}
func (m *DeleteConnectReply) XXX_Size() int {
	return xxx_messageInfo_DeleteConnectReply.Size(m)
}
func (m *DeleteConnectReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteConnectReply.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteConnectReply proto.InternalMessageInfo

func (m *DeleteConnectReply) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *DeleteConnectReply) GetDeleteError() string {
	if m != nil {
		return m.DeleteError
	}
	return ""
}

func init() {
	proto.RegisterType((*Pod)(nil), "testdataplane.Pod")
	proto.RegisterType((*BuildConnectRequest)(nil), "testdataplane.BuildConnectRequest")
	proto.RegisterType((*BuildConnectReply)(nil), "testdataplane.BuildConnectReply")
	proto.RegisterType((*DeleteConnectRequest)(nil), "testdataplane.DeleteConnectRequest")
	proto.RegisterType((*DeleteConnectReply)(nil), "testdataplane.DeleteConnectReply")
	proto.RegisterEnum("testdataplane.NSMPodType", NSMPodType_name, NSMPodType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BuildConnectClient is the client API for BuildConnect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuildConnectClient interface {
	RequestBuildConnect(ctx context.Context, in *BuildConnectRequest, opts ...grpc.CallOption) (*BuildConnectReply, error)
}

type buildConnectClient struct {
	cc *grpc.ClientConn
}

func NewBuildConnectClient(cc *grpc.ClientConn) BuildConnectClient {
	return &buildConnectClient{cc}
}

func (c *buildConnectClient) RequestBuildConnect(ctx context.Context, in *BuildConnectRequest, opts ...grpc.CallOption) (*BuildConnectReply, error) {
	out := new(BuildConnectReply)
	err := c.cc.Invoke(ctx, "/testdataplane.BuildConnect/RequestBuildConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildConnectServer is the server API for BuildConnect service.
type BuildConnectServer interface {
	RequestBuildConnect(context.Context, *BuildConnectRequest) (*BuildConnectReply, error)
}

func RegisterBuildConnectServer(s *grpc.Server, srv BuildConnectServer) {
	s.RegisterService(&_BuildConnect_serviceDesc, srv)
}

func _BuildConnect_RequestBuildConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildConnectServer).RequestBuildConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdataplane.BuildConnect/RequestBuildConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildConnectServer).RequestBuildConnect(ctx, req.(*BuildConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BuildConnect_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testdataplane.BuildConnect",
	HandlerType: (*BuildConnectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestBuildConnect",
			Handler:    _BuildConnect_RequestBuildConnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testdataplane.proto",
}

// DeleteConnectClient is the client API for DeleteConnect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeleteConnectClient interface {
	RequestDeleteConnect(ctx context.Context, in *DeleteConnectRequest, opts ...grpc.CallOption) (*DeleteConnectReply, error)
}

type deleteConnectClient struct {
	cc *grpc.ClientConn
}

func NewDeleteConnectClient(cc *grpc.ClientConn) DeleteConnectClient {
	return &deleteConnectClient{cc}
}

func (c *deleteConnectClient) RequestDeleteConnect(ctx context.Context, in *DeleteConnectRequest, opts ...grpc.CallOption) (*DeleteConnectReply, error) {
	out := new(DeleteConnectReply)
	err := c.cc.Invoke(ctx, "/testdataplane.DeleteConnect/RequestDeleteConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteConnectServer is the server API for DeleteConnect service.
type DeleteConnectServer interface {
	RequestDeleteConnect(context.Context, *DeleteConnectRequest) (*DeleteConnectReply, error)
}

func RegisterDeleteConnectServer(s *grpc.Server, srv DeleteConnectServer) {
	s.RegisterService(&_DeleteConnect_serviceDesc, srv)
}

func _DeleteConnect_RequestDeleteConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeleteConnectServer).RequestDeleteConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdataplane.DeleteConnect/RequestDeleteConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeleteConnectServer).RequestDeleteConnect(ctx, req.(*DeleteConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeleteConnect_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testdataplane.DeleteConnect",
	HandlerType: (*DeleteConnectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestDeleteConnect",
			Handler:    _DeleteConnect_RequestDeleteConnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testdataplane.proto",
}

func init() { proto.RegisterFile("testdataplane.proto", fileDescriptor_testdataplane_70255cbe4608ea4b) }

var fileDescriptor_testdataplane_70255cbe4608ea4b = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x5f, 0x8f, 0xd2, 0x40,
	0x14, 0xc5, 0x2d, 0x44, 0x81, 0xcb, 0xb2, 0x8b, 0x03, 0x0f, 0x48, 0x62, 0x64, 0xab, 0x0f, 0x1b,
	0x63, 0x68, 0x44, 0x7d, 0xf2, 0x69, 0x5d, 0x6a, 0xa2, 0x59, 0xb0, 0x16, 0x7c, 0x32, 0xb1, 0x29,
	0x9d, 0x1b, 0x76, 0x42, 0xdb, 0x3b, 0x76, 0xa6, 0x9a, 0x7e, 0x00, 0xbf, 0xb7, 0xe9, 0x1f, 0x75,
	0xdb, 0xac, 0x3c, 0xf5, 0xce, 0xe9, 0xb9, 0xf7, 0xcc, 0x6f, 0xda, 0x81, 0x91, 0x46, 0xa5, 0xb9,
	0xaf, 0x7d, 0x19, 0xfa, 0x31, 0xce, 0x65, 0x42, 0x9a, 0xd8, 0xa0, 0x26, 0x4e, 0xed, 0xbd, 0xd0,
	0x37, 0xe9, 0x6e, 0x1e, 0x50, 0x64, 0x85, 0x62, 0xef, 0x6b, 0xb2, 0x62, 0xd4, 0x3f, 0x29, 0x39,
	0x28, 0x4c, 0x7e, 0x88, 0x00, 0x23, 0x54, 0x37, 0x96, 0x3c, 0xec, 0xad, 0x58, 0x45, 0x96, 0x2f,
	0x85, 0xb2, 0x02, 0x8a, 0x22, 0x8a, 0xab, 0x47, 0x39, 0xd5, 0x74, 0xa1, 0xed, 0x10, 0x67, 0x2f,
	0xa0, 0x1b, 0xa1, 0xf6, 0xf3, 0xf1, 0x13, 0x63, 0x66, 0x5c, 0xf4, 0x17, 0xc3, 0x79, 0xe5, 0x5b,
	0x55, 0xba, 0xfb, 0xd7, 0xc1, 0x1e, 0x03, 0x08, 0xe9, 0xf9, 0x9c, 0x27, 0xa8, 0xd4, 0xa4, 0x35,
	0x33, 0x2e, 0x7a, 0x6e, 0x4f, 0xc8, 0xcb, 0x52, 0x30, 0x7f, 0x19, 0x30, 0x7a, 0x97, 0x8a, 0x90,
	0x5f, 0x51, 0x1c, 0x63, 0xa0, 0x5d, 0xfc, 0x9e, 0xa2, 0xd2, 0xec, 0x25, 0x80, 0xa2, 0x34, 0x09,
	0xd0, 0x93, 0xc4, 0xab, 0x18, 0x36, 0xaf, 0xb3, 0x3a, 0xc4, 0xdd, 0x5e, 0xe9, 0xca, 0xf7, 0xf5,
	0x16, 0xce, 0x38, 0x2a, 0x2d, 0x62, 0x5f, 0x0b, 0x8a, 0x8b, 0xbe, 0xd6, 0x7f, 0xfb, 0x4e, 0x6f,
	0x59, 0x1d, 0xe2, 0xe6, 0x47, 0x78, 0x58, 0xdf, 0x86, 0x0c, 0x33, 0x36, 0x86, 0xfb, 0xbb, 0x54,
	0x84, 0xba, 0xc8, 0xef, 0xba, 0xe5, 0x82, 0x3d, 0x81, 0x7e, 0x5e, 0x70, 0x0f, 0x93, 0x84, 0x92,
	0x0a, 0x09, 0x0a, 0xc9, 0xce, 0x15, 0x33, 0x81, 0xf1, 0x12, 0x43, 0xd4, 0xd8, 0x60, 0x7a, 0x06,
	0xed, 0xe3, 0x30, 0xf9, 0x6b, 0xf6, 0x1a, 0xba, 0x92, 0xb8, 0xa7, 0x33, 0x89, 0xc5, 0xec, 0xd3,
	0xc5, 0xa3, 0x86, 0x75, 0xbd, 0x59, 0x39, 0xc4, 0xb7, 0x99, 0x44, 0xb7, 0x23, 0xcb, 0xc2, 0xfc,
	0x0c, 0xac, 0x91, 0x99, 0x03, 0x4c, 0xa0, 0xc3, 0x0b, 0x95, 0x57, 0x08, 0x7f, 0x96, 0xec, 0x1c,
	0x4e, 0xca, 0xb2, 0x46, 0xd1, 0x2f, 0xb5, 0x02, 0xe3, 0xf9, 0x1b, 0x80, 0x7f, 0x49, 0xec, 0x0c,
	0xfa, 0x4b, 0xfb, 0xfd, 0xe5, 0x97, 0xeb, 0xad, 0xe7, 0x7c, 0x5a, 0x0e, 0xef, 0xb1, 0x01, 0xf4,
	0xd6, 0x9b, 0xd5, 0xd5, 0xf5, 0x07, 0x7b, 0xbd, 0x1d, 0x1a, 0xac, 0x03, 0xed, 0xf5, 0xc6, 0x1e,
	0xb6, 0x16, 0x07, 0x38, 0xb9, 0x7d, 0x92, 0xec, 0x2b, 0x8c, 0xaa, 0x03, 0xa8, 0xc9, 0x66, 0x03,
	0xea, 0x8e, 0x9f, 0x60, 0x3a, 0x3b, 0xea, 0x91, 0x61, 0xb6, 0x20, 0x18, 0xd4, 0xb0, 0xd9, 0x37,
	0x18, 0x57, 0xdd, 0x75, 0xfd, 0x69, 0x63, 0xd4, 0x5d, 0x1f, 0x68, 0x7a, 0x7e, 0xdc, 0x24, 0xc3,
	0x6c, 0xf7, 0xa0, 0xb8, 0x0a, 0xaf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x20, 0x2d, 0x43, 0xb0,
	0x77, 0x03, 0x00, 0x00,
}
