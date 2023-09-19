// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmicether/ethquery/ethquery.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type EthQuery struct {
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	QueryType   string `protobuf:"bytes,2,opt,name=query_type,json=queryType,proto3" json:"query_type,omitempty"`
	Request     []byte `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	CallbackId  string `protobuf:"bytes,4,opt,name=callback_id,json=callbackId,proto3" json:"callback_id,omitempty"`
	BlockHeight uint64 `protobuf:"varint,5,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *EthQuery) Reset()         { *m = EthQuery{} }
func (m *EthQuery) String() string { return proto.CompactTextString(m) }
func (*EthQuery) ProtoMessage()    {}
func (*EthQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3a1027a4e9d5bd3, []int{0}
}
func (m *EthQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EthQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthQuery.Merge(m, src)
}
func (m *EthQuery) XXX_Size() int {
	return m.Size()
}
func (m *EthQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_EthQuery.DiscardUnknown(m)
}

var xxx_messageInfo_EthQuery proto.InternalMessageInfo

func (m *EthQuery) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EthQuery) GetQueryType() string {
	if m != nil {
		return m.QueryType
	}
	return ""
}

func (m *EthQuery) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *EthQuery) GetCallbackId() string {
	if m != nil {
		return m.CallbackId
	}
	return ""
}

func (m *EthQuery) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*EthQuery)(nil), "cosmicether.ethquery.EthQuery")
}

func init() {
	proto.RegisterFile("cosmicether/ethquery/ethquery.proto", fileDescriptor_b3a1027a4e9d5bd3)
}

var fileDescriptor_b3a1027a4e9d5bd3 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xce, 0x2f, 0xce,
	0xcd, 0x4c, 0x4e, 0x2d, 0xc9, 0x48, 0x2d, 0xd2, 0x4f, 0x2d, 0xc9, 0x28, 0x2c, 0x4d, 0x2d, 0xaa,
	0x84, 0x33, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44, 0x90, 0x14, 0xe9, 0xc1, 0xe4, 0xa4,
	0x24, 0x41, 0xa2, 0xf9, 0xc5, 0xf1, 0x60, 0x35, 0xfa, 0x10, 0x0e, 0x44, 0x83, 0x94, 0x48, 0x7a,
	0x7e, 0x7a, 0x3e, 0x44, 0x1c, 0xc4, 0x82, 0x88, 0x2a, 0x4d, 0x67, 0xe4, 0xe2, 0x70, 0x2d, 0xc9,
	0x08, 0x04, 0xe9, 0x16, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x62, 0xca, 0x4c, 0x11, 0x92, 0xe5, 0xe2, 0x02, 0x1b, 0x1b, 0x5f, 0x52, 0x59, 0x90, 0x2a, 0xc1,
	0x04, 0x16, 0xe7, 0x04, 0x8b, 0x84, 0x54, 0x16, 0xa4, 0x0a, 0x49, 0x70, 0xb1, 0x17, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0x48, 0x30, 0x2b, 0x30, 0x6a, 0xf0, 0x04, 0xc1, 0xb8, 0x42, 0xf2, 0x5c,
	0xdc, 0xc9, 0x89, 0x39, 0x39, 0x49, 0x89, 0xc9, 0xd9, 0xf1, 0x99, 0x29, 0x12, 0x2c, 0x60, 0x9d,
	0x5c, 0x30, 0x21, 0xcf, 0x14, 0x21, 0x45, 0x2e, 0x9e, 0xa4, 0x9c, 0xfc, 0xe4, 0xec, 0xf8, 0x8c,
	0xd4, 0xcc, 0xf4, 0x8c, 0x12, 0x09, 0x56, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x6e, 0xb0, 0x98, 0x07,
	0x58, 0xc8, 0xc9, 0xe7, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63,
	0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x8c, 0xd2,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x13, 0xb3, 0x12, 0xf3, 0x8a, 0x13,
	0x8b, 0x32, 0x2d, 0x4d, 0xf5, 0x91, 0x43, 0xad, 0x02, 0x11, 0x6e, 0x20, 0xb7, 0x17, 0x27, 0xb1,
	0x81, 0xbd, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x1e, 0x93, 0x3f, 0x5c, 0x01, 0x00,
	0x00,
}

func (m *EthQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintEthquery(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x28
	}
	if len(m.CallbackId) > 0 {
		i -= len(m.CallbackId)
		copy(dAtA[i:], m.CallbackId)
		i = encodeVarintEthquery(dAtA, i, uint64(len(m.CallbackId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Request) > 0 {
		i -= len(m.Request)
		copy(dAtA[i:], m.Request)
		i = encodeVarintEthquery(dAtA, i, uint64(len(m.Request)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.QueryType) > 0 {
		i -= len(m.QueryType)
		copy(dAtA[i:], m.QueryType)
		i = encodeVarintEthquery(dAtA, i, uint64(len(m.QueryType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintEthquery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEthquery(dAtA []byte, offset int, v uint64) int {
	offset -= sovEthquery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EthQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovEthquery(uint64(l))
	}
	l = len(m.QueryType)
	if l > 0 {
		n += 1 + l + sovEthquery(uint64(l))
	}
	l = len(m.Request)
	if l > 0 {
		n += 1 + l + sovEthquery(uint64(l))
	}
	l = len(m.CallbackId)
	if l > 0 {
		n += 1 + l + sovEthquery(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovEthquery(uint64(m.BlockHeight))
	}
	return n
}

func sovEthquery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEthquery(x uint64) (n int) {
	return sovEthquery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EthQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEthquery
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
			return fmt.Errorf("proto: EthQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthquery
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
				return ErrInvalidLengthEthquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthquery
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
				return ErrInvalidLengthEthquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEthquery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEthquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Request = append(m.Request[:0], dAtA[iNdEx:postIndex]...)
			if m.Request == nil {
				m.Request = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthquery
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
				return ErrInvalidLengthEthquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallbackId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEthquery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEthquery
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
func skipEthquery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEthquery
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
					return 0, ErrIntOverflowEthquery
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
					return 0, ErrIntOverflowEthquery
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
				return 0, ErrInvalidLengthEthquery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEthquery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEthquery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEthquery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEthquery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEthquery = fmt.Errorf("proto: unexpected end of group")
)
