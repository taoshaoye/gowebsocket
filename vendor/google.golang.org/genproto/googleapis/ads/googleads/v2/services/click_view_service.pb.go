// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/click_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [ClickViewService.GetClickView][google.ads.googleads.v2.services.ClickViewService.GetClickView].
type GetClickViewRequest struct {
	// The resource name of the click view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClickViewRequest) Reset()         { *m = GetClickViewRequest{} }
func (m *GetClickViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetClickViewRequest) ProtoMessage()    {}
func (*GetClickViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fade0b32fb74bb5, []int{0}
}

func (m *GetClickViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClickViewRequest.Unmarshal(m, b)
}
func (m *GetClickViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClickViewRequest.Marshal(b, m, deterministic)
}
func (m *GetClickViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClickViewRequest.Merge(m, src)
}
func (m *GetClickViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetClickViewRequest.Size(m)
}
func (m *GetClickViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClickViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetClickViewRequest proto.InternalMessageInfo

func (m *GetClickViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetClickViewRequest)(nil), "google.ads.googleads.v2.services.GetClickViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/click_view_service.proto", fileDescriptor_7fade0b32fb74bb5)
}

var fileDescriptor_7fade0b32fb74bb5 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x4b, 0xfb, 0x40,
	0x1c, 0x25, 0xf9, 0xc3, 0x1f, 0x0c, 0x15, 0x24, 0x22, 0x96, 0xe8, 0x50, 0x6a, 0x07, 0x29, 0xe5,
	0x0e, 0xa2, 0x0e, 0x9e, 0x38, 0xa4, 0x0e, 0x75, 0x92, 0x52, 0x21, 0x83, 0x04, 0x4a, 0x4c, 0x8e,
	0x70, 0x98, 0xe4, 0x6a, 0x7e, 0xd7, 0x74, 0x10, 0x17, 0xbf, 0x82, 0xdf, 0xc0, 0x4d, 0x3f, 0x4a,
	0x57, 0x47, 0x57, 0x27, 0x3f, 0x83, 0x83, 0xa4, 0x97, 0x4b, 0x5b, 0xb1, 0x74, 0x7b, 0xfc, 0x7e,
	0xef, 0xbd, 0xdf, 0xbb, 0xc7, 0x19, 0xa7, 0x11, 0xe7, 0x51, 0x4c, 0xb1, 0x1f, 0x02, 0x96, 0xb0,
	0x40, 0xb9, 0x8d, 0x81, 0x66, 0x39, 0x0b, 0x28, 0xe0, 0x20, 0x66, 0xc1, 0xdd, 0x30, 0x67, 0x74,
	0x32, 0x2c, 0x67, 0x68, 0x94, 0x71, 0xc1, 0xcd, 0x86, 0xe4, 0x23, 0x3f, 0x04, 0x54, 0x49, 0x51,
	0x6e, 0x23, 0x25, 0xb5, 0xec, 0x55, 0xe6, 0x19, 0x05, 0x3e, 0xce, 0x96, 0xdd, 0xa5, 0xab, 0xb5,
	0xaf, 0x34, 0x23, 0x86, 0xfd, 0x34, 0xe5, 0xc2, 0x17, 0x8c, 0xa7, 0x50, 0x6e, 0x77, 0x17, 0xb6,
	0x41, 0xcc, 0x68, 0x2a, 0xe4, 0xa2, 0x49, 0x8c, 0xed, 0x1e, 0x15, 0x17, 0x85, 0x9b, 0xcb, 0xe8,
	0x64, 0x40, 0xef, 0xc7, 0x14, 0x84, 0x79, 0x60, 0x6c, 0xaa, 0x5b, 0xc3, 0xd4, 0x4f, 0x68, 0x5d,
	0x6b, 0x68, 0x87, 0x1b, 0x83, 0x9a, 0x1a, 0x5e, 0xf9, 0x09, 0xb5, 0x3f, 0x34, 0x63, 0xab, 0x52,
	0x5e, 0xcb, 0xf0, 0xe6, 0xab, 0x66, 0xd4, 0x16, 0x1d, 0xcd, 0x13, 0xb4, 0xee, 0xbd, 0xe8, 0x8f,
	0x04, 0x56, 0x67, 0xa5, 0xac, 0x2a, 0x01, 0x55, 0xa2, 0xe6, 0xf1, 0xd3, 0xfb, 0xe7, 0xb3, 0x8e,
	0xcc, 0x4e, 0xd1, 0xd2, 0xc3, 0x52, 0xf4, 0xf3, 0x60, 0x0c, 0x82, 0x27, 0x34, 0x03, 0xdc, 0x96,
	0xb5, 0x15, 0x0a, 0xc0, 0xed, 0x47, 0x6b, 0x6f, 0xea, 0xd4, 0xe7, 0xd6, 0x25, 0x1a, 0x31, 0x40,
	0x01, 0x4f, 0xba, 0xdf, 0x9a, 0xd1, 0x0a, 0x78, 0xb2, 0x36, 0x7d, 0x77, 0xe7, 0x77, 0x07, 0xfd,
	0xa2, 0xd9, 0xbe, 0x76, 0x73, 0x59, 0x4a, 0x23, 0x1e, 0xfb, 0x69, 0x84, 0x78, 0x16, 0xe1, 0x88,
	0xa6, 0xb3, 0xde, 0xf1, 0xfc, 0xd8, 0xea, 0x2f, 0x74, 0xa6, 0xc0, 0x8b, 0xfe, 0xaf, 0xe7, 0x38,
	0x6f, 0x7a, 0xa3, 0x27, 0x0d, 0x9d, 0x10, 0x90, 0x84, 0x05, 0x72, 0x6d, 0x54, 0x1e, 0x86, 0xa9,
	0xa2, 0x78, 0x4e, 0x08, 0x5e, 0x45, 0xf1, 0x5c, 0xdb, 0x53, 0x94, 0x2f, 0xbd, 0x25, 0xe7, 0x84,
	0x38, 0x21, 0x10, 0x52, 0x91, 0x08, 0x71, 0x6d, 0x42, 0x14, 0xed, 0xf6, 0xff, 0x2c, 0xe7, 0xd1,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0xcd, 0x5b, 0x6b, 0xe9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClickViewServiceClient is the client API for ClickViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClickViewServiceClient interface {
	// Returns the requested click view in full detail.
	GetClickView(ctx context.Context, in *GetClickViewRequest, opts ...grpc.CallOption) (*resources.ClickView, error)
}

type clickViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewClickViewServiceClient(cc *grpc.ClientConn) ClickViewServiceClient {
	return &clickViewServiceClient{cc}
}

func (c *clickViewServiceClient) GetClickView(ctx context.Context, in *GetClickViewRequest, opts ...grpc.CallOption) (*resources.ClickView, error) {
	out := new(resources.ClickView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.ClickViewService/GetClickView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClickViewServiceServer is the server API for ClickViewService service.
type ClickViewServiceServer interface {
	// Returns the requested click view in full detail.
	GetClickView(context.Context, *GetClickViewRequest) (*resources.ClickView, error)
}

func RegisterClickViewServiceServer(s *grpc.Server, srv ClickViewServiceServer) {
	s.RegisterService(&_ClickViewService_serviceDesc, srv)
}

func _ClickViewService_GetClickView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClickViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClickViewServiceServer).GetClickView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.ClickViewService/GetClickView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClickViewServiceServer).GetClickView(ctx, req.(*GetClickViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClickViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.ClickViewService",
	HandlerType: (*ClickViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClickView",
			Handler:    _ClickViewService_GetClickView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/click_view_service.proto",
}
