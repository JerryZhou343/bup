// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: usermgt/passport/passport.proto

package usermgt_passport_api_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	userbase "usermgt/userbase"
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

type HelloReq struct {
	Msg                  string             `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Info                 *userbase.UserInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HelloReq) Reset()         { *m = HelloReq{} }
func (m *HelloReq) String() string { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()    {}
func (*HelloReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6b1dee8ab7d8e21, []int{0}
}
func (m *HelloReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReq.Merge(m, src)
}
func (m *HelloReq) XXX_Size() int {
	return m.Size()
}
func (m *HelloReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReq.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReq proto.InternalMessageInfo

func (m *HelloReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *HelloReq) GetInfo() *userbase.UserInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type HelloRsp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRsp) Reset()         { *m = HelloRsp{} }
func (m *HelloRsp) String() string { return proto.CompactTextString(m) }
func (*HelloRsp) ProtoMessage()    {}
func (*HelloRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6b1dee8ab7d8e21, []int{1}
}
func (m *HelloRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRsp.Merge(m, src)
}
func (m *HelloRsp) XXX_Size() int {
	return m.Size()
}
func (m *HelloRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRsp.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRsp proto.InternalMessageInfo

func (m *HelloRsp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloReq)(nil), "usermgt.passport.api.v1.HelloReq")
	proto.RegisterType((*HelloRsp)(nil), "usermgt.passport.api.v1.HelloRsp")
}

func init() { proto.RegisterFile("usermgt/passport/passport.proto", fileDescriptor_c6b1dee8ab7d8e21) }

var fileDescriptor_c6b1dee8ab7d8e21 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x2d, 0x4e, 0x2d,
	0xca, 0x4d, 0x2f, 0xd1, 0x2f, 0x48, 0x2c, 0x2e, 0x2e, 0xc8, 0x2f, 0x42, 0x30, 0xf4, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0xc4, 0xa1, 0x0a, 0xf4, 0xe0, 0xe2, 0x89, 0x05, 0x99, 0x7a, 0x65, 0x86,
	0x52, 0x42, 0x30, 0x9d, 0x20, 0x1a, 0xa2, 0x58, 0xc9, 0x85, 0x8b, 0xc3, 0x23, 0x35, 0x27, 0x27,
	0x3f, 0x28, 0xb5, 0x50, 0x48, 0x80, 0x8b, 0x39, 0xb7, 0x38, 0x5d, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x33, 0x08, 0xc4, 0x14, 0x52, 0xe3, 0x62, 0xc9, 0xcc, 0x4b, 0xcb, 0x97, 0x60, 0x52, 0x60, 0xd4,
	0xe0, 0x36, 0x12, 0xd2, 0x03, 0x69, 0x4c, 0x4a, 0x2c, 0x4e, 0xd5, 0x0b, 0x2d, 0x4e, 0x2d, 0xf2,
	0xcc, 0x4b, 0xcb, 0x0f, 0x02, 0xcb, 0x2b, 0xc9, 0xc1, 0x4c, 0x29, 0x2e, 0x10, 0x12, 0xe2, 0x62,
	0x49, 0xce, 0x4f, 0x49, 0x05, 0x1b, 0xc3, 0x1a, 0x04, 0x66, 0x1b, 0x25, 0x70, 0xf1, 0x83, 0x74,
	0x80, 0xd4, 0x06, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0xf9, 0x72, 0xb1, 0x82, 0xb5, 0x08,
	0x29, 0xea, 0xe1, 0x70, 0xaf, 0x1e, 0xcc, 0x61, 0x52, 0x84, 0x94, 0x14, 0x17, 0x38, 0x09, 0x9c,
	0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31,
	0x24, 0xb1, 0x81, 0x3d, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x5b, 0xeb, 0xbb, 0x30,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserinfoServiceClient is the client API for UserinfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserinfoServiceClient interface {
	Hello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRsp, error)
}

type userinfoServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserinfoServiceClient(cc *grpc.ClientConn) UserinfoServiceClient {
	return &userinfoServiceClient{cc}
}

func (c *userinfoServiceClient) Hello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRsp, error) {
	out := new(HelloRsp)
	err := c.cc.Invoke(ctx, "/usermgt.passport.api.v1.UserinfoService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserinfoServiceServer is the server API for UserinfoService service.
type UserinfoServiceServer interface {
	Hello(context.Context, *HelloReq) (*HelloRsp, error)
}

// UnimplementedUserinfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserinfoServiceServer struct {
}

func (*UnimplementedUserinfoServiceServer) Hello(ctx context.Context, req *HelloReq) (*HelloRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}

func RegisterUserinfoServiceServer(s *grpc.Server, srv UserinfoServiceServer) {
	s.RegisterService(&_UserinfoService_serviceDesc, srv)
}

func _UserinfoService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserinfoServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgt.passport.api.v1.UserinfoService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserinfoServiceServer).Hello(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserinfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "usermgt.passport.api.v1.UserinfoService",
	HandlerType: (*UserinfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _UserinfoService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usermgt/passport/passport.proto",
}

func (m *HelloReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HelloReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Info != nil {
		{
			size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPassport(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintPassport(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HelloRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HelloRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Code != 0 {
		i = encodeVarintPassport(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPassport(dAtA []byte, offset int, v uint64) int {
	offset -= sovPassport(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HelloReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovPassport(uint64(l))
	}
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovPassport(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HelloRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovPassport(uint64(m.Code))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPassport(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPassport(x uint64) (n int) {
	return sovPassport(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HelloReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPassport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HelloReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPassport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPassport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPassport
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPassport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &userbase.UserInfo{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPassport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPassport
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPassport
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HelloRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPassport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HelloRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPassport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPassport
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPassport
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPassport(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPassport
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPassport
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPassport
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPassport
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPassport
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPassport
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPassport        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPassport          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPassport = fmt.Errorf("proto: unexpected end of group")
)
