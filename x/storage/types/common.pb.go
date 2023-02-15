// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/storage/common.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SourceType int32

const (
	SOURCE_TYPE_ORIGIN          SourceType = 0
	SOURCE_TYPE_BSC_CROSS_CHAIN SourceType = 1
)

var SourceType_name = map[int32]string{
	0: "SOURCE_TYPE_ORIGIN",
	1: "SOURCE_TYPE_BSC_CROSS_CHAIN",
}

var SourceType_value = map[string]int32{
	"SOURCE_TYPE_ORIGIN":          0,
	"SOURCE_TYPE_BSC_CROSS_CHAIN": 1,
}

func (x SourceType) String() string {
	return proto.EnumName(SourceType_name, int32(x))
}

func (SourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{0}
}

// TODO: Need to confirm the read quota enum which determined by payment module
// TODO: Make this field be configured through governance
type ReadQuota int32

const (
	READ_QUOTA_FREE ReadQuota = 0
	READ_QUOTA_1G   ReadQuota = 1
	READ_QUOTA_10G  ReadQuota = 2
)

var ReadQuota_name = map[int32]string{
	0: "READ_QUOTA_FREE",
	1: "READ_QUOTA_1G",
	2: "READ_QUOTA_10G",
}

var ReadQuota_value = map[string]int32{
	"READ_QUOTA_FREE": 0,
	"READ_QUOTA_1G":   1,
	"READ_QUOTA_10G":  2,
}

func (x ReadQuota) String() string {
	return proto.EnumName(ReadQuota_name, int32(x))
}

func (ReadQuota) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{1}
}

func init() {
	proto.RegisterEnum("bnbchain.greenfield.storage.SourceType", SourceType_name, SourceType_value)
	proto.RegisterEnum("bnbchain.greenfield.storage.ReadQuota", ReadQuota_name, ReadQuota_value)
}

func init() { proto.RegisterFile("greenfield/storage/common.proto", fileDescriptor_4eff6c0fa4aaf4c9) }

var fileDescriptor_4eff6c0fa4aaf4c9 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4d, 0x4a, 0xc3, 0x40,
	0x14, 0x80, 0x27, 0x22, 0x82, 0x03, 0x6a, 0x1c, 0x45, 0xb0, 0x85, 0xe9, 0xbe, 0x60, 0x46, 0xf1,
	0x04, 0x6d, 0x8c, 0x31, 0x88, 0x8d, 0x9d, 0xa4, 0x0b, 0xdd, 0x0c, 0x99, 0x74, 0x4c, 0x03, 0x26,
	0x2f, 0xe4, 0x07, 0xec, 0x0d, 0x5c, 0x7a, 0x07, 0x2f, 0xe3, 0xb2, 0x4b, 0x97, 0x92, 0x5c, 0x44,
	0x4c, 0x44, 0xb3, 0x9b, 0xf9, 0xf8, 0xde, 0x83, 0xf7, 0xe1, 0x51, 0x94, 0x2b, 0x95, 0x3e, 0xc5,
	0xea, 0x79, 0xc9, 0x8a, 0x12, 0xf2, 0x20, 0x52, 0x2c, 0x84, 0x24, 0x81, 0xd4, 0xc8, 0x72, 0x28,
	0x81, 0x0c, 0x65, 0x2a, 0xc3, 0x55, 0x10, 0xa7, 0xc6, 0xbf, 0x69, 0xfc, 0x9a, 0x83, 0xd3, 0x10,
	0x8a, 0x04, 0x0a, 0xd1, 0xaa, 0xac, 0xfb, 0x74, 0x73, 0x83, 0xe3, 0x08, 0x22, 0xe8, 0xf8, 0xcf,
	0xab, 0xa3, 0xe3, 0x5b, 0x8c, 0x3d, 0xa8, 0xf2, 0x50, 0xf9, 0xeb, 0x4c, 0x91, 0x13, 0x4c, 0x3c,
	0x77, 0xc1, 0x4d, 0x4b, 0xf8, 0x0f, 0xf7, 0x96, 0x70, 0xb9, 0x63, 0x3b, 0x33, 0x1d, 0x91, 0x11,
	0x1e, 0xf6, 0xf9, 0xd4, 0x33, 0x85, 0xc9, 0x5d, 0xcf, 0x13, 0xe6, 0xcd, 0xc4, 0x99, 0xe9, 0xda,
	0x60, 0xfb, 0xf5, 0x9d, 0xa2, 0xf1, 0x1d, 0xde, 0xe5, 0x2a, 0x58, 0xce, 0x2b, 0x28, 0x03, 0x72,
	0x84, 0x0f, 0xb8, 0x35, 0xb9, 0x12, 0xf3, 0x85, 0xeb, 0x4f, 0xc4, 0x35, 0xb7, 0x2c, 0x1d, 0x91,
	0x43, 0xbc, 0xd7, 0x83, 0x17, 0xb6, 0xae, 0x11, 0x82, 0xf7, 0xfb, 0xe8, 0xdc, 0xd6, 0xb7, 0xba,
	0x75, 0x53, 0xe7, 0xa3, 0xa6, 0xda, 0xa6, 0xa6, 0xda, 0x57, 0x4d, 0xb5, 0xb7, 0x86, 0xa2, 0x4d,
	0x43, 0xd1, 0x67, 0x43, 0xd1, 0x23, 0x8b, 0xe2, 0x72, 0x55, 0x49, 0x23, 0x84, 0x84, 0xc9, 0x54,
	0x9e, 0xb5, 0x3d, 0x58, 0xaf, 0xdc, 0xcb, 0x5f, 0xbb, 0x72, 0x9d, 0xa9, 0x42, 0xee, 0xb4, 0xd7,
	0x5e, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x80, 0x8c, 0xca, 0xd5, 0x5e, 0x01, 0x00, 0x00,
}