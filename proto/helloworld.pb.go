// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/helloworld.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type SendMessageRequest struct {
	ToWho                string   `protobuf:"bytes,1,opt,name=toWho,proto3" json:"toWho,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageRequest) Reset()         { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()    {}
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_helloworld_3c6183f333534a35, []int{0}
}
func (m *SendMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageRequest.Unmarshal(m, b)
}
func (m *SendMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageRequest.Marshal(b, m, deterministic)
}
func (dst *SendMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageRequest.Merge(dst, src)
}
func (m *SendMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SendMessageRequest.Size(m)
}
func (m *SendMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageRequest proto.InternalMessageInfo

func (m *SendMessageRequest) GetToWho() string {
	if m != nil {
		return m.ToWho
	}
	return ""
}

func (m *SendMessageRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SendMessageResponse struct {
	FromWho              string   `protobuf:"bytes,1,opt,name=fromWho,proto3" json:"fromWho,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageResponse) Reset()         { *m = SendMessageResponse{} }
func (m *SendMessageResponse) String() string { return proto.CompactTextString(m) }
func (*SendMessageResponse) ProtoMessage()    {}
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_helloworld_3c6183f333534a35, []int{1}
}
func (m *SendMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageResponse.Unmarshal(m, b)
}
func (m *SendMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageResponse.Marshal(b, m, deterministic)
}
func (dst *SendMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageResponse.Merge(dst, src)
}
func (m *SendMessageResponse) XXX_Size() int {
	return xxx_messageInfo_SendMessageResponse.Size(m)
}
func (m *SendMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageResponse proto.InternalMessageInfo

func (m *SendMessageResponse) GetFromWho() string {
	if m != nil {
		return m.FromWho
	}
	return ""
}

func (m *SendMessageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SendMessageRequest)(nil), "helloworld.SendMessageRequest")
	proto.RegisterType((*SendMessageResponse)(nil), "helloworld.SendMessageResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageClient interface {
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
}

type messageClient struct {
	cc *grpc.ClientConn
}

func NewMessageClient(cc *grpc.ClientConn) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, "/helloworld.Message/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServer is the server API for Message service.
type MessageServer interface {
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
}

func RegisterMessageServer(s *grpc.Server, srv MessageServer) {
	s.RegisterService(&_Message_serviceDesc, srv)
}

func _Message_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Message/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Message_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _Message_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/helloworld.proto",
}

func init() { proto.RegisterFile("proto/helloworld.proto", fileDescriptor_helloworld_3c6183f333534a35) }

var fileDescriptor_helloworld_3c6183f333534a35 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x03, 0x0b, 0x08,
	0x71, 0x21, 0x44, 0x94, 0x5c, 0xb8, 0x84, 0x82, 0x53, 0xf3, 0x52, 0x7c, 0x53, 0x8b, 0x8b, 0x13,
	0xd3, 0x53, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0x4b, 0xf2, 0xc3,
	0x33, 0xf2, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x09, 0x2e, 0xf6, 0x5c,
	0x88, 0x3a, 0x09, 0x26, 0xb0, 0x38, 0x8c, 0xab, 0xe4, 0xc9, 0x25, 0x8c, 0x62, 0x4a, 0x71, 0x41,
	0x7e, 0x5e, 0x71, 0x2a, 0x48, 0x43, 0x5a, 0x51, 0x7e, 0x2e, 0xc2, 0x20, 0x18, 0x17, 0xb7, 0x51,
	0x46, 0xd1, 0x5c, 0xec, 0x50, 0x63, 0x84, 0x02, 0xb8, 0xb8, 0x91, 0x4c, 0x15, 0x92, 0xd3, 0x43,
	0xf2, 0x09, 0xa6, 0xa3, 0xa5, 0xe4, 0x71, 0xca, 0x43, 0x9c, 0xa3, 0xc4, 0xe0, 0xc4, 0x1e, 0xc5,
	0x0a, 0x0e, 0x82, 0x24, 0x36, 0x30, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x74, 0xc4, 0x74,
	0xd0, 0x23, 0x01, 0x00, 0x00,
}
