// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: usermgt/code.proto

package userbase

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

type UserCode int32

const (
	UserCode_OK     UserCode = 0
	UserCode_fAILED UserCode = 1
)

var UserCode_name = map[int32]string{
	0: "OK",
	1: "fAILED",
}

var UserCode_value = map[string]int32{
	"OK":     0,
	"fAILED": 1,
}

func (x UserCode) String() string {
	return proto.EnumName(UserCode_name, int32(x))
}

func (UserCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2b37fb293fc032e9, []int{0}
}

func init() {
	proto.RegisterEnum("userbase.UserCode", UserCode_name, UserCode_value)
}

func init() { proto.RegisterFile("usermgt/code.proto", fileDescriptor_2b37fb293fc032e9) }

var fileDescriptor_2b37fb293fc032e9 = []byte{
	// 113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x2d, 0x4e, 0x2d,
	0xca, 0x4d, 0x2f, 0xd1, 0x4f, 0xce, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x00, 0x89, 0x25, 0x25, 0x16, 0xa7, 0x6a, 0xc9, 0x73, 0x71, 0x86, 0x16, 0xa7, 0x16, 0xc5, 0x83,
	0x24, 0x85, 0xd8, 0xb8, 0x98, 0xfc, 0xbd, 0x05, 0x18, 0x84, 0xb8, 0xb8, 0xd8, 0xd2, 0x1c, 0x3d,
	0x7d, 0x5c, 0x5d, 0x04, 0x18, 0x9d, 0x94, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1,
	0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0xa2, 0x04, 0x60, 0x06, 0xc2, 0x0c, 0x49, 0x62,
	0x03, 0x9b, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x63, 0x5d, 0x41, 0x8a, 0x6b, 0x00, 0x00,
	0x00,
}