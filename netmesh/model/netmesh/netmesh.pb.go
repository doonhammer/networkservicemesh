// Code generated by protoc-gen-go. DO NOT EDIT.
// source: netmesh.proto

/*
Package netmesh is a generated protocol buffer package.

It is generated from these files:
	netmesh.proto

It has these top-level messages:
	NetworkServiceEndpoint
	NetworkService
*/
package netmesh

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

type NetworkServiceEndpoint struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Uuid string `protobuf:"bytes,2,opt,name=uuid" json:"uuid,omitempty"`
}

func (m *NetworkServiceEndpoint) Reset()                    { *m = NetworkServiceEndpoint{} }
func (m *NetworkServiceEndpoint) String() string            { return proto.CompactTextString(m) }
func (*NetworkServiceEndpoint) ProtoMessage()               {}
func (*NetworkServiceEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NetworkServiceEndpoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkServiceEndpoint) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type NetworkService struct {
	Name     string                           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Uuid     string                           `protobuf:"bytes,2,opt,name=uuid" json:"uuid,omitempty"`
	Selector string                           `protobuf:"bytes,3,opt,name=selector" json:"selector,omitempty"`
	Channels []*NetworkService_NetmeshChannel `protobuf:"bytes,4,rep,name=channels" json:"channels,omitempty"`
}

func (m *NetworkService) Reset()                    { *m = NetworkService{} }
func (m *NetworkService) String() string            { return proto.CompactTextString(m) }
func (*NetworkService) ProtoMessage()               {}
func (*NetworkService) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NetworkService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkService) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *NetworkService) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *NetworkService) GetChannels() []*NetworkService_NetmeshChannel {
	if m != nil {
		return m.Channels
	}
	return nil
}

type NetworkService_NetmeshChannel struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
}

func (m *NetworkService_NetmeshChannel) Reset()         { *m = NetworkService_NetmeshChannel{} }
func (m *NetworkService_NetmeshChannel) String() string { return proto.CompactTextString(m) }
func (*NetworkService_NetmeshChannel) ProtoMessage()    {}
func (*NetworkService_NetmeshChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

func (m *NetworkService_NetmeshChannel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkService_NetmeshChannel) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func init() {
	proto.RegisterType((*NetworkServiceEndpoint)(nil), "netmesh.NetworkServiceEndpoint")
	proto.RegisterType((*NetworkService)(nil), "netmesh.NetworkService")
	proto.RegisterType((*NetworkService_NetmeshChannel)(nil), "netmesh.NetworkService.NetmeshChannel")
}

func init() { proto.RegisterFile("netmesh.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4b, 0x2d, 0xc9,
	0x4d, 0x2d, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x1c, 0xb8,
	0xc4, 0xfc, 0x52, 0x4b, 0xca, 0xf3, 0x8b, 0xb2, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x5d,
	0xf3, 0x52, 0x0a, 0xf2, 0x33, 0xf3, 0x4a, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x90, 0x58, 0x69, 0x69, 0x66, 0x8a, 0x04, 0x13,
	0x44, 0x0c, 0xc4, 0x56, 0xba, 0xc1, 0xc8, 0xc5, 0x87, 0x6a, 0x04, 0xb1, 0x5a, 0x85, 0xa4, 0xb8,
	0x38, 0x8a, 0x53, 0x73, 0x52, 0x93, 0x4b, 0xf2, 0x8b, 0x24, 0x98, 0xc1, 0xe2, 0x70, 0xbe, 0x90,
	0x13, 0x17, 0x47, 0x72, 0x46, 0x62, 0x5e, 0x5e, 0x6a, 0x4e, 0xb1, 0x04, 0x8b, 0x02, 0xb3, 0x06,
	0xb7, 0x91, 0x9a, 0x1e, 0xcc, 0x0f, 0xa8, 0xd6, 0x81, 0xb8, 0x20, 0x61, 0x67, 0x88, 0xf2, 0x20,
	0xb8, 0x3e, 0x29, 0x3b, 0xb0, 0xcb, 0x90, 0xe4, 0xb0, 0xba, 0x4c, 0x82, 0x8b, 0xbd, 0x20, 0xb1,
	0x32, 0x27, 0x3f, 0x11, 0xe6, 0x38, 0x18, 0x37, 0x89, 0x0d, 0x1c, 0x58, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x0f, 0x2a, 0xe9, 0xb2, 0x3d, 0x01, 0x00, 0x00,
}
