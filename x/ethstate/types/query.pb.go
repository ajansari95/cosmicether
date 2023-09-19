// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmicether/ethstate/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QuerySlotDataRequest struct {
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	Slot            string `protobuf:"bytes,2,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (m *QuerySlotDataRequest) Reset()         { *m = QuerySlotDataRequest{} }
func (m *QuerySlotDataRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySlotDataRequest) ProtoMessage()    {}
func (*QuerySlotDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e621100987cf2fcb, []int{0}
}
func (m *QuerySlotDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySlotDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySlotDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySlotDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySlotDataRequest.Merge(m, src)
}
func (m *QuerySlotDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySlotDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySlotDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySlotDataRequest proto.InternalMessageInfo

func (m *QuerySlotDataRequest) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *QuerySlotDataRequest) GetSlot() string {
	if m != nil {
		return m.Slot
	}
	return ""
}

type QuerySlotDataRequestResponse struct {
	SlotData *SlotData `protobuf:"bytes,1,opt,name=slot_data,json=slotData,proto3" json:"slot_data,omitempty"`
}

func (m *QuerySlotDataRequestResponse) Reset()         { *m = QuerySlotDataRequestResponse{} }
func (m *QuerySlotDataRequestResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySlotDataRequestResponse) ProtoMessage()    {}
func (*QuerySlotDataRequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e621100987cf2fcb, []int{1}
}
func (m *QuerySlotDataRequestResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySlotDataRequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySlotDataRequestResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySlotDataRequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySlotDataRequestResponse.Merge(m, src)
}
func (m *QuerySlotDataRequestResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySlotDataRequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySlotDataRequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySlotDataRequestResponse proto.InternalMessageInfo

func (m *QuerySlotDataRequestResponse) GetSlotData() *SlotData {
	if m != nil {
		return m.SlotData
	}
	return nil
}

type QueryContractDataRequest struct {
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
}

func (m *QueryContractDataRequest) Reset()         { *m = QueryContractDataRequest{} }
func (m *QueryContractDataRequest) String() string { return proto.CompactTextString(m) }
func (*QueryContractDataRequest) ProtoMessage()    {}
func (*QueryContractDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e621100987cf2fcb, []int{2}
}
func (m *QueryContractDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContractDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContractDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractDataRequest.Merge(m, src)
}
func (m *QueryContractDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryContractDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractDataRequest proto.InternalMessageInfo

func (m *QueryContractDataRequest) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

type QueryContractDataRequestResponse struct {
	Slots []SlotData `protobuf:"bytes,1,rep,name=slots,proto3" json:"slots"`
}

func (m *QueryContractDataRequestResponse) Reset()         { *m = QueryContractDataRequestResponse{} }
func (m *QueryContractDataRequestResponse) String() string { return proto.CompactTextString(m) }
func (*QueryContractDataRequestResponse) ProtoMessage()    {}
func (*QueryContractDataRequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e621100987cf2fcb, []int{3}
}
func (m *QueryContractDataRequestResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContractDataRequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractDataRequestResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContractDataRequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractDataRequestResponse.Merge(m, src)
}
func (m *QueryContractDataRequestResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryContractDataRequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractDataRequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractDataRequestResponse proto.InternalMessageInfo

func (m *QueryContractDataRequestResponse) GetSlots() []SlotData {
	if m != nil {
		return m.Slots
	}
	return nil
}

type QueryEthBlockRequest struct {
	BlockHeight uint64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *QueryEthBlockRequest) Reset()         { *m = QueryEthBlockRequest{} }
func (m *QueryEthBlockRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEthBlockRequest) ProtoMessage()    {}
func (*QueryEthBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e621100987cf2fcb, []int{4}
}
func (m *QueryEthBlockRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEthBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEthBlockRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEthBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEthBlockRequest.Merge(m, src)
}
func (m *QueryEthBlockRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEthBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEthBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEthBlockRequest proto.InternalMessageInfo

func (m *QueryEthBlockRequest) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type QueryEthBlockRequestResponse struct {
	BlockData *BlockData `protobuf:"bytes,1,opt,name=block_data,json=blockData,proto3" json:"block_data,omitempty"`
}

func (m *QueryEthBlockRequestResponse) Reset()         { *m = QueryEthBlockRequestResponse{} }
func (m *QueryEthBlockRequestResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEthBlockRequestResponse) ProtoMessage()    {}
func (*QueryEthBlockRequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e621100987cf2fcb, []int{5}
}
func (m *QueryEthBlockRequestResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEthBlockRequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEthBlockRequestResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEthBlockRequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEthBlockRequestResponse.Merge(m, src)
}
func (m *QueryEthBlockRequestResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEthBlockRequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEthBlockRequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEthBlockRequestResponse proto.InternalMessageInfo

func (m *QueryEthBlockRequestResponse) GetBlockData() *BlockData {
	if m != nil {
		return m.BlockData
	}
	return nil
}

func init() {
	proto.RegisterType((*QuerySlotDataRequest)(nil), "cosmicether.ethstate.QuerySlotDataRequest")
	proto.RegisterType((*QuerySlotDataRequestResponse)(nil), "cosmicether.ethstate.QuerySlotDataRequestResponse")
	proto.RegisterType((*QueryContractDataRequest)(nil), "cosmicether.ethstate.QueryContractDataRequest")
	proto.RegisterType((*QueryContractDataRequestResponse)(nil), "cosmicether.ethstate.QueryContractDataRequestResponse")
	proto.RegisterType((*QueryEthBlockRequest)(nil), "cosmicether.ethstate.QueryEthBlockRequest")
	proto.RegisterType((*QueryEthBlockRequestResponse)(nil), "cosmicether.ethstate.QueryEthBlockRequestResponse")
}

func init() { proto.RegisterFile("cosmicether/ethstate/query.proto", fileDescriptor_e621100987cf2fcb) }

var fileDescriptor_e621100987cf2fcb = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x8a, 0x13, 0x41,
	0x10, 0xc6, 0x33, 0x6b, 0x56, 0x92, 0xda, 0x05, 0xa5, 0xc9, 0x21, 0x84, 0x65, 0x36, 0x8e, 0x17,
	0xf5, 0x30, 0x2d, 0x11, 0x83, 0x61, 0x45, 0xd9, 0xe8, 0x82, 0x07, 0x2f, 0x8e, 0x78, 0x11, 0x61,
	0xe8, 0x4c, 0x9a, 0x99, 0xd1, 0xec, 0x74, 0x76, 0xba, 0x02, 0x2e, 0x21, 0x17, 0x9f, 0x40, 0x10,
	0x5f, 0xc0, 0xbb, 0x57, 0x9f, 0x61, 0x8f, 0x0b, 0x5e, 0x3c, 0x89, 0x24, 0x3e, 0x88, 0x74, 0xcf,
	0x9f, 0x8d, 0xa1, 0x13, 0x56, 0x6f, 0x35, 0xd5, 0xd5, 0x5f, 0xfd, 0xa6, 0xbe, 0xa2, 0xa1, 0x1d,
	0x08, 0x79, 0x1c, 0x07, 0x1c, 0x23, 0x9e, 0x52, 0x8e, 0x91, 0x44, 0x86, 0x9c, 0x9e, 0x4c, 0x78,
	0x7a, 0xea, 0x8e, 0x53, 0x81, 0x82, 0x34, 0x96, 0x2a, 0xdc, 0xa2, 0xa2, 0x75, 0xd3, 0x78, 0xaf,
	0x08, 0xb2, 0xab, 0xad, 0x46, 0x28, 0x42, 0xa1, 0x43, 0xaa, 0xa2, 0x3c, 0xbb, 0x17, 0x0a, 0x11,
	0x8e, 0x38, 0x65, 0xe3, 0x98, 0xb2, 0x24, 0x11, 0xc8, 0x30, 0x16, 0x89, 0xcc, 0x4e, 0x9d, 0x57,
	0xd0, 0x78, 0xa1, 0xba, 0xbf, 0x1c, 0x09, 0x7c, 0xca, 0x90, 0x79, 0xfc, 0x64, 0xc2, 0x25, 0x92,
	0xdb, 0x70, 0x3d, 0x10, 0x09, 0xa6, 0x2c, 0x40, 0x9f, 0x0d, 0x87, 0x29, 0x97, 0xb2, 0x69, 0xb5,
	0xad, 0x5b, 0x75, 0xef, 0x5a, 0x91, 0x3f, 0xcc, 0xd2, 0x84, 0x40, 0x55, 0x8e, 0x04, 0x36, 0xb7,
	0xf4, 0xb1, 0x8e, 0x9d, 0x37, 0xb0, 0x67, 0x92, 0xf5, 0xb8, 0x1c, 0x8b, 0x44, 0x72, 0xf2, 0x10,
	0xea, 0xaa, 0xce, 0x1f, 0x32, 0x64, 0x5a, 0x77, 0xa7, 0xb3, 0xef, 0x9a, 0xfe, 0xdc, 0x2d, 0xcb,
	0xbc, 0x9a, 0xcc, 0xc5, 0x9c, 0x23, 0x68, 0x6a, 0xf5, 0x27, 0x39, 0xc9, 0xff, 0x81, 0x3b, 0x3e,
	0xb4, 0xd7, 0xc9, 0x94, 0xa0, 0x07, 0xb0, 0xad, 0xda, 0x2a, 0x8d, 0x2b, 0x97, 0x80, 0xec, 0x57,
	0xcf, 0x7e, 0xee, 0x57, 0xbc, 0xec, 0x8e, 0xd3, 0xcb, 0x87, 0x7b, 0x84, 0x51, 0x7f, 0x24, 0x82,
	0x77, 0x05, 0xe3, 0x0d, 0xd8, 0x1d, 0xa8, 0x6f, 0x3f, 0xe2, 0x71, 0x18, 0xa1, 0xe6, 0xab, 0x7a,
	0x3b, 0x3a, 0xf7, 0x4c, 0xa7, 0x1c, 0x3f, 0x1f, 0xe0, 0xca, 0xd5, 0x92, 0xeb, 0x31, 0x40, 0x26,
	0xb1, 0x34, 0xc1, 0xb6, 0x19, 0xee, 0xa2, 0xce, 0xab, 0xeb, 0x58, 0xfd, 0x67, 0xe7, 0x73, 0x15,
	0xb6, 0x75, 0x07, 0xf2, 0xd5, 0x82, 0x5a, 0xe1, 0x13, 0xb9, 0x63, 0xd6, 0x30, 0x99, 0xd9, 0xea,
	0x5c, 0xbe, 0xb6, 0xe0, 0x76, 0x0e, 0x3f, 0x7c, 0xff, 0xfd, 0x69, 0xeb, 0x80, 0xf4, 0xa8, 0x71,
	0xa3, 0xcb, 0x41, 0xd2, 0xe9, 0xaa, 0x8b, 0x33, 0x3a, 0x55, 0xa7, 0x33, 0xf2, 0xcd, 0x82, 0xdd,
	0x65, 0xcb, 0x88, 0xbb, 0x81, 0xc3, 0xe0, 0x6d, 0xab, 0xfb, 0x6f, 0xf5, 0x25, 0xfb, 0x23, 0xcd,
	0xfe, 0x80, 0x74, 0xcd, 0xec, 0x25, 0xf0, 0x1a, 0x7e, 0xf2, 0xc5, 0x82, 0x5a, 0xe1, 0xe7, 0xc6,
	0x41, 0xaf, 0x98, 0xbe, 0x71, 0xd0, 0x6b, 0x16, 0xc4, 0xe9, 0x6a, 0xd8, 0xbb, 0xc4, 0xa5, 0xeb,
	0x9e, 0x0e, 0x5f, 0x2f, 0x03, 0x9d, 0x2e, 0xaf, 0xe2, 0xac, 0xff, 0xfc, 0x6c, 0x6e, 0x5b, 0xe7,
	0x73, 0xdb, 0xfa, 0x35, 0xb7, 0xad, 0x8f, 0x0b, 0xbb, 0x72, 0xbe, 0xb0, 0x2b, 0x3f, 0x16, 0x76,
	0xe5, 0x75, 0x27, 0x8c, 0x31, 0x9a, 0x0c, 0xdc, 0x40, 0x1c, 0x53, 0xf6, 0x96, 0x25, 0x92, 0xa5,
	0x71, 0xef, 0xfe, 0x5f, 0xf2, 0xef, 0x2f, 0x1a, 0xe0, 0xe9, 0x98, 0xcb, 0xc1, 0x55, 0xfd, 0xca,
	0xdc, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x88, 0x87, 0x7b, 0xa2, 0xf8, 0x04, 0x00, 0x00,
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
	SlotData(ctx context.Context, in *QuerySlotDataRequest, opts ...grpc.CallOption) (*QuerySlotDataRequestResponse, error)
	ContractData(ctx context.Context, in *QueryContractDataRequest, opts ...grpc.CallOption) (*QueryContractDataRequestResponse, error)
	EthBlock(ctx context.Context, in *QueryEthBlockRequest, opts ...grpc.CallOption) (*QueryEthBlockRequestResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) SlotData(ctx context.Context, in *QuerySlotDataRequest, opts ...grpc.CallOption) (*QuerySlotDataRequestResponse, error) {
	out := new(QuerySlotDataRequestResponse)
	err := c.cc.Invoke(ctx, "/cosmicether.ethstate.Query/SlotData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ContractData(ctx context.Context, in *QueryContractDataRequest, opts ...grpc.CallOption) (*QueryContractDataRequestResponse, error) {
	out := new(QueryContractDataRequestResponse)
	err := c.cc.Invoke(ctx, "/cosmicether.ethstate.Query/ContractData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EthBlock(ctx context.Context, in *QueryEthBlockRequest, opts ...grpc.CallOption) (*QueryEthBlockRequestResponse, error) {
	out := new(QueryEthBlockRequestResponse)
	err := c.cc.Invoke(ctx, "/cosmicether.ethstate.Query/EthBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	SlotData(context.Context, *QuerySlotDataRequest) (*QuerySlotDataRequestResponse, error)
	ContractData(context.Context, *QueryContractDataRequest) (*QueryContractDataRequestResponse, error)
	EthBlock(context.Context, *QueryEthBlockRequest) (*QueryEthBlockRequestResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) SlotData(ctx context.Context, req *QuerySlotDataRequest) (*QuerySlotDataRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SlotData not implemented")
}
func (*UnimplementedQueryServer) ContractData(ctx context.Context, req *QueryContractDataRequest) (*QueryContractDataRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractData not implemented")
}
func (*UnimplementedQueryServer) EthBlock(ctx context.Context, req *QueryEthBlockRequest) (*QueryEthBlockRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EthBlock not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_SlotData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySlotDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SlotData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmicether.ethstate.Query/SlotData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SlotData(ctx, req.(*QuerySlotDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ContractData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContractDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContractData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmicether.ethstate.Query/ContractData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContractData(ctx, req.(*QueryContractDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EthBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEthBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EthBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmicether.ethstate.Query/EthBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EthBlock(ctx, req.(*QueryEthBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmicether.ethstate.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SlotData",
			Handler:    _Query_SlotData_Handler,
		},
		{
			MethodName: "ContractData",
			Handler:    _Query_ContractData_Handler,
		},
		{
			MethodName: "EthBlock",
			Handler:    _Query_EthBlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmicether/ethstate/query.proto",
}

func (m *QuerySlotDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySlotDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySlotDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Slot) > 0 {
		i -= len(m.Slot)
		copy(dAtA[i:], m.Slot)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Slot)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySlotDataRequestResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySlotDataRequestResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySlotDataRequestResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SlotData != nil {
		{
			size, err := m.SlotData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryContractDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryContractDataRequestResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractDataRequestResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractDataRequestResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Slots) > 0 {
		for iNdEx := len(m.Slots) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Slots[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryEthBlockRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEthBlockRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEthBlockRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryEthBlockRequestResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEthBlockRequestResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEthBlockRequestResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockData != nil {
		{
			size, err := m.BlockData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
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
func (m *QuerySlotDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Slot)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySlotDataRequestResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SlotData != nil {
		l = m.SlotData.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryContractDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryContractDataRequestResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Slots) > 0 {
		for _, e := range m.Slots {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryEthBlockRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovQuery(uint64(m.BlockHeight))
	}
	return n
}

func (m *QueryEthBlockRequestResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockData != nil {
		l = m.BlockData.Size()
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
func (m *QuerySlotDataRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySlotDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySlotDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
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
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slot", wireType)
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
			m.Slot = string(dAtA[iNdEx:postIndex])
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
func (m *QuerySlotDataRequestResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySlotDataRequestResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySlotDataRequestResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlotData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SlotData == nil {
				m.SlotData = &SlotData{}
			}
			if err := m.SlotData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *QueryContractDataRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryContractDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
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
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
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
func (m *QueryContractDataRequestResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryContractDataRequestResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractDataRequestResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slots", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slots = append(m.Slots, SlotData{})
			if err := m.Slots[len(m.Slots)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *QueryEthBlockRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryEthBlockRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEthBlockRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
func (m *QueryEthBlockRequestResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryEthBlockRequestResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEthBlockRequestResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockData == nil {
				m.BlockData = &BlockData{}
			}
			if err := m.BlockData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
