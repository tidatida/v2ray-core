// Code generated by protoc-gen-go.
// source: v2ray.com/core/proxy/vmess/outbound/config.proto
// DO NOT EDIT!

/*
Package outbound is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/proxy/vmess/outbound/config.proto

It has these top-level messages:
	Config
*/
package outbound

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_protocol1 "v2ray.com/core/common/protocol"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	Receiver []*v2ray_core_common_protocol1.ServerSpecPB `protobuf:"bytes,1,rep,name=Receiver,json=receiver" json:"Receiver,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Config) GetReceiver() []*v2ray_core_common_protocol1.ServerSpecPB {
	if m != nil {
		return m.Receiver
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "v2ray.core.proxy.vmess.outbound.Config")
}

func init() { proto.RegisterFile("v2ray.com/core/proxy/vmess/outbound/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0xce, 0x31, 0x6b, 0xc3, 0x30,
	0x10, 0x05, 0x60, 0x4a, 0xc1, 0x18, 0x79, 0xf3, 0x54, 0xba, 0xb4, 0xb4, 0x8b, 0xa7, 0x53, 0x71,
	0xd7, 0x4e, 0x6e, 0xe6, 0x60, 0xec, 0x2d, 0x4b, 0x88, 0x2f, 0x97, 0x60, 0x88, 0x7c, 0xe2, 0x64,
	0x8b, 0xf8, 0xdf, 0x87, 0xc8, 0x11, 0x84, 0x2c, 0x59, 0x8f, 0xf7, 0xbe, 0x77, 0xea, 0xc7, 0x97,
	0xb2, 0x9b, 0x01, 0xd9, 0x68, 0x64, 0x21, 0x6d, 0x85, 0xcf, 0xb3, 0xf6, 0x86, 0x9c, 0xd3, 0x3c,
	0x8d, 0x1d, 0x4f, 0xc3, 0x5e, 0x23, 0x0f, 0x87, 0xfe, 0x08, 0x56, 0x78, 0xe4, 0xfc, 0x23, 0x36,
	0x84, 0x20, 0xa4, 0x21, 0xa4, 0x21, 0xa6, 0xdf, 0x1f, 0x49, 0x64, 0x63, 0x78, 0xd0, 0xa1, 0x8d,
	0x7c, 0xd2, 0x8e, 0xc4, 0x93, 0x6c, 0x9d, 0x25, 0x5c, 0xc8, 0xaf, 0xb5, 0x4a, 0xfe, 0xc3, 0x44,
	0xbe, 0x52, 0x69, 0x43, 0x48, 0xbd, 0x27, 0x79, 0x7b, 0xf9, 0x7c, 0x2d, 0xb2, 0xb2, 0x80, 0xbb,
	0xbd, 0x85, 0x82, 0x48, 0x41, 0x1b, 0xa8, 0xd6, 0x12, 0xd6, 0x55, 0x93, 0xca, 0xad, 0x59, 0xfd,
	0xa9, 0x6f, 0x64, 0x03, 0x4f, 0x1e, 0xad, 0xb2, 0x65, 0xb4, 0xbe, 0x6a, 0x9b, 0x34, 0x9e, 0xbb,
	0x24, 0xf0, 0xbf, 0x97, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x86, 0x34, 0xb5, 0x1b, 0x01, 0x00,
	0x00,
}
