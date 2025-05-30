// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aether/template/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type SayMyNameRequest struct {
}

func (m *SayMyNameRequest) Reset()         { *m = SayMyNameRequest{} }
func (m *SayMyNameRequest) String() string { return proto.CompactTextString(m) }
func (*SayMyNameRequest) ProtoMessage()    {}
func (*SayMyNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2251439876cfef6e, []int{0}
}
func (m *SayMyNameRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SayMyNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SayMyNameRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SayMyNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayMyNameRequest.Merge(m, src)
}
func (m *SayMyNameRequest) XXX_Size() int {
	return m.Size()
}
func (m *SayMyNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SayMyNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SayMyNameRequest proto.InternalMessageInfo

type SayMyNameResponse struct {
	SayMyName string `protobuf:"bytes,1,opt,name=say_my_name,json=sayMyName,proto3" json:"say_my_name,omitempty"`
}

func (m *SayMyNameResponse) Reset()         { *m = SayMyNameResponse{} }
func (m *SayMyNameResponse) String() string { return proto.CompactTextString(m) }
func (*SayMyNameResponse) ProtoMessage()    {}
func (*SayMyNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2251439876cfef6e, []int{1}
}
func (m *SayMyNameResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SayMyNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SayMyNameResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SayMyNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayMyNameResponse.Merge(m, src)
}
func (m *SayMyNameResponse) XXX_Size() int {
	return m.Size()
}
func (m *SayMyNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SayMyNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SayMyNameResponse proto.InternalMessageInfo

func (m *SayMyNameResponse) GetSayMyName() string {
	if m != nil {
		return m.SayMyName
	}
	return ""
}

func init() {
	proto.RegisterType((*SayMyNameRequest)(nil), "aether.template.v1.SayMyNameRequest")
	proto.RegisterType((*SayMyNameResponse)(nil), "aether.template.v1.SayMyNameResponse")
}

func init() { proto.RegisterFile("aether/template/v1/query.proto", fileDescriptor_2251439876cfef6e) }

var fileDescriptor_2251439876cfef6e = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x4c, 0x2d, 0xc9,
	0x48, 0x2d, 0xd2, 0x2f, 0x49, 0xcd, 0x2d, 0xc8, 0x49, 0x2c, 0x49, 0xd5, 0x2f, 0x33, 0xd4, 0x2f,
	0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x82, 0xc8, 0xeb, 0xc1,
	0xe4, 0xf5, 0xca, 0x0c, 0x95, 0x84, 0xb8, 0x04, 0x82, 0x13, 0x2b, 0x7d, 0x2b, 0xfd, 0x12, 0x73,
	0x53, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0x8c, 0xb9, 0x04, 0x91, 0xc4, 0x8a, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xe4, 0xb8, 0xb8, 0x8b, 0x13, 0x2b, 0xe3, 0x73, 0x2b, 0xe3, 0xf3,
	0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x38, 0x8b, 0x61, 0xea, 0x8c, 0x12,
	0xb9, 0x58, 0x03, 0x41, 0x76, 0x09, 0x45, 0x70, 0x71, 0xc2, 0x75, 0x0b, 0xa9, 0xe8, 0x61, 0xda,
	0xa9, 0x87, 0x6e, 0xa1, 0x94, 0x2a, 0x01, 0x55, 0x10, 0x27, 0x38, 0x79, 0x9c, 0x78, 0x24, 0xc7,
	0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c,
	0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x5e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72,
	0x7e, 0xae, 0x3e, 0xc4, 0x28, 0xdd, 0x82, 0xa2, 0xfc, 0x2c, 0x28, 0x5b, 0xbf, 0x02, 0x11, 0x24,
	0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0x00, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xe7, 0x5a, 0xde, 0x08, 0x32, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	SayMyName(ctx context.Context, in *SayMyNameRequest, opts ...grpc.CallOption) (*SayMyNameResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) SayMyName(ctx context.Context, in *SayMyNameRequest, opts ...grpc.CallOption) (*SayMyNameResponse, error) {
	out := new(SayMyNameResponse)
	err := c.cc.Invoke(ctx, "/aether.template.v1.Query/SayMyName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	SayMyName(context.Context, *SayMyNameRequest) (*SayMyNameResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) SayMyName(ctx context.Context, req *SayMyNameRequest) (*SayMyNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayMyName not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_SayMyName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayMyNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SayMyName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aether.template.v1.Query/SayMyName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SayMyName(ctx, req.(*SayMyNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var Query_serviceDesc = _Query_serviceDesc
var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "aether.template.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayMyName",
			Handler:    _Query_SayMyName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aether/template/v1/query.proto",
}

func (m *SayMyNameRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SayMyNameRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SayMyNameRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *SayMyNameResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SayMyNameResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SayMyNameResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SayMyName) > 0 {
		i -= len(m.SayMyName)
		copy(dAtA[i:], m.SayMyName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.SayMyName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SayMyNameRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *SayMyNameResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SayMyName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SayMyNameRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: SayMyNameRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SayMyNameRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SayMyNameResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: SayMyNameResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SayMyNameResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SayMyName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SayMyName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
